# BIN Lookup API

Returns card metadata for a Bank Identification Number (BIN/IIN) — the first 6–8
digits of a payment card number.

## Endpoint

`GET /v1/finance/bin/{bin}`

## Path Parameters

| Parameter | Type   | Required | Description                                                         |
| --------- | ------ | -------- | ------------------------------------------------------------------- |
| `bin`     | string | Yes      | 6–8 digit BIN prefix. Dashes and spaces are stripped automatically. |

## Response

```json
{
  "data": {
    "bin": "424242",
    "scheme": "visa",
    "card_type": "credit",
    "card_level": "classic",
    "issuer_name": "Chase",
    "issuer_url": "www.chase.com",
    "issuer_phone": "+18002324000",
    "country_code": "US",
    "country_name": "United States",
    "prepaid": false,
    "luhn_prefix_valid": true,
    "confidence": 0.92
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Response Fields

| Field          | Type    | Description                                                                                                                     |
| -------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `bin`          | string  | The normalised BIN prefix used for the lookup                                                                                   |
| `scheme`       | string  | Card network: `visa`, `mastercard`, `amex`, `discover`, `jcb`, `diners`, `unionpay`, `maestro`, `mir`, `rupay`, `private_label` |
| `card_type`    | string  | `credit`, `debit`, `prepaid`, or `charge`                                                                                       |
| `card_level`   | string  | `classic`, `gold`, `platinum`, `infinite`, `business`, `signature`, `standard`                                                  |
| `issuer_name`  | string  | Name of the card-issuing bank                                                                                                   |
| `issuer_url`   | string  | Bank website URL                                                                                                                |
| `issuer_phone` | string  | Bank customer service phone number                                                                                              |
| `country_code` | string  | ISO 3166-1 alpha-2 country code of the issuing bank                                                                             |
| `country_name` | string  | Full country name                                                                                                               |
| `prepaid`      | boolean | Whether the card is a prepaid card                                                                                              |
| `luhn_prefix_valid` | boolean | Whether the BIN prefix (not a full card number) passes the Luhn algorithm check                                            |
| `confidence`   | number  | Data quality score (0.00–1.00). Multi-source confirmed records score higher                                                     |

## BIN vs 8-digit IIN

The ISO 7812-1:2017 standard expanded BINs from 6 digits to 8 digits. This
endpoint accepts both:

- 6-digit lookup: matches the original 6-digit database entries
- 8-digit lookup: tries an exact 8-digit match first, then falls back to the
  6-digit prefix if no 8-digit record exists

## Scheme Detection

The `scheme` field is derived from the BIN database. If a record has no scheme,
the API falls back to ISO/IEC 7812 prefix-range detection:

| Scheme      | Prefix Ranges                    |
| ----------- | -------------------------------- |
| Visa        | 4                                |
| Mastercard  | 51–55, 2221–2720                 |
| Amex        | 34, 37                           |
| Discover    | 6011, 622126–622925, 644–649, 65 |
| JCB         | 3528–3589                        |
| Diners Club | 300–305, 36, 38                  |
| UnionPay    | 62, 81                           |
| Maestro     | 6304, 6759, 6761–6763            |
| Mir         | 2200–2204                        |
| RuPay       | 60, 6521, 6522                   |

## Luhn Check

The `luhn_prefix_valid` field runs the Luhn algorithm on the BIN prefix itself
(not a full card number). For a 6–8 digit prefix this is a partial check — it
reflects whether the prefix has a valid Luhn structure, not whether a specific
full card number is valid.

## Error Codes

| Code             | Status | When                                                   |
| ---------------- | ------ | ------------------------------------------------------ |
| `bad_request`    | 400    | BIN is not 6–8 digits or contains non-digit characters |
| `not_found`      | 404    | BIN prefix not found in the database                   |
| `internal_error` | 500    | Unexpected server error                                |

## Data Sources

BIN data is aggregated from multiple open-source datasets and normalised with
confidence scoring. Records confirmed across multiple sources receive a higher
`confidence` score.
