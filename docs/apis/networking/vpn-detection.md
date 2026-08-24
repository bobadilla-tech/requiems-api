# VPN & Proxy Detection API

## Status

✅ **Live** — Available at `GET /v1/networking/ip/vpn/{ip}`

## Overview

Detect if an IP address belongs to a VPN, proxy, Tor exit node, or hosting
provider. Returns threat scores and fraud indicators for fraud prevention, risk
assessment, and bot detection.

## Endpoint

### Check IP Address

`GET /v1/networking/ip/vpn/{ip}`

### Path Parameters

| Parameter | Required | Description                        |
| --------- | -------- | ---------------------------------- |
| `ip`      | Yes      | IP address to check (IPv4 or IPv6) |

### Response

```json
{
  "data": {
    "ip": "8.8.8.8",
    "is_vpn": false,
    "is_proxy": false,
    "is_tor": false,
    "is_hosting": true,
    "score": 1,
    "threat": "low",
    "fraud_score": 0,
    "asn_org": "GOOGLE-ASN"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

### Detection Types

| Field        | Description                                                         |
| ------------ | ------------------------------------------------------------------- |
| `is_vpn`     | True when the IP belongs to a known VPN provider                    |
| `is_proxy`   | True when the IP is a known public or web proxy                     |
| `is_tor`     | True when the IP is a known Tor exit node                           |
| `is_hosting` | True when the IP belongs to a data-centre or hosting provider (DCH) |

### Threat Scoring

| Field         | Type   | Description                                                                                       |
| ------------- | ------ | ------------------------------------------------------------------------------------------------- |
| `score`       | int    | Raw threat score (0-9+). Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1 |
| `threat`      | string | Threat level derived from score: none, low, medium, high, or critical                             |
| `fraud_score` | int    | Fraud risk score from 0 (no risk) to 100 (high risk). Available with IP2Proxy PX5+ database       |
| `asn_org`     | string | Organization name owning the Autonomous System containing the IP (e.g., "DIGITALOCEAN-ASN")       |

### Threat Levels

| Level      | Score Range | Interpretation                        |
| ---------- | ----------- | ------------------------------------- |
| `none`     | 0           | No threat indicators detected         |
| `low`      | 1           | Minimal risk (e.g., hosting provider) |
| `medium`   | 2-3         | Moderate risk (VPN/proxy detected)    |
| `high`     | 4-5         | Elevated risk (multiple indicators)   |
| `critical` | 6+          | High risk (Tor + VPN/proxy combined)  |

### Error Codes

| Code             | Status | When                             |
| ---------------- | ------ | -------------------------------- |
| `bad_request`    | 400    | IP address is missing or invalid |
| `internal_error` | 500    | Unexpected server error          |

## Code Examples

### cURL

```bash
curl "https://requiems.xyz/v1/networking/ip/vpn/8.8.8.8" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

url = "https://requiems.xyz/v1/networking/ip/vpn/8.8.8.8"
headers = {"requiems-api-key": "YOUR_API_KEY"}

response = requests.get(url, headers=headers)
result = response.json()['data']

print(f"IP: {result['ip']}")
print(f"Is VPN: {result['is_vpn']}")
print(f"Threat Level: {result['threat']}")
print(f"Fraud Score: {result['fraud_score']}")
```

### JavaScript

```javascript
const response = await fetch(
  "https://requiems.xyz/v1/networking/ip/vpn/8.8.8.8",
  {
    headers: { "requiems-api-key": "YOUR_API_KEY" },
  },
);

const { data } = await response.json();
console.log(`IP: ${data.ip}`);
console.log(`Is VPN: ${data.is_vpn}`);
console.log(`Threat Level: ${data.threat}`);
console.log(`Fraud Score: ${data.fraud_score}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

uri = URI('https://requiems.xyz/v1/networking/ip/vpn/8.8.8.8')
request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

data = JSON.parse(response.body)['data']
puts "IP: #{data['ip']}"
puts "Is VPN: #{data['is_vpn']}"
puts "Threat Level: #{data['threat']}"
puts "Fraud Score: #{data['fraud_score']}"
```

## Use Cases

- **Fraud Prevention** - Identify high-risk IP addresses in e-commerce and
  fintech
- **Risk Scoring** - Assess user authentication risk based on IP reputation
- **Bot Detection** - Detect users hiding behind VPNs, proxies, or Tor
- **Compliance** - Meet regulatory requirements for IP intelligence
- **Geo-blocking Circumvention** - Identify users attempting to bypass
  restrictions

## FAQ

**What databases are used for detection?** The API uses the IP2Proxy database
for VPN, proxy, Tor, and hosting detection. ASN information is retrieved from a
separate MaxMind ASN database.

**Why is fraud_score 0 for some IPs?** The fraud_score is only available when
using IP2Proxy database tier PX5 or higher. Lower tiers return 0, indicating the
value is unavailable for the current database.

**Does this support IPv6?** Yes, both IPv4 and IPv6 addresses are fully
supported.

**What does is_hosting mean?** When `is_hosting` is true, the IP belongs to a
data-centre or cloud hosting provider (DCH in IP2Proxy terminology). This is
often an indicator of automated traffic or server-based access rather than a
residential user.
