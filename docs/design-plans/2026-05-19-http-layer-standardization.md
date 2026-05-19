# HTTP Layer Standardization

The 2026-05-18 service layer cleanup consolidated types and removed `type.go`
files. That work left three structural problems unaddressed: a marker interface
that has outlived its purpose, GET endpoints that bypass the standard
`Handle`/`HandleBatch` validation path, and domain types in `service.go` that
carry HTTP-platform concerns.

This document captures the current state of each problem and the acceptance
criteria that define "done."

---

## Problem 1: `Data` Interface Provides No Semantic Value

`httpx.Data` is a marker interface with a single empty method `IsData()`.
It was introduced so `Handle` and `httpx.JSON` could use it as a type
constraint, preventing arbitrary types from being passed as response payloads.
In practice it enforces nothing — any type can implement it with a one-liner —
and it has already been abandoned in `HandleBatch`, which constrains `Item` as
`any`. The interface now creates work without providing safety.

The deeper cost is coupling: every domain type that flows through a
`Handle`-wrapped endpoint must implement `IsData()`. This leaks an HTTP-platform
concern into service-layer files that should know nothing about transport.

### Current Status

`IsData()` is implemented by **65+ types** spread across every service package.
Every domain entity in the codebase has been annotated with this marker:

| Category | Example types with `IsData()` |
| --- | --- |
| `technology/` | `Result`, `Response`, `Password`, `Base64Response`, `User`, `Counter` |
| `places/` | `PostalCode`, `Info`, `Holiday`, `HolidayList`, `City`, `WorkingDays`, `GeocodeResponse` |
| `networking/` | `LookupResponse`, `IPCheckResponse`, `IPAddressASNResponse`, `InfoResponse`, `CheckEmailResponse` |
| `health/` | `Exercise`, `ExerciseList`, `StringList` |
| `entertainment/` | `Horoscope`, `Quote`, `Emoji`, `Question`, `Advice`, `Puzzle`, `Fact`, `DadJoke` |
| `finance/` | `Price`, `RateResponse`, `ConvertResponse`, `LookupResponse`, `CommodityPrice`, `ParseResponse`, `Response` |
| `text/` | `Word`, `DictionaryEntry`, `Result`, `Lorem`, `EmailNormalization` |
| `validation/` | `Validation`, `ValidateResponse`, `BatchValidateResponse`, `Result` |

`BatchResponse[T]` in `apps/api/platform/httpx/httpx.go:23` also implements
`IsData()` — that is acceptable since it lives in the platform package, but it
signals the constraint is purely structural.

`Handle` and `httpx.JSON` are the only remaining callers of the `Data`
constraint. `HandleBatch` already uses `any`.

### Acceptance Criteria

- [ ] `Data` interface removed from `apps/api/platform/httpx/httpx.go`
- [ ] `Handle[Req any, Res Data]` → `Handle[Req any, Res any]`
- [ ] `JSON[T Data]` → `JSON[T any]`, `Response[T Data]` → `Response[T any]`
- [ ] All `IsData()` method implementations removed from `services/`
- [ ] Tests that exist solely to call `IsData()` removed (e.g.
      `technology/units/service_test.go:217`, `health/exercises/service_test.go:13-28`,
      `finance/commodities/service_test.go:39`)
- [ ] `go build ./...` and `go test ./...` pass

---

## Problem 2: GET Endpoints Bypass the Handle Pattern

`Handle` and `HandleBatch` provide a consistent contract: bind the request,
validate via struct tags, call the service, write a structured response or a
typed error envelope. Every POST endpoint uses this path. Every GET endpoint
does not — because `Handle` reads from the JSON body and GET requests have no
body.

Instead, GET handlers call `httpx.BindQuery` manually, then branch on the error
themselves. This produces duplicated boilerplate across 22 files, and the error
responses produced by manual branches are not guaranteed to match the `Handle`
envelope shape.

### Current Status

**28 `httpx.BindQuery` call sites across 22 `transport_http.go` files:**

