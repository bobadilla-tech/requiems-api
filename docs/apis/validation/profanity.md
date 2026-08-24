# Profanity Filter API

Detect and censor profanity in text for content moderation.

## Status

✅ **Live** — Available at `POST /v1/validation/profanity`

## Endpoint

`POST /v1/validation/profanity`

## Request

```json
{
  "text": "Some text to check"
}
```

| Field | Type   | Required | Description       |
| ----- | ------ | -------- | ----------------- |
| text  | string | Yes      | The text to check |

## Response

```json
{
  "data": {
    "has_profanity": false,
    "censored": "Some text to check",
    "flagged_words": []
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field         | Type             | Description                                              |
| ------------- | ---------------- | -------------------------------------------------------- |
| has_profanity | boolean          | Whether any profanity was detected                       |
| censored      | string           | Input text with profane words replaced by `*` characters |
| flagged_words | array of strings | Deduplicated list of detected profane words (lowercase)  |

## Behaviour

- Detection is **case-insensitive** — `BULLSHIT`, `Bullshit`, and `bullshit` all
  match.
- Censoring replaces each character of a flagged word with `*`, preserving word
  length.
- Surrounding punctuation and whitespace are left unchanged.
- `flagged_words` contains each word only once, even if it appears multiple
  times in the input.

## Code Examples

### cURL

```bash
curl -X POST https://requiems.xyz/v1/validation/profanity \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"text": "Some text to check"}'
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/validation/profanity"
headers = {
    "requiems-api-key": "YOUR_API_KEY",
    "Content-Type": "application/json"
}
payload = {"text": "Some text to check"}

response = requests.post(url, headers=headers, json=payload)
result = response.json()['data']
print(f"Has profanity: {result['has_profanity']}")
print(f"Censored: {result['censored']}")
```

### JavaScript

```javascript
const response = await fetch(
  "https://requiems.xyz/v1/validation/profanity",
  {
    method: "POST",
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ text: "Some text to check" }),
  },
);

const { data } = await response.json();
console.log(`Has profanity: ${data.has_profanity}`);
console.log(`Censored: ${data.censored}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/validation/profanity')
request = Net::HTTP::Post.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'
request['Content-Type'] = 'application/json'
request.body = { text: 'Some text to check' }.to_json

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

data = JSON.parse(response.body)['data']
puts "Has profanity: #{data['has_profanity']}"
puts "Censored: #{data['censored']}"
```

## Error Codes

| Code                | Status | When                                     |
| ------------------- | ------ | ---------------------------------------- |
| `validation_failed` | 422    | The `text` field is missing or empty     |
| `bad_request`       | 400    | The request body is missing or malformed |
| `internal_error`    | 500    | Unexpected server error                  |
