# SEO, Performance & UX Audit — Fix Pass

A technical audit of `requiemsapi.com` (`apps/dashboard`, Rails 8) triggered by a
pasted PageSpeed Insights / Lighthouse report (Performance 85–97, Accessibility
100, Best Practices 100, SEO 100, Agentic Browsing 3/3 depending on
device/throttling) and a Google `site:requiemsapi.com` search showing no
sitelinks/rich sub-page snippets, unlike a comparison site the owner pointed at
(Bobadilla Tech). Goal: fix everything controllable from this repo that moves
either the Lighthouse numbers or the search-result quality, and keep going
until subagent review rounds stop finding anything substantial.

## Context

`requiemsapi.com` went live as the dashboard/marketing host **today**
(2026-08-24), per
[`docs/plans/2026-08-23-requiemsapi-domain-role-swap.md`](2026-08-23-requiemsapi-domain-role-swap.md) —
`requiems.xyz` swapped roles to become the bare API host, and `requiemsapi.com`
is the new home for the Rails app, locale routing, and all public marketing
pages. This matters for two reasons:

1. **Sitelinks are algorithmic and take time.** Google generates sitelinks (the
   indented sub-page links with their own snippets, as in the owner's Bobadilla
   Tech example) from crawl history, internal-link structure, and site
   authority accumulated *on that exact domain*. A domain that started serving
   real content today cannot have sitelinks yet, no matter what's fixed here.
   This audit cannot manufacture that outcome — it can only fix the technical
   signals that make it possible once Google has had time to re-crawl and
   re-establish authority on the new domain.
2. **The migration left a real, live bug** (found below) that actively works
   against re-indexing on the new domain, which is worth fixing regardless of
   the sitelinks question.

## Findings

### SEO

1. **`public/robots.txt` still points at the old domain — real bug.**
   [`apps/dashboard/public/robots.txt`](../../apps/dashboard/public/robots.txt)
   hardcodes `Sitemap: https://requiems.xyz/sitemap.xml`, but
   [`config/sitemap.rb`](../../apps/dashboard/config/sitemap.rb) generates the
   sitemap with `default_host = "https://requiemsapi.com"`. The domain-swap
   plan's Phase 3 (§6, item 4) explicitly called out `robots.txt` as needing
   this edit — it did not land. Today, any crawler reading
   `requiemsapi.com/robots.txt` is told to fetch a sitemap that lives on what
   is now the bare API host, not the marketing site. This is the single
   highest-priority fix here: it's actively working against re-indexing on the
   new domain.

