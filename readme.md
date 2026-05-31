<p align="center">
  <p align="center">
    <a href="https://requiems.xyz/?utm_source=github&utm_medium=logo" target="_blank">
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
data delivered through one unified API.

Build production-ready systems without rebuilding backend infrastructure.

[![CI](https://github.com/bobadilla-tech/requiems-api/actions/workflows/ci.yml/badge.svg)](https://github.com/bobadilla-tech/requiems-api/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/bobadilla-tech/requiems-api/graph/badge.svg?token=N3O0R9J0SN)](https://codecov.io/gh/bobadilla-tech/requiems-api)
[![Get Started](https://img.shields.io/badge/Get_Started-→-blue)](https://requiems.xyz/en)
[![Documentation](https://img.shields.io/badge/Documentation-📖-green)](https://requiems.xyz/en/apis)

## Systems

One product, four systems, each focused on a real SaaS problem.

- **Identity & Risk System** - Protect your product from fake users, fraud, and
  bad data before it reaches your database.
- **Payments Intelligence System** - Validate and enrich financial data to
  reduce failed payments and detect risky transactions.
- **Global Data System** - Power international products with accurate, real-time
  location and compliance-ready data.
- **Data Integrity System** - Clean, validate, and standardize user input across
  your entire platform.

## Engines

Need a decision, not just raw data? Use a composed engine.

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
parallel, then returns one structured result you can use immediately.

## Why Teams Use It

- **Signup protection** - Block fake accounts, bots, and abusive signups.
- **Fintech onboarding** - Validate financial data and reduce onboarding
  friction.
- **Marketplace fraud prevention** - Detect risky users and transactions in real
  time.
- **Global product support** - Handle international users with accurate data
  everywhere.

## Developer Experience

- **Live API playground** - Test every endpoint directly in the docs.
- **Copy-paste examples** - cURL, Python, JavaScript, Go, and Markdown-ready
  snippets.
- **Precise documentation** - Every parameter, response field, and error code
  documented.
- **Built for AI agents** - llms.txt, Markdown docs, and one-click examples for
  Claude, ChatGPT, and coding agents.
- **Official client libraries** - Installable clients for JavaScript,
  TypeScript, Python, Go, Ruby, C#, and more:
  [requiems-api-clients](https://github.com/bobadilla-tech/requiems-api-clients)
- **Agent Skills** - Installable skills for AI agents and copilots:
  [requiems-api-skills](https://github.com/bobadilla-tech/requiems-api-skills)

## Quick Start

Get your API key at [requiems.xyz](https://requiems.xyz), then try it out:

```bash
# Example: Protect a signup with one call
curl -X POST https://api.requiems.xyz/v1/signup/protect \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"email": "user@tempmail.io", "ip_address": "45.33.32.156"}'
```

Explore the full catalog in the [documentation](https://requiems.xyz/apis), or
start with the [systems overview](https://requiems.xyz/en/systems) to pick the
problem you want to solve.
