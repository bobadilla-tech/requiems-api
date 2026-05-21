# Batch APIs Expansion

Nine existing single-item endpoints have no batch equivalent.

All nine services exist and are wired into their domain routers. No new domains,
new top-level routes, or new dependencies are required. Every change is
additive: new types, a new `BatchXxx` service method, and a new
`httpx.HandleBatch` registration in each service's existing `transport_http.go`.

---

## Standards Applied

All batch endpoints in this expansion follow `docs/core/batch-apis.md` without
exception:

| Topic                            | Standard                                                         |
| -------------------------------- | ---------------------------------------------------------------- |
| HTTP method                      | `POST`                                                           |
| Path                             | `/resource/batch` under the same prefix as the single-item route |
| Go handler                       | `httpx.HandleBatch` — sets `X-Usage-Count` automatically         |
| Max items (in-memory / local DB) | 50                                                               |
| Max items (outbound HTTP)        | 20 — Nominatim usage policy                                      |
| Error model                      | Partial success: per-item `error` field; top-level 200           |
| Service layer                    | Real batch method — no loop over the single-item method for I/O  |
| Result ordering                  | Same order as input array                                        |

---

## Endpoint Inventory

| API                    | New endpoint                            | Domain       | Batch strategy                          |
| ---------------------- | --------------------------------------- | ------------ | --------------------------------------- |
| Timezone               | `POST /v1/places/timezone/batch`        | `places`     | Sequential (in-memory `tzf`)            |
| QR Code (base64)       | `POST /v1/technology/qr/base64/batch`   | `technology` | Sequential (in-memory)                  |
| Profanity Filter       | `POST /v1/validation/profanity/batch`   | `validation` | Sequential (in-memory `go-away`)        |
| Postal Code            | `POST /v1/places/postal/batch`          | `places`     | Sequential (in-memory map)              |
| Password Generator     | `POST /v1/technology/password/batch`    | `technology` | Sequential (`crypto/rand`)              |
| Number Base Conversion | `POST /v1/technology/base/batch`        | `technology` | Sequential (pure math)                  |
| IP Geolocation         | `POST /v1/networking/ip/info/batch`     | `networking` | Sequential (local MaxMind DB)           |
| ASN Lookup             | `POST /v1/networking/ip/asn/batch`      | `networking` | Sequential (local MaxMind DB)           |
| Geocoding — forward    | `POST /v1/places/geocode/batch`         | `places`     | Goroutines + semaphore (Nominatim HTTP) |
| Geocoding — reverse    | `POST /v1/places/reverse-geocode/batch` | `places`     | Goroutines + semaphore (Nominatim HTTP) |

> **Phone validation** (`POST /v1/validation/phone/batch`) is already live and
> is excluded from this work.
>
> **QR PNG** (`GET /v1/technology/qr`) gets no batch variant — it returns
> `image/png` binary, which is incompatible with the JSON batch envelope. Only
> the base64 JSON variant (`/qr/base64`) is batched.

---

## Per-Endpoint Specs

### 1. Profanity Filter — `POST /v1/validation/profanity/batch`

**Files:** `apps/api/services/validation/profanity/`

Service uses `go-away` (in-memory). Pure computation — sequential loop is
correct.

**Request / Response:**

```json
// POST /v1/validation/profanity/batch
{
  "texts": ["hello world", "offensive text here"]
}

// 200 OK
{
  "data": {
    "results": [
      { "text": "hello world",        "result": { "has_profanity": false, "flagged_words": [] } },
      { "text": "offensive text here","result": { "has_profanity": true,  "flagged_words": ["offensive"] } }
    ],
    "total": 2
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchRequest struct {
    Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

type BatchProfanityItem struct {
    Text   string  `json:"text"`
    Result *Result `json:"result,omitempty"`
    Error  string  `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) CheckBatch(ctx context.Context, texts []string) []BatchProfanityItem
