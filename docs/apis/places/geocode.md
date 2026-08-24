# Geocoding API

Convert human-readable addresses into geographic coordinates (geocoding) and
convert coordinates back into addresses (reverse geocoding). Powered by
[Nominatim](https://nominatim.openstreetmap.org/) (OpenStreetMap).

## Endpoints

### Forward Geocoding

`GET /v1/places/geocode`

| Parameter | Location | Required | Description                  |
| --------- | -------- | -------- | ---------------------------- |
| `address` | query    | Yes      | Free-text address to geocode |

```bash
curl "https://requiems.xyz/v1/places/geocode?address=1600+Pennsylvania+Ave+NW" \
  -H "requiems-api-key: YOUR_API_KEY"
```

Response:

```json
{
  "data": {
    "address": "White House, 1600, Pennsylvania Avenue Northwest, Washington, DC, 20500, United States",
    "city": "Washington",
    "country": "US",
    "lat": 38.8976763,
    "lon": -77.0365298
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

### Reverse Geocoding

`GET /v1/places/reverse-geocode`

| Parameter | Location | Required | Description             |
| --------- | -------- | -------- | ----------------------- |
| `lat`     | query    | Yes      | Latitude (-90 to 90)    |
| `lon`     | query    | Yes      | Longitude (-180 to 180) |

```bash
curl "https://requiems.xyz/v1/places/reverse-geocode?lat=38.8977&lon=-77.0365" \
  -H "requiems-api-key: YOUR_API_KEY"
```

Response:

```json
{
  "data": {
    "lat": 38.8977,
    "lon": -77.0365,
    "address": "White House, 1600, Pennsylvania Avenue Northwest, Washington, DC, 20500, United States",
    "city": "Washington",
    "country": "US"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Caching

Results are cached in Redis for **24 hours**. Repeated calls with the same
address or coordinates return instantly without hitting Nominatim.

## City Resolution

The `city` field is resolved from the Nominatim `address` block in priority
order: `city` → `town` → `village` → `county`. This handles rural areas where no
formal city is present.

## Error Codes

| Code             | Status | When                                            |
| ---------------- | ------ | ----------------------------------------------- |
| `not_found`      | 404    | No results found for the address or coordinates |
| `upstream_error` | 503    | Nominatim is temporarily unavailable            |
| `bad_request`    | 400    | Required parameter is missing or invalid        |

## Notes

- Address geocoding works best when the query includes country context (e.g.
  "Eiffel Tower, Paris, France" rather than just "Eiffel Tower").
- Reverse geocoding precision depends on OSM coverage density at the given
  coordinates. Cities and major towns are well covered worldwide.
- Nominatim ToS requires a `User-Agent` header and caching — both are handled
  automatically by the API.
