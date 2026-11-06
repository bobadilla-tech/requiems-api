# Release-driven sync: manual setup guide

Status: pending manual steps. Written 2026-11-06.

The tutorials, clients and skills repos were merged into `requiems-api`
(`apps/tutorials`, `apps/clients`, `apps/skills`) and the weekly cron jobs were
replaced by release-driven workflows. The code is merged; the steps below need
repository-admin or npm-owner access and cannot be done from a workflow.

## What runs when

```
publish a GitHub release
  -> release-sync.yml
       clients job: regenerate SDKs from https://requiems.xyz/openapi.json
                    -> PR "Automated: regenerate API clients" (only if changed)
       skills job:  regenerate skills from apps/dashboard/config/api_docs
                    -> PR "Automated: regenerate agent skills" (only if changed,
                       bumps package.json + .claude-plugin/plugin.json)
merge the skills PR   -> publish-skills.yml  -> npm publish @requiems/api-skills
merge the clients PR  -> publish-clients.yml -> placeholder (publishing not built yet)
```

There is no cron for skills or clients anymore (only the weekly CodeQL scan
remains). `release-sync.yml` can also be started by hand (Actions ->
Release Sync -> Run workflow).

## Step 1: create the `SYNC_PR_TOKEN` secret

PRs opened with the default `GITHUB_TOKEN` do not trigger workflows, so the
required `CI Success` check never runs on them and they cannot merge.

1. GitHub -> Settings -> Developer settings -> Personal access tokens ->
   Fine-grained tokens -> Generate new token (a GitHub App installation token
   works too).
2. Resource owner: `bobadilla-tech`. Repository access: only
   `bobadilla-tech/requiems-api`.
3. Repository permissions: **Contents: Read and write**, **Pull requests: Read
   and write**, **Metadata: Read**.
4. Repo -> Settings -> Secrets and variables -> Actions -> New repository secret
   named `SYNC_PR_TOKEN`.

With the secret set, the PRs trigger CI and `gh pr merge --auto --squash` is
enabled on them automatically. Without it, the PRs are still created but need
an admin merge.

Confirm "Allow auto-merge" is on (Settings -> General -> Pull Requests) and that
the "CI Checks" ruleset still requires `CI Success`.

## Step 2: re-point the npm Trusted Publisher

The first run of `publish-skills.yml` failed with `404 Not Found - PUT
https://registry.npmjs.org/@requiems%2fapi-skills` because the package's
trusted publisher still points at the old `requiems-api-skills` repository and
`publish.yml`.

1. Sign in to npmjs.com as an owner of `@requiems/api-skills`.
2. Package -> Settings -> Trusted Publisher. Remove the old GitHub Actions entry
   and add:
   - Organization or user: `bobadilla-tech`
   - Repository: `requiems-api`
   - Workflow filename: `publish-skills.yml`
   - Environment: leave empty
3. npm currently lists version 1.0.0 while `apps/skills/package.json` is at
   1.0.5, so earlier versions never reached npm. The first successful run will
   publish 1.0.5 (or the next bump).

## Step 3: enable the publish workflow

`publish-skills.yml` is gated so main stays green until step 2 is done.

1. Repo -> Settings -> Secrets and variables -> Actions -> Variables ->
   New repository variable: `NPM_PUBLISH_ENABLED` = `true`.
2. Actions -> Publish Skills -> Run workflow (branch `main`). Check that it
   publishes, and that re-running it says the version is already on npm.

## Step 4: dry-run the release flow

1. Actions -> Release Sync -> Run workflow (branch `main`). Nothing opens if the
   spec and docs have not changed since the last regeneration; that is correct.
2. For a real check: publish a release after the deploy has finished (the
   clients job reads the live spec at `requiems.xyz/openapi.json`, so a release
   published before the deploy completes regenerates from the old spec).
3. Expect up to two PRs. CI must run on them (needs step 1). Merge the skills
   PR and confirm `publish-skills.yml` publishes the new version.

Releases created by another workflow using `GITHUB_TOKEN` do not fire
`release: published`. Publish releases from the UI, the CLI as a user, or with
the same token as step 1.

## Step 5: tell Claude Code users about the new marketplace

The marketplace moved. Anyone who ran `/plugin marketplace add
bobadilla-tech/requiems-api-skills` must run:

```
/plugin marketplace add bobadilla-tech/requiems-api
/plugin install requiems-api-skills@requiems-api-skills
```

Optionally leave a note on the old repository's README until it is deleted.

## Step 6: delete the old repositories

Already deleted: `requiems-api-tutorials`, `requiems-api-clients`.
Remaining: `requiems-api-skills` (no open PRs or issues; full history is in
`requiems-api` under `apps/skills`). Deleting needs the `delete_repo` scope:
`gh auth refresh -h github.com -s delete_repo`, then
`gh repo delete bobadilla-tech/requiems-api-skills`. Do this only after step 2,
because npm's trusted publisher is tied to the old repo until it is re-pointed.

## Step 7 (later): implement client publishing

`publish-clients.yml` only prints a notice. When publishing is built, it fires
when `apps/clients/VERSION` changes on `main`. Registries to set up, ideally
with Trusted Publishing and no long-lived tokens: TypeScript -> npm, Python ->
PyPI, Ruby -> RubyGems, C# -> NuGet, Go -> a tagged module release.

## Checklist

- [ ] `SYNC_PR_TOKEN` secret created
- [ ] npm Trusted Publisher points at `requiems-api` / `publish-skills.yml`
- [ ] `NPM_PUBLISH_ENABLED=true` set and a manual Publish Skills run succeeded
- [ ] Release Sync run manually once; PRs (if any) got CI and merged
- [ ] Old `requiems-api-skills` repository deleted
- [ ] Client publishing implemented (tracked separately)

## Notes

- Dependabot directories for the clients were wrong after the move
  (`/apps/clients/typescript` instead of `/apps/clients/clients/typescript`,
  `/apps/example-apps/...` instead of `/apps/clients/example-apps/...`) and
  failed on every run; fixed in the same change as this guide. The generated
  client manifests are rewritten on every regeneration, so bumps there are
  short-lived.
- Skills validation (`apps/skills/scripts/validate.ts`) runs in CI and in
  Release Sync. It checks SKILL.md frontmatter and that `package.json` and
  `.claude-plugin/plugin.json` have the same version.
