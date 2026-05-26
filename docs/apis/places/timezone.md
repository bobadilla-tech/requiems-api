# Timezone API

## Status

✅ **MVP** - Basic timezone lookup by coordinates and city name

## Overview

Get timezone information for any location. This endpoint provides timezone data
for geographic coordinates (latitude/longitude) or city names worldwide.

## Live Endpoints

### Get Timezone Information

**Endpoint:** `GET /v1/places/timezone`

Get timezone information for a location by coordinates or city name.

#### Query Parameters

| Parameter | Type   | Required | Description                                               |
| --------- | ------ | -------- | --------------------------------------------------------- |
| `lat`     | float  | \*       | Latitude (-90 to 90). Required when using coordinates.    |
| `lon`     | float  | \*       | Longitude (-180 to 180). Required when using coordinates. |
| `city`    | string | \*       | City name. Required when not using coordinates.           |

\* Either `city` **or** both `lat` + `lon` must be provided.

#### Example Requests

```
GET /v1/places/timezone?lat=51.5&lon=-0.1
GET /v1/places/timezone?city=Tokyo
```

#### Example Response

```json
{
  "data": {
    "timezone": "Europe/London",
    "offset": "+00:00",
    "current_time": "2024-12-15T14:30:00Z",
    "is_dst": false
  },
  "metadata": {
    "timestamp": "2024-12-15T14:30:00Z"
  }
}
```

#### Response Fields

| Field          | Type    | Description                                       |
| -------------- | ------- | ------------------------------------------------- |
| `timezone`     | string  | IANA timezone identifier (e.g. `"Europe/London"`) |
| `offset`       | string  | UTC offset in `+HH:MM` / `-HH:MM` format          |
| `current_time` | string  | Current UTC time in RFC 3339 format               |
| `is_dst`       | boolean | Whether the location is currently observing DST   |

---

## Batch IP VPN/Proxy Detection

**Endpoint:** `POST /v1/networking/ip/vpn/batch`

Analyze multiple IP addresses in a single request to determine if they belong to a VPN, proxy, Tor exit node, or hosting provider. Returns threat indicators, fraud scores, and ASN information for each IP.

---

### Request Body

| Field | Type     | Required | Description                                            |
| ----- | -------- | -------- | ------------------------------------------------------ |
| `ips` | string[] | Yes      | List of IP addresses (IPv4 or IPv6). Maximum 50 items. |


### Example Request

```http
POST /v1/networking/ip/vpn/batch
Content-Type: application/json
```

```json
{
  "ips": ["8.8.8.8", "1.1.1.1", "2001:4860:4860::8888"]
}
```

**Response:**

```json
{
  "data": {
    "results": [
      {
        "ip": "8.8.8.8",
        "is_vpn": false,
        "is_proxy": false,
        "is_tor": false,
        "is_hosting": true,
        "score": 1,
        "threat": 1,
        "fraud_score": 0,
        "asn_org": "GOOGLE-ASN"
      },
      {
        "ip": "1.1.1.1",
        "is_vpn": false,
        "is_proxy": false,
        "is_tor": false,
        "is_hosting": true,
        "score": 1,
        "threat": 1,
        "fraud_score": 0,
        "asn_org": "CLOUDFLARE-ASN"
      },
      {
        "ip": "invalid-ip",
        "is_vpn": false,
        "is_proxy": false,
        "is_tor": false,
        "is_hosting": false,
        "score": 0,
        "threat": 0,
        "fraud_score": 0,
        "asn_org": ""
      }
    ]
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

---

### Response Fields

| Field         | Type    | Description                                             |
| ------------- | ------- | ------------------------------------------------------- |
| `results`     | array   | List of IP detection results (preserves input order)    |
| `ip`          | string  | IP address analyzed                                     |
| `is_vpn`      | boolean | True if IP belongs to a VPN provider                    |
| `is_proxy`    | boolean | True if IP is a proxy                                   |
| `is_tor`      | boolean | True if IP is a Tor exit node                           |
| `is_hosting`  | boolean | True if IP belongs to a data-centre or hosting provider |
| `score`       | integer | Threat score (Tor +3, VPN +2, Proxy +2, Hosting +1)     |
| `threat`      | integer | 0=None, 1=Low, 2-3=Medium, 4-5=High, 6+=Critical        |
| `fraud_score` | integer | Fraud risk score from 0 to 100                          |
| `asn_org`     | string  | ASN organization name                                   |

---

### Validation Rules

- Maximum 50 IPs per request
- At least 1 IP required
- Invalid IPs do not fail the request
- Response preserves input order
- Each IP is evaluated independently
