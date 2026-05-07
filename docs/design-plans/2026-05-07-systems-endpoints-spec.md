# Systems Layer — Endpoints Spec

This document defines every endpoint that needs to be built to power the five
Requiems Systems. These endpoints are distinct from the existing low-level
primitives (divisions). They are composed, decision-driven, and designed for
production SaaS use.

All new endpoints share the following conventions:

- Base path: `/v1/`
- Authentication: `Authorization: Bearer <api_key>` (same key as existing APIs)
- Content-Type: `application/json`
- Standard wrapper fields on every response: `request_id`, `latency_ms`
- Score fields always range `0 to 1` (0 = safe/low-risk, 1 = unsafe/high-risk)

Each system section leads with its **Engine** — the primary composed endpoint
shown in the dashboard. Supporting endpoints follow.

---

## 1. Identity & Risk System

### Engine: `POST /v1/signup/protect`

Full signup gate. Returns a single `is_safe` decision with risk score,
confidence, and flags. Designed for real-time use at the registration boundary.

**Request**

```json
{
  "email": "user@tempmail.io",
  "ip_address": "45.33.32.156",
  "phone": "+14155552671"
}
```

All fields optional, but at least one must be present. More signals = higher
confidence.

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 52,
  "risk_score": 0.87,
  "is_safe": false,
  "confidence": 0.94,
  "flags": [
    "disposable_email",
    "vpn_detected"
  ],
  "signals": {
    "email_valid": true,
    "phone_valid": false,
    "vpn_detected": true,
    "disposable_email": true
  }
}
```

**Notes**

- `is_safe: false` means the signup should be blocked or challenged (CAPTCHA,
  SMS OTP)
- `flags` is an array of machine-readable flag codes
- `is_safe` is always derived: `risk_score < 0.5 && confidence > 0.6`
- Recommended: surface `is_safe` directly to your registration controller

**Internal APIs composed**

- Email validation (`/v1/validation/email`)
- Disposable domain check (`/v1/networking/disposable/check`)
- Phone validation (`/v1/validation/phone`)
- IP geolocation (`/v1/networking/ip/lookup`)
- VPN/proxy detection (`/v1/networking/ip/vpn`)

---

### `POST /v1/risk/score`

Score a user based on the combination of email, phone, and IP signals. Returns
`risk_score` and `confidence` without the full decision envelope. Lower latency
than `/signup/protect`.

**Request**

```json
{
  "email": "user@tempmail.io",
  "phone": "+14155552671",
  "ip_address": "45.33.32.156",
  "user_agent": "Mozilla/5.0 ..."
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 48,
  "risk_score": 0.87,
  "confidence": 0.94,
  "flags": ["disposable_email", "vpn_detected"],
  "signals": {
    "email_valid": true,
    "email_disposable": true,
    "phone_valid": false,
    "vpn_detected": true,
    "ip_threat_level": "high"
  }
}
```

---

### `POST /v1/user/verify`

Verify a set of identity signals and return a structured confidence score. Lower
latency than `/signup/protect`, suitable for background re-scoring of existing
users.

**Request**

```json
{
  "email": "user@example.com",
  "ip_address": "1.2.3.4",
  "domain": "example.com"
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 38,
  "verified": true,
  "confidence": 0.91,
  "risk_score": 0.12,
  "flags": []
}
```

---

## 2. Payments Intelligence System

### Engine: `POST /v1/payment/validate`

Validate financial data (BIN, IBAN, SWIFT) and assess transaction risk in a
single call. Include whichever fields apply to the transaction.

**Request**

```json
{
  "card_bin": "424242",
  "iban": "DE89370400440532013000",
  "ip_address": "92.168.1.1",
  "billing_country": "DE"
}
```

Optional additional field: `"swift": "COBADEFFXXX"`

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 61,
  "valid": true,
  "risk_score": 0.14,
  "country_mismatch": false,
  "bank": {
    "name": "Deutsche Bank",
    "country": "DE",
    "brand": "Visa"
  },
  "signals": {
    "bin_valid": true,
    "iban_valid": true,
    "ip_country": "DE",
    "billing_country": "DE"
  }
}
```

