# PR: Batch APIs Expansion (9 new batch endpoints)

## Summary

Adds `POST .../batch` endpoints to nine existing single-item APIs. Phone validation already had a batch endpoint and served as the reference implementation throughout. Every endpoint follows the RFC in `docs/core/batch-apis.md` without deviation.

Also fixes two pre-existing `camelCase` JSON field name violations (`wordCount` → `word_count`, `partOfSpeech` → `part_of_speech`) caught while touching adjacent code.

---

## New Endpoints

| Endpoint | Domain | Max items | Batch strategy |
|----------|--------|-----------|----------------|
| `POST /v1/validation/profanity/batch` | `validation` | 50 | Sequential (in-memory `go-away`) |
| `POST /v1/technology/base/batch` | `technology` | 50 | Sequential (pure math) |
| `POST /v1/technology/password/batch` | `technology` | 50 | Sequential (`crypto/rand`) |
| `POST /v1/places/postal/batch` | `places` | 50 | Sequential (in-memory GeoNames map) |
| `POST /v1/places/timezone/batch` | `places` | 50 | Sequential (in-memory `tzf`) |
| `POST /v1/technology/qr/base64/batch` | `technology` | 50 | Sequential (in-memory QR gen) |
| `POST /v1/networking/ip/info/batch` | `networking` | 50 | Sequential (local MaxMind DB) |
| `POST /v1/networking/ip/asn/batch` | `networking` | 50 | Sequential (local MaxMind DB) |
| `POST /v1/places/geocode/batch` | `places` | 20 | Goroutines + semaphore (Nominatim HTTP) |
| `POST /v1/places/reverse-geocode/batch` | `places` | 20 | Goroutines + semaphore (Nominatim HTTP) |

> **QR PNG** (`GET /v1/technology/qr`) has no batch variant — it returns `image/png` binary, incompatible with the JSON batch envelope. Only `/qr/base64/batch` is added.

---

## Design Decisions

**Error model — partial success.** Every batch endpoint returns HTTP 200 with the full result set. Per-item failures are absorbed in-band via an `error` string field on that item. Top-level errors are only returned for systemic failures (context cancellation, unrecoverable service init failure). This matches the existing phone batch and the RFC.

**Billing.** All endpoints use `httpx.HandleBatch`, which sets the `X-Usage-Count` response header to `len(results)`. The auth gateway reads this header and charges N units for N items. No entries in `apps/workers/shared/src/config.ts` are needed — the header supersedes the static multiplier map.

**Sequential vs. concurrent.** Services backed by in-memory data or local MaxMind DB files use a sequential loop — no I/O, negligible latency per item. Only geocoding fans out with goroutines + semaphore (`maxWorkers = 10`, `context.WithTimeout` of 3 s per item) because it makes real HTTP calls to the Nominatim API.

**Geocoding cap is 20, not 50.** Nominatim's usage policy limits request rate. A cap of 20 items with a 10-worker semaphore and 3 s per-item timeout means worst-case wall time stays bounded and well within reasonable HTTP timeout limits.

**`httpx.Guard` on nullable services.** Timezone, IP info, and ASN services can return `nil` from their constructors (MaxMind DB missing, finder init failure). Their batch routes are wrapped with `httpx.Guard(svc, ...)` — matching the existing single-item routes in those packages.

**Timezone multi-mode input.** The single-item API has two routes (`/time/*` by IANA name, `/timezone` by city or coords). The batch collapses this into one unified item struct with priority: IANA name → city name → lat+lon. Items supplying none of these return an in-band error.

**Password length default.** The GET handler applies `Length = 16` via HandleGet defaults. The batch POST handler cannot use that mechanism — the service method applies the default when `Length == 0`. The validate tag uses `omitempty` so zero is accepted and defaulted, not rejected.

---

## Changed Files

```
apps/api/services/validation/profanity/service.go        + CheckBatch, BatchResult
apps/api/services/validation/profanity/transport_http.go + BatchRequest, POST /profanity/batch

apps/api/services/technology/numbase/service.go          + ConvertQuery, BatchResult, ConvertBatch
apps/api/services/technology/numbase/transport_http.go   + BatchConvertRequest, POST /base/batch

apps/api/services/technology/password/service.go         + PasswordQuery, BatchResult, GenerateBatch
apps/api/services/technology/password/transport_http.go  + BatchPasswordRequest, POST /password/batch

apps/api/services/places/postal/service.go               + PostalQuery, BatchResult, LookupBatch
apps/api/services/places/postal/transport_http.go        + BatchPostalRequest, POST /postal/batch

apps/api/services/places/timezone/service.go             + TimezoneQuery, BatchResult, BatchLookup
apps/api/services/places/timezone/transport_http.go      + BatchTimezoneRequest, POST /timezone/batch

apps/api/services/technology/qr/service.go               + QRQuery, BatchBase64Item, GenerateBatch
apps/api/services/technology/qr/transport_http.go        + BatchQRRequest, POST /qr/base64/batch

apps/api/services/networking/ip/info/service.go          + BatchIPInfoItem, CheckInfoBatch
apps/api/services/networking/ip/info/transport_http.go   + BatchInfoRequest, POST /ip/info/batch

apps/api/services/networking/ip/asn/service.go           + BatchASNItem, CheckASNBatch
apps/api/services/networking/ip/asn/transport_http.go    + BatchASNRequest, POST /ip/asn/batch

apps/api/services/places/geocode/service.go              + ReverseQuery, BatchGeocodeItem,
                                                           BatchReverseGeocodeItem, GeocodeBatch,
                                                           ReverseGeocodeBatch
apps/api/services/places/geocode/transport_http.go       + GeocodeBatchRequest, ReverseGeocodeBatchRequest,
                                                           POST /geocode/batch, POST /reverse-geocode/batch
```

