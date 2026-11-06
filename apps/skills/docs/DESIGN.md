# Requiems API Skills Spec

Installable skill packages for
[Requiems API](https://github.com/bobadilla-tech/requiems-api), enabling AI
agents to interact with Requiems API endpoints without manual copy-pasting of
documentation.

This document is the living spec: it describes what the pipeline actually does
today, not what was originally proposed. See the
[Appendix](#appendix-original-poc-plan) for the historical POC plan this was
built from.

## Overview

Requiems API exposes its endpoint documentation as structured YAML files
(`apps/dashboard/config/api_docs/*.yml`). This repository packages that
documentation as installable skill bundles so that AI agents (Claude, Cursor,
Copilot, etc.) can discover, understand, and call Requiems API endpoints
natively.

Instead of copying documentation by hand, users install this package once and
their agent immediately gains full context on every available endpoint.

```bash
npm install @requiems/api-skills
```

## Problem

Getting Requiems API's docs into an agent's context isn't the hard part — every
API doc page on requiems.xyz already has a "copy as markdown" button that dumps
a fully-formed doc a user can paste straight into their agent. The actual
problem is staleness: that copy is a one-time paste. The moment the underlying
`api_docs/*.yml` changes, every already-pasted copy silently drifts out of date,
with nothing to notice or fix it. Multiply that by every user who's pasted a
given API's docs into their own agent config, and it's a slow, invisible drift
with no way to correct it in bulk.

This pipeline turns that manual, one-time paste into a versioned npm dependency:
install once, and pick up current docs through a normal `npm update` instead of
remembering to re-copy from the site.

## Why We Believe This Is Worth Solving

AI-assisted development is a standard part of developer workflows. As Requiems
API grows in endpoint count and user base, the gap between "API exists" and
"agent still has last month's copy-pasted docs" widens.

Specifically:

- Every change to `requiems-api` requires every user who's already copy-pasted
  those docs to notice and manually re-paste them. Nothing prompts them to; most
  won't. This does not scale.
- Stale or inconsistent agent context leads to incorrect API calls, which
  generates avoidable support load.
- Competing API platforms (e.g., Stripe, Twilio) already ship AI-ready context
  packages. Not having one puts Requiems API at a disadvantage in developer
  experience.
- The documentation infrastructure already exists in `api_docs/`. The marginal
  cost of packaging it properly is low compared to the adoption benefit.

## Approaches Considered

### Option A — Direct Git submodule (discarded)

Users add `requiems-api` as a git submodule and point their agent at the
`api_docs/` folder directly. Simple to set up, but requires users to manage
submodule updates manually and exposes them to the full repo rather than a
clean, minimal package.

### Option B — Static file download via CDN (discarded)

Publish the skill files as static assets on a CDN (e.g., jsDelivr over GitHub
raw). No packaging step needed. Downside: no versioning, no dependency
management, no way to pin a specific version. If a file moves or is renamed, all
users break silently.

### Option C — Per-skill individual packages (deferred)

Publish one npm package per endpoint (e.g., `@requiems/skill-sentiment`,
`@requiems/skill-sudoku`). Maximally granular, but multiplies maintenance
overhead significantly. Suitable as a future improvement once the single-package
approach is proven — see [No granular versioning](#known-limitations-and-risks).

### Option D — npm package with automated sync pipeline (chosen, implemented)

Package all skills into a single installable npm package. A GitHub Action
regenerates the content whenever a `requiems-api` release is published and opens
a PR if anything changed; a maintainer reviews and merges it; merging a version
bump publishes to npm. This is what's
implemented today — see [Current Architecture](#current-architecture).

## Current Architecture

The package lives in the `requiems-api` monorepo (`apps/skills`) next to the
YAML it is generated from. The pipeline has three stages, driven by releases
rather than a cron:

1. **Regenerate** (on release) — when a GitHub release is published,
   `.github/workflows/release-sync.yml` regenerates the skills from
   `apps/dashboard/config/api_docs/` and opens a PR (with a patch version bump)
   only if anything changed.
2. **Review & merge** (human, or auto-merge when `SYNC_PR_TOKEN` is set) — a
   reviewer spot-checks the PR, which links the live doc page for each API that
   changed.
3. **Publish** (on merge) — a merged change to `apps/skills/package.json` runs
   `.github/workflows/publish-skills.yml`, which publishes to npm if that
   version is not already published.

## Repository Structure

```
apps/skills/   (workflows live in the repo root `.github/`:
               release-sync.yml, publish-skills.yml, scripts/build-skills-pr-body.sh)
├── scripts/
│   └── build/
│       ├── index.ts               # transform pipeline (YAML -> SKILL.md)
│       └── types.ts               # ApiDoc / Endpoint / Parameter shapes
├── skills/
│   └── <api_id>-<method>-<last-path-segment>/
│       └── SKILL.md               # one generated skill per endpoint
├── docs/
│   ├── DESIGN.md                  # this file
│   ├── DEVELOPERS.md              # local regen quick-start
│   └── HOW_TO_BUILD_AGENT_SKILLS.md  # the generalized pattern, write-up form
├── .nvmrc                         # pins the Node version for CI and local dev
├── package.json
├── LICENSE
└── README.md
```

## Toolchain

The transform script (`scripts/build/`) runs directly on **Node.js (≥26, see
`.nvmrc`)** as plain TypeScript — no `ts-node`, no bundler, no build step.
Node's native type-stripping erases the `interface`/`as` syntax at load time, so
`node scripts/build/index.ts` just works. YAML parsing uses `js-yaml` (a
`devDependency`, since it's build tooling, not part of the published package).

This wasn't always the case: the script originally ran on Deno while everything
else in the repo (`package.json`, publishing) was Node — two runtimes, two
lockfiles, and a `denoland/setup-deno` step in CI just to run one file. The
transform itself never needed anything Deno-specific (no untrusted input, no
need for its permission sandboxing), so the split bought nothing but maintenance
overhead. Consolidating onto Node removed a whole toolchain. See
[HOW_TO_BUILD_AGENT_SKILLS.md](HOW_TO_BUILD_AGENT_SKILLS.md#keep-the-toolchain-boring)
for the general version of this lesson.

## Skill Format

Each endpoint becomes its own directory under `skills/`, named
`<api_id>-<method>-<last-path-segment>`, containing a single `SKILL.md`.
Front-matter carries the machine-readable fields; the body is natural-language
plus conditionally-rendered sections (Parameters, Request Example, Response
Example, Response Fields, Errors — only emitted if the source YAML has that
data). Real example, generated as-is from `api_docs/advice.yml`:

````markdown
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

```json
{
  "data": {
    "id": 42,
    "advice": "Don't compare yourself to others. Compare yourself to the person you were yesterday."
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Response Fields

| Field    | Type    | Description                      |
| -------- | ------- | -------------------------------- |
| `id`     | integer | Unique identifier for the advice |
| `advice` | string  | A random piece of advice         |
````

One file per endpoint (not one per API) keeps each skill small enough for an
agent to load on demand, and keeps regen PR diffs scoped to exactly the
endpoints that changed.

## Automation

### release-sync.yml (`skills` job)

Triggers: a published GitHub release, and manual `workflow_dispatch`. Also
regenerates the API clients in a sibling job.

1. Checks out `main` (the YAML source is in the same repo).
2. Installs Node from `apps/skills/.nvmrc`, runs `npm ci`.
3. Runs `node scripts/build/index.ts --source ../dashboard/config/api_docs
   --output ./skills` from `apps/skills`.
4. Checks `git status --porcelain -- apps/skills/skills/`. If nothing changed,
   the job ends: no PR, no version bump.
5. Otherwise bumps the patch version (`npm version patch --no-git-tag-version`).
6. Builds the PR body via `.github/scripts/build-skills-pr-body.sh`, linking the
   `requiems.xyz/en/apis/<api_id>` page of each changed API.
7. Opens or updates a PR from `chore/automated-skill-regeneration` into `main`.
   Set the `SYNC_PR_TOKEN` secret so CI runs on that PR and auto-merge is
   enabled; PRs created with the default token do not trigger CI.

### publish-skills.yml

Trigger: push to `main` touching `apps/skills/package.json` (and manual
dispatch). Checks whether `name@version` already exists on npm, and if not runs
`npm publish --access public --provenance` using npm Trusted Publishing (OIDC).
The trusted publisher on npmjs.com must reference this repository and workflow
file.

## Known Limitations and Risks

**Sync lag**: Skills are regenerated when a release is published. If an endpoint
is added or changed between releases, the package won't reflect it until the
next release (or a manual `workflow_dispatch` run of `release-sync.yml`).

**Format coupling**: The transform assumes the current shape of `api_docs/` YAML
files. A schema change in `requiems-api` breaks the sync until
`scripts/build/index.ts` is updated.

**Manual release step is easy to forget**: Merging a regen PR only updates
`main` — nothing gets published until a maintainer pushes a matching tag. If
that step is skipped, the published package silently drifts behind what's
actually in `main`.

**Review ping isn't an enforced gate**: `@bobadilla-tech/requiems-api` is
mentioned directly in the PR body text (not requested as a formal reviewer via
the GitHub API), because the default `GITHUB_TOKEN` used by the workflow usually
lacks the org-scope permission team review requests need. This notifies the team
but doesn't block merging without an approval.

**Version bump only tracks `skills/`**: The patch-bump gate diffs `skills/`
specifically. A change to `scripts/build/index.ts` itself (e.g. an output format
change) won't bump the version on its own — only a resulting change in generated
output will.

**CODEOWNERS / review team naming**: `.github/CODEOWNERS` lists
`@bobadilla-tech/requiems-api-core-team`, while the automated regen PR pings
`@bobadilla-tech/requiems-api` — two different team slugs. Worth confirming both
teams exist, are visible to this repo, and are the intended reviewers for their
respective purpose.

**npm registry dependency**: Publishing depends on an org-level npm token. Token
rotation or org changes require updating the `NPM_TOKEN` secret in this
repository's settings.

**Agent compatibility**: The `SKILL.md` convention is broad enough to work with
Claude Code, OpenCode, and GitHub Copilot (see the README for per-agent install
steps), but each still expects the file in a different directory. There's no
per-agent adapter beyond documenting those paths.

**No granular versioning**: All skills ship as a single package. If only one
endpoint changes, users re-download everything. A per-skill package approach
(Option C, above) is a possible future improvement.

## Appendix: Original POC Plan

This is the original proof-of-concept plan the automation pipeline above was
built from. Kept for the record — every step below was completed and the success
criteria were met before the CI automation was added.

Before building the full pipeline, a small proof of concept validated the core
assumption: that the existing `api_docs/` YAML files could be automatically
transformed into skill bundles and packaged for local installation.

### Scope

The POC covered only the transformation and local packaging step — no GitHub
Actions, no npm publishing, no CI/CD. The goal was to answer one question: _does
the data we have today produce a usable skill package?_

### Steps

**Step 1 — Pick two or three representative endpoints.** Endpoints with
different characteristics from `api_docs/` — simple parameters and a flat
response (e.g., `sudoku`), a richer nested response structure (e.g.,
`exercises`), and an external API dependency (e.g., `spellcheck`, which calls
the LanguageTool API). This tested that the transformer handles different levels
of complexity, not just the happy path.

**Step 2 — Write the transform script.** A minimal script that:

- Reads a `.yml` file from `api_docs/`
- Extracts the relevant fields (name, method, path, description, parameters,
  errors, examples)
- Writes a `.md` file in the skill format defined above

No GitHub Actions, no automation. Just:
`node scripts/build/index.ts --source ./sample-docs --output ./skills`

**Step 3 — Set up `package.json`.** Minimal `package.json` including only the
`skills/` directory in the published files. Verified locally with `npm
pack` —
this generates a `.tgz` file showing exactly what a user would download.

**Step 4 — Install locally and verify.**

```bash
npm pack
npm install ./requiems-api-skills-0.0.1.tgz
```

Opened the installed files and confirmed a real agent tool (Claude Code) could
load and use a skill correctly.

### Success criteria (all met)

- The transform ran without errors on the three selected endpoints.
- The generated `.md` files matched valid skill format (correct front-matter,
  readable body).
- `npm pack` produced a `.tgz` with only the `skills/` folder inside.
- At least one generated skill loaded in Claude Code without manual editing.

### Out of scope for the POC (addressed later, above)

- Automated sync with `requiems-api`
- Publishing to the npm registry
- Handling every edge case in `api_docs/` structure
- Tests
