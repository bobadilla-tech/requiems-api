# Global Data System — Engine Spec

The Global Data System powers international products with accurate, real-time
location and calendar data. It answers: where is this location, what time is it
there, and what does the business calendar look like?

**Distinct from other systems:**

| Question                                                     | System                |
| ------------------------------------------------------------ | --------------------- |
| Is this data valid?                                          | Data Integrity        |
| Is this user fraudulent?                                     | Identity & Risk       |
| Is this transaction safe?                                    | Payments Intelligence |
| Where is this, what time is it, what are the local holidays? | **Global Data**       |

IP geolocation appears in both Identity & Risk and Global Data. The distinction:
Identity & Risk uses IP to extract FRAUD signals (is this user hiding behind a
VPN?). Global Data uses IP to extract LOCATION CONTEXT (what timezone is this
user in?). Same underlying service (`services/networking/ip/info`), different
output shape, different call site.

---

## Conventions

- Base path: `/v1/`
- Every response wrapped in
  `{"data": {...}, "metadata": {"timestamp": "...", "trace_id": "..."}}`
- `flags` always `string[]` of machine-readable codes
- Timestamps in ISO 8601 / RFC 3339
- No risk scores — this system is neutral; it returns facts, not decisions

---

## Endpoints

### Engine: `POST /v1/location/resolve`

Primary composed endpoint. Resolves any location input into a full context
object: normalized address, coordinates, city, country, timezone, current local
time, holiday status, and working day count for the current month.

Replaces 4–5 individual API calls (geocode + timezone + world time + holidays +
working days) and eliminates the sequential dependency chain (geocode must
complete before timezone lookup can begin — this endpoint handles that
internally).

**Internal APIs composed:**

1. Geocoding (`services/places/geocode`) — resolves `address` to coordinates
2. Then parallel fan-out:
   - Timezone (`services/places/timezone`) — from coordinates
   - Holidays (`services/places/holidays`) — from country + current year
   - Working Days (`services/places/working-days`) — current month for country

**Request** — provide `address` OR `coordinates` (one required)

```json
{
  "address": "Alexanderplatz 1, Berlin",
  "country_code": "DE"
}
```

```json
{
  "coordinates": { "lat": 52.52, "lng": 13.405 }
}
```

`country_code` is optional when `address` is provided (geocoder infers it).
Required for holidays/working-days fallback if geocoder cannot determine
country.

**Response**

```json
{
  "address": "Alexanderplatz 1, 10178 Berlin, Germany",
  "city": "Berlin",
  "country": "Germany",
  "country_code": "DE",
  "coordinates": { "lat": 52.5219, "lng": 13.4132 },
  "timezone": "Europe/Berlin",
  "utc_offset": "+02:00",
  "current_time": "2026-05-25T19:14:00+02:00",
  "is_holiday_today": false,
  "working_days_this_month": 21,
  "next_holiday": {
    "date": "2026-06-04",
    "name": "Whit Thursday",
    "type": "national"
  }
}
```

`next_holiday` is `null` if no upcoming holiday found within the next 90 days.

**Degradation**: if geocoder fails, return 422 (cannot proceed without
coordinates). If timezone/holidays/working-days fail, omit those fields and add
flags: `timezone_unavailable`, `calendar_unavailable`.

**Flags:** `timezone_unavailable`, `calendar_unavailable`, `country_inferred`
(when country was inferred from geocode result, not provided explicitly)

---

### `GET /v1/timezone/from-ip/{ip}`

Detect a user's timezone from their IP address. Two-step composition: IP
geolocation → city/country → timezone lookup. Returns the current local time,
offset, and DST status.

**Distinct from Identity & Risk IP use**: here the IP is used for LOCATION
ENRICHMENT only. No VPN check, no fraud score, no `is_safe` decision.

**Internal APIs composed (sequential — timezone lookup needs country from
geo):**

1. IP Geolocation (`services/networking/ip/info`) — extract country, city
2. Timezone (`services/places/timezone`) — from city/country

**Request**

```
GET /v1/timezone/from-ip/203.0.113.42
```

If `{ip}` is omitted or set to `me`, falls back to caller IP (same pattern as
`services/networking/ip/info`).

**Response**

```json
{
  "ip": "203.0.113.42",
  "city": "Berlin",
  "country_code": "DE",
  "timezone": "Europe/Berlin",
  "utc_offset": "+02:00",
  "dst_active": true,
  "current_time": "2026-05-25T19:14:00+02:00"
}
```

**Error cases:**

- Private/reserved IP → 422 with `private_ip` error code
- IP not found in geolocation database → 404 with `ip_not_found`
- Timezone resolution fails (city found but no timezone match) → return response
  with `timezone: null`, `current_time: null`, flag `timezone_unavailable`

---

### `GET /v1/business-calendar/{country}`

Return the business calendar for a country and time period. Combines holiday
data with working day count so that scheduling, invoicing, and SLA systems can
operate without building their own calendar logic per country.

**Internal APIs composed (parallel):**

- Holidays (`services/places/holidays`)
- Working Days (`services/places/working-days`)

**Request**

```
GET /v1/business-calendar/DE?year=2026&month=6
```

`country` — ISO 3166-1 alpha-2 (path param, required)\
`year` — optional, defaults to current year\
`month` — optional, defaults to current month. When provided, `working_days`
scoped to that month. When omitted, returns full-year holidays + year working
day total.

