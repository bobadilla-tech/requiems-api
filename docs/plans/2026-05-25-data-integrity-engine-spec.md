# Data Integrity System — Engine Spec

The Data Integrity System validates, normalizes, and scores user-supplied data
so applications can trust their own databases. The system is for catching bad
data at ingest (forms, APIs, imports), cleaning existing records, and assessing
whether business-critical fields are real and safe.

Unlike the Identity & Risk System (which gates signups) and the Payments
Intelligence System (which gates transactions), the Data Integrity System is
neutral, it returns structured facts and scores without rendering a blocking
decision. The caller decides what to do.

**Primary use cases:**

- CRM / mailing list hygiene: validate and normalize contact records in bulk
- SaaS onboarding forms: real-time field validation beyond syntax checks
- B2B lead qualification: verify that an email domain is a real business
- Content moderation: check user-generated text before persistence
- Data pipeline quality gates: score import batches before writing to production

## Conventions

- Base path: `/v1/`
- All scores: `0.0` (no risk / highest quality) to `1.0` (certain risk / worst
  quality)
- All quality/trust scores: `0.0` (unusable) to `1.0` (verified good)
- Every response wrapped in
  `{"data": {...}, "metadata": {"timestamp": "...", "trace_id": "..."}}`
- `flags` is always a `string[]` of machine-readable codes, never freeform text
- Score 0-1, never 0-100

## Endpoints

### Engine: `POST /v1/input/validate`

Primary composed endpoint shown in the dashboard. Validates and normalizes a
contact record (email + phone) and optionally checks user-generated text for
safety. One call replaces up to four individual API round-trips.

All fields are optional but at least one must be present.

**Internal APIs composed (all fanned out in parallel):**

- Email Validator (`services/validation/email`)
- Phone Validation (`services/validation/phone`)
- Profanity Filter (`services/validation/profanity`) — only when `text` is
  provided
- Sentiment Analysis (`services/text/sentiment`) — only when `text` is provided

**Request**

```json
{
  "email": "  User+tag@Googlemail.COM  ",
  "phone": "004915123456789",
  "text": "This is some user-generated content"
}
```

**Response**

```json
{
  "email": {
    "valid": true,
    "normalized": "user@gmail.com",
    "original": "  User+tag@Googlemail.COM  ",
    "syntax_valid": true,
    "mx_valid": true,
    "disposable": false,
    "suggestion": null,
    "quality_score": 0.97
  },
  "phone": {
    "valid": true,
    "normalized": "+4915123456789",
    "country": "DE",
    "type": "mobile",
    "carrier": { "name": "Telekom" },
    "risk": { "is_voip": false, "is_virtual": false },
    "quality_score": 0.95
  },
  "text": {
    "is_safe": true,
    "toxicity_score": 0.01,
    "sentiment": "neutral",
    "flags": []
  },
  "overall_quality_score": 0.96
}
```

`overall_quality_score` is the weighted mean of present signals: email (weight
0.5) + phone (weight 0.4) + text safety inverse of toxicity (weight 0.1). Omit
omitted fields from calculation.

**Quality score breakdown (email):**

| Condition               | Penalty           |
| ----------------------- | ----------------- |
| Invalid syntax          | -1.0 (floor at 0) |
| Invalid MX              | -0.6              |
| Disposable domain       | -0.5              |
| No normalization change | +0 (no bonus)     |
| Has typo suggestion     | -0.1              |

**Quality score breakdown (phone):**

| Condition                     | Penalty           |
| ----------------------------- | ----------------- |
| Invalid number                | -1.0 (floor at 0) |
| VoIP or virtual               | -0.3              |
| Unknown type                  | -0.1              |
| Landline (lower reachability) | -0.05             |

**Flags (email):** `email_invalid`, `email_syntax_invalid`, `email_mx_invalid`,
`email_disposable`, `email_has_suggestion`

**Flags (phone):** `phone_invalid`, `phone_voip`, `phone_virtual`,
`phone_unknown_type`

**Flags (text):** `text_profanity`, `text_toxic`, `text_hate_speech`,
`text_spam`

---

### `POST /v1/input/validate/batch`

Batch variant. Process up to 50 contact records in a single call. Per-record
errors are in-band (same pattern as email/phone batch endpoints).

**Request**

```json
{
  "items": [
    { "email": "alice@example.com", "phone": "+14155552671" },
    { "email": "bad@@email", "phone": null }
  ]
}
```

**Response**

```json
{
  "results": [
    {
      "index": 0,
      "email": {
        "valid": true,
        "normalized": "alice@example.com",
        "quality_score": 0.97,
        "disposable": false
      },
      "phone": {
        "valid": true,
        "normalized": "+14155552671",
        "quality_score": 0.95
      },
      "overall_quality_score": 0.96,
      "error": null
    },
    {
      "index": 1,
      "email": {
        "valid": false,
        "normalized": null,
        "quality_score": 0.0,
        "disposable": false
      },
      "phone": null,
      "overall_quality_score": 0.0,
      "error": null
    }
  ],
  "total": 2,
  "valid_count": 1,
  "invalid_count": 1,
  "average_quality_score": 0.48
}
```

