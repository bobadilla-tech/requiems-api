# Services Reorganization — Aligning Code with Dashboard Categories

The Go backend (`apps/api/services/`) historically grew organically and ended up
with 10 directories whose names did not match the 8 categories exposed to users
in the dashboard, the marketing site, or `docs/apis/`. This made it hard to find
code for a given API, created inconsistent URL prefixes (`/v1/tech/`, `/v1/ai/`,
`/v1/misc/`, etc.), and made the CLAUDE.md code-organization section misleading.

**Old directories (10):** `ai/`, `convert/`, `email/`, `entertainment/`,
`finance/`, `fitness/`, `misc/`, `places/`, `tech/`, `text/`

**New directories (8):** `entertainment/`, `finance/`, `health/`, `networking/`,
`places/`, `technology/`, `text/`, `validation/`

---

## Goals

1. One code directory per public API category, named identically to the URL
   segment it serves.
2. URL paths updated to match the new structure — no more `/v1/tech/`,
   `/v1/ai/`, `/v1/misc/`, `/v1/email/`, `/v1/fitness/`, `/v1/convert/`.
3. All downstream consumers updated: docs, dashboard Live Demo panels, Rails
   tests, integration tests, load tests, CLAUDE.md.
4. No behaviour change — all services keep identical logic; only file layout and
   URL prefixes change.

---

## Mapping

| Old directory       | Old URL prefix             | New directory       | New URL prefix            |
| ------------------- | -------------------------- | ------------------- | ------------------------- |
| `ai/`               | `/v1/ai/`                  | `text/`             | `/v1/text/`               |
| `convert/`          | `/v1/convert/`             | `technology/`       | `/v1/technology/`         |
| `email/disposable/` | `/v1/email/disposable/`    | `networking/`       | `/v1/networking/`         |
| `email/normalize/`  | `/v1/email/normalize/`     | `text/`             | `/v1/text/`               |
| `email/validate/`   | `/v1/email/validate/`      | `validation/`       | `/v1/validation/`         |
| `entertainment/`    | `/v1/entertainment/`       | `entertainment/`    | `/v1/entertainment/`      |
| `finance/`          | `/v1/finance/`             | `finance/`          | `/v1/finance/`            |
| `fitness/`          | `/v1/fitness/`             | `health/`           | `/v1/health/`             |
| `misc/convert/`     | `/v1/misc/convert/`        | `technology/units/` | `/v1/technology/convert/` |
| `misc/counter/`     | `/v1/misc/counter/`        | `technology/`       | `/v1/technology/`         |
| `misc/random_user/` | `/v1/misc/random-user/`    | `technology/`       | `/v1/technology/`         |
| `places/`           | `/v1/places/`              | `places/`           | `/v1/places/`             |
| `tech/barcode/`     | `/v1/tech/barcode/`        | `technology/`       | `/v1/technology/`         |
| `tech/domain/`      | `/v1/tech/domain/`         | `networking/`       | `/v1/networking/`         |
| `tech/ip/`          | `/v1/tech/ip/`             | `networking/`       | `/v1/networking/`         |
| `tech/mx/`          | `/v1/tech/mx/`             | `networking/`       | `/v1/networking/`         |
| `tech/password/`    | `/v1/tech/password/`       | `technology/`       | `/v1/technology/`         |
| `tech/phone/`       | `/v1/tech/validate/phone/` | `validation/`       | `/v1/validation/phone/`   |
| `tech/qr/`          | `/v1/tech/qr/`             | `technology/`       | `/v1/technology/`         |
| `tech/useragent/`   | `/v1/tech/useragent/`      | `technology/`       | `/v1/technology/`         |
| `tech/whois/`       | `/v1/tech/whois/`          | `networking/`       | `/v1/networking/`         |
| `text/advice/`      | `/v1/text/advice/`         | `entertainment/`    | `/v1/entertainment/`      |
| `text/lorem/`       | `/v1/text/lorem/`          | `text/`             | `/v1/text/`               |
| `text/profanity/`   | `/v1/text/profanity/`      | `validation/`       | `/v1/validation/`         |
| `text/quotes/`      | `/v1/text/quotes/`         | `entertainment/`    | `/v1/entertainment/`      |
| `text/spellcheck/`  | `/v1/text/spellcheck/`     | `text/`             | `/v1/text/`               |
| `text/thesaurus/`   | `/v1/text/thesaurus/`      | `text/`             | `/v1/text/`               |
| `text/words/`       | `/v1/text/words/`          | `text/`             | `/v1/text/`               |

