# Lorem Ipsum Generator API

Generate classic Lorem Ipsum placeholder text for design mockups, prototypes,
and testing.

## Status

✅ **Live** - Available now at `GET /v1/text/lorem`

## Endpoint

`GET /v1/text/lorem`

## Query Parameters

| Parameter  | Type    | Required | Default | Range | Description                       |
| ---------- | ------- | -------- | ------- | ----- | --------------------------------- |
| paragraphs | integer | No       | 1       | 1-20  | Number of paragraphs to generate  |
| sentences  | integer | No       | 5       | 1-20  | Number of sentences per paragraph |

## Response

```json
{
  "data": {
    "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
    "paragraphs": 1,
    "wordCount": 45
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field      | Type    | Description                       |
| ---------- | ------- | --------------------------------- |
| text       | string  | Generated Lorem Ipsum text        |
| paragraphs | integer | Number of paragraphs generated    |
| wordCount  | integer | Total number of words in the text |

## Error Codes

| Code          | Status | When                                |
| ------------- | ------ | ----------------------------------- |
| `bad_request` | 400    | paragraphs must be between 1 and 20 |
| `bad_request` | 400    | sentences must be between 1 and 20  |

## Code Examples

### cURL

```bash
# Generate 3 paragraphs with 5 sentences each
curl "https://requiems.xyz/v1/text/lorem?paragraphs=3&sentences=5" \
  -H "requiems-api-key: YOUR_API_KEY"

# Generate 1 paragraph with 10 sentences
curl "https://requiems.xyz/v1/text/lorem?sentences=10" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/text/lorem"
headers = {"requiems-api-key": "YOUR_API_KEY"}
params = {"paragraphs": 3, "sentences": 5}

response = requests.get(url, headers=headers, params=params)
result = response.json()['data']
print(result['text'])
print(f"Word count: {result['wordCount']}")
```

### JavaScript

```javascript
const params = new URLSearchParams({ paragraphs: 3, sentences: 5 });
const response = await fetch(
  `https://requiems.xyz/v1/text/lorem?${params}`,
  {
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
    },
  },
);

const { data } = await response.json();
console.log(data.text);
console.log(`Word count: ${data.wordCount}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/text/lorem')
uri.query = URI.encode_www_form(paragraphs: 3, sentences: 5)

request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

data = JSON.parse(response.body)['data']
puts data['text']
puts "Word count: #{data['wordCount']}"
```

## Use Cases

- **Web Design Mockups** - Fill content areas in design prototypes and
  wireframes
- **Development Testing** - Generate realistic-looking text for testing layouts
  and typography
- **Content Placeholders** - Create temporary content while waiting for final
  copy
- **Typography Demos** - Showcase fonts and text styling with varied content
  lengths

## FAQ

**What is Lorem Ipsum?** Lorem Ipsum is dummy text used in the printing and
typesetting industry since the 1500s. It's a scrambled version of Latin text
that creates natural-looking placeholder content.

**Why use Lorem Ipsum instead of 'Test test test'?** Lorem Ipsum has a more
natural distribution of letters and word lengths, making it better for testing
typography and layout. It also looks more professional in mockups.

**Can I control both paragraphs and sentences?** Yes! Use the `paragraphs`
parameter to set how many paragraphs you want (1-20), and the `sentences`
parameter to set how many sentences per paragraph (1-20).

**Is the generated text always the same?** The text is generated using the
Lorelai library, which creates varied Lorem Ipsum content based on the classic
Lorem Ipsum corpus.
