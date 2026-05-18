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
| `lat`     | float  | *        | Latitude (-90 to 90). Required when using coordinates.    |
| `lon`     | float  | *        | Longitude (-180 to 180). Required when using coordinates. |
| `city`    | string | *        | City name. Required when not using coordinates.           |

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


## Batch Timezone 

**Endpoint:** `POST /v1/places/timezone/batch`

Get timezone information for multiple cities in a single request.

The response preserves the same order as the input list. Invalid or unknown
cities return `"info": null`.

### Request Body

| Field    | Type     | Required | Description                        |
| -------- | -------- | -------- | ---------------------------------- |
| `cities` | string[] | Yes      | List of city names (max 50 items). |

### Example Request

```http
POST /v1/places/timezone/batch
Content-Type: application/json
````

```json
{
  "cities": [
    "Tokyo",
    "Lima",
    "New York"
  ]
}
```

### Example Response

```json
{
  "data": {
    "results": [
      {
        "city": "Tokyo",
        "info": {
          "timezone": "Asia/Tokyo",
          "offset": "+09:00",
          "current_time": "2024-12-15T14:30:00Z",
          "is_dst": false
        }
      },
      {
        "city": "Lima",
        "info": {
          "timezone": "America/Lima",
          "offset": "-05:00",
          "current_time": "2024-12-15T14:30:00Z",
          "is_dst": false
        }
      },
      {
        "city": "Atlantis",
        "info": null
      }
    ]
  },
  "metadata": {
    "timestamp": "2024-12-15T14:30:00Z"
  }
}
```

### Response Fields

| Field            | Type         | Description                                      |
| ---------------- | ------------ | ------------------------------------------------ |
| `results`        | object[]     | List of timezone lookup results                  |
| `results[].city` | string       | Original city name from the request              |
| `results[].info` | object | nil | Timezone information or `null` if city not found |

### Info Object Fields

| Field          | Type    | Description                                     |
| -------------- | ------- | ----------------------------------------------- |
| `timezone`     | string  | IANA timezone identifier                        |
| `offset`       | string  | UTC offset in `+HH:MM` / `-HH:MM` format        |
| `current_time` | string  | Current UTC time in RFC 3339 format             |
| `is_dst`       | boolean | Whether the location is currently observing DST |

### Validation Rules

* `cities` must contain at least 1 city
* Maximum 50 cities per request
* Invalid city names do not fail the request
* Results preserve request ordering

