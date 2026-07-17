# Identity & Risk System — Engine Spec

The Identity & Risk System assesses whether a user or action is trustworthy. It
returns blocking decisions (`is_safe`) and risk scores backed by email, phone,
IP, and domain intelligence. It is designed for real-time use at registration
boundaries, account review workflows, and any step where a fraudulent actor
could cause downstream harm.

**Distinct from other systems:**

| Question                     | System                |
| ---------------------------- | --------------------- |
| Is this data real and clean? | Data Integrity        |
| Is this user trustworthy?    | **Identity & Risk**   |
| Is this payment valid?       | Payments Intelligence |
| Where is this user?          | Global Data           |

---

## Conventions

- Base path: `/v1/`
- Every response wrapped in
  `{"data": {...}, "metadata": {"timestamp": "...", "trace_id": "..."}}`
- `risk_score`: `0.0` (no risk) to `1.0` (certain risk)
- `confidence`: `0.0` to `1.0` — proportion of signals present and resolved
- `is_safe` always derived: `risk_score < 0.5 && confidence > 0.6`
- `flags` always `string[]` of machine-readable codes, never freeform

---

## Endpoints

### Engine: `POST /v1/signup/protect`

Primary composed endpoint shown in the dashboard. Full signup gate: returns a
single `is_safe` decision with `risk_score`, `confidence`, and a per-signal
breakdown. Designed for real-time use at the registration boundary.

One call replaces: email validation + disposable check + phone validation + IP
geolocation + VPN detection — plus the cross-signal country consistency check
that no individual endpoint can provide.

**Internal APIs composed (all fanned out in parallel):**

- Email Validator (`services/validation/email`)
- Phone Validation (`services/validation/phone`) — only when `phone` provided
- VPN/Proxy Detection (`services/networking/ip/vpn`) — only when `ip_address`
  provided
- IP Geolocation (`services/networking/ip/info`) — only when `ip_address`
  provided

All fields optional, but at least one must be present. More signals = higher
`confidence`.

**Request**

```json
{
  "email": "user@tempmail.io",
  "phone": "+14155552671",
  "ip_address": "45.33.32.156"
}
```

**Response**

```json
{
  "risk_score": 0.87,
  "is_safe": false,
  "confidence": 0.94,
  "flags": ["disposable_email", "vpn_detected"],
  "signals": {
    "email": {
      "valid": true,
      "disposable": true,
      "mx_valid": true,
      "suggestion": null
    },
    "phone": {
      "valid": true,
      "country": "US",
      "is_voip": false,
      "is_virtual": false
    },
    "ip": {
      "country_code": "NL",
      "is_vpn": true,
      "is_proxy": false,
      "is_tor": false,
      "is_hosting": false,
      "fraud_score": 0.72
    }
  }
}
```

`signals.*` is `null` for any field not provided in the request.

**Risk score weights (additive, cap 1.0):**

| Signal                                          | Weight               |
| ----------------------------------------------- | -------------------- |
| Email disposable                                | +0.30                |
| Email invalid / syntax failed                   | +0.40                |
| Email no MX records                             | +0.25                |
| Phone invalid                                   | +0.25                |
| Phone VoIP or virtual                           | +0.15                |
| IP is TOR                                       | +0.40                |
| IP is proxy                                     | +0.25                |
| IP is VPN                                       | +0.20                |
| IP is hosting provider                          | +0.10                |
| IP `fraud_score` contribution                   | `fraud_score × 0.30` |
| Geo mismatch: email domain country ≠ IP country | +0.20                |
| Geo mismatch: phone country ≠ IP country        | +0.10                |

`confidence` = signals resolved / signals requested (3 signals requested, all
resolved = 1.0; 1 of 3 resolved = 0.33).

**Flags:** `disposable_email`, `email_invalid`, `email_no_mx`, `phone_invalid`,
`phone_voip`, `phone_virtual`, `vpn_detected`, `proxy_detected`, `tor_detected`,
`hosting_ip`, `geo_mismatch_email_ip`, `geo_mismatch_phone_ip`

---

### `POST /v1/risk/score`

Lighter, faster version of `/v1/signup/protect`. Same inputs, same risk
computation, but no per-signal breakdown in the response. Lower latency —
suitable for background re-scoring of existing users where the detailed
breakdown is not needed.

**Request** — identical to `/v1/signup/protect`

**Response**

```json
{
  "risk_score": 0.87,
  "is_safe": false,
  "confidence": 0.94,
  "flags": ["disposable_email", "vpn_detected"]
}
```

No `signals` object. Flags and derived `is_safe` are identical to the full
endpoint.

**Internal APIs composed:** identical to `/v1/signup/protect`

---

### `POST /v1/user/verify`

Verify an existing user's identity signals. Composes email validity with domain
legitimacy (WHOIS + DNS + MX) to assess whether a registered email address
belongs to a real, established domain. Optionally includes IP re-check.

**Distinct from Data Integrity `/v1/domain/trust`:** that endpoint scores a
domain in isolation for B2B email qualification. This endpoint assesses a
specific user's identity — it combines the email address format check WITH the
domain trust signals to return a `verified` decision at the user level, not the
domain level.