---

## What Changed

### Go backend (`apps/api/`)

- `apps/api/app/routes_v1.go` — rewritten to mount 8 service categories instead
  of 10. Old mounts removed, new mounts added:

  ```go
  r.Mount("/entertainment", entertainment.Router(pool))
  r.Mount("/finance",       finance.Router(pool))
  r.Mount("/health",        health.Router(pool))
  r.Mount("/networking",    networking.Router(ctx, cfg))
  r.Mount("/places",        places.Router(pool))
  r.Mount("/technology",    technology.Router(ctx, pool, rdb))
  r.Mount("/text",          text.Router(pool))
  r.Mount("/validation",    validation.Router())
  ```

- New router files created: `services/health/router.go`,
  `services/networking/router.go`, `services/technology/router.go`,
  `services/validation/router.go`.
- Updated router files: `services/entertainment/router.go` (added advice,
  quotes), `services/text/router.go` (added detectlanguage, sentiment,
  similarity, normalize; removed advice, quotes, profanity).
- `package` declarations and import paths updated in all moved files.
- Two sub-path fixes:
  - `validation/email`: route changed from `/validate` → `/email` (to produce
    `/v1/validation/email` not `/v1/validation/validate`).
  - `validation/phone`: route changed from `/validate/phone` → `/phone` (same
    reason).

### Docs (`docs/`)

- All `docs/apis/**/*.md` files updated to reference new URL paths.
- `docs/core/adding-go-endpoints.md` — domain table updated to list 8
  directories; example endpoint updated.
- `CLAUDE.md` — Code Organization → Go API → `services/` section rewritten to
  describe the 8 new directories with URL prefixes.

### Rails dashboard (`apps/dashboard/`)

- `config/api_docs/*.yml` (27 files) — all URL examples updated to new paths.
  These power the Live Demo panels on the API key detail page.
- `app/views/dashboard/api_keys/show_key.html.erb` — hardcoded URL examples
  updated.
- `config/locales/en/home.en.yml`, `config/locales/es/home.es.yml` — URL
  examples updated.
- `test/controllers/admin/analytics_controller_test.rb` — endpoint fixtures
  updated.

### Tests (`tests/`)

- `tests/integration/src/suites/email.test.ts` — paths updated.
- `tests/integration/src/suites/tech.test.ts` — paths updated.
- `tests/integration/src/suites/convert.test.ts` — paths updated.
- `tests/integration/src/suites/misc.test.ts` — paths updated.
- `tests/integration/src/suites/text.test.ts` — paths updated; assertion field
  names corrected (see below).
- `tests/integration/src/suites/entertainment.test.ts` — field name corrected.
- `tests/integration/src/suites/finance.test.ts` — field name corrected.
- `tests/load/config.ts` — `SAMPLE_ENDPOINTS` updated to new paths.

### Integration test setup fix

`tests/integration/src/setup.ts` had a bug where `ROOT` resolved to `tests/`
instead of `tests/integration/`, so the `.env` file was never loaded. Fixed by
changing `"../../"` → `"../"`.

---

## Test Field-Name Corrections

While running integration tests against production, several tests were found to
assert incorrect JSON field names. These were bugs in the tests, not in the API.
Fixed alongside the URL updates:

| Endpoint                              | Wrong field            | Correct field                                               |
| ------------------------------------- | ---------------------- | ----------------------------------------------------------- |
| `GET /v1/entertainment/chuck-norris`  | `joke`                 | `fact`                                                      |
| `GET /v1/entertainment/quotes/random` | `quote`                | `text`                                                      |
| `GET /v1/text/lorem`                  | `data` as plain string | `data.text` (object with `text`, `paragraphs`, `wordCount`) |
| `POST /v1/validation/profanity`       | `is_profane`           | `has_profanity`                                             |
| `POST /v1/text/spellcheck`            | `errors`               | `corrections`                                               |
| `GET /v1/finance/crypto/{symbol}`     | `price`                | `price_usd`                                                 |

---

## Integration Test Status (2026-04-15)

Tests run against production (`https://api.requiems.xyz`). Production is still
serving the **old** URL structure — the backend changes have not been deployed
yet.

**Passing (29/51):** All tests that target unchanged URLs or use the old paths
still live on production.

| Suite                   | Status  | Notes                                                                                         |
| ----------------------- | ------- | --------------------------------------------------------------------------------------------- |
| `gateway.test.ts`       | 10/10 ✓ | `/healthz`, auth checks, rate limiting                                                        |
| `entertainment.test.ts` | 7/7 ✓   | Unchanged `/v1/entertainment/` prefix                                                         |
| `finance.test.ts`       | 6/6 ✓   | Unchanged `/v1/finance/` prefix                                                               |
| `text.test.ts`          | 6/9 ✓   | words, lorem, dictionary, thesaurus, spellcheck pass; advice/quotes/profanity 404 (new paths) |

**Failing with 404 (22/51) — pending production deployment:**

| Suite             | Failing tests | New path (not yet live)                                                                           |
| ----------------- | ------------- | ------------------------------------------------------------------------------------------------- |
| `tech.test.ts`    | 7             | `/v1/networking/*`, `/v1/technology/*`, `/v1/validation/phone`                                    |
| `convert.test.ts` | 5             | `/v1/technology/base64`, `/v1/technology/base`, `/v1/technology/color`, `/v1/technology/markdown` |
| `email.test.ts`   | 5             | `/v1/networking/disposable/*`, `/v1/validation/email`, `/v1/text/normalize`                       |
| `misc.test.ts`    | 2             | `/v1/technology/random-user`, `/v1/technology/convert`                                            |
| `text.test.ts`    | 3             | `/v1/entertainment/advice`, `/v1/entertainment/quotes/random`, `/v1/validation/profanity`         |

All 22 failures share the same root cause: the new URL prefixes do not exist on
production until the backend is redeployed. The test assertions themselves are
correct.

---

## Deployment Checklist

Before deploying, verify:

- [ ] `docker exec requiem-dev-api-1 go build ./...` — clean build
- [ ] `docker exec requiem-dev-api-1 go test ./...` — all tests pass
- [ ] Spot-check key renames manually against local dev stack:
  - `GET /v1/health/exercises` (was `/v1/fitness/exercises`)
  - `GET /v1/networking/ip` (was `/v1/tech/ip`)
  - `POST /v1/validation/phone` (was `/v1/tech/validate/phone`)
  - `POST /v1/validation/email` (was `/v1/email/validate`)
  - `POST /v1/technology/counter/test` (was `/v1/misc/counter/test`)
- [ ] Re-run `pnpm test` in `tests/integration/` against production after deploy
      — expect 51/51 pass.

---

## What Was Not Changed

- **Business logic** — no service behaviour was altered.
- **Response shapes** — all JSON responses are identical.
- **Auth Gateway** — no changes to `apps/workers/auth-gateway/`; it proxies all
  paths verbatim.
- **API Management** — no changes to `apps/workers/api-management/`.
- **`ENDPOINT_MULTIPLIERS`** in `apps/workers/shared/src/config.ts` — the only
  affected entries (`GET /v1/text/words/define`, `GET /v1/text/words/synonyms`)
  are in `text/words/` which did not move.
- **Database schema** — no migrations were added or modified.
