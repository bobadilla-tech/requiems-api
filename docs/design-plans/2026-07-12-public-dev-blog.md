# Public Developer Blog

`/blog` was a dead stub: `HomeController#blog` was an empty action rendering a
static "coming soon" page, and `config/sitemap.rb` reserved the `/blog` slot at
priority 0.6 with nothing behind it. There was no blog engine, no `Post` model,
no markdown pipeline anywhere in the repo.

This document describes the work that replaced the stub with a real, file-based
blog engine and shipped the first post: "Building an MCP Server on Top of 87
REST Endpoints."

The goal is a developer-blog-targeting-search-rankings play — real engineering
writeups (starting with how we built our MCP server) rather than marketing copy,
to build authority as an API platform.

---

## Problem

| What existed                               | What was missing                          |
| ------------------------------------------ | ----------------------------------------- |
| `/blog` route → static "coming soon" view  | Any way to publish and list real posts    |
| `sitemap.rb` `/blog` entry at priority 0.6 | Per-post sitemap entries                  |
| `case_studies` content pattern (i18n+ERB)  | A pattern that works for code-heavy posts |
| `highlight.js` wired via Stimulus          | A markdown → HTML pipeline to feed it     |

The closest existing content pattern is `case_studies` — controller + hardcoded
slug array + ERB views + i18n copy. That pattern is a poor fit for technical
posts: long-form prose and fenced code blocks inside YAML strings is painful to
author and diff, and translating code-heavy deep-dives into `es`/`fr` has near
zero ROI for a developer audience.

---

## Decisions

**Markdown files, not ERB views or a DB-backed model.** Posts are authored as
`.md` files with YAML frontmatter, parsed at request time and memoized. No
migration, no admin UI — appropriate for a single-author technical blog that
publishes occasionally. Fast to write, git-diff friendly, real code blocks.

**English-only article body.** Site chrome (nav/footer/breadcrumbs) stays
translated via the existing i18n system; the post body itself is not translated.
The audience is developers reading English docs anyway.

**No new syntax-highlighting dependency.** The site already ships
`highlight.js` + a `github-dark` theme (preloaded in the layout) and a Stimulus
`highlight` controller (`app/javascript/controllers/highlight_controller.js`)
that runs `hljs.highlightElement` on every `pre code` inside an element tagged
`data-controller="highlight"` — used on `home/ai`, `home/for_llms`,
`systems/show`, and others. Kramdown's GFM parser already emits
`<pre><code class="language-xxx">`, which is exactly what `hljs` expects, so the
blog reuses that controller instead of adding `rouge` for server-side
highlighting.

**No `@tailwindcss/typography` plugin.** This app has no npm/JS build — Tailwind
v4 runs via the `tailwindcss-rails` gem's bundled CLI (`@import "tailwindcss"`
only, no `package.json` plugins). Markdown-rendered HTML has no classes on its
tags (`<h2>`, `<p>`, `<pre>`), so utility classes can't be hand-placed per
element the way they are in authored ERB. A `.blog-prose` block in
`app/assets/tailwind/application.css` styles the bare tags instead, matching the
site's existing dark-mode conventions.

---

## Architecture

### Content

```
apps/dashboard/content/blog/*.md
```

One file per post, e.g.
`2026-07-12-building-an-mcp-server-on-top-of-87-rest-endpoints.md`. Frontmatter:

```yaml
---
title: "..."
slug: "..."        # optional — falls back to filename minus leading date
date: 2026-07-12
author: "..."
description: "..." # meta description, OG description, index-card excerpt
---
```

### `app/models/blog_post.rb` — PORO, not ActiveRecord

- `BlogPost.all` — globs `content/blog/*.md`, splits frontmatter from body via
  `FRONTMATTER_RE`, parses frontmatter with
  `YAML.safe_load(permitted_classes:
  [Date])`, memoizes the list sorted by
  date descending.
- `BlogPost.find(slug)` — linear lookup over `.all`, `nil` if missing.
- `#to_html` —
  `Kramdown::Document.new(body, input: "GFM", hard_wrap:
  false).to_html`,
  memoized per instance.
- `#reading_time_minutes` — word count / 200, minimum 1.

**`hard_wrap: false` is load-bearing.** Kramdown's GFM input mode follows
GitHub's behavior of turning every single newline into `<br>`. Markdown source
files are word-wrapped at ~80 columns for readable diffs, so without
`hard_wrap: false` every wrapped line inside a paragraph rendered as a hard
break — paragraphs came out as a stack of one-line fragments instead of flowing
text. Caught during verification (`bin/rails runner` against the real post),
fixed by disabling hard wrap explicitly.

