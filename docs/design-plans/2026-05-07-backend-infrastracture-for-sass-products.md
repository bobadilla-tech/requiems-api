# Backend Infrastructure for SaaS Products — Repositioning Plan

Requiems API has been perceived as a large API catalog . This document captures
the full strategic pivot to position it as **production-grade backend
infrastructure for SaaS products**. The goal is to compete at a higher price
point, attract SaaS founders and engineering teams (not just hobbyists), and
differentiate clearly from API aggregators.

This document is the source of truth for all go-to-market messaging, homepage
copy, navigation structure, and product naming. Engineering and GTM should
reference it when making decisions about what to build, name, or publish.

---

## Positioning Statement

> **All-in-one backend for SaaS products.** Build production-ready systems
> without rebuilding core infrastructure.

### Supporting one-liners (rotate / A-B test)

- "Ship faster with battle-tested backend systems."
- "From signup to fraud detection—handled."
- "Replace 10+ integrations with one platform."

### Mental model shift

| Before               | After                                                        |
| -------------------- | ------------------------------------------------------------ |
| API marketplace      | Backend infrastructure layer for SaaS                        |
| Large API collection | Four outcome-focused systems                                 |
| Cheap tool           | Production-grade, proprietary pipelines                      |
| Raw data             | Decisions: risk scores, validation results, enriched signals |

---

## Three-Layer Architecture (Public Mental Model)

Requiems API exposes three layers. Each serves a different audience need.

### Layer 1 — Divisions (Technical, SEO)

The existing 8 divisions map 1:1 to API categories in the technical catalog.
They are preserved for SEO value and for developers who want a precise,
category-level view.

- finance, validation, networking, places, text, technology, entertainment,
  health
- URL: `/finance`, `/validation`, etc.
- Audience: developers looking for specific endpoints
- Keep unchanged in navigation and routing

### Layer 2 — Systems (Marketing, Outcome-focused)

Systems group Divisions by business outcome. They are the primary marketing
surface — the way a SaaS founder or product engineer understands what Requiems
solves for them. Systems live at `/systems` and each system page shows the
relevant APIs from multiple divisions.

**The 5 Systems:**

#### 🔐 Identity & Risk System

> Protect your product from fake users, fraud, and bad data — before it reaches
> your database.

- Risk scoring
- Email, phone, and IP intelligence
- Domain and behavioral signals
- Powered by: validation, networking, text divisions

#### 💳 Payments Intelligence System

> Validate and enrich financial data to reduce failed payments and detect risky
> transactions.

- BIN & card validation
- IBAN / SWIFT checks
- Transaction risk signals
- Powered by: finance, networking divisions

#### 🌍 Global Data System

> Power international products with accurate, real-time location and
> compliance-ready data.

- Geocoding
- Timezones & holidays
- Address intelligence
- Powered by: places division

#### 🧠 Data Integrity System

> Clean, validate, and standardize user input across your entire platform.

- Input validation
- Text normalization
- Content moderation
- Powered by: validation, text divisions

#### 🧰 Developer Utilities

> Useful tools you shouldn't have to rebuild.

- QR codes
- Encoding & conversions
- Misc utilities
- Powered by: technology, entertainment, health divisions
- Presented as lower-priority "extras," not a core system

### Layer 3 — Engines (Composed Endpoints, Premium)

Engines are the key differentiator and the primary reason to charge premium
pricing. An Engine is a single endpoint that combines multiple Divisions /
signals into one decision, so the developer does not have to build the
orchestration themselves.

**This is the main product bet:** developers pay for the aggregation, the
scoring, and the decision — not for raw data.

**Example Engine — Signup Protection:**

```
POST /v1/engines/signup/protect
Content-Type: application/json

{ "email": "user@tempmail.io", "ip": "45.33.32.156" }
```

```json
{
  "risk_score": 0.87,
  "is_safe": false,
  "reasons": ["disposable_email", "vpn_detected"],
  "signals": {
    "email": { "disposable": true, "domain_age_days": 12 },
    "ip": { "is_vpn": true, "country": "US", "threat_level": "high" }
  }
}
```

**Planned Engines:**

| Engine           | Endpoint                            | Signals combined        |
| ---------------- | ----------------------------------- | ----------------------- |
| Signup Protect   | `POST /v1/engines/signup/protect`   | email + IP + phone      |
| Payment Validate | `POST /v1/engines/payment/validate` | BIN + IBAN + risk       |
| User Risk Score  | `POST /v1/engines/user/risk-score`  | email + IP + behavioral |
| Onboarding Check | `POST /v1/engines/onboarding/check` | identity + financial    |

---

## Homepage Copy Spec

### Hero

```
Headline:    All-in-one backend
             for SaaS products.

Subheadline: Authentication, validation, fraud detection, payments intelligence,
             and global data—delivered through one unified API.

Supporting:  Build production-ready systems without rebuilding core infrastructure.

CTA Primary:   Get API key
CTA Secondary: Explore systems
```

