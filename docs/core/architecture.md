# Architecture

## Overview

```
┌──────────────────────────────────────────────────────────────────┐
│                           INTERNET                               │
└──────────────────────────────────────────────────────────────────┘
           │                                │
           │ API Requests                   │ Web Users
           ▼                                ▼
┌──────────────────────────────┐  ┌──────────────────────────────┐
│   api.requiems.xyz           │  │   requiems.xyz               │
│   (Cloudflare Worker)        │  │   (Rails Dashboard)          │
│                              │  │                              │
│  ┌────────┐  ┌────────┐      │  │  ┌────────┐  ┌─────────┐     │
│  │Auth KV │  │Usage   │      │  │  │Landing │  │Dashboard│     │
│  │        │  │  D1    │      │  │  │ Page   │  │ & Admin │     │
│  └────────┘  └────────┘      │  │  └────────┘  └─────────┘     │
│                              │  │                              │
│  requiems-api-key validation │  │  User management             │
│  Rate limiting               │  │  API key creation            │
│  Request tracking            │  │  Usage stats                 │
└──────────────────────────────┘  └──────────────────────────────┘
           │                                │
           │ X-Backend-Secret               │ DB queries
           ▼                                ▼
┌───────────────────────────────────────────────────────────────────┐
│              internal.requiems.xyz (Go Backend)                   │
│                      (Hetzner VPS)                                │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐      │
│  │              PostgreSQL (Shared Database)               │      │
│  │                                                         │      │
│  │  Go Tables:                                             │      │
│  │    - advice, quotes, words (business data)              │      │
│  │                                                         │      │
│  │  Rails Tables:                                          │      │
│  │    - users, api_keys, subscriptions                     │      │
│  │    - usage_logs, daily_usage_summaries                  │      │
│  │    - credit_adjustments, audit_logs, abuse_reports      │      │
│  └─────────────────────────────────────────────────────────┘      │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐      │
│  │                      Redis                              │      │
│  │   Sidekiq background jobs · Go API counter storage      │      │
│  └─────────────────────────────────────────────────────────┘      │
└───────────────────────────────────────────────────────────────────┘
```

## Components

### 1. Cloudflare Worker (api.requiems.xyz)

**Purpose:** Public API gateway with global edge distribution

**Responsibilities:**

- API key validation (Cloudflare KV)
- Rate limiting (KV counters)
- Request tracking (D1 SQLite)
- Request forwarding to internal backend

**Technology:** TypeScript on Cloudflare Workers

### 2. Rails Dashboard (requiems.xyz)

**Purpose:** Public-facing web application

**Responsibilities:**

- Landing page and marketing
- User registration and authentication
- User dashboard (`/dashboard/*`)
  - API key management
  - Usage statistics
  - Billing/subscription management
- Admin panel (`/admin/*`)
  - User management
  - System monitoring
  - Revenue tracking
  - Abuse detection
- API key sync to Cloudflare KV

**Technology:** Rails, Tailwind CSS, Hotwire (Turbo + Stimulus)

### 3. Go Backend (internal.requiems.xyz)

**Purpose:** Internal business logic API

**Responsibilities:**

- Execute business logic for all API endpoints
- Database queries for business data
- Only accessible with X-Backend-Secret header

**Technology:** Go, Chi router

### 4. Shared PostgreSQL

**Purpose:** Single source of truth for all data

**Schema Separation:**

- **Go migrations:** Business data tables
- **Rails migrations:** User/account data tables
- Separate migration tracking tables to avoid conflicts

### 5. Redis

**Purpose:** Background job queue for Rails and real-time counter store for the
Go API

**Used for:**

- Go API atomic counter increments (`INCR counter:{namespace}`)
- Counter persistence: background sync worker flushes Redis values to PostgreSQL
  every 60 seconds
- Sidekiq background jobs
- Usage sync from Cloudflare D1
- Email sending
- Scheduled tasks

## Data Stores Explained

We use **3 data stores**, each optimized for its specific use case:

### 1. Cloudflare KV (Key-Value)

**What:** Globally distributed key-value store with <10ms reads worldwide.

**Used for:**

- **API Key lookup** (`key:{api_key}` → user data, plan, etc.)
- **Rate limiting counters** (`rl:m:{key}:{minute}`)

