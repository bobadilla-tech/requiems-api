# Move Kamal deploys off the `bobadilla-tech/kamal` fork

Status: IMPLEMENTED in #1067 (relocate scripts) and #1068 (official gem), merged
2026-11-06. Originally a proposal; review-bot corrections folded in below and
recorded under **Final notes**.

## Context

`infra/kamal/Gemfile` installs Kamal from `github.com/bobadilla-tech/kamal`
(`branch: main`, locked to `275c561` in `Gemfile.lock`). The fork is a GitHub
fork of `crydafan/kamal` and carries the gem as `kamal 2.12.0`. Goal: use the
official gem so we get upstream bug fixes and drop fork maintenance.

## Verified findings

- `gh api repos/bobadilla-tech/kamal/compare/basecamp:main...bobadilla-tech:main`:
  fork is **1 commit ahead, 39 behind** upstream `basecamp/kamal` main.
- The single patch is `Pass additional build contexts via builder.additional_contexts`
  (touches `builder/base.rb`, `configuration/builder.rb`, `validator.rb`, docs,
  tests). Nothing else differs.
- Upstream has the same feature as open PR
  [basecamp/kamal#1929](https://github.com/basecamp/kamal/pull/1929) (opened
  2026-08-18, `mergeable_state: clean`, 1 comment, last update 2026-09-13). Not
  merged, not in a release. Latest release / rubygems version is `2.12.0`.
- The only consumer of the patch is `infra/kamal/deploy.dashboard.yml`:
  `builder.additional_contexts: { scripts: scripts }`. `api` and `mcp` configs do
  not use it.
- That context feeds the `pdf-builder` stage of
  `infra/docker/dashboard.Dockerfile` (`COPY --from=scripts ...`), which runs
  `scripts/generate-docs.mjs` to produce the gitignored
  `apps/dashboard/public/documents/*.pdf`. `docker-compose.yml` passes the same
  context (`dashboard` and `worker` services) via native Compose support.
- `scripts/` also holds `gen-og-image.mjs` (manual, writes
  `apps/dashboard/public/og-image.png`) sharing `scripts/package.json`.
- PDFs are gitignored. `deploy.dashboard.yml` sets an explicit `builder.context`,
  which Kamal treats as the local checkout (not a git clone), and the dashboard
  `.dockerignore` does not exclude `public/documents`. So "generate in CI, then
  build" is technically possible, but needs Dockerfile changes plus a generate
  step in `cd.yml` and every other build workflow, and splits PDF generation
  from the image build. Not worth it versus Option A.
