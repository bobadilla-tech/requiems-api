# Payments Intelligence System — Engine Spec

The Payments Intelligence System validates financial instruments and assesses
transaction fraud risk. It answers two distinct questions: is this payment data
structurally valid and internally consistent? And is this transaction risky at
runtime?

**Distinct from other systems:**

| Question                                             | System                    |
| ---------------------------------------------------- | ------------------------- |
| Is this email/phone data clean?                      | Data Integrity            |
| Is this user fraudulent?                             | Identity & Risk           |
| Is this payment instrument valid / transaction safe? | **Payments Intelligence** |
| Where is this user located?                          | Global Data               |

IP geolocation and VPN detection appear in both Identity & Risk and Payments
Intelligence. The distinction: Identity & Risk uses IP to assess who the USER is
at signup. Payments Intelligence uses IP to assess whether a TRANSACTION at
checkout is geographically consistent with the payment instrument. Different
call sites, different questions, different outputs.

---

## Conventions

- Base path: `/v1/`
- Every response wrapped in
  `{"data": {...}, "metadata": {"timestamp": "...", "trace_id": "..."}}`
- `risk_score`: `0.0` (no risk) to `1.0` (certain risk)
- `is_safe` always derived: `risk_score < 0.5`
- `flags` always `string[]` of machine-readable codes, never freeform
- Score 0–1, never 0–100

---

## Endpoints

### Engine: `POST /v1/payment/validate`

Primary composed endpoint. Validate one or more payment instruments in a single
call and get a cross-instrument consistency check. Replaces up to three separate
API calls (BIN, IBAN, SWIFT) and adds the country consistency layer that none of
those endpoints can provide alone.

At least one of `bin`, `iban`, `swift` must be present.

**Internal APIs composed (whichever fields are provided, all parallel):**

- BIN Lookup (`services/finance/bin`)
- IBAN Validator (`services/finance/iban`)
- SWIFT Lookup (`services/finance/swift`)

**Request**

```json
{
  "bin": "424242",
  "iban": "GB29NWBK60161331926819",
  "swift": "NWBKGB2L"
}
```

**Response**

```json
{
  "bin": {
    "valid": true,
    "scheme": "visa",
    "card_type": "credit",
    "card_level": "classic",
    "country_code": "US",
    "issuer": "Chase",
    "prepaid": false,
    "luhn": true
  },
  "iban": {
    "valid": true,
    "country_code": "GB",
    "bank_code": "NWBK",
    "account_number": "31926819"
  },
  "swift": {
    "valid": true,
    "institution": "NatWest",
    "country": "GB",
    "branch": "London"
  },
  "consistency": {
    "ok": false,
    "flags": ["country_mismatch_bin_iban"]
  }
}
```

Fields omitted from the request are `null` in the response. `consistency.ok` is
`true` only when all provided instruments agree on country and bank.

**Consistency flags:**

| Flag                          | Condition                                                        |
| ----------------------------- | ---------------------------------------------------------------- |
| `country_mismatch_bin_iban`   | BIN `country_code` ≠ IBAN `country_code`                         |
| `country_mismatch_bin_swift`  | BIN `country_code` ≠ SWIFT `country`                             |
| `country_mismatch_iban_swift` | IBAN `country_code` ≠ SWIFT `country`                            |
| `bank_mismatch_iban_swift`    | IBAN `bank_code` does not match SWIFT BIC prefix (first 4 chars) |

Consistency check only runs between instruments actually provided. If only `bin`
is present, `consistency.ok: true` and `consistency.flags: []` (single
instrument has no cross-checks to fail).

---

### `POST /v1/transaction/risk`

Score a payment transaction for fraud risk at checkout time. Combines card
intelligence with IP signals to detect geographic inconsistencies and known
fraud patterns that emerge only when the instruments are seen together in
transaction context.

**Distinct from `/v1/payment/validate`:** validate checks instrument FORMAT and
CONSISTENCY (are these instruments structurally valid and internally
compatible?). transaction/risk checks FRAUD SIGNALS at payment execution time
(is this checkout attempt suspicious given the IP, card origin, and billing
country?).

**Internal APIs composed (all fanned out in parallel):**

- BIN Lookup (`services/finance/bin`)
- VPN/Proxy Detection (`services/networking/ip/vpn`)
- IP Geolocation (`services/networking/ip/info`)

**Request**

```json
{
  "card_bin": "424242",
  "ip_address": "92.168.1.1",
  "billing_country": "US",
  "amount_usd": 299.0
}
```

Required: `card_bin` + `ip_address`. Optional: `billing_country`, `amount_usd`.

