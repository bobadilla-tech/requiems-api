# Adding a New Endpoint to the Go Backend

This guide walks through every step required to ship a new endpoint: from
writing the Go code to updating the dashboard catalog and API documentation.
Follow it in order — the checklist at the end maps to each section.

## Architecture Refresher

Before writing any code, understand where the Go backend sits in the request
flow:

```
User → Auth Gateway (api.requiems.xyz)
         ↓  validates API key, enforces rate limits
       Go Backend (apps/api, port 8080)
         ↓  business logic, database queries
       PostgreSQL
```

### Domain Directory Layout

Every feature lives inside a domain:

```
apps/api/
├── app/
│   ├── app.go            # Bootstrap: DB, Redis, middleware, router
│   └── routes_v1.go      # Mounts all /v1 domain routers
├── platform/
│   ├── httpx/            # JSON, Error, Handle, BindQuery helpers
│   ├── svcerr/           # Transport-agnostic service errors (use in services)
│   ├── config/           # Env config
│   ├── db/               # PostgreSQL connection + migrations
│   └── middleware/       # BackendSecretAuth
└── services/
    └── {domain}/
        ├── router.go         # Wires services → chi.Router for this domain
        └── {feature}/
            ├── type.go           # Request + response types
            ├── service.go        # Business logic
            └── transport_http.go # HTTP handlers
```

Existing top-level domains and their `/v1` prefixes:

| Domain folder             | URL prefix          |
| ------------------------- | ------------------- |
| `services/entertainment/` | `/v1/entertainment` |
| `services/finance/`       | `/v1/finance`       |
| `services/health/`        | `/v1/health`        |
| `services/networking/`    | `/v1/networking`    |
| `services/places/`        | `/v1/places`        |
| `services/technology/`    | `/v1/technology`    |
| `services/text/`          | `/v1/text`          |
| `services/validation/`    | `/v1/validation`    |

---

## Before Writing Any Code — Check for Existing Libraries

**Do not write a service from scratch if a battle-tested library already solves
the problem.**

Go has a rich ecosystem and several `bobadilla-tech` packages already in
`go.mod` that were purpose-built for this platform. Writing your own
implementation of something that already exists adds maintenance burden,
reintroduces bugs that libraries have already fixed, and makes the codebase
harder to onboard.

### Check in this order

1. **`go.mod` — existing dependencies first** Before anything else, open
   `apps/api/go.mod` and scan the current dependency list. The problem may
   already be solved.

2. **`bobadilla-tech` org on GitHub** Check whether a new `bobadilla-tech`
   package exists for the domain you are building. These are first-party
   libraries designed to slot directly into this platform.

3. **Well-known Go ecosystem packages** For common problems, prefer established
   packages over custom implementations:
   - Parsing / tokenising: `golang.org/x/text`, `mvdan.cc/xurls`, etc.
   - Cryptography: standard library `crypto/*` — never roll your own
   - Time zones / date math: standard library `time` + `golang.org/x/time`
   - HTTP client retries: `hashicorp/go-retryablehttp`
   - UUID generation: `google/uuid`

4. **Only write it yourself if:**
   - No library exists for the domain
   - Available libraries are unmaintained, have known CVEs, or have prohibitive
     licenses
   - The logic is so simple (< 20 lines, no edge cases) that a dependency would
     be overkill
   - You have discussed it with the team and documented the decision

### When you add a new library

```bash
docker exec requiem-dev-api-1 go get github.com/some/library@latest
docker exec requiem-dev-api-1 go mod tidy
```

Commit both `go.mod` and `go.sum`. Never edit them by hand.

---

## Step 1 — Write the Go Code

Create four files in this order (each builds on the previous).

### 1a. `type.go` — Define Your Types

> **Rule: always use `snake_case` for JSON field names.** Every `json:"..."` tag
> in this codebase uses lower_snake_case. Never use camelCase or PascalCase. ✅
> `json:"has_profanity"` `json:"flagged_words"` `json:"browser_version"` ❌
> `json:"hasProfanity"` `json:"flaggedWords"` `json:"browserVersion"`