### `app/controllers/blog_controller.rb`

Mirrors `case_studies_controller.rb`'s shape: `index` sets
`@posts =
BlogPost.all`; `show` sets `@post = BlogPost.find(params[:slug])` and
`head :not_found` if missing.

### Routes

```ruby
get "blog", to: "blog#index"
get "blog/:slug", to: "blog#show", as: :blog_post
```

Replaces `get "blog", to: "home#blog"`. Stays inside the existing `(:locale)`
scope like every other page — locale only affects chrome, not the (English-only)
article body.

### Views

- `app/views/blog/index.html.erb` — card list: title, date, description, reading
  time, link. Follows the `case_studies/index.html.erb` card styling (rounded
  border, hover elevation, dark-mode variants).
- `app/views/blog/show.html.erb` — `content_for :title`/`:description` from post
  frontmatter (drives `page_title`/`og_meta_tags` automatically), byline
  (author, date, reading time), article body rendered inside
  `<div class="blog-prose" data-controller="highlight">`.

### SEO

- `og_meta_tags`, `seo_head_tags`, `json_ld_tags` (Organization/WebSite/
  Breadcrumb) — existing helpers, work unchanged once `content_for` slots are
  set.
- New `blog_post_json_ld(post)` in `application_helper.rb`, same shape as
  `case_study_json_ld` — schema.org `BlogPosting` with `headline`,
  `datePublished`, `author` (`Person`), `publisher`, `mainEntityOfPage`.
- `BREADCRUMB_NAMES["blog"]` already existed.

### Sitemap

`config/sitemap.rb` gained a `BLOG_POST_PAGES` array built from `BlogPost.all`,
looped the same way as `CASE_STUDY_PAGES` and `DIVISION_MARKETING_PAGES` —
locale-prefixed URLs with hreflang alternates, reduced priority for non-`en`
locales. New posts appear in the sitemap automatically without editing this file
again.

---

## Files changed

### Created

| File                                                                            | Purpose                                               |
| ------------------------------------------------------------------------------- | ----------------------------------------------------- |
| `app/models/blog_post.rb`                                                       | Frontmatter parsing, markdown rendering, reading time |
| `app/controllers/blog_controller.rb`                                            | `index`/`show`                                        |
| `app/views/blog/index.html.erb`                                                 | Post list                                             |
| `app/views/blog/show.html.erb`                                                  | Article page                                          |
| `content/blog/2026-07-12-building-an-mcp-server-on-top-of-87-rest-endpoints.md` | First post                                            |

### Modified

| File                                   | Change                                                 |
| -------------------------------------- | ------------------------------------------------------ |
| `Gemfile` / `Gemfile.lock`             | Add `kramdown`, `kramdown-parser-gfm`                  |
| `config/routes.rb`                     | Replace `home#blog` stub with `blog#index`/`blog#show` |
| `app/controllers/home_controller.rb`   | Remove empty `blog` action                             |
| `app/views/home/blog.html.erb`         | Deleted (dead stub view)                               |
| `app/helpers/application_helper.rb`    | Add `blog_post_json_ld`                                |
| `app/assets/tailwind/application.css`  | Add `.blog-prose` block                                |
| `config/sitemap.rb`                    | Add `BLOG_POST_PAGES` + generation loop                |
| `config/locales/{en,es,fr}/home.*.yml` | Remove unused `home.blog.*` "coming soon" keys         |

---

## Verification

Booted the dev server (Puma clustered mode crashes on this machine due to a
known macOS Objective-C fork-safety issue — unrelated to this change; worked
around with `OBJC_DISABLE_INITIALIZE_FORK_SAFETY=YES WEB_CONCURRENCY=0`):

- `/en/blog` → 200, lists the post with correct title/date/description/reading
  time.
- `/en/blog/building-an-mcp-server-on-top-of-87-rest-endpoints` → 200, article
  renders with correct headings/paragraphs (post-`hard_wrap` fix), fenced code
  blocks emit `<pre><code class="language-typescript">`.
- `/en/blog/nope-not-a-post` → 404.
- View-source: `<title>`, meta description, OG/Twitter tags, canonical +
  hreflang links, and the `BlogPosting` JSON-LD block all present and correct.
- `bin/rails tailwindcss:build` compiles `.blog-prose` cleanly (pre-existing,
  unrelated `@import` ordering warning for the flatpickr CDN import).
- `bin/rails sitemap:refresh` — confirmed the new post URL appears (×3 locales,
  with hreflang alternates) in `sitemap_static.xml`, then reverted the
  regenerated sitemap XML files since they're verification byproducts, not part
  of this change (their `lastmod`/ordering churns on every run).
