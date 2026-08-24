# World Time API

## Status

✅ **Live**

## Overview

Get the current time for any location in the world by specifying an IANA
timezone identifier. Returns the timezone name, UTC offset, current time, and
DST status.

## Live Endpoints

### Get Current Time by Timezone

**Endpoint:** `GET /v1/places/time/{timezone}`

Returns the current time for the given IANA timezone identifier.

#### Path Parameters

| Parameter  | Type   | Required | Description                                                      |
| ---------- | ------ | -------- | ---------------------------------------------------------------- |
| `timezone` | string | Yes      | IANA timezone identifier (e.g. `America/New_York`, `Asia/Tokyo`) |

#### Example Request

```bash
curl "https://requiems.xyz/v1/places/time/America/New_York" \
  -H "requiems-api-key: YOUR_API_KEY"
```

#### Example Response

```json
{
  "timezone": "America/New_York",
  "offset": "-05:00",
  "current_time": "2024-12-15T14:30:00Z",
  "is_dst": false
}
```

#### Response Fields

| Field          | Type    | Description                                            |
| -------------- | ------- | ------------------------------------------------------ |
| `timezone`     | string  | IANA timezone identifier                               |
| `offset`       | string  | UTC offset in `+HH:MM` or `-HH:MM` format              |
| `current_time` | string  | Current UTC time in RFC 3339 format                    |
| `is_dst`       | boolean | Whether the timezone is observing daylight saving time |
