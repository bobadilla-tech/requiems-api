# SEO Audit & Site-Wide JSON-LD Coverage

A technical SEO audit of requiems.xyz (`apps/dashboard`, Rails 8), followed by a
locale-link correctness pass, a sitemap hygiene fix, and a site-wide push to
give every crawlable page structured data — plus a verification pass against a
subsequent code review of the changes.

---

## Context

The app is a Rails 8 marketing + API-directory/docs hybrid (not Next.js, as
initially assumed) with i18n locale routing (`en`/`es`/`fr`, prefix-always,
enforced via a 301 in
[`application_controller.rb`](../../apps/dashboard/app/controllers/application_controller.rb)).
Before this session:

- Every page already got baseline `Organization` + `WebSite` + `BreadcrumbList`
  JSON-LD globally via `json_ld_tags` in
  [`application_helper.rb`](../../apps/dashboard/app/helpers/application_helper.rb),
  called from the layout. Breadcrumbs are auto-derived from the URL path — no
  per-page setup needed.
- Only 5 pages had page-specific schema on top of that baseline: `home/faq`
  (FAQPage), `blog/show` (BlogPosting), `apis/show` (WebAPI),
  `case_studies/show` (Article, via a helper already written for that page), and
  none for pricing.
- 7 tool FAQ partials linked to `/contact` with a raw `<a href>` instead of the
  locale-aware `contact_path` helper — bypassing `default_url_options[:locale]`
  and risking a silent language flip.
- `website_json_ld` and `api_json_ld` hardcoded `/en/` in two URLs, ignoring the
  visitor's actual locale.
- `public/core-sitemap.xml` (8,245 lines) was orphaned — marked "retired" in a
  rake task comment but still on disk and crawlable.
- The sitemap (`sitemap_generator` gem, static XML files) was only regenerated
  manually — no CI or cron job called `rake sitemap:refresh`.
- Plan pricing (`$0/$30/$75/$150`) was duplicated across
  `home/pricing.html.erb`, `admin/dashboard_controller.rb`, and
  `AnalyticsRevenueService`, even though a canonical `PlanConfig`
  (`app/lib/plan_config.rb`) already existed and was only used for billing/Lemon
  Squeezy variant lookups.

## Approach

### 1. Locale-link fixes

Swapped the 7 hardcoded `<a href="/contact">` anchors for
`link_to ..., contact_path`, and replaced the 2 hardcoded `/en/` JSON-LD URLs
with `I18n.locale` interpolation, so every locale gets a self-referential URL
instead of always pointing at English.

### 2. Sitemap hygiene

Deleted the orphaned `core-sitemap.xml` (confirmed unreferenced by `robots.txt`
and the current sitemap index). Added
[`RefreshSitemapJob`](../../apps/dashboard/app/jobs/refresh_sitemap_job.rb),
scheduled nightly at 03:00 UTC via
[`sidekiq_schedule.yml`](../../apps/dashboard/config/sidekiq_schedule.yml) (the
live scheduler here is sidekiq-cron, not Solid Queue — `config/recurring.yml`
exists but is dead config, confirmed via `queue_adapter = :sidekiq` in
`config/environments/production.rb` and no `solid_queue` gem in the Gemfile).
The job lazily loads rake tasks and executes `sitemap:refresh` in-process.

### 3. Pricing JSON-LD + PlanConfig consolidation

Added `pricing_json_ld` (Product + one Offer per tier) to
`home/pricing.html.erb`. While building it, found the same plan-price literals
duplicated in three other places — including a genuine semantic drift, not just
repetition: `AnalyticsRevenueService::PLAN_PRICES` stored the yearly value as a
monthly-equivalent (e.g. `62.5`), while `PlanConfig::PLANS` stored the yearly
value as the annual total (e.g. `750`). Any future price change would have had
to be applied correctly in two different unit conventions by hand. Fixed by
adding `PlanConfig.price_yearly_monthly` and rebuilding all three call sites
(`pricing.html.erb`, `admin/dashboard_controller.rb`,
`AnalyticsRevenueService::PLAN_PRICES`) from `PlanConfig` directly, plus
`PlanConfig.formatted_requests` / `formatted_rate_limit` to stop
`pricing.html.erb` from hand-formatting `"100k"` / `"5,000/min"` strings (this
also fixed a pre-existing inconsistency where the developer tier showed
`"5,000/min"` on the pricing card but `"5k/min"` in the comparison table on the
same page).

### 4. Site-wide JSON-LD coverage

Mapped all ~35 remaining page types via an Explore pass and split them into:
pages with a strong schema.org fit (built), and purely static pages with none
(explicitly skipped — forcing schema onto a privacy/terms/status page has no SEO
upside and was flagged as bad practice rather than done anyway).

