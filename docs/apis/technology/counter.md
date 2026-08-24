# Counter API

## Status

✅ **Live** - Production-ready

## Overview

A high-performance, namespace-isolated hit counter. Increments are atomic and
non-blocking — writes go to Redis instantly and are durably flushed to
PostgreSQL in the background every 60 seconds.

## Base URL

All endpoints are mounted under `/v1/technology`

## Endpoints

### 1. Increment Counter

Atomically increment a counter by 1 and return the updated value.

**Endpoint:** `POST /v1/technology/counter/{namespace}`

**Path Parameters:**

- `namespace` — Counter name. Must be 1–64 characters: alphanumeric, hyphens
  (`-`), or underscores (`_`).

**Response:** `200 OK`

```json
{
  "namespace": "page-views",
  "value": 42
}
```

**Example:**

```bash
curl -X POST https://requiems.xyz/v1/technology/counter/page-views \
  -H "requiems-api-key: YOUR_API_KEY"
```

---

### 2. Get Counter Value

Retrieve the current value of a counter. Reads from Redis; falls back to
PostgreSQL on a cache miss and re-hydrates the cache automatically.

**Endpoint:** `GET /v1/technology/counter/{namespace}`

**Path Parameters:**

- `namespace` — Counter name (same validation rules as above).

**Response:** `200 OK`

```json
{
  "namespace": "page-views",
  "value": 42
}
```

**Example:**

```bash
curl https://requiems.xyz/v1/technology/counter/page-views \
  -H "requiems-api-key: YOUR_API_KEY"
```

---

## Namespace Rules

| Rule               | Detail                             |
| ------------------ | ---------------------------------- |
| Min length         | 1 character                        |
| Max length         | 64 characters                      |
| Allowed characters | `a-z`, `A-Z`, `0-9`, `-`, `_`      |
| Examples           | `hits`, `page-views`, `my_counter` |

Invalid namespaces return `400 Bad Request`.

---

## Error Responses

```json
{
  "error": "error message description"
}
```

| Code              | Reason                   |
| ----------------- | ------------------------ |
| `400 Bad Request` | Invalid namespace format |

---

## Architecture Notes

- **Redis** (`INCR counter:{namespace}`) is the primary store — increments are
  O(1) and never block on a database write.
- **PostgreSQL** is the source of truth. A background sync worker scans all
  `counter:*` keys every 60 seconds and upserts the totals.
- On a GET cache miss, the PostgreSQL value is fetched and written back into
  Redis automatically.

---

## Use Cases

### Page View Tracking

```bash
# Increment on every page load
curl -X POST https://requiems.xyz/v1/technology/counter/homepage-views \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Feature Flag Hit Counting

```javascript
// Increment whenever a feature is used
await fetch(
  "https://requiems.xyz/v1/technology/counter/dark-mode-enabled",
  {
    method: "POST",
    headers: { "requiems-api-key": "YOUR_API_KEY" },
  },
);

// Read the current count
const res = await fetch(
  "https://requiems.xyz/v1/technology/counter/dark-mode-enabled",
  { headers: { "requiems-api-key": "YOUR_API_KEY" } },
);
const { value } = await res.json();
console.log(`Dark mode has been enabled ${value} times`);
```

### Event Counting

```python
import httpx

def track_event(event_name: str):
    httpx.post(
        f"https://requiems.xyz/v1/technology/counter/{event_name}",
        headers={"requiems-api-key": "YOUR_API_KEY"},
    )

track_event("user-signup")
track_event("checkout-completed")
```

---
