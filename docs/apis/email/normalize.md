# Email Normalizer API

Normalize email addresses to a canonical form: trimming, lowercasing,
provider-specific rules (Gmail dots and plus tags, `googlemail.com` →
`gmail.com`), and a structured list of changes applied.

## Endpoints

| Method | Path                            | Description                    |
| ------ | ------------------------------- | ------------------------------ |
| POST   | `/v1/text/normalize`            | Normalize a single address     |
| POST   | `/v1/text/normalize/batch`      | Normalize up to 100 addresses |

Usage is billed **per email** on the batch route (see `X-Usage-Count` / gateway
docs).

---

## POST /v1/text/normalize

### Request body

```json
{
  "email": "Te.st.User+spam@Googlemail.com"
}
```

| Field   | Type   | Required | Description           |
| ------- | ------ | -------- | --------------------- |
| `email` | string | Yes      | Address to normalize |

### Responses

- **200** — Success; `data` matches the single-item shape (`original`,
  `normalized`, `local`, `domain`, `changes`).
- **400** — Invalid JSON, unknown fields, or address cannot be normalized.
- **422** — Validation failure (e.g. missing `email`).

---

## POST /v1/text/normalize/batch

Normalizes up to **100** addresses in one request. Results are in the **same
order** as `emails`. Each row uses the **phone batch model**: HTTP **200** with
`valid: true` or `false` per item; invalid addresses do not fail the whole
request.

### Request body

```json
{
  "emails": ["user@example.com", "not-an-email", "te.st@gmail.com"]
}
```

| Field    | Type     | Required | Description                          |
| -------- | -------- | -------- | ------------------------------------ |
| `emails` | string[] | Yes      | Min 1, max 100; each entry non-empty |

### Response (example)

```json
{
  "data": {
    "results": [
      {
        "original": "user@example.com",
        "normalized": "user@example.com",
        "local": "user",
        "domain": "example.com",
        "changes": [],
        "valid": true
      },
      {
        "original": "not-an-email",
        "valid": false,
        "message": "…"
      },
      {
        "original": "te.st@gmail.com",
        "normalized": "test@gmail.com",
        "local": "test",
        "domain": "gmail.com",
        "changes": ["lowercased", "removed_dots"],
        "valid": true
      }
    ],
    "total": 3
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

When `valid` is `false`, only `original`, `valid`, and `message` are populated
for that row.

### Errors

- **422** — Body validation (empty `emails`, too many items, empty string in
  the array).
- **400** — Malformed JSON or unknown fields.

Dashboard catalog and examples: `email-normalize` API doc in the dashboard
config (`email-normalize.yml`).