Added 8 reusable helpers to `application_helper.rb` — `collection_json_ld`
(generic `CollectionPage`/`ItemList`, reused by all 11 index/listing pages
instead of one bespoke method per controller), `service_json_ld` (systems,
industries), `software_application_json_ld` (all 20 tool detail pages),
`article_json_ld` (comparisons, examples, and `api_reference` — a real gap the
final coverage grep caught that hadn't been in the original inventory),
`contact_page_json_ld` (contact, suggest-an-api, talk-to-sales,
private-deployment), `about_page_json_ld`, `team_json_ld`, and
`glossary_json_ld` (`DefinedTermSet`, built from the existing 12-entry glossary
i18n data).

One bug caught mid-implementation: industry and comparison slugs contain hyphens
(`real-estate`, `api-ninjas`) but their i18n translation keys use underscores —
the first pass on those two index pages would have raised
`I18n::MissingTranslationData` in production.

### 5. Review-fix pass

A subsequent review of the diff raised 6 findings. Verified each against current
code before touching anything (task-list findings can be stale):

- **Real, fixed** — `admin/dashboard_controller#calculate_mrr` only grouped by
  `plan_name`, pricing every subscription at its monthly rate even when billed
  yearly. Fixed to group by `[plan_name, plan]` (billing cycle) and price yearly
  subs via `PlanConfig.price_yearly_monthly`, matching
  `AnalyticsRevenueService`'s existing logic.
- **Real, fixed** — the two locale-fixed JSON-LD URLs still hardcoded the
  `https://requiems.xyz` host; swapped to `request.base_url`.
- **Real, fixed** — `RefreshSitemapJob` had no test coverage; added
  `test/jobs/refresh_sitemap_job_test.rb` covering the success path, the
  lazy-`load_tasks` path, and the log-and-reraise error path, all via
  `Minitest::Mock`/`Object#stub` rather than touching the real rake task.
- **Real, fixed** — `PlanConfig.humanize_count` used integer division
  (`n / 1_000`), which would silently truncate any future non-round value (e.g.
  `1500` → `"1k"` instead of `"1.5k"`). All current values happen to divide
  evenly, so this wasn't live yet, but the fix was cheap and closes a real
  correctness gap for any future price/limit tier.
- **Real, fixed** — `AnalyticsRevenueService#calculate_mrr` still had
  `sub.plan == "yearly" ? (price * 12 / 12.0) : price`, a no-op left over from
  before `PLAN_PRICES[:yearly]` was made monthly-equivalent. Simplified to
  return the looked-up price directly.
- **Judgment call, fixed defensively** — added explicit `tz: UTC` to the
  `refresh_sitemap` cron entry. Couldn't confirm the production container's
  system timezone from this environment, so added it rather than leave the 03:00
  UTC intent dependent on an unverified assumption.

## Verification

`ruby -c` on every edited `.rb` file; ERB files were compiled via
`ERB.new(file, trim_mode: "-").src` and syntax-checked the same way, since the
bundle isn't fully installable in this sandbox (missing `tzinfo-data`,
`simplecov`) and `bin/rails test` couldn't be run directly. A final coverage
grep (`grep -rL 'application/ld+json' app/views/**/*.html.erb`, excluding
partials/layouts/admin/dashboard/devise/mailers/tool_demos) confirmed exactly
the 10 intentionally-skipped static pages remained uncovered, and caught the
`api_reference` gap mentioned above.

## Final notes

Shipped as planned, with the `api_reference` page added as an in-flight addition
once the coverage grep surfaced it — not in the original inventory. No live
render or Google Rich Results Test pass was possible in this environment;
recommended before merge. `categories/show` and `divisions/show` were both given
`collection_json_ld` for consistency even though they list the same underlying
category record at two different URLs (`/categories/:id` vs `/:division_slug`) —
a pre-existing near-duplicate-content pattern, flagged but not restructured, out
of scope for this pass. The sitemap refresh still relies solely on sidekiq-cron;
a CI/deploy-pipeline trigger was considered but not built, since sidekiq-cron
already covers the "never regenerates" gap that mattered for the audit.

## Files Changed

- [`apps/dashboard/app/helpers/application_helper.rb`](../../apps/dashboard/app/helpers/application_helper.rb)
  — 8 new JSON-LD helpers, locale/host fixes to `website_json_ld` and
  `api_json_ld`
- [`apps/dashboard/app/lib/plan_config.rb`](../../apps/dashboard/app/lib/plan_config.rb)
  — `price_yearly_monthly`, `formatted_requests`, `formatted_rate_limit`,
  `humanize_count` fix
- [`apps/dashboard/app/jobs/refresh_sitemap_job.rb`](../../apps/dashboard/app/jobs/refresh_sitemap_job.rb)
  — new, +
  [`test/jobs/refresh_sitemap_job_test.rb`](../../apps/dashboard/test/jobs/refresh_sitemap_job_test.rb)
- [`apps/dashboard/config/sidekiq_schedule.yml`](../../apps/dashboard/config/sidekiq_schedule.yml)
  — nightly `refresh_sitemap` entry
- [`apps/dashboard/app/controllers/admin/dashboard_controller.rb`](../../apps/dashboard/app/controllers/admin/dashboard_controller.rb),
  [`apps/dashboard/app/services/analytics_revenue_service.rb`](../../apps/dashboard/app/services/analytics_revenue_service.rb)
  — consolidated onto `PlanConfig`, fixed billing-cycle MRR bug
- [`apps/dashboard/app/views/home/pricing.html.erb`](../../apps/dashboard/app/views/home/pricing.html.erb)
  — `pricing_json_ld`, sourced from `PlanConfig` instead of hardcoded literals
- 7 tool FAQ partials (`app/views/partials/tools/*/​_faq.html.erb`) —
  locale-aware `contact_path` instead of raw `href="/contact"`
- ~34 views wired with page-specific JSON-LD: 11 index/listing pages, 20 tool
  detail pages, systems/industries show pages, comparisons/examples show pages,
  `home/api_reference`, `home/glossary`, `home/about`, `home/team`, and 4
  lead-gen forms (contact, suggest-an-api, talk-to-sales, private-deployment)
- `apps/dashboard/public/core-sitemap.xml` — deleted (orphaned, retired)