2. **Sitelinks search box's `SearchAction` targets the wrong query param —
   real bug.**
   [`application_helper.rb:161-169`](../../apps/dashboard/app/helpers/application_helper.rb#L161-L169)
   (`website_json_ld`) declares
   `urlTemplate: ".../apis?search={search_term_string}"`, but
   [`apis_controller.rb:11`](../../apps/dashboard/app/controllers/apis_controller.rb#L11)
   reads `params[:q]`, not `params[:search]`. If Google ever renders a
   sitelinks search box for this site (structurally correct today, but
   contingent on the same authority/time problem as sitelinks generally), a
   user's search would silently return unfiltered results. Fix: change the
   `urlTemplate` param name to `q` to match the controller (cheaper and lower
   risk than changing the controller's param name, which is public API surface
   for bookmarked/shared URLs).

3. **`organization_json_ld`'s `sameAs` is a hardcoded empty array — real gap.**
   [`application_helper.rb:137-153`](../../apps/dashboard/app/helpers/application_helper.rb#L137-L153)
   ships `"sameAs" => []` even though real profile URLs already exist in
   [`config/initializers/external_links.rb`](../../apps/dashboard/config/initializers/external_links.rb):
   `ExternalLinks::GITHUB[:organization]` and `ExternalLinks::SOCIAL[:linkedin]`.
   `sameAs` is one of the stronger signals Google uses to build an entity's
   Knowledge Graph presence (which sitelinks/knowledge-panel features draw
   from) — currently it's declared but empty, wasting the schema slot. Fix:
   populate from `ExternalLinks` instead of hardcoding `[]`.

4. **No Google Search Console verification found in-repo.** Could be
   DNS-verified (outside repo, unverifiable from here) or genuinely missing for
   the new property. Flagged as a manual/external check, not a code fix — see
   §Manual runbook.

### Performance (from the pasted Lighthouse report)

5. **GTM/`gtag.js` is the single biggest lever.** 185 KB transferred, 74.3 KB
   of it flagged as unused, and it's the source of the two longest main-thread
   tasks (107 ms + 81 ms) in the mobile trace.
   [`application.html.erb:4-15`](../../apps/dashboard/app/views/layouts/application.html.erb#L4-L15)
   loads it with a plain `async` tag in `<head>` for every signed-out visitor —
   `async` still competes for bandwidth and parses on arrival rather than
   deferring past first interaction. Fix: delay the GTM bootstrap until the
   page is idle or the user has interacted (`requestIdleCallback` with a
   `load`/interaction fallback), so analytics no longer competes with the
   critical rendering path.

6. **Three non-composited animations, all one root cause.**
   [`hero_demo_controller.js:194-200`](../../apps/dashboard/app/javascript/controllers/hero_demo_controller.js#L194-L200)
   (`_updateDots`) sets `width`, `height`, and `background-color` directly via
   `style.cssText` under `transition: all 0.3s`. Animating those properties
   forces layout + paint on every frame (non-composited), which is exactly what
   Lighthouse's "Avoid non-composited animations" insight flagged (3 elements,
   the hero-demo progress dots). Fix: keep the dots the same visual size and
   express the "active" state as `transform: scale(...)` + `opacity`/color
   changes on a `background-color` that's cheap because it isn't matched with a
   layout-affecting property change in the same transition — transform-only
   sizing is compositor-driven and doesn't force layout.

7. **Forced reflow, same file.**
   [`hero_demo_controller.js:163`](../../apps/dashboard/app/javascript/controllers/hero_demo_controller.js#L163)
   does `void el.offsetWidth` immediately after a style write, to force the
   browser to commit the "from" state before starting a CSS transition to the
   "to" state. That's a real synchronous layout read-after-write — the forced
   reflow Lighthouse measured (45–73 ms, unattributed in the report, but this
   is the only place in the codebase doing this pattern deliberately). Fix:
   replace the `offsetWidth` read with a double `requestAnimationFrame`, which
   achieves the same "commit the from-state first" effect without a forced
   synchronous layout.

8. **`flatpickr` ships an unnecessary polyfill.**
   [`config/importmap.rb`](../../apps/dashboard/config/importmap.rb) pins
   `flatpickr` from `esm.sh` with `?target=es2020`, which includes an
   `Object.assign` polyfill (10.3 KB) that no evergreen browser needs. Fix:
   bump the `target` query param to a newer baseline (`es2022`) to drop the
   polyfill from the CDN-transpiled bundle.

9. **Turbo/ActionCable unused-code flag — investigated, not fixed.**
   Lighthouse flags `turbo.min-*.js` (24.8 KB unused) including
   `@rails/actioncable`'s `connection.js` / `connection_monitor.js`. Confirmed
   by grep: **ActionCable is not used anywhere in this app**
   (`grep -rln "ActionCable" app` → no matches). The `@hotwired/turbo-rails`
   browser build always bundles Cable-Stream support as part of the single
   `turbo.min.js` file — there's no official "no cable" build to swap in
   without adopting a real JS bundler (esbuild/rollup) to tree-shake it
   ourselves, which is a meaningfully bigger infra change than anything else in
   this pass. **Recorded here as an accepted trade-off**, not fixed — flagging
   it for a future "adopt a JS bundler" project rather than bundling that scope
   into an SEO/perf audit.

10. **Cache-TTL flags on third-party/Cloudflare assets — investigated, not
    fixed.** `beacon.min.js` (Cloudflare Web Analytics) and
    `email-decode.min.js` (Cloudflare's Email Address Obfuscation feature) are
    both Cloudflare-injected, not app-controlled. Checked whether Email
    Obfuscation is actually earning its keep: confirmed real `mailto:` links
    exist in `home/security`, `home/privacy`, `home/about`, `home/terms`, and
    the contact partial — so the feature is doing real anti-scraping work, not
    dead weight. **Left as-is**, not worth trading 5 KiB of cache TTL for
    exposed email addresses.

### UX

Beyond what §1–8 already double as UX fixes (working sitelinks search box,
faster/cleaner hero animation), this pass didn't surface additional
high-confidence UX bugs from a first read of the layout/footer/nav. Review
round 1 (§Review findings log) found real ones — see #11–19 below.

## Review findings log

### Round 1

An independent subagent re-verified findings #1–10 above against live code
(not just this doc) — all ten **CONFIRMED**, including exact line numbers, so
none of the fix descriptions above needed correction. It then hunted for
additional real SEO/performance/UX issues in areas this pass hadn't looked at.
New findings, folded in and renumbered to continue the sequence above:

11. **HIGH — `divisions.en.yml` is missing 97 translation keys that `es`/`fr`
    both have; English division pages likely render visible "translation
    missing" text.**
    [`config/locales/en/divisions.en.yml`](../../apps/dashboard/config/locales/en/divisions.en.yml)
    has zero `final_cta.*` or `how_it_works.*` keys for any of the 8 division
    verticals (`finance, validation, networking, places, text, technology,
    entertainment, health` — see
    [`app/lib/division_slugs.rb`](../../apps/dashboard/app/lib/division_slugs.rb)),
    plus is missing `divisions.index.hero.secondary_cta`. Both `es` and `fr`
    locale files have all 97 keys. These are rendered unconditionally, no
    `default:` fallback, in
    [`partials/divisions/_final_cta.html.erb:4,5,8,11`](../../apps/dashboard/app/views/partials/divisions/_final_cta.html.erb)
    and
    [`_how_it_works.html.erb:19,20`](../../apps/dashboard/app/views/partials/divisions/_how_it_works.html.erb),
    called from
    [`divisions/show.html.erb:18,45`](../../apps/dashboard/app/views/divisions/show.html.erb)
    on every division page. `config/application.rb` sets `default_locale = :en`
    with `i18n.fallbacks = true`, but English has nowhere else to fall back to
    — Rails renders literal `translation missing: en.divisions.X...` spans in
    place of this content. This hits the highest-intent pages (conversion CTA
    + explainer) for English visitors and Google's crawler alike — directly
    opposed to the snippet-quality problem that triggered this audit. Fix:
    backfill the 97 keys with native English copy (not a mechanical
    translation of the es/fr strings back to en).

12. **MEDIUM-HIGH — ~10 real pages share the homepage's meta description
    (duplicate-description problem).** `page_description`
    (`application_helper.rb:397`) falls back to `t("home.index.hero.description")`
    when a page sets no `content_for(:description)`. Ten pages with unique
    titles have no description override, so Google sees the homepage's
    description for all of them: `categories/show`, `divisions/index`,
    `systems/index`, `home/changelog`, `home/error_codes`, `home/for_llms`,
    `home/privacy`, `home/security`, `home/status`, `home/terms` (all
    `.html.erb`, all under `app/views/`). `og:description`/`twitter:description`
    inherit the same fallback, so this also duplicates on social shares. Ruled
    out as false leads: `tool_demos/*` (POST-only turbo-frame fragments) and
    Devise `/users/*` views (robots.txt-disallowed) also lack descriptions but
    don't matter for SEO. Fix: add page-specific `content_for(:description,
    ...)` to each of the 10 files.

13. **MEDIUM — several tool-page meta titles/descriptions are long enough to
    truncate in SERP snippets.** Sampled `*.seo.title`/`*.seo.description` in
    `config/locales/en/tools.en.yml`: titles up to 89 chars (`timezone.seo.title`,
    line ~1389) and 86 chars (`mortgage`) against Google's practical ~55-60
    char budget; descriptions up to 193 chars (`vpn_detection.seo.description`,
    lines ~1737-1739) and 190 chars (`bin_lookup`, line ~488) against the
    ~155-160 char budget, with `random_user` (177), `sentiment_analysis` (168),
    `useragent` (157), and `timezone` (161) close behind. The homepage's own
    title (50) and description (119) are fine by comparison — this is specific
    to tool-page copy. Fix: trim the worst offenders; bundle with #12 since
    it's the same locale files.

14. **MEDIUM — `division_hub_controller.js` layout-thrashes once per slide,
    worse than the hero-demo pattern already found.** `syncSlideHeights()`
    ([`division_hub_controller.js:33-50`](../../apps/dashboard/app/javascript/controllers/division_hub_controller.js#L33-L50))
    loops over slides and, **inside the loop**, toggles the `hidden` class then
    immediately reads `el.offsetHeight` (line 45) — a write-then-read repeated
    once per slide, forcing N synchronous reflows in a row instead of the
    single one #7 already covers. Runs on `connect()` (double-rAF, correct
    technique, lines 14-16), on every debounced `resize` (lines 27-30), and
    again after `document.fonts.ready` (lines 17-19). Same root-cause family as
    #6/#7, on a page (division carousels) this pass hadn't looked at. Fix:
    measure all slide heights in one pass without alternating writes and reads
    — e.g. temporarily `position: absolute; visibility: hidden` on all slides
    and read every `offsetHeight` without toggling `hidden` in between.

15. **MEDIUM — no `prefers-reduced-motion` support anywhere in the codebase,
    and the hero-demo's infinite auto-play has no pause control.** `grep -rn
    "prefers-reduced-motion" app/` returns zero results, in JS or CSS.
    `hero_demo_controller.js`'s `_loop()` (lines 131-169) runs forever on a 5s
    interval, no exposed pause/stop, no `matchMedia` check. A WCAG 2.2.2
    (Level A) gap Lighthouse's automated a11y score doesn't check (Lighthouse
    only auto-evaluates a subset of axe-core rules and never covers
    motion-preference handling). Fix: gate the loop start on
    `!window.matchMedia('(prefers-reduced-motion: reduce)').matches`, and/or
    add a visible pause control.

16. **MEDIUM — hero-demo widget has no ARIA at all for its auto-updating
    content.** No `role`, `aria-live`, or `aria-hidden` anywhere in
    `_hero.html.erb` or `hero_demo_controller.js`. The widget cycles 3 fake
    JSON responses every ~5s, mutating text content with no live-region
    announcement and no `aria-hidden` marking it decorative. Since the content
    is simulated (not real API output), the correct fix is `aria-hidden="true"`
    on the `data-controller="hero-demo"` container
    ([`partials/home_index/_hero.html.erb:50`](../../apps/dashboard/app/views/partials/home_index/_hero.html.erb#L50)),
    not a live region.

17. **MEDIUM — navbar/mobile search is a custom combobox with none of the
    standard ARIA combobox pattern.** No `role="combobox"`, `aria-expanded`,
    `aria-controls`, `aria-activedescendant` on the input
    ([`partials/navbar/_search.html.erb:16-25`](../../apps/dashboard/app/views/partials/navbar/_search.html.erb#L16-L25)),
    no `role="listbox"`/`role="option"` on the results dropdown (lines 39-46),
    and neither `navbar_search_controller.js` nor `api_search_controller.js`
    sets these dynamically. Keyboard nav works (`handleKeydown`), so this is
    exactly the "works by keyboard, invisible to screen readers" gap Lighthouse
    can't detect. Present in both desktop and the duplicated mobile search
    markup.

18. **LOW-MEDIUM — two hardcoded English strings in the search empty-state,
    breaking es/fr parity on an always-rendered component.**
    `partials/navbar/_search.html.erb:58` (`Browse APIs →`) and `:61` (`Browse
    Examples →`) are plain text, not `t(...)`, even though the surrounding
    line (:55) correctly uses `t('navbar.search.try_different_keywords_or_browse_all')`.
    Spanish/French users with a no-results search see these two CTAs in
    English. Duplicated verbatim in the mobile menu's inline search block too.
    Fix: move both strings into the locale files.

19. **LOW — Devise's inline error summary isn't associated with individual
    invalid fields.** `devise/shared/_error_messages.html.erb` renders one
    error list at the top of the form; nothing in
    `devise/registrations/new.html.erb` sets `aria-invalid`/`aria-describedby`
    on the specific field that failed. A real WCAG 3.3.1/4.1.2 gap, weaker than
    #16/#17 since it's a full-page re-render (no `aria-live` needed).

20. **LOW — `highlight.js`'s CDN stylesheet is preloaded + noscript-loaded on
    every page, but only ~10-12 of ~55 templates use it.**
    `application.html.erb:69-70` unconditionally emits the preload/noscript
    tags in the shared layout. Actual usage (`hljs`/highlight-controller) is
    confined to `apis/show`, `blog/show`, `home/ai`, `home/for_llms`,
    `systems/show`, and a couple of code-sample partials — homepage, pricing,
    team, contact, FAQ, all 8 division pages, tools index, examples, etc. pay
    for a resource hint (and, for noscript users, a render-blocking external
    stylesheet) they never use. Same class of issue as #5 (GTM), smaller. Fix:
    move the tag into a `content_for :head` block on only the pages that
    render code blocks.

21. **LOW — `scroll_spy_controller.js` reads layout on every scroll event with
    no throttling.** `_update()` (lines 18-24) reads `target.offsetTop` for
    every tracked heading on every fired (passive) scroll event, no rAF
    batching. Used on `apis/show` and `systems/show`. Lower severity than
    #7/#14 — no preceding write in the same tick, so it's a cached-layout read,
    not a forced reflow — but still worth batching via rAF.

**Checked and ruled out** (recorded so they aren't re-investigated): image
`alt` coverage is solid (three apparent misses were multi-line `alt=`
attributes on a wrapped line, not real gaps). Internal linking/crawlability is
solid — the navbar's docs dropdown links APIs/Examples/Tools/Divisions/
Comparisons/Industries on every page. `devise.en.yml`'s 38 missing
`devise.failure.*`/`devise.sessions.*` keys (present in es/fr) are very likely
intentional — the Devise gem ships its own English defaults, and the app only
overrides es/fr because the gem doesn't localize those — flagged as a caveat,
not a finding, since gem-fallback resolution couldn't be verified live here.
`home_index.*.yml`'s `api_requiems_xyz_v1` key name still looks stale but its
*value* is correctly `requiems.xyz/v1/` in all three locales — cosmetic only.

**Risk notes on the plan's proposed fixes, raised before implementation:**

- **GTM defer (#5) — `requestIdleCallback` has no Safari/WebKit support at
  all.** The fallback (`setTimeout`, feature-detected via `'requestIdleCallback'
  in window`) needs to be the *primary* path for Safari/iOS traffic, not an
  edge case, or a meaningful fraction of visitors get no improvement.
- **GTM defer (#5) — real analytics-accuracy trade-off, worth a one-line
  callout on ship, not a blocker.** Delaying `gtag('config', ...)` until
  idle/interaction means a visitor who bounces before that fires never
  registers a pageview. Normally an accepted trade-off, but notable timing
  here specifically: the domain moved yesterday, and clean before/after
  traffic comparison is probably something the owner cares about right now.
- **`flatpickr` target bump (#8) — cannot be verified live in this
  environment** (no running server/browser here). `es2022` is a safe evergreen
  baseline, but this is the one fix that most needs a real smoke test (open a
  page with a flatpickr date-range picker) before/after, not just a syntax
  check.
- **Hero-demo animation fix (#6) — technique confirmed sound**, not just
  plausible: background-color-only transitions don't force layout, so pairing
  `transform: scale()` for size with a plain `background-color` transition
  (not combined with `width`/`height` in the same `transition: all`) is
  correct.
- **Double-rAF instead of `offsetWidth` (#7) — low-risk, has local precedent**:
  `division_hub_controller.js:14-16` already does this exact pattern
  elsewhere in this codebase for a different purpose.
- No risk found in #1–3 (robots.txt, SearchAction, sameAs) — all three are
  single-line and mechanical.

### Round 2

A second, independent read-only agent re-verified #11–21 line-by-line (not
trusting round 1's numbers) and kept hunting in areas neither prior pass had
covered (blog, comparisons, tool catalog, API docs page's interactive
widgets). Two corrections to round 1, plus 9 new findings — two of them HIGH:

**Corrections:**

- **#11's key count is exactly right (97), but its composition needs a
  precise breakdown.** Only **80** of the 97 missing keys
  (`final_cta.*` × 8 divisions + `how_it_works.*` × 8) are actually rendered
  anywhere and therefore produce a live "translation missing" bug. The
  remaining **17** (`{division}.hero.secondary_cta` × 8,
  `{division}.view_case` × 8, `divisions.index.hero.secondary_cta` × 1) are
  never referenced by any view (`grep -rn "secondary_cta\|view_case"
  app/views/divisions/ app/views/partials/divisions/` only matches
  `final_cta.secondary_cta`) — orphaned in es/fr, not a live bug in en. Still
  worth backfilling for locale-file symmetry and in case they're wired up
  later, but the fix should distinguish "these 80 are urgent" from "these 17
  are cleanup."
- **#18's "duplicated in the mobile menu" claim is wrong.** The mobile search
  block is a separate, hand-rolled implementation
  ([`layouts/_navbar.html.erb:160-205`](../../apps/dashboard/app/views/layouts/_navbar.html.erb#L160-L205)),
  not a re-render of the desktop partial — its empty state already uses
  correct `t(...)` calls and has no "Browse APIs"/"Browse Examples" CTAs at
  all. The bug is real but confined to the desktop search partial
  (`_search.html.erb:58,61`) only — smaller blast radius, same fix.

**New findings:**

22. **HIGH — the entire `/tools` catalog (index + ~29 individual tool pages ×
    3 locales, ~90 URLs) is absent from the sitemap.**
    [`config/sitemap.rb`](../../apps/dashboard/config/sitemap.rb) has
    dedicated groups for static pages, case studies, blog, divisions,
    comparisons, industries, systems, categories, examples, and API pages —
    zero mention of `tools` (`grep -n "tool" config/sitemap.rb` → nothing).
    `resources :tools, only: [:index, :show]`
    ([`routes.rb:191`](../../apps/dashboard/config/routes.rb#L191)) is a real,
    locale-scoped, fully public route (not disallowed in `robots.txt`), and
    every tool page has its own unique title/description and
    `SoftwareApplication` JSON-LD — these are built as independently-indexed
    landing pages, just never listed anywhere Google is told to look.
    Confirmed today's generated `public/sitemap*.xml` files contain zero
    `/tools` URLs. Given this whole audit started from a "no sitelinks"
    complaint, a missing sitemap section for one of the site's largest content
    types is a direct, mechanical contributor — arguably bigger in reach than
    #1. Fix: add a `TOOL_PAGES` list + `group(filename: :sitemap_tools)` block
    to `sitemap.rb`, mirroring the existing `EXAMPLES`/`CASE_STUDY_PAGES`
    pattern, and wire `sitemap_tools.xml` into the sitemap index.

23. **HIGH — the `/tools` catalog's names/descriptions/JSON-LD are hardcoded
    in English at the controller level, site-wide, silently breaking es/fr.**
    [`tools_controller.rb:6-162`](../../apps/dashboard/app/controllers/tools_controller.rb#L6-L162)'s
    `TOOLS_METADATA` hash hardcodes English `name:`/`description:` for all
    ~29 tools with no i18n. `tools/index.html.erb:109,113` render
    `tool[:name]`/`tool[:description]` directly (not `t()`) for every card;
    line 2's `content_for :description` is a hardcoded English literal (unlike
    line 1's title, which correctly uses `t(...)`). 30 of 31
    `tools/*/show.html.erb` pages source their `software_application_json_ld`
    from the same English-only hash, so the **structured data is wrong for
    every non-English tool page**, not just the visible copy. 5 of those 31
    show pages additionally hardcode their own `content_for :description` in
    English. Combined with #22, the entire tools surface — likely the
    highest-traffic conversion path after the homepage — is silently
    all-English for es/fr visitors, with no visible error the way #11 has.
    Fix: move `TOOLS_METADATA` name/description into locale files
    (`tools.catalog.<id>.{name,description}`), swap the controller/views to
    `t(...)`, fix the 5 hardcoded per-page descriptions and the index page's
    hardcoded description.

24. **MEDIUM — `partials/tools/quotes/_cta.html.erb` is uniquely hand-rolled
    in English where all 29 other tool CTA partials use `t(...)`.** Lines 15,
    20, 25 hardcode "Get your free API key" / "View API documentation" /
    "Free plan includes 1,000 requests/month." — verified against every other
    tool's `_cta.html.erb`, all of which call
    `t("tools.<id>.cta.cta_primary")`/`cta_docs`; `tools.en.yml`'s
    `quotes.cta` block only defines different keys
    (`ready_to_add_quotes_to_your_app`), confirming this partial was written
    from scratch rather than just missing translations. Fix: rewrite to match
    the shared `partials/tools/shared/_cta.html.erb` pattern, backed by real
    `tools.quotes.cta.cta_primary`/`cta_docs` locale keys.

25. **MEDIUM — `blog/index.html.erb` and `blog/show.html.erb` hardcode their
    entire page chrome in English, unlike every sibling content-hub page.**
    Meta title/description, H1, hero paragraph, empty state, "min read", and
    "Read post"/"← Back to blog" are all plain literals
    (`blog/index.html.erb:1,2,12,14,22,31,38`, `blog/show.html.erb:10,22`) —
    zero `t()` calls, even though `/es/blog` and `/fr/blog` are real
    locale-scoped routes and `case_studies/index.html.erb` /
    `comparisons/index.html.erb` fully localize identical UI chrome. (Blog
    post *body* content is intentionally English-only, file-backed Markdown
    with no locale variants — out of scope to translate.) Fix: move the
    chrome strings into locale files, matching the case_studies/comparisons
    pattern.

26. **MEDIUM — API code-example tabs have no ARIA tabs pattern, on every
    endpoint of every `apis/show` page.**
    [`code_tabs_controller.js`](../../apps/dashboard/app/javascript/controllers/code_tabs_controller.js)
    sets no ARIA attributes;
    [`partials/apis_show/_endpoint_documentation.html.erb:139-151`](../../apps/dashboard/app/views/partials/apis_show/_endpoint_documentation.html.erb#L139-L151)
    (tab buttons) has no `role="tablist"`/`role="tab"`/`aria-selected`; lines
    154-155 (panels) have no `role="tabpanel"`. Same bug category as #17, a
    different widget, firing once per code-example block across the whole API
    catalog. Fix: add the standard ARIA tabs attributes, toggling
    `aria-selected` in the controller's `switch()`.

27. **LOW-MEDIUM — homepage FAQ accordion has no ARIA, while an
    identical-purpose widget elsewhere in the same codebase does it
    correctly.**
    [`tabs_controller.js`](../../apps/dashboard/app/javascript/controllers/tabs_controller.js)
    (used only by `partials/home_index/_faq.html.erb`) never sets
    `aria-expanded`; contrast
    [`faq_accordion_controller.js:17`](../../apps/dashboard/app/javascript/controllers/faq_accordion_controller.js#L17)
    (used on `apis/show`'s FAQ) which correctly does
    `button.setAttribute("aria-expanded", ...)`. Fix: reuse
    `faq_accordion_controller.js` for the homepage FAQ, or port its toggle
    into `tabs_controller.js`.

28. **LOW-MEDIUM — screen-reader-only "Yes"/"No" text on the comparisons
    benchmark table is hardcoded English.**
    `comparisons/index.html.erb:112,117,129,134` — `<span
    class="sr-only">Yes</span>`/`No`, plain text. The codebase already has a
    working pattern for this (`badge_yes`/`badge_no` in `tools.en.yml`), so
    this is an isolated oversight, not a missing-infrastructure problem. Fix:
    wrap in the equivalent `t(...)` calls.

29. **LOW — three pages that legitimately need highlight.js load its
    stylesheet twice, from two different CDNs — extends #20, don't just
    layer a fix on top.** `apis/show.html.erb:4`, `home/for_llms.html.erb:4`,
    `systems/show.html.erb:6` each inject their own plain (non-preloaded,
    render-blocking) `<link>` for the same `github-dark.min.css` from
    `cdnjs.cloudflare.com`, on top of the global layout's
    `cdn.jsdelivr.net` preload+swap tag (#20). Fix: when implementing #20's
    "scope the tag to pages that need it," delete these 3 pages' own
    duplicate tags rather than adding page-scoping on top of what's already
    there.

30. **LOW — `dropdown_controller.js` sets no `aria-expanded`/`aria-haspopup`,
    same root cause as #17/#26.** Used on `apis/show`'s copy-page button and
    the site-wide navbar user menu / docs dropdown. Toggles `.hidden` only.
    Lower priority than #17/#26 (simpler show/hide menu, not a full
    combobox/tabs widget) but same fix category.

**Round 2 verdict**: not clean — two HIGH findings (#22, #23) comparable in
scope to round 1's biggest items, plus 7 smaller real ones and 2 corrections
to round 1. The reviewing agent's explicit recommendation: don't do a third
open-ended hunting round — #23 alone touches ~30 view files, which is exactly
the kind of larger mechanical change likely to introduce a fresh small bug.
Instead, treat implementation itself as the next round: fix everything above,
then have the next round verify the fixes (render/smoke-test a couple of
`/es/tools/:id` pages, confirm the regenerated sitemap includes tool URLs)
rather than hunting for more net-new scope. This is folded into Plan of
action below.

### Round 3 (implementation + live verification)

All 27 items in the Plan of action below were implemented. Two corrections to
earlier rounds surfaced during implementation, plus one real (already-fixed)
bug and one real (already-fixed) minor regression surfaced by an 8-agent
code-review pass over the full diff at the end.

**Corrections to round 1/2's findings:**

- **#20's "only ~10–12 of ~55 templates use highlight.js" undercount.** Actual
  usage is far broader: `home/index` (via `partials/home_index/_engine_spotlight`,
  `data-controller="highlight"` already present pre-audit) and **all ~29
  `tools/*/show.html.erb` pages** (via `partials/tools/shared/_code_sample`,
  used 1–2× per page in every `_what_it_does` partial) both render
  highlighted code — round 1 missed this because it only grepped for
  `highlight-controller`-style literal strings, not the actual
  `data-controller="highlight"` usage in the shared code-sample partial. The
  scoping fix (§20) is still correct and worth doing — pricing, team, contact,
  FAQ, all 8 division pages, tools index, examples, comparisons, industries,
  case studies, and all of `home/*`'s legal/status pages still don't need it —
  just note the "pages that need it" set is roughly 35 templates, not ~12.
  Implemented via a new `partials/shared/_highlight_head.html.erb` partial,
  opted into per-page via `content_for(:head)`, with a request-scoped
  `@highlight_head_rendered` guard inside `_code_sample.html.erb` so pages
  that render it 2× (e.g. request + response samples) don't double-emit the
  tag (verified live: `email-validator`'s show page renders exactly 2
  `github-dark.min.css` occurrences — 1 preload + 1 noscript — not 4).
- **#24's "every other tool's `_cta.html.erb` already calls `t(...)`" was
  wrong in two ways, not one.** Round 2 verified only the *heading/body*
  locals passed into `partials/tools/shared/_cta.html.erb`, not the shared
  partial's own button markup. The shared partial hardcoded
  `"Get your free API key"` / `"Read the docs"` in English, unconditionally,
  for **all 27 tools that route through it** — not just `quotes`. On top of
  that, `email_validator` and `unit_conversion`'s own `_cta.html.erb` files
  passed literal English `heading`/`body` strings too (round 2 said "every
  other tool" used `t(...)` here; these two didn't). Fix scope grew from
  "rewrite one partial" to: give the shared partial `cta_primary`/`cta_docs`
  required locals (removing the hardcoded strings), wire all 27 callers to
  pass `t("tools.<id>.cta.cta_primary"/".cta_docs")`, and backfill those two
  keys into `tools.en/es/fr.yml` for the ~19 tools that didn't already have
  them (9 tools — `inflation`, `qr_code`, `profanity_filter`, `spell_check`,
  `random_user`, `useragent`, `sudoku`, `trivia`, `email_normalizer` — already
  had these exact keys defined but orphaned/unused, the same
  defined-but-never-wired-up pattern as `divisions.en.yml`'s #11).

**Real bug found and fixed during this round's own verification (not by an
earlier round):** `partials/phone_validator/_cta.html.erb` — this tool's
partials live directly under `app/views/partials/phone_validator/`, not
nested under `partials/tools/phone_validator/` like the other ~29 tools. The
initial fix for #24 above was scripted against
`app/views/partials/tools/*/_cta.html.erb` and silently missed this file
since it doesn't match that glob. Caught by live-rendering all 93 tool pages
(31 tools × 3 locales) after implementation: `/en|es|fr/tools/phone-validator`
all 500'd with `ActionView::Template::Error (key not found: :cta_primary)`.
Fixed by applying the same `cta_primary`/`cta_docs` wiring to that file and
its locale keys; re-verified all 93 pages render 200 with zero errors.

**Real (minor) regression found and fixed by the code-review pass:** removing
the 3 duplicate highlight.js `<link>` tags (§20/§29) dropped a Subresource
Integrity hash — `systems/show.html.erb`'s own duplicate tag had
`integrity="sha512-..." crossorigin="anonymous"` pinning the cdnjs asset;
neither the global layout's original preload tag nor `apis/show`/`home/for_llms`'s
duplicate tags had one. Consolidating onto the new shared partial (which
mirrored the un-pinned original) silently removed that one page's SRI
protection. Fixed by computing a real SHA-384 hash for the jsdelivr URL
(`openssl dgst -sha384` against the live asset) and adding
`integrity`/`crossorigin` to the shared partial itself, so **all** pages using
it now get SRI, not just the one that happened to have it before.

**Live verification performed** (a local Postgres dev database was created
and schema-loaded specifically to run these — `bin/rails db:create db:schema:load`
— since none existed in this sandbox; this is separate from the project's
docker-based dev stack on port 5433):

- All 8 division pages × 3 locales (`en`/`es`/`fr`) rendered via
  `ActionDispatch::Integration::Session`: 200 status, **zero**
  `translation missing:` spans in the response body.
- All 31 tools × 3 locales (93 pages) rendered the same way: 200 status,
  zero `translation missing:` spans, zero exceptions (after the
  `phone_validator` fix above). Spot-checked that `/es/tools/timezone`
  actually contains "Zona Horaria" (not the English name) and
  `/fr/tools/timezone` contains "Fuseau Horaire".
- `/es/tools` (catalog index) renders with localized card copy (e.g.
  "Búsqueda de BIN").
- ~37 additional pages (blog, comparisons, all `home/*` legal/status pages,
  systems index, homepage, `/ai`, `/faq`) rendered clean across all 3
  locales.
- Ran `config/sitemap.rb` against a scratch output directory: confirmed
  `sitemap_tools.xml` is generated (96 links = 32 pages × 3 locales) and
  `sitemap.xml`'s index references it.
- Fetched the real `flatpickr@4.6.13` bundles for both `es2020` and `es2022`
  esm.sh targets directly (`curl`) to check §8's claimed 10.3 KB
  `Object.assign`-polyfill saving: **the claim doesn't hold for this exact
  package/version** — both targets are 50877 vs 50879 bytes (2-byte
  difference), and neither contains an `Object.assign` polyfill. The `es2022`
  bump is still correct to keep (safe, more accurate target for evergreen
  browsers, doesn't regress anything — confirmed 200 response, identical
  functionality) but it does **not** deliver the bundle-size win the original
  audit predicted. Flagged here since round 1 explicitly said this fix "most
  needs a real smoke test" and couldn't get one at the time.
- Ran an 8-agent parallel code-review pass over the full diff (all 117
  changed files). Findings besides the two above were architecture/DRY
  suggestions (CTA button copy duplicated per-tool across 3 locale files
  instead of a shared default with overrides; no shared JS `rafThrottle`
  helper despite 3 controllers now hand-rolling similar
  `requestAnimationFrame` patterns; `ToolsController#show` doesn't reuse
  `tool_data` the way `#index` does, so ~29 view files independently rebuild
  the same i18n-key string) — real but not bugs, and fixing them means
  touching another ~30-90 files at this scope for a working, already-verified
  diff. Left as follow-ups, not applied.
- One agent flagged the homepage FAQ accordion switching from
  first-item-open (`tabs_controller.js`'s old default) to
  fully-collapsed (`faq_accordion_controller.js` has no `connect()` hook) as
  a possible regression. Checked: `partials/pricing/_faq_item.html.erb`
  already uses `faq_accordion_controller.js` and already starts fully
  collapsed — this is the codebase's existing accordion convention, not a
  regression introduced here. Not changed.

## Plan of action

1. **Fix `robots.txt`** — one-line change, §1.
2. **Fix `SearchAction` param + populate `sameAs`** — both in
   `application_helper.rb`, §2–3.
3. **Fix hero-demo animation + forced reflow** — `hero_demo_controller.js`,
   §6–7.
4. **Bump `flatpickr` target** — `importmap.rb`, §8.
5. **Defer GTM bootstrap**, primary path safe for Safari (per round 1's risk
   note) — `application.html.erb`, §5.
6. **Backfill `divisions.en.yml`'s 97 missing keys** with native English copy
   — §11. Highest-impact item in this whole pass; do it early and verify by
   rendering (or at minimum ERB-compiling) all 8 division pages after.
7. **Add page-specific `content_for(:description, ...)`** to the 10 pages
   sharing the homepage's meta description — §12.
8. **Trim the worst oversized tool-page titles/descriptions** in
   `tools.en.yml` — §13.
9. **Fix `division_hub_controller.js`'s per-slide layout thrash** — §14.
10. **Add `prefers-reduced-motion` gating** to the hero-demo loop, plus a
    pause control — §15.
11. **Add `aria-hidden="true"`** to the hero-demo container — §16.
12. **Add ARIA combobox pattern** to navbar/mobile search — §17.
13. **i18n the two hardcoded search empty-state strings** — §18.
14. **Add `aria-invalid`/`aria-describedby`** to Devise form fields — §19.
15. **Scope `highlight.js`'s preload/noscript tag** to only the pages that use
    it — §20.
16. **Throttle `scroll_spy_controller.js`'s scroll handler** via rAF — §21.
17. **Add `/tools` (index + all ~29 show pages) to `config/sitemap.rb`** —
    §22. Comparable priority to §1; do this early.
18. **Move `TOOLS_METADATA` name/description into locale files**, fix
    controller + ~30 views + the 5 hardcoded per-page descriptions + the
    index page's hardcoded description — §23. Largest single change in this
    pass; do it carefully, one locale/page family at a time, and smoke-test
    (render, don't just syntax-check) a handful of `/es/tools/:id` and
    `/fr/tools/:id` pages afterward.
19. **Rewrite `partials/tools/quotes/_cta.html.erb`** to match the shared
    pattern — §24.
20. **i18n `blog/index.html.erb` and `blog/show.html.erb`'s chrome** — §25.
21. **Add ARIA tabs pattern** to `code_tabs_controller.js` +
    `_endpoint_documentation.html.erb` — §26.
22. **Fix homepage FAQ accordion ARIA** (reuse `faq_accordion_controller.js`
    or port its toggle) — §27.
23. **i18n the sr-only Yes/No text** on the comparisons table — §28.
24. **Remove the 3 duplicate highlight.js stylesheet tags** on `apis/show`,
    `home/for_llms`, `systems/show` — do this as part of §20/§29 together, not
    as two separate passes.
25. **Add `aria-expanded`/`aria-haspopup`** to `dropdown_controller.js` — §30.
26. **Implementation-as-round-3**: per round 2's explicit recommendation,
    the next round should verify the fixes above rather than hunt for
    further net-new scope — §23 alone touches ~30 files, which is exactly
    the kind of change likely to introduce a fresh small bug. After
    implementing, verify by: rendering (or ERB-compiling, if no live server)
    a sample of `/es/tools/:id` and `/fr/tools/:id` pages, confirming the
    regenerated sitemap includes `/tools` URLs, and re-running a normal
    code-review pass over the full diff. If that surfaces something
    substantial, fold it in and repeat; if not, the review is closed.
27. **Final verification pass**: syntax-check every edited file; note what
    could not be verified live in this environment (no running Rails server /
    real Lighthouse run available here) — flatpickr (§8's risk note), the
    division pages (§11), and the tools catalog (§23) most need a real
    render, not just a syntax check.

## Manual runbook (outside this repo — not part of the code diff)

- Confirm/complete Google Search Console property setup for `requiemsapi.com`
  (add property if missing, submit `sitemap.xml`, and decide the fate of the
  old `requiems.xyz` property per the domain-swap plan §9 item 2 — this audit
  doesn't repeat that decision, just flags that it's a prerequisite for
  meaningful indexing feedback going forward).
- Same for Bing Webmaster Tools, if used.
- Nothing to change in Cloudflare (Email Obfuscation is earning its keep, per
  §10).

## Explicitly out of scope

- Manufacturing Google Sitelinks directly — algorithmic, not a code change
  (§Context).
- Adopting a JS bundler to tree-shake Turbo/ActionCable — real fix for §9, but
  a separate, much larger project.
- Any live Search Console / Cloudflare dashboard changes — no credentials in
  this session; documented as a manual runbook instead, consistent with the
  domain-swap plan's own §14 pattern.

## Files changed

All 27 Plan of action items landed. ~117 files touched; grouped by area
rather than listed individually.

**SEO / crawlability**
- `public/robots.txt` — sitemap URL fixed to `requiemsapi.com` (§1).
- `app/helpers/application_helper.rb` — `SearchAction` `urlTemplate` param
  `search`→`q`; `organization_json_ld`'s `sameAs` populated from
  `ExternalLinks` (§2–3).
- `config/sitemap.rb` — new `TOOL_PAGES` list + `sitemap_tools` group,
  mirroring the existing `INDUSTRY_PAGES` pattern (§22).

**Performance**
- `app/views/layouts/application.html.erb` — GTM bootstrap deferred via
  `requestIdleCallback` with `setTimeout` as an equally-primary path for
  Safari (§5); global unconditional highlight.js tag removed (§20).
- `app/javascript/controllers/hero_demo_controller.js` — dot animation now
  `transform: scale()` + `background-color` only, no layout-affecting
  properties in the same transition (§6); `offsetWidth` forced-reflow read
  replaced with double-`requestAnimationFrame` (§7); `prefers-reduced-motion`
  gating + pause/resume control (§15).
- `app/javascript/controllers/division_hub_controller.js` — `syncSlideHeights`
  rewritten to batch all writes then all reads instead of alternating
  per-slide (§14).
- `app/javascript/controllers/scroll_spy_controller.js` — scroll handler
  throttled via `requestAnimationFrame` (§21).
- `config/importmap.rb` — `flatpickr` pin bumped to `?target=es2022` (§8; see
  round 3's live-verification note above — no measurable size win for this
  package/version, kept anyway as a safe/correct target).
- `app/views/partials/shared/_highlight_head.html.erb` (new) — highlight.js
  preload/noscript tags with a real SRI hash, opted into per-page instead of
  loaded unconditionally (§20/§29); wired into `apis/show`, `blog/show`,
  `systems/show`, `home/ai`, `home/for_llms`, `home/index` (via
  `_engine_spotlight`), and `partials/tools/shared/_code_sample.html.erb`
  (request-scoped dedup guard for pages rendering it 2×).

**Accessibility**
- `app/views/devise/registrations/new.html.erb` — `aria-invalid`/
  `aria-describedby` on all 5 form fields (§19).
- `app/views/partials/navbar/_search.html.erb`,
  `app/views/layouts/_navbar.html.erb`,
  `app/javascript/controllers/navbar_search_controller.js` — full ARIA
  combobox pattern (desktop + mobile instances, unique listbox ids) (§17).
- `app/views/partials/apis_show/_endpoint_documentation.html.erb`,
  `app/javascript/controllers/code_tabs_controller.js` — ARIA tabs pattern,
  per-endpoint unique ids (§26).
- `app/views/partials/home_index/_faq.html.erb` — migrated to
  `faq_accordion_controller.js` (matches `pricing`'s existing accordion);
  `app/javascript/controllers/tabs_controller.js` deleted as dead code once
  its only caller was migrated (§27).
- `app/javascript/controllers/dropdown_controller.js`,
  `app/views/partials/navbar/_docs_dropdown.html.erb`,
  `app/views/layouts/_navbar.html.erb`, `app/views/apis/show.html.erb` —
  `aria-expanded`/`aria-haspopup` wired via a new `trigger` target (§30).
- `app/views/partials/home_index/_hero.html.erb` — `aria-hidden="true"` on
  the simulated hero-demo content, with the pause control kept outside it so
  it stays keyboard/screen-reader reachable (§16).

**i18n**
- `config/locales/{en,es,fr}/divisions.en.yml` — 97 keys backfilled (§11; 80
  live-rendered, 17 orphaned-but-symmetric per round 2's correction).
- `config/locales/{en,es,fr}/tools.en.yml` — new `tools.catalog.<id>.{name,description}`
  namespace (31 tools); 5 tools' hardcoded per-page `content_for(:description)`
  moved into `show.description` keys; `cta_primary`/`cta_docs` added across
  all 27 tool CTA namespaces incl. `phone_validator` (§23–24, see round 3
  correction above); worst-offending `seo.title`/`seo.description` trimmed
  for 7 tools (§13).
- `app/controllers/tools_controller.rb` — `TOOLS_METADATA` stripped to
  `icon_classes` only; new `tool_data` builds localized name/description via
  `I18n.t` at request time (§23).
- ~30 `app/views/tools/*/show.html.erb` + `app/views/tools/index.html.erb` —
  `ToolsController::TOOLS_METADATA[...][...]` swapped for `t("tools.catalog...")`.
- `partials/tools/quotes/_cta.html.erb` rewritten to the shared pattern (§24);
  the shared `partials/tools/shared/_cta.html.erb` and all 26 remaining
  `_cta.html.erb` callers (incl. `partials/phone_validator/_cta.html.erb`)
  wired for localized button text.
- `config/locales/{en,es,fr}/blog.en.yml` (new files) + `blog/index.html.erb`,
  `blog/show.html.erb` — chrome strings localized (§25).
- `app/views/partials/navbar/_search.html.erb` — "Browse APIs"/"Browse
  Examples" moved to `navbar.search.browse_apis`/`browse_examples` (§18).
- `config/locales/{en,es,fr}/comparisons.en.yml` +
  `app/views/comparisons/index.html.erb` — sr-only Yes/No wrapped in
  `t("comparisons.hub.badge_yes"/"badge_no")` (§28).
- `config/locales/{en,es,fr}/home.en.yml` — new `pause_animation`/
  `resume_animation` (hero-demo), `privacy.meta_description`,
  `terms.meta_description` keys.
- 10 pages (`categories/show`, `divisions/index`, `systems/index`,
  `home/changelog`, `home/error_codes`, `home/for_llms`, `home/privacy`,
  `home/security`, `home/status`, `home/terms`) — page-specific
  `content_for(:description, ...)` added, each reusing an existing locale
  key where a good one already existed rather than inventing new copy (§12).

## Manual runbook status

Not attempted — outside the repo, no credentials in this session. Left as
documented open follow-ups (Search Console property setup + sitemap
submission, deciding `requiems.xyz`'s old property, Bing Webmaster Tools).
