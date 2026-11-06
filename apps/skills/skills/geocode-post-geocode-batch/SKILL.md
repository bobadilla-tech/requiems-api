---
name: geocode-post-geocode-batch
api: "Geocoding"
method: POST
path: /v1/places/geocode/batch
base_url: https://requiems.xyz
description: "Convert up to 20 addresses to coordinates in a single request. Processed concurrently; results are returned in input order. Per-item errors are reported inline. Cache hits avoid outbound network calls."
---

## Endpoint

**POST https://requiems.xyz/v1/places/geocode/batch**

## Batch Geocode Addresses

Convert up to 20 addresses to coordinates in a single request. Processed concurrently; results are returned in input order. Per-item errors are reported inline. Cache hits avoid outbound network calls.

## Parameters

| Name | Type | Required | Location | Description |
| ---- | ---- | -------- | -------- | ----------- |
| `addresses` | array<string> | yes | body | List of addresses to geocode (1–20 items, each non-empty). Max 20 due to Nominatim usage policy. |

## Request Example

```json
{
  "addresses": [
    "1600 Pennsylvania Ave NW, Washington DC",
    "Eiffel Tower, Paris"
  ]
}
```

## Response Example

```json
{
  "data": {
    "results": [
      {
        "address": "1600 Pennsylvania Ave NW, Washington DC",
        "result": { "address": "White House, 1600, Pennsylvania Avenue Northwest...", "city": "Washington", "country": "US", "lat": 38.8977, "lon": -77.0365 }
      },
      {
        "address": "Eiffel Tower, Paris",
        "result": { "address": "Tour Eiffel, 5, Avenue Anatole France, Paris...", "city": "Paris", "country": "FR", "lat": 48.8584, "lon": 2.2945 }
      }
    ],
    "total": 2
  },
  "metadata": { "timestamp": "2026-01-01T00:00:00Z" }
}
```

## Response Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| `results` | array<object> | Per-item results in input order. |
| `results[].address` | string | The original address as provided in the request. |
| `results[].result` | object | Geocoding result (omitted on error). Same fields as the single-item endpoint. |
| `results[].error` | string | Error message if geocoding failed (omitted on success). |
| `total` | integer | Total number of items processed. |

## Errors

- `422` **validation_failed** — The addresses array is missing, empty, or exceeds 20 items.
- `400` **bad_request** — The request body is missing or malformed.
