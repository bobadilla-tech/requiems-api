# Consuming `openapi.json` Directly for API Docs — Gap Analysis

Research spec only. No implementation here. Answers: "are we ready to stop
hand-maintaining `apps/dashboard/config/api_docs/*.yml` and read everything
from `apps/mcp/openapi.json` instead?"

Short answer: **not yet.** Three separate blockers, in order of severity below.
None are hard — this is a real direction — but none are done either.

## Context

This follows [2026-08-03-api-docs-codegen-spec.md](2026-08-03-api-docs-codegen-spec.md),
which built `ApiDocs::SnippetGenerator` to derive `code_examples` from the
YAML's own `parameters`/`response_fields` — deliberately scoped to avoid the
bigger question of where that YAML's data should come from in the first
place. That non-goal is what this doc investigates, prompted by the fact that
other PRs are currently improving `apps/mcp/openapi.json` and there's a hope
of eventually consuming it directly and dropping `config/api_docs/*.yml`
entirely.

## Blocker 1 — coverage gap: `openapi.json` doesn't describe every endpoint

Diffed `(method, path)` across every `config/api_docs/*.yml` endpoint against
every operation in `apps/mcp/openapi.json` (129 endpoints vs 119 operations).
**10 endpoints in api_docs have no corresponding operation in openapi.json at
all**, and one has drifted to a different path:

| Endpoint | Status |
|---|---|
| `POST /v1/technology/base64/encode` | missing |
| `POST /v1/technology/base64/decode` | missing |
| `POST /v1/technology/base64/encode/batch` | missing |
| `POST /v1/technology/base64/decode/batch` | missing |
| `POST /v1/technology/barcode/batch` | missing |
| `POST /v1/technology/markdown` | missing |
| `POST /v1/systems/input/validate` | missing |
| `POST /v1/systems/input/validate/batch` | missing |
| `GET /v1/systems/domain/trust/{domain}` | missing |
| `GET /v1/technology/qr/base64` | missing |
| `GET /v1/entertainment/advice` (api_docs) vs `GET /v1/text/advice` (openapi.json) | path drift — same endpoint, different category prefix |

Whole endpoint families (all of `base64.yml`) are absent, not just individual
operations. Until `openapi.json` covers 129/129, it can't be the sole source —
it would silently drop pages from the dashboard.

## Blocker 2 — fidelity gap: what's there is present but less accurate than api_docs today

Spot-checked several operations (`sudoku`, `barcode`, `bin-lookup`,
`email-validate`) against their `openapi.json` equivalents:

