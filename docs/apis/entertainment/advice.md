# Random Advice API

Get random pieces of advice and wisdom for inspiration, daily motivation, or
content generation.

## Status

✅ **Live** - Available now at `GET /v1/entertainment/advice`

## Endpoint

`GET /v1/entertainment/advice`

## Query Parameters

None required.

## Response

```json
{
  "data": {
    "id": 42,
    "advice": "Don't compare yourself to others. Compare yourself to the person you were yesterday."
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field  | Type    | Description                      |
| ------ | ------- | -------------------------------- |
| id     | integer | Unique identifier for the advice |
| advice | string  | A random piece of advice         |

## Error Codes

| Code                  | Status | When                            |
| --------------------- | ------ | ------------------------------- |
| `service_unavailable` | 503    | No advice available in database |

## Code Examples

### cURL

```bash
curl https://requiems.xyz/v1/entertainment/advice \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/entertainment/advice"
headers = {"requiems-api-key": "YOUR_API_KEY"}

response = requests.get(url, headers=headers)
advice = response.json()['data']
print(f"Advice #{advice['id']}: {advice['advice']}")
```

### JavaScript

```javascript
const response = await fetch(
  "https://requiems.xyz/v1/entertainment/advice",
  {
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
    },
  },
);

const { data } = await response.json();
console.log(`Advice #${data.id}: ${data.advice}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/entertainment/advice')
request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

data = JSON.parse(response.body)['data']
puts "Advice ##{data['id']}: #{data['advice']}"
```

## Use Cases

- **Daily Motivation Apps** - Provide users with daily wisdom and inspiration
- **Chatbot Responses** - Add helpful advice to conversational AI responses
- **Content Placeholders** - Fill content areas during development
- **Quote Widgets** - Display rotating advice on websites and dashboards

## Features

- Curated collection of advice and wisdom
- Simple REST API with no parameters
- Fast response times
- Unique ID for each piece of advice

## FAQ

**Can I request specific types of advice?** Currently, the API returns random
advice from our curated collection. Category filtering is planned for a future
update.

**How many pieces of advice are in the database?** Our collection contains over
200 pieces of curated advice and wisdom, and we're constantly adding more.

**Will I get the same advice on consecutive calls?** No, advice is selected
randomly on each request, so consecutive calls will typically return different
advice.