**Bug fixes:**
```
apps/api/services/text/lorem/service.go   wordCount     → word_count     (snake_case fix)
apps/api/services/text/words/service.go   partOfSpeech  → part_of_speech (snake_case fix)
```

---

## Testing

All existing tests pass. No existing test was modified.

```
go test -count=1 ./...
```

```
ok  requiems-api/services/validation/profanity   0.327s
ok  requiems-api/services/technology/numbase     0.177s
ok  requiems-api/services/technology/password    0.486s
ok  requiems-api/services/places/postal          0.653s
ok  requiems-api/services/places/timezone        1.367s
ok  requiems-api/services/technology/qr          1.263s
ok  requiems-api/services/networking/ip/info     0.957s
ok  requiems-api/services/networking/ip/asn      1.101s
ok  requiems-api/services/places/geocode         1.402s
... (all 66 packages, zero failures)
```

---

## Smoke Tests

```bash
# Profanity batch
curl -X POST http://localhost:8080/v1/validation/profanity/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"texts":["hello world","offensive word here"]}'

# Number base conversion batch
curl -X POST http://localhost:8080/v1/technology/base/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"items":[{"from":10,"to":2,"value":"255"},{"from":16,"to":10,"value":"FF"}]}'

# Password batch
curl -X POST http://localhost:8080/v1/technology/password/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"items":[{"length":16,"uppercase":true,"numbers":true},{"length":32,"symbols":true}]}'

# Postal code batch
curl -X POST http://localhost:8080/v1/places/postal/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"items":[{"code":"10001","country":"US"},{"code":"SW1A1AA","country":"GB"}]}'

# Timezone batch (mixed query modes)
curl -X POST http://localhost:8080/v1/places/timezone/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"items":[{"timezone":"America/New_York"},{"city":"Tokyo"},{"lat":51.5074,"lon":-0.1278}]}'

# QR base64 batch
curl -X POST http://localhost:8080/v1/technology/qr/base64/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"items":[{"data":"https://requiems.xyz"},{"data":"hello world","size":128}]}'

# IP geolocation batch
curl -X POST http://localhost:8080/v1/networking/ip/info/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"ips":["8.8.8.8","1.1.1.1"]}'

# ASN lookup batch
curl -X POST http://localhost:8080/v1/networking/ip/asn/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"ips":["8.8.8.8","1.1.1.1"]}'

# Geocode batch (forward)
curl -X POST http://localhost:8080/v1/places/geocode/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"addresses":["Eiffel Tower, Paris","Times Square, New York"]}'

# Reverse geocode batch
curl -X POST http://localhost:8080/v1/places/reverse-geocode/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"items":[{"lat":48.8584,"lon":2.2945},{"lat":40.7580,"lon":-73.9855}]}'

# Billing check — X-Usage-Count must equal item count
curl -sv -X POST http://localhost:8080/v1/validation/profanity/batch \
  -H "X-Backend-Secret: $SECRET" -H "Content-Type: application/json" \
  -d '{"texts":["a","b","c"]}' 2>&1 | grep X-Usage-Count
# Expected: X-Usage-Count: 3
```

---

## Checklist

- [x] `go build ./...` — clean
- [x] `go test -count=1 ./...` — all 66 packages pass, zero failures
- [x] All batch endpoints use `httpx.HandleBatch` (sets `X-Usage-Count` automatically)
- [x] No sequential loop over single-item method for I/O-backed services (geocoding fans out with goroutines)
- [x] Partial success: per-item `error` field, top-level 200
- [x] Result order matches input order
- [x] Nullable services (timezone, ip/info, ip/asn) use `httpx.Guard`
- [x] Struct validation enforces `min=1,max=N` on all batch arrays
- [x] No changes to `router.go` or `routes_v1.go` — all services already mounted
- [x] No changes to `apps/workers/shared/src/config.ts` — `HandleBatch` sets header dynamically
- [ ] Dashboard YAML docs updated (follow-up)
- [ ] `api_catalog.yml` `endpoints_count` incremented (follow-up)
