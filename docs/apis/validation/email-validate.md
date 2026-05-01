# Email Validate API

Full email validation in a single call using the standard library `net/mail`
parser for syntax, live MX DNS lookups, the
[is-email-disposable](https://github.com/bobadilla-tech/is-email-disposable)
blocklist, the
[go-email-normalizer](https://github.com/bobadilla-tech/go-email-normalizer)
package for normalization, and a Levenshtein-based typo detector for common
provider domains.

## Endpoint

`POST /v1/validation/email`

## What It Checks

| Check           | How                                                       |
| --------------- | --------------------------------------------------------- |
| Syntax          | RFC 5322 parse via `net/mail.ParseAddress`                |
| MX record       | Live DNS lookup — domain must have ≥ 1 MX record          |
| Disposable      | Checked against the 90,000+ domain blocklist              |
| Normalization   | Lowercase, plus-tag removal, alias-domain resolution      |
| Typo suggestion | Levenshtein distance ≤ 2 against ~20 well-known providers |

## Request

### Single Email
```json
{
  "email": "user@gmial.com"
}
```

**Validation:** `email` is required. A missing field returns
`422 Unprocessable
Entity`. Syntax errors do not return a 4xx — they return
`200 OK` with `syntax_valid: false` and `valid: false`.


### Batch (multiple emails)

Accepts up to **50 emails** per request. Results are returned in the same
order as the input array. All emails are processed, and validation results
are returned for each item regardless of validity.

Each email in the request counts as **1 credit** (`X-Usage-Count` is set to
the number of emails in the request).

`POST /v1/validation/email/batch`

```json
  { "emails": [ "user@gmail.com", "user@gmial.com", "bad-email" ] }
```

**Validation:** `emails` is required. A missing or empty array returns
`422 Unprocessable Entity`. Requests with more than 50 items also return
`422 Unprocessable Entity`. Each item must be a string. Syntax errors are
handled per item and do not return a 4xx — the API returns `200 OK` with
`syntax_valid: false` and `valid: false` for each invalid email.


## Response Envelope

All responses are wrapped in the standard envelope:

```json
{
  "data": { ... },
  "metadata": { "timestamp": "2026-01-01T00:00:00Z" }
}
```

## Response Fields

| Field          | Type         | Description                                                                  |
| -------------- | ------------ | ---------------------------------------------------------------------------- |
| `email`        | string\|null | Exact input supplied by the caller; `null` when syntax is invalid            |
| `valid`        | boolean      | `true` only when `syntax_valid` and `mx_valid` are both `true`               |
| `syntax_valid` | boolean      | Passes RFC 5322 syntax check                                                 |
| `mx_valid`     | boolean      | Domain has at least one MX record                                            |
| `disposable`   | boolean      | Domain is on the disposable email blocklist                                  |
| `normalized`   | string\|null | Canonical address after normalization; `null` when syntax is invalid         |
| `domain`       | string\|null | Domain part (after `@`); `null` when syntax is invalid                       |
| `suggestion`   | string\|null | Closest well-known domain when the input looks like a typo; `null` otherwise |

## Examples

### Valid address

```bash
curl -X POST https://api.requiems.xyz/v1/validation/email \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"email": "user@gmail.com"}'
```

Response:

```json
{
  "data": {
    "email": "user@gmail.com",
    "valid": true,
    "syntax_valid": true,
    "mx_valid": true,
    "disposable": false,
    "normalized": "user@gmail.com",
    "domain": "gmail.com",
    "suggestion": null
  },
  "metadata": { "timestamp": "2026-01-01T00:00:00Z" }
}
```

Batch Response:

```json
{
  "data": {
    "results": [
      {
        "email": "user@gmail.com",
        "valid": true,
        "syntax_valid": true,
        "mx_valid": true,
        "disposable": false,
        "normalized": "user@gmail.com",
        "domain": "gmail.com",
        "suggestion": null
      },
      {
        "email": "user@gmial.com",
        "valid": false,
        "syntax_valid": true,
        "mx_valid": false,
        "disposable": false,
        "normalized": "user@gmial.com",
        "domain": "gmial.com",
        "suggestion": "gmail.com"
      },
      {
        "email": null,
        "valid": false,
        "syntax_valid": false,
        "mx_valid": false,
        "disposable": false,
        "normalized": null,
        "domain": null,
        "suggestion": null
      }
    ],
    "total": 3
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

### Typo in domain

```bash
curl -X POST https://api.requiems.xyz/v1/validation/email \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"email": "user@gmial.com"}'
```

Response:

```json
{
  "data": {
    "email": "user@gmial.com",
    "valid": false,
    "syntax_valid": true,
    "mx_valid": false,
    "disposable": false,
    "normalized": "user@gmial.com",
    "domain": "gmial.com",
    "suggestion": "gmail.com"
  },
  "metadata": { "timestamp": "2026-01-01T00:00:00Z" }
}
```

### Invalid syntax

```json
{
  "data": {
    "email": null,
    "valid": false,
    "syntax_valid": false,
    "mx_valid": false,
    "disposable": false,
    "normalized": null,
    "domain": null,
    "suggestion": null
  },
  "metadata": { "timestamp": "2026-01-01T00:00:00Z" }
}
```

## Error Codes

| Code                | Status | When                                                      |
| ------------------- | ------ | --------------------------------------------------------- |
| `validation_failed` | 422    | `email` field is missing from the request body            |
| `bad_request`       | 400    | Body is missing, invalid JSON, or contains unknown fields |
| `internal_error`    | 500    | Unexpected server error                                   |

## Typo Detection

The `suggestion` field compares the supplied domain against a curated list of
well-known providers using Levenshtein edit distance. A suggestion is returned
only when the distance is ≤ 2. Covered providers include:

`gmail.com`, `googlemail.com`, `yahoo.com`, `yahoo.co.uk`, `outlook.com`,
`hotmail.com`, `icloud.com`, `me.com`, `mac.com`, `aol.com`, `protonmail.com`,
`proton.me`, `live.com`, `msn.com`, `yandex.com`, `yandex.ru`, `mail.com`,
`zoho.com`

## Use Cases

### Inline typo correction at signup

Show a prompt when `suggestion` is non-null before the user submits the form.

```javascript
const { data } = await validateEmail(userInput);
if (data.suggestion) {
  showPrompt(
    `Did you mean ${userInput.replace(data.domain, data.suggestion)}?`,
  );
}
```

### Block undeliverable addresses

Reject the address early when `valid` is `false` to avoid wasted sends.

```javascript
const { data } = await validateEmail(userInput);
if (!data.valid) {
  throw new ValidationError("Please enter a valid email address.");
}
```

### Flag disposable signups

```javascript
const { data } = await validateEmail(userInput);
if (data.disposable) {
  flagForReview(userInput);
}
```
