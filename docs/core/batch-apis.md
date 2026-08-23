# Batch APIs

> request shape, billing, and Go handler pattern

## 1. Goals

1. **Predictable JSON** — Request and response shapes are easy to validate and
   document.
2. **Fair usage** — Customers are charged in proportion to **work done** (items
   processed), not only per HTTP call.
3. **Operational safety** — Strict body size limits, bounded batch sizes, and
   clear errors when limits are exceeded.
4. **Gateway compatibility** — The auth gateway can apply the correct **usage
   multiplier** without custom per-route hacks where possible.

## Reference implementation

**This document** and live **`httpx.HandleBatch`** handlers are the normative
reference. Public request/response details live next to each API’s docs.

Example: **Phone validation** — `POST /v1/validation/phone/batch`

**Go wiring:** `httpx.HandleBatch` in the validation phone router. The handler
returns `(response, itemCount, error)`; the platform sets response header
**`X-Usage-Count`** to `itemCount` (stringified integer).

**Request (illustrative):**

- JSON object with an array field (here `numbers`).
- Validation tags on the struct enforce **min/max length** of the array (phone:
  1–50 items).

**Response:**

- Same **`{"data": …, "metadata": …}`** success envelope as all `httpx.JSON` /
  `Handle` / `HandleGet` / `HandleBatch` responses.

**Errors:**

- Malformed input or validation failure → **`422`** with
  `error: "validation_failed"` and `fields` (standard `httpx` behavior).
- Domain errors from the service → map with `httpx.AppError` as today.

**Ordering:** Results MUST stay in the **same order** as the input array unless
the doc for that specific API documents otherwise.

---

## Auth gateway and billing

1. The Go process attaches **`X-Usage-Count`** on the response to the gateway
   (internal; not a public contract for API clients).
2. The auth gateway reads that header when recording usage. If it is a
   **positive integer**, it becomes the **effective multiplier** for that
   request; otherwise the gateway falls back to the static **per-route
   multiplier** map (default **1**).
3. The gateway **strips** `X-Usage-Count` before returning the response to the
   public client, so keys are not exposed to billing internals.

**Important:** Any new batch endpoint that must consume **N** units of quota for
**N** items **must** use `HandleBatch` (or equivalent that sets `X-Usage-Count`
consistently). Otherwise a single HTTP request that processes 50 items might
still count as **1** request.

---

## Standards for _new_ batch endpoints

| Topic                        | Standard                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **HTTP method**              | `POST`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| **Path naming**              | Prefer **`/resource/batch`** under the same prefix as single-item routes (e.g. `/phone/batch`). If the domain already shipped a different path, document it; new work should prefer `/batch`.                                                                                                                                                                                                                                                                                                                                                           |
| **Go handler**               | Use **`httpx.HandleBatch`**; return item count = number of billable units processed in the success path (typically length of input slice after dedupe rules are applied—document per API).                                                                                                                                                                                                                                                                                                                                                              |
| **Request body limit**       | Same as other JSON handlers today: **1 MiB** max body (`http.MaxBytesReader`). Keep arrays small enough to parse safely.                                                                                                                                                                                                                                                                                                                                                                                                                                |
| **Batch size**               | Enforce with `validate` struct tags (`min`, `max`, `dive`). Choose limits with latency and upstream rate limits in mind; **document** the max in public API docs.                                                                                                                                                                                                                                                                                                                                                                                       |
| **Response type**            | Any struct; include a **`total`** field if it helps clients when partial results are impossible—optional but phone does it.                                                                                                                                                                                                                                                                                                                                                                                                                             |
| **Partial failure**          | **Prefer partial success over total failure.** When one item encounters an infrastructure or validation error, absorb it in-band (e.g. `valid: false`, `found: false`, or an `error` field on that item) and continue processing the rest. Return **200** with the full result set. Only propagate a top-level error for systemic failures that affect _all_ items (e.g. context cancellation, unrecoverable infra outage). “All or nothing” semantics require a separate RFC discussion and explicit error codes—do not mix models without versioning. |
| **OpenAPI / dashboard docs** | Any new route must be added to dashboard API YAML so **`openapi.json`** on the gateway stays accurate.                                                                                                                                                                                                                                                                                                                                                                                                                                                  |

