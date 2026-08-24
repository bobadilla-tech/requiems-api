# Holidays API

Get public holiday information for any country and year. This endpoint provides
a list of national holidays with their dates.

## Status

✅ **Live** - Available now at `GET /v1/places/holidays`

## Endpoint

`GET /v1/places/holidays`

## Query Parameters

| Parameter | Type   | Required | Description                                  |
| --------- | ------ | -------- | -------------------------------------------- |
| `country` | string | Yes      | ISO 3166-1 alpha-2 country code (e.g., `US`) |
| `year`    | int    | Yes      | Year (e.g., `2025`)                          |

## Response

```json
{
  "data": {
    "country": "US",
    "year": 2025,
    "holidays": [
      {
        "date": "2025-01-01",
        "name": "New Year's Day"
      },
      {
        "date": "2025-07-04",
        "name": "Independence Day"
      }
    ],
    "total": 11
  },
  "metadata": {
    "timestamp": "2025-01-15T10:30:00Z"
  }
}
```

| Field      | Type             | Description                                   |
| ---------- | ---------------- | --------------------------------------------- |
| `country`  | string           | ISO 3166-1 alpha-2 country code               |
| `year`     | int              | Year                                          |
| `holidays` | array of objects | List of holidays with date and name           |
| `total`    | int              | Total number of holidays for the country/year |

### Holiday Object

| Field  | Type   | Description                 |
| ------ | ------ | --------------------------- |
| `date` | string | Date in `YYYY-MM-DD` format |
| `name` | string | Name of the holiday         |

## Error Codes

| Code          | Status | When                                                 |
| ------------- | ------ | ---------------------------------------------------- |
| `bad_request` | 400    | Missing or invalid country code or year              |
| `not_found`   | 404    | No holidays found for the specified country and year |

## Code Examples

### cURL

```bash
# Get US holidays for 2025
curl "https://requiems.xyz/v1/places/holidays?country=US&year=2025" \
  -H "requiems-api-key: YOUR_API_KEY"

# Get UK holidays for 2025
curl "https://requiems.xyz/v1/places/holidays?country=GB&year=2025" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/places/holidays"
headers = {"requiems-api-key": "YOUR_API_KEY"}
params = {"country": "US", "year": 2025}

response = requests.get(url, headers=headers, params=params)
result = response.json()['data']
print(f"Found {result['total']} holidays in {result['country']}")
for holiday in result['holidays']:
    print(f"  {holiday['date']}: {holiday['name']}")
```

### JavaScript

```javascript
const params = new URLSearchParams({
  country: "US",
  year: 2025,
});

const response = await fetch(
  `https://requiems.xyz/v1/places/holidays?${params}`,
  {
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
    },
  },
);

const { data } = await response.json();
console.log(`Found ${data.total} holidays in ${data.country}`);
data.holidays.forEach((h) => console.log(`  ${h.date}: ${h.name}`));
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/places/holidays')
uri.query = URI.encode_www_form(country: 'US', year: 2025)

request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

result = JSON.parse(response.body)['data']
puts "Found #{result['total']} holidays in #{result['country']}"
result['holidays'].each do |h|
  puts "  #{h['date']}: #{h['name']}"
end
```

## Use Cases

- **Calendar Applications** - Display national holidays in scheduling apps
- **Compliance Tools** - Track regulatory deadlines and observances
- **Travel Planning** - Account for local holidays when planning trips
- **HR Systems** - Manage employee leave and holiday entitlements
- **Business Intelligence** - Analyze holiday patterns across regions

## FAQ

**Which countries are supported?** The API supports holidays for over 190
countries using ISO 3166-1 alpha-2 country codes. This includes all UN member
states plus many territories and dependencies.

**What types of holidays are included?** The API returns national and public
holidays for each country, including federal holidays, bank holidays, and widely
observed holidays. Religious and regional holidays may vary by country.

**Can I get holidays for multiple years?** Currently, holidays are returned for
a single year per request. To get holidays across multiple years, make separate
requests for each year.

**Are holiday dates stable across years?** No. Holidays like "Thanksgiving" fall
on different dates each year. The API returns the correct date for the specified
year.

**What about moveable holidays?** Some holidays (like Easter) are calculated
based on astronomical events. The API handles these automatically and returns
the correct date for the specified year and country.
