# Base64 Encode / Decode API

## Status

✅ **Live** - Available now

## Overview

Encode text to Base64 and decode Base64 strings back to plain text. Supports
both standard Base64 and URL-safe Base64 (base64url) variants.

## Base URL

All endpoints are mounted under `/v1/convert`

## Endpoints

### 1. Encode

Encode a plain-text string to Base64.

**Endpoint:** `POST /v1/technology/base64/encode`

**Request body:**

```json
{
  "value": "Hello, world!",
  "variant": "standard"
}
```

| Field     | Type   | Required | Description                              |
| --------- | ------ | -------- | ---------------------------------------- |
| `value`   | string | ✅       | The string to encode                     |
| `variant` | string | ❌       | `standard` (default) or `url` (URL-safe) |

**Response:** `200 OK`

```json
{
  "data": {
    "result": "SGVsbG8sIHdvcmxkIQ=="
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
curl -X POST https://api.requiems.xyz/v1/technology/base64/encode \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"value": "Hello, world!"}'
```

---

### 2. Decode

Decode a Base64-encoded string back to plain text.

**Endpoint:** `POST /v1/technology/base64/decode`

**Request body:**

```json
{
  "value": "SGVsbG8sIHdvcmxkIQ==",
  "variant": "standard"
}
```

| Field     | Type   | Required | Description                              |
| --------- | ------ | -------- | ---------------------------------------- |
| `value`   | string | ✅       | The Base64 string to decode              |
| `variant` | string | ❌       | `standard` (default) or `url` (URL-safe) |

**Response:** `200 OK`

```json
{
  "data": {
    "result": "Hello, world!"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
curl -X POST https://api.requiems.xyz/v1/technology/base64/decode \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"value": "SGVsbG8sIHdvcmxkIQ=="}'
```

---

## Variants

| Variant  | Key        | Alphabet                                          | Padding |
| -------- | ---------- | ------------------------------------------------- | ------- |
| Standard | `standard` | `A-Z a-z 0-9 + /`                                 | `=`     |
| URL-safe | `url`      | `A-Z a-z 0-9 - _` (replaces `+` → `-`, `/` → `_`) | `=`     |

Use `url` when the encoded output will appear in a URL, filename, or HTTP header
where `+` and `/` are problematic.

---

## Error Responses

| HTTP Status | Error Code       | Reason                         |
| ----------- | ---------------- | ------------------------------ |
| 400         | `bad_request`    | Missing or empty `value` field |
| 422         | `invalid_base64` | Input is not valid Base64      |

---

## Code Examples

### cURL

```bash
# Encode
curl -X POST https://api.requiems.xyz/v1/technology/base64/encode \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"value": "Hello, world!"}'

# Decode
curl -X POST https://api.requiems.xyz/v1/technology/base64/decode \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"value": "SGVsbG8sIHdvcmxkIQ=="}'

# URL-safe variant
curl -X POST https://api.requiems.xyz/v1/technology/base64/encode \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"value": "Hello, world!", "variant": "url"}'
```

### Python

```python
import requests

url_base = "https://api.requiems.xyz/v1/technology/base64"
headers = {
    "requiems-api-key": "YOUR_API_KEY",
    "Content-Type": "application/json",
}

# Encode
encode_resp = requests.post(f"{url_base}/encode", json={"value": "Hello, world!"}, headers=headers)
encoded = encode_resp.json()["data"]["result"]
print(f"Encoded: {encoded}")  # SGVsbG8sIHdvcmxkIQ==

# Decode
decode_resp = requests.post(f"{url_base}/decode", json={"value": encoded}, headers=headers)
decoded = decode_resp.json()["data"]["result"]
print(f"Decoded: {decoded}")  # Hello, world!
```

### JavaScript

```javascript
const BASE = "https://api.requiems.xyz/v1/technology/base64";
const headers = {
  "requiems-api-key": "YOUR_API_KEY",
  "Content-Type": "application/json",
};

// Encode
const encodeRes = await fetch(`${BASE}/encode`, {
  method: "POST",
  headers,
  body: JSON.stringify({ value: "Hello, world!" }),
});
const { result: encoded } = (await encodeRes.json()).data;
console.log(encoded); // SGVsbG8sIHdvcmxkIQ==

// Decode
const decodeRes = await fetch(`${BASE}/decode`, {
  method: "POST",
  headers,
  body: JSON.stringify({ value: encoded }),
});
const { result: decoded } = (await decodeRes.json()).data;
console.log(decoded); // Hello, world!
```

### Ruby

```ruby
require 'net/http'
require 'json'

def base64_request(path, payload)
  uri = URI("https://api.requiems.xyz/v1/technology/base64/#{path}")
  req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json', 'requiems-api-key' => 'YOUR_API_KEY')
  req.body = payload.to_json
  Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) { |http| http.request(req) }
end

encoded = JSON.parse(base64_request('encode', { value: 'Hello, world!' }).body).dig('data', 'result')
puts encoded  # SGVsbG8sIHdvcmxkIQ==

decoded = JSON.parse(base64_request('decode', { value: encoded }).body).dig('data', 'result')
puts decoded  # Hello, world!
```