When `swift` is provided, `signals` will also include `"swift_valid": true`.

**Internal APIs composed**

- BIN lookup (`/v1/finance/bin/lookup`)
- IBAN validation (`/v1/finance/iban/validate`)
- SWIFT validation (`/v1/finance/swift/validate`)
- IP geolocation (`/v1/networking/ip/lookup`)

---

### `POST /v1/transaction/risk`

Score a transaction for fraud risk. Use before authorizing high-value actions.

**Request**

```json
{
  "ip_address": "92.168.1.1",
  "billing_country": "US",
  "card_bin": "424242",
  "amount_usd": 299.00
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 44,
  "risk_score": 0.71,
  "is_safe": false,
  "flags": ["country_mismatch", "high_value_vpn"],
  "signals": {
    "ip_country": "RU",
    "billing_country": "US",
    "vpn_detected": true,
    "country_mismatch": true
  }
}
```

---

## 3. Global Data System

### Engine: `POST /v1/location/resolve`

Resolve an address or coordinates into enriched location data including
timezone, working days, and holiday status.

**Request**

```json
{
  "address": "Alexanderplatz 1, Berlin",
  "country_code": "DE"
}
```

OR

```json
{
  "coordinates": { "lat": 52.52, "lng": 13.405 }
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 72,
  "country": "DE",
  "city": "Berlin",
  "timezone": "Europe/Berlin",
  "utc_offset": "+02:00",
  "working_days_this_month": 21,
  "is_holiday": false,
  "coordinates": {
    "lat": 52.52,
    "lng": 13.405
  }
}
```

**Internal APIs composed**

- Geocoding (`/v1/places/geocoding/forward`)
- Timezone resolution (`/v1/places/timezone`)
- Holiday calendar (`/v1/places/holidays`)

---

### `GET /v1/timezone/resolve`

Resolve timezone from IP address or location string.

**Query params:** `ip` OR `location` (string)

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 29,
  "timezone": "America/New_York",
  "utc_offset": "-05:00",
  "dst_active": false,
  "country": "US"
}
```

---

### `GET /v1/holidays`

List public holidays for a country and year.

**Query params:** `country` (ISO 3166-1 alpha-2), `year`

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 18,
  "country": "DE",
  "year": 2026,
  "holidays": [
    { "date": "2026-01-01", "name": "New Year's Day", "type": "national" },
    { "date": "2026-04-03", "name": "Good Friday", "type": "national" }
  ]
}
```

---

## 4. Data Integrity System

### Engine: `POST /v1/input/validate`

Validate and normalize a set of user input fields in a single call.

**Request**

```json
{
  "email": "  User@EXAMPLE.COM  ",
  "phone": "004915123456789",
  "text": "This is some user-generated content"
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 41,
  "email": {
    "valid": true,
    "normalized": "user@example.com",
    "disposable": false
  },
  "phone": {
    "valid": true,
    "normalized": "+4915123456789",
    "country": "DE"
  },
  "text": {
    "is_safe": true,
    "toxicity_score": 0.01,
    "sentiment": "neutral"
  }
}
```

**Internal APIs composed**

- Email validation + normalization
- Phone validation + normalization
- Text toxicity check + sentiment

---

### `POST /v1/content/moderate`

Check a block of text for toxicity, profanity, and policy violations.

**Request**

```json
{
  "text": "This is user-generated content to moderate.",
  "language": "en"
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 55,
  "is_safe": true,
  "toxicity_score": 0.03,
  "sentiment": "neutral",
  "language": "en",
  "flags": [],
  "categories": {
    "profanity": false,
    "hate_speech": false,
    "spam": false,
    "violence": false
  }
}
```

---

### `POST /v1/text/normalize`

Clean and standardize a string: trim whitespace, fix encoding, normalize case.

**Request**

```json
{
  "text": "  héllo   WORLD  ",
  "operations": ["trim", "lowercase", "normalize_unicode"]
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 12,
  "original": "  héllo   WORLD  ",
  "normalized": "héllo world",
  "operations_applied": ["trim", "lowercase", "normalize_unicode"]
}
```

---

## 5. Developer Utilities

