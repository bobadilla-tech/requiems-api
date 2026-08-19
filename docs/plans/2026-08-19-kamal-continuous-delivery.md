# Kamal Continuous Delivery — GitHub Actions

Automate the manual `kamal deploy -c infra/kamal/deploy.<app>.yml` flow with a
GitHub Actions CD workflow that builds, pushes to GHCR, and deploys to the VPS on
every push to `main`, path-filtered so only the changed apps redeploy.

## Context

Kamal deployment support was added recently (the three `infra/kamal/deploy.*.yml`
configs plus a shared secrets file), but deploys were still run by hand from a
developer machine: source the gitignored root `.env` into the shell, then run
`kamal deploy -c infra/kamal/deploy.<app>.yml`. Nothing in CI built or pushed
images, and nothing deployed.

The Kamal configs are YAML rendered as ERB, so host and registry values already
come from the environment — `servers.web` uses `ENV["HETZNER_VPS_IP"]`, and
`registry.password` references the `KAMAL_REGISTRY_PASSWORD` secret. Secrets
live in `infra/kamal/secrets`, which is committed but only contains `KEY=$KEY`
indirection: Kamal parses it with dotenv and resolves each value from the
process environment. So the same "export the real values as env vars" pattern
that powers manual deploys is exactly what a CI runner must do — no new secret
storage mechanism is needed, just GitHub environment secrets mapped to env vars.

The existing CI workflow (`ci.yml`) already scopes per-app jobs with
`dorny/paths-filter` and runs on `ubuntu-latest`; the CD workflow mirrors that
pattern.

Two things needed deciding before writing it:

1. **Registry auth.** The deploy configs hardcoded `username: bobadilla-tech`
   with a `KAMAL_REGISTRY_PASSWORD` secret. `GITHUB_TOKEN` was preferred for CI
   (no long-lived credential to rotate), but it authenticates as the triggering
   actor, not the org name. Package visibility is unaffected by the token
   choice, so `GITHUB_TOKEN` was adopted, with the registry username made
   overridable via a `KAMAL_REGISTRY_USERNAME` env var that defaults to
   `bobadilla-tech` so manual deploys keep working unchanged.
2. **Where secrets live.** A GitHub `production` environment, so deploy secrets
   are scoped there and can gain required-reviewer protection later.

## Approach

**Registry override (three one-line edits).** Each `infra/kamal/deploy.*.yml`
changed `username: bobadilla-tech` to
`username: <%= ENV["KAMAL_REGISTRY_USERNAME"] || "bobadilla-tech" %>`. This is
the only config change — the password already flows through the secrets file, and
CI just sets `KAMAL_REGISTRY_USERNAME=${{ github.actor }}` and
`KAMAL_REGISTRY_PASSWORD=${{ github.token }}`.

**New `.github/workflows/cd.yml`.** A single workflow with:

- **Triggers:** `push` to `main` with a broad `paths` filter, plus
  `workflow_dispatch` with an `app` choice (`all`/`api`/`dashboard`/`mcp`) for
  manual redeploys and rollbacks.
- **Permissions:** `contents: read` and `packages: write` (for the GHCR push).
- **Concurrency:** a single `cd-deploy` group with `cancel-in-progress: false`,
  so deployments serialize — all three apps share one VPS and the `db`/`redis`
  accessories.
- **`detect-changes` job** (push only) using `dorny/paths-filter` to decide which
  apps changed. The dashboard filter also watches `scripts/**` and
  `infra/docker/dashboard.Dockerfile` because its image consumes the `scripts`
  build context; all three watch `infra/kamal/**` and the workflow file itself.
- **Three deploy jobs** (`deploy-api`, `deploy-dashboard`, `deploy-mcp`), each
  `environment: production`, gated on the filter output (or the dispatch input).
  Each job sets its own `env` (secrets mapped to env vars) and then delegates to
  a shared composite action (`.github/actions/kamal-deploy`) that runs the common
  steps: `ruby/setup-ruby` → configure bundler (`BUNDLE_GEMFILE` points at
  `infra/kamal/Gemfile`, which installs the patched Kamal fork from
  `github.com/bobadilla-tech/kamal`) → `bundle install` → configure SSH (write
  the private key to `~/.ssh/id_ed25519`, add the host via `ssh-keyscan`) → a
  guard that fails fast if any required secret is empty → `bundle exec kamal
  deploy -c infra/kamal/deploy.<app>.yml`. The composite action avoids repeating
  those steps across the three jobs and keeps the SSH/registry plumbing in one
  place. The fork is pinned to a commit via the committed `Gemfile.lock`.

The per-app env mapping keeps the surface minimal: `api` and `dashboard` get the
Postgres, `DATABASE_URL`/`REDIS_URL`, and `BACKEND_SECRET` values; `dashboard`
additionally gets `SECRET_KEY_BASE`, `API_MANAGEMENT_API_KEY`,
`LEMONSQUEEZY_SIGNING_SECRET`, and `SMTP_PASSWORD`; `mcp` only needs the VPS IP
and registry values. The guard step exists because dotenv silently leaves a
`$VAR` reference literal when the corresponding env var is missing — an explicit
empty-check prevents deploying a malformed connection string.

**Docs.** `docs/core/deployment.md` gained a "Continuous Delivery (GitHub
Actions)" section documenting the workflow, the required `production`
environment secrets, and the manual redeploy/rollback path. `infra/readme.md`
got a pointer to it.

## Final notes

- Shipped: registry username override in the three deploy configs,
  `.github/workflows/cd.yml`, a shared `.github/actions/kamal-deploy` composite
  action (extracted after review to avoid repeating the SSH/Ruby/Kamal steps
  across the three jobs), and the two doc updates above.
- Kamal is installed from the patched fork (`bobadilla-tech/kamal`) via
  `infra/kamal/Gemfile` + committed `Gemfile.lock` (pinned to the fork's `main`
  at the time of writing). To deploy a newer patch, update the branch/ref in the
  Gemfile and re-run `BUNDLE_GEMFILE=infra/kamal/Gemfile bundle lock`.
- Manual setup still required (not something code can do): create the
  `production` GitHub environment and add the secrets
  (`HETZNER_VPS_IP`, `SSH_PRIVATE_KEY`, `POSTGRES_USER`, `POSTGRES_PASSWORD`,
  `POSTGRES_DB`, `DATABASE_URL`, `REDIS_URL`, `BACKEND_SECRET`,
  `SECRET_KEY_BASE`, `API_MANAGEMENT_API_KEY`, `LEMONSQUEEZY_SIGNING_SECRET`,
  `SMTP_PASSWORD`), then verify the first run.
- Known limitation: `GITHUB_TOKEN` is ephemeral, so out-of-band `kamal rollback`
  on the VPS after a run can't pull images. Rollbacks should be triggered via
  `workflow_dispatch`, or a `KAMAL_REGISTRY_PASSWORD` PAT added for server-side
  pulls.
