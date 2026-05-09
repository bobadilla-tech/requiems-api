# Spell Check API

## Status

✅ **Live**

## Overview

Check spelling and get correction suggestions. This endpoint identifies
misspelled words in the input text, provides the best correction per word, and
returns a rebuilt version of the text with all corrections applied.

## Endpoint

`POST /v1/text/spellcheck`

## Request

```json
{
  "text": "Ths is a smiple tset"
}
```

| Field  | Type   | Required | Description          |
| ------ | ------ | -------- | -------------------- |
| `text` | string | ✅       | The text to analyse. |

## Response

```json
{
  "data": {
    "corrected": "This is a simple test",
    "corrections": [
      { "original": "Ths", "suggested": "This", "position": 0 },
      { "original": "smiple", "suggested": "simple", "position": 9 },
      { "original": "tset", "suggested": "test", "position": 16 }
    ]
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field         | Type             | Description                                                   |
| ------------- | ---------------- | ------------------------------------------------------------- |
| `corrected`   | string           | Input text with misspelled words replaced by their correction |
| `corrections` | array of objects | One entry per misspelled word (see below)                     |

Each `corrections` entry:

| Field       | Type   | Description                                            |
| ----------- | ------ | ------------------------------------------------------ |
| `original`  | string | The misspelled word as it appears in the original text |
| `suggested` | string | The best correction found                              |
| `position`  | int    | 0-based character offset of the word in the input text |

## Notes

- Only **English** is supported.
- Positions are **0-based** character offsets in the original input string.
- Capitalisation is **preserved**: if the original word starts with an uppercase
  letter, the suggestion will too.
- When no mistakes are found, `corrections` is an empty array (`[]`) and
  `corrected` equals the input.
- Only **ASCII letter sequences** (`[a-zA-Z]`) are spell-checked. Non-ASCII
  characters (accented letters, CJK, emoji, etc.) are passed through unchanged
  and do not affect position counting.

---

## Batch Endpoint

`POST /v1/text/spellcheck/batch`

Check spelling for multiple texts in a single request. Results are returned in the same order as the input array.

### Request

```json
{
  "texts": ["Ths is a tset", "Smiple example"]
}
```

| Field   | Type             | Required | Description                                   |
| ------- | ---------------- | -------- | --------------------------------------------- |
| `texts` | array of strings | ✅       | Texts to spell-check. Between 1 and 50 items. |

### Response

```json
{
  "data": {
    "results": [
      {
        "corrected": "This is a test",
        "corrections": [
          { "original": "Ths", "suggested": "This", "position": 0 },
          { "original": "tset", "suggested": "test", "position": 9 }
        ]
      },
      {
        "corrected": "Simple example",
        "corrections": [
          { "original": "Smiple", "suggested": "Simple", "position": 0 }
        ]
      }
    ],
    "total": 2
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

| Field     | Type             | Description                                                                                                                     |
| --------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `results` | array of objects | One entry per input text, in input order. Each item has `corrected` and `corrections` (same shape as the single endpoint). |
| `total`   | integer          | Number of texts processed.                                                                                                      |

### Limits

- Minimum **1** text per request, maximum **50**.
- Request body capped at **1 MiB**.

### Billing

Each text in the batch counts as **1 unit** of quota. A request with 10 texts consumes 10 units.

### Partial failure

This endpoint always returns `200` with a result for every input text. There are no per-item errors — spell-checking is pure local computation with no external dependencies. If the service is unavailable, the entire request fails with `500`.

### Batch Error Codes

| Code                | Status | When                                                       |
| ------------------- | ------ | ---------------------------------------------------------- |
| `validation_failed` | 422    | `texts` is missing, empty, or contains more than 50 items  |
| `bad_request`       | 400    | Request body is missing or not JSON                        |
| `internal_error`    | 500    | Unexpected failure                                         |

---

## Error Codes

| Code                | Status | When                                |
| ------------------- | ------ | ----------------------------------- |
| `validation_failed` | 422    | `text` field is missing or empty    |
| `bad_request`       | 400    | Request body is missing or not JSON |
| `internal_error`    | 500    | Unexpected failure                  |