These endpoints are existing primitives promoted to cleaner, stable paths under
the `/v1/` namespace. No composition logic, just cleaner DX.

### Engine: `GET /v1/qr/generate`

Generate a QR code for any URL or string.

**Query params:** `content` (required), `format` (`png`|`svg`, default `png`),
`size` (pixels, default `256`)

**Request**

```
GET /v1/qr/generate
  ?content=https://requiems.xyz
  &format=png
  &size=256
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 34,
  "format": "png",
  "url": "https://cdn.requiems.xyz/qr/abc123.png",
  "size": 256,
  "expires_at": "2026-05-14T00:00:00Z"
}
```

---

### `POST /v1/encoding/base64`

Encode or decode a Base64 string.

**Request**

```json
{
  "value": "Hello, World!",
  "operation": "encode"
}
```

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 5,
  "result": "SGVsbG8sIFdvcmxkIQ==",
  "operation": "encode"
}
```

---

### `GET /v1/words/random`

Get random words with configurable count and length constraints.

**Query params:** `count` (default 5), `min_length`, `max_length`

**Response**

```json
{
  "request_id": "req_01HX...",
  "latency_ms": 8,
  "words": ["harbor", "eclipse", "monarch", "syntax", "verdant"]
}
```

---

## Standard Error Response

All endpoints return errors in this shape:

```json
{
  "request_id": "req_01HX...",
  "error": {
    "code": "INVALID_INPUT",
    "message": "email is required when phone and ip_address are not provided",
    "field": "email"
  }
}
```

| HTTP Status | Code                | When                                      |
| ----------- | ------------------- | ----------------------------------------- |
| 400         | `INVALID_INPUT`     | Missing or malformed request fields       |
| 401         | `UNAUTHORIZED`      | Missing or invalid API key                |
| 422         | `VALIDATION_FAILED` | Request is valid but fails business rules |
| 429         | `RATE_LIMITED`      | Per-minute or monthly quota exceeded      |
| 500         | `INTERNAL_ERROR`    | Unexpected server error                   |

---

## Implementation Notes

### Scoring consistency

- All `risk_score` values: `0.0` (no risk) to `1.0` (certain risk)
- All `confidence` values: `0.0` (no confidence) to `1.0` (certain)
- `is_safe` is always derived: `risk_score < 0.5 && confidence > 0.6`
- `flags` is always an array of machine-readable string codes (never `reasons`)

### Composition pattern

Each system endpoint calls internal Go service handlers directly (not HTTP). The
orchestration logic lives in the Go API (`apps/api/`) under a new `systems/`
package. Each system handler:

1. Validates the composite request
2. Fans out to internal service calls in parallel where possible
3. Runs the scoring/decision logic
4. Returns the structured system response

### Latency targets

| Tier                     | Target P50 | Target P99 |
| ------------------------ | ---------- | ---------- |
| Simple (1 signal)        | < 30ms     | < 100ms    |
| Composite (2-3 signals)  | < 60ms     | < 200ms    |
| Full system (4+ signals) | < 100ms    | < 350ms    |

### Go package structure

```
apps/api/
  systems/
    identity_risk/
      handler.go       # HTTP handler for /v1/signup/protect, /v1/risk/score, /v1/user/verify
      scorer.go        # Scoring logic
      signals.go       # Internal API fan-out
    payments/
      handler.go       # HTTP handler for /v1/payment/validate, /v1/transaction/risk
      scorer.go
      signals.go
    global_data/
      handler.go       # HTTP handler for /v1/location/resolve, /v1/timezone/resolve, /v1/holidays
      resolver.go
    data_integrity/
      handler.go       # HTTP handler for /v1/input/validate, /v1/content/moderate, /v1/text/normalize
      normalizer.go
    utilities/
      handler.go       # HTTP handler for /v1/qr/generate, /v1/encoding/base64, /v1/words/random
```

### Dashboard integration

The API catalog (`config/api_catalog.yml`) needs new entries for each system
endpoint so they appear in the developer docs and live playground. These should
be categorized under a new top-level `"systems"` category in the catalog with a
`composed: true` flag to distinguish them from primitive endpoints.