---

## Service-layer implementation rules

The transport layer (`httpx.HandleBatch`) handles HTTP concerns. The service
method it calls must itself be a real batch — not a loop over the single-item
method.

### The anti-pattern (do not do this)

```go
// WRONG: N serial round-trips disguised as a batch.
func (s *Service) LookupBatch(ctx context.Context, domains []string) []Result {
    results := make([]Result, len(domains))
    for i, d := range domains {
        results[i], _ = s.Lookup(ctx, d) // one DNS/HTTP/DB call per iteration
    }
    return results
}
```

This compiles, tests green, and silently gives callers no throughput benefit
over N individual API calls. A 50-item batch takes 50× the latency of one call.

### Rule 1 — Database: write a batch query, do not call the single-item method

If the single-item method runs a SQL query, the batch version must issue **one
query** that fetches all items, not N queries in a loop.

```go
// WRONG — N queries
for _, id := range ids {
    row := s.db.QueryRow(ctx, `SELECT ... FROM t WHERE id = $1`, id)
    ...
}

// RIGHT — 1 query using ANY / IN / LIMIT n (pick the form that fits the schema)
rows, err := s.db.Query(ctx, `SELECT ... FROM t WHERE id = ANY($1::int[])`, ids)

// For "N random rows" patterns:
rows, err := s.db.Query(ctx, `SELECT ... FROM t ORDER BY random() LIMIT $1`, n)
```

Reusing the single-item method for the batch case is only acceptable when the
single-item method performs **no I/O** (pure in-memory computation).

### Rule 2 — Network I/O (DNS, HTTP): fan out with goroutines + semaphore

When each item requires an independent network call (DNS lookup, outbound HTTP),
fan out concurrently. Cap concurrency with a semaphore to avoid exhausting the
upstream or the connection pool.

```go
const maxWorkers = 10

results := make([]BatchItem, len(items))
sem := make(chan struct{}, maxWorkers)
var wg sync.WaitGroup

for i, item := range items {
    wg.Add(1)
    sem <- struct{}{}
    go func(i int, item string) {
        defer wg.Done()
        defer func() { <-sem }()

        itemCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
        defer cancel()

        r, err := s.singleItemCall(itemCtx, item)
        if err != nil {
            results[i] = BatchItem{Found: false, Error: err.Error()}
            return
        }
        results[i] = BatchItem{Found: true, Data: r}
    }(i, item)
}

wg.Wait()
```

Results are written to pre-allocated index slots, so input order is preserved
without sorting. See `networking/whois/service.go` and
`validation/email/service.go` for the canonical in-codebase reference.

### Rule 3 — Error handling: in-band, never fail-fast

Consistent with the **Partial failure** standard above, goroutine-based batch
methods must absorb per-item errors in-band (`Found: false` / `Valid: false` /
`error` field on that item) and continue. Do not use `errgroup` with its
fail-fast semantics unless the RFC for that endpoint explicitly documents
all-or-nothing behaviour.

### Summary table

| Underlying operation        | Correct batch strategy                       |
| --------------------------- | -------------------------------------------- |
| DB query (rows by key / ID) | Single `WHERE id = ANY($1)` or equivalent    |
| DB query (N random rows)    | Single `ORDER BY random() LIMIT $1`          |
| DNS lookup                  | Goroutines + semaphore (`maxWorkers ≈ 10`)   |
| Outbound HTTP               | Goroutines + semaphore (`maxWorkers ≈ 10`)   |
| Pure in-memory computation  | Sequential loop is fine — no I/O, no latency |

---

## Checklist before merging a new batch endpoint

- [ ] Uses `httpx.HandleBatch`
- [ ] Service method is a real batch — no sequential loop over the single-item
      method for DB or network I/O (see **Service-layer implementation rules**).
- [ ] Struct validation documents **min/max** batch size; tests cover boundary
      and oversize.
- [ ] Public docs include curl example and error cases.
- [ ] Confirmed with the Go API owner that **usage** in PostgreSQL matches
      expectations for a sample batch (e.g. 10 items → 10 units).
