# System Documentation Pages

The four backend systems (Identity & Risk, Payments Intelligence, Global Data,
Data Integrity) and the Developer Utilities grouping existed as live Go services
and as marketing overview pages at `/en/systems/:slug`. They had no technical
documentation: no parameter tables, no code examples, no endpoint reference.
Engineers landing on a system page saw a dark hero, a request/response demo, and
bullet points. There was no path from the API directory to the systems.

This document describes the work that created full technical documentation for
every system endpoint and wired the systems into the existing API directory.

---

## Problem

| What existed                                | What was missing                                     |
| ------------------------------------------- | ---------------------------------------------------- |
| Marketing pages at `/en/systems/:slug`      | Parameter tables, code examples, response field docs |
| Engine showcase (request/response demo)     | Full endpoint reference for supporting endpoints     |
| `/en/systems` index                         | Systems cards in `/en/apis` for discoverability      |
| Individual API docs with sidebar nav        | No sidebar nav or scroll-spy on system pages         |
| YAML doc files for 58+ individual APIs      | No YAML docs for the 5 systems                       |
| `ApisHelper#api_documentation` loading YAML | `SystemsController` not using it                     |

Additionally, the engine endpoint paths shown in the marketing showcase were
wrong — they showed `/v1/signup/protect` instead of `/v1/systems/signup/protect`
(the systems router is mounted at `/v1/systems` in `app/routes_v1.go`).

---

## Approach

**Hybrid layout** — preserve the existing dark hero and engine showcase (the
marketing hook), then transition into a full API Reference section that reuses
the same patterns as the individual API docs pages.

Two additional goals:

1. Add a **Backend Systems** section at the top of `/en/apis` so systems are
   discoverable from the API directory.
2. Fix the engine endpoint paths in all three locale files.

---

## Architecture

### Data layer

Each system gets its own YAML doc file in `config/api_docs/`, matching the exact
schema used by individual API docs. The same `ApisHelper#api_documentation`
method loads them — no new code needed for data access.

```
config/api_docs/
  identity-risk.yml
  payments-intelligence.yml
  global-data.yml
  data-integrity.yml
  developer-utilities.yml
```

YAML schema per file:

```yaml
api_id: identity-risk
api_name: Identity & Risk System
description: ...
base_url: https://api.requiems.xyz
overview:
  use_cases: [...]
  features: [...]
endpoints:
  - name: Protect Signup
    method: POST
    path: /v1/systems/signup/protect
    description: ...
    parameters:
      - { name, type, required, location, description, example }
    request_example: |
      { ... }
    response_example: |
      { "data": { ... }, "metadata": { ... } }
    response_fields:
      - { name, type, description }
    errors:
      - { code, message, description }
    code_examples:
      curl: |
      python: |
      javascript: |
      ruby: |
performance:
  p50_ms: 55
  p95_ms: 140
  p99_ms: 280
  samples: 100000
  measured_at: "2026-04-01"
faq:
  - { question, answer }
```

All endpoint paths are sourced from the Go transport files — not inferred.

### Confirmed endpoint inventory

| System                | Endpoint                                  | Method |
| --------------------- | ----------------------------------------- | ------ |
| Identity & Risk       | `/v1/systems/signup/protect`              | POST   |
| Identity & Risk       | `/v1/systems/risk/score`                  | POST   |
| Identity & Risk       | `/v1/systems/user/verify`                 | POST   |
| Payments Intelligence | `/v1/systems/payment/validate`            | POST   |
| Payments Intelligence | `/v1/systems/transaction/risk`            | POST   |
| Global Data           | `/v1/systems/location/resolve`            | POST   |
| Global Data           | `/v1/systems/timezone/from-ip/{ip}`       | GET    |
| Global Data           | `/v1/systems/business-calendar/{country}` | GET    |
| Data Integrity        | `/v1/systems/content/moderate`            | POST   |
| Data Integrity        | `/v1/systems/text/normalize`              | POST   |
| Developer Utilities   | `/v1/technology/qr/base64`                | GET    |
| Developer Utilities   | `/v1/technology/base64/encode`            | POST   |
| Developer Utilities   | `/v1/technology/base64/decode`            | POST   |
| Developer Utilities   | `/v1/technology/markdown`                 | POST   |

**Developer Utilities note:** there is no `/v1/systems/*` mount for utilities —
they are individual endpoints from the technology division. The YAML groups them
under a single doc file but each endpoint path is its full absolute path.

### Controller

`SystemsController` now includes `ApisHelper` and loads `@documentation` in the
`show` action:

```ruby
class SystemsController < ApplicationController
  include ApisHelper

  def show
    @slug = params[:system_slug]
    @system = SYSTEMS.find { |s| s[:slug] == @slug }
    head :not_found and return unless @system
    @documentation = api_documentation(@slug)
  end
end
```

`api_documentation` reads `config/api_docs/{slug}.yml` and memoizes per request.
If the file does not exist it returns `nil` — the view guards with
`if @documentation`.

### Page layout — hybrid

```
┌─────────────────────────────────────────────┐
│  Dark gradient hero                         │  ← unchanged
├─────────────────────────────────────────────┤
│  Engine showcase (dark, request/response)   │  ← unchanged
├─────────────────────────────────────────────┤
│  Purpose paragraph (light)                  │  ← unchanged
├─────────────────────────────────────────────┤
│  API Reference (gray-50 / gray-950)         │  ← new
│  ┌──────────────┬──────────────────────────┐│
│  │ Sidebar      │ Main (white card)         ││
│  │ Overview     │ Overview                  ││
│  │ Endpoints    │   Use cases + features    ││
│  │  ↳ ep1       │ API Endpoints             ││
│  │  ↳ ep2       │   _endpoint_documentation ││
│  │  ↳ ep3       │   (params, code, errors)  ││
│  │ FAQ          │ FAQ accordion             ││
│  │ Powered by   │                           ││
│  │ Performance  │                           ││
│  └──────────────┴──────────────────────────┘│
├─────────────────────────────────────────────┤
│  CTA (dark)                                 │  ← unchanged
└─────────────────────────────────────────────┘
```