- **Type accuracy is worse, not better.** `openapi.json`'s `sudoku` puzzle/
  solution fields are typed `"type": "string"` for what's actually a 9×9
  array-of-arrays (the `example` payload shows the real shape; the `schema`
  doesn't). Batch array parameters use `"items": {}` — no element type at
  all. api_docs already went through a closed-vocabulary type migration
  (`docs/plans/2026-08-03-api-docs-codegen-spec.md` step 1); openapi.json
  hasn't had an equivalent pass.
- **No machine-readable error codes.** api_docs errors carry both a `code:`
  slug (`validation_failed`, `bad_request`) and a `status:` — openapi.json
  responses only have an HTTP status key and a prose `description`. Losing
  `code` is a real regression for anything downstream that branches on it
  (this dashboard's own error tables render `code` as a column).
- **No binary-response marker.** api_docs has `response_kind: binary`
  (2 files: `barcode.yml`, `qr-code.yml`) — exactly what this migration's
  spec called out as needing an explicit flag rather than sniffing prose.
  openapi.json has no equivalent; a 200 response with no `content` block is
  the only signal, which is the same "infer from context" fragility the
  original spec explicitly moved away from.
- **No performance data.** `performance:` (p50/p95/p99 latency, sample
  count) is populated in 38/61 api_docs files and rendered in both the text
  and markdown doc views (`ApisHelper#api_documentation_as_text/_markdown`).
  No OpenAPI construct maps to this; it would need an `x-*` vendor extension
  that doesn't exist today.
- **No dashboard prose fields.** `overview.use_cases`, `overview.features`,
  and `faq` don't exist in OpenAPI at all — confirmed by grep across the
  whole spec file. Each `tags[].description` gives one sentence per API,
  nothing per-parameter or Q&A-shaped. This was already a stated non-goal
  in the original spec ("Auto-generating prose... stays hand-written") and
  stays true here regardless of which file is the source.
- **Catalog metadata lives elsewhere anyway.** `status` (live/beta/
  deprecated), `popular`, `categories`, and all of `api_catalog.yml`'s
  category-level marketing copy (icon, descriptions) has no OpenAPI
  equivalent and isn't per-endpoint data — this would stay hand-authored no
  matter what happens to `code_examples`/`parameters`/`response_fields`.

## Blocker 3 — provenance: is `openapi.json` even auto-generated?

This is the one worth confirming with whoever owns the other in-flight PRs
before investing further. `apps/mcp/scripts/fetch-spec.ts` pulls the file from
`https://api.requiems.xyz/openapi.json` at build time — so within *this*
repo, `openapi.json` is a fetched snapshot, not something generated here.
Searching `apps/api` (the Go backend that presumably serves it) for any
OpenAPI/Swagger generation code turned up nothing — no `swaggo`, no `huma`,
no annotation-based generator, no hand-rolled JSON builder referencing
"openapi" anywhere in `apps/api/**/*.go`.

That means one of two things is true, and it changes everything about
whether this migration is worth doing:

- **If** `openapi.json` is generated from Go route/struct definitions in a
  part of the stack not visible here (a service deployed from elsewhere),
  then it genuinely is closer to a single source of truth, and the
  coverage/fidelity gaps above are just backlog items on that generator.
- **If** it's hand-maintained (by a human, or by another LLM, editing JSON
  directly) rather than derived from code, then migrating api_docs's
  hand-maintenance burden onto openapi.json's hand-maintenance doesn't
  eliminate hand-editing anywhere — it just moves which file gets edited,
  with no automation gained and a worse type/error-code vocabulary today.

This repo can't answer that question by itself. It needs an answer before
Blockers 1–2 are worth fixing, since fixing them by hand (adding the missing
10 endpoints, tightening types, adding vendor extensions) is exactly the kind
of manual YAML/JSON editing this whole effort is trying to get away from —
pointless if the target file isn't actually generated either.

## What "done" looks like, if this is pursued

In rough dependency order:

1. Confirm `openapi.json`'s generation story (Blocker 3) with the team
   running those other PRs. If it's hand-maintained today with a generator
   planned, get a rough timeline before scheduling the rest of this.
2. Close the coverage gap (Blocker 1) — the 10 missing endpoints, `advice`'s
   path drift — presumably as part of whatever generates the spec, not by
   hand-patching the fetched JSON here.
3. Tighten `openapi.json`'s type vocabulary to at least parity with
   api_docs's closed vocabulary (fix `items: {}`, fix string-typed nested
   arrays), add a `code` field to error responses (or accept status-only and
   update the dashboard's error tables to drop the code column), and add an
   explicit binary-response marker (an `x-*` extension, or standardize on
   `content` being absent for binary — matching what `response_kind: binary`
   already encodes in api_docs).
4. Decide where use_cases/features/faq/performance/catalog metadata live once
   `parameters`/`response_fields`/`errors` no longer need hand-authoring in
   api_docs — likely a much smaller sidecar file per API (just the prose +
   performance + catalog fields), replacing the current all-in-one YAML
   rather than eliminating a YAML file per API entirely.
5. Point `ApisHelper#api_documentation` at openapi.json + sidecar instead of
   `config/api_docs/*.yml`, behind the same kind of proof-before-cutover
   process the code_examples migration used (`golden_diff_api_docs.rb` is a
   direct template for "diff generated-from-openapi output against the
   current hand-written YAML before trusting it").

## Recommendation

Don't start 2–5 yet. Step 1 (confirming provenance) is cheap — a conversation,
not code — and determines whether the rest of this is real automation or
busywork moved sideways. Until that's answered, `config/api_docs/*.yml`
remains the source of truth; nothing here is ready to move to a `generated/`
directory. This matches the original spec's own non-goal: "A full OpenAPI
migration... reconciling the two is a bigger, separate question outside this
spec."
