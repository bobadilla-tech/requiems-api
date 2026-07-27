---
title: "Building AI Agent Skills as Build Artifacts"
slug: "building-ai-agent-skills-as-build-artifacts"
date: 2026-07-27
author: "Eliaz Bobadilla"
description: "How one YAML doc per endpoint became the single source of truth for our OpenAPI spec, our public docs, our AI agent skills, and five generated API clients — and the generalized pipeline pattern behind it."
---

Every API doc page on requiems.xyz has a "copy as markdown" button that dumps a
fully-formed doc a user can paste straight into their agent. That part was
never the hard part. The hard part is that the copy is a one-time paste: the
moment the underlying doc changes, every already-pasted copy silently drifts
out of date, with nothing to notice or fix it. Multiply that by every user
who's pasted a given API's docs into their own agent config, and you have a
slow, invisible staleness problem with no way to correct it in bulk.

So we stopped treating "docs for agents" as a document at all, and started
treating it as a **build artifact** — the same YAML that already powers our
public docs pages compiles into an OpenAPI spec, into installable AI-agent
skills, into an MCP server, and into five language SDKs. Nobody hand-writes
any of them. This is a walkthrough of that pipeline, using our actual
production setup as the worked example, so you can steal the pattern for your
own API.

## The source of truth

Every endpoint in Requiems API is described once, in a plain YAML file, not in
prose:

```yaml
# apps/dashboard/config/api_docs/advice.yml (trimmed)
api_id: advice
api_name: Random Advice
description: Get random pieces of advice and wisdom for inspiration, daily motivation, or content generation.
base_url: https://api.requiems.xyz
endpoints:
  - name: Get Random Advice
    method: GET
    path: /v1/entertainment/advice
    description: Returns a random piece of advice
    parameters: []
    response_example: |
      {
        "data": {
          "id": 42,
          "advice": "Don't compare yourself to others. Compare yourself to the person you were yesterday."
        }
      }
    response_fields:
      - name: id
        type: integer
        description: Unique identifier for the advice
      - name: advice
        type: string
        description: A random piece of advice
```

