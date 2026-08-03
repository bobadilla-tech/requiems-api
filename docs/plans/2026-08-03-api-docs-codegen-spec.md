# API Docs Code-Example Generator — Spec

Research spec only. No implementation here — written for another engineer to
build, validate, and test.

## Context

### Current state

Every public API has a hand-maintained YAML file at
`apps/dashboard/config/api_docs/{api_id}.yml` (61 files, 129 endpoints, 16,510
lines total). Each endpoint block carries structured metadata (`parameters`,
`response_fields`, `errors`, etc.) plus a `code_examples` map with **four
hand-written language snippets**: `curl`, `python`, `javascript`, `ruby`.

```
apps/dashboard/config/api_docs/{api_id}.yml
  → loaded raw by ApisHelper#api_documentation (YAML.load_file, no parsing/model layer)
  → rendered by app/views/partials/apis_show/_endpoint_documentation.html.erb (tabs over code_examples.each)
  → also rendered as plain text/markdown by ApisHelper#api_documentation_as_text/_markdown
  → schema-validated by test/config/api_docs_test.rb (structural checks only; doesn't touch code_examples)
```

Measured: **5,875 of 16,510 lines (35.6%) are `code_examples` blocks.** Across
129 endpoints, most `code_examples` are mechanical: build a URL, set the
`requiems-api-key` header, serialize params, call, print one field of `data`.
The four snippets are near-identical in shape, differing only in per-language
syntax — exactly the boilerplate a generator can produce from data that's
already sitting in the YAML (`method`, `path`, `parameters`, `response_fields`).

This is why they drift: changing an endpoint's path or adding a parameter means
editing prose _and_ four code blocks by hand, and nothing enforces the code
blocks stay correct — `api_docs_test.rb` never asserts a snippet's URL matches
the endpoint's `path`.

### What makes this non-trivial

Pulled directly from the 61 files, these are the real obstacles a generator must
clear — not hypothetical ones:

1. **`type` is free text, not an enum.** `api_docs_test.rb` requires `type` to
   be present but never constrains its values. Observed in the wild: `string`,
   `array`, `array of objects`, `array of strings`, `array[string]`,
   `array[array[integer]]`, `string|null`, `boolean or
   null`, `bytes`. A
   generator that branches on type to build a representative literal (JS object
   vs array vs scalar) cannot reliably parse this today.
2. **Two response shapes.** Most endpoints return the standard
   `{"data": {...}, "metadata": {...}}` JSON envelope. Three files
   (`barcode.yml`, `qr-code.yml`, `developer-utilities.yml`) have ene a blob
   URL. Nothing in the schema currently marks an endpoint as binary; it's only
   inferable from prose in `description`.
3. **Path params.** 23 endpoints dpoints that return a raw PNG — the snippets
   there don't parse JSON, they write bytes to a file / creatuse `{name}`-style
   path interpolation (`/v1/finance/bin/{bin}`, `/v1/text/dictionary/{word}`,
   …), consistent syntax, `location: path` params. Straightforward, just needs
   handling.
4. **Auth is uniform** — every single file uses the `requiems-api-key` header,
   no per-endpoint variation. One less axis to handle.
5. **14 of 61 files (23%) hand-illustrate more than one call per language**
   inside a single `code_examples` block — e.g. `base64.yml`'s curl example
   shows both the default and `variant: url` calls in one snippet, with a
   comment separating them. A 1-endpoint-to-1-snippet generator can't reproduce
   this without a structured way to express "show N example invocations."

None of these are blockers, but they mean "just template it" undersells the work
— the type taxonomy in particular has to be tightened before generated code is
trustworthy.

## Approach

### Principle: additive, not destructive

Ship the generator, prove it against every existing file, and only delete a
file's hand-written `code_examples` once its generated output has been diffed
against the original and accepted. Never let a YAML edit and a
generator-behavior edit land in the same PR — they need independently reviewable
diffs.

### 1. Tighten the `type` field (prerequisite)

Introduce a closed vocabulary the generator can switch on, and update
`api_docs_test.rb` to enforce it going forward:

```
string | integer | number | boolean | object | array<string> | array<object> | array<number>
```

Nullability becomes a separate `nullable: true` flag rather than baked into the
type string (`string|null` → `type: string, nullable: true`). This is a schema
migration across all 61 files' `parameters` and `response_fields` blocks —
mechanical, scriptable, but touches everything, so it's its own phase before the
generator is trusted.

### 2. Mark binary endpoints explicitly

Add an optional endpoint-level field: `response_kind: binary` (default, when
absent: `json`). Only the 3 affected files need it. This is a one-line addition
per binary endpoint, not a structural change.

