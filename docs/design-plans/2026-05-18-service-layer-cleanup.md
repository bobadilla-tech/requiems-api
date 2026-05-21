# Service Layer Cleanup & Standardization

The Go backend (`apps/api/services/`) has ~60 modules each typically organized
around two primary files: `service.go` and `transport_http.go`. Historically
types were maintained separately which led to types being collected away from
their owning code. Additionally, 4 modules independently define an identical
`querier` interface, and some service methods duplicate validation already
enforced by struct tags.

---

## Goals

1. Types should live in the file that owns them.
2. `service.go` — domain entities + service logic only. No HTTP-specific
   imports.
3. `transport_http.go` — HTTP handlers + request/response types (with
   `validate:` tags).
4. No redundant validation in service methods for rules already covered by
   struct tags.
5. Standardize the repeated `querier` interface into a shared `platform/db`
   type.

---

## Type Placement Rules

| Type category                                               | Target file         |
| ----------------------------------------------------------- | ------------------- |
| Domain entities (`Quote`, `Exercise`, `ExerciseList`)       | `service.go`        |
| Service parameter structs (no HTTP-specific tags)           | `service.go`        |
| HTTP request structs with `json:`/`query:`/`validate:` tags | `transport_http.go` |
| HTTP response type aliases (`= httpx.BatchResponse[X]`)     | `transport_http.go` |

Since all types in a module share one Go package, placement is semantic, not
functional.

---

## Shared Querier Interface

Four modules define this identically:

```go
type querier interface {
    QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
```

Modules: `entertainment/quotes`, `entertainment/advice`, `finance/iban`,
`text/words`.

**Fix:** Add `platform/db/interfaces.go`:

```go
package db

import (
    "context"
    "github.com/jackc/pgx/v5"
)

// Querier is the minimal DB interface for services that only need single-row queries.
// *pgxpool.Pool satisfies this interface directly.
type Querier interface {
    QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
```

Each of the 4 services drops its local interface and uses `db.Querier` instead.

**Note:** `health/exercises` has a more complex `dbPool`/`dbRows`/`poolWrapper`
pattern that enables `Query` mocking in tests. That is intentional and is not
changed.

---

## Validation Rules

Service methods must not re-check constraints already enforced by struct tags.
The `httpx.BindAndValidate` / `httpx.Handle` / `httpx.HandleBatch` wrappers
validate before the service is called — duplicate guards are dead code in
production.

**Remove:** bounds/range checks that mirror `validate:"min=X,max=Y"` tags.
**Keep:** semantic invariants the HTTP layer cannot enforce (e.g. checking DB
results are non-empty, context propagation).

---

## Module Change Table

