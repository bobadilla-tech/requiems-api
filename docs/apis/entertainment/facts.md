# Facts API

## Status

⏳ **Planned** - Not yet implemented

## Overview

Get random interesting facts. This endpoint will provide fascinating and
educational facts about various topics.

## Planned Endpoints

### Get Random Fact

**Planned Endpoint:** `GET /v1/entertainment/facts`

Get a random interesting fact.

### Batch Random Facts

**`POST /v1/entertainment/facts/batch`**

Returns one randomly selected fact per requested category in a single HTTP
call. Invalid categories are returned with an `error` field and do not fail
the rest of the batch. Results are always in the same order as the input array.

**Billing:** Each category submitted counts as one unit of quota (the gateway
reads `X-Usage-Count` from the response, strips it, and charges accordingly).
A request with 10 categories costs 10 units.

#### Request body

```json
{
  "categories": ["science", "history", "space"]
}
```

| Field        | Type             | Required | Constraints                                         |
| ------------ | ---------------- | -------- | --------------------------------------------------- |
| `categories` | array of strings | Yes      | 1 – 50 items. Case-insensitive. Duplicates allowed. |

#### Response

```json
{
  "data": {
    "results": [
      {
        "category": "science",
        "fact": "Octopuses have three hearts and blue blood.",
        "source": "National Geographic"
      },
      {
        "category": "history",
        "fact": "Oxford University is older than the Aztec Empire.",
        "source": "Oxford University"
      },
      {
        "category": "dragons",
        "error": "no facts found for category: dragons"
      }
    ],
    "total": 3
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

#### Response fields

| Field                | Type    | Description                                        |
| -------------------- | ------- | -------------------------------------------------- |
| `results`            | array   | One entry per requested category, in input order.  |
| `results[].category` | string  | The requested category.                            |
| `results[].fact`     | string  | The fact text. Omitted on error.                   |
| `results[].source`   | string  | Source attribution. Omitted on error.              |
| `results[].error`    | string  | Present only when the item could not be fulfilled. |
| `total`              | integer | Length of the results array.                       |

#### Errors

| Code | Error               | Cause                                                |
| ---- | ------------------- | ---------------------------------------------------- |
| 400  | `bad_request`       | Malformed JSON or invalid request body.              |
| 422  | `validation_failed` | Struct validation failure (empty or >50 categories). |

Per-item failures (unknown category) are **not** top-level errors — they appear
as `error` fields on the relevant result slot.

#### Examples

```bash
curl -X POST https://api.requiems.xyz/v1/entertainment/facts/batch \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"categories": ["science", "history", "space"]}'
```

```python
import requests

results = requests.post(
    "https://api.requiems.xyz/v1/entertainment/facts/batch",
    headers={"requiems-api-key": "YOUR_API_KEY"},
    json={"categories": ["science", "history", "space"]},
).json()["data"]["results"]

for item in results:
    if "error" in item:
        print(f"[{item['category']}] ERROR: {item['error']}")
    else:
        print(f"[{item['category']}] {item['fact']} — {item['source']}")
```

```javascript
const { data } = await fetch(
  "https://api.requiems.xyz/v1/entertainment/facts/batch",
  {
    method: "POST",
    headers: {
      "requiems-api-key": "YOUR_API_KEY",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ categories: ["science", "history", "space"] }),
  },
).then((r) => r.json());

for (const item of data.results) {
  if (item.error) console.error(`[${item.category}] ${item.error}`);
  else console.log(`[${item.category}] ${item.fact} — ${item.source}`);
}
```

```ruby
require 'net/http'
require 'json'

uri = URI('https://api.requiems.xyz/v1/entertainment/facts/batch')
req = Net::HTTP::Post.new(uri)
req['requiems-api-key'] = 'YOUR_API_KEY'
req['Content-Type']     = 'application/json'
req.body = { categories: ['science', 'history', 'space'] }.to_json

results = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) { |h|
  JSON.parse(h.request(req).body)['data']['results']
}

results.each do |item|
  if item['error']
    puts "[#{item['category']}] ERROR: #{item['error']}"
  else
    puts "[#{item['category']}] #{item['fact']} — #{item['source']}"
  end
end
```

---

## Categories

| Category     | Description                            |
| ------------ | -------------------------------------- |
| `science`    | Biology, physics, chemistry, astronomy |
| `history`    | Historical events and civilisations    |
| `technology` | Computing, internet, and inventions    |
| `nature`     | Animals, plants, and natural phenomena |
| `space`      | Planets, stars, and the cosmos         |
| `food`       | Culinary facts and botanical oddities  |

---

## FAQ

**What happens if I pass an invalid category to the batch endpoint?**
That result slot gets an `error` field; the rest of the batch is unaffected and
the overall response is still `200 OK`.

**Can I request the same category multiple times?**
Yes. Each slot is drawn independently, so you may get the same fact more than
once if the pool for that category is small.

**How does billing work for the batch endpoint?**
Each category counts as one unit. Submitting 10 categories costs 10 units,
regardless of how many succeed.

**Does every fact have a source?**
Yes. Every entry in the database carries a source from a reputable publication
or institution.

---