- Kamal has no passthrough for arbitrary `docker buildx` flags
  (`builder.output_options`, basecamp/kamal#1843, was closed), and `builder.args`
  only sets `--build-arg`. No config-only workaround exists.

## Options

| # | Option | Verdict |
|---|--------|---------|
| A | Move `scripts/` into the dashboard build context; drop `additional_contexts` | **Recommended.** Works on official Kamal, removes the need rather than patching around it. |
| B | `builder.context: .` (repo root) for dashboard + `dashboard.Dockerfile.dockerignore` | Works, but rewrites every `COPY` path, needs a big-context ignore file, changes compose too. More churn and cache risk. |
| C | Generate PDFs in CI (or commit them) and delete the `pdf-builder` stage | Feasible (see findings) but needs Dockerfile + `cd.yml` + other build-workflow changes, or binary churn in git. Rejected. |
| D | Build/publish a PDF image in CI, `COPY --from=<registry image>` | Extra pipeline, registry auth in build. Overkill for 3 PDFs. |
| E | Wait for #1929, keep fork until released | Unbounded; keeps the fork and the 39-commit lag. Fallback only. |
| F | Rebase fork on upstream main | Fixes lag but keeps a fork to maintain. Not the goal. |

## Proposal (Option A), two stacked PRs

### PR 1: relocate scripts (safe under fork *and* official Kamal)

- `git mv scripts apps/dashboard/tools` (name avoids clash with Rails `script/`).
  Includes `package.json`, `package-lock.json`, both `.mjs` files.
- Fix paths in the moved scripts:
  - `generate-docs.mjs`: `OUT_DIR` `../apps/dashboard/public/documents` -> `../public/documents`.
  - `gen-og-image.mjs`: `ROOT` becomes `resolve(__dirname, "..")` (dashboard root);
    `LOGO` -> `app/assets/images/logo.png`, `OUT` -> `public/og-image.png`.
- `dashboard.Dockerfile` `pdf-builder` stage: drop `--from=scripts`; set
  `WORKDIR /build/tools` *before* `npm ci` (a lockfile copied to `/build/tools`
  is not found by `npm ci` run in `/build`), then
  `COPY tools/package.json tools/package-lock.json ./`, `npm ci --omit=dev`,
  `COPY tools/ ./`, `node generate-docs.mjs`; final `COPY --from=pdf-builder /build/public/documents ./public/documents`.
  The stage still only hashes `tools/`, so it does not invalidate the Ruby layers
  (the `COPY . .` in `build` is after bundle install as today).
- Add `**/node_modules` to `apps/dashboard/.dockerignore` and `.gitignore`
  (the `build` stage `COPY . .` must not ship `tools/node_modules`; also
  `rm -rf tools` in the build stage's cleanup if the final image should not carry it).
- Remove `additional_contexts` from `docker-compose.yml` (`dashboard`, `worker`).
  Also remove `additional_contexts` from `deploy.dashboard.yml` in PR 1: the fork
  still passes `--build-context scripts=scripts`, and buildx versions (0.15.1,
  0.24.0) check that path even when the Dockerfile does not use it, so a clean
  checkout without `scripts/` could fail to deploy. The fork treats an absent key
  as no extra contexts, so this is safe on the fork and on official Kamal.
- `cd.yml`: replace `scripts/**` path filters (top-level and `dashboard:` filter) with nothing
  (already covered by `apps/dashboard/**`).
- Update any doc/readme mention of `scripts/` and `docs:` references
  (`grep -rn "scripts/" docs infra apps/dashboard/README.md`).

### PR 2: switch to official Kamal

- `infra/kamal/Gemfile`: `gem "kamal", "~> 2.12"` from rubygems (remove `git:` and the
  "Patched Kamal fork" comment). If an unreleased upstream fix is needed later, pin
  `git: "https://github.com/basecamp/kamal.git", ref: "<sha>"`, never a branch.
- Regenerate lock: `BUNDLE_GEMFILE=infra/kamal/Gemfile bundle lock --add-platform arm64-darwin-23 x86_64-linux ruby`
  (keeps the three existing platforms). Review diff: `GIT` block disappears, `kamal (2.12.0)` moves to `GEM`.
- `deploy.dashboard.yml`: `builder.additional_contexts` must already be gone (removed in
  PR 1); the official validator rejects unknown keys.
- `.github/actions/kamal-deploy/action.yml`: no change needed (`bundle install` + `bundle exec kamal`).
- Docs: `docs/core/deployment.md` / `infra/readme.md` should state Kamal is the
  official gem pinned via `infra/kamal/Gemfile.lock`. Update the "patched fork" wording
  in no historical artifact; the old plan stays as is.

## Validation (before merging each PR)

1. `docker build -f infra/docker/dashboard.Dockerfile apps/dashboard` succeeds and the image
   contains `public/documents/{api-auth-policy,dpa-template,security-packet}.pdf`.
2. `docker compose -f infra/docker/docker-compose.yml build dashboard worker` succeeds.
3. PR 2, with `export BUNDLE_GEMFILE=$PWD/infra/kamal/Gemfile` and dummy env vars for the
   `${...}`/secret references: run `bundle exec kamal config -c infra/kamal/deploy.<app>.yml`
   once per file (`-c` takes a single value; brace expansion does not validate all three).
   `kamal version` alone does not prove the source (fork also reports 2.12.0): check
   `bundle info kamal` and that `Gemfile.lock` has no `GIT` block.
4. Merging PR 2 touches `infra/kamal/**`, so `cd.yml` auto-deploys all three apps on push; a
   manual mcp -> dashboard -> api order cannot run first. Either watch that run and verify
   `/healthz`/dashboard health per app, or (if a staged rollout is wanted) temporarily
   disable the push trigger before merging and use `workflow_dispatch` per app.

## Risks and rollback

- Fork carries no other patches, so behavior drift risk is limited to upstream's own
  changes between our lock (`275c561`, based on 2.12.0) and the released 2.12.0 gem; the
  fork is *behind* upstream main, so main's extra fixes are not in 2.12.0 either. If a
  needed fix is only on main, use the `ref:` pin above.
- Rollback for PR 2: revert it (restores fork Gemfile and lock; the config key stays removed
  because the fork accepts its absence). PR 1 is independent and stays. A revert push also
  auto-deploys all three apps.
- Open question for the owner: if the fork has uses beyond this repo, keep it; this plan
  only removes *this* repo's dependency.
- Follow-up (optional): comment on #1929 that we no longer depend on it, or leave it for others.

## Final notes

- #1067 and #1068 merged 2026-11-06; merge of #1068 queued an automatic `cd.yml` run
  (all apps), so the staged manual rollout in Validation step 4 was not used.
- In practice `additional_contexts` was removed from `deploy.dashboard.yml` by the
  merged changes; main has no remaining reference.
- Review-bot corrections applied to this plan: explicit `builder.context` = local
  checkout (option C feasibility), `npm ci` working dir, remove stale named context in
  PR 1, per-file `kamal config`, `BUNDLE_GEMFILE` for the version/source check, and the
  CD auto-trigger on `infra/kamal/**`.
