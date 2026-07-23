# SEO Tool Pages — Barcode, Advice, Base64, WHOIS, Lorem Ipsum, Working Days

Shipped 6 new "SEO demo tool" pages on the Rails dashboard (`apps/dashboard`) —
live-demo form → real API call → rendered result, plus keyword-rich copy and
CTAs — following the exact pattern already proven by the `mortgage`, `markdown`,
and `unit-conversion` tool pages. Localized in EN, ES, and FR from the start.

## Context

Growth filed a batch of tickets: "[Growth][SEO tool] Live demo + landing" for
`barcode`, `advice`, `base64`, `lorem-ipsum`, `unit-conversion`, `whois`, plus
"[Growth][SEO tool][ES] Localize demo + copy" for `advice`, `base64`,
`lorem-ipsum`, `markdown`, `mortgage`, `unit-conversion`, `working-days`,
`whois`. The raw ticket text was jumbled (titles and bodies mismatched across
copy-paste), so the real scope had to be verified against the codebase rather
than taken at face value.

Verification (two Explore passes + direct file reads of
`apps/dashboard/app/controllers/tools_controller.rb`,
`tool_demos_controller.rb`, `config/routes.rb`, and the locale files) found:

- `mortgage`, `markdown`, `unit-conversion` already had complete, high-quality
  **EN + ES + FR** pages and translations in production — nothing to build.
  One ES typo found in passing: `tools.unit_conversion.hero.meters` was
  `"Medidores"` (should be `"Metros"`).
- `barcode`, `advice`, `base64`, `whois`, `lorem-ipsum`, `working-days` had
  **no page at all** — no controller action, route, view, or locale keys in
  any language. The "[ES] Localize" tickets for `lorem-ipsum` and
  `working-days` were mis-scoped: there was no EN page to translate from, so
  they needed full builds like the other four, not a translation pass.
- `/tools/demos/*` POST routes were not covered by Rack::Attack (only
  `/api/proxy` was), despite each new tool page's acceptance criteria calling
  for "rate-limit messaging."

The app's `available_locales` is `[:en, :es, :fr]`
(`config/application.rb`), and `config/locales/fr/tools.fr.yml` already
mirrored the ES file's completeness for the 3 existing tools — so all 6 new
tools were built with EN/ES/FR from the start, per user decision, rather than
EN-only with ES/FR as follow-ups.

## Approach

Every new tool follows the same file set, copied from the `mortgage` /
`markdown` / `number_base_conversion` / `mx_lookup` / `qr_code` /
`random_user` tools as templates:

- `ToolsController::SUPPORTED_TOOLS` / `TOOLS_METADATA` / `CATEGORIES` — id,
  name, description, icon classes, category bucket.
- `ToolDemosController#<tool>` — validates params, calls
  `ApiProxyService.call` via the existing private `api_call` helper, branches
  on `429` / `422` / `404` / non-200 / nil-data, renders a Turbo Frame result
  partial or the shared `demo_error` partial.
- `config/routes.rb` — one `post "tools/demos/<slug>"` line per tool.
- `app/views/tools/<tool>/show.html.erb` + 6 partials
  (`_hero`, `_what_it_does`, `_use_cases`, `_api_combinations`, `_faq`, `_cta`)
  under `app/views/partials/tools/<tool>/`.
- `app/views/tool_demos/<tool>.html.erb` — Turbo Frame result partial.
- `app/javascript/controllers/<tool>_demo_controller.js` — Stimulus, handles
  `turbo:submit-start`/`turbo:submit-end`, client-side validation, spinner.
- `tools.<tool>.*` locale keys in `config/locales/{en,es,fr}/tools.{en,es,fr}.yml`.
- Tests in `tools_controller_test.rb` (GET show page) and
  `tool_demos_controller_test.rb` (POST demo — success / empty / invalid /
  429 / non-200 / nil-data, plus tool-specific cases).

Per-tool specifics (endpoint shapes confirmed against
`config/api_docs/*.yml` and cross-checked with the Go backend source):

- **barcode** — `GET /v1/technology/barcode/base64`, `data` + `type` (enum of
  5 formats). Renders the returned base64 PNG inline, same pattern as
  `qr_code`.
- **advice** — `GET /v1/entertainment/advice`, no params, no-input
  click-to-fetch form (`random_user` pattern). See note below on the initial
  wrong endpoint.
- **base64** — the one tool needing a mode toggle (`encode`/`decode`) plus an
  optional `variant` (`standard`/`url`); single controller action branches on
  `params[:mode]`, surfaces `422` as a distinct "invalid Base64" message.
- **whois** — `GET /v1/networking/whois/{domain}`, reuses the existing
  `normalize_domain`/`valid_domain?` private helpers; `404` → "not found"
  message distinct from the generic error (`mx_lookup` pattern).
- **lorem-ipsum** — `GET /v1/text/lorem`, optional `paragraphs`/`sentences`
  (1–20, defaults 1/5), validated both client- and server-side.
