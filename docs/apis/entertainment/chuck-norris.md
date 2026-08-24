# Chuck Norris API

## Status

✅ **Live** - Available now

## Overview

Get a random Chuck Norris fact from a curated built-in database. Every call
returns a different fact selected with a cryptographically secure random number
generator.

## Endpoints

### Get Random Chuck Norris Fact

**Endpoint:** `GET /v1/entertainment/chuck-norris`

Returns a randomly selected Chuck Norris fact.

#### Response

```json
{
  "data": {
    "id": "cn_0",
    "fact": "Chuck Norris can divide by zero."
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

#### Response Fields

| Field  | Type   | Description                                       |
| ------ | ------ | ------------------------------------------------- |
| `id`   | string | Unique fact identifier in the format `cn_<index>` |
| `fact` | string | The Chuck Norris fact text                        |

#### Error Responses

| Status | Description                       |
| ------ | --------------------------------- |
| `401`  | Missing `requiems-api-key` header |
| `403`  | Invalid API key                   |

## Code Examples

### cURL

```bash
curl https://requiems.xyz/v1/entertainment/chuck-norris \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/entertainment/chuck-norris"
headers = {"requiems-api-key": "YOUR_API_KEY"}

response = requests.get(url, headers=headers)
result = response.json()['data']
print(f"[{result['id']}] {result['fact']}")
```

### JavaScript

```javascript
const response = await fetch(
  "https://requiems.xyz/v1/entertainment/chuck-norris",
  { headers: { "requiems-api-key": "YOUR_API_KEY" } },
);
const { data } = await response.json();
console.log(`[${data.id}] ${data.fact}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/entertainment/chuck-norris')
request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

result = JSON.parse(response.body)['data']
puts "[#{result['id']}] #{result['fact']}"
```