```go
package riddle

// Request for POST endpoints with JSON body.
// Use validate tags for automatic validation.
type GenerateRequest struct {
    Category string `json:"category" validate:"required,oneof=general science history"`
}

// Response types.
type Riddle struct {
    ID       int    `json:"id"`
    Question string `json:"question"`
    Answer   string `json:"answer"`
    Category string `json:"category"`
}

// For collections or richer responses:
type RiddleList struct {
    Items []Riddle `json:"items"`
    Total int      `json:"total"`
}
```

**Validation tag reference** (from `go-playground/validator`):

| Tag             | Meaning                                |
| --------------- | -------------------------------------- |
| `required`      | Field must be present and non-zero     |
| `email`         | Must be a valid email address          |
| `oneof=a b c`   | Must be one of the listed values       |
| `min=1,max=100` | Numeric range (or string/slice length) |
| `dive,email`    | Validate each element of a slice       |
| `url`           | Must be a valid URL                    |

> **Rule: always declare validation in struct tags — never write it by hand in
> handlers.** `validate` tags work on both `json` (POST body via `httpx.Handle`)
> and `query` (GET params via `httpx.BindQuery`) structs. The framework enforces
> them automatically and returns consistent error responses.
>
> ✅ Correct — validation declared in the struct:
>
> ```go
> type ParseRequest struct {
>     UA string `query:"ua" validate:"required"`
> }
> ```
>
> ❌ Wrong — manual inline check that bypasses the validation pipeline:
>
> ```go
> ua := r.URL.Query().Get("ua")
> if ua == "" {
>     httpx.Error(w, http.StatusBadRequest, "bad_request", "ua parameter is required")
>     return
> }
> ```

### 1b. `service.go` — Business Logic

Services receive dependencies through their constructor. Keep this layer free of
HTTP concerns.

**Service Layer Rules — read these before writing a single line:**

1. **Never import `net/http` or `requiems-api/platform/httpx` in a service.**
   Services must not know about HTTP status codes. All error information is
   communicated through `requiems-api/platform/svcerr`.

2. **Return `*svcerr.Error` for expected failure cases** (not found, invalid
   input, upstream unavailable). Return a plain `error` only for unexpected
   infrastructure failures that should become `500`.

3. **Never silently clamp or default invalid input.** If `difficulty` is not one
   of `easy|medium|hard`, return `svcerr.Invalid(...)`. If `page > maxPages`,
   return `svcerr.NotFound(...)`. Silent defaults hide bugs in callers.

**`svcerr` constructor reference:**

| Constructor                  | HTTP status | Use when                                   |
| ---------------------------- | ----------- | ------------------------------------------ |
| `svcerr.NotFound(code, msg)` | 404         | Resource does not exist                    |
| `svcerr.Invalid(code, msg)`  | 400         | Bad input the caller can fix               |
| `svcerr.Unknown(code, msg)`  | 422         | Unprocessable but structurally valid input |
| `svcerr.Upstream(code, msg)` | 503         | Third-party service / DB unavailable       |

`httpx.Handle`, `httpx.HandleGet`, and `httpx.HandleBatch` map `*svcerr.Error`
to the correct HTTP response automatically. `KindUpstream` errors are also
forwarded to Sentry.

```go
package riddle

import (
    "context"
    "errors"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"

    "requiems-api/platform/svcerr"
)

type Service struct {
    db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
    return &Service{db: db}
}

func (s *Service) GetByID(ctx context.Context, id int) (Riddle, error) {
    row := s.db.QueryRow(ctx, `
        SELECT id, question, answer, category
        FROM riddles WHERE id = $1`, id)

    var r Riddle
    if err := row.Scan(&r.ID, &r.Question, &r.Answer, &r.Category); err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return Riddle{}, svcerr.NotFound("not_found", "riddle not found")
        }
        return Riddle{}, err // unexpected DB failure → 500
    }
    return r, nil
}
```

For services with **no database dependency** (e.g., in-memory computation,
embedded data):

```go
type Service struct{}

func NewService() *Service { return &Service{} }
```

### 1c. `transport_http.go` — HTTP Handlers

Choose the right pattern based on the endpoint's input method.

