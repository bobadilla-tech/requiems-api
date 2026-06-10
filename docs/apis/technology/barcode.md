# Barcode API

## Status

⏳ **Planned** - Not yet implemented

## Overview

Generate and validate barcodes. This endpoint will support various barcode
formats and standards.

## Planned Endpoints

### Generate Barcode

**Planned Endpoint:** `GET /v1/internet-technology/barcode`

Generate or validate a barcode.

### Generate Barcodes (Batch)

`POST /v1/technology/barcode/batch`

Generate up to 20 barcodes in a single request.

Each barcode is processed independently. Invalid items do not fail the entire
request. Results are returned in the same order as the input.

## Request Body

```json
{
  "items": [
    {
      "data": "123456789",
      "type": "code128"
    },
    {
      "data": "1234567",
      "type": "ean8"
    }
  ]
}
```

## Response Example

```json
{
  "data": {
    "results": [
      {
        "image": "iVBORw0KGgoAAAANSUhEUgAA...",
        "type": "code128",
        "width": 300,
        "height": 100,
        "success": true,
        "error": ""
      },
      {
        "image": "iVBORw0KGgoAAAANSUhEUgAA...",
        "type": "ean8",
        "width": 300,
        "height": 100,
        "success": true,
        "error": ""
      }
    ],
    "total": 2
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Response Fields

| Field | Type | Description |
| ------- | ------ | ----------- |
| `results` | array | Batch results preserving input order |
| `results[].image` | string | Base64-encoded PNG image |
| `results[].type` | string | Barcode format used |
| `results[].width` | integer | Image width in pixels |
| `results[].height` | integer | Image height in pixels |
| `results[].success` | boolean | Whether barcode generation succeeded |
| `results[].error` | string | Error message when generation fails |
| `total` | integer | Number of items processed |

## Response Headers

| Header | Description |
| ---------- | ----------- |
| `X-Usage-Count` | Number of barcodes processed in the batch |

## Examples

### Batch Request

```bash
curl -X POST "https://api.requiems.xyz/v1/technology/barcode/batch" \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {
        "data": "123456789",
        "type": "code128"
      },
      {
        "data": "HELLO",
        "type": "code93"
      }
    ]
  }'
```

### Python

```python
import requests

url = "https://api.requiems.xyz/v1/technology/barcode/batch"

response = requests.post(
    url,
    headers={
        "requiems-api-key": "YOUR_API_KEY",
        "Content-Type": "application/json"
    },
    json={
        "items": [
            {
                "data": "123456789",
                "type": "code128"
            },
            {
                "data": "HELLO",
                "type": "code93"
            }
        ]
    }
)

print(response.json())
```

## Batch Validation Rules

- Minimum items: 1
- Maximum items: 20
- Supported types: `code128`, `code93`, `code39`, `ean8`, `ean13`
- Results preserve request order
- Individual failures are reported per item
- Successful items are returned even if other items fail