- **working-days** — `GET /v1/places/working-days`, required `from`/`to`
  dates plus optional `country`; validates date parsing and `to >= from`
  server-side via a small `parse_iso_date` helper (avoids an inline `rescue`
  modifier to stay consistent with existing style).

Added matching Rack::Attack throttles (`tool_demos/ip` 5/min anonymous,
`tool_demos/user` 15/min free-tier authenticated) mirroring the existing
`api_proxy/*` throttles, and extended `throttled_responder` to recognize
`tool_demos/ip` for the same "create a free account" messaging.

Fixed the one-line ES typo (`Medidores` → `Metros`) alongside the rest of the
locale work since it was in the same file.

## Verification

`bin/rails test` from the host process hung indefinitely (0 bytes of output
after 25+ minutes) — the dashboard actually runs inside Docker
(`requiem-dev-dashboard-1`, via `infra/docker/docker-compose.dev.yml`), and a
bare host Ruby process can't reach the container's Postgres
(`127.0.0.1:5433`, non-default port) or its network-aliased `INTERNAL_API_URL`.
Switched to running tests inside the container:
`docker exec requiem-dev-dashboard-1 bin/rails test ...`.

First run: 208 tests, 3 failures. Two were test bugs I introduced — HTML
entity-encoding mismatches (`isn't` / `Don't` render as `isn&#39;t` /
`Don&#39;t` in the response body, so literal-apostrophe `assert_match` strings
never matched); fixed by matching substrings without the apostrophe, same
pattern the existing `mortgage` test already used for `>` / `&gt;`. The third
(`ToolsControllerTest#test_tools_index_renders_successfully`, expecting a
`.tools-grid` CSS class) is pre-existing and unrelated — confirmed via
`git log` that `app/views/tools/index.html.erb` wasn't touched this session
and the class was removed in an earlier "redesign tools index" commit without
updating the test. Left as-is.

Second run: 208 tests, 207 passing, only the pre-existing `.tools-grid`
failure remains.

## Final notes

The user manually tested the `advice` tool page in the running dev container
and got "Request Failed — Something went wrong. Try again." on every submit.
Investigation (curl against the live `api` container on port 8080) found the
real bug wasn't in this session's code: the documented endpoint
(`config/api_docs/advice.yml`, and by extension `api_catalog.yml`'s implied
contract) says `GET /v1/text/advice`, but the actual Go route — confirmed by
reading `apps/api/services/entertainment/advice/transport_http.go` and its
mount in `apps/api/app/routes_v1.go` — is `GET /v1/entertainment/advice`
(category `entertainment`, not `text`). This is a pre-existing production
documentation bug (the public `/apis/advice` docs page has been serving a
broken curl example), unrelated to any other tool built in this batch — the
other five were spot-checked against their actual Go route registrations and
all matched their `api_docs/*.yml` paths exactly. Fixed both
`ToolDemosController#advice` (now calls `/v1/entertainment/advice`) and
`config/api_docs/advice.yml` (all 5 occurrences of the wrong path corrected)
so the public docs page and the new tool demo agree with the real backend.

No live browser walkthrough of the other 5 new tool pages was done in this
session beyond the automated test suite — worth a manual pass before/after
merge, particularly for `base64`'s mode toggle and `working-days`' date
inputs, which are the two most interactive forms in this batch.

## Files Changed

- `apps/dashboard/app/controllers/tools_controller.rb` — added
  `barcode`/`advice`/`base64`/`whois`/`lorem-ipsum`/`working-days` to
  `SUPPORTED_TOOLS`, `TOOLS_METADATA`, and `CATEGORIES`.
- `apps/dashboard/app/controllers/tool_demos_controller.rb` — 6 new actions +
  a `parse_iso_date` private helper.
- `apps/dashboard/config/routes.rb` — 6 new `tools/demos/*` POST routes.
- `apps/dashboard/app/views/tools/{barcode,advice,base64,whois,lorem_ipsum,working_days}/show.html.erb`
  and the corresponding `partials/tools/<tool>/` directories (6 partials each).
- `apps/dashboard/app/views/tool_demos/{barcode,advice,base64,whois,lorem_ipsum,working_days}.html.erb`
- `apps/dashboard/app/javascript/controllers/{barcode,advice,base64,whois,lorem_ipsum,working_days}_demo_controller.js`
- `apps/dashboard/config/locales/{en,es,fr}/tools.{en,es,fr}.yml` — new
  `tools.<tool>.*` blocks for all 6 tools in all 3 languages; ES
  `unit_conversion.hero.meters` typo fix.
- `apps/dashboard/config/initializers/rack_attack.rb` — `tool_demos/ip` and
  `tool_demos/user` throttles, `throttled_responder` message branch.
- `apps/dashboard/test/controllers/tools_controller_test.rb` — 6 new page
  tests.
- `apps/dashboard/test/controllers/tool_demos_controller_test.rb` — 6 new
  test groups (success / validation / rate-limit / error cases).
- `apps/dashboard/config/api_docs/advice.yml` — fixed 5 occurrences of the
  wrong endpoint path (`/v1/text/advice` → `/v1/entertainment/advice`).
