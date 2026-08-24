# Data Format Conversion API

## Status

✅ **Live** - Available now

## Overview

Convert structured data between JSON, YAML, CSV, XML, and TOML in a single API
call. Accepts content in any supported format and returns it serialized in the
target format.

## Base URL

All endpoints are mounted under `/v1/convert`

## Endpoint

### Convert Format

**Endpoint:** `POST /v1/technology/format`

**Request body:**

```json
{
  "from": "json",
  "to": "yaml",
  "content": "{\"name\": \"Alice\", \"age\": 30}"
}
```

| Field     | Type   | Required | Description                                            |
| --------- | ------ | -------- | ------------------------------------------------------ |
| `from`    | string | ✅       | Source format: `json`, `yaml`, `csv`, `xml`, or `toml` |
| `to`      | string | ✅       | Target format: `json`, `yaml`, `csv`, `xml`, or `toml` |
| `content` | string | ✅       | The content to convert, as a string                    |

**Response:** `200 OK`

```json
{
  "data": {
    "result": "age: 30\nname: Alice\n"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
curl -X POST https://requiems.xyz/v1/technology/format \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"from":"json","to":"yaml","content":"{\"name\":\"Alice\",\"age\":30}"}'
```

---

## Supported Formats

| Format | `from`/`to` value | Notes                                                          |
| ------ | ----------------- | -------------------------------------------------------------- |
| JSON   | `json`            | RFC 8259. Numbers preserve int/float types.                    |
| YAML   | `yaml`            | YAML 1.2 via `gopkg.in/yaml.v3`.                               |
| CSV    | `csv`             | First row is the header. Requires array of objects for output. |
| XML    | `xml`             | Input parsed to map; output wrapped in `<root>` element.       |
| TOML   | `toml`            | Top-level value must be an object (no bare arrays/scalars).    |

---

## Format-Specific Notes

### CSV

- **Input (from: csv):** The first row is treated as the header row. Each
  subsequent row becomes a JSON object keyed by the headers. Rows with more
  columns than the header return a `422 invalid_csv` error.
- **Output (to: csv):** The input must be a JSON array of objects with
  consistent string keys. The keys of the first object are used as headers
  (sorted alphabetically). Non-object array elements return a
  `422
  conversion_error`.

### XML

- **Input (from: xml):** The root element is unwrapped. Repeated sibling
  elements with the same tag name are collapsed into a slice. Text-only elements
  are returned as `{"#text": "value"}`.
- **Output (to: xml):** The entire structure is wrapped in a `<root>` element.
  Arrays of items are each wrapped in an `<item>` element.

### TOML

- **Output (to: toml):** TOML requires a top-level table (object). Attempting to
  convert a JSON array or scalar to TOML returns a `422 conversion_error`.

---

## Size Limit

The `content` field is capped at **512 KB**. Requests exceeding this return
`413 content_too_large`.

---

## Error Responses

| HTTP Status | Error Code          | When                                                     |
| ----------- | ------------------- | -------------------------------------------------------- |
| 413         | `content_too_large` | `content` exceeds 512 KB                                 |
| 422         | `validation_failed` | Missing or invalid `from`/`to`/`content` field           |
| 422         | `invalid_json`      | `content` is not valid JSON (when `from=json`)           |
| 422         | `invalid_yaml`      | `content` is not valid YAML (when `from=yaml`)           |
| 422         | `invalid_csv`       | `content` is not valid CSV or row/header column mismatch |
| 422         | `invalid_xml`       | `content` is not valid XML (when `from=xml`)             |
| 422         | `invalid_toml`      | `content` is not valid TOML (when `from=toml`)           |
| 422         | `conversion_error`  | Data structure is incompatible with the target format    |
| 500         | `internal_error`    | Unexpected failure during serialization                  |

---

## Code Examples

### cURL

```bash
# JSON → YAML
curl -X POST https://requiems.xyz/v1/technology/format \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"from":"json","to":"yaml","content":"{\"name\":\"Alice\",\"age\":30}"}'

# CSV → JSON
curl -X POST https://requiems.xyz/v1/technology/format \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"from":"csv","to":"json","content":"name,age\nAlice,30\nBob,25"}'

# JSON → TOML
curl -X POST https://requiems.xyz/v1/technology/format \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"from":"json","to":"toml","content":"{\"title\":\"Config\",\"debug\":false}"}'
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/technology/format"
headers = {
    "requiems-api-key": "YOUR_API_KEY",
    "Content-Type": "application/json",
}

# JSON → YAML
response = requests.post(url, json={
    "from": "json",
    "to": "yaml",
    "content": '{"name": "Alice", "age": 30}',
}, headers=headers)
print(response.json()["data"]["result"])
# age: 30
# name: Alice

# TOML → JSON
response = requests.post(url, json={
    "from": "toml",
    "to": "json",
    "content": 'name = "Alice"\nage = 30\n',
}, headers=headers)
print(response.json()["data"]["result"])
```

### JavaScript

```javascript
const url = "https://requiems.xyz/v1/technology/format";
const headers = {
  "requiems-api-key": "YOUR_API_KEY",
  "Content-Type": "application/json",
};

// JSON → XML
const res = await fetch(url, {
  method: "POST",
  headers,
  body: JSON.stringify({
    from: "json",
    to: "xml",
    content: JSON.stringify({ name: "Alice", age: 30 }),
  }),
});

const { data } = await res.json();
console.log(data.result);
// <root>
//   <age>30</age>
//   <name>Alice</name>
// </root>
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/technology/format')
request = Net::HTTP::Post.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'
request['Content-Type'] = 'application/json'
request.body = {
  from: 'json',
  to: 'yaml',
  content: '{"name":"Alice","age":30}'
}.to_json

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

puts JSON.parse(response.body)['data']['result']
# age: 30
# name: Alice
```