Sixty of these live under `apps/dashboard/config/api_docs/`, one per API, and
a Minitest suite (`api_docs_test.rb`) enforces the schema on every PR — every
endpoint needs `name`/`method`/`path`/`description`, every `path` has to start
with `/v1/`, every parameter needs a `location` (`path`, `query`, or `body` —
deliberately not OpenAPI's `in`, so this format doesn't accidentally couple
itself to OpenAPI's vocabulary). That schema check is what makes everything
downstream trustworthy: nothing consumes this YAML without knowing it's
already valid.

## The fan-out

From that one file, four independent things get generated, and none of them
know about each other:

```
apps/dashboard/config/api_docs/*.yml
        │
        ├──▶ apps/workers/auth-gateway (scripts/openapi/*.ts)
        │       └──▶ src/generated/openapi.ts ──▶ served at
        │            https://api.requiems.xyz/openapi.json
        │                  │
        │                  ├──▶ apps/mcp (fetch-spec.ts + generate.ts)
        │                  │       └──▶ generated/tools/*.ts (MCP tool wrappers)
        │                  │
        │                  └──▶ requiems-api-clients (weekly workflow)
        │                          └──▶ openapi-generator-cli ──▶ TypeScript /
        │                               C# / Python / Ruby / Go SDKs
        │
        ├──▶ requiems-api-skills (scripts/build/index.ts, reads YAML directly)
        │       └──▶ skills/<api>-<method>-<path>/SKILL.md ──▶ published to npm
        │
        └──▶ apps/dashboard/app/helpers/apis_helper.rb
                └──▶ human-facing docs page + "open in Claude/ChatGPT" link
```

The OpenAPI spec itself is never hand-edited — `auth-gateway`'s
`package.json` runs `generate:openapi` as a `predev` and `predeploy` hook, so
the spec regenerates from the YAML on every local dev boot and every
deploy. It's structurally impossible for the served `/openapi.json` to lag
behind the source docs.

Two of the four branches — the MCP server and the client SDKs — go through
that generated `/openapi.json`. The skills package doesn't: it reads the
YAML directly, on purpose, which is worth its own section.

## Skills as a build artifact, not a document

`@requiems/api-skills` is an npm package that installs `SKILL.md` files —
plain Markdown with a small YAML front-matter header — into an agent's skills
directory. That shape (front-matter + Markdown body) is the deliberate compile
target, because it's the lowest common denominator across Claude Code,
OpenCode, and GitHub Copilot's agent-skills support. No proprietary schema, no
SDK to install.

The transform is one pure function, `buildSkillMarkdown()` in
`requiems-api-skills/scripts/build/index.ts`: it takes a parsed YAML doc and
one of its endpoints, and returns a Markdown string. No file I/O, no CLI
parsing inside it — that mattered directly when this script migrated off Deno
onto Node: the markdown-building logic needed zero changes, only the I/O shell
around it did.

Running it against `advice.yml` produces:

```markdown
---
name: advice-get-advice
api: Random Advice
method: GET
path: /v1/entertainment/advice
base_url: https://api.requiems.xyz
description: Returns a random piece of advice
---

## Endpoint

**GET https://api.requiems.xyz/v1/entertainment/advice**

## Get Random Advice

Returns a random piece of advice

## Response Example

​```json
{
  "data": {
    "id": 42,
    "advice": "Don't compare yourself to others. Compare yourself to the person you were yesterday."
  }
}
​```

## Response Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| `id` | integer | Unique identifier for the advice |
| `advice` | string | A random piece of advice |
```

One file per *endpoint*, not one per API — `skills/advice-get-advice/SKILL.md`,
not `skills/advice.md`. That's deliberate: a single file covering ten
endpoints means every regeneration touches that one file, and a PR reviewer
can't tell which endpoint actually changed without reading the whole diff.
Per-endpoint files keep the diff scoped to whatever actually moved upstream.

## Regeneration, without tracking state

The weekly job (`.github/workflows/regenerate-skills.yml`) doesn't try to
figure out what changed in `requiems-api` since last time — that would mean
keeping state, and state drifts. Instead it checks out both repos fresh,
reruns the entire transform from scratch, and lets `git status --porcelain`
tell it what's different:

```yaml
- name: Checkout requiems-api (source of YAML docs)
  uses: actions/checkout@v7
  with:
    repository: bobadilla-tech/requiems-api
    path: requiems-api

- name: Regenerate skills
  run: |
    node scripts/build/index.ts \
      --source requiems-api/apps/dashboard/config/api_docs \
      --output ./skills

- name: Detect changes in skills/
  id: diff
  run: |
    if [ -n "$(git status --porcelain -- skills/)" ]; then
      echo "changed=true" >> "$GITHUB_OUTPUT"
    else
      echo "changed=false" >> "$GITHUB_OUTPUT"
    fi

- name: Bump patch version
  if: steps.diff.outputs.changed == 'true'
  run: npm version patch --no-git-tag-version
```

The version bump is gated on the diff, not on the cron tick firing. A weekly
job that always bumps the version — even on a no-op week — produces a
version-only PR every single week, forever. Gating on the actual diff means a
quiet week produces no PR at all, and a missed run isn't a lost event, either
— the next run just regenerates from scratch and produces a bigger diff.

That regeneration opens a PR. It does not publish anything. Publishing is a
second, separate, human-triggered step: pushing a `vX.Y.Z` tag fires
`publish.yml`, which cross-checks the tag against `package.json` before
touching npm at all —

```yaml
- name: Verify tag matches package.json version
  run: |
    TAG_VERSION="${GITHUB_REF_NAME#v}"
    PKG_VERSION="$(node -p "require('./package.json').version")"
    if [ "$TAG_VERSION" != "$PKG_VERSION" ]; then
      echo "::error::Tag v$TAG_VERSION does not match package.json version $PKG_VERSION"
      exit 1
    fi

- name: Publish to npm
  run: npm publish --access public --provenance
```

— and publishes using npm's Trusted Publishing (OIDC), so there's no
long-lived npm token sitting in repo secrets. "This reflects current docs"
(merge) and "this is now public" (tag push) are two decisions a maintainer
makes at different times, on purpose. The tempting version of this pipeline
auto-publishes whenever `main` changes; we didn't build that, because it means
nothing reaches the public registry that a human didn't explicitly decide to
ship — which matters a lot more once real users depend on the package.

## Design decisions worth stealing

- **One generated file per endpoint.** Keeps regeneration diffs reviewable.
- **Gate the version bump on the diff, not the schedule.** No-op weeks produce
  no noise.
- **Decouple merged from published.** A docs update should never silently
  become a public release.
- **Guard the publish step by cross-checking tag against package version.**
  Three lines of bash beats discovering the mismatch as a cryptic npm 403.
- **Keep the toolchain boring.** This script briefly ran on Deno for its nicer
  `--allow-read`/`--allow-write` permission model, while the rest of the repo
  was Node. In practice the script never touched anything untrusted, so the
  sandboxing bought nothing — it just meant two lockfiles and an extra
  `setup-deno` step to run one file. Consolidating onto Node's native
  TypeScript stripping (no `ts-node`, no bundler) removed a whole toolchain for
  free.

## Applying this to your own API

1. Make sure your API docs are structured data somewhere — YAML, JSON, a
   database table — not prose. If they're a wiki page today, write that
   structured layer first; there's no reliably transforming free text.
2. Write one pure function: doc in, target format out (Markdown for a skill,
   an OpenAPI operation object, whatever your target is).
3. Wrap it in a CLI you can run on your own laptop and get the same result CI
   gets.
4. Schedule a job that reruns the transform from scratch and diffs the output,
   rather than tracking what changed upstream.
5. Keep publishing a separate, explicit action from merging.
6. Mark every generated file `// AUTO-GENERATED — do not edit`, and pick one
   boring runtime to generate it with.

## Where it's still rough

Two honest gaps, since this is a production system and not a case study: the
`requiems-api-clients` pipeline regenerates all five language SDKs weekly and
opens a PR, but there's no publish step wired up yet — nothing pushes to npm,
PyPI, NuGet, or RubyGems automatically, unlike the skills package. And because
`openapi-generator-cli`'s `packageName`/`gemName` options were never set for
every language, a couple of the generated clients still carry generator
boilerplate names (`Org.OpenAPITools` for C#, `openapi_client` as the Ruby gem
name) instead of anything Requiems-branded.

Both are next on the list — the pattern proved itself on the skills side
first, so that's the one we're writing about today.