The old sections 4 (signals + additional endpoints) and 5 (use cases) from the
i18n-driven layout are **removed**. Their content is superseded by the
YAML-driven API Reference section which is richer and accurate.

The `_endpoint_documentation` partial from `partials/apis_show/` is reused
unchanged — it renders parameters, live playground, response fields, code tabs,
and error codes. No new rendering code needed.

`highlight.js` stylesheet is added via `content_for :head` to syntax-highlight
the code blocks (same as `apis/show.html.erb`).

### Sidebar details

The sidebar mirrors `apis/show.html.erb`:

- Sticky nav with scroll-spy (Stimulus `scroll-spy` controller)
- Overview, Endpoints, per-endpoint links with colored method badges, FAQ
- "Powered by" division badges linking to `/en/:division_slug`
- Performance box (p50/p95/p99) when present in YAML

### Systems section on `/en/apis`

A new partial `partials/apis_index/_systems_section.html.erb` renders 5 system
cards above the "Most Popular" section in the API directory.

Each card shows:

- Color-coded dot and "System" badge (or "Extras" for developer utilities)
- System name and blurb from `systems.en.yml`
- Two signal pills from `systems.en.yml`
- Arrow icon that brightens on hover

Content is read directly from `SystemsController::SYSTEMS` (for color/slug) and
`t("systems.index.cards.#{key}.*")` (for copy). No new controller or model
needed.

Grid layout: `sm:grid-cols-2 lg:grid-cols-3` — the first 4 systems fill the grid
naturally; Developer Utilities spans `sm:col-span-2 lg:col-span-1` on the last
row.

---

## Files changed

### Created

| File                                                  | Purpose                                    |
| ----------------------------------------------------- | ------------------------------------------ |
| `config/api_docs/identity-risk.yml`                   | Full endpoint docs for 3 endpoints         |
| `config/api_docs/payments-intelligence.yml`           | Full endpoint docs for 2 endpoints         |
| `config/api_docs/global-data.yml`                     | Full endpoint docs for 3 endpoints         |
| `config/api_docs/data-integrity.yml`                  | Full endpoint docs for 2 endpoints         |
| `config/api_docs/developer-utilities.yml`             | Full endpoint docs for 4 utility endpoints |
| `views/partials/apis_index/_systems_section.html.erb` | Backend Systems cards for API directory    |

### Modified

| File                                    | Change                                                                                       |
| --------------------------------------- | -------------------------------------------------------------------------------------------- |
| `app/controllers/systems_controller.rb` | `include ApisHelper`, load `@documentation` in `show`                                        |
| `app/views/systems/show.html.erb`       | Add API Reference section; remove old signals/endpoints/use-cases sections; add highlight.js |
| `app/views/apis/index.html.erb`         | Render systems_section partial above most_popular                                            |
| `config/locales/en/systems.en.yml`      | Fix engine endpoint paths (all 5 systems)                                                    |
| `config/locales/es/systems.es.yml`      | Same path fixes                                                                              |
| `config/locales/fr/systems.fr.yml`      | Same path fixes                                                                              |

---

## Engine path fix

The i18n key `systems.{slug}.engine.endpoint_path` is displayed in the engine
showcase section on every system page. All five paths were wrong — they pointed
to paths that do not exist (e.g., `/v1/signup/protect` vs the actual
`/v1/systems/signup/protect`).

| System                | Old path               | Correct path                   |
| --------------------- | ---------------------- | ------------------------------ |
| Identity & Risk       | `/v1/signup/protect`   | `/v1/systems/signup/protect`   |
| Payments Intelligence | `/v1/payment/validate` | `/v1/systems/payment/validate` |
| Global Data           | `/v1/location/resolve` | `/v1/systems/location/resolve` |
| Data Integrity        | `/v1/input/validate`   | `/v1/systems/content/moderate` |
| Developer Utilities   | `/v1/qr/base64`        | `/v1/technology/qr/base64`     |

Fixed in all three locale files (`en`, `es`, `fr`).

---

## Design decisions

**Reuse `_endpoint_documentation` unchanged.** The existing partial handles
parameters, playground, response fields, code tabs, and error codes. Adding a
systems-specific variant would create drift. The only integration point is
`base_url` (passed from `@documentation["base_url"]`) and `index` (for anchor
IDs).

**Keep the marketing sections.** The dark hero and engine showcase are the
strongest differentiator for the systems pages — they communicate what the
system does before any technical detail. Replacing them with a pure docs layout
(like `apis/show`) would lose that context for developers arriving from search
or a link.

**YAML over i18n for technical content.** The existing signals, use_cases, and
additional endpoints were stored in i18n YAML, which is hard to keep in sync
with the backend and produces shallow docs. Moving to `api_docs/*.yml` gives the
same structure as individual API docs, enables reuse of all rendering
infrastructure, and makes it easy to add parameters, code examples, and errors.

**No new controller for the API directory systems section.**
`SystemsController::SYSTEMS` is a frozen constant — readable from any view
without going through a controller action. The partial reads it directly.

**Developer Utilities is a grouping, not a system.** It has no `/v1/systems/*`
mount. The doc file groups the most useful individual endpoints from the
technology division under one page so they are documented alongside the systems.