#### Pattern A: `httpx.Handle` — POST with JSON body (recommended for mutations)

Use this when the request has a JSON body. The wrapper handles decoding,
validation, and error mapping automatically. The service just returns errors —
the handler does not need to inspect them.

```go
package riddle

import (
    "context"

    "github.com/go-chi/chi/v5"
    "requiems-api/platform/httpx"
)

func RegisterRoutes(r chi.Router, svc *Service) {
    r.Post("/riddle/generate", httpx.Handle(
        func(ctx context.Context, req GenerateRequest) (Riddle, error) {
            return svc.Random(ctx, req.Category)
        },
    ))
}
```

`httpx.Handle` automatically:

- Caps the request body at 1 MB
- Decodes JSON (strict mode — unknown fields are rejected)
- Runs struct validation; returns `422 Unprocessable Entity` with field-level
  errors on failure
- Unwraps `*svcerr.Error` and maps it to the appropriate HTTP status
- Captures `KindUpstream` errors in Sentry before responding
- Returns `500 Internal Server Error` for any other unhandled error

---

#### Pattern B: `httpx.HandleGet` — GET with query parameters

Use this when input comes from query string parameters. Pass an initialised
`Req` value as the second argument to set defaults — `BindQuery` skips missing
params, so the defaults remain in place. Always declare required fields and
constraints using `validate` tags on the struct — do not add manual checks in
the handler body.

```go
package riddle

import (
    "context"

    "github.com/go-chi/chi/v5"
    "requiems-api/platform/httpx"
)

func RegisterRoutes(r chi.Router, svc *Service) {
    r.Get("/riddles", httpx.HandleGet(
        func(ctx context.Context, req ListRequest) (RiddleList, error) {
            return svc.List(ctx, req.Page, req.PerPage)
        },
        ListRequest{Page: 1, PerPage: 20}, // defaults applied before binding
    ))
}
```

`httpx.HandleGet` automatically:

- Binds query parameters to the struct (`query:` tags)
- Runs struct validation; returns `422 Unprocessable Entity` with field-level
  errors on failure (parse errors return `400`)
- Unwraps `*svcerr.Error` and maps it to the appropriate HTTP status
- Captures `KindUpstream` errors in Sentry before responding
- Returns `500 Internal Server Error` for any other unhandled error

> **Keep inline handlers** (`r.Get(...)` with manual `BindQuery`) only when
> `HandleGet` cannot apply: path parameters via `chi.URLParam`, pre-bind URL
> mutation (e.g. uppercasing a country code before binding), or non-JSON
> responses (e.g. raw PNG output).

---

#### Pattern C: `chi.URLParam` — GET with URL path parameters

Use this when the identifier is part of the path (e.g., `/riddles/{id}`).

```go
func RegisterRoutes(r chi.Router, svc *Service) {
    r.Get("/riddles/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")
        if id == "" {
            httpx.Error(w, http.StatusBadRequest, "bad_request", "id is required")
            return
        }

        riddle, err := svc.GetByID(r.Context(), id)
        if err != nil {
            if se, ok := errors.AsType[*svcerr.Error](err); ok {
                httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
                return
            }
            httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
            return
        }

        httpx.JSON(w, http.StatusOK, riddle)
    })
}
```

> The service returns `svcerr.NotFound(...)` when the riddle does not exist. The
> handler propagates that decision rather than assuming a 404 — if the service
> later returns a 422 for a malformed id, the handler still does the right thing
> without any changes.

---

**Error response reference:**

| Situation                    | Status | Code (snake_case)    | Source                 |
| ---------------------------- | ------ | -------------------- | ---------------------- |
| Missing/invalid JSON body    | 400    | `bad_request`        | `httpx` (automatic)    |
| Failed struct validation     | 422    | `validation_failed`  | `httpx` (automatic)    |
| Bad input the caller can fix | 400    | anything descriptive | `svcerr.Invalid(...)`  |
| Resource not found           | 404    | anything descriptive | `svcerr.NotFound(...)` |
| Valid input, unprocessable   | 422    | anything descriptive | `svcerr.Unknown(...)`  |
| Third-party / DB unavailable | 503    | `upstream_error`     | `svcerr.Upstream(...)` |
| Unexpected failure           | 500    | `internal_error`     | plain `error` from svc |

