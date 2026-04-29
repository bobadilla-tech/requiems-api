# Requiems API — Q2 2026 Product & Engineering Roadmap

**Document type:** Planning and handoff (Confluence-ready Markdown)  
**Planning anchor:** 28 April 2026  
**Horizon:** Calendar Q2 2026 (April–June 2026)  
**Audience:** Leadership, senior engineers, junior engineers, growth (non-technical QA and copy)

---

## Executive summary

Q2 2026 is **not** a major API version bump and **does not** assume a new `/v2` HTTP surface. The Go API continues to ship under **`/v1`** behind the existing auth gateway. This quarter focuses on **growth and acquisition** (case studies, SEO landings, comparators, industry and division pages), **conversion instrumentation** (CTAs, attribution, funnel visibility), **developer distribution** (OpenAPI-backed SDKs, AI skills, MCP), and **incremental product depth** (batch APIs beyond today’s narrow implementation, RFC-level clarity for billing and errors).

Work is sized so **juniors can execute in parallel** with clear acceptance criteria, while **seniors** own architecture, gateway/OpenAPI contracts, and analytics truth. **Alexandra** (growth) participates in QA, copy, and product direction; high-intent non-signup paths route to a **strategy call** where appropriate.

---

## Program boundaries

| In scope for Q2 | Explicitly out of scope (unless reprioritized) |
|-----------------|--------------------------------------------------|
| New marketing and SEO pages on the Rails dashboard | Renaming or replacing the entire `/v1` tree with `/v2` |
| External SDK monorepo bootstrap and publish story | Full OAuth product for MCP (defer past MVP) |
| MCP + Cursor skills repos (narrow, validation-first MVP) | Building generic “any OpenAPI” MCP as the final product shape |
| Batch API **design + next implementations** after phone | Guaranteeing every catalog API has a batch twin in Q2 |
| Funnel events + reporting views | Full data warehouse rebuild |

