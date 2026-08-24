<p align="center">
  <p align="center">
    <a href="https://requiemsapi.com/?utm_source=github&utm_medium=logo" target="_blank">
      <img src="https://raw.githubusercontent.com/bobadilla-tech/requiems-api/refs/heads/main/apps/dashboard/app/assets/images/logo.png" alt="Requiems API" width="280" />
    </a>
  </p>
  <p align="center">
    All-in-one backend for SaaS products.
  </p>
  <p align="center">
    <i>A product by <a href="https://bobadilla.tech">Bobadilla Technologies</a></i>
  </p>
</p>

# Requiems API

Authentication, validation, fraud detection, payments intelligence, and global
data behind one API key.

Stop rebuilding the same backend plumbing for every SaaS product.

[![CI](https://github.com/bobadilla-tech/requiems-api/actions/workflows/ci.yml/badge.svg)](https://github.com/bobadilla-tech/requiems-api/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/bobadilla-tech/requiems-api/graph/badge.svg?token=N3O0R9J0SN)](https://codecov.io/gh/bobadilla-tech/requiems-api)
[![Get Started](https://img.shields.io/badge/Get_Started-→-blue)](https://requiemsapi.com/en)
[![Documentation](https://img.shields.io/badge/Documentation-📖-green)](https://requiemsapi.com/en/apis)

## Systems

Four systems, each built around a problem SaaS teams actually hit.

- **Identity & Risk** - keep fake users, fraud, and bad data out before they
  reach your database.
- **Payments Intelligence** - validate and enrich financial data to cut failed
  payments and catch risky transactions.
- **Global Data** - accurate, real-time location and compliance data for
  international products.
- **Data Integrity** - clean, validate, and standardize user input across your
  platform.

## Engines

Need a decision, not just raw data? Compose an engine.

```
POST /v1/signup/protect
{
  "email": "user@tempmail.io",
  "ip_address": "45.33.32.156",
  "phone": "+14155552671"
}
```

```json
{
  "risk_score": 0.87,
  "is_safe": false,
  "confidence": 0.94,
  "flags": ["disposable_email", "vpn_detected"],
  "signals": {
    "email_valid": true,
    "phone_valid": false,
    "vpn_detected": true,
    "disposable_email": true
  }
}
```

Each engine fans out across validation, networking, and intelligence APIs in
parallel, then hands back one structured result.

## Where It Fits

- **Signup protection** - block fake accounts, bots, and abusive signups.
- **Fintech onboarding** - validate financial data and cut onboarding friction.
- **Marketplace fraud prevention** - catch risky users and transactions as they
  happen.
- **Global products** - handle international users with data that's actually
  accurate.

## Developer Experience

- **Live API playground** - test every endpoint straight from the docs.
- **Copy-paste examples** - cURL, Python, JavaScript, Go, and Markdown ready to
  drop in.
- **Precise documentation** - every parameter, response field, and error code
  written down.
- **Built for AI agents** - llms.txt, Markdown docs, and one-click examples for
  Claude, ChatGPT, and coding agents.
- **Official client libraries** - JavaScript, TypeScript, Python, Go, Ruby, C#,
  and more:
  [requiems-api-clients](https://github.com/bobadilla-tech/requiems-api-clients)
- **MCP server & Agent Skills** - connect via MCP or install skills straight
  into Claude and other coding agents:
  [requiemsapi.com/ai](https://requiemsapi.com/en/ai)

## Quick Start

Grab an API key at [requiemsapi.com](https://requiemsapi.com), then try it:

```bash
# Example: Protect a signup with one call
curl -X POST https://requiems.xyz/v1/signup/protect \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"email": "user@tempmail.io", "ip_address": "45.33.32.156"}'
```

Explore the full catalog in the [documentation](https://requiemsapi.com/en/apis),
or start with the [systems overview](https://requiemsapi.com/en/systems) to pick
the problem you want to solve.