`svcerr.Upstream` errors are captured in Sentry automatically by `httpx.Handle`,
`httpx.HandleGet`, and `httpx.HandleBatch`. `NotFound`, `Invalid`, and `Unknown`
are not — they are expected outcomes, not bugs.

### 1d. `router.go` — Wire the Domain

This file instantiates the service and calls the feature's `RegisterRoutes`.

```go
package text  // parent domain package

import (
    "github.com/go-chi/chi/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "requiems-api/services/text/riddle"
)

func RegisterRoutes(r chi.Router, pool *pgxpool.Pool) {
    // ... existing features
    svc := riddle.NewService(pool)
    riddle.RegisterRoutes(r, svc)
}
```

---

## Step 2 — Mount the Router

**Adding to an existing domain** (most common): just add the lines above to the
existing `router.go` for that domain. No other files need changing.

**Creating a brand-new top-level domain**: add two lines to
`apps/api/app/routes_v1.go`:

```go
func registerV1Routes(ctx context.Context, r chi.Router, pool *pgxpool.Pool, rdb *redis.Client) {
    // ... existing mounts

    puzzlesRouter := chi.NewRouter()
    puzzles.RegisterRoutes(puzzlesRouter, pool)
    r.Mount("/puzzles", puzzlesRouter)
}
```

---

## Step 3 — Database Migrations (if needed)

If the feature needs new tables or columns, create a SQL migration file:

```
apps/api/migrations/0005_add_riddles_table.up.sql
apps/api/migrations/0005_add_riddles_table.down.sql
```

Use the next sequential 4-digit number after the last existing migration.

Example:

```sql
-- 0005_add_riddles_table.up.sql
CREATE TABLE riddles (
    id       SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    answer   TEXT NOT NULL,
    category VARCHAR(50) NOT NULL DEFAULT 'general'
);

-- 0005_add_riddles_table.down.sql
DROP TABLE riddles;
```

**No Go registration is needed.** The app calls `db.MigrateWithRetry()` on
startup and discovers all `*.up.sql` files in that directory automatically.

---

## Step 4 — Tests

Tests are **required before merge**. There are two layers to cover.

### Service Unit Tests (`service_test.go`)

Use table-driven tests. Keep each case small and named clearly.

```go
package riddle

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestService_Random(t *testing.T) {
    t.Parallel()
    // For services with no DB: svc := NewService()
    // For DB-backed services: use a test DB or interface/mock.

    tests := []struct {
        name     string
        category string
        wantErr  bool
    }{
        {name: "valid category", category: "general", wantErr: false},
        {name: "empty category", category: "", wantErr: true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()

            svc := NewService()
            result, err := svc.Random(tt.category)
            if tt.wantErr {
                require.Error(t, err)
            } else {
                require.NoError(t, err)
                assert.NotEmpty(t, result.Question)
            }
        })
    }
}
```

### HTTP Handler Tests (`transport_http_test.go`)

Test the full HTTP layer using `httptest`. Always include sad paths.

```go
package riddle

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/go-chi/chi/v5"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"

    "requiems-api/platform/httpx"
)

func setupRouter(svc *Service) chi.Router {
    r := chi.NewRouter()
    RegisterRoutes(r, svc)
    return r
}

func TestRiddle_Generate_HappyPath(t *testing.T) {
    t.Parallel()

    r := setupRouter(NewService())

    body := `{"category":"general"}`
    req := httptest.NewRequest(http.MethodPost, "/riddle/generate", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    require.Equal(t, http.StatusOK, w.Code)

    var resp httpx.Response[Riddle]
    err := json.NewDecoder(w.Body).Decode(&resp)
    require.NoError(t, err)
    assert.NotEmpty(t, resp.Data.Question)
}

func TestRiddle_Generate_MissingCategory(t *testing.T) {
    t.Parallel()

    r := setupRouter(NewService())
    req := httptest.NewRequest(http.MethodPost, "/riddle/generate", strings.NewReader(`{}`))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
```

