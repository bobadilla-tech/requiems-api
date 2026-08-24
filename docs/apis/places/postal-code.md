# Postal Code API

Look up city, state, and geographic coordinates for any postal code worldwide.
Data sourced from the
[GeoNames postal code dataset](https://www.geonames.org/postal-codes/).

## Endpoint

`GET /v1/places/postal/{code}`

| Parameter | Location | Required | Description                                     |
| --------- | -------- | -------- | ----------------------------------------------- |
| `code`    | path     | Yes      | Postal/zip code to look up                      |
| `country` | query    | No       | ISO 3166-1 alpha-2 country code (default: `US`) |

## Response

```json
{
  "data": {
    "postal_code": "10001",
    "city": "New York City",
    "state": "New York",
    "country": "US",
    "lat": 40.7484,
    "lon": -73.9967
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Country Support

100+ countries are supported. Pass the ISO 3166-1 alpha-2 code as `country`:

```bash
# US zip code (default)
curl "https://requiems.xyz/v1/places/postal/10001" \
  -H "requiems-api-key: YOUR_API_KEY"

# UK postcode
curl "https://requiems.xyz/v1/places/postal/SW1A1AA?country=GB" \
  -H "requiems-api-key: YOUR_API_KEY"

# German postcode
curl "https://requiems.xyz/v1/places/postal/10115?country=DE" \
  -H "requiems-api-key: YOUR_API_KEY"
```

## Postal Code Format

Pass the code in the format native to the country. The API normalises case
automatically:

| Country | Example               |
| ------- | --------------------- |
| US      | `10001`               |
| UK      | `SW1A1AA` (no spaces) |
| Canada  | `M5V3L9` (no spaces)  |
| Germany | `10115`               |
| Japan   | `1000001`             |

## Error Codes

| Code             | Status | When                                        |
| ---------------- | ------ | ------------------------------------------- |
| `not_found`      | 404    | Postal code not found for the given country |
| `internal_error` | 500    | Unexpected failure                          |

## Notes

- When a postal code spans multiple places the primary administrative entry is
  returned.
- The dataset is updated periodically from GeoNames; some very new or obsolete
  codes may not be present.