**Internal APIs composed (all fanned out in parallel):**

- Email Validator (`services/validation/email`)
- WHOIS Lookup (`services/networking/whois`)
- Domain Info (`services/networking/domain`)
- MX Lookup (`services/networking/mx`)
- VPN/Proxy Detection (`services/networking/ip/vpn`) — only when `ip_address`
  provided

**Request**

```json
{
  "email": "alice@example.com",
  "ip_address": "203.0.113.1"
}
```

**Response**

```json
{
  "verified": true,
  "confidence": 0.91,
  "risk_score": 0.12,
  "flags": [],
  "signals": {
    "email": {
      "valid": true,
      "disposable": false,
      "mx_valid": true
    },
    "domain": {
      "age_days": 11231,
      "has_mx": true,
      "has_a_records": true,
      "available": false
    },
    "ip": {
      "is_vpn": false,
      "is_proxy": false,
      "fraud_score": 0.04
    }
  }
}
```

`verified` is derived:
`risk_score < 0.3 && confidence > 0.5 && email.valid &&
domain.has_mx && !domain.available`

**Risk signals:**

| Signal                 | Weight                       |
| ---------------------- | ---------------------------- |
| Email invalid          | +0.50 (floor verified=false) |
| Email disposable       | +0.30                        |
| Domain age < 30 days   | +0.25                        |
| Domain age 30–180 days | +0.10                        |
| No MX records          | +0.30                        |
| Domain not registered  | +0.50 (floor verified=false) |
| IP VPN/proxy/TOR       | +0.20                        |

---

## Go package structure

```text
apps/api/services/systems/
└── identity_risk/
    ├── router.go                        # mounts all sub-routers
    ├── signup_protect/
    │   ├── service.go                   # composes email + phone + vpn + ip/info
    │   ├── transport_http.go            # POST /v1/signup/protect
    │   └── transport_http_test.go
    ├── risk_score/
    │   ├── service.go                   # same composition as signup_protect, no breakdown
    │   ├── transport_http.go            # POST /v1/risk/score
    │   └── transport_http_test.go
    └── user_verify/
        ├── service.go                   # composes email + whois + domain + mx + optional vpn
        ├── transport_http.go            # POST /v1/user/verify
        └── transport_http_test.go
```

`risk_score/service.go` can share the scoring logic from
`signup_protect/service.go` via an internal package — do not duplicate the
weight table.

---

## Testing requirements

### `/v1/signup/protect`

| Case                          | Expected                                                            |
| ----------------------------- | ------------------------------------------------------------------- |
| All signals clean             | 200, `is_safe: true`, `risk_score < 0.5`                            |
| Disposable email only         | 200, `disposable_email` in flags, score ≥ 0.30                      |
| TOR IP                        | 200, `tor_detected` in flags, `is_safe: false`, score ≥ 0.40        |
| VoIP phone                    | 200, `phone_voip` in flags                                          |
| Email domain ≠ IP country     | 200, `geo_mismatch_email_ip` in flags                               |
| Only email provided           | 200, `signals.phone: null`, `signals.ip: null`, `confidence ≤ 0.34` |
| No fields provided            | 422                                                                 |
| All composed services stubbed | ✓                                                                   |

### `/v1/risk/score`

| Case                          | Expected                                |
| ----------------------------- | --------------------------------------- |
| Same inputs as signup/protect | 200, identical `risk_score` and `flags` |
| Response has no `signals` key | ✓                                       |
| No fields provided            | 422                                     |

### `/v1/user/verify`

| Case                            | Expected                                                       |
| ------------------------------- | -------------------------------------------------------------- |
| Valid email, established domain | 200, `verified: true`, `risk_score < 0.3`                      |
| Disposable email                | 200, `disposable_email` flag, `verified: false`                |
| New domain (< 30 days)          | 200, score elevated, `young_domain`-equivalent flag            |
| Domain not registered           | 200, `verified: false`, high score                             |
| WHOIS failure                   | 200, `whois_unavailable` flag, domain signals from DNS/MX only |
| No email provided               | 422                                                            |

---

## Open questions for engineer

1. **Shared scoring logic** — `signup_protect` and `risk_score` use identical
   risk computation. Extract to `identity_risk/internal/scorer.go` to avoid
   duplication. Confirm package visibility before implementing.
2. **Geo mismatch: email domain country** — detecting country from email domain
   requires a WHOIS lookup on the domain, which adds latency.But anyways we
   choose what makes a better service.
3. **IP fallback to caller IP** — No.

## Ticket breakdown

| # | Endpoint                  | Composes                          | Priority |
| - | ------------------------- | --------------------------------- | -------- |
| 1 | `POST /v1/risk/score`     | email + phone + vpn + ip/info     | medium   |
| 2 | `POST /v1/signup/protect` | same + full signal breakdown      | high     |
| 3 | `POST /v1/user/verify`    | email + whois + domain + mx + vpn | medium   |

Implement in order: risk/score first (simpler response shape), then
signup/protect (reuses scorer), then user/verify (different signal set).