**Run tests:**

```bash
# Domain-scoped (fast feedback during development)
docker exec requiem-dev-api-1 go test ./services/text/riddle/...

# Full suite (required before pushing)
docker exec requiem-dev-api-1 go test ./...

# With race detection and coverage (CI equivalent)
docker exec requiem-dev-api-1 go test -race -coverprofile=coverage.out ./...
```

---

## Step 5 — Update the Dashboard API Catalog

> Skip this step if you are adding a new endpoint to an **existing** API (the
> catalog entry already exists).

File: `apps/dashboard/config/api_catalog.yml`

Add an entry under the appropriate category:

```yaml
- id: riddles # MUST use hyphens (not underscores) and match YAML filename
  name: Riddles
  category: text
  description: Generate random riddles across categories like general, science, and history.
  endpoints_count: 2 # Note: it's endpoints_count, not endpoints
  status: live # or coming_soon
  popular: false
  documentation_url: /apis/riddles # MUST match the id
  tags:
    - riddles
    - trivia
    - fun
```

**CRITICAL**: The `id` field MUST:

- Use hyphens (not underscores): `random-word` not `random_word`
- Match the YAML documentation filename exactly
- Match the `documentation_url` path (after `/apis/`)

**Example of correct matching:**

- Catalog: `id: random-word`, `documentation_url: /apis/random-word`
- YAML file: `apps/dashboard/config/api_docs/random-word.yml`
- URL: `http://localhost:3000/apis/random-word`

---

## Step 6 — Add the API Documentation YAML

File: `apps/dashboard/config/api_docs/riddles.yml`

This YAML powers the interactive API documentation page in the dashboard. Every
field is rendered in the UI — do not leave any required section empty.

**CRITICAL RULES FOR YAML DOCUMENTATION:**

1. **File Naming**: Use hyphens (not underscores) and match the catalog ID
   exactly
   - ✅ `random-word.yml` matching catalog `id: random-word`
   - ❌ `random_word.yml` with catalog `id: random-word` (will not load!)

2. **Response Format**: ALL responses MUST include `data` and `metadata`
   wrappers

   ```json
   {
     "data": {
       // Your actual response fields here
     },
     "metadata": {
       "timestamp": "2026-01-01T00:00:00Z"
     }
   }
   ```

   This is enforced by `httpx.JSON` in the Go backend.

3. **Field Naming**: Always use `snake_case`, never `camelCase`
   - ✅ `is_disposable`, `part_of_speech`, `has_more`
   - ❌ `isDisposable`, `partOfSpeech`, `hasMore`

4. **Authentication**: Always use `requiems-api-key` header
   - ✅ `-H "requiems-api-key: YOUR_API_KEY"`
   - ❌ `-H "Authorization: Bearer YOUR_API_KEY"`

5. **Path Parameters**: For endpoints with URL path parameters (e.g.,
   `/counter/{namespace}`):

   ```yaml
   parameters:
     - name: namespace
       type: string
       required: true
       location: path # CRITICAL - enables input field in playground
       description: "Counter namespace identifier"
       example: page-views
   ```

6. **Query Parameters**: For GET endpoints with query strings:

   ```yaml
   parameters:
     - name: page
       type: integer
       required: false
       location: query
       description: "Page number (default: 1)"
       example: 1
   ```

7. **Body Parameters**: For POST/PUT JSON body parameters:

   ```yaml
   parameters:
     - name: email
       type: string
       required: true
       location: body # Can be omitted, defaults to body
       description: The email address to check
       example: test@example.com
   ```

8. **YAML Quoting**: Always quote strings containing colons to avoid parse
   errors
   - ✅ `description: "Page number (default: 1)"`
   - ❌ `description: Page number (default: 1)` (will cause YAML syntax error!)

9. **No Pricing Section**: Do NOT include pricing information (it's global, not
   per-API)
   - Pricing is displayed site-wide, not in individual API docs

10. **Hot Reload**: The development environment auto-reloads YAML changes - just
    refresh your browser