`index` preserves input order. `error` is set only if the item itself caused a
processing error (malformed JSON, internal timeout) — field-level failures are
surfaced in `email.valid`/`phone.valid`, not in `error`.

---

### `GET /v1/domain/trust/{domain}`

Assess whether a domain is a real, established business domain. Designed for B2B
lead qualification and business email verification. Also useful for
anti-phishing checks when a domain looks like a known brand.

**Internal APIs composed (all fanned out in parallel):**

- WHOIS Lookup (`services/networking/whois`)
- Domain Info (`services/networking/domain`)
- MX Lookup (`services/networking/mx`)

**Request**

```
GET /v1/domain/trust/example.com
```

**Response**

```json
{
  "domain": "example.com",
  "trust_score": 0.91,
  "trust_level": "high",
  "whois": {
    "registrar": "GoDaddy, LLC",
    "created_at": "1995-08-14T04:00:00Z",
    "expires_at": "2026-08-13T04:00:00Z",
    "age_days": 11231,
    "status": ["clientDeleteProhibited", "clientUpdateProhibited"]
  },
  "dns": {
    "has_a_records": true,
    "has_mx_records": true,
    "has_ns_records": true,
    "available": false
  },
  "mx_records": [
    { "hostname": "aspmx.l.google.com", "priority": 1 },
    { "hostname": "alt1.aspmx.l.google.com", "priority": 5 }
  ],
  "flags": []
}
```

**Trust score logic:**

| Condition                                 | Effect                  |
| ----------------------------------------- | ----------------------- |
| Base                                      | 1.0                     |
| Domain not registered (`available: true`) | -1.0 (floor at 0, stop) |
| Domain age < 30 days                      | -0.5 (`new_domain`)     |
| Domain age 30–180 days                    | -0.25 (`young_domain`)  |
| No MX records                             | -0.35 (`no_mx`)         |
| No A records                              | -0.2 (`no_a_records`)   |
| Domain expires in < 14 days               | -0.15 (`expiring_soon`) |

`trust_level` is derived: `>= 0.75` → `high`, `0.4–0.74` → `medium`, `< 0.4` →
`low`

**Flags:** `new_domain`, `young_domain`, `no_mx`, `no_a_records`,
`no_ns_records`, `domain_not_registered`, `expiring_soon`

---

### `POST /v1/content/moderate`

Check user-generated text for toxicity, profanity, and policy violations before
persisting. Returns a structured breakdown across content categories.

**Internal APIs composed:**

- Profanity Filter (`services/validation/profanity`)
- Sentiment Analysis (`services/text/sentiment`)
- Language Detection (`services/text/detectlanguage`) — to gate on language
  support

**Request**

```json
{
  "text": "This is user-generated content to moderate.",
  "language": "en"
}
```

`language` is optional. If omitted, detected automatically using Language
Detection. Provide it when known to save latency.

**Response**

```json
{
  "is_safe": true,
  "toxicity_score": 0.03,
  "sentiment": "neutral",
  "language": "en",
  "language_confidence": 0.99,
  "flags": [],
  "categories": {
    "profanity": false,
    "hate_speech": false,
    "spam": false,
    "violence": false
  }
}
```

`is_safe` is derived: `toxicity_score < 0.5` and `categories` all `false`.

---

### `POST /v1/text/normalize`

Clean and standardize a string: trim whitespace, fix encoding, normalize
Unicode, standardize case. Useful as a pre-processing step before persisting
user input or running comparisons.

**Request**

```json
{
  "text": "  héllo   WORLD  ",
  "operations": ["trim", "lowercase", "normalize_unicode"]
}
```

Valid operations: `trim`, `lowercase`, `uppercase`, `normalize_unicode`,
`collapse_whitespace`, `strip_html`, `remove_punctuation`

**Response**

```json
{
  "original": "  héllo   WORLD  ",
  "normalized": "héllo world",
  "operations_applied": ["trim", "lowercase", "normalize_unicode"],
  "changed": true
}
```

---

## Go package structure

Extends the pattern established in the general systems spec:

```
apps/api/services/systems/
└── data_integrity/
    ├── router.go                     # mounts all sub-routers
    ├── validate/
    │   ├── service.go                # composes email + phone + profanity + sentiment
    │   ├── transport_http.go         # POST /v1/input/validate
    │   └── transport_http_test.go
    ├── validate_batch/
    │   ├── service.go
    │   ├── transport_http.go         # POST /v1/input/validate/batch
    │   └── transport_http_test.go
    ├── domain_trust/
    │   ├── service.go                # composes whois + domain_info + mx
    │   ├── transport_http.go         # GET /v1/domain/trust/{domain}
    │   └── transport_http_test.go
    ├── content_moderate/
    │   ├── service.go                # composes validation/profanity + text/sentiment + text/detectlanguage
    │   ├── transport_http.go         # POST /v1/content/moderate
    │   └── transport_http_test.go
    └── text_normalize/
        ├── service.go                # wraps existing text/normalize service
        ├── transport_http.go         # POST /v1/text/normalize
        └── transport_http_test.go
```

