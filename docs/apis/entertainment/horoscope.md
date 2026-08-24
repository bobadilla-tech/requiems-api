# Horoscope API

## Status

✅ **Live** - Available now

## Overview

Get daily horoscope readings for any of the 12 zodiac signs. Each reading is
deterministically generated per sign per day, so the same sign returns the same
reading throughout the day.

## Endpoints

### Get Daily Horoscope

**Endpoint:** `GET /v1/entertainment/horoscope/:sign`

Returns a daily horoscope reading for the given zodiac sign.

#### Path Parameters

| Parameter | Type   | Required | Description                                                |
| --------- | ------ | -------- | ---------------------------------------------------------- |
| `sign`    | string | Yes      | Zodiac sign (case-insensitive). One of the 12 signs below. |

#### Supported Signs

`aries`, `taurus`, `gemini`, `cancer`, `leo`, `virgo`, `libra`, `scorpio`,
`sagittarius`, `capricorn`, `aquarius`, `pisces`

#### Response

```json
{
  "data": {
    "sign": "aries",
    "date": "2024-12-15",
    "horoscope": "Today is a great day for new beginnings. Trust your instincts and take that first step toward your goals.",
    "lucky_number": 7,
    "mood": "energetic"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

#### Response Fields

| Field          | Type    | Description                               |
| -------------- | ------- | ----------------------------------------- |
| `sign`         | string  | Normalized zodiac sign (lowercase)        |
| `date`         | string  | Today's date in `YYYY-MM-DD` format (UTC) |
| `horoscope`    | string  | Daily horoscope reading                   |
| `lucky_number` | integer | Lucky number for the day (1–99)           |
| `mood`         | string  | Suggested mood for the day                |

#### Error Responses

| Status | Description                       |
| ------ | --------------------------------- |
| `400`  | Invalid zodiac sign provided      |
| `401`  | Missing `requiems-api-key` header |
| `403`  | Invalid API key                   |

## Code Examples

### cURL

```bash
curl https://requiems.xyz/v1/entertainment/horoscope/aries \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/entertainment/horoscope/aries"
headers = {"requiems-api-key": "YOUR_API_KEY"}

response = requests.get(url, headers=headers)
result = response.json()['data']
print(f"{result['sign']}: {result['horoscope']}")
print(f"Lucky number: {result['lucky_number']}, Mood: {result['mood']}")
```

### JavaScript

```javascript
const response = await fetch(
  "https://requiems.xyz/v1/entertainment/horoscope/aries",
  { headers: { "requiems-api-key": "YOUR_API_KEY" } },
);
const { data } = await response.json();
console.log(`${data.sign}: ${data.horoscope}`);
console.log(`Lucky number: ${data.lucky_number}, Mood: ${data.mood}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/entertainment/horoscope/aries')
request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

result = JSON.parse(response.body)['data']
puts "#{result['sign']}: #{result['horoscope']}"
puts "Lucky number: #{result['lucky_number']}, Mood: #{result['mood']}"
```