```yaml
api_id: riddles
api_name: Riddles
description: Generate random riddles across multiple categories for trivia apps, brain teasers, and educational content.
base_url: https://api.requiems.xyz

overview:
  use_cases:
    - Trivia apps and quiz games
    - Educational platforms
    - Ice-breaker bots
    - Daily challenge widgets

  features:
    - Multiple categories (general, science, history)
    - Deterministic question/answer pairs
    - Fast, low-latency responses

endpoints:
  - name: Generate Random Riddle
    method: POST
    path: /v1/text/riddle/generate
    description: Returns a random riddle from the specified category.

    parameters:
      - name: category
        type: string
        required: true
        location: body
        description: Category of riddle to return.
        example: general

    request_example: |
      {
        "category": "general"
      }

    response_example: |
      {
        "data": {
          "id": 42,
          "question": "What has keys but no locks?",
          "answer": "A keyboard",
          "category": "general"
        },
        "metadata": {
          "timestamp": "2026-01-01T00:00:00Z"
        }
      }

    response_fields:
      - name: id
        type: integer
        description: Unique riddle identifier
      - name: question
        type: string
        description: The riddle question
      - name: answer
        type: string
        description: The answer to the riddle
      - name: category
        type: string
        description: Category the riddle belongs to

    errors:
      - code: validation_failed
        status: 422
        description: The category field is missing or contains an invalid value.
      - code: internal_error
        status: 500
        description: Unexpected server error.

    code_examples:
      curl: |
        curl -X POST https://api.requiems.xyz/v1/text/riddle/generate \
          -H "requiems-api-key: YOUR_API_KEY" \
          -H "Content-Type: application/json" \
          -d '{"category": "general"}'

      python: |
        import requests

        url = "https://api.requiems.xyz/v1/text/riddle/generate"
        headers = {
            "requiems-api-key": "YOUR_API_KEY",
            "Content-Type": "application/json"
        }
        payload = {"category": "general"}

        response = requests.post(url, headers=headers, json=payload)
        print(response.json())

      javascript: |
        const response = await fetch('https://api.requiems.xyz/v1/text/riddle/generate', {
          method: 'POST',
          headers: {
            'requiems-api-key': 'YOUR_API_KEY',
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ category: 'general' })
        });

        const data = await response.json();
        console.log(data.data.question);

      ruby: |
        require 'net/http'
        require 'json'

        uri = URI('https://api.requiems.xyz/v1/text/riddle/generate')
        request = Net::HTTP::Post.new(uri)
        request['requiems-api-key'] = 'YOUR_API_KEY'
        request['Content-Type'] = 'application/json'
        request.body = { category: 'general' }.to_json

        response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
          http.request(request)
        end

        data = JSON.parse(response.body)
        puts data['data']['question']

faq:
  - question: Can I request riddles from multiple categories at once?
    answer: Not currently — each request returns one riddle from one category. Batch support is planned.

  - question: Are riddle IDs stable across requests?
    answer: Yes, IDs are stable database identifiers. You can use them to avoid showing a riddle twice.
```

Use `apps/dashboard/config/api_docs/disposable-email.yml` as the reference
template for complex multi-endpoint APIs and `advice.yml` for simple
single-endpoint APIs.

---

## Step 7 — Add Markdown API Docs

File: `docs/apis/{category}/{api-name}.md`

This is the developer-facing plain-text companion in the repo. It should explain
the _why_ and _edge cases_ that the YAML cannot express clearly.

```markdown
# Riddles API

Returns random riddles across multiple categories.

## Endpoint

`POST /v1/text/riddle/generate`

## Categories

| Value     | Description                        |
| --------- | ---------------------------------- |
| `general` | Everyday wordplay and logic        |
| `science` | Biology, physics, chemistry        |
| `history` | Historical facts framed as riddles |

## Response Envelope

All responses are wrapped in the standard envelope:

\`\`\`json { "data": { ... }, "metadata": { "timestamp": "2026-01-01T00:00:00Z"
} } \`\`\`

## Error Codes

| Code                | Status | When                        |
| ------------------- | ------ | --------------------------- |
| `validation_failed` | 422    | Invalid or missing category |
| `internal_error`    | 500    | Unexpected failure          |
```

