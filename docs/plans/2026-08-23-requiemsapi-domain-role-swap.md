# requiemsapi.com / requiems.xyz Domain Role Swap — Implementation Plan

Status: proposed, not started.

Supersedes nothing. Per this repo's own convention (see
`docs/audits/2026-08-22-domain-migration-audit-current-state.md`, itself
following
`docs/plans/2026-08-22-go-auth-foundation-phase-8-9-workers-retirement-completion.md`
§9.1), historical audit/plan docs are left unedited as migration records. This
plan is new and answers the open question in
`docs/audits/2026-08-22-domain-migration-audit-current-state.md` §5: the owner
confirmed (conversation, 2026-08-23) that `requiemsapi.com` is already purchased
and the migration is live-wanted, with this exact shape:

> "requiemsapi becomes the dashboard/user-facing [domain]. The api itself should
> be requiems.xyz. The mcp should be [stay] mcp.requiems.xyz. All direct
> UI-facing [traffic] should be requiemsapi.com and the tech stuff at
> requiems.xyz."

This is **not** the original 2026-08-21 audit's plan (Rails stays put, API moves
to a new apex with a path-split). It is a **role swap**: the two existing
production domains trade jobs, and `requiemsapi.com` enters as the new home for
the role `requiems.xyz` vacates.

## 1. Target architecture

| Host                | Today                                                                  | After this plan                                                                                                                                                       |
| ------------------- | ---------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `requiemsapi.com`   | Does not exist anywhere (DNS, Cloudflare zone, or repo)                | **New home for the Rails app**: public site, locale routes, Devise auth/sessions, dashboard, admin, `/api/proxy`, docs, blog, webhooks (Lemon Squeezy)                |
| `requiems.xyz`      | Rails app (see above)                                                  | **New home for the public Go API** — same role `api.requiems.xyz` has today: Caddy with Cloudflare Origin Pull (AOP) mTLS enforced → Kamal → Go, `/v1/*` + `/healthz` |
| `api.requiems.xyz`  | Public Go API                                                          | **Retired.** DNS/Caddy vhost removed once cutover is verified.                                                                                                        |
| `mcp.requiems.xyz`  | MCP HTTP server, upstream `REQUIEMS_BASE_URL=https://api.requiems.xyz` | **Unchanged hostname.** Upstream target updates to `https://requiems.xyz`.                                                                                            |
| `mail.requiems.xyz` | SMTP sending domain (Resend)                                           | **Unchanged** — email deliverability DNS is independent of which app serves web traffic. Not moving unless the owner says otherwise (see §2.2).                       |

```text
Client (browser, dashboard UI)
  -> Cloudflare (zone: requiemsapi.com, new)
  -> Caddy requiemsapi.com { } (no AOP, matches today's requiems.xyz block)
  -> Kamal Go Proxy -> Rails (requiems-dashboard container)

Client (API consumer: curl, SDK, MCP, Rails /api/proxy)
  -> Cloudflare (zone: requiems.xyz, existing)
  -> Caddy requiems.xyz { tls client_auth AOP } (matches today's api.requiems.xyz block)
  -> Kamal Go Proxy -> Go API (requiems-api container)

MCP
  -> Cloudflare (zone: requiems.xyz, existing, unchanged)
  -> Caddy mcp.requiems.xyz { } (unchanged)
  -> Kamal Go Proxy -> MCP container, REQUIEMS_BASE_URL=https://requiems.xyz
```

`requiems.xyz` and `requiemsapi.com` are **different root domains → different
Cloudflare zones**. This is a genuinely new zone to provision, not a DNS record
inside an existing zone. `mcp.requiems.xyz` stays in the existing `requiems.xyz`
zone untouched.

