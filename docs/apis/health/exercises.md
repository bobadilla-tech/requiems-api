# Fitness Exercises API

Browse 1,500+ exercises with step-by-step instructions, target muscles,
equipment requirements, and body part filters.

## Endpoints

| Method | Path                          | Description                           |
| ------ | ----------------------------- | ------------------------------------- |
| GET    | `/v1/health/exercises`        | Paginated list with optional filters  |
| GET    | `/v1/health/exercises/{id}`   | Single exercise by ID                 |
| GET    | `/v1/health/exercises/random` | Random exercise with optional filters |
| POST   | `/v1/health/exercises/batch`  | Fetch up to 50 exercises by ID        |
| GET    | `/v1/health/body-parts`       | All valid body part values            |
| GET    | `/v1/health/equipment`        | All valid equipment values            |
| GET    | `/v1/health/muscles`          | All valid muscle values               |

## Listing and Filtering

`GET /v1/health/exercises`

All parameters are optional and combinable:

| Parameter   | Type    | Description                           |
| ----------- | ------- | ------------------------------------- |
| `body_part` | string  | Filter to a specific body part        |
| `equipment` | string  | Filter to a specific equipment type   |
| `muscle`    | string  | Filter by target or secondary muscle  |
| `search`    | string  | Full-text search on exercise name     |
| `page`      | integer | Page number (default: 1)              |
| `per_page`  | integer | Results per page, 1–100 (default: 20) |

Filter values are case-sensitive and must exactly match the strings returned by
the metadata endpoints. Use `body_part=upper legs` not `body_part=Upper Legs`.

### Pagination

The response includes `total`, `page`, and `per_page` fields. To iterate all
results:

```python
page = 1
while True:
    resp = requests.get(url, params={"page": page, "per_page": 100}, headers=headers)
    data = resp.json()["data"]
    process(data["items"])
    if page * data["per_page"] >= data["total"]:
        break
    page += 1
```

## Random Exercise

`GET /v1/health/exercises/random`

Accepts the same filter parameters as the list endpoint. Returns 404 when no
exercises match the given filters.

```bash
# Random bodyweight back exercise
curl "https://api.requiems.xyz/v1/health/exercises/random?body_part=back&equipment=body+weight" \
  -H "requiems-api-key: YOUR_API_KEY"
```

## Batch Lookup

`POST /v1/health/exercises/batch`

Fetch up to 50 exercises in a single request. Results are returned in the same order as the input IDs. IDs that do not exist are silently skipped.

```bash
curl -X POST "https://api.requiems.xyz/v1/health/exercises/batch" \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"ids": [1, 7, 42]}'
```

**Request body:**

| Field | Type             | Required | Constraints                        |
| ----- | ---------------- | -------- | ---------------------------------- |
| `ids` | array of integer | yes      | 1–50 items, each must be ≥ 1       |

**Response:**

```json
{
  "data": {
    "results": [
      {
        "id": 1,
        "name": "Barbell Squat",
        "body_parts": ["upper legs"],
        "equipment": ["barbell"],
        "target_muscles": ["quadriceps"],
        "secondary_muscles": ["glutes", "hamstrings"],
        "instructions": ["Stand with feet shoulder-width apart..."]
      }
    ],
    "total": 1
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

> **Note:** IDs 7 and 42 were not found and were silently skipped. Only existing IDs appear in `results`.

**Error codes:**

| Code                | Status | When                                             |
| ------------------- | ------ | ------------------------------------------------ |
| `bad_request`       | 400    | Missing or malformed JSON body                   |
| `validation_failed` | 422    | Empty list, more than 50 IDs, or any ID below 1  |

## Metadata Endpoints

Use these to discover all valid filter values before building filter UIs or
validating user input:

- `GET /v1/health/body-parts` — ~10 values (chest, back, upper legs, …)
- `GET /v1/health/equipment` — ~28 values (barbell, dumbbell, cable, …)
- `GET /v1/health/muscles` — ~51 values (biceps, glutes, pectorals, …)

All return `{ items: string[], total: number }`.

## Response Envelope

All responses follow the standard envelope:

```json
{
  "data": { ... },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

## Error Codes

| Code             | Status | When                                                                |
| ---------------- | ------ | ------------------------------------------------------------------- |
| `bad_request`    | 400    | Invalid query parameter (e.g. `per_page=0`) or non-numeric ID       |
| `not_found`      | 404    | No exercise with the given ID, or no exercises match random filters |
| `internal_error` | 500    | Unexpected server failure                                           |

## Data Notes

- No media (GIFs, images) is served. Only text data.
- Exercise IDs are stable database integers — safe to store and reference.
- Instructions have the `Step:N` prefix stripped; they are plain sentences in
  order.
- The `muscle` filter matches both `target_muscles` and `secondary_muscles`.

## Seeding

The exercise data is not stored in git. To populate the database:

```bash
# Dry run (no writes)
docker exec requiem-dev-api-1 go run ./cmd/seed-exercises \
  --data-dir /path/to/json --dry-run

# Production seed
docker exec requiem-dev-api-1 go run ./cmd/seed-exercises \
  --data-dir /path/to/json \
  --db-url "postgres://requiem:requiem@db:5432/requiem"
```

The seeder expects an `exercises.json` file in `--data-dir`. It is idempotent —
re-running updates existing records without changing their IDs.
