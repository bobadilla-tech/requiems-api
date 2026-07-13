# Adding a Static Page

> route, controller action, view, i18n, and the sitemap step that keeps getting
> forgotten

A "static page" here means a simple marketing/info page with no live demo (for
that, see [Adding a Tool Page](./adding-tools.md)) — things like `/pricing`,
`/faq`, `/security`, `/ai`. This guide exists because we shipped several of
these pages without ever adding them to the sitemap, so they were invisible to
Google despite being live, linked, indexable pages.

---

## 1. Checklist

- [ ] Route in `apps/dashboard/config/routes.rb` (public scope, not under
      `dashboard`/`admin` namespaces)
- [ ] Empty action in `apps/dashboard/app/controllers/home_controller.rb` (or a
      dedicated controller if the page isn't a simple `home#` page)
- [ ] View in `apps/dashboard/app/views/home/{page}.html.erb`
- [ ] `content_for :title` and `content_for :description` at the top of the view
      — every public page needs these, see below
- [ ] Strings added to `config/locales/en|es|fr/home.en.yml` (or the relevant
      locale file) — no hardcoded English
- [ ] **Add the path to `STATIC_PAGES` in
      [apps/dashboard/config/sitemap.rb](../../apps/dashboard/config/sitemap.rb)**
      — this is the step that has been missed repeatedly
- [ ] Regenerate and commit the sitemap: `bundle exec rake sitemap:refresh` from
      `apps/dashboard/`
- [ ] Verify the new URL shows up:
      `grep -o '<loc>[^<]*/{page}/</loc>'
      apps/dashboard/public/sitemap_static.xml`

If the page is a redirect to another page (e.g. an old URL kept for
compatibility), do **not** add it to the sitemap — see §4.

---

## 2. Route + controller + view

```ruby
# config/routes.rb, inside the locale-scoped block
get "my-new-page", to: "home#my_new_page"
```

```ruby
# app/controllers/home_controller.rb
def my_new_page
end
```

```erb
<%# app/views/home/my_new_page.html.erb %>
<% content_for :title, t("home.my_new_page.title") %>
<% content_for :description, t("home.my_new_page.subheading") %>

<div class="min-h-screen ...">
  ...
</div>
```

`content_for :title` / `:description` matter because the layout's
`seo_head_tags` helper (`app/helpers/application_helper.rb:308`) falls back to
the homepage's title/description when a view doesn't set its own — silently
producing duplicate-title/description pages in Google's eyes. This actually
happened to `/docs` (fixed 2026-07).

Canonical tags and hreflang alternates (en/es/fr + x-default) are generated
automatically from `request.path` by the same helper — no per-page work needed
there.

---

## 3. The sitemap step

`config/sitemap.rb` is a hand-maintained allowlist — routes.rb existing is
**not** enough for a page to be crawlable via the sitemap. Add the path to
`STATIC_PAGES`:

```ruby
STATIC_PAGES = [
  { path: "",                changefreq: "weekly",  priority: 1.0 },
  ...
  { path: "/my-new-page",    changefreq: "monthly", priority: 0.5 },
  ...
].freeze
```

`priority`/`changefreq` are advisory to Google, not load-bearing — match a
similar existing page (a lead-gen form like `/talk-to-sales` is `0.4`, a
trust/info page like `/about` is `0.5`, a core product page is `0.7`+).

Then regenerate:

```bash
cd apps/dashboard
bundle exec rake sitemap:refresh
```

This rewrites all six `public/sitemap_*.xml` files (locales × every page in
every group), so expect a large diff even for a one-line `sitemap.rb` change —
that's expected, not a bug. See
[Sitemap & Crawler Fetchability Hardening](../design-plans/2026-04-15-sitemap-crawler-fetchability-hardening.md)
for why the output is a static pre-generated file rather than a dynamic
controller.

If the page is one of a larger group (an industry page, a comparison page, an
API page), it belongs in its own array/group instead of `STATIC_PAGES` — see
`COMPARISON_PAGES`, `INDUSTRY_PAGES`, `DIVISION_MARKETING_PAGES` in
`config/sitemap.rb` for the pattern (an allowlist module in `lib/` + `array.map`
into the same `add "/#{locale}#{path}/"` shape).

---

## 4. Pages that should NOT be in the sitemap

- Anything under `/dashboard`, `/admin` — already blocked by
  `Disallow: /*/dashboard/` / `/*/admin/` in `public/robots.txt`, and the
  authenticated layouts set `<meta name="robots" content="noindex, nofollow">`.
- A page whose controller action just `redirect_to`s somewhere else (e.g.
  `/docs` → `/apis`). Sitemap should only list the final destination URL.
- Anything gated behind auth/params that isn't meaningfully indexable.