**Why KV:**

- Extremely fast reads (critical for every API request)
- TTL expiration (rate limit keys auto-delete)
- Simple key-value, no complex queries needed
- Data is small (just key → JSON blob)

**Example data:**

```json
// key:requiem_abc123xyz
{
  "userId": "user_456",
  "plan": "starter",
  "createdAt": "2024-01-15T00:00:00Z",
  "billingCycleStart": "2024-12-01T00:00:00Z"
}
```

### 2. Cloudflare D1 (SQLite at Edge)

**What:** SQLite database running on Cloudflare's edge network.

**Used for:**

- **Request usage tracking** (INSERT on every API call, SUM queries for totals)
- **Usage analytics** (which endpoints, when, how much)

**Why D1 (not KV):**

- Need SQL aggregations (`SUM`, `GROUP BY`, date ranges)
- Need to query "usage since X date" for billing periods
- Historical data for analytics/billing
- KV can't do `WHERE used_at >= '2024-12-01'`

**Schema:**

```sql
CREATE TABLE credit_usage (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  api_key TEXT NOT NULL,
  endpoint TEXT NOT NULL,
  credits_used INTEGER NOT NULL,
  used_at TEXT NOT NULL
);
```

**Example query:**

```sql
-- Get this month's usage for billing
SELECT SUM(credits_used) FROM credit_usage
WHERE api_key = ? AND used_at >= '2024-12-01T00:00:00Z'
```

### 3. PostgreSQL (Backend Database)

**What:** Traditional relational database running alongside the Go backend.

**Used for:**

- **All business data** (advice, quotes, words, etc.)
- **User accounts**
- **Anything the API actually returns**

**Why PostgreSQL:**

- Complex queries, joins, full-text search
- Larger datasets (millions of rows)
- Backend runs on dedicated server, not edge
- D1 is edge-only, can't be accessed from Go backend

**Example tables:**

```sql
-- Business data in PostgreSQL
CREATE TABLE advice (id SERIAL, content TEXT, category TEXT, ...);
CREATE TABLE quotes (id SERIAL, text TEXT, author TEXT, ...);
CREATE TABLE words (id SERIAL, word TEXT, definition TEXT, ...);
```

## Why This Split?

| Store          | Location                | Latency | Use Case                |
| -------------- | ----------------------- | ------- | ----------------------- |
| **KV**         | Edge (global)           | <10ms   | Auth, rate limits       |
| **D1**         | Edge (global)           | ~50ms   | Usage tracking, billing |
| **PostgreSQL** | Backend (single region) | ~100ms  | Business data           |

**The Gateway (Worker) handles:**

- Auth + rate limits (KV) - must be ultra-fast
- Request tracking (D1) - needs SQL, still at edge

**The Backend (Go) handles:**

- Business logic - complex queries on PostgreSQL
- No auth overhead - trusts the gateway

## Request Flow

```
1. User calls: GET api.requiems.xyz/v1/text/advice
   └─ Header: requiems-api-key: requiem_abc123

2. Worker receives request
   └─ KV.get("key:requiem_abc123") → { plan: "starter", ... }

3. Check rate limit
   └─ KV.get("rl:m:requiem_abc123:28377600") → "150"
   └─ Under limit? Continue. Over? Return 429.

4. Check usage quota
   └─ D1: SELECT SUM(credits_used) WHERE api_key = ? AND used_at >= ?
   └─ Under limit? Continue. Over? Return 429 (free) or allow (paid).

5. Forward to backend
   └─ fetch(BACKEND_URL + "/v1/text/advice", { headers: { X-Backend-Secret } })

6. Backend processes
   └─ PostgreSQL: SELECT * FROM advice ORDER BY RANDOM() LIMIT 1
   └─ Returns JSON response

7. Record usage
   └─ D1: INSERT INTO credit_usage (api_key, endpoint, credits_used, ...)

8. Return response with headers
   └─ X-Requests-Used: 1
   └─ X-Requests-Remaining: 29849
   └─ X-RateLimit-Remaining: 149
```

## Security Model

1. **API Keys** stored in KV, looked up on every request
2. **Rate limits** enforced at edge before backend is touched
3. **Backend Secret** header required
4. **Backend trusts gateway** - no redundant auth checks
