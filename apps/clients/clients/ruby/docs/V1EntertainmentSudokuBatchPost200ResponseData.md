# OpenapiClient::V1EntertainmentSudokuBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Generated puzzles in the same order as the input array. Each item has the same fields as the single-puzzle endpoint. | [optional] |
| **results_difficulty** | **String** | The difficulty level of the puzzle (easy, medium, or hard) | [optional] |
| **results_puzzle** | **String** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] |
| **results_solution** | **String** | 9×9 grid containing the complete, valid solution | [optional] |
| **total** | **Integer** | Number of puzzles returned. Matches the length of the input array. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentSudokuBatchPost200ResponseData.new(
  results: null,
  results_difficulty: null,
  results_puzzle: null,
  results_solution: null,
  total: null
)
```

