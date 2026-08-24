# Working Days Calculator API

Calculate the number of working days (business days) between two dates, with
optional support for country-specific holidays.

## Status

✅ **Live** - Available now at `GET /v1/places/working-days`

## Endpoint

`GET /v1/places/working-days`

## Query Parameters

| Parameter   | Type   | Required | Description                                                                        |
| ----------- | ------ | -------- | ---------------------------------------------------------------------------------- |
| from        | string | Yes      | Start date in YYYY-MM-DD format (ISO 8601)                                         |
| to          | string | Yes      | End date in YYYY-MM-DD format (ISO 8601). Must be >= from date.                    |
| country     | string | No       | ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "FR"). Excludes holidays.       |
| subdivision | string | No       | ISO 3166-2 subdivision code for state/region (e.g., "NY", "CA"). Requires country. |

## Response

```json
{
  "data": {
    "working_days": 4,
    "from": "2024-02-23",
    "to": "2024-02-28",
    "country": "US",
    "subdivision": "NY"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field        | Type    | Description                                                          |
| ------------ | ------- | -------------------------------------------------------------------- |
| working_days | integer | Number of working days (excluding weekends and optionally holidays)  |
| from         | string  | Start date (echoed from request)                                     |
| to           | string  | End date (echoed from request)                                       |
| country      | string  | Country code (echoed from request, empty string if not provided)     |
| subdivision  | string  | Subdivision code (echoed from request, empty string if not provided) |

## Error Codes

| Code          | Status | When                                                           |
| ------------- | ------ | -------------------------------------------------------------- |
| `bad_request` | 400    | Required parameters missing, invalid date format, or to < from |

## Code Examples

### cURL

```bash
# Without country (only excludes weekends)
curl "https://requiems.xyz/v1/places/working-days?from=2024-02-23&to=2024-02-28" \
  -H "requiems-api-key: YOUR_API_KEY"

# With country (excludes weekends and US federal holidays)
curl "https://requiems.xyz/v1/places/working-days?from=2024-02-23&to=2024-02-28&country=US" \
  -H "requiems-api-key: YOUR_API_KEY"

# With country and subdivision (excludes weekends and US + NY holidays)
curl "https://requiems.xyz/v1/places/working-days?from=2024-02-23&to=2024-02-28&country=US&subdivision=NY" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/places/working-days"
headers = {"requiems-api-key": "YOUR_API_KEY"}
params = {
    "from": "2024-02-23",
    "to": "2024-02-28",
    "country": "US",
    "subdivision": "NY"
}

response = requests.get(url, headers=headers, params=params)
result = response.json()['data']
print(f"{result['working_days']} working days between {result['from']} and {result['to']}")
```

### JavaScript

```javascript
const params = new URLSearchParams({
  from: "2024-02-23",
  to: "2024-02-28",
  country: "US",
  subdivision: "NY",
});

const response = await fetch(
  `https://requiems.xyz/v1/places/working-days?${params}`,
  {
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
    },
  },
);

const { data } = await response.json();
console.log(
  `${data.working_days} working days between ${data.from} and ${data.to}`,
);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/places/working-days')
uri.query = URI.encode_www_form(
  from: '2024-02-23',
  to: '2024-02-28',
  country: 'US',
  subdivision: 'NY'
)

request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

result = JSON.parse(response.body)['data']
puts "#{result['working_days']} working days between #{result['from']} and #{result['to']}"
```

## Use Cases

- **Project Deadline Calculations** - Estimate realistic project timelines
- **Delivery Time Estimates** - Calculate shipping and delivery dates
- **Payroll and Billing** - Calculate working days for invoicing
- **SLA Tracking** - Monitor service level agreement compliance
- **Vacation Planning** - Calculate available work days

## FAQ

**What defines a working day?** By default, working days are Monday through
Friday, excluding weekends (Saturday and Sunday). When a country is specified,
public holidays for that country are also excluded.

**Which countries are supported for holiday calculations?** The API supports
holiday calendars for most countries worldwide using ISO 3166-1 alpha-2 country
codes. If a country is not recognized, the API returns only weekend-excluding
calculations.

**How are holidays determined?** The API uses the business-days-calculator
library which includes comprehensive holiday calendars for each country,
including federal/national holidays and optional state/regional holidays when a
subdivision is specified.

**What if the to date is before the from date?** The API will return a 400 Bad
Request error. The to date must be greater than or equal to the from date.

**Are partial days counted?** No. The calculation counts full calendar days
between the from and to dates (inclusive), then excludes weekends and holidays
as appropriate.