**Important gap found in plan review, not in the table above**: the
`requiems.xyz` zone also carries a live, paying-customer feature this plan
otherwise never mentions — Private Deployment tenant subdomains
(`{slug}.requiems.xyz`, e.g. `acme-team.requiems.xyz`), each a dedicated
per-tenant VPS with **no Cloudflare proxying in front** (DNS-only A records, per
`docs/plans/2026-04-10-private-deployments.md:42,111-115`), unrelated to
`HETZNER_VPS_IP` and to this plan's Caddy/Kamal changes.
`apps/dashboard/app/models/private_deployment_request.rb:84` generates these
URLs and `apps/dashboard/app/mailers/private_deployment_mailer.rb:34` emails
them to customers paying $200–$960/mo as "Your Private API is Live." These
should be structurally unaffected (DNS-only records aren't touched by
proxied-zone AOP scoping, and they're separate VPS targets), but **Phase 0 step
5's live AOP verification must explicitly include confirming a real tenant
subdomain still resolves correctly before and after**, since it's easy to verify
only the two hostnames this plan is actively changing and miss a third live
thing sharing the same zone. There's also a product-copy consequence worth
flagging to the owner: marketing language calling `requiems.xyz` "the tech
stuff" now describes the same zone that customers' branded `{slug}.requiems.xyz`
private deployments live under.

## 2. Decisions the owner must confirm before or during execution

These are called out explicitly rather than assumed, because they're each hard
to reverse cleanly once live.

### 2.1 SEO / redirects for the old `requiems.xyz` marketing pages

`requiems.xyz` currently ranks in search, has backlinks, and is submitted to
Search Console with a live sitemap (`apps/dashboard/config/sitemap.rb`). Once
`requiems.xyz` becomes the API host, the old marketing/dashboard URLs
(`requiems.xyz/en/pricing/`, etc.) will 404 or hit Go's router, unless a
deliberate decision is made:

- **Clean break (matches this repo's stated precedent):** the Worker retirement
  plan explicitly chose "no redirects" over redirect complexity. Applying the
  same policy here means accepting SEO loss on `requiems.xyz`'s existing
  rankings/backlinks.
- **Redirect old paths to `requiemsapi.com`:** requires either a Cloudflare Bulk
  Redirect / Redirect Rule at the `requiems.xyz` zone edge (before traffic
  reaches Caddy/Go) or a small Go handler. Adds a permanent piece of infra that
  has to keep working alongside the API.

This plan defaults to **no redirects** (clean break, consistent with repo
precedent) unless the owner says otherwise before Phase 6. **Confirm before
executing Phase 0.**

### 2.2 Mail domain

`mail.requiems.xyz` (SMTP sending domain, SPF/DKIM/DMARC already provisioned for
it) and `noreply@requiems.xyz` / `noreply@mail.requiems.xyz` sender addresses
are left unchanged by this plan — email deliverability DNS has no relationship
to which app serves web traffic on the apex. If the owner wants outbound mail to
eventually read `@requiemsapi.com`, that is a **separate, independent**
DNS/deliverability project (new SPF/DKIM records, warm-up period) and is out of
scope here.

### 2.3 CORS on the Go API

Once the dashboard and the API are genuinely different root domains (not
subdomains of one site), any client-side JS on `requiemsapi.com` calling
`requiems.xyz` directly is a real cross-origin request, and the API currently
emits **zero** CORS headers (confirmed:
`grep -rn
"Access-Control\|cors\." apps/api --include='*.go'` → no matches).
This plan adds minimal CORS middleware to Go for `/v1/*` (Phase 5) rather than
leaving it an accident, per the current-state audit's explicit recommendation.

## 3. Phase 0 — Cloudflare & DNS prerequisites

Do this first and let it fully propagate before touching app config, so the new
zone is ready the moment Caddy/Kamal need it.

1. Add `requiemsapi.com` as a new zone in the Cloudflare account (owner confirms
   it's already purchased/registered; verify nameservers point at Cloudflare, or
   that a registrar delegation step is still pending — this is the one step this
   plan cannot verify from the repo and needs to be confirmed live in the
   Cloudflare dashboard/API before proceeding).
2. Add a proxied (orange-cloud) `A` record for `requiemsapi.com` (and
   `www.requiemsapi.com` if the owner wants the `www` variant, redirecting to
   apex or vice versa — confirm preference) pointing at the same VPS IP
   (`HETZNER_VPS_IP`) already used by `requiems.xyz`.
3. Confirm Universal SSL / Always Use HTTPS / minimum TLS version on the new
   zone matches the `requiems.xyz` zone's current settings (whatever those are —
   read them from the live zone, don't assume).
4. **Do not** enable Authenticated Origin Pulls (AOP) on the `requiemsapi.com`
   zone — this hostname takes over the Rails role, which has never been behind
   AOP mTLS (only the API host has).
5. On the existing `requiems.xyz` zone: confirm AOP is (or becomes) enforced
   however it's scoped today — check whether AOP was configured zone-wide or
   per-hostname/per-Page-Rule for `api.requiems.xyz` specifically. If it's
   scoped to the hostname rather than the zone, that scoping needs to move from
   `api.requiems.xyz` to bare `requiems.xyz` alongside the Caddy vhost change in
   Phase 1. (`infra/caddy/certs/cloudflare-origin-pull-ca.pem` already exists
   and is zone-independent — the CA cert is reusable as-is, nothing to
   regenerate there.)
6. Mirror any WAF rules / rate-limiting rules / Page Rules that exist on the
   `requiems.xyz` zone today for the Rails app onto the new `requiemsapi.com`
   zone (read current rules from the live zone; the repo has no infra-as-code
   for Cloudflare zone config to diff against, so this is a manual parity
   check).
7. Leave `mail.requiems.xyz` DNS untouched (§2.2).
8. Check for CAA records restricting certificate issuance on the new
   `requiemsapi.com` zone before relying on Caddy's automatic HTTPS to provision
   a cert for it. Low likelihood of an actual restrictive record on a freshly
   purchased domain, but a silent cert-issuance failure at cutover time is cheap
   to rule out now and expensive to debug live.

**Do not proceed to Phase 1 until `https://requiemsapi.com` resolves through
Cloudflare and Caddy can obtain a certificate for it** (verify with a
`curl
-v https://requiemsapi.com` returning _some_ TLS handshake, even a 404,
before cutting the app over).

## 4. Phase 1 — Infra: Caddy + Kamal swap

All three files below need coordinated edits — a partial deploy would route
`requiems.xyz` traffic to the wrong app.

### 4.1 `infra/caddy/Caddyfile`

Rewrite so the vhost blocks trade bodies: the block currently named
`requiems.xyz` (Rails reverse proxy, gzip/zstd, locale-bare redirect) moves its
_behavior_ onto a `requiemsapi.com` block; the block currently named
`api.requiems.xyz` (AOP mTLS `tls { client_auth ... }`) moves its behavior onto
a `requiems.xyz` block. `mcp.requiems.xyz` is untouched. Delete the
`api.requiems.xyz` block once cutover is verified (Phase 7) rather than in this
same edit, so there's a fallback host during the verification window (see §11
rollback).

### 4.2 `infra/kamal/deploy.dashboard.yml`

- `proxy.host: requiems.xyz` → `proxy.host: requiemsapi.com`
- `env.clear.MAILER_HOST: requiems.xyz` (already explicitly set at
  `infra/kamal/deploy.dashboard.yml:71` — this is not defaulting from
  `AppConfig`, it's overriding it in production today) → change the value to
  `requiemsapi.com`
- `env.clear.API_BASE_URL: https://api.requiems.xyz` → `https://requiems.xyz`

### 4.3 `infra/kamal/deploy.api.yml`

- `proxy.host: api.requiems.xyz` → `proxy.host: requiems.xyz`

### 4.4 `infra/kamal/deploy.mcp.yml`

- `env.clear.REQUIEMS_BASE_URL: https://api.requiems.xyz` →
  `https://requiems.xyz`

### 4.5 `infra/docker/docker-compose.yml` and `infra/docker/.env.example`

Update the equivalent local-dev-facing defaults (`REQUIEMS_BASE_URL`, any
`API_BASE_URL`) for consistency, even though local dev typically targets
`localhost:8080` — check each value rather than assuming; the earlier grep
showed `infra/docker/docker-compose.yml:49` has
`REQUIEMS_BASE_URL=https://api.requiems.xyz` baked in as a non-local default.

## 5. Phase 2 — Rails app config (dashboard moves to `requiemsapi.com`)

### 5.1 `apps/dashboard/app/lib/app_config.rb`

- `@api_base_url` default: `"https://api.requiems.xyz"` →
  `"https://requiems.xyz"`
- `@mailer_host` default: `"requiems.xyz"` → `"requiemsapi.com"`

### 5.2 `apps/dashboard/config/initializers/external_links.rb`

- `WEBSITE[:home]`: `"https://requiems.xyz"` → `"https://requiemsapi.com"`
- `WEBSITE[:api_docs]`: `"https://requiems.xyz/apis"` →
  `"https://requiemsapi.com/apis"`

### 5.3 `apps/dashboard/config/sitemap.rb`

- `SitemapGenerator::Sitemap.default_host` and all inline
  `"https://requiems.xyz/..."` string interpolations (alternate-language `href`s
  — 21 occurrences total including `default_host`, verified count) →
  `https://requiemsapi.com`. Regenerate the sitemap after
  (`rails sitemap:refresh` or whatever task this repo uses) and re-submit it in
  Search Console for the new property (see Phase 6).

### 5.4 `apps/dashboard/app/helpers/application_helper.rb` and `case_studies_helper.rb`

JSON-LD structured data: 12 hardcoded `"https://requiems.xyz"` /
`"https://requiems.xyz/logo.png"` / `"https://requiems.xyz/og-image.png"`
strings → `requiemsapi.com` equivalents. (`request.base_url` used for the
`<link rel="canonical">` tag is dynamic and needs no change — it already follows
whatever host serves the request.)

Also in `application_helper.rb` (line ~114): the FAQ's
`Authorization: Bearer YOUR_API_KEY` string is the **Bearer-vs-header-key
documentation bug** flagged in the current-state audit (§4/§6) — fix it here to
say `requiems-api-key: YOUR_API_KEY` (the actual auth contract per
`apps/api/platform/middleware/apikeyauth.go:18`), not just the hostname.

### 5.5 Mailers

- `apps/dashboard/app/mailers/private_deployment_mailer.rb`:
  `@docs_url = "https://requiems.xyz/docs"` (×2) →
  `https://requiemsapi.com/docs`
- `apps/dashboard/app/mailers/application_mailer.rb` default `from`
  (`noreply@requiems.xyz`) — **leave unchanged**, it's an email address, not
  covered by §2.2's decision to leave the mail domain alone.
- `config/environments/production.rb`'s `ActionMailer::Base.default_url_options`
  reads `AppConfig.mailer_host` — no direct edit needed here, it inherits from
  §5.1/§4.2. This is the mechanism that makes password-reset and confirmation
  email links point at the right host; **this is the single highest-risk item in
  Phase 2** — a stale value here silently breaks password reset for every user
  until caught. Test explicitly (see §7).

### 5.6 `config.hosts` production allowlist (bonus fix, bundled here)

The current-state audit flags this as **STILL OPEN, HIGH** independent of this
migration (`grep -n "config.hosts"
apps/dashboard/config/**/*.rb` → nothing in
production). Since this plan is already editing
`config/environments/production.rb`, add the allowlist here rather than as a
separate, easy-to-forget follow-up:

```ruby
config.hosts << "requiemsapi.com"
config.hosts << "www.requiemsapi.com" # only if the www variant is served, not just redirected
```

**Risk flagged in plan review, HIGH — sequencing change recommended.**
`config.hosts` is currently empty everywhere in this app (confirmed: no other
file sets or bypasses it) — Rails' `ActionDispatch::HostAuthorization`
middleware today permits _any_ `Host` header, including whatever Kamal's own
health-check probe sends. Bundling the **first-ever** host restriction this app
has had into the same deploy as the domain cutover risks the probe's `Host`
header (container-internal address, a stale/mismatched value during DNS
propagation, etc.) getting a 403 on `/up`, which Kamal would read as an
unhealthy deploy and could roll back — mid-cutover, for a reason unrelated to
whether the new domain actually works. **Recommend deploying this specific line
as its own small, low-risk change _before_ Phase 1–5's cutover deploy** (with
the _old_ host, `requiems.xyz`, allowlisted first, to prove `config.hosts`
doesn't break Kamal's healthcheck at all before also depending on it during the
higher-stakes cutover deploy), rather than as one more line in the same commit
as everything else.

## 6. Phase 3 — API-facing bulk reference sweep (`api.requiems.xyz` → `requiems.xyz`)

`grep -rn "api\.requiems\.xyz"` currently matches **~600 lines across ~90
files**, but almost all of it is generated or mechanically identical. Confirm
the shape before bulk-editing:

1. **`apps/dashboard/config/api_docs/*.yml`** (~65 files, the large majority of
   the matches): each has a `base_url: https://api.requiems.xyz` field plus
   repeated inline example URLs using the same string. **Correction from plan
   review**: these files are _not_ a generation source for `docs/apis/**/*.md` —
   verified by reading the actual code path. `ApiDocs::SnippetGenerator`
   (`apps/dashboard/app/services/api_docs/snippet_generator.rb`) is invoked at
   runtime only, by `apps/dashboard/app/helpers/apis_helper.rb:79`, to backfill
   missing `code_examples` for the in-app `/apis/{id}` page — it never writes to
   `docs/apis/`. `apps/dashboard/script/golden_diff_api_docs.rb` compares
   generator output against the YAML's own hand-written `code_examples` and
   writes its report to `docs/plans/2026-08-21-golden-diff-report.md`, not
   `docs/apis/`. No script anywhere writes `docs/apis/*.md`. **This means both
   the ~65 YAML files and the ~25 markdown files under `docs/apis/` need
   independent hand-edits (or a purpose-built bulk find/replace across both
   trees), and there is no drift-detection between them today** — treat them as
   two separate, unrelated bulk-replace passes, not a generate-from-source step.
2. **`docs/apis/**/*.md`** (~25 files) — independent hand-edit/bulk-replace
   pass, per the correction above (not a regeneration).
3. **`apps/dashboard/config/locales/{en,es,fr}/home_index.*.yml`** (all three
   locales, not just en/es — `fr` has the identical key too) — literal
   `api_requiems_xyz_v1: api.requiems.xyz/v1/` string → `requiems.xyz/v1/`.
4. **`readme.md`** (root), **`docs/core/adding-tools.md`**,
   **`docs/core/infrastructure.md`**, **`docs/core/deployment.md`**, and
   **`apps/dashboard/docs/app-config.md`** (`API_BASE_URL`'s documented default)
   — hand edit, small number of occurrences each, some are prose describing the
   _old_ Worker-era topology that's already stale for unrelated reasons
   (`docs/core/infrastructure.md` still says "Cloudflare terminates the public
   edge and requires the Cloudflare Origin Pull client certificate at Caddy"
   attached to `api.requiems.xyz` specifically — update this description to
   match the new host along with the rename).
   **`apps/dashboard/public/robots.txt`** also needs its own hand-edit here — it
   hardcodes `Sitemap: https://requiems.xyz/sitemap.xml`, same SEO-risk class as
   §5.3 but a file that sweep missed.
5. **`tests/integration/src/config.ts`, `tests/integration/src/reporter.ts`,
   `tests/integration/README.md`, `tests/load/.env.example`** — default
   `API_BASE_URL` fallback values, `https://api.requiems.xyz` →
   `https://requiems.xyz`.
6. **`apps/dashboard/test/scripts/golden_diff_api_docs_test.rb`,
   `apps/dashboard/test/services/api_docs/snippet_generator_test.rb`** — test
   fixtures hardcode the old host; update alongside step 1 or these tests will
   fail against the regenerated docs.
7. **`apps/dashboard/content/blog/2026-07-27-*.md`** — a blog post with a worked
   example using `api.requiems.xyz/openapi.json` and a curl example. Judgment
   call: editing published blog content changes its historical accuracy.
   Recommend leaving blog _post_ prose as a historical snapshot (it describes
   what was true when published) unless the owner wants it corrected — flag for
   an explicit decision rather than silently rewriting or silently skipping.
8. **`apps/mcp/scripts/fetch-spec.ts`, `apps/mcp/scripts/generate.ts`** — see
   Phase 4, bundled with the `/openapi.json` fix.
9. **Committed, static sitemap XML files** — `apps/dashboard/public/sitemap.xml`
   (5 matches), `apps/dashboard/public/sitemap_apis.xml` (840 matches),
   `apps/dashboard/public/sitemap_static.xml` (1035 matches) are git-tracked,
   pre-generated files, not runtime output — `infra/docker/dashboard.Dockerfile`
   has no sitemap-generation build step. Regenerating via the sitemap rake task
   (§5.3) only helps if the regenerated files are **committed to git before the
   image is built** — a Docker image is immutable at runtime, so "regenerate
   after deploy" as originally phrased in §5.3 is not sufficient by itself;
   regeneration has to happen pre-build, in the same commit/PR as the rest of
   Phase 2–3.
10. **`apps/api/services/places/geocode/service.go:43`** —
    `userAgent = "requiems-api/1.0 (https://requiems.xyz)"`, an outbound
    contact-URL string sent to the Nominatim/OpenStreetMap geocoding provider
    per their usage-policy requirements. Post-swap this becomes
    self-referentially correct (points at the API host itself) but should still
    be a deliberate edit, not a missed one — confirm intent rather than leaving
    it un-inventoried.

**Do not edit** `docs/plans/*.md` or `docs/audits/*.md` — those are historical
records per repo convention and several of the matches above (phase 7/8-9 plan
docs, the two domain-migration audits) are exactly that.

## 7. Phase 4 — MCP

1. `infra/kamal/deploy.mcp.yml`: `REQUIEMS_BASE_URL` → `https://requiems.xyz`
   (done in Phase 1.4, listed here for completeness of the MCP checklist).
2. `apps/mcp/scripts/fetch-spec.ts`: `SPEC_URL` →
   `https://requiems.xyz/openapi.json`
3. `apps/mcp/scripts/generate.ts`: comment referencing the same URL, update for
   consistency.
4. This is **blocked on Phase 5.1** (`/openapi.json` currently 404s on Go —
   confirmed by the current-state audit, `apps/api/app/app.go:56` only registers
   `/healthz`). Fetching the spec from the new host will fail exactly the same
   way it already silently fails today unless Phase 5.1 restores the endpoint
   first. Verify `apps/mcp`'s actual current spec source before assuming
   `fetch-spec.ts` is even live in the current build pipeline — the
   current-state audit explicitly flagged this as "not yet verified."

## 8. Phase 5 — Findings this migration surfaces or must not leave broken

These are bundled into this plan (not deferred) because the migration directly
touches the same files/hosts, per the user's ask to catch "extra things that
might be needed" along the way.

### 8.1 Restore `/openapi.json` on the Go API (current-state audit §4, LOW→ but blocking Phase 4)

Go's router (`apps/api/app/app.go:56`) only registers `/healthz`. Decide and
implement one of:

- Add a `router.Get("/openapi.json", ...)` handler serving a build-time or
  embedded spec file, matching what the retired Worker used to do, **or**
- Serve it from the Rails side instead (`requiemsapi.com/openapi.json` or
  `/apis/openapi.json`) and update `fetch-spec.ts` to point there instead of the
  API host.

Either is acceptable; leaving it a silent 404 is not — MCP's build depends on
it.

### 8.2 Add CORS middleware to Go for `/v1/*` (current-state audit §4, MEDIUM)

Per §2.3 above: add minimal CORS headers (`Access-Control-Allow-Origin`,
preflight `OPTIONS` handling) scoped to the public `/v1/*` surface in
`apps/api/app/app.go` / a new middleware file alongside the existing
`apikeyauth.go`/`ratelimit.go`. Confirm with the owner whether this should be
wildcard `*` (matches the retired Worker's prior behavior) or restricted to
known origins (e.g., `requiemsapi.com`) — wildcard is simpler and matches prior
behavior; restricted is tighter but breaks arbitrary third-party browser callers
the docs currently imply are supported.

### 8.3 Fix the Bearer-vs-header-key documentation contradiction, repo-wide

Not just `application_helper.rb` (§5.4) — also `readme.md:103` and the three
`apps/dashboard/config/locales/{en,es,fr}/comparisons.*.yml` files. **Scope
correction**: each locale file has **10** `Authorization: Bearer <key>`
examples, not one (verified: `comparisons.en.yml` lines 66, 146, 228, 308, 392,
472, 555, 639, 714, 798 — paired 1:1 with the 10 `api.requiems.dev` occurrences
fixed in §8.4 below, same lines). One coherent pass across all 30 occurrences
(10 × 3 locales): all examples should show `requiems-api-key: YOUR_API_KEY`,
matching `apps/api/platform/middleware/apikeyauth.go:18`.

### 8.4 Fix stale `api.requiems.dev` marketing copy

30 lines across `apps/dashboard/config/locales/{en,es,fr}/comparisons.*.yml`
reference a fictional `api.requiems.dev` in competitor-comparison marketing copy
(e.g., "Replace `api.ipstack.com/{ip}` with
`api.requiems.dev/v1/ip/lookup?ip={ip}`"). Replace with the real new API host
`requiems.xyz` in the same pass as 8.3 (both are in the same three files).

## 9. Phase 6 — External services & manual steps (not repo edits)

These happen outside the codebase and need explicit sequencing around the DNS
cutover, not before or long after:

1. **Lemon Squeezy webhook URL** — currently posts to
   `requiems.xyz/webhooks/lemonsqueezy` (Rails). Must be updated in the Lemon
   Squeezy dashboard to `requiemsapi.com/webhooks/lemonsqueezy` **at the same
   time** DNS/Caddy cut over, or subscription webhooks will either 404 (pointed
   at the new API host with no Rails behind it) or silently stop arriving
   (pointed at old host once Rails moves off it). Test with Lemon Squeezy's
   webhook test-send feature immediately after cutover. 1a. **Lemon Squeezy
   checkout success/cancel redirect — separate from the webhook URL, found in
   plan review.**
   `apps/dashboard/app/controllers/dashboard/billing_controller.rb:98-115`
   builds the checkout URL with no `redirect_url` query param, meaning the
   post-payment return destination is configured on the Lemon Squeezy
   product/checkout settings themselves (LS dashboard side), not derivable from
   this repo. If left pointed at `requiems.xyz`, a customer who just paid gets
   redirected to the bare Go API host immediately after checkout — the
   subscription still activates via the webhook, but the customer sees what
   looks like a broken payment flow. **Update this alongside the webhook URL in
   the Lemon Squeezy dashboard**, not just the webhook endpoint.
2. **Google Search Console** — add `requiemsapi.com` as a new property, submit
   the regenerated sitemap (§5.3). Decide whether to also keep the
   `requiems.xyz` property (now showing 404s for old paths, per §2.1's default
   no-redirect policy) or remove it.
3. **GA4** — confirm `GA4_MEASUREMENT_ID` (`infra/kamal/deploy.dashboard.yml`,
   unchanged value) is configured in Google Analytics to accept
   `requiemsapi.com` as a valid stream hostname; GA4 properties are usually
   host-agnostic per measurement ID, but verify rather than assume.
4. **Any OAuth app / webhook allow-lists at third parties** that reference
   `requiems.xyz` by exact origin (check GitHub OAuth app settings if one
   exists, any other integration with a configured redirect/callback URL) —
   audit and update case by case; nothing found in-repo pointing at one, but
   this class of config lives outside the repo by definition.

## 10. Phase 7 — Cutover sequencing & verification

Recommended order to minimize the window where something is half-migrated:

1. Complete Phase 0 fully; confirm `requiemsapi.com` resolves through Cloudflare
   with a valid cert, _before_ changing any app config.
2. Merge Phases 1–5 code changes together (they're interdependent — a partial
   deploy misroutes traffic) and deploy via the existing Kamal CD path for all
   three services (`dashboard`, `api`, `mcp`) together. 2a. **Explicitly
   force-restart the Caddy accessory, don't assume CD does it.** Caddy runs as a
   Kamal accessory (`infra/kamal/deploy.api.yml:96-106`) bind-mounting
   `infra/caddy/Caddyfile` read-only into an already-running container; the
   existing CD workflow always runs `kamal setup`, never `kamal deploy`, for
   accessories (`.github/actions/kamal-deploy/action.yml:52`). Whether
   `kamal setup` reliably force-recreates an _already-running_ accessory to pick
   up a changed mounted file — versus treating it as already-booted and leaving
   the stale Caddyfile mounted — is untested here. If it doesn't, every
   app-level deploy can report success while Caddy silently keeps routing the
   old way, invisible to Kamal's own health signal. Run
   `kamal accessory reboot caddy -c infra/kamal/deploy.api.yml` (or equivalent
   forced restart) as an explicit, named step, not an assumed side effect of the
   app deploys.
3. Immediately after deploy, run the smoke-test matrix below. Note that nothing
   in `.github/workflows/` currently runs any curl/health/smoke check
   post-deploy (confirmed: no matches for `curl|health|smoke` in the CD
   workflow) — this matrix is a **manual runbook**, not CI-enforced. Assign who
   runs it before cutover, not during.
4. Update the Lemon Squeezy webhook URL (Phase 6.1) within the same maintenance
   window — not before (would point live webhooks at a host with no Rails yet
   during Phase 0–2) or long after (would drop events in the gap).
5. Only after the new topology is verified live, delete the old
   `api.requiems.xyz` Caddy vhost/DNS record (§4.1's deferred cleanup) and the
   corresponding Cloudflare AOP scoping if it was hostname-specific.

### Smoke-test matrix (mirrors the rigor of the phase 8/9 Worker-retirement doc)

| Check                                                                                                     | Expected                                                                                                                                                                                                      |
| --------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `curl https://requiemsapi.com/`                                                                           | Rails home page, 200                                                                                                                                                                                          |
| `curl https://requiemsapi.com/healthz`-equivalent (whatever Rails exposes, e.g. `/up`)                    | 200                                                                                                                                                                                                           |
| Full Devise login/logout on `requiemsapi.com`                                                             | Session cookie set/read correctly on the new domain, no CSRF/host errors                                                                                                                                      |
| Password reset email → click link                                                                         | Link host is `requiemsapi.com`, not stale `requiems.xyz` (validates §5.5)                                                                                                                                     |
| `curl https://requiems.xyz/v1/entertainment/advice -H "requiems-api-key: ..."`                            | 200, real API response                                                                                                                                                                                        |
| `curl -v https://requiems.xyz/v1/...` **without** a valid Cloudflare AOP client cert, direct to origin IP | TLS `certificate_required` rejection (validates AOP moved correctly, matches the phase 8/9 doc's own verification method)                                                                                     |
| `curl https://requiems.xyz/healthz`                                                                       | 200                                                                                                                                                                                                           |
| `curl https://requiems.xyz/openapi.json`                                                                  | 200 with valid spec (validates §8.1)                                                                                                                                                                          |
| `curl -H "Origin: https://requiemsapi.com" -X OPTIONS https://requiems.xyz/v1/...`                        | CORS preflight succeeds (validates §8.2)                                                                                                                                                                      |
| `curl https://mcp.requiems.xyz/...`                                                                       | Unchanged behavior, confirms `REQUIEMS_BASE_URL` repoint didn't break MCP                                                                                                                                     |
| MCP spec fetch script (`apps/mcp/scripts/fetch-spec.ts`)                                                  | Successfully downloads from `requiems.xyz/openapi.json`                                                                                                                                                       |
| Lemon Squeezy test webhook send                                                                           | Delivered and processed at `requiemsapi.com/webhooks/lemonsqueezy`                                                                                                                                            |
| `curl -I https://requiemsapi.com/` and `https://requiems.xyz/`                                            | `Strict-Transport-Security` header present (this was already an unverified open item in the current-state audit — verify now while touching both hosts, even though HSTS itself isn't part of this migration) |
| Rails `/api/proxy` (server-to-server, private network path)                                               | Still works unaffected — confirms `INTERNAL_API_URL` (private Docker network) wasn't accidentally touched                                                                                                     |

## 11. Rollback plan

The code/config layer of Phase 1–5 is an ordinary redeploy-the-previous-commit
rollback, same as any other change. **What plan review flagged as understated**:
a domain swap has real, non-code state that a code rollback cannot undo, so
"redeploy the previous commit" is necessary but not sufficient. Specifically:

1. **DNS propagation lag**: if Cloudflare DNS for either zone needs to revert,
   allow for TTL/propagation delay — this is not instant like a Kamal redeploy.
   Keep the old `api.requiems.xyz` Caddy vhost and DNS record alive (per §10.5,
   deferred deletion) specifically so there's a known-good fallback host
   reachable during the verification window without needing a DNS change to roll
   back.
2. **Session cookies are not portable across a rollback.** Once real users
   authenticate on `requiemsapi.com` during the migration window, their session
   cookie is scoped to that domain (confirmed: no shared-cookie mechanism or
   explicit `domain:` override exists anywhere in `apps/dashboard/config`).
   Rolling Rails back to `requiems.xyz` **logs out every user who signed in
   during the new-domain window**, not just new sign-ups after cutover — treat
   this as a real, user-visible incident to warn users about if a rollback
   happens, not a silent no-op.
3. **Lemon Squeezy state written during the window survives a rollback.**
   Webhooks processed against `requiemsapi.com` during the migration window
   write real subscription state to Postgres; reverting code doesn't revert that
   data. If rolling back, also revert the webhook URL _and_ the checkout
   redirect URL (§9.1/§9.1a) back to the pre-migration host in the same rollback
   window — the same drop/misroute risk applies in reverse that applies on the
   forward cutover (§10.4) — and manually reconcile any subscription events that
   landed during the window before the revert completes.

## 12. Explicitly out of scope / deferred backlog

Surfaced during the investigation for this plan but **not** bundled in, because
they're not touched by the domain swap itself and bundling them in would expand
this into an unrelated change:

- **Webhook event-ID idempotency** on
  `apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb` — still
  HIGH-equivalent risk per the current-state audit (duplicate delivery can
  repeat subscription side effects), untouched by v2, unrelated to which host
  the webhook arrives on.
- **Live HSTS verification** — added to the smoke-test matrix above as a quick
  check since both hosts are being touched anyway, but fixing it (if missing) is
  separate work in Caddy/Cloudflare config, not part of this plan's diff.
- **MCP fetch/runtime/snapshot URL layer consistency** and the
  `(method, path, operationId)` drift check (e.g., `advice`
  `/v1/entertainment/advice` vs `/v1/text/advice`) that the original 2026-08-21
  audit flagged — not reverified in the current-state audit, not reverified here
  either; independent of hostnames.
- **Blog post historical accuracy** (§6, item 7) — explicit owner decision
  needed, not resolved by this plan.
- **`www.requiemsapi.com` handling** — flagged as an open question in Phase 0.2,
  not resolved by this plan; needs an explicit owner answer (redirect to apex,
  serve directly, or don't provision it at all).

## 13. Review findings log

This plan was reviewed by two independent passes (correctness-against-repo, and
completeness/risk-for-production-cutover) before implementation started. HIGH
and most MEDIUM findings from both are folded into the plan body above, at point
of relevance, with an inline note where they came from a review correction. This
section is the complete record, including items judged LOW priority or
explicitly out of scope, so nothing found along the way gets lost.

### Folded into the plan body (see cross-references above)

- Phase 3's generator claim was factually wrong — corrected (§6 item 1–2).
- Private Deployment tenant subdomains were entirely missing — added (§1).
- `config.hosts` bundled-into-cutover rollback risk — added (§5.6).
- Lemon Squeezy checkout redirect URL (separate from webhook URL) — added (§9,
  item 1a).
- Rollback plan understated session/webhook irreversibility — strengthened
  (§11).
- `robots.txt` and committed static sitemap XML files missing from inventory —
  added (§6 items 4, 9).
- Kamal `kamal setup` vs. Caddy accessory restart uncertainty — added (§10, step
  2a).
- No CI-enforced smoke test exists — noted (§10, step 3).
- CAA record check missing from Phase 0 — added (§3, item 8).
- `MAILER_HOST` "currently absent" was a false premise (it's already set in
  `deploy.dashboard.yml:71`) — corrected (§4.2).
- Bearer-auth example count in `comparisons.*.yml` was understated 10x (10 per
  locale, not 1) — corrected (§8.3).
- `home_index.fr.yml` was missing from the locale sweep list — corrected (§6
  item 3).
- `sitemap.rb`/helper JSON-LD occurrence counts were undercounted (~15→21,
  ~9→12) — corrected (§5.3, §5.4).
- `apps/api/services/places/geocode/service.go:43` outbound User-Agent URL —
  added to inventory (§6 item 10).

### LOW priority, not folded into the plan body (informational, low stakes)

- `apps/dashboard/docs/app-config.md:10` documents `API_BASE_URL`'s default —
  added to the docs list (§6 item 4) rather than given its own subsection.
- `apps/dashboard/config/initializers/devise.rb:27` — `config.mailer_sender`
  fallback default is `noreply@requiems.xyz`, inconsistent with the real
  production value (`noreply@mail.requiems.xyz` from `deploy.dashboard.yml:70`).
  Pre-existing, dead in production (the Kamal env var always wins), unrelated to
  this migration's correctness — worth a drive-by fix but not blocking.
- `infra/docker/.env.example:104` — `MAILER_HOST=requiems.xyz` dev default,
  low-stakes, not production-facing.
- `infra/caddy/Caddyfile:2` — global ACME contact `email admin@requiems.xyz` —
  informational only, no action needed; unrelated to which app serves which
  host.
- No `.well-known/` directory exists in this app today (confirmed) — genuinely
  not applicable, not a regression.
- CSP (`apps/dashboard/config/initializers/content_security_policy.rb`) is
  entirely commented out — confirmed zero active directives, so there's no CSP
  host-allowlist risk to address.
- `tests/integration/.env.example:10,16` has the same
  `API_BASE_URL=https://api.requiems.xyz` default as `tests/load/.env.example`
  (already listed in §6 item 5) — add this sibling file to the same bulk-replace
  pass.
- `apps/dashboard/test/mailers/private_deployment_mailer_test.rb`,
  `apps/dashboard/test/models/private_deployment_request_test.rb` — reference
  tenant subdomain URLs (`https://acme-team.requiems.xyz`); correct to leave
  unchanged since tenant subdomains aren't moving (§1's new tenant-subdomain
  note), called out explicitly here so it reads as a deliberate no-op, not a
  miss.

### Explicitly out of scope — candidate backlog items, not part of this migration

- **Webhook event-ID idempotency** on `Webhooks::LemonsqueezyController` —
  already listed in §12, reconfirmed by both reviews as still-open and unrelated
  to hostnames.
- **Live HSTS verification** — already in the smoke-test matrix (§10) as a quick
  check; the underlying fix (if missing) is separate Caddy/Cloudflare work.
- **Private Deployment tenant infrastructure has no infra-as-code in this repo
  at all** (provisioning is entirely manual per
  `docs/plans/2026-04-10-private-deployments.md:111-115`) — unrelated to this
  domain swap, but means nobody can audit from the repo how many live tenant
  subdomains exist or how they're configured. Worth a separate inventory task.
- **MCP fetch/runtime/snapshot URL layer consistency** and the
  `(method, path, operationId)` drift check from the original 2026-08-21 audit —
  not reverified by either review pass, independent of hostnames.

## 14. File change summary (for implementation tracking)

| Area              | Files                                                                                                                                                                                                                                                                   | Nature of change                                                                             |
| ----------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| Cloudflare/DNS    | _(no repo files — live dashboard/API config)_                                                                                                                                                                                                                           | New zone, DNS, AOP scoping, WAF parity                                                       |
| Caddy             | `infra/caddy/Caddyfile`                                                                                                                                                                                                                                                 | Swap vhost bodies, add `requiemsapi.com`, retire `api.requiems.xyz` after verification       |
| Kamal             | `infra/kamal/deploy.dashboard.yml`, `deploy.api.yml`, `deploy.mcp.yml`                                                                                                                                                                                                  | `proxy.host` swaps, env var updates                                                          |
| Docker local dev  | `infra/docker/docker-compose.yml`, `.env.example`                                                                                                                                                                                                                       | Consistency, non-blocking                                                                    |
| Rails config      | `app_config.rb`, `external_links.rb`, `sitemap.rb`, `application_helper.rb`, `case_studies_helper.rb`, `private_deployment_mailer.rb`, `config/environments/production.rb`, `public/robots.txt`, `public/sitemap*.xml` (committed, regenerate pre-build)                | Host defaults, JSON-LD, mailer, `config.hosts` (ship separately first, see §5.6)             |
| API docs source   | `apps/dashboard/config/api_docs/*.yml` (~65 files)                                                                                                                                                                                                                      | Bulk `base_url`/example replace — independent pass, no generator links these to `docs/apis/` |
| API docs (site)   | `docs/apis/**/*.md` (~25 files)                                                                                                                                                                                                                                         | Independent bulk replace — **not** generated from `config/api_docs/*.yml`, see §6 item 1     |
| Locales           | `config/locales/{en,es,fr}/comparisons.*.yml`, `home.*.yml`, `home_index.*.yml` (all 3 locales)                                                                                                                                                                         | Bearer-auth fix (30 occurrences), `api.requiems.dev` fix (30 occurrences), hostname strings  |
| Root docs         | `readme.md`, `docs/core/{infrastructure,deployment,adding-tools}.md`, `apps/dashboard/docs/app-config.md`                                                                                                                                                               | Hand edit                                                                                    |
| Tests             | `tests/integration/src/{config,reporter}.ts`, `tests/integration/README.md`, `tests/integration/.env.example`, `tests/load/.env.example`, `apps/dashboard/test/scripts/golden_diff_api_docs_test.rb`, `apps/dashboard/test/services/api_docs/snippet_generator_test.rb` | Default URL fixtures                                                                         |
| MCP               | `apps/mcp/scripts/fetch-spec.ts`, `generate.ts`                                                                                                                                                                                                                         | New spec URL, blocked on §8.1                                                                |
| Go API            | `apps/api/app/app.go` (+ new middleware file), `apps/api/services/places/geocode/service.go:43`                                                                                                                                                                         | Restore `/openapi.json`, add CORS, outbound User-Agent URL                                   |
| External services | Lemon Squeezy webhook URL, Lemon Squeezy checkout redirect URL, Search Console, GA4                                                                                                                                                                                     | Manual, sequenced with cutover                                                               |
| Infra ops         | Kamal accessory restart for Caddy (`kamal accessory reboot caddy`)                                                                                                                                                                                                      | Explicit step, not assumed side effect of app deploys — see §10, step 2a                     |

4
