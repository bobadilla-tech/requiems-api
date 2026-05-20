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
curl https://api.requiems.xyz/v1/entertainment/advice \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://api.requiems.xyz/v1/entertainment/advice"
headers = {"requiems-api-key": "YOUR_API_KEY"}

response = requests.get(url, headers=headers)
advice = response.json()['data']
print(f"Advice #{advice['id']}: {advice['advice']}")
```

### JavaScript

```javascript
const response = await fetch(
  "https://api.requiems.xyz/v1/entertainment/advice",
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

uri = URI('https://api.requiems.xyz/v1/entertainment/advice')
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

## POST /v1/entertainment/advice/batch

Returns multiple random pieces of advice in a single request.

### Request Body

```json
{
  "count": 3
}
```

| Field   | Type    | Required | Description                                                  |
| ------- | ------- | -------- | ------------------------------------------------------------ |
| `count` | integer | ✅       | Number of advice items to return. Must be greater than zero. |

### Response

```json
{
  "data": {
    "results": [
      {
        "id": 12,
        "advice": "Stay consistent."
      },
      {
        "id": 87,
        "advice": "Do one thing every day that scares you."
      },
      {
        "id": 34,
        "advice": "Don't compare yourself to others."
      }
    ]
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field              | Type    | Description                      |
| ------------------ | ------- | -------------------------------- |
| `results`          | array   | List of random advice items      |
| `results[].id`     | integer | Unique identifier for the advice |
| `results[].advice` | string  | A random piece of advice         |

### Error Codes

| Code                  | Status | When                            |
| --------------------- | ------ | ------------------------------- |
| `invalid_request`     | 400    | Invalid or malformed JSON body  |
| `invalid_count`       | 422    | `count` is zero or negative     |
| `service_unavailable` | 503    | No advice available in database |

### Code Examples

#### cURL

```bash
curl -X POST https://api.requiems.xyz/v1/entertainment/advice/batch \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"count": 3}'
```

#### Python

```python
import requests

url = "https://api.requiems.xyz/v1/entertainment/advice/batch"
headers = {"requiems-api-key": "YOUR_API_KEY"}
payload = {"count": 3}

response = requests.post(url, json=payload, headers=headers)
results = response.json()['data']['results']
for item in results:
    print(f"Advice #{item['id']}: {item['advice']}")
```

#### JavaScript

```javascript
const response = await fetch(
  "https://api.requiems.xyz/v1/entertainment/advice/batch",
  {
    method: "POST",
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ count: 3 }),
  },
);

const { data } = await response.json();
data.results.forEach((item) => {
  console.log(`Advice #${item.id}: ${item.advice}`);
});
```

#### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://api.requiems.xyz/v1/entertainment/advice/batch')
request = Net::HTTP::Post.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'
request['Content-Type'] = 'application/json'
request.body = JSON.generate({ count: 3 })

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

results = JSON.parse(response.body)['data']['results']
results.each { |item| puts "Advice ##{item['id']}: #{item['advice']}" }
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