**Single source of truth for “which APIs exist”:** [`apps/dashboard/config/api_catalog.yml`](https://github.com/bobadilla-tech/requiems-api/blob/main/apps/dashboard/config/api_catalog.yml) (today **56** public API products under **8** categories). Internal narrative docs (for example [`docs/core/v1-strategy.md`](core/v1-strategy.md)) may lag reality; prefer the catalog + Go routes when scoping.

---

## Visitor and conversion model

```mermaid
flowchart LR
  landing[Landing_and_catalog]
  hub[Case_studies_hub]
  case[Individual_case_or_SEO_tool]
  docs[Docs_and_examples]
  signup[API_key_signup]
  call[Strategy_call_cal_com]
  landing --> hub
  hub --> case
  case --> docs
  docs --> signup
  hub --> call
  case --> call
```

- **Primary conversion:** API key / signup for developers ready to integrate.  
- **Secondary conversion:** Book a strategy call with growth — [`cal.com/alexandra-flores/book-a-strategy-call`](https://cal.com/alexandra-flores/book-a-strategy-call) for industry and division CTAs, partnerships, and testimonials (per product roadmap discussions).

---

## Strategic pillars

### Pillar A — Acquisition and SEO

- **Case studies hub** — Information architecture from the roadmap spec: hero, differentiator line, problem-based filters, featured carousel (three cases + closing card), six-case grid, scale proof, how-it-works, UX flow strip, final CTA. Responsive desktop/mobile.
- **Featured client narratives** — VeriGeo and CompileStrength (and similar): dedicated pages with outcome-first copy, metrics, professional cookie consent copy where needed, and cross-links to relevant APIs.
- **Business cases listing** — “Requiems helps businesses in X industries” with short paragraphs and deep links to case pages; CTA for production users (consultation / partnership / testimonial exchange).
- **Industry pages** — One page per target industry with use cases, CTAs to the strategy call, and internal links to APIs and tools.
- **Division pages** — e.g. Validation division: common use cases, industries, “why us vs alternatives,” testimonials, CTA to get started + strategy call.
- **Comparator SEO** — Reusable template for high-intent keywords (e.g. vs incumbent validators); metadata, canonical, OG, structured data; minimum pilot set of published pages with internal links.
- **Programmatic SEO tools** — Simple, indexable pages: form + live API demo + long-form keyword-oriented content + CTA to official docs. Target **40+** tools aligned to catalog endpoints and high-intent phrases (e.g. email validation API, phone validation API). Serves both SEO and product education.
- **Technical SEO** — Build on existing sitemap and crawler work ([`docs/design-plans/2026-04-15-sitemap-crawler-fetchability-hardening.md`](design-plans/2026-04-15-sitemap-crawler-fetchability-hardening.md)): sitemap partitioning by page type, crawl health checks, indexability guardrails, optional CI or release checklist automation.

### Pillar B — Conversion and attribution

- **CTA matrix** — Rules for which pages push API key vs docs vs strategy call; consistent primary/secondary buttons.
- **UTM and events** — Instrument CTAs and key steps so growth can answer “which template and channel drove signups.”
- **Funnel reporting** — Lightweight dashboard or exportable view: SEO landing → case or tool → docs → signup; dimensions for campaign, channel, template.

### Pillar C — Developer distribution

- **OpenAPI today** — Auth gateway exposes **`GET /openapi.json`** generated from dashboard API docs YAML (`pnpm generate:openapi` in `apps/workers/auth-gateway`). Contract is already **repo-driven**; Q2 work is **governance** (always fresh on release, changelog expectations) and **downstream consumers**.
- **Multilanguage SDK monorepo** — New repo (see Appendix): TypeScript, Python, Go clients generated or wrapped from OpenAPI, with quickstarts that run from clean environments.
- **MCP server (MVP)** — Dedicated server, not an unbounded proxy: start with **validation** tools only; streamable HTTP transport; resource exposing spec/docs; read-only/bounded tools first; FastMCP or official TS SDK per engineering choice after spike.
- **Cursor / AI skills repo** — Markdown-forward skills documenting how agents should call Requiems; minimal runnable examples.

### Pillar D — API product depth (incremental under `/v1`)

- **Batch roadmap** — RFC: payload shape, partial failure semantics, idempotency, async if needed, alignment with **usage multipliers** in gateway ([`docs/core/auth-gateway.md`](core/auth-gateway.md), shared plan limits). **Today in code:** `httpx.HandleBatch` and **`POST` phone batch** under validation; most other domains are **single-item** only until implemented.
- **OpenAPI and docs parity** — Any new route must appear in dashboard `api_docs` YAML so gateway OpenAPI stays accurate.

### Pillar E — Quality and junior-friendly execution

- **Paired QA tickets** — For major dev tickets (carousel, comparator, SEO tool, CTA wiring), add QA tickets with explicit matrices (mobile/desktop, accessibility, links, events).
- **Alexandra** — QA on copy and conversion paths; input on prioritization of industries and case studies.

---

## Code-backed current state (April 2026)

Facts below are from this repository; use them when scoping tickets so work is not duplicated or mis-described.

| Area | What exists today | Implication for Q2 |
|------|-------------------|---------------------|
| **Go HTTP routes** | Only **`/v1/...`** under backend-secret middleware ([`apps/api/app/app.go`](../apps/api/app/app.go)) | New endpoints ship as **`/v1/...`** unless leadership opens a separate versioning project. |
| **OpenAPI spec** | Served at **`/openapi.json`** on the auth gateway ([`apps/workers/auth-gateway/src/index.ts`](../apps/workers/auth-gateway/src/index.ts)); generated from [`apps/workers/auth-gateway/scripts/openapi/main.ts`](../apps/workers/auth-gateway/scripts/openapi/main.ts) | SDK and MCP work should treat this spec as the **public contract**; CI should fail releases if generation is stale. |
| **Batch in Go** | `HandleBatch` helper ([`apps/api/platform/httpx/handler.go`](../apps/api/platform/httpx/handler.go)); batch route found for **phone validation** ([`apps/api/services/validation/phone/transport_http.go`](../apps/api/services/validation/phone/transport_http.go)) | Batch expansion is **greenfield** for other domains; gateway usage headers must be validated per endpoint. |
| **Rails public surface** | Routes include home, docs, pricing, APIs catalog, categories, examples, blog, contact, sitemap helpers, etc. ([`apps/dashboard/config/routes.rb`](../apps/dashboard/config/routes.rb)) | No dedicated **case studies hub**, **industry**, **division**, **comparator**, or **per-API SEO tool** routes yet — these are Q2 **net new** routes, controllers, views, and content models. |
| **API catalog** | **56** APIs, **8** categories in [`api_catalog.yml`](../apps/dashboard/config/api_catalog.yml) | SEO tool count and sitemap partitioning should align to **catalog ids** to avoid orphan pages. |
| **Edge** | Auth gateway + API management ([`docs/core/auth-gateway.md`](core/auth-gateway.md), [`docs/core/api-management.md`](core/api-management.md)) | Usage and rate limits remain edge concerns; batch billing semantics must stay consistent with workers. |

---

## Prioritized backlog (executable themes)

Priorities follow business impact: **SEO and conversion surfaces first**, then **measurement**, then **SDK/MCP**, then **batch expansion**. Ticket titles below avoid embedding calendar dates (team writing standard).

### P0 — Urgent (start or unblock in late April–May)

| Theme | Description | Example acceptance criteria |
|-------|-------------|----------------------------|
| Case studies hub IA | Implement hub structure: nav, hero, differentiator, filters, carousel, grid, scale proof, steps, UX flow, CTA | All sections render on desktop and mobile; primary CTA to signup, secondary to catalog; no blocking a11y regressions on primary paths |
| Carousel and filters | Three featured cases + closing card; problem-based filters | Keyboard usable controls; filters never produce dead-end empty states without guidance |
| Comparator template | Reusable “Requiems vs X” page with SEO and internal linking | Valid metadata + canonical + structured data; at least pilot set of live pages linked from hub or catalog |
| CTA and routing | Document and implement routing rules (signup vs call vs docs) + basic event payloads | Every CTA variant has documented destination; analytics payload includes source and CTA id |
| Funnel instrumentation (MVP) | Events from key templates into something queryable or exportable | Can answer signup rate by landing template within a tolerance; QA verifies a sample path end-to-end |

### P1 — High (May–June)

| Theme | Description | Example acceptance criteria |
|-------|-------------|----------------------------|
| Business case landings | Outcome-driven pages (fraud, deliverability, onboarding, data ops) | Shared content blocks; each page maps pain → APIs → docs CTA |
| Industry and division pages | Industry-specific and API-division pages with strategy-call CTAs | Data model or YAML-driven generation where possible; no orphan pages in sitemap |
| OpenAPI and SDK assessment | Document generation pipeline, release gates, and monorepo layout | TS/Python/Go quickstarts run from clean clone; versioning policy written |
| Batch API RFC + first expansion | Beyond phone: validation batch patterns, errors, partial success | RFC approved; at least one additional batch contract implemented or explicitly deferred with rationale |
| SEO utilities | Sitemap segments, crawl checks, indexability | Automated or checklist-based regression for critical templates |
| MCP / skills starter | Narrow MVP repos with docs and examples | Validation-only tools; spec as resource; security review on samples |

### P2 — Medium (June or carryover)

| Theme | Description |
|-------|-------------|
| AI skills expansion | More skills and scenarios beyond bootstrap |
| SEO tools at scale | Move toward **40+** live tools with shared layout and content guidelines |
| Comparator library | Scale comparators only where keyword research justifies |

---

## QA and ownership matrix (suggested)

| Work type | Typical implementer | QA / validation partner |
|-----------|----------------------|-------------------------|
| Rails new pages (case studies, SEO tools) | Junior + senior review | Alexandra (copy/conversion), rotating engineer (a11y/links) |
| Comparator and metadata | Mid or senior | Daniel / Mariana / Gustavo — rotate per sprint |
| CTA and analytics | Senior (Leonardo / Eliaz) | Mariana + Leonardo cross-QA on event integrity |
| OpenAPI / gateway | Senior | Carlos — quickstart and contract QA |
| Go batch APIs | Senior + junior pair | Leonardo — contract and edge cases |

Names align with team structure described in internal planning docs; adjust assignees to match Plane or squad reality.

---

## Dependencies

1. **Content** — Client approvals for named case studies; legal review for claims and logos.  
2. **Design** — Hub and carousel need layout decisions; this roadmap assumes **content architecture** is specified (roadmap narrative in internal planning), not final visual design.  
3. **OpenAPI** — Dashboard `api_docs` YAML must stay complete for any shipped Go route so `openapi.json` remains truthful.  
4. **External repos** — SDK and MCP may live outside the monorepo; document sync on release tags.

---

## Risks and mitigations

| Risk | Mitigation |
|------|------------|
| Junior overload (Go + Git + Docker + Rails) | Time-box spikes; one primary stack per sprint per person; senior office hours |
| SEO thin content penalties | Unique copy blocks per tool; avoid auto-generated duplicate paragraphs; human review for top keywords |
| MCP security and abuse | Read-only MVP; tool allowlist; no user secrets in examples; rate limits unchanged at edge |
| Analytics mistrust | Sampled reconciliation of events vs database signups monthly |
| Stale internal docs | Link engineers to catalog + gateway OpenAPI as SoT |

---

## Quarter phasing (high level)

| Phase | Window | Focus |
|-------|--------|--------|
| **A** | Late April – mid-May | Hub + carousel + filters; first comparator pilots; CTA matrix MVP; first funnel events |
| **B** | Mid-May – June | Industry/division templates; first wave of SEO tools; OpenAPI/SDK monorepo bootstrap; batch RFC + second batch implementation if RFC lands |
| **C** | June (and carryover) | Scale SEO tools and comparators; MCP hardening; technical SEO automation; backlog burn-down |

Exact dates for milestones live in issue tracking (for example Plane); this document stays stable for Confluence paste.

---

## Appendix — External repositories (from product planning)

- [`github.com/bobadilla-tech/requiems-api-skills`](https://github.com/bobadilla-tech/requiems-api-skills) — Cursor / agent skills  
- [`github.com/bobadilla-tech/requiems-api-mcp`](https://github.com/bobadilla-tech/requiems-api-mcp) — MCP server  
- [`github.com/bobadilla-tech/requiems-api-clients`](https://github.com/bobadilla-tech/requiems-api-clients) — Multilanguage SDK monorepo (to be populated)

## Appendix — Internal technical references

- [Auth gateway](core/auth-gateway.md) — Public entry, rate limits, usage, `openapi.json`  
- [API management](core/api-management.md) — Internal keys, analytics, usage export  
- [Rails app](core/rails-app.md) — Dashboard patterns  
- [Adding Go endpoints](core/adding-go-endpoints.md) — New `/v1` endpoints  
- [Sitemap hardening design plan](design-plans/2026-04-15-sitemap-crawler-fetchability-hardening.md)  
- [Home / social proof design plan](design-plans/2026-04-15-home-page-social-proof-ai-features.md) — Related growth surface

---

## Document maintenance

- **Owner:** Product + engineering lead (update monthly during Q2).  
- **Do not** paste API keys, Plane tokens, or webhook secrets into Confluence or this file.  
- When Q2 closes, archive this page and link to Q3 planning.