| File | Lines |
| --- | --- |
| `technology/useragent/transport_http.go` | 32 |
| `technology/password/transport_http.go` | 23 |
| `technology/numbase/transport_http.go` | 23 |
| `technology/color/transport_http.go` | 26 |
| `technology/qr/transport_http.go` | 26, 46 |
| `technology/barcode/transport_http.go` | 23, 43 |
| `places/timezone/transport_http.go` | 65 |
| `places/holidays/transport_http.go` | 33 |
| `places/working-days/transport_http.go` | 24 |
| `places/geocode/transport_http.go` | 20, 45 |
| `networking/disposable/transport_http.go` | 57 |
| `health/exercises/transport_http.go` | 58, 75 |
| `entertainment/emoji/transport_http.go` | 26 |
| `entertainment/trivia/transport_http.go` | 22 |
| `entertainment/sudoku/transport_http.go` | 29 |
| `entertainment/facts/transport_http.go` | 22 |
| `finance/exchange/transport_http.go` | 47, 76 |
| `finance/swift/transport_http.go` | 41 |
| `finance/inflation/transport_http.go` | 51 |
| `finance/mortgage/transport_http.go` | 43 |
| `text/lorem/transport_http.go` | 26 |
| `validation/phone/transport_http.go` | 26 |

The repeating pattern in each file:

```go
r.Get("/endpoint", func(w http.ResponseWriter, r *http.Request) {
    var req Request
    if err := httpx.BindQuery(r, &req); err != nil {
        httpx.ValidationError(w, errors.AsType[*httpx.ValidationFailure](err))
        return
    }
    httpx.JSON(w, http.StatusOK, svc.Do(r.Context(), req))
})
```

Some files (e.g. `places/sudoku`) set struct field defaults before calling
`BindQuery`. `HandleGet` must accommodate a pre-bind hook or the conversion
must handle defaults via struct tag defaults.

### Acceptance Criteria

- [ ] `HandleGet[Req any, Res any]` added to
      `apps/api/platform/httpx/handler.go` — same error path as `Handle` but
      binds query params instead of JSON body
- [ ] All 28 manual `BindQuery` call sites replaced with `HandleGet`
- [ ] Any per-endpoint defaults set before `BindQuery` migrated to struct field
      initializers or `HandleGet` pre-bind option
- [ ] `go build ./...` and `go test ./...` pass
- [ ] Spot-check: validation error shape from a bad GET query matches the shape
      returned by a bad POST body

---

## Problem 3: Service Types Carry HTTP Concerns

A domain type in `service.go` should represent a business concept. It should
not know whether it will be served over HTTP, gRPC, or a CLI. The `IsData()`
marker (Problem 1) is the most visible symptom, but the broader rule is:
`service.go` must not import `net/http`, `chi`, or `httpx`. The HTTP layer
calls the service; the service does not reach back into the HTTP layer.

The previous cleanup plan (2026-05-18) established this rule and moved types
from `type.go` into the correct files. It did not remove `IsData()` from
service types or enforce the import constraint programmatically.

### Current Status

**Import violations in `service.go` files** (inbound HTTP — wrong):

All 65+ `IsData()` implementations listed in Problem 1 represent implicit
coupling: the service type is shaped by the HTTP platform's constraint. Removing
`Data` (Problem 1) eliminates this coupling without requiring import changes,
since `IsData()` itself does not pull in any import.

**Legitimate `net/http` usage in `service.go` files** (outbound clients — keep):

| File | Purpose |
| --- | --- |
| `places/geocode/service.go` | `http.Client` calling Nominatim geocoding API |
| `finance/crypto/service.go` | `http.Client` calling CoinGecko API |
| `finance/exchange/service.go` | `http.Client` calling Frankfurter exchange-rate API |

These are integration clients. The service layer calling an external HTTP API is
correct; the service layer handling inbound HTTP is not.

**Response struct placement:** `transport_http.go` should define the HTTP
request and response structs with `json:`, `query:`, and `validate:` tags.
Where a batch endpoint returns `httpx.BatchResponse[DomainType]`, the type
alias or construction must live in `transport_http.go`, not `service.go`.

### Acceptance Criteria

- [ ] No `service.go` file (outside the three legitimate outbound clients above)
      imports `net/http`, `github.com/go-chi/chi/v5`, or any `httpx` sub-package
- [ ] No `service.go` file defines a type whose sole purpose is satisfying an
      HTTP-layer interface (all such types removed as part of Problem 1)
- [ ] Every `transport_http.go` has a clearly named `Request` struct (with
      `validate:` tags) and a `Response` struct or uses `httpx.BatchResponse[T]`
      directly as the return type
- [ ] `go build ./...` and `go test ./...` pass
- [ ] Enforced by CI: `grep -r 'IsData' apps/api/services/` returns empty