```

---

### 2. Number Base Conversion — `POST /v1/technology/base/batch`

**Files:** `apps/api/services/technology/numbase/`

Pure integer arithmetic. Sequential loop. The existing single-item handler uses
inline `BindQuery` — the batch handler uses `httpx.HandleBatch` with a JSON
body.

**Request / Response:**

```json
// POST /v1/technology/base/batch
{
  "items": [
    { "from": 10, "to": 2,  "value": "255" },
    { "from": 16, "to": 10, "value": "FF"  },
    { "from": 10, "to": 16, "value": "bad" }
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "from": 10, "to": 2,  "input": "255", "result": "11111111" },
      { "from": 16, "to": 10, "input": "FF",  "result": "255"      },
      { "from": 10, "to": 16, "input": "bad", "error": "invalid value for base 10" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchConvertRequest struct {
    Items []ConvertQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

type ConvertQuery struct {
    From  int    `json:"from"  validate:"required,oneof=2 8 10 16"`
    To    int    `json:"to"    validate:"required,oneof=2 8 10 16"`
    Value string `json:"value" validate:"required"`
}

type BatchConvertItem struct {
    From   int    `json:"from"`
    To     int    `json:"to"`
    Input  string `json:"input"`
    Result string `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) ConvertBatch(items []ConvertQuery) []BatchConvertItem
```

---

### 3. Password Generator — `POST /v1/technology/password/batch`

**Files:** `apps/api/services/technology/password/`

Each item may have independent config (length, charset flags). Pure
`crypto/rand` — no I/O. Default `Length = 16` when zero.

**Request / Response:**

```json
// POST /v1/technology/password/batch
{
  "items": [
    { "length": 16, "uppercase": true, "numbers": true, "symbols": false },
    { "length": 32, "uppercase": true, "numbers": true, "symbols": true  }
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "result": { "password": "Abc1Def2Ghi3Jkl4", "length": 16, "entropy": 95.2 } },
      { "result": { "password": "Xy!z9...", "length": 32, "entropy": 190.5 } }
    ],
    "total": 2
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchPasswordRequest struct {
    Items []PasswordQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

type PasswordQuery struct {
    Length    int  `json:"length"    validate:"min=8,max=128"`
    Uppercase bool `json:"uppercase"`
    Numbers   bool `json:"numbers"`
    Symbols   bool `json:"symbols"`
}

type BatchPasswordItem struct {
    Result *Password `json:"result,omitempty"`
    Error  string    `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) GenerateBatch(items []PasswordQuery) []BatchPasswordItem
// Apply Length default of 16 when item.Length == 0
```

---

### 4. Postal Code — `POST /v1/places/postal/batch`

**Files:** `apps/api/services/places/postal/`

In-memory map (loaded from GeoNames TSV at startup). O(1) per lookup. Country
defaults to `"US"` when empty. "Not found" is a valid outcome — returned via
`found: false`, not an error.

**Request / Response:**

```json
// POST /v1/places/postal/batch
{
  "items": [
    { "code": "10001", "country": "US" },
    { "code": "SW1A1AA", "country": "GB" },
    { "code": "00000", "country": "US" }
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "code": "10001",   "country": "US", "found": true,  "result": { "postal_code": "10001", "city": "New York", "state": "New York", "country": "US", "lat": 40.748, "lon": -73.996 } },
      { "code": "SW1A1AA", "country": "GB", "found": true,  "result": { ... } },
      { "code": "00000",   "country": "US", "found": false }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchPostalRequest struct {
    Items []PostalQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

type PostalQuery struct {
    Code    string `json:"code"    validate:"required"`
    Country string `json:"country" validate:"omitempty,len=2"`
}

type BatchPostalItem struct {
    Code    string      `json:"code"`
    Country string      `json:"country"`
    Found   bool        `json:"found"`
    Result  *PostalCode `json:"result,omitempty"`
    Error   string      `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) LookupBatch(items []PostalQuery) []BatchPostalItem
// Default Country to "US" when empty; map to uppercase before lookup
```

---

### 5. Timezone — `POST /v1/places/timezone/batch`

**Files:** `apps/api/services/places/timezone/`

In-memory `tzf` lookup. Three query modes per item, evaluated in priority order:
IANA timezone name → city name → lat+lon. Items with none of these set return an
in-band error.

**Request / Response:**

```json
// POST /v1/places/timezone/batch
{
  "items": [
    { "timezone": "America/New_York" },
    { "city": "Tokyo" },
    { "lat": 51.5074, "lon": -0.1278 },
    {}
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "info": { "timezone": "America/New_York", "utc_offset": "-05:00", ... } },
      { "info": { "timezone": "Asia/Tokyo", ... } },
      { "info": { "timezone": "Europe/London", ... } },
      { "error": "provide timezone name, city, or lat+lon" }
    ],
    "total": 4
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchTimezoneRequest struct {
    Items []TimezoneQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

type TimezoneQuery struct {
    Timezone string  `json:"timezone"`
    City     string  `json:"city"`
    Lat      float64 `json:"lat" validate:"min=-90,max=90"`
    Lon      float64 `json:"lon" validate:"min=-180,max=180"`
}

type BatchTimezoneItem struct {
    Info  *Info  `json:"info,omitempty"`
    Error string `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) BatchLookup(items []TimezoneQuery) []BatchTimezoneItem
// Priority: Timezone > City > lat+lon (non-zero coords)
// Error in-band when no valid query mode detected
```

---

### 6. QR Code (base64) — `POST /v1/technology/qr/base64/batch`

**Files:** `apps/api/services/technology/qr/`

In-memory QR generation. Only the base64 JSON variant is batched — the raw PNG
endpoint (`GET /qr`) is incompatible with the JSON batch envelope. Default
`Size = 256` when zero.

**Request / Response:**

```json
// POST /v1/technology/qr/base64/batch
{
  "items": [
    { "data": "https://requiems.xyz", "size": 256, "recovery": "medium" },
    { "data": "hello world" },
    { "data": "" }
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "data": "https://requiems.xyz", "image": "iVBORw0KGgo...", "width": 256, "height": 256 },
      { "data": "hello world",          "image": "iVBORw0KGgo...", "width": 256, "height": 256 },
      { "data": "",                     "error": "data is required" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchQRRequest struct {
    Items []QRQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

type QRQuery struct {
    Data     string `json:"data"     validate:"required"`
    Size     int    `json:"size"     validate:"min=50,max=1000"`
    Recovery string `json:"recovery" validate:"omitempty,oneof=low medium high highest"`
}

type BatchBase64Item struct {
    Data   string `json:"data"`
    Image  string `json:"image,omitempty"`
    Width  int    `json:"width,omitempty"`
    Height int    `json:"height,omitempty"`
    Error  string `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) GenerateBatch(items []QRQuery) []BatchBase64Item
// Apply Size default of 256 when item.Size == 0
```

---

### 7. IP Geolocation — `POST /v1/networking/ip/info/batch`

**Files:** `apps/api/services/networking/ip/info/`

Uses `ipi.Client` from `bobadilla-tech/go-ip-intelligence/v2`, which wraps local
MaxMind DB files. No outbound network I/O — sequential loop is correct.

**Request / Response:**

```json
// POST /v1/networking/ip/info/batch
{
  "ips": ["8.8.8.8", "1.1.1.1", "not-an-ip"]
}

// 200 OK
{
  "data": {
    "results": [
      { "ip": "8.8.8.8", "result": { "ip": "8.8.8.8", "country": "United States", "country_code": "US", "city": "Mountain View", "isp": "Google LLC", "is_vpn": false } },
      { "ip": "1.1.1.1", "result": { ... } },
      { "ip": "not-an-ip", "error": "invalid IP address" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchInfoRequest struct {
    IPs []string `json:"ips" validate:"required,min=1,max=50,dive,ip"`
}

type BatchIPInfoItem struct {
    IP     string          `json:"ip"`
    Result *LookupResponse `json:"result,omitempty"`
    Error  string          `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) CheckInfoBatch(ctx context.Context, ips []string) []BatchIPInfoItem
```

---

### 8. ASN Lookup — `POST /v1/networking/ip/asn/batch`

**Files:** `apps/api/services/networking/ip/asn/`

Same `ipi.Client` local MaxMind DB as IP Geolocation. Private/reserved IPs
return `BatchASNItem{IP: ip}` with no error, matching the existing single-item
behavior.

**Request / Response:**

```json
// POST /v1/networking/ip/asn/batch
{
  "ips": ["8.8.8.8", "1.1.1.1"]
}

// 200 OK
{
  "data": {
    "results": [
      { "ip": "8.8.8.8", "result": { "ip": "8.8.8.8", "asn": "AS15169", "org": "Google LLC", "isp": "Google LLC", "domain": "google.com", "route": "8.8.8.0/24", "type": "hosting" } },
      { "ip": "1.1.1.1", "result": { ... } }
    ],
    "total": 2
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchASNRequest struct {
    IPs []string `json:"ips" validate:"required,min=1,max=50,dive,ip"`
}

type BatchASNItem struct {
    IP     string                `json:"ip"`
    Result *IPAddressASNResponse `json:"result,omitempty"`
    Error  string                `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) CheckASNBatch(ctx context.Context, ips []string) []BatchASNItem
```

---

### 9. Geocoding — `POST /v1/places/geocode/batch` + `POST /v1/places/reverse-geocode/batch`

**Files:** `apps/api/services/places/geocode/`

**⚠ Outbound HTTP.** The service calls the Nominatim API via `httpClient.Do()`.
Redis caching is applied per-item inside `Geocode()`/`ReverseGeocode()` — cache
hits avoid network I/O. Worst case is N outbound HTTP calls; use goroutines +
semaphore as required by the RFC.

Max items is **20** (Nominatim usage policy).

Goroutine pattern: follow `networking/whois/service.go` or
`validation/email/service.go` (canonical references cited in
`docs/core/batch-apis.md`).

**Request / Response — forward:**

```json
// POST /v1/places/geocode/batch
{
  "addresses": ["1600 Amphitheatre Parkway, Mountain View, CA", "Eiffel Tower, Paris"]
}

// 200 OK
{
  "data": {
    "results": [
      { "address": "1600 Amphitheatre Parkway, Mountain View, CA", "result": { "address": "...", "city": "Mountain View", "country": "US", "lat": 37.422, "lon": -122.084 } },
      { "address": "Eiffel Tower, Paris", "result": { ... } }
    ],
    "total": 2
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**Request / Response — reverse:**

```json
// POST /v1/places/reverse-geocode/batch
{
  "items": [
    { "lat": 37.4224764, "lon": -122.0842499 },
    { "lat": 48.8584,    "lon": 2.2945 }
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "lat": 37.4224764, "lon": -122.0842499, "result": { "lat": 37.4224764, "lon": -122.0842499, "address": "...", "city": "Mountain View", "country": "US" } },
      { "lat": 48.8584,    "lon": 2.2945,        "result": { ... } }
    ],
    "total": 2
  },
  "metadata": { "timestamp": "2026-05-21T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
// Forward
type GeocodeBatchRequest struct {
    Addresses []string `json:"addresses" validate:"required,min=1,max=20,dive,required"`
}

type BatchGeocodeItem struct {
    Address string           `json:"address"`
    Result  *GeocodeResponse `json:"result,omitempty"`
    Error   string           `json:"error,omitempty"`
}

// Reverse
type ReverseGeocodeBatchRequest struct {
    Items []ReverseQuery `json:"items" validate:"required,min=1,max=20,dive"`
}

type ReverseQuery struct {
    Lat float64 `json:"lat" validate:"required,min=-90,max=90"`
    Lon float64 `json:"lon" validate:"required,min=-180,max=180"`
}

type BatchReverseGeocodeItem struct {
    Lat    float64                 `json:"lat"`
    Lon    float64                 `json:"lon"`
    Result *ReverseGeocodeResponse `json:"result,omitempty"`
    Error  string                  `json:"error,omitempty"`
}
```

**New service methods:**

```go
const maxGeoWorkers = 10

func (s *Service) GeocodeBatch(ctx context.Context, addresses []string) []BatchGeocodeItem
// Goroutines + semaphore; per-item context.WithTimeout(ctx, 3s)
// In-band errors; do NOT use errgroup

func (s *Service) ReverseGeocodeBatch(ctx context.Context, items []ReverseQuery) []BatchReverseGeocodeItem
// Same concurrency pattern as GeocodeBatch
```

---

## Implementation Order

| Step | Service                     | Reason                                                      |
| ---- | --------------------------- | ----------------------------------------------------------- |
| 1    | Profanity Filter            | Simplest: single string slice, no result variants           |
| 2    | Number Base Conversion      | Pure math, per-item error is common, good test case         |
| 3    | Password Generator          | Stateless generation; validates default-value handling      |
| 4    | Postal Code                 | In-memory map; introduces `found` bool pattern              |
| 5    | Timezone                    | Multi-mode query resolution                                 |
| 6    | QR Code base64              | Validates PNG-exclusion documentation                       |
| 7    | IP Geolocation + ASN Lookup | Implement together; same `ipi.Client` pattern               |
| 8    | Geocoding forward + reverse | Most complex; goroutines + semaphore + Nominatim rate limit |

---

## Documentation Requirements (per endpoint)

- [ ] Batch types + service method added to existing `service.go` and
      `transport_http.go`
- [ ] Route registered in existing `RegisterRoutes()` — no `router.go` changes
      needed
- [ ] `transport_http_test.go` covers: happy path (N items), empty array → 422,
      oversize → 422, mixed valid/invalid → 200 partial
- [ ] `service_test.go` covers `BatchXxx` method when non-trivial logic exists
- [ ] `apps/dashboard/config/api_docs/{name}.yml` — batch endpoint section
      added:
  - Correct `method: POST` and path
  - `request_example` and `response_example` with `data/metadata` wrapper
  - Parameters with `min=1,max=N` documented
  - Error table: `validation_failed` (422), `internal_error` (500)
  - Code examples in curl, Python, JavaScript, Ruby using `requiems-api-key`
    header
- [ ] `apps/dashboard/config/api_catalog.yml` — `endpoints_count` incremented
      for each affected API
- [ ] `docs/apis/{category}/{name}.md` — batch endpoint section added
- [ ] `apps/workers/shared/src/config.ts` — no changes; `HandleBatch` sets
      `X-Usage-Count` dynamically

---

## Verification

```bash
# Per-domain (fast feedback during development)
docker exec requiem-dev-api-1 go test ./services/validation/profanity/...
docker exec requiem-dev-api-1 go test ./services/technology/numbase/...
docker exec requiem-dev-api-1 go test ./services/technology/password/...
docker exec requiem-dev-api-1 go test ./services/technology/qr/...
docker exec requiem-dev-api-1 go test ./services/places/postal/...
docker exec requiem-dev-api-1 go test ./services/places/timezone/...
docker exec requiem-dev-api-1 go test ./services/places/geocode/...
docker exec requiem-dev-api-1 go test ./services/networking/ip/...

# Full suite with race detection (required before PR)
docker exec requiem-dev-api-1 go test -race -coverprofile=coverage.out ./...

# Smoke test (replace X-Backend-Secret as needed)
curl -X POST http://localhost:8080/v1/validation/profanity/batch \
  -H "X-Backend-Secret: local_secret" \
  -H "Content-Type: application/json" \
  -d '{"texts":["clean text","offensive word"]}'

# Billing verification — X-Usage-Count must equal item count
curl -v -X POST http://localhost:8080/v1/validation/profanity/batch \
  -H "X-Backend-Secret: local_secret" \
  -H "Content-Type: application/json" \
  -d '{"texts":["a","b","c"]}' 2>&1 | grep X-Usage-Count
# Expected: X-Usage-Count: 3

# YAML validation (run for each updated doc)
docker exec requiem-dev-dashboard-1 ruby -ryaml \
  -e "YAML.load_file('config/api_docs/profanity-filter.yml'); puts 'OK'"

# Dashboard playground
open http://localhost:3000/apis/profanity-filter
```
