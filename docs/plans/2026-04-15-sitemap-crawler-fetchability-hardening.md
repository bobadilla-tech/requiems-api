# Sitemap & Crawler Fetchability Hardening

Google Search Console reported "Couldn't fetch" / "Unknown" type for
`https://requiems.xyz/sitemap.xml` while direct `curl` requests returned valid
XML with `application/xml; charset=utf-8`. Chrome rendered the sitemap as plain
text rather than a structured view.

---

## Problem Statement

Four distinct bugs were found through investigation:

### Bug 1 — Rails appends `; charset=utf-8` to `application/xml`

When `render ..., content_type: "application/xml"` is called, Rails
automatically appends `; charset=utf-8`, producing
`Content-Type: application/xml; charset=utf-8`.

RFC 7303 §6.1 states the charset parameter **SHOULD NOT** be used with
`application/xml`; the encoding is determined from the XML declaration instead.

### Bug 2 — Caddy streaming gzip yields no `Content-Length` on HEAD requests

Caddy's `encode gzip` compresses response bodies as a streaming pass-through.
For HEAD requests there is no body to compress, so Caddy cannot compute the
compressed content-length and omits it entirely.

Google Search Console (and some other crawlers) issue a HEAD request before
fetching to check file metadata. A missing `Content-Length` contributed to GSC
reporting "Unknown" type and "Couldn't fetch".

### Bug 3 — Dynamic controller approach was fragile

The original controller used `respond_to { |f| f.xml }` with a `before_action`
for content-type. Under proxy/CDN configurations MIME negotiation could override
the intended header. The controller approach also left all content-type fighting
to be solved at request time rather than eliminating it structurally.

### Bug 4 — sitemap_generator minifies XML onto one line

After migrating to `sitemap_generator`, the gem's `SitemapFile` builder runs
`.gsub!(/\s+/, ' ')` on the XML wrapper and appends each `<url>` block with no
newlines, producing 49 KB on a single line. Chrome renders this as a wall of
unformatted text. Without an XSLT stylesheet, browsers show raw XML source
regardless of formatting.

---

## Goals

1. Serve `sitemap.xml` with exactly `Content-Type: application/xml` (no
   charset).
2. Ensure HEAD requests return a real `Content-Length` for the sitemap path.
3. Render the sitemap as a styled HTML page in browsers via XSLT.
4. Eliminate all dynamic controller logic for the sitemap — static file only.

---

## Non-Goals

1. Changing sitemap XML content or hreflang structure.
2. Introducing long-lived cache for llms/api-doc endpoints.

---

## Design

### 1. Migrate sitemap to a pre-generated static file (`sitemap_generator` gem)

Replaced the dynamic `SitemapController#sitemap` action and `GET /sitemap.xml`
route with a static `public/sitemap.xml` served by `ActionDispatch::Static`.
This eliminates charset fighting entirely — Rack's static file handler sets
`application/xml` with no charset.

The file is generated via a Rake task and committed to the repo. No database
connection required — the data source is the static `config/api_catalog.yml`.

```bash
docker exec requiem-dev-dashboard-1 bin/rails sitemap:refresh
git add apps/dashboard/public/sitemap.xml
git commit -m "chore: regenerate sitemap"
```

Gem added: `gem "sitemap_generator"` in `apps/dashboard/Gemfile`.

Config:
[apps/dashboard/config/sitemap.rb](../../apps/dashboard/config/sitemap.rb)

### 2. Post-process: pretty-print + inject XSLT stylesheet reference

`sitemap_generator` outputs minified single-line XML. A Rake task enhancement
runs after `sitemap:refresh` to:

1. Parse the generated file with `REXML::Formatters::Pretty` (2-space indent).
2. Inject `<?xml-stylesheet type='text/xsl' href='/sitemap.xsl'?>` immediately
   after the XML declaration.

```ruby
Rake::Task["sitemap:refresh"].enhance do
  require "rexml/document"
  path = Rails.root.join("public", "sitemap.xml")
  doc  = REXML::Document.new(path.read)
  fmt  = REXML::Formatters::Pretty.new(2)
  fmt.compact = true
  out  = +""
  fmt.write(doc, out)
  out.sub!(
    "<?xml version='1.0' encoding='UTF-8'?>",
    "<?xml version='1.0' encoding='UTF-8'?>\n<?xml-stylesheet type='text/xsl' href='/sitemap.xsl'?>"
  )
  path.write("#{out}\n")
end
```

Implementation:
[apps/dashboard/lib/tasks/sitemap.rake](../../apps/dashboard/lib/tasks/sitemap.rake)

### 3. XSLT stylesheet (`public/sitemap.xsl`)

A static XSL file in `public/` transforms the sitemap XML into a styled HTML
table in the browser. Chrome, Firefox, and Safari all apply XSLT client-side
when the `<?xml-stylesheet?>` PI is present. The stylesheet sorts URLs by
priority descending and trims `lastmod` to the date portion only.

Implementation:
[apps/dashboard/public/sitemap.xsl](../../apps/dashboard/public/sitemap.xsl)

### 4. Caddy — exclude `/sitemap.xml` from gzip

The Caddyfile `encode` block uses an inline `match` to skip gzip for the sitemap
path only. Everything else stays compressed. This lets Rails' static handler
return a real `Content-Length` on both GET and HEAD:

```caddy
requiems.xyz {
  encode {
    gzip
    match {
      not path /sitemap.xml
    }
  }
  reverse_proxy dashboard:80
}
```

Implementation: [infra/caddy/Caddyfile](../../infra/caddy/Caddyfile)

---

## Validation

### Must-pass Rails checks

```bash
docker exec requiem-dev-dashboard-1 bin/rails test
docker exec requiem-dev-dashboard-1 bundle exec bundler-audit
docker exec requiem-dev-dashboard-1 bin/importmap audit
docker exec requiem-dev-dashboard-1 bundle exec brakeman --no-pager
```

### Verify generated file

```bash
docker exec requiem-dev-dashboard-1 bin/rails sitemap:refresh
head -4 apps/dashboard/public/sitemap.xml
# <?xml version='1.0' encoding='UTF-8'?>
# <?xml-stylesheet type='text/xsl' href='/sitemap.xsl'?>
# <urlset ...>
#   <url>
```

### Manual verification after deploy

```bash
# Content-Type must be exactly "application/xml" — no charset
curl -sI https://requiems.xyz/sitemap.xml | grep content-type
# → content-type: application/xml

# HEAD must return a non-zero Content-Length (gzip bypassed for this path)
curl -sI https://requiems.xyz/sitemap.xml | grep content-length
# → content-length: <actual size>

# XSL file must be served
curl -sI https://requiems.xyz/sitemap.xsl | grep content-type
# → content-type: application/xml

# Body must include both PIs
curl -s https://requiems.xyz/sitemap.xml | head -3

# Valid XML
curl -s https://requiems.xyz/sitemap.xml | \
  python3 -c "import sys,xml.etree.ElementTree as ET; ET.parse(sys.stdin); print('valid')"
```

### Browser check

Open `https://requiems.xyz/sitemap.xml` in Chrome/Firefox/Safari. The browser
should render a styled HTML table listing all 150 URLs sorted by priority — not
raw XML source.

### GSC

After deploy, go to GSC → Sitemaps → click "Retest". Status should change to
"Success" within ~24 h. Use "Test Live URL" in URL Inspection to trigger an
immediate recrawl.

---

## Tradeoffs

### Why pre-generated static file instead of dynamic controller?

Eliminates the entire class of content-type header bugs. Static files served by
`ActionDispatch::Static` use Rack's MIME registry directly — no Rails rendering
pipeline, no charset appending, no format negotiation. `Content-Length` is set
by the file size. Any new API added to `api_catalog.yml` requires running
`bin/rails sitemap:refresh` and committing the result, which is an explicit,
reviewable step.

### Why XSLT instead of just pretty-printing?

Pretty-printing alone only helps human readability of the source. Chrome still
shows raw XML markup (syntax-highlighted source) regardless of indentation. XSLT
causes the browser to apply a full transformation and render the result as HTML
— the same technique used by Yoast/WordPress and other major sitemap
implementations. Google's crawler ignores the stylesheet PI and processes the
raw XML directly, so this is purely a browser UX improvement with no impact on
indexing.

### Why exclude only `/sitemap.xml` from gzip instead of disabling gzip globally?

Caddy's streaming gzip is by design — it avoids buffering large responses. The
trade-off (no `Content-Length` on HEAD) only matters for sitemap crawlers.
Disabling gzip globally would hurt page load performance for all HTML/JS/CSS.
Excluding only the sitemap path is a targeted, low-risk fix.

---

## Rollout

1. Regenerate sitemap:
   `docker exec requiem-dev-dashboard-1 bin/rails sitemap:refresh`
2. Commit `public/sitemap.xml` and `public/sitemap.xsl`.
3. Deploy `infra/caddy/Caddyfile` (Caddy reload required).
4. Deploy `apps/dashboard`.
5. Hard-refresh browser (Cmd+Shift+R) to bypass cached response.
6. Re-submit sitemap in Google Search Console.

---

## Files Changed

1. [apps/dashboard/Gemfile](../../apps/dashboard/Gemfile) — added
   `sitemap_generator`
2. [apps/dashboard/config/sitemap.rb](../../apps/dashboard/config/sitemap.rb) —
   new generator config
3. [apps/dashboard/lib/tasks/sitemap.rake](../../apps/dashboard/lib/tasks/sitemap.rake)
   — pretty-print + XSLT PI injection
4. [apps/dashboard/public/sitemap.xml](../../apps/dashboard/public/sitemap.xml)
   — pre-generated static file
5. [apps/dashboard/public/sitemap.xsl](../../apps/dashboard/public/sitemap.xsl)
   — XSLT browser stylesheet
6. [apps/dashboard/app/controllers/sitemap_controller.rb](../../apps/dashboard/app/controllers/sitemap_controller.rb)
   — removed `sitemap` action
7. [apps/dashboard/config/routes.rb](../../apps/dashboard/config/routes.rb) —
   removed `GET /sitemap.xml` route
8. [apps/dashboard/test/controllers/sitemap_controller_test.rb](../../apps/dashboard/test/controllers/sitemap_controller_test.rb)
   — removed sitemap action test
9. [apps/dashboard/app/views/sitemap/sitemap.xml.erb](../../apps/dashboard/app/views/sitemap/)
   — deleted
10. [infra/caddy/Caddyfile](../../infra/caddy/Caddyfile) — exclude sitemap from
    gzip