**Response**

```json
{
  "risk_score": 0.71,
  "is_safe": false,
  "flags": ["country_mismatch", "vpn_detected"],
  "signals": {
    "ip_country": "RU",
    "billing_country": "US",
    "bin_country": "US",
    "vpn_detected": true,
    "is_proxy": false,
    "is_tor": false,
    "country_mismatch": true
  }
}
```

**Risk score weights (additive, cap 1.0):**

| Signal                                        | Weight               |
| --------------------------------------------- | -------------------- |
| IP country ≠ billing_country                  | +0.35                |
| BIN country ≠ billing_country                 | +0.25                |
| VPN detected                                  | +0.20                |
| Proxy detected                                | +0.30                |
| TOR detected                                  | +0.40                |
| IP `fraud_score` contribution                 | `fraud_score × 0.25` |
| High-value + VPN (`amount_usd > 500` and VPN) | extra +0.15          |

`country_mismatch` flag is set when IP country ≠ billing_country OR BIN country
≠ billing_country.

**Flags:** `country_mismatch`, `vpn_detected`, `proxy_detected`, `tor_detected`,
`high_value_vpn`, `bin_country_mismatch`, `ip_country_mismatch`

---

## Go package structure

```text
apps/api/services/systems/
└── payments/
    ├── router.go                         # mounts both sub-routers
    ├── payment_validate/
    │   ├── service.go                    # composes bin + iban + swift; consistency check
    │   ├── transport_http.go             # POST /v1/payment/validate
    │   └── transport_http_test.go
    └── transaction_risk/
        ├── service.go                    # composes bin + vpn + ip/info; risk scoring
        ├── transport_http.go             # POST /v1/transaction/risk
        └── transport_http_test.go
```

---

## Testing requirements

### `/v1/payment/validate`

| Case                               | Expected                                                                    |
| ---------------------------------- | --------------------------------------------------------------------------- |
| Valid BIN only                     | 200, `bin.valid: true`, `iban: null`, `swift: null`, `consistency.ok: true` |
| BIN + IBAN, same country           | 200, `consistency.ok: true`, no flags                                       |
| BIN (US) + IBAN (GB)               | 200, `country_mismatch_bin_iban` in consistency.flags                       |
| IBAN + SWIFT, bank code mismatch   | 200, `bank_mismatch_iban_swift` in flags                                    |
| All three provided, all consistent | 200, `consistency.ok: true`, flags empty                                    |
| No instruments provided            | 422                                                                         |
| Invalid BIN format                 | 200, `bin.valid: false`, `bin.luhn: false`                                  |
| All composed services stubbed      | ✓                                                                           |

### `/v1/transaction/risk`

| Case                               | Expected                                                 |
| ---------------------------------- | -------------------------------------------------------- |
| IP country matches billing, no VPN | 200, `is_safe: true`, `risk_score < 0.3`                 |
| IP country ≠ billing country       | 200, `country_mismatch` flag, score ≥ 0.35               |
| VPN detected                       | 200, `vpn_detected` flag, score ≥ 0.20                   |
| TOR detected                       | 200, `tor_detected` flag, score ≥ 0.40, `is_safe: false` |
| High value (> $500) + VPN          | 200, `high_value_vpn` flag, extra +0.15 on score         |
| BIN country ≠ billing_country      | 200, `bin_country_mismatch` flag                         |
| `card_bin` missing                 | 422                                                      |
| `ip_address` missing               | 422                                                      |
| All composed services stubbed      | ✓                                                        |

---

## Open questions for engineer

1. **BIN country for consistency** — the BIN service returns `country_code` (ISO
   2-char). SWIFT returns `country` (may be full name or ISO depending on data
   source). Verify both fields are ISO 2-char before comparing, normalize if
   not.
2. **`billing_country` absent in transaction/risk** — when `billing_country` is
   not provided, the `country_mismatch` check can only compare BIN country vs IP
   country. Document this in the response: set `billing_country: null` in
   signals and skip the IP-vs-billing comparison.
3. **`fraud_score` field name** — VPN service returns `fraud_score` as integer
   0–100. Normalize to 0.0–1.0 before applying the weight table.

## Ticket breakdown

| # | Endpoint                    | Composes            | Priority |
| - | --------------------------- | ------------------- | -------- |
| 1 | `POST /v1/payment/validate` | bin + iban + swift  | high     |
| 2 | `POST /v1/transaction/risk` | bin + vpn + ip/info | high     |

Both are high priority — they unlock the Payments Intelligence system landing
page. Implement validate first (no IP calls, simpler), then transaction/risk.
