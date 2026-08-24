# SWIFT Code API

## Status

✅ **Available**

## Overview

Validate and look up bank information by SWIFT/BIC code. Returns the bank name,
city, country, and the parsed components of the code.

SWIFT/BIC codes follow the ISO 9362 standard and are either 8 characters
(primary office) or 11 characters (specific branch). 8-character lookups
transparently resolve to the primary office record.

## Endpoints

### Get SWIFT Code

**Endpoint:** `GET /v1/finance/swift/{code}`

Look up bank metadata for a SWIFT/BIC code.

#### Path Parameters

| Parameter | Type   | Required | Description                                                              |
| --------- | ------ | -------- | ------------------------------------------------------------------------ |
| `code`    | string | Yes      | SWIFT/BIC code — 8 characters (primary office) or 11 characters (branch) |

#### Response Fields

| Field           | Type    | Description                                                |
| --------------- | ------- | ---------------------------------------------------------- |
| `swift_code`    | string  | Full 11-character BIC (8-char input expanded with `XXX`)   |
| `bank_code`     | string  | Institution code (characters 1–4)                          |
| `country_code`  | string  | ISO 3166-1 alpha-2 country code (characters 5–6)           |
| `location_code` | string  | Location code (characters 7–8)                             |
| `branch_code`   | string  | Branch code (characters 9–11); `XXX` = primary office      |
| `bank_name`     | string  | Full name of the bank or institution                       |
| `city`          | string  | City of the branch                                         |
| `country_name`  | string  | Country name                                               |
| `is_primary`    | boolean | `true` if this is the primary office (`branch_code = XXX`) |

#### Example Request

```bash
# Primary office (8-char)
curl https://requiems.xyz/v1/finance/swift/DEUTDEDB \
  -H "requiems-api-key: YOUR_API_KEY"

# Specific branch (11-char)
curl https://requiems.xyz/v1/finance/swift/DEUTDEDB001 \
  -H "requiems-api-key: YOUR_API_KEY"
```

#### Example Response

```json
{
  "data": {
    "swift_code": "DEUTDEDBXXX",
    "bank_code": "DEUT",
    "country_code": "DE",
    "location_code": "DB",
    "branch_code": "XXX",
    "bank_name": "Deutsche Bank",
    "city": "Frankfurt am Main",
    "country_name": "Germany",
    "is_primary": true
  },
  "metadata": {
    "timestamp": "2025-01-15T10:30:00Z"
  }
}
```

#### Error Responses

| Status | Code             | Description                                            |
| ------ | ---------------- | ------------------------------------------------------ |
| 400    | `bad_request`    | Invalid SWIFT code format (wrong length or characters) |
| 404    | `not_found`      | SWIFT code not found in the database                   |
| 500    | `internal_error` | Internal server error                                  |

#### Validation Rules

- Must be exactly 8 or 11 characters
- Characters 1–4 (bank code): letters only
- Characters 5–6 (country code): letters only
- Characters 7–8 (location code): alphanumeric
- Characters 9–11 (branch code, 11-char only): alphanumeric

### List SWIFT Codes

**Endpoint:** `GET /v1/finance/swift`

List SWIFT records with optional filters and pagination.

#### Query Parameters

| Parameter      | Type    | Required | Description                                          |
| -------------- | ------- | -------- | ---------------------------------------------------- |
| `country_code` | string  | No       | 2-letter country code filter (e.g. `DE`)             |
| `bank_code`    | string  | No       | 4-letter bank code filter (e.g. `DEUT`)              |
| `q`            | string  | No       | Search term for `swift_code`, `bank_name`, or `city` |
| `limit`        | integer | No       | Max rows to return (default `50`, max `200`)         |
| `offset`       | integer | No       | Rows to skip (default `0`)                           |

#### Example Request

```bash
curl "https://requiems.xyz/v1/finance/swift?country_code=DE&limit=3" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### List SWIFT Codes By Country

**Endpoint:** `GET /v1/finance/swift/country/{country_code}`

List SWIFT records for a country with optional search/pagination.

#### Path Parameters

| Parameter      | Type   | Required | Description           |
| -------------- | ------ | -------- | --------------------- |
| `country_code` | string | Yes      | 2-letter country code |

#### Example Request

```bash
curl "https://requiems.xyz/v1/finance/swift/country/US?q=chase&limit=2" \
  -H "requiems-api-key: YOUR_API_KEY"
```
