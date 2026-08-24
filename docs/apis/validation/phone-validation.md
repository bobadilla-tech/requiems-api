# Validate Phone API

Validate phone numbers globally. Detect carrier, country, number type, and VOIP
or virtual risk using only phone metadata. No external lookups.

## Endpoints

| Method | Path                         | Description               |
| ------ | ---------------------------- | ------------------------- |
| GET    | `/v1/validation/phone`       | Validate a single number  |
| POST   | `/v1/validation/phone/batch` | Validate up to 50 numbers |

---

## GET /v1/validation/phone

### Query Parameters

| Parameter | Type   | Required | Description                                         |
| --------- | ------ | -------- | --------------------------------------------------- |
| `number`  | string | Yes      | Phone number in E.164 format (e.g. `+447400123456`) |

### Response

```json
{
  "data": {
    "number": "+447400123456",
    "valid": true,
    "country": "GB",
    "type": "mobile",
    "formatted": "+44 7400 123456",
    "carrier": {
      "name": "Three",
      "source": "metadata"
    },
    "risk": {
      "is_voip": false,
      "is_virtual": false
    }
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

When the number is invalid, `valid` is `false` and all optional fields are
omitted:

```json
{
  "data": {
    "number": "12345",
    "valid": false
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

---

## POST /v1/validation/phone/batch

Validates up to 50 numbers in one request. Results are returned in the same
order as the input.

### Request Body

```json
{
  "numbers": ["+447400123456", "+12015551234", "12345"]
}
```

| Field     | Type     | Required | Description                              |
| --------- | -------- | -------- | ---------------------------------------- |
| `numbers` | string[] | Yes      | Array of phone numbers (min: 1, max: 50) |

### Response

```json
{
  "data": {
    "results": [
      {
        "number": "+447400123456",
        "valid": true,
        "country": "GB",
        "type": "mobile",
        "formatted": "+44 7400 123456",
        "carrier": { "name": "Three", "source": "metadata" },
        "risk": { "is_voip": false, "is_virtual": false }
      },
      {
        "number": "+12015551234",
        "valid": true,
        "country": "US",
        "type": "landline_or_mobile",
        "formatted": "+1 201-555-1234",
        "risk": { "is_voip": false, "is_virtual": false }
      },
      {
        "number": "12345",
        "valid": false
      }
    ],
    "total": 3
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

---

## Number Types

| Value                | Description                                     |
| -------------------- | ----------------------------------------------- |
| `mobile`             | Mobile / cell phone                             |
| `landline`           | Fixed-line telephone                            |
| `landline_or_mobile` | Cannot be distinguished (common for US numbers) |
| `toll_free`          | Toll-free number                                |
| `premium_rate`       | Premium-rate number                             |
| `shared_cost`        | Shared-cost number                              |
| `voip`               | Voice over IP                                   |
| `personal_number`    | Personal number                                 |
| `pager`              | Pager                                           |
| `uan`                | Universal Access Number                         |
| `voicemail`          | Voicemail number                                |
| `unknown`            | Type could not be determined                    |

---

## Response Fields

| Field             | Type    | Description                                                                                              |
| ----------------- | ------- | -------------------------------------------------------------------------------------------------------- |
| `number`          | string  | The original number as supplied in the request                                                           |
| `valid`           | boolean | Whether the number is a valid, dialable phone number                                                     |
| `country`         | string  | ISO 3166-1 alpha-2 country code (omitted when valid is false)                                            |
| `type`            | string  | Number type (see table above, omitted when valid is false)                                               |
| `formatted`       | string  | International format of the number (omitted when valid is false)                                         |
| `carrier.name`    | string  | Carrier name from prefix metadata (omitted when carrier cannot be determined)                            |
| `carrier.source`  | string  | How the carrier was determined. Always `"metadata"` when present                                         |
| `risk.is_voip`    | boolean | `true` when the number type is `voip`                                                                    |
| `risk.is_virtual` | boolean | `true` for types not tied to a physical SIM or fixed line (voip, personal_number, uan, pager, voicemail) |

---

## Error Codes

| Code                | Status | When                                                       |
| ------------------- | ------ | ---------------------------------------------------------- |
| `bad_request`       | 400    | The `number` query parameter is missing (single endpoint)  |
| `validation_failed` | 422    | The `numbers` array is missing, empty, or exceeds 50 items |

---

## Code Examples

### cURL

```bash
# Single
curl "https://requiems.xyz/v1/validation/phone?number=%2B447400123456" \
  -H "requiems-api-key: YOUR_API_KEY"

# Batch
curl -X POST "https://requiems.xyz/v1/validation/phone/batch" \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"numbers":["+447400123456","+12015551234"]}'
```

### Python

```python
import requests

headers = {"requiems-api-key": "YOUR_API_KEY"}

# Single
r = requests.get(
    "https://requiems.xyz/v1/validation/phone",
    headers=headers,
    params={"number": "+447400123456"},
)
result = r.json()["data"]
if result["valid"]:
    print(f"Valid {result['type']} number in {result['country']}")
    if result.get("carrier"):
        print(f"Carrier: {result['carrier']['name']}")
    if result.get("risk", {}).get("is_voip"):
        print("Warning: VOIP number detected")

# Batch
r = requests.post(
    "https://requiems.xyz/v1/validation/phone/batch",
    headers={**headers, "Content-Type": "application/json"},
    json={"numbers": ["+447400123456", "+12015551234"]},
)
for item in r.json()["data"]["results"]:
    print(item["number"], item["valid"])
```

### JavaScript

```javascript
const headers = { "requiems-api-key": "YOUR_API_KEY" };

// Single
const single = await fetch(
  `https://requiems.xyz/v1/validation/phone?number=${
    encodeURIComponent("+447400123456")
  }`,
  { headers },
);
const { data } = await single.json();
if (data.valid) {
  console.log(`Valid ${data.type} in ${data.country}`);
  if (data.carrier) console.log(`Carrier: ${data.carrier.name}`);
  if (data.risk?.is_voip) console.warn("VOIP number");
}

// Batch
const batch = await fetch(
  "https://requiems.xyz/v1/validation/phone/batch",
  {
    method: "POST",
    headers: { ...headers, "Content-Type": "application/json" },
    body: JSON.stringify({ numbers: ["+447400123456", "+12015551234"] }),
  },
);
const { data: batchData } = await batch.json();
batchData.results.forEach((r) => console.log(r.number, r.valid));
```

### Ruby

```ruby
require 'net/http'
require 'json'

headers = { 'requiems-api-key' => 'YOUR_API_KEY' }

# Single
uri = URI('https://requiems.xyz/v1/validation/phone')
uri.query = URI.encode_www_form(number: '+447400123456')
req = Net::HTTP::Get.new(uri, headers)
res = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) { |h| h.request(req) }
data = JSON.parse(res.body)['data']
puts "#{data['type']} in #{data['country']}" if data['valid']