**Response (month scope)**

```json
{
  "country_code": "DE",
  "year": 2026,
  "month": 6,
  "working_days": 21,
  "total_days": 30,
  "weekend_days": 8,
  "holidays": [
    {
      "date": "2026-06-04",
      "name": "Whit Thursday",
      "type": "national"
    }
  ],
  "holiday_count": 1,
  "next_holiday": {
    "date": "2026-06-04",
    "name": "Whit Thursday",
    "type": "national"
  }
}
```

**Response (year scope — month param omitted)**

```json
{
  "country_code": "DE",
  "year": 2026,
  "working_days": 252,
  "holidays": [ ... ],
  "holiday_count": 12,
  "next_holiday": { ... }
}
```

**Error cases:**

- Unknown country code → 422 with `unsupported_country`
- Year out of supported range → 422 with `year_out_of_range`

---

## Go package structure

```
apps/api/services/systems/
└── global_data/
    ├── router.go                          # mounts all sub-routers
    ├── location_resolve/
    │   ├── service.go                     # geocode → parallel timezone + holidays + working-days
    │   ├── transport_http.go              # POST /v1/location/resolve
    │   └── transport_http_test.go
    ├── timezone_ip/
    │   ├── service.go                     # sequential: ip/info → timezone
    │   ├── transport_http.go              # GET /v1/timezone/from-ip/{ip}
    │   └── transport_http_test.go
    └── business_calendar/
        ├── service.go                     # parallel: holidays + working-days
        ├── transport_http.go              # GET /v1/business-calendar/{country}
        └── transport_http_test.go
```

Note: `location_resolve/service.go` is the only endpoint in the whole system
where composition is **sequential then parallel** (geocode must complete first
to get coordinates, then timezone/holidays/working-days fan out). All others are
fully parallel.

---

## Testing requirements

### `/v1/location/resolve`

| Case                                    | Expected                                                                     |
| --------------------------------------- | ---------------------------------------------------------------------------- |
| Valid address                           | 200, coordinates populated, timezone non-null, `working_days_this_month > 0` |
| Valid coordinates (no address)          | 200, `address` field shows reverse-geocoded result                           |
| Both address and coordinates provided   | Use coordinates, ignore address (coordinates are authoritative)              |
| Address geocode fails                   | 422                                                                          |
| Timezone resolution fails after geocode | 200, `timezone: null`, `timezone_unavailable` in flags                       |
| Holiday lookup fails                    | 200, calendar fields null, `calendar_unavailable` in flags                   |
| Neither address nor coordinates         | 422                                                                          |
| All composed services stubbed           | ✓                                                                            |

### `/v1/timezone/from-ip/{ip}`

| Case                               | Expected                                                |
| ---------------------------------- | ------------------------------------------------------- |
| Valid public IP                    | 200, `timezone` non-null, `current_time` valid ISO 8601 |
| Private IP (192.168.x.x)           | 422, `private_ip` error code                            |
| Unknown IP                         | 404, `ip_not_found`                                     |
| IP found, timezone fails           | 200, `timezone: null`, `timezone_unavailable` flag      |
| No IP in path (caller IP fallback) | 200, resolved from X-Forwarded-For / RemoteAddr         |
| All composed services stubbed      | ✓                                                       |

### `/v1/business-calendar/{country}`

| Case                                    | Expected                                                  |
| --------------------------------------- | --------------------------------------------------------- |
| Valid country + month                   | 200, `working_days > 0`, `holidays` array present         |
| Valid country + year scope (no month)   | 200, full year `holidays`, `working_days` is annual total |
| Country with no holidays this month     | 200, `holidays: []`, `holiday_count: 0`                   |
| Unknown country code                    | 422, `unsupported_country`                                |
| Missing country param                   | 422                                                       |
| `next_holiday` when no upcoming holiday | 200, `next_holiday: null`                                 |
| All composed services stubbed           | ✓                                                         |

---

## Open questions for engineer

1. **Sequential composition in location/resolve** — geocode runs first, then
   timezone/holidays/working-days fan out in parallel. Use a two-phase approach:
   phase 1 = geocode (with timeout), phase 2 = parallel fan-out using resolved
   country + coordinates. Clarify timeout budget: geocode gets 3s, fan-out
   services get 2s each.
2. **`next_holiday` lookback window** — how far ahead to scan for the next
   holiday when returning `next_holiday`. Suggest 90 days as default. If no
   holiday found in that window, return `null`.
3. **Working days and month scope** — `services/places/working-days` may accept
   date ranges differently than month/year integers. Verify the service's input
   contract before building the business calendar wrapper.

## Ticket breakdown

| # | Endpoint                              | Composes                                     | Priority |
| - | ------------------------------------- | -------------------------------------------- | -------- |
| 1 | `GET /v1/business-calendar/{country}` | holidays + working-days                      | medium   |
| 2 | `GET /v1/timezone/from-ip/{ip}`       | ip/info + timezone                           | medium   |
| 3 | `POST /v1/location/resolve`           | geocode + timezone + holidays + working-days | high     |

Implement in order: business-calendar first (pure parallel, no IP), then
timezone/from-ip (sequential, adds IP), then location/resolve (most complex —
sequential + parallel + degradation handling).
