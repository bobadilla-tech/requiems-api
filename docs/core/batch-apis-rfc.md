# RFC: Batch APIs (request shape, billing, and Go handler pattern)

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

**1. Phone validation** — `POST /v1/validation/phone/batch`

**Go wiring:** `httpx.HandleBatch` in the validation phone router. The handler
returns `(response, itemCount, error)`; the platform sets response header
**`X-Usage-Count`** to `itemCount` (stringified integer).

**2. Disposable email batch** — `POST /v1/networking/disposable/batch`

**Go wiring:** `httpx.HandleBatch` in the networking disposable router. Item
count is **`len(emails)`** after validation (1–100), matching one billable unit
per email checked. Public details: disposable email API docs.

**Request (illustrative):**

- JSON object with an array field (here `numbers`).
- Validation tags on the struct enforce **min/max length** of the array (phone:
  1–50 items).

**Response:**

- Same **`{"data": …, "metadata": …}`** success envelope as all `httpx.JSON` /
  `Handle` / `HandleBatch` responses (`data` must implement `IsData()`).

**Errors:**

- Malformed input or validation failure → **`422`** with
  `error: "validation_failed"` and `fields` (standard `httpx` behavior).
- Domain errors from the service → map with `httpx.AppError` as today.

**Ordering:** Results MUST stay in the **same order** as the input array unless
the RFC for that specific API documents otherwise.

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

| Topic                        | Standard                                                                                                                                                                                                                                                                          |
| ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **HTTP method**              | `POST`                                                                                                                                                                                                                                                                            |
| **Path naming**              | Prefer **`/resource/batch`** under the same prefix as single-item routes (e.g. `/phone/batch`). If the domain already shipped a different path, document it; new work should prefer `/batch`.                                                                                     |
| **Go handler**               | Use **`httpx.HandleBatch`**; return item count = number of billable units processed in the success path (typically length of input slice after dedupe rules are applied—document per API).                                                                                        |
| **Request body limit**       | Same as other JSON handlers today: **1 MiB** max body (`http.MaxBytesReader`). Keep arrays small enough to parse safely.                                                                                                                                                          |
| **Batch size**               | Enforce with `validate` struct tags (`min`, `max`, `dive`). Choose limits with latency and upstream rate limits in mind; **document** the max in public API docs.                                                                                                                 |
| **Response type**            | A struct implementing **`IsData()`**; include a **`total`** field if it helps clients when partial results are impossible—optional but phone does it.                                                                                                                             |
| **Partial failure**          | **Phone model:** always **200** with per-item results (each item may indicate invalid / error state in band). If a future API needs “all or nothing” semantics, that must be a **separate** RFC section and explicit error codes—do not mix models without versioning discussion. |
| **OpenAPI / dashboard docs** | Any new route must be added to dashboard API YAML so **`openapi.json`** on the gateway stays accurate.                                                                                                                                                                            |

---

## Checklist before merging a new batch endpoint

- [ ] Uses `httpx.HandleBatch` and returns correct **item count**.
- [ ] Struct validation documents **min/max** batch size; tests cover boundary
      and oversize.
- [ ] Public docs include curl example and error cases.
- [ ] Confirmed with gateway owner that **usage** in D1/KV matches expectations
      for a sample batch (e.g. 10 items → 10 units).
