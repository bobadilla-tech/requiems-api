# Requiems API v2 — Deployment Playbook

Small, running checklist of things not to forget when actually deploying
the Go-auth-foundation rewrite (PR #966) to production. Not a full
deployment guide — see `docs/core/deployment.md` for that. This is just
the "don't forget X" list accumulated while building v2.

- [ ] **Provision `PLAYGROUND_API_KEY`.** As of
      `docs/plans/2026-08-22-go-auth-foundation-phase-6-usage-multiplier-and-loose-ends.md`
      Phase 6b item 3, production's `PLAYGROUND_API_KEY` was not confirmed —
      check whether it's still `app_config.rb`'s
      `requiem_notprovisioned0000000000` default. If so, create a real,
      dedicated `ApiKey` row via the normal `ApiKeyGenerator` path on a
      bounded (not `enterprise`) plan, and set the production env var.
      Until this is done, the public playground/demo forms
      (`ApiProxyService`) are non-functional.
