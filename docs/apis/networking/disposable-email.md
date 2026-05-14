# Disposable Email Detection API

This API provides endpoints to detect disposable/temporary email addresses and
domains using the
[is-email-disposable](https://github.com/bobadilla-tech/is-email-disposable)
package.

## Base URL

All endpoints are mounted under `/v1/networking`

## Endpoints

### 1. Check Single Email

Check if a single email address is disposable.

**Endpoint:** `POST /v1/networking/disposable/check`

**Request Body:**

```json
{
  "email": "user@tempmail.com"
}
```

**Response:** `200 OK`

```json
{
  "data": {
    "email": "user@tempmail.com",
    "is_disposable": true,
    "domain": "tempmail.com"
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
# Production
curl -X POST https://api.requiems.xyz/v1/networking/disposable/check \
  -H "Content-Type: application/json" \
  -H "requiems-api-key: YOUR_API_KEY" \
  -d '{"email": "user@mailinator.com"}'
```

---

### 2. Check Multiple Emails (Batch)

Check multiple email addresses in a single request. Maximum 100 emails per
batch.

**Endpoint:** `POST /v1/networking/disposable/batch`

**Request Body:**

```json
{
  "emails": ["user1@tempmail.com", "user2@gmail.com", "user3@10minutemail.com"]
}
```

**Response:** `200 OK`

```json
{
  "data": {
    "results": [
      {
        "email": "user1@tempmail.com",
        "is_disposable": true,
        "domain": "tempmail.com"
      },
      {
        "email": "user2@gmail.com",
        "is_disposable": false,
        "domain": "gmail.com"
      },
      {
        "email": "user3@10minutemail.com",
        "is_disposable": true,
        "domain": "10minutemail.com"
      }
    ],
    "total": 3
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
# Production
curl -X POST https://api.requiems.xyz/v1/networking/disposable/batch \
  -H "Content-Type: application/json" \
  -H "requiems-api-key: YOUR_API_KEY" \
  -d '{
    "emails": [
      "user1@mailinator.com",
      "user2@gmail.com",
      "user3@10minutemail.com"
    ]
  }'
```

**Validation:**

- Maximum 100 emails per batch
- Returns `400 Bad Request` if limit exceeded

---

### 3. Check Domain

Check if a specific domain is in the disposable list.

**Endpoint:** `GET /v1/networking/disposable/domain/{domain}`

**Path Parameters:**

- `domain` - The domain to check (e.g., `tempmail.com`)

**Response:** `200 OK`

```json
{
  "data": {
    "domain": "tempmail.com",
    "is_disposable": true
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
# Production
curl https://api.requiems.xyz/v1/networking/disposable/domain/guerrillamail.com \
  -H "requiems-api-key: YOUR_API_KEY"
```

---

### 4. List All Disposable Domains

Get a paginated list of all disposable email domains in the blocklist.

**Endpoint:** `GET /v1/networking/disposable/domains`

**Query Parameters:**

- `page` (optional) - Page number, default: `1`
- `per_page` (optional) - Results per page (1-1000), default: `100`

**Response:** `200 OK`

```json
{
  "data": {
    "domains": [
      "0-mail.com",
      "0815.ru",
      "0clickemail.com",
      "...more domains..."
    ],
    "total": 15432,
    "page": 1,
    "per_page": 100,
    "has_more": true
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
# Production - Get first page (default 100 results)
curl https://api.requiems.xyz/v1/networking/disposable/domains \
  -H "requiems-api-key: YOUR_API_KEY"

# Production - Get page 2 with 50 results per page
curl "https://api.requiems.xyz/v1/networking/disposable/domains?page=2&per_page=50" \
  -H "requiems-api-key: YOUR_API_KEY"
```

**Notes:**

- There are thousands of domains in the blocklist
- Use pagination to avoid overwhelming responses
- `has_more` indicates if there are more pages available

---

### 5. Get Statistics

Get statistics about the disposable domains blocklist.

**Endpoint:** `GET /v1/networking/disposable/stats`

**Response:** `200 OK`

```json
{
  "data": {
    "total_domains": 15432
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

**Example:**

```bash
# Production
curl https://api.requiems.xyz/v1/networking/disposable/stats \
  -H "requiems-api-key: YOUR_API_KEY"
```

---

## Error Responses

All endpoints may return error responses in the following format:

```json
{
  "error": "error message description"
}
```

### Common Error Codes

- `400 Bad Request` - Invalid input (missing email, invalid JSON, batch limit
  exceeded)
- `503 Service Unavailable` - Service error

---

## Use Cases

### User Registration Validation

Prevent users from signing up with disposable email addresses:

```javascript
const response = await fetch(
  "https://api.requiems.xyz/v1/networking/disposable/check",
  {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "requiems-api-key": "YOUR_API_KEY",
    },
    body: JSON.stringify({ email: userEmail }),
  },
);

const { data } = await response.json();

if (data.is_disposable) {
  throw new Error("Please use a permanent email address");
}
```

### Bulk Email List Cleaning

Clean a list of email addresses by checking them in batches:

```javascript
const emails = [
  /* array of up to 100 emails */
];

const response = await fetch(
  "https://api.requiems.xyz/v1/networking/disposable/batch",
  {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "requiems-api-key": "YOUR_API_KEY",
    },
    body: JSON.stringify({ emails }),
  },
);

const { data } = await response.json();
const validEmails = data.results
  .filter((r) => !r.is_disposable)
  .map((r) => r.email);
```

### Domain Allowlist/Blocklist

Check if a domain should be blocked:

```javascript
const domain = email.split("@")[1];
const response = await fetch(
  `https://api.requiems.xyz/v1/networking/disposable/domain/${domain}`,
  {
    headers: { "requiems-api-key": "YOUR_API_KEY" },
  },
);
const { data } = await response.json();

if (data.is_disposable) {
  console.log("This domain is on the blocklist");
}
```

---
