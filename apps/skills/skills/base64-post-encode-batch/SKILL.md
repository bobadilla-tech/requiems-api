---
name: base64-post-encode-batch
api: "Base64 Encode / Decode"
method: POST
path: /v1/technology/base64/encode/batch
base_url: https://requiems.xyz
description: "Encode multiple strings to Base64 in a single request"
---

## Endpoint

**POST https://requiems.xyz/v1/technology/base64/encode/batch**

## Encode Batch

Encode multiple strings to Base64 in a single request

## Parameters

| Name | Type | Required | Location | Description |
| ---- | ---- | -------- | -------- | ----------- |
| `values` | array<string> | yes | body | List of values to encode |
| `variant` | string | no | body | Encoding variant (standard or url) |

## Response Example

```json
{
  "data": {
    "results": [
      {
        "result": "SGVsbG8="
      },
      {
        "result": "V29ybGQ="
      }
    ],
    "total": 2
  }
}
```

## Response Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| `results` | array<object> | Encoded results |
| `total` | integer | Number of processed items |

## Errors

- `undefined` **422** — Invalid variant or values list size outside 1-50
