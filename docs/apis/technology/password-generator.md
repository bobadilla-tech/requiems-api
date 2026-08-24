# Password Generator API

Generate cryptographically secure random passwords with customizable complexity.

## Status

✅ **Live** - Available now at `GET /v1/technology/password`

## Endpoint

`GET /v1/technology/password`

## Query Parameters

| Parameter | Type    | Required | Default | Range | Description                                     |
| --------- | ------- | -------- | ------- | ----- | ----------------------------------------------- |
| length    | integer | No       | 16      | 8-128 | Password length                                 |
| uppercase | boolean | No       | false   | -     | Include uppercase letters (A-Z)                 |
| numbers   | boolean | No       | false   | -     | Include numbers (0-9)                           |
| symbols   | boolean | No       | false   | -     | Include special characters (!@#$%^&*()-_=+etc.) |

## Response

```json
{
  "data": {
    "password": "aB3#cDeFgHiJkLmN",
    "length": 16,
    "strength": "strong"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field    | Type    | Description                                         |
| -------- | ------- | --------------------------------------------------- |
| password | string  | The generated password                              |
| length   | integer | Length of the generated password                    |
| strength | string  | Password strength assessment (weak, medium, strong) |

## Password Strength

Strength is calculated based on length and character set diversity:

- **Weak** (≤1 point) - Short with limited character sets
- **Medium** (2-3 points) - Reasonable length with some variety
- **Strong** (≥4 points) - 16+ characters with multiple character sets

Scoring: 1 point for length ≥ 8, +1 for length ≥ 16, +1 per enabled charset

## Error Codes

| Code             | Status | When                               |
| ---------------- | ------ | ---------------------------------- |
| `bad_request`    | 400    | length must be between 8 and 128   |
| `internal_error` | 500    | Failed to generate password (rare) |

## Code Examples

### cURL

```bash
# Generate 16-character password with all character types
curl "https://requiems.xyz/v1/technology/password?length=16&uppercase=true&numbers=true&symbols=true" \
  -H "requiems-api-key: YOUR_API_KEY"

# Generate simple 12-character password (lowercase only)
curl "https://requiems.xyz/v1/technology/password?length=12" \
  -H "requiems-api-key: YOUR_API_KEY"

# Generate 20-character alphanumeric password
curl "https://requiems.xyz/v1/technology/password?length=20&uppercase=true&numbers=true" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/technology/password"
headers = {"requiems-api-key": "YOUR_API_KEY"}
params = {
    "length": 16,
    "uppercase": True,
    "numbers": True,
    "symbols": True
}

response = requests.get(url, headers=headers, params=params)
result = response.json()['data']
print(f"Password: {result['password']}")
print(f"Strength: {result['strength']}")
```

### JavaScript

```javascript
const params = new URLSearchParams({
  length: 16,
  uppercase: true,
  numbers: true,
  symbols: true,
});

const response = await fetch(
  `https://requiems.xyz/v1/technology/password?${params}`,
  {
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
    },
  },
);

const { data } = await response.json();
console.log(`Password: ${data.password}`);
console.log(`Strength: ${data.strength}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/technology/password')
uri.query = URI.encode_www_form(
  length: 16,
  uppercase: true,
  numbers: true,
  symbols: true
)

request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

result = JSON.parse(response.body)['data']
puts "Password: #{result['password']}"
puts "Strength: #{result['strength']}"
```

## Use Cases

- **User Registration Systems** - Generate temporary passwords for new users
- **Password Reset Functionality** - Create secure reset tokens
- **Security Tools** - Provide password suggestions to users
- **Temporary Credentials** - Generate one-time access codes
- **API Key Generation** - Create random secure keys

## FAQ

**Are the passwords truly random?** Yes. The API uses Go's crypto/rand package
which provides cryptographically secure random number generation suitable for
security-sensitive applications.

**Can I generate passwords with only specific character sets?** Yes. By default,
lowercase letters are always included. You can add uppercase, numbers, and/or
symbols by setting the respective parameters to true.

**Is there a minimum length requirement?** Yes, passwords must be at least 8
characters long. The maximum length is 128 characters.

**Does the API guarantee character distribution?** Yes. When you enable a
character set (uppercase, numbers, symbols), the algorithm guarantees at least
one character from each enabled set will be included. Characters are then
shuffled using Fisher-Yates algorithm for randomness.

**How is password strength calculated?** Strength is based on length and
character set diversity. Weak (≤1 point) - short with limited character sets.
Medium (2-3 points) - reasonable length with some variety. Strong (≥4 points) -
16+ characters with multiple character sets.
