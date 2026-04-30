# Inflation API

Annual CPI inflation rates for 241 countries, sourced from the World Bank.

## Endpoints

### Single country

`GET /v1/finance/inflation?country=US`

| Name      | Type   | Required | Description                                                      |
| --------- | ------ | -------- | ---------------------------------------------------------------- |
| `country` | string | yes      | ISO 3166-1 alpha-2 country code (e.g. US, GB). Case-insensitive. |

```json
{
  "data": {
    "country": "US",
    "rate": 2.9495,
    "period": "2024",
    "historical": [
      { "period": "2023", "rate": 4.1163 },
      { "period": "2022", "rate": 8.0028 }
    ]
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

### Batch (multiple countries)

`POST /v1/finance/inflation/batch`

Accepts up to **50 countries** per request. Results are returned in the same
order as the input array. Countries with no data are included with `found: false`
instead of failing the entire request.

Each country in the request counts as **1 credit** (`X-Usage-Count` is set to
the number of countries in the request).

**Request body:**

```json
{ "countries": ["US", "AR", "DE"] }
```

**Response:**

```json
{
  "data": {
    "results": [
      {
        "country": "US",
        "found": true,
        "rate": 2.9495,
        "period": "2024",
        "historical": [
          { "period": "2023", "rate": 4.1163 }
        ]
      },
      {
        "country": "AR",
        "found": true,
        "rate": 211.4,
        "period": "2024",
        "historical": []
      },
      {
        "country": "XX",
        "found": false
      }
    ],
    "total": 3
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

- `found: false` — the country code is valid but has no data in the World Bank set.
- `total` — total number of items returned (always equals the number of countries sent).

## Data Source

World Bank indicator `FP.CPI.TOTL.ZG`. Updated annually; re-run
`cmd/seed-inflation` to pull new figures when the World Bank publishes them.

## Error Codes

| Code                | Status | When                                                       |
| ------------------- | ------ | ---------------------------------------------------------- |
| `bad_request`       | 400    | Missing or invalid country code (single endpoint)          |
| `not_found`         | 404    | No data for that country (single endpoint)                 |
| `validation_failed` | 422    | Invalid batch body — empty array, over 50 items, bad codes |
| `internal_error`    | 500    | Unexpected failure                                         |
