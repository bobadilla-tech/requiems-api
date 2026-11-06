# Batch APIs Expansion — Wave 2

Twelve existing single-item endpoints have no batch equivalent. One API
(dictionary) is deferred pending single-item implementation.

All twelve services exist and are wired into their domain routers. No new
domains, new top-level routes, or new dependencies are required. Every change is
additive: new types, a new `BatchXxx` service method, and a new
`httpx.HandleBatch` registration in each service's existing `transport_http.go`.

---

## Standards Applied

All batch endpoints in this expansion follow `docs/core/batch-apis.md` without
exception:

| Topic                         | Standard                                                         |
| ----------------------------- | ---------------------------------------------------------------- |
| HTTP method                   | `POST`                                                           |
| Path                          | `/resource/batch` under the same prefix as the single-item route |
| Go handler                    | `httpx.HandleBatch` — sets `X-Usage-Count` automatically         |
| Max items (in-memory / local) | 50                                                               |
| Max items (outbound HTTP)     | 20                                                               |
| Max items (domain-info DNS)   | 10 — each domain fans out 5 DNS queries internally               |
| Error model                   | Partial success: per-item `error` field; top-level 200           |
| Service layer                 | Real batch method — no loop over single-item method for I/O      |
| Result ordering               | Same order as input array                                        |

---

## Endpoint Inventory

| API                    | New endpoint                                | Domain          | Batch strategy                                    |
| ---------------------- | ------------------------------------------- | --------------- | ------------------------------------------------- |
| Color Conversion       | `POST /v1/technology/color/batch`           | `technology`    | Sequential (pure in-memory math)                  |
| Data Format Conversion | `POST /v1/technology/format/batch`          | `technology`    | Sequential (in-memory encoding libs); max 20      |
| Language Detection     | `POST /v1/text/detect-language/batch`       | `text`          | Sequential (in-memory `lingua-go`)                |
| Thesaurus              | `POST /v1/text/thesaurus/batch`             | `text`          | Sequential (in-memory `thesaurusData` map)        |
| Cities                 | `POST /v1/places/cities/batch`              | `places`        | Sequential (in-memory GeoNames map)               |
| Chuck Norris Facts     | `POST /v1/entertainment/chuck-norris/batch` | `entertainment` | Sequential N random draws (`crypto/rand`)         |
| Dad Jokes              | `POST /v1/entertainment/jokes/dad/batch`    | `entertainment` | Sequential N random draws (`math/rand/v2`)        |
| BIN Lookup             | `POST /v1/finance/bin/batch`                | `finance`       | Single `WHERE bin_prefix = ANY($1)` DB query      |
| Commodity Prices       | `POST /v1/finance/commodities/batch`        | `finance`       | Single `WHERE slug = ANY($1)` DB query; max 20    |
| Crypto Prices          | `POST /v1/finance/crypto/batch`             | `finance`       | Redis cache check + single CoinGecko call; max 20 |
| Domain Info            | `POST /v1/networking/domain/batch`          | `networking`    | Goroutines + semaphore (DNS); max 10              |
| Counter                | `POST /v1/technology/counter/batch`         | `technology`    | Redis pipeline (INCR per namespace)               |

> **Dictionary** (`/v1/text/dictionary/batch`) is deferred —
> `apps/api/services/text/dictionary/` does not exist. Spec this endpoint only
> after single-item service is live.

---

## Per-Endpoint Specs

### 1. Color Format Conversion — `POST /v1/technology/color/batch`

**Files:** `apps/api/services/technology/color/`

Stateless pure math. Sequential loop is correct — no I/O.

**Request / Response:**

```json
// POST /v1/technology/color/batch
{
  "items": [
    { "from": "hex",  "to": "rgb",  "value": "#FF5733" },
    { "from": "rgb",  "to": "hsl",  "value": "255,87,51" },
    { "from": "hex",  "to": "cmyk", "value": "bad" }
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "from": "hex", "to": "rgb",  "input": "#FF5733",    "result": { "r": 255, "g": 87, "b": 51 } },
      { "from": "rgb", "to": "hsl",  "input": "255,87,51",  "result": { "h": 11, "s": 100, "l": 60 } },
      { "from": "hex", "to": "cmyk", "input": "bad",        "error": "invalid hex value" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchColorRequest struct {
    Items []ColorQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

type ColorQuery struct {
    From  string `json:"from"  validate:"required,oneof=hex rgb hsl cmyk"`
    To    string `json:"to"    validate:"required,oneof=hex rgb hsl cmyk"`
    Value string `json:"value" validate:"required"`
}

type BatchColorItem struct {
    From   string    `json:"from"`
    To     string    `json:"to"`
    Input  string    `json:"input"`
    Result *Response `json:"result,omitempty"`
    Error  string    `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) ConvertBatch(items []ColorQuery) []BatchColorItem
