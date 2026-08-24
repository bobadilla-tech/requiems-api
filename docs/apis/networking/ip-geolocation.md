# IP Geolocation API

Get geolocation data for any IP address including country, city, ISP, and VPN
detection.

## Endpoints

### Lookup Caller IP

`GET /v1/networking/ip`

Returns geolocation information for the requesting client's IP address. Useful
when you want information about the user making the request without specifying
an IP explicitly.

**Parameters:** None

### Lookup Specific IP

`GET /v1/networking/ip/{ip}`

Returns geolocation information for a specific IP address.

| Name | Type   | Required | Description                          |
| ---- | ------ | -------- | ------------------------------------ |
| `ip` | string | Yes      | IP address to look up (IPv4 or IPv6) |

## Response Envelope

All responses are wrapped in the standard envelope:

```json
{
  "data": {
    "ip": "8.8.8.8",
    "country": "United States",
    "country_code": "US",
    "city": "Mountain View",
    "isp": "Google Public DNS",
    "is_vpn": false
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Response Fields

| Field          | Type    | Description                                    |
| -------------- | ------- | ---------------------------------------------- |
| `ip`           | string  | The IP address that was looked up              |
| `country`      | string  | Country name where the IP is located           |
| `country_code` | string  | Two-letter ISO country code (e.g., "US", "GB") |
| `city`         | string  | City name where the IP is located              |
| `isp`          | string  | Internet Service Provider                      |
| `is_vpn`       | boolean | True if the IP belongs to a known VPN          |

## Error Codes

| Code             | Status | When               |
| ---------------- | ------ | ------------------ |
| `bad_request`    | 400    | Invalid IP address |
| `internal_error` | 500    | Unexpected failure |

## Examples

### Lookup Caller IP

```bash
curl "https://requiems.xyz/v1/networking/ip" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Lookup Specific IP

```bash
curl "https://requiems.xyz/v1/networking/ip/8.8.8.8" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests

# Lookup specific IP
url = "https://requiems.xyz/v1/networking/ip/8.8.8.8"
headers = {"requiems-api-key": "YOUR_API_KEY"}

response = requests.get(url, headers=headers)
print(response.json())
```

## FAQ

**How accurate is the geolocation data?**

Accuracy varies by IP type. ISP and hosting provider IPs typically have
city-level accuracy (80-95%). Residential IPs can be accurate to within a few
kilometers. Mobile IPs are generally less accurate.

**Does this support IPv6?**

Yes, both IPv4 and IPv6 addresses are fully supported.

**What happens with private IP addresses?**

Private IP addresses (192.168.x.x, 10.x.x.x, 172.16-31.x.x) do not have
geolocation data. The API returns the IP with empty location fields.
