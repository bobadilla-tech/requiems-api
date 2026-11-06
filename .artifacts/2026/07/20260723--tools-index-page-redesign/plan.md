# Redesign /tools Index Page to Match Site Design System

Rebuilt the `/en/tools` (and `/es`, `/fr`) tools directory page — previously a
hand-rolled, off-brand "mini design system" — to use the same Tailwind
conventions, shared partials, and iconography as the rest of the dashboard app.
Done in two passes: an initial full rebuild, then a follow-up trim (drop the
bottom CTA section, add category icons) requested after review.

## Context

`apps/dashboard/app/views/tools/index.html.erb` was the only view in the whole
app carrying its own CSS-in-ERB system: a 440-line `<style>` block with
light/dark CSS custom properties, Google-Fonts-loaded `Inter` /
`JetBrains
Mono`, and bespoke classes (`.tool-card`, `.tools-hero`,
`.tools-cta`, etc.). Every other page — home, blog, systems, the `/apis`
directory, every tool `show` page — uses Tailwind utility classes exclusively,
the default system font stack, and shared partials (`partials/shared/_card`,
`_button`, and per-feature CTA partials). The page also carried a "NEW — New
tools added" top banner the user wanted gone regardless of the redesign.

The closest sibling page, `/apis`
(`apps/dashboard/app/views/apis/index.html.erb`), solves the same problem — a
sidebar-filterable grid of catalog items — with plain Tailwind and existing
helpers, and became the reference pattern.

`ToolsController` (`SUPPORTED_TOOLS`, `TOOLS_METADATA`, `CATEGORIES`) already
had Tailwind-ready per-tool icon colors (`icon_classes`, e.g.
`"bg-emerald-50
dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400"`)
sitting completely unused by the old view, which instead colored cards through
its own CSS-variable category-color system — a second, redundant color scheme
layered on top of the first.

## Approach

**Pass 1 — full rebuild.** Removed the Google Fonts `<link>` tags and the entire
`<style>` block, and deleted the "NEW" topbar markup outright (plus its
`badge`/`banner`/`eyebrow` locale keys in `en`/`es`/`fr`, all now unreferenced).
Kept the existing `icon_svg` lambda (tool-id → inline SVG path map) since it's
pure icon geometry, not styling.

Rebuilt the markup on the `/apis` pattern: a `min-h-screen flex` shell with a
sticky `<aside>` sidebar (desktop) and a horizontal-scroll pill row (mobile),
both linking to in-page `#tools-<category>` anchors rather than separate routes,
since this is one page rather than `/apis`' per-category pages. Tool cards were
rebuilt on `partials/apis_index/_api_card.html.erb`'s shape —
`bg-white dark:bg-gray-800 rounded-lg shadow-sm hover:shadow-md ... border
p-4 group`
— with the icon badge now driven by each tool's own
`TOOLS_METADATA[:icon_classes]` instead of the old CSS-variable category colors,
dropping the CSS-only hover-reveal snippet interaction that had no counterpart
anywhere else in the app. The bottom CTA was swapped for the already-existing
shared partial `partials/tools/shared/_cta.html.erb` — the same block every tool
`show` page renders — instead of a hand-rolled one. Category section headings
lost their color-dot styling in favor of the plain
`text-xl font-bold ... border-b` convention used elsewhere. Controller-side,
dropped the now-pointless `color:` field from each `CATEGORIES` entry (it only
fed the deleted CSS-variable system).

Verification here hit an infra wall: this sandbox has no seeded Postgres DB for
the dashboard app (`Database not found: requiem`), so a real HTTP round-trip via
`curl`/Puma couldn't complete. Confirmed correctness instead by rendering the
view in-process with `ApplicationController.renderer.render` (bypasses the DB
entirely since the controller action does no queries), checking the rendered
HTML for the expected 22 tool links, 6 category sections, absence of the old
banner/font references, and by running `bin/rubocop` on the controller. Also
cross-checked a `erb -x -T -` syntax warning against `git show HEAD:...` of the
original file — the same warning fired on the pre-existing code, confirming it's
a quirk of the plain `erb` CLI mishandling Rails' multi-line block-ERB
(`link_to ... do`), not a real bug.

**Pass 2 — follow-up trim.** User asked to (a) remove the bottom "One key. Every
API." CTA section entirely and (b) give the six category labels (Validation,
Text & Language, Networking & Internet, Finance, Entertainment, Developer Tools)
icons "such as our apis" — i.e. matching `/apis`' own category iconography.

For (a), deleted the `render 'partials/tools/shared/cta', ...` line added in
pass 1, and removed the now-orphaned `tools.index.cta.*` locale keys (`heading`,
`body`, `get_key`, `browse_apis`) from all three locale files — confirmed via
grep they weren't referenced anywhere else first.

For (b), rather than inventing new icons, reused what `/apis` already has:
`ApisHelper::CATEGORY_ICON_SVGS` / `CATEGORY_COLORS`, exposed via the
`category_icon_svg(id, size:)` / `category_color_classes(id)` helpers (already
global Rails helpers, no include needed). Those are keyed by the API-catalog's
own category ids (`finance`, `validation`, `networking`, `places`, `text`,
`technology`, `entertainment`, `health`), which only partially overlap the tools
page's own category keys (`validation`, `text`, `network`, `finance`,
`entertainment`, `dev`). Rather than renaming the tools page's keys — used
throughout locale files and anchor hrefs — added a small `icon_category:` field
to each `CATEGORIES` entry in `tools_controller.rb` mapping the tools key to the
matching apis-helper id (`network → networking`, `dev → technology`, the rest
1:1), threaded through `@categories` in the `index` action. Wired the icon +
color circle into all three places a category label appears: the desktop sidebar
link, the mobile pill nav, and each section `<h2>`, matching the icon-circle
treatment `partials/apis_index/_category_section.html.erb` already uses for its
own headings.

Re-verified with the same in-process render technique — confirmed the CTA copy
is fully gone, 18 category-icon `<svg>` instances render (6 categories × 3
placements), all 23 tool links still resolve (`SUPPORTED_TOOLS` actually has 23
entries, not 22 as miscounted in the pass-1 verification — recounted directly
from the array literal this time), rubocop clean, and all three locale YAMLs
still parse.

## Final Notes

- Files touched: `apps/dashboard/app/views/tools/index.html.erb` (full rewrite
  across both passes), `apps/dashboard/app/controllers/tools_controller.rb`
  (dropped `color:`, added `icon_category:`), and
  `config/locales/{en,es,fr}/tools.{en,es,fr}.yml` (removed `badge`, `banner`,
  `eyebrow`, and `cta.*` keys under `tools.index`).
- No commit was made — changes are sitting in the working tree, pending user
  review/commit.
- Verification was render-only (`ApplicationController.renderer`) plus
  rubocop/YAML syntax checks; no real browser/dark-mode/mobile-viewport check
  happened because this sandbox has no running Postgres instance seeded for the
  dashboard app. Recommended the user eyeball it in their own dev environment
  before shipping — dark mode toggle and mobile width especially.
- The `partials/tools/shared/_cta.html.erb` partial itself was left untouched
  (still used by every tool `show` page) — only the `/tools` index's call to it
  was removed.