```

---

### 2. Data Format Conversion — `POST /v1/technology/format/batch`

**Files:** `apps/api/services/technology/format/`

Stateless in-memory encoding. Sequential loop. Max **20** — each item may be up
to 512 KB; total body must stay ≤ 1 MiB (`http.MaxBytesReader`). Enforce
per-item 512 KB limit in service and return in-band error if exceeded (do not
reject the whole batch).

**Request / Response:**

```json
// POST /v1/technology/format/batch
{
  "items": [
    { "from": "json", "to": "yaml", "content": "{\"key\":\"value\"}" },
    { "from": "yaml", "to": "toml", "content": "key: value" }
  ]
}

// 200 OK
{
  "data": {
    "results": [
      { "from": "json", "to": "yaml", "result": "key: value\n" },
      { "from": "yaml", "to": "toml", "result": "key = \"value\"\n" }
    ],
    "total": 2
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchFormatRequest struct {
    Items []FormatQuery `json:"items" validate:"required,min=1,max=20,dive"`
}

type FormatQuery struct {
    From    string `json:"from"    validate:"required,oneof=json yaml csv xml toml"`
    To      string `json:"to"      validate:"required,oneof=json yaml csv xml toml"`
    Content string `json:"content" validate:"required"`
}

type BatchFormatItem struct {
    From   string `json:"from"`
    To     string `json:"to"`
    Result string `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) ConvertBatch(items []FormatQuery) []BatchFormatItem
// Return in-band error for items where len(content) > maxContentSize (512 KB)
```

---

### 3. Language Detection — `POST /v1/text/detect-language/batch`

**Files:** `apps/api/services/text/detectlanguage/`

Stateless `lingua-go` in-memory. Sequential loop.

**Request / Response:**

```json
// POST /v1/text/detect-language/batch
{
  "texts": ["Hello world", "Bonjour le monde", ""]
}

// 200 OK
{
  "data": {
    "results": [
      { "text": "Hello world",       "result": { "language": "English", "code": "en", "confidence": 0.99 } },
      { "text": "Bonjour le monde",  "result": { "language": "French",  "code": "fr", "confidence": 0.98 } },
      { "text": "",                   "error": "text is required" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchDetectRequest struct {
    Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

type BatchDetectItem struct {
    Text   string  `json:"text"`
    Result *Result `json:"result,omitempty"`
    Error  string  `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) DetectBatch(texts []string) []BatchDetectItem
```

---

### 4. Thesaurus — `POST /v1/text/thesaurus/batch`

**Files:** `apps/api/services/text/thesaurus/`

Stateless in-memory `thesaurusData` map. Sequential loop.

**Request / Response:**

```json
// POST /v1/text/thesaurus/batch
{
  "words": ["happy", "fast", "zzzzz"]
}

// 200 OK
{
  "data": {
    "results": [
      { "word": "happy", "result": { "word": "happy", "synonyms": ["joyful","glad"], "antonyms": ["sad"] } },
      { "word": "fast",  "result": { "word": "fast",  "synonyms": ["quick","swift"], "antonyms": ["slow"] } },
      { "word": "zzzzz", "error": "word not found" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchThesaurusRequest struct {
    Words []string `json:"words" validate:"required,min=1,max=50,dive,required"`
}

type BatchThesaurusItem struct {
    Word   string  `json:"word"`
    Result *Result `json:"result,omitempty"`
    Error  string  `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) LookupBatch(words []string) []BatchThesaurusItem
```

---

### 5. Cities — `POST /v1/places/cities/batch`

**Files:** `apps/api/services/places/cities/`

In-memory GeoNames map. O(1) per lookup. Sequential loop. "Not found" is a valid
outcome — `found: false`, not an error.

**Request / Response:**

```json
// POST /v1/places/cities/batch
{
  "names": ["New York", "Paris", "Faketown"]
}

// 200 OK
{
  "data": {
    "results": [
      { "name": "New York", "found": true,  "result": { "name": "New York", "country": "US", "lat": 40.714, "lon": -74.006, "population": 8336817 } },
      { "name": "Paris",    "found": true,  "result": { ... } },
      { "name": "Faketown", "found": false }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchCitiesRequest struct {
    Names []string `json:"names" validate:"required,min=1,max=50,dive,required"`
}

type BatchCityItem struct {
    Name   string `json:"name"`
    Found  bool   `json:"found"`
    Result *City  `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) FindBatch(names []string) []BatchCityItem
// Reuse existing Find() — it is pure in-memory, no I/O
```

---

### 6. Chuck Norris Facts — `POST /v1/entertainment/chuck-norris/batch`

**Files:** `apps/api/services/entertainment/chucknorris/`

In-memory facts slice. No per-item input — caller requests N random facts.
Repeats are allowed.

**Request / Response:**

```json
// POST /v1/entertainment/chuck-norris/batch
{
  "count": 3
}

// 200 OK
{
  "data": {
    "results": [
      { "id": 42, "fact": "Chuck Norris counted to infinity. Twice." },
      { "id": 7,  "fact": "..." },
      { "id": 19, "fact": "..." }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchChuckRequest struct {
    Count int `json:"count" validate:"required,min=1,max=50"`
}
```

**New service method:**

```go
func (s *Service) RandomBatch(n int) []Fact
// Call crypto/rand for each draw; repeats permitted
```

---

### 7. Dad Jokes — `POST /v1/entertainment/jokes/dad/batch`

**Files:** `apps/api/services/entertainment/jokes/`

In-memory jokes slice. No per-item input — caller requests N random jokes.
Repeats are allowed.

**Request / Response:**

```json
// POST /v1/entertainment/jokes/dad/batch
{
  "count": 3
}

// 200 OK
{
  "data": {
    "results": [
      { "id": "joke_1", "setup": "Why don't scientists trust atoms?", "punchline": "Because they make up everything." },
      { "id": "joke_8", "setup": "...", "punchline": "..." },
      { "id": "joke_3", "setup": "...", "punchline": "..." }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchJokesRequest struct {
    Count int `json:"count" validate:"required,min=1,max=50"`
}
```

**New service method:**

```go
func (s *Service) RandomBatch(n int) []DadJoke
// Use math/rand/v2.IntN(); repeats permitted
```

---

### 8. BIN Lookup — `POST /v1/finance/bin/batch`

**Files:** `apps/api/services/finance/bin/`

PostgreSQL `bin_data` table. Single query fetches all matching rows; re-apply
exact-match vs. prefix fallback logic in-memory per row after fetch.

**Request / Response:**

```json
// POST /v1/finance/bin/batch
{
  "bins": ["424242", "400000", "000000"]
}

// 200 OK
{
  "data": {
    "results": [
      { "bin": "424242", "found": true,  "result": { "scheme": "visa", "bank": "Stripe", "country": "US", "type": "debit" } },
      { "bin": "400000", "found": true,  "result": { ... } },
      { "bin": "000000", "found": false }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchBINRequest struct {
    BINs []string `json:"bins" validate:"required,min=1,max=50,dive,required"`
}

type BatchBINItem struct {
    BIN    string          `json:"bin"`
    Found  bool            `json:"found"`
    Result *LookupResponse `json:"result,omitempty"`
    Error  string          `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) LookupBatch(ctx context.Context, bins []string) []BatchBINItem
// Single query: SELECT ... FROM bin_data WHERE LEFT(bin_number, 6) = ANY($1::text[])
// Group by exact match first, then 6-digit prefix — same precedence as single-item Lookup()
```

---

### 9. Commodity Prices — `POST /v1/finance/commodities/batch`

**Files:** `apps/api/services/finance/commodities/`

PostgreSQL `commodity_price_history` table. Single query with
`WHERE slug = ANY($1)`. Max **20** — each commodity returns up to 11 years of
history (historyDepth = 11).

**Request / Response:**

```json
// POST /v1/finance/commodities/batch
{
  "slugs": ["gold", "silver", "unobtanium"]
}

// 200 OK
{
  "data": {
    "results": [
      { "slug": "gold",       "found": true,  "result": { "name": "Gold", "unit": "troy oz", "currency": "USD", "history": [...] } },
      { "slug": "silver",     "found": true,  "result": { ... } },
      { "slug": "unobtanium", "found": false }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchCommoditiesRequest struct {
    Slugs []string `json:"slugs" validate:"required,min=1,max=20,dive,required"`
}

type BatchCommodityItem struct {
    Slug   string          `json:"slug"`
    Found  bool            `json:"found"`
    Result *CommodityPrice `json:"result,omitempty"`
    Error  string          `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) GetBatch(ctx context.Context, slugs []string) []BatchCommodityItem
// Single query:
//   SELECT slug, year, price, name, unit, currency
//   FROM commodity_price_history
//   WHERE slug = ANY($1)
//   ORDER BY slug, year DESC
// Group rows by slug in-memory; apply historyDepth=11 cap per slug
```

---

### 10. Crypto Prices — `POST /v1/finance/crypto/batch`

**Files:** `apps/api/services/finance/crypto/`

CoinGecko HTTP API + Redis cache (5-min TTL). CoinGecko's `/simple/price`
endpoint accepts a comma-separated `ids` param, so cache misses collapse to
**one outbound HTTP call**. Max **20** — CoinGecko free-tier rate-limit caution.

**Request / Response:**

```json
// POST /v1/finance/crypto/batch
{
  "symbols": ["BTC", "ETH", "XYZ"]
}

// 200 OK
{
  "data": {
    "results": [
      { "symbol": "BTC", "result": { "symbol": "BTC", "name": "Bitcoin",  "price_usd": 65000.00, "change_24h": 1.2 } },
      { "symbol": "ETH", "result": { ... } },
      { "symbol": "XYZ", "error": "unsupported symbol" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchCryptoRequest struct {
    Symbols []string `json:"symbols" validate:"required,min=1,max=20,dive,required"`
}

type BatchCryptoItem struct {
    Symbol string `json:"symbol"`
    Result *Price `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *Service) GetPriceBatch(ctx context.Context, symbols []string) []BatchCryptoItem
// 1. Validate each symbol against coinMap; mark unknown ones in-band (no further processing)
// 2. Check Redis cache for each valid symbol; collect cache misses
// 3. Single CoinGecko call for all misses:
//    GET /simple/price?ids=<comma-sep-ids>&vs_currencies=usd&include_24hr_change=true
// 4. Populate results array; cache fresh responses with 5-min TTL
```

---

### 11. Domain Info — `POST /v1/networking/domain/batch`

**Files:** `apps/api/services/networking/domain/`

Live DNS lookups. The existing `GetInfo()` already fans out 5 parallel DNS
queries per domain via `sync.WaitGroup`. Max **10** — 10 domains × 5 DNS queries
= up to 50 concurrent lookups.

Validate domain format against existing regex before dispatching goroutines;
mark invalid domains in-band without dispatching a goroutine.

**Request / Response:**

```json
// POST /v1/networking/domain/batch
{
  "domains": ["google.com", "github.com", "not-a-valid-domain"]
}

// 200 OK
{
  "data": {
    "results": [
      { "domain": "google.com",       "result": { "a_records": [...], "mx_records": [...], "available": false } },
      { "domain": "github.com",       "result": { ... } },
      { "domain": "not-a-valid-domain", "error": "invalid domain format" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchDomainRequest struct {
    Domains []string `json:"domains" validate:"required,min=1,max=10,dive,required"`
}

type BatchDomainItem struct {
    Domain string        `json:"domain"`
    Result *InfoResponse `json:"result,omitempty"`
    Error  string        `json:"error,omitempty"`
}
```

**New service method:**

```go
const maxDomainWorkers = 5

func (s *Service) GetInfoBatch(ctx context.Context, domains []string) []BatchDomainItem
// Validate domain format inline (existing domainRegex); mark invalid in-band
// Goroutines + semaphore (maxDomainWorkers); per-item context.WithTimeout(ctx, 5s)
// Reuse GetInfo() directly — it already parallelises its own 5 DNS queries
```

---

### 12. Counter — `POST /v1/technology/counter/batch`

**Files:** `apps/api/services/technology/counter/`

Redis (hot) + PostgreSQL (cold). Batch increments all namespaces by 1 each,
matching single-item `POST /counter/{namespace}` semantics. Redis pipeline
reduces N INCR calls to one round-trip. Validate each namespace against existing
regex `^[a-zA-Z0-9_-]{1,64}$` in-band before pipeline dispatch.

> `UpsertBatch()` in `repository.go` is an internal sync-worker method — do
> **not** reuse it here.

**Request / Response:**

```json
// POST /v1/technology/counter/batch
{
  "namespaces": ["page_views", "api_calls", "bad namespace!"]
}

// 200 OK
{
  "data": {
    "results": [
      { "namespace": "page_views",     "value": 1024 },
      { "namespace": "api_calls",      "value": 512 },
      { "namespace": "bad namespace!", "error": "invalid namespace format" }
    ],
    "total": 3
  },
  "metadata": { "timestamp": "2026-06-08T00:00:00Z" }
}
```

**New types (transport_http.go):**

```go
type BatchCounterRequest struct {
    Namespaces []string `json:"namespaces" validate:"required,min=1,max=50,dive,required"`
}

type BatchCounterItem struct {
    Namespace string `json:"namespace"`
    Value     int64  `json:"value,omitempty"`
    Error     string `json:"error,omitempty"`
}
```

**New service method:**

```go
func (s *service) IncrementBatch(ctx context.Context, namespaces []string) []BatchCounterItem
// Validate each namespace against existing namespaceRegex; mark invalid ones in-band
// Redis pipeline: INCR for each valid namespace in one round-trip
// For cold namespaces (missing from Redis), apply existing bootstrapFromDB logic before pipeline
```

---

## Implementation Order

| Step | Service                | Reason                                                               |
| ---- | ---------------------- | -------------------------------------------------------------------- |
| 1    | Chuck Norris Facts     | Simplest: `count` only, no per-item fields, no failure cases         |
| 2    | Dad Jokes              | Same pattern; validates entertainment domain wiring                  |
| 3    | Language Detection     | Single string slice, in-band error for empty text                    |
| 4    | Thesaurus              | Introduces `found`-equivalent via error field                        |
| 5    | Cities                 | Introduces explicit `found: bool` pattern                            |
| 6    | Color Conversion       | Per-item fields, in-band error on bad value                          |
| 7    | Data Format Conversion | Per-item fields + per-item size enforcement                          |
| 8    | BIN Lookup             | First DB batch query; validates `WHERE ... ANY($1)` pattern          |
| 9    | Commodity Prices       | DB batch with in-memory grouping + depth cap                         |
| 10   | Crypto Prices          | Cache-aware batch + single outbound HTTP consolidation               |
| 11   | Counter                | Redis pipeline; validates stateful batch semantics                   |
| 12   | Domain Info            | Most complex; goroutines + semaphore + nested concurrency in GetInfo |

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
- [ ] `apps/dashboard/config/api_catalog.yml` — `batch_eligible: true` and
      `endpoints_count` incremented for each affected API
- [ ] `docs/apis/{category}/{name}.md` — batch endpoint section added
- [ ] `apps/workers/shared/src/config.ts` — no changes; `HandleBatch` sets
      `X-Usage-Count` dynamically

---

## Verification

```bash
# Per-domain (fast feedback during development)
docker exec requiem-dev-api-1 go test ./services/entertainment/chucknorris/...
docker exec requiem-dev-api-1 go test ./services/entertainment/jokes/...
docker exec requiem-dev-api-1 go test ./services/text/detectlanguage/...
docker exec requiem-dev-api-1 go test ./services/text/thesaurus/...
docker exec requiem-dev-api-1 go test ./services/places/cities/...
docker exec requiem-dev-api-1 go test ./services/technology/color/...
docker exec requiem-dev-api-1 go test ./services/technology/format/...
docker exec requiem-dev-api-1 go test ./services/finance/bin/...
docker exec requiem-dev-api-1 go test ./services/finance/commodities/...
docker exec requiem-dev-api-1 go test ./services/finance/crypto/...
docker exec requiem-dev-api-1 go test ./services/technology/counter/...
docker exec requiem-dev-api-1 go test ./services/networking/domain/...

# Full suite with race detection (required before PR)
docker exec requiem-dev-api-1 go test -race -coverprofile=coverage.out ./...

# Smoke test example (replace X-Backend-Secret as needed)
curl -X POST http://localhost:8080/v1/text/thesaurus/batch \
  -H "X-Backend-Secret: local_secret" \
  -H "Content-Type: application/json" \
  -d '{"words":["happy","fast","zzzzz"]}'

# Billing verification — X-Usage-Count must equal item count (or count field value)
curl -v -X POST http://localhost:8080/v1/text/detect-language/batch \
  -H "X-Backend-Secret: local_secret" \
  -H "Content-Type: application/json" \
  -d '{"texts":["hello","world","!"]}' 2>&1 | grep X-Usage-Count
# Expected: X-Usage-Count: 3

# YAML validation (run for each updated doc)
docker exec requiem-dev-dashboard-1 ruby -ryaml \
  -e "YAML.load_file('config/api_docs/color-conversion.yml'); puts 'OK'"

# Catalog verification
grep -A 3 'id: disposable-email' apps/dashboard/config/api_catalog.yml
# Expected: batch_eligible: true
```