| Module                    | querier update | types → service.go                 | types → transport_http.go                                                  | validation removal                 |
| ------------------------- | -------------- | ---------------------------------- | ------------------------------------------------------------------------- | ---------------------------------- |
| entertainment/advice      | `db.Querier`   | `Advice`                           | —                                                                         | —                                  |
| entertainment/chucknorris | —              | `ChuckNorrisFact`                  | —                                                                         | —                                  |
| entertainment/emoji       | —              | `Emoji`                            | `SearchRequest`, `BatchGetRequest`, `EmojiResponse`, `BatchEmojiResponse` | —                                  |
| entertainment/facts       | —              | `Fact`, `Category`                 | —                                                                         | —                                  |
| entertainment/horoscope   | —              | `Horoscope`, `Sign`                | `HoroscopeRequest`                                                        | —                                  |
| entertainment/jokes       | —              | `Joke`                             | `BatchRandomRequest`, `BatchRandomResponse`                               | —                                  |
| entertainment/quotes      | `db.Querier`   | `Quote`                            | `BatchRandomRequest`, `BatchRandomResponse`                               | remove `if n < 1` in `RandomBatch` |
| entertainment/sudoku      | —              | `Board`                            | `GenerateRequest`, `SolveRequest`, `GenerateResponse`, `SolveResponse`    | —                                  |
| entertainment/trivia      | —              | `Question`                         | `BatchRandomRequest`, `BatchRandomResponse`                               | —                                  |
| finance/bin               | —              | `BINLookup`                        | `LookupRequest`                                                           | —                                  |
| finance/commodities       | —              | `Commodity`                        | `BatchGetRequest`, `BatchCommodityResponse`                               | —                                  |
| finance/crypto            | —              | `CryptoQuote`                      | `BatchGetRequest`, `BatchCryptoResponse`                                  | —                                  |
| finance/exchange          | —              | `ExchangeRate`, `ConversionResult` | `RateRequest`, `ConvertRequest`                                           | —                                  |
| finance/iban              | `db.Querier`   | `IBANResult`                       | `ValidateRequest`                                                         | —                                  |
| finance/inflation         | —              | `InflationRate`                    | `BatchGetRequest`, `BatchInflationResponse`                               | —                                  |
| finance/mortgage          | —              | `MortgageResult`                   | `CalculateRequest`                                                        | —                                  |
| finance/swift             | —              | `SWIFTRecord`                      | `LookupRequest`                                                           | —                                  |
| health/exercises          | —              | (already in service.go)            | `ListParams`, `BatchGetRequest`, `BatchExerciseResponse`                  | —                                  |
| networking/disposable     | —              | domain types                       | —                                                                         | —                                  |
| networking/domain         | —              | domain types                       | request types                                                             | —                                  |
| networking/ip/asn         | —              | domain types                       | request types                                                             | —                                  |
| networking/ip/info        | —              | domain types                       | request types                                                             | —                                  |
| networking/ip/vpn         | —              | domain types                       | request types                                                             | —                                  |
| networking/mx             | —              | domain types                       | request types                                                             | —                                  |
| networking/whois          | —              | domain types                       | request types                                                             | —                                  |
| places/cities             | —              | domain types                       | request types                                                             | —                                  |
| places/geocode            | —              | domain types                       | request types                                                             | —                                  |
| places/holidays           | —              | domain types                       | request types                                                             | —                                  |
| places/postal             | —              | domain types                       | request types                                                             | —                                  |
| places/timezone           | —              | domain types                       | request types                                                             | —                                  |
| places/working-days       | —              | domain types                       | request types                                                             | —                                  |
| technology/barcode        | —              | domain types                       | request types                                                             | —                                  |
| technology/base64         | —              | domain types                       | request types                                                             | —                                  |
| technology/color          | —              | domain types                       | request types                                                             | —                                  |
| technology/format         | —              | domain types                       | request types                                                             | —                                  |
| technology/markdown       | —              | domain types                       | request types                                                             | —                                  |
| technology/numbase        | —              | domain types                       | request types                                                             | —                                  |
| technology/password       | —              | domain types                       | request types                                                             | —                                  |
| technology/qr             | —              | domain types                       | request types                                                             | —                                  |
| technology/random_user    | —              | domain types                       | request types                                                             | —                                  |
| technology/units          | —              | domain types                       | request types                                                             | —                                  |
| technology/useragent      | —              | domain types                       | request types                                                             | —                                  |
| text/detectlanguage       | —              | domain types                       | request types                                                             | —                                  |
| text/lorem                | —              | domain types                       | request types                                                             | —                                  |
| text/normalize            | —              | domain types                       | request types                                                             | —                                  |
| text/sentiment            | —              | domain types                       | request types                                                             | —                                  |
| text/similarity           | —              | domain types                       | request types                                                             | —                                  |
| text/spellcheck           | —              | domain types                       | request types                                                             | —                                  |
| text/thesaurus            | —              | domain types                       | request types                                                             | —                                  |
| text/words                | `db.Querier`   | domain types                       | request types                                                             | —                                  |
| validation/email          | —              | domain types                       | request types                                                             | —                                  |
| validation/phone          | —              | domain types                       | request types                                                             | —                                  |
| validation/profanity      | —              | domain types                       | request types                                                             | —                                  |

---

## Acceptance Criteria

- `go build ./...` — clean
- `go test ./...` — all pass
- No standalone type-only files in `services/`
- `service.go` has no `net/http`, `chi`, `httpx` imports (except `counter/`
  which has its own structure)
- `transport_http.go` has no `querier`/`dbPool` interface definitions
- No duplicate range/min/max guards in service methods

---

## Deployment Checklist

- [ ] `go build ./...` passes locally
- [ ] `go test ./...` passes locally
- [ ] No standalone type-only files remain (types are declared in `service.go` or `transport_http.go`)
      returns 0
- [ ] Spot-check: `GET /v1/entertainment/quotes/random` still works
- [ ] Spot-check: `POST /v1/health/exercises/batch` still works

---

## What Was Not Changed

- Business logic — no service behaviour altered
- Response shapes — all JSON responses identical
- HTTP routes — no path changes
- Database schema — no migrations
- `counter/` module — unique repository pattern, left as-is