### 3. Build the generator as a Rails service, not a Rake task

`app/services/api_docs/snippet_generator.rb` — one class per language
(`CurlSnippet`, `PythonSnippet`, `JavascriptSnippet`, `RubySnippet`), sharing
one interface: `#call(endpoint, base_url) -> String`. Each implementation:

- Builds the URL: substitutes `{param}` path segments from `location: path`
  params, appends a query string from `location: query` params.
- Builds the request body from `location: body` params, using each param's
  `example:` value when present (already authored on ~most params) and a
  type-appropriate placeholder otherwise.
- Injects the `requiems-api-key` header.
- For `response_kind: binary`, uses the binary-download template (write bytes to
  file / blob URL) instead of the JSON-parse template.
- For JSON responses, prints `response.json()["data"]` (or the
  language-idiomatic equivalent) as a whole, rather than trying to guess which
  single field is "the interesting one" — the hand-written examples cherry-pick
  a field (`data["result"]`, `data["lat"]`, …), but replicating that
  heuristically means encoding per-endpoint knowledge the generator shouldn't
  need. This is a deliberate fidelity trade for robustness; flag it to the
  implementing engineer as a call worth revisiting once the generator is live
  and reviewers can judge if it reads worse in practice.

### 4. Wire it in without touching the view layer

`ApisHelper#api_documentation` currently does a flat `YAML.load_file`. Change it
so that when an endpoint's YAML has no `code_examples` key at all, the helper
calls the generator lazily and produces the same
`{"curl" => "...", "python" => "...", ...}` hash shape the ERB partial and the
text/markdown renderers already expect. **Presence of `code_examples:` in the
YAML means "manual override, keep it verbatim."** No new YAML field needed for
the escape hatch — absence of the key _is_ the signal. This keeps
`_endpoint_documentation.html.erb` and `api_documentation_as_text/_markdown`
completely unchanged.

### 5. Prove it before deleting anything: golden-diff test

Before any YAML file loses its `code_examples` block, add a test/script that
runs the generator against all 61 files' endpoint metadata and diffs the output
against the existing hand-written snippets. Bucket the 129 endpoints into:

- **Safe to auto-generate** (single example, JSON response, no notes) — majority
  of the 129; exact count depends on how close generated output lands to
  hand-written, decide a normalized-diff threshold.
- **Needs manual override** — the 14 multi-example files, plus the 3 binary
  files until `response_kind` handling is verified, plus any batch endpoint
  whose `notes:` describe partial-success semantics the generator can't infer
  from params alone.

Only endpoints in the first bucket get their `code_examples` block deleted from
YAML in this migration. The rest keep hand-written blocks — the
`code_examples:`-key-present escape hatch means this is a permanent, not
temporary, state for endpoints where hand-authoring is cheaper than generalizing
the generator.

### 6. Rollout

Batch deletions by domain (mirrors
`docs/plans/2026-07-17-bulk-pr-review-merge-playbook.md`'s approach to
reviewable PR sizing) rather than one 61-file PR — each PR: N YAML files with
`code_examples` deleted + the golden-diff report as the PR description's
evidence that output is equivalent.

### 7. Update the "add a new endpoint" doc

`docs/core/adding-go-endpoints.md` currently doesn't mention `code_examples` at
all (confirm during implementation), but if it does, update it: new endpoints
should omit `code_examples` by default and only hand-write it when the
generator's single-example/JSON-only model doesn't fit.

## Non-goals

- **New target languages** (C#, Go, PHP, etc.). Once the generator exists,
  adding a language is "write one more `XSnippet` class" — cheap, but that's
  follow-up work, not part of this migration.
- **Auto-generating prose** — `description`, `overview.use_cases`, `faq`,
  `notes` stay hand-written. Only `code_examples` is in scope.
- **Changing the Go backend, routes, or response payloads.**
- **A full OpenAPI migration.** `apps/mcp/openapi.json` exists separately and
  isn't the source of truth for these docs today — reconciling the two is a
  bigger, separate question outside this spec.

## Open questions for the implementing engineer

- Where's the line for "safe to auto-generate" in the golden-diff step — exact
  string match, or an accepted normalized diff (e.g. ignoring variable naming)?
  This determines how big bucket 1 actually is; estimate here is unvalidated.
- Is losing the cherry-picked response field (`response.json()["data"]`
  whole-object print vs `data["result"]`) an acceptable readability trade-off?
  Worth a design review with whoever owns the docs UX before deleting real
  files.
- Should `type` tightening (step 1) land as its own standalone PR series before
  the generator work starts, given it touches all 61 files independently of
  code-example generation?

## Final notes

_(Fill in once the work lands: what shipped, what deviated from this approach,
follow-ups.)_