# Batch
uri = URI('https://requiems.xyz/v1/validation/phone/batch')
req = Net::HTTP::Post.new(uri, headers.merge('Content-Type' => 'application/json'))
req.body = { numbers: ['+447400123456', '+12015551234'] }.to_json
res = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) { |h| h.request(req) }
JSON.parse(res.body)['data']['results'].each { |r| puts "#{r['number']}: #{r['valid']}" }
```

---

## Use Cases

- **User Registration** - Validate phone numbers during signup and reject
  non-dialable inputs
- **VOIP Screening** - Flag VOIP or virtual numbers before sending SMS
  verification codes or one-time passwords
- **Fraud Prevention** - Combine number type and carrier metadata to identify
  suspicious phone patterns
- **Data Normalization** - Store phone numbers in a consistent international
  format to prevent duplicates

---

## FAQ

**What format should I use for the phone number?** Always include the country
calling code prefixed with a plus sign (E.164 format), for example +447400123456
or +12015551234. Numbers without a country code cannot be validated reliably.

**What does landline_or_mobile mean?** Some numbering plans (notably the United
States) do not distinguish between mobile and landline numbers at the format
level. When the type cannot be determined more precisely, the API returns
`landline_or_mobile`.

**What happens when a number is invalid?** The API returns HTTP 200 with `valid`
set to `false`. All optional fields (`country`, `type`, `formatted`, `carrier`,
`risk`) are omitted from the response.

**Will carrier always be present for valid numbers?** No. Carrier detection
relies on prefix metadata and coverage varies by country. When the carrier
cannot be determined, the `carrier` object is omitted entirely.

**How is VOIP detection done?** VOIP and virtual flags come from the phone
number type in libphonenumber metadata. No external lookups are made. A number
is flagged as VOIP when its type is `voip`, and as virtual when its type is
`voip`, `personal_number`, `uan`, `pager`, or `voicemail`.
