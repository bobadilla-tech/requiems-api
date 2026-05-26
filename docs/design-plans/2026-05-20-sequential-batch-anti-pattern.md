# Sequential Batch Anti-pattern Fix

Several `*Batch` functions across the API advertised batch throughput but
implemented it as a sequential `for` loop over the single-item method — making N
serial I/O calls. For network endpoints (DNS, HTTP) this means wall-clock time
scales linearly with input size, eliminating any benefit over N individual
calls.

## Problem

Go has no type or lint rule that prevents a "batch" function from being a hidden
sequential loop. The functions compiled and tested correctly — the anti-pattern
was invisible until a reviewer measured the actual call graph.

```go
// Looks like a batch. Makes N serial DNS lookups.
func (s *Service) LookupBatch(ctx context.Context, domains []string) []BatchLookupItem {
    results := make([]BatchLookupItem, len(domains))
    for i, d := range domains {
        result, err := s.Lookup(ctx, d) // one DNS round-trip per iteration
        ...
        results[i] = ...
    }
    return results
}
```

For 50 domains, this is 50 sequential DNS queries. A caller expecting batch
throughput gets the same wall time as 50 individual API calls.

---

## Status Before Changes

### Affected Functions

| File                                       | Function            | I/O type | Before behavior                            |
| ------------------------------------------ | ------------------- | -------- | ------------------------------------------ |
| `services/networking/mx/service.go`        | `LookupBatch`       | DNS      | Sequential `for` loop over `Lookup`        |
| `services/finance/inflation/service.go`    | `GetInflationBatch` | DB query | Sequential `for` loop over `GetInflation`  |
| `services/finance/iban/service.go`         | `ParseBatch`        | DB query | Sequential `for` loop over `Parse`         |
| `services/entertainment/quotes/service.go` | `RandomBatch`       | DB query | N × `SELECT ... ORDER BY random() LIMIT 1` |

### Already-Correct Implementations (reference)

Two services already had the right pattern before this change:

- `services/networking/whois/service.go` — `LookupBatch`: goroutines + semaphore
  channel, `maxWorkers=10`, 3s per-item timeout
- `services/validation/email/service.go` — `ValidateEmailBatch`: goroutines +
  semaphore channel, `maxWorkers=8`

These served as the implementation template.

---

## Fix

Two strategies applied depending on the I/O type.

### Pattern A — Goroutines + Semaphore (network / DB fan-out)

Used when each item needs its own I/O call and calls are independent.

```go
const maxWorkers = 10

results := make([]BatchLookupItem, len(domains))
sem := make(chan struct{}, maxWorkers)
var wg sync.WaitGroup

for i, d := range domains {
    wg.Add(1)
    sem <- struct{}{}
    go func(i int, d string) {
        defer wg.Done()
        defer func() { <-sem }()

        itemCtx, cancel := context.WithTimeout(ctx, perItemTimeout)
        defer cancel()

        result, err := s.Lookup(itemCtx, d)
        if err != nil {
            results[i] = BatchLookupItem{Domain: d, Found: false, Error: err.Error()}
            return
        }
        results[i] = BatchLookupItem{Domain: d, Found: true, Data: result}
    }(i, d)
}

wg.Wait()
```

The semaphore caps concurrency so a large batch can't exhaust the DNS resolver
or DB connection pool. Results are written to pre-allocated index slots, so
input order is always preserved without a sort.

Applied to:

| Function                      | `maxWorkers` | Per-item timeout        |
| ----------------------------- | ------------ | ----------------------- |
| `mx.LookupBatch`              | 10           | 3s                      |
| `inflation.GetInflationBatch` | 8            | — (inherits caller ctx) |
| `iban.ParseBatch`             | 8            | — (inherits caller ctx) |

### Pattern B — Single Batch Query (DB, when all items share one table)

Used when the N items can be collapsed into one SQL statement.

```go
// Before: N queries
for i := 0; i < n; i++ {
    row := s.db.QueryRow(ctx, `SELECT ... FROM quotes ORDER BY random() LIMIT 1`)
    ...
}

// After: 1 query
rows, err := s.db.Query(ctx, `
SELECT id, text, author
FROM quotes
ORDER BY random()
LIMIT $1`, n)
```

Applied to `quotes.RandomBatch` — reduced N DB round-trips to 1.

### Interface Change (quotes only)

`platform/db.Querier` only exposes `QueryRow`. Expanding it would cascade into
all mocks across four services. Instead, a local `quotesDB` interface was
defined in `quotes/service.go`:

```go
type quotesDB interface {
    QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
    Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}
```

`*pgxpool.Pool` satisfies it directly. The shared `db.Querier` interface remains
unchanged.

---

## Acceptance Criteria

- [ ] `go test ./services/networking/mx/...` — pass
- [ ] `go test ./services/finance/inflation/...` — pass
- [ ] `go test ./services/finance/iban/...` — pass
- [ ] `go test ./services/entertainment/quotes/...` — pass
- [ ] `go build ./...` — clean
- [ ] No `*Batch` function with network/DB I/O uses a sequential `for` loop over
      the single-item method

All criteria met at time of merge.

---

## Lessons Learned

**1. "Batch" is a contract, not a label.** A function named `*Batch` promises
throughput benefit. A sequential loop breaks that promise silently — no compiler
error, no test failure, no runtime panic. Reviewers must check the call graph,
not just the function signature.

**2. Two failure modes, two fixes.** "Sequential I/O" can mean two different
things:

- N independent calls that could run in parallel → goroutines + semaphore
- N queries that could be collapsed into one → SQL `IN` / `ANY` / `LIMIT n`

Applying goroutines to `quotes.RandomBatch` would still make N DB round-trips.
The right diagnosis determines the right fix.

**3. Minimal shared interfaces have hidden costs.** `db.Querier` was
intentionally narrow (`QueryRow` only — "services that only need single-row
queries"). When one service evolved beyond that, expanding the shared interface
would have required updating mocks in four unrelated packages. Scoping a local
interface to the one package that needs it is the right tradeoff: isolation over
uniformity.

**4. Tests encode the implementation, not just the contract.**
`quotes.RandomBatch` tests used `multiMockQuerier` to simulate N sequential
`QueryRow` calls and verify ordering. When the implementation changed to a
single `Query` call, the tests were invalid — not just failing, but testing the
wrong thing. Implementation changes that affect the call shape require
rethinking the test setup, not just adjusting assertions.

**5. Positive examples are the cheapest template.** `whois.LookupBatch` already
had the correct goroutine + semaphore pattern. Every new fix was a copy of that
pattern with adjusted constants. No design work was needed — just recognition
that the reference already existed.

**6. Batch functions must always use in-band errors.** PR #650 review suggested
`errgroup` (one error aborts the entire batch). That is wrong for this codebase.
A failed item gets `Found: false` / `Valid: false` / an `error` field — the
batch continues and returns 200. Fail-fast (`errgroup`) is never acceptable for
batch service methods: it punishes all items in a request for one item's failure
and breaks the partial-success contract documented in `docs/core/batch-apis.md`.
