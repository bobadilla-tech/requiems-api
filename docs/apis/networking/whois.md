# Whois API

## Status

✅ **Live**

## Overview

Get WHOIS registration details for any domain name. Returns registrar, name
servers, status flags, registration dates, and DNSSEC information.

Supports both single-domain lookups and batch WHOIS queries.

## Endpoint

### WHOIS Lookup

`GET /v1/networking/whois/{domain}`

Returns WHOIS information for a domain.

| Parameter | Type   | Required | Description                                 |
| --------- | ------ | -------- | ------------------------------------------- |
| `domain`  | string | Yes      | Domain name to look up (e.g. `example.com`) |

## Response

```json
{
  "data": {
    "domain": "example.com",
    "registrar": "RESERVED-Internet Assigned Numbers Authority",
    "name_servers": ["A.IANA-SERVERS.NET", "B.IANA-SERVERS.NET"],
    "status": [
      "clientDeleteProhibited",
      "clientTransferProhibited",
      "clientUpdateProhibited"
    ],
    "created_date": "1995-08-14T04:00:00Z",
    "updated_date": "2023-08-14T07:01:38Z",
    "expiry_date": "2024-08-13T04:00:00Z",
    "dnssec": true
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Error Codes

| Code             | Status | When                                            |
| ---------------- | ------ | ----------------------------------------------- |
| `bad_request`    | 400    | Domain name format is invalid                   |
| `not_found`      | 404    | Domain is not registered or no WHOIS data found |
| `internal_error` | 500    | Upstream WHOIS query failed                     |

## Batch WHOIS lookup

`POST /v1/networking/whois/batch`

Accepts up to **50 domains** per request. Results are returned in the same
order as the input array.

Domains that do not exist or have no WHOIS data are returned with:

```json
{
  "found": false
}
```

instead of failing the entire request.

Each domain in the request counts as **1 credit** (`X-Usage-Count` equals the
number of domains submitted).

---

## Request body

```json
{
  "domains": ["example.com", "google.com"]
}
```

---

## Batch response

```json
{
  "data": {
    "results": [
      {
        "domain": "example.com",
        "found": true,
        "data": {
          "domain": "example.com",
          "registrar": "RESERVED-Internet Assigned Numbers Authority",
          "name_servers": ["A.IANA-SERVERS.NET", "B.IANA-SERVERS.NET"],
          "status": ["clientDeleteProhibited"],
          "created_date": "1995-08-14T04:00:00Z",
          "updated_date": "2023-08-14T07:01:38Z",
          "expiry_date": "2024-08-13T04:00:00Z",
          "dnssec": true
        }
      },
      {
        "domain": "doesnotexist.com",
        "found": false,
        "error": "domain not found"
      }
    ],
    "total": 2
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

---

# Batch behavior

- Maximum of **50 domains** per request.
- Results preserve the original request order.
- Lookups run concurrently for improved performance.
- Individual domain failures do not fail the entire batch.
- `total` always equals the number of domains submitted.

---

# Error Codes

| Code                | Status | When                                                      |
| ------------------- | ------ | --------------------------------------------------------- |
| `bad_request`       | 400    | Invalid JSON body or malformed request                    |
| `validation_failed` | 422    | Empty batch, more than 50 domains, or invalid domain name |
| `not_found`         | 404    | Domain not found (single lookup endpoint only)            |
| `internal_error`    | 500    | WHOIS query failed unexpectedly                           |

---

# Validation Rules

Batch requests validate the following constraints:

| Field       | Rule                         |
| ----------- | ---------------------------- |
| `domains`   | Required                     |
| `domains`   | Minimum 1 item               |
| `domains`   | Maximum 50 items             |
| `domains[]` | Must be a valid RFC1123 host |