Each `service.go` receives its dependencies as interfaces (same pattern as BIN,
email, and other services — stubs in tests, real services in production). The
router mounts under the `systems` prefix defined in `routes_v1.go`.

---

## Testing requirements

### `/v1/input/validate`

| Case                        | Expected                                                                        |
| --------------------------- | ------------------------------------------------------------------------------- |
| Valid email + phone         | 200, both `valid: true`, `overall_quality_score` > 0.9                          |
| Disposable email            | 200, `email.disposable: true`, `email_disposable` in flags, quality_score ≤ 0.5 |
| Invalid email syntax        | 200, `email.valid: false`, quality_score 0.0                                    |
| Valid email only (no phone) | 200, `phone: null`, `overall_quality_score` based on email only                 |
| VoIP phone                  | 200, `phone.risk.is_voip: true`, `phone_voip` in flags                          |
| No fields present           | 422                                                                             |
| Email with typo suggestion  | 200, `email.suggestion` non-null, `email_has_suggestion` in flags               |

### `/v1/input/validate/batch`

| Case                           | Expected                                             |
| ------------------------------ | ---------------------------------------------------- |
| 2 valid items                  | 200, `valid_count: 2`, `average_quality_score` > 0.9 |
| Mix of valid and invalid       | 200, per-item `valid` differs, `invalid_count` > 0   |
| Over 50 items                  | 422                                                  |
| Empty array                    | 422                                                  |
| One item causes internal error | 200, that item's `error` non-null, others unaffected |

### `/v1/domain/trust/{domain}`

| Case                                  | Expected                                                  |
| ------------------------------------- | --------------------------------------------------------- |
| Established domain (e.g. example.com) | 200, `trust_level: "high"`, no flags                      |
| New domain (< 30 days old)            | 200, `new_domain` in flags, `trust_score` ≤ 0.5           |
| No MX records                         | 200, `no_mx` in flags, trust_score reduced                |
| Domain not registered                 | 200, `domain_not_registered` in flags, `trust_score: 0.0` |
| Invalid domain format                 | 422                                                       |

### `/v1/content/moderate`

| Case                   | Expected                                                                       |
| ---------------------- | ------------------------------------------------------------------------------ |
| Clean text             | 200, `is_safe: true`, `toxicity_score` < 0.1                                   |
| Text with profanity    | 200, `is_safe: false`, `categories.profanity: true`, `text_profanity` in flags |
| Empty text             | 422                                                                            |
| Language auto-detected | 200, `language` field populated from detection                                 |

### `/v1/text/normalize`

| Case                              | Expected                                   |
| --------------------------------- | ------------------------------------------ |
| Text with leading/trailing spaces | 200, `normalized` trimmed, `changed: true` |
| Already clean text                | 200, `changed: false`                      |
| Invalid operation in array        | 422                                        |
| Empty text                        | 422                                        |

---

## Ticket breakdown

One ticket per sub-service. Suggested implementation order (simple to complex):

| # | Ticket                    | Endpoint                        | Composes                                |
| - | ------------------------- | ------------------------------- | --------------------------------------- |
| 1 | Text Normalize endpoint   | `POST /v1/text/normalize`       | No external services                    |
| 2 | Content Moderate endpoint | `POST /v1/content/moderate`     | Profanity + Sentiment + Language        |
| 3 | Domain Trust endpoint     | `GET /v1/domain/trust/{domain}` | WHOIS + Domain Info + MX                |
| 4 | Input Validate single     | `POST /v1/input/validate`       | Email + Phone + (Profanity + Sentiment) |
| 5 | Input Validate batch      | `POST /v1/input/validate/batch` | Same as #4, with concurrency pool       |

Each ticket should reference the relevant section of this document for request
shape, response shape, score logic, flag constants, and test cases.

## Open questions for engineer

1. **Language detection fallback** — if `services/text/detectlanguage` fails in
   `/v1/content/moderate`, return `language: null` and
   `language_confidence: null` rather than erroring, so callers still get the
   moderation result.
2. **WHOIS reliability** — WHOIS lookups can time out or return partial data for
   some TLDs. The domain trust endpoint should degrade gracefully: if WHOIS
   fails, omit `whois` from response and set a `whois_unavailable` flag rather
   than returning 503.
3. **Batch concurrency** — `/v1/input/validate/batch` will fan out email + phone
   per item. With 50 items and 2 calls each, that's up to 100 in-flight calls.
   Use a semaphore (8-10 workers, matching the email batch pattern) not
   unbounded goroutines.
