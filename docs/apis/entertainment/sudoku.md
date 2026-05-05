# Sudoku API

## Status

✅ **MVP** - Live

## Overview

Generate Sudoku puzzles of varying difficulty levels, complete with their
solutions. Each request returns a freshly generated, unique puzzle.

## Endpoints

### Batch Generate Sudoku Puzzles

**Endpoint:** `POST /v1/entertainment/sudoku/batch`

Generate up to 20 Sudoku puzzles in a single request. Results are returned in the same order as the input array. Each puzzle in the batch counts as one unit of API usage (billed via `X-Usage-Count`).

#### Request Body

| Field      | Type            | Required | Description                                                                               |
| ---------- | --------------- | -------- | ----------------------------------------------------------------------------------------- |
| `puzzles`  | array of string | Yes      | List of difficulty levels to generate. Each must be `easy`, `medium`, or `hard` (min: 1, max: 20). |

#### Response Fields

| Field           | Type         | Description                                                 |
| --------------- | ------------ | ----------------------------------------------------------- |
| `results`       | array        | Generated puzzles in the same order as the input array.     |
| `results[].difficulty` | string | The difficulty level of each puzzle.                 |
| `results[].puzzle`     | array[array] | 9×9 grid — `0` represents an empty cell.          |
| `results[].solution`   | array[array] | 9×9 grid containing the complete solution.        |
| `total`         | integer      | Number of puzzles returned. Matches the length of the input. |

#### Example Request

```json
{
  "puzzles": ["easy", "hard"]
}
```

#### Example Response

```json
{
  "data": {
    "results": [
      {
        "difficulty": "easy",
        "puzzle": [[5, 3, 0, 0, 7, 0, 0, 0, 0], ["..."]],
        "solution": [[5, 3, 4, 6, 7, 8, 9, 1, 2], ["..."]]
      },
      {
        "difficulty": "hard",
        "puzzle": [[0, 0, 0, 0, 0, 0, 0, 0, 0], ["..."]],
        "solution": [[1, 2, 3, 4, 5, 6, 7, 8, 9], ["..."]]
      }
    ],
    "total": 2
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

#### Billing

Each puzzle in the batch is billed as one request unit. Sending 10 puzzles consumes 10 units of quota. The gateway receives the item count via the internal `X-Usage-Count` response header (stripped before reaching the client).

#### Errors

| Code                | Status | When                                                              |
| ------------------- | ------ | ----------------------------------------------------------------- |
| `validation_failed` | 422    | `puzzles` is missing, empty, exceeds 20 items, or contains an invalid difficulty value. |

---

### Get Sudoku Puzzle

**Endpoint:** `GET /v1/entertainment/sudoku`

Generate a random Sudoku puzzle.

#### Query Parameters

| Parameter    | Type   | Required | Description                                                           |
| ------------ | ------ | -------- | --------------------------------------------------------------------- |
| `difficulty` | string | No       | Puzzle difficulty: `easy`, `medium`, or `hard`. Defaults to `medium`. |

#### Response Fields

| Field        | Type         | Description                                         |
| ------------ | ------------ | --------------------------------------------------- |
| `difficulty` | string       | The difficulty level of the returned puzzle.        |
| `puzzle`     | array[array] | 9×9 grid — `0` represents an empty cell to fill in. |
| `solution`   | array[array] | 9×9 grid containing the complete solution.          |

#### Example Response

```json
{
  "data": {
    "difficulty": "hard",
    "puzzle": [
      [5, 3, 0, 0, 7, 0, 0, 0, 0],
      [6, 0, 0, 1, 9, 5, 0, 0, 0],
      [0, 9, 8, 0, 0, 0, 0, 6, 0],
      [8, 0, 0, 0, 6, 0, 0, 0, 3],
      [4, 0, 0, 8, 0, 3, 0, 0, 1],
      [7, 0, 0, 0, 2, 0, 0, 0, 6],
      [0, 6, 0, 0, 0, 0, 2, 8, 0],
      [0, 0, 0, 4, 1, 9, 0, 0, 5],
      [0, 0, 0, 0, 8, 0, 0, 7, 9]
    ],
    "solution": [
      [5, 3, 4, 6, 7, 8, 9, 1, 2],
      [6, 7, 2, 1, 9, 5, 3, 4, 8],
      [1, 9, 8, 3, 4, 2, 5, 6, 7],
      [8, 5, 9, 7, 6, 1, 4, 2, 3],
      [4, 2, 6, 8, 5, 3, 7, 9, 1],
      [7, 1, 3, 9, 2, 4, 8, 5, 6],
      [9, 6, 1, 5, 3, 7, 2, 8, 4],
      [2, 8, 7, 4, 1, 9, 6, 3, 5],
      [3, 4, 5, 2, 8, 6, 1, 7, 9]
    ]
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```
