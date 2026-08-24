# Cities API

Look up metadata for cities worldwide including population, IANA timezone,
country, and geographic coordinates. Covers ~22,000 cities with a population
above 15,000 from the [GeoNames dataset](https://www.geonames.org/).

## Endpoint

`GET /v1/places/cities/{city}`

| Parameter | Location | Required | Description                  |
| --------- | -------- | -------- | ---------------------------- |
| `city`    | path     | Yes      | City name (case-insensitive) |

```bash
curl "https://requiems.xyz/v1/places/cities/london" \
  -H "requiems-api-key: YOUR_API_KEY"
```

## Response

```json
{
  "data": {
    "name": "London",
    "country": "GB",
    "population": 7556900,
    "timezone": "Europe/London",
    "lat": 51.5085,
    "lon": -0.1257
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Name Resolution

- Lookup is case-insensitive: `london`, `London`, and `LONDON` all work.
- When multiple cities share the same name the most populous one is returned
  (e.g. `london` → London, UK not London, Ontario).
- Multi-word names work: `/v1/places/cities/new%20york%20city`.

## Timezone

The `timezone` field is an IANA timezone identifier compatible with the
`/v1/places/timezone` endpoint. Use it directly in `time.LoadLocation` (Go),
`ZoneInfo` (Python), or `Intl.DateTimeFormat` (JavaScript).

## Error Codes

| Code             | Status | When                                  |
| ---------------- | ------ | ------------------------------------- |
| `not_found`      | 404    | No city with that name in the dataset |
| `internal_error` | 500    | Unexpected failure                    |

## Coverage

The GeoNames `cities15000` dataset includes all cities with a population above
15,000 as of the last dataset update. This covers essentially all significant
towns and cities worldwide. Very small settlements, villages, and newly
incorporated cities may not be present.