---

## Step 8 — Update the Request Multiplier (if non-default)

File: `apps/workers/shared/src/config.ts`

Most endpoints cost **1 request** and do not need a config entry. Only add an
entry if your endpoint explicitly counts as more than 1 request (rare — document
this clearly in the endpoint's API doc):

```ts
// In ENDPOINT_MULTIPLIERS map:
["/v1/text/riddle/generate", 1],   // default — omit this line unless non-1
["/v1/text/summarize", 5],         // expensive AI call
["/v1/translate/text", 3],         // ML translation
```

The gateway uses `getRequestMultiplier(method, pathname)` to look up the
multiplier before tracking requests. If no entry exists, it defaults to `1`.

---

## Step 9 — Validate Documentation

Before testing, validate the YAML syntax to catch errors early:

```bash
# Validate YAML syntax
docker exec requiem-dev-dashboard-1 ruby -ryaml -e "YAML.load_file('config/api_docs/riddles.yml'); puts '✅ YAML is valid'"

# Check if all catalog IDs have matching YAML files
cd apps/dashboard
for id in $(grep "documentation_url:" config/api_catalog.yml | awk '{print $2}' | sed 's|/apis/||'); do
  if [ -f "config/api_docs/$id.yml" ]; then
    echo "✅ $id.yml exists"
  else
    echo "❌ $id.yml MISSING"
  fi
done
```

---

## Common Errors & Troubleshooting

### "API not found" when clicking the API in the dashboard

**Cause**: Mismatch between catalog `id` and the YAML filename or
`documentation_url`

**Fix**:

1. Check that catalog has `id: random-word` (with hyphens)
2. Check that YAML file is named `random-word.yml` (matching the id)
3. Check that `documentation_url: /apis/random-word` matches the id
4. Refresh the page (hot reload should pick up changes)

### "Documentation not available for this API yet"

**Cause**: YAML file is missing or has a syntax error

**Fix**:

1. Run the YAML validation command above
2. Check for common YAML errors:
   - Unquoted strings containing colons:
     `description: "Use quotes (like: this)"`
   - Incorrect indentation (use spaces, not tabs)
   - Missing required sections

### "Request failed: bad component(expected absolute path component)"

**Cause**: Path parameters are not filled in or `location: path` is not set

**Fix**:

1. Ensure path parameters have `location: path` in the YAML
2. Fill in all required path parameters before clicking "Send Request"
3. Example: For `/counter/{namespace}`, the user must input a namespace value

### "Mapping values are not allowed in this context" (YAML parse error)

**Cause**: Unquoted string containing a colon or other special characters

**Fix**: Quote any description or value containing `:`, `{`, `}`, or `#`

```yaml
# ❌ Wrong
description: Page number (default: 1)

  # ✅ Correct
  description: "Page number (default: 1)"
```

---

## Docker Considerations

**No Docker changes are needed** for adding Go endpoints. The dev container
(`requiem-dev-api-1`) runs with Air hot reload — new `.go` files are compiled
and reloaded automatically when saved.

If your feature adds a **new external service dependency** (e.g., a new
third-party HTTP client, a new Redis data structure), check
`infra/docker/docker-compose.dev.yml` for any environment variables or service
additions required, but this is uncommon.

---

## Pre-Merge Verification Checklist

Work through these in order before opening a PR.

```
Go code
  [ ] go test ./... passes in the container (zero failures)
  [ ] go test -race -coverprofile=coverage.out ./... passes (no races)
  [ ] golangci-lint run passes or only advisory warnings remain

Manual smoke test
  [ ] curl -X POST http://localhost:8080/v1/{domain}/{endpoint} \
        -H "X-Backend-Secret: your_local_secret" \
        -H "Content-Type: application/json" \
        -d '{...}' returns expected response
  [ ] Invalid input returns correct 4xx with descriptive error

Database (if applicable)
  [ ] Migration file follows naming convention: 000X_description.up.sql + matching .down.sql
  [ ] App starts cleanly after migration (no startup errors in docker logs api)

Documentation
  [ ] apps/dashboard/config/api_docs/{name}.yml created with all sections
  [ ] YAML filename uses hyphens (not underscores)
  [ ] Catalog id matches YAML filename exactly
  [ ] All responses include data/metadata wrappers
  [ ] All fields use snake_case (not camelCase)
  [ ] All code examples use requiems-api-key header
  [ ] Path parameters have location: path set
  [ ] Strings with colons are quoted
  [ ] NO pricing section included (pricing is global)
  [ ] YAML validation passes: docker exec requiem-dev-dashboard-1 ruby -ryaml -e "YAML.load_file('config/api_docs/{name}.yml'); puts 'Valid'"
  [ ] apps/dashboard/config/api_catalog.yml updated (new API only)
  [ ] docs/apis/{category}/{name}.md created
  [ ] Tested in playground at http://localhost:3000/apis/{name}

Workers
  [ ] apps/workers/shared/src/config.ts updated if request multiplier != 1
  [ ] pnpm run typecheck passes in apps/workers/shared/ (if config.ts was touched)
```

---

## Quick Reference: Documentation Parameter Types

When defining parameters in your YAML documentation:

**Parameter Locations:**

- `location: path` - Part of the URL path (e.g., `/counter/{namespace}`)
- `location: query` - Query string parameter (e.g., `?page=1&per_page=20`)
- `location: body` - JSON request body parameter (default, can be omitted)

**Parameter Types:**

- `string` - Text values
- `integer` - Whole numbers
- `number` - Decimal numbers
- `boolean` - true/false values
- `array` - JSON arrays (users input as JSON: `["item1", "item2"]`)
- `object` - JSON objects (users input as JSON: `{"key": "value"}`)

**Required Fields for Each Parameter:**

```yaml
parameters:
  - name: param_name # Parameter identifier
    type: string # One of the types above
    required: true # true or false
    location: query # path, query, or body
    description: "What it does" # Quote if contains colons
    example: example-value # Shown as placeholder in playground
```

---

## Worked Example: `GET /v1/text/riddle/random`

Below is a complete minimal implementation for a riddle endpoint backed by an
in-memory list (no database).

**`apps/api/services/text/riddle/type.go`**

```go
package riddle

type Riddle struct {
    Question string `json:"question"`
    Answer   string `json:"answer"`
}
```

**`apps/api/services/text/riddle/service.go`**

```go
package riddle

import "math/rand"

var riddles = []Riddle{
    {Question: "What has keys but no locks?", Answer: "A keyboard"},
    {Question: "What gets wetter the more it dries?", Answer: "A towel"},
}

type Service struct{}

func NewService() *Service { return &Service{} }

func (s *Service) Random() Riddle {
    return riddles[rand.Intn(len(riddles))]
}
```

**`apps/api/services/text/riddle/transport_http.go`**

```go
package riddle

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "requiems-api/platform/httpx"
)

func RegisterRoutes(r chi.Router, svc *Service) {
    r.Get("/riddle/random", func(w http.ResponseWriter, r *http.Request) {
        httpx.JSON(w, http.StatusOK, svc.Random())
    })
}
```

**`apps/api/services/text/riddle/transport_http_test.go`**

```go
package riddle

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/go-chi/chi/v5"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"

    "requiems-api/platform/httpx"
)

func newRouter() chi.Router {
    r := chi.NewRouter()
    RegisterRoutes(r, NewService())
    return r
}

func TestRiddle_Random(t *testing.T) {
    t.Parallel()

    req := httptest.NewRequest(http.MethodGet, "/riddle/random", http.NoBody)
    w := httptest.NewRecorder()
    newRouter().ServeHTTP(w, req)

    require.Equal(t, http.StatusOK, w.Code)

    var resp httpx.Response[Riddle]
    err := json.NewDecoder(w.Body).Decode(&resp)
    require.NoError(t, err)
    assert.NotEmpty(t, resp.Data.Question)
    assert.NotEmpty(t, resp.Data.Answer)
}
```

**Add to `apps/api/services/text/router.go`:**

```go
svc := riddle.NewService()
riddle.RegisterRoutes(r, svc)
```

That is everything needed to ship a working, tested, documented endpoint.