### Trust / Proof Strip (below subheadline, before CTAs)

Four inline badges:

- Proprietary data pipelines
- Real-time processing
- Consistent API design across all systems
- Designed for production workloads

### "What you get" — Systems Grid

Section heading: "What you get" Subheading: "Four systems that cover the
infrastructure every SaaS needs."

Each system card: icon + outcome title + 3 bullet signals + "Explore system →"
link. Developer Utilities shown as a fifth, lower-emphasis card or below the
main 4.

### "How it works" — 3 Steps

```
1. Choose a system
   Pick the problem you want to solve — identity, payments, global data, or
   data integrity.

2. Call a single endpoint
   Use high-level endpoints or compose workflows based on your needs.

3. Get decisions, not raw data
   Receive risk scores, validation results, and enriched insights — ready to
   use in your product.
```

### Engine Feature Spotlight

```
Headline: Stop bad users before they sign up

POST /v1/engines/signup/protect

Response:
{
  "risk_score": 0.87,
  "is_safe": false,
  "reasons": ["disposable_email", "vpn_detected"]
}

Sub: Combine multiple signals into a single decision—without building the
     pipeline yourself.
```

### "Why Requiems" — Differentiator Section

```
Heading: Infrastructure, not integration glue

Points:
- No reliance on third-party APIs
- Proprietary scoring and validation systems
- Unified schema across all endpoints
- Low-latency, high-throughput design
- Built for developers shipping real products
```

### Use Cases

```
SaaS signup protection   — Prevent fake accounts, bots, and abuse from day one.
Fintech onboarding       — Validate financial data and reduce onboarding friction.
Marketplace fraud        — Detect risky users and transactions in real time.
Global product support   — Handle international users with accurate data everywhere.
```

### Developer Experience

```
One API key
Consistent authentication
Predictable response formats
Copy-ready examples

Integrate once. Expand without friction.
```

### Final CTA

```
Heading:   Stop rebuilding backend infrastructure. Start shipping product.

CTA:       Get API key
           Talk to sales
```

### Footer Tagline (update)

```
Before: One API key for validation, utilities, and more, without the integration sprawl
After:  Backend infrastructure for SaaS products — delivered through one unified API
```

---

## Navigation Changes

| Item      | Current                | Target                      | Notes                           |
| --------- | ---------------------- | --------------------------- | ------------------------------- |
| Divisions | "Divisions" (megamenu) | keep as-is                  | SEO + dev audience              |
| APIs      | "APIs"                 | keep as-is                  | direct link to catalog          |
| Systems   | —                      | add "Systems" link          | new `/systems` route            |
| Engines   | —                      | add "Engines" link (future) | new `/engines` route when ready |

The nav will have: APIs · Examples · Case Studies · Divisions · Systems ·
Pricing

---

## Messaging Do's and Don'ts

### ❌ Avoid

- "APIs for everything"
- "large catalog"
- "many endpoints"
- "cheap" or "affordable" framing
- Positioning around quantity of APIs

### ✅ Emphasize

- Systems → outcomes
- Engines → decisions
- Infrastructure → reliability
- Proprietary pipelines → differentiation
- One key, consistent DX → low switching cost

---

## Implementation Roadmap

### Phase 1 — Homepage + Nav (immediate)

- [ ] Add `/systems` route, controller, and hub page (5 system cards)
- [ ] Update homepage hero copy
- [ ] Add proof strip to hero
- [ ] Replace categories grid with systems grid
- [ ] Add "How it works" section
- [ ] Add Engine spotlight section
- [ ] Add "Why Requiems" section
- [ ] Add use cases section
- [ ] Update final CTA copy
- [ ] Update footer tagline
- [ ] Add "Systems" to navbar

### Phase 2 — System Detail Pages

- [ ] Individual system detail pages at `/systems/:slug`
- [ ] Each page shows purpose, key signals, and APIs from relevant divisions
- [ ] System-specific hero, use cases, and how-it-works
- [ ] Megamenu or dropdown for Systems in navbar

### Phase 3 — Engines (Product + Marketing)

- [ ] Design Engine API contract (request/response shape)
- [ ] Implement first Engine: `POST /v1/engines/signup/protect`
- [ ] `/engines` marketing hub page
- [ ] Individual Engine detail pages
- [ ] Pricing integration (Engines as premium tier or add-on)
- [ ] Add "Engines" to navbar

### Phase 4 — Positioning Hardening

- [ ] Update About page to use infrastructure framing
- [ ] Update Pricing page description to emphasize systems/engines
- [ ] Update SEO meta titles to include "backend infrastructure" language
- [ ] Case studies reframed around systems, not individual APIs
- [ ] Update ES/FR translations for all new copy

---

## Pricing Direction

Engines justify higher price points because they:

1. Save developers from building orchestration logic
2. Deliver a decision (is_safe: true/false) not raw data
3. Create stickiness — switching means rebuilding the pipeline
