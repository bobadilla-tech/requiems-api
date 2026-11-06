# OpenapiClient::V1EntertainmentSudokuGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **difficulty** | **String** | The difficulty level of the returned puzzle (easy, medium, or hard) | [optional] |
| **puzzle** | **String** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] |
| **solution** | **String** | 9×9 grid containing the complete, valid solution | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentSudokuGet200ResponseData.new(
  difficulty: null,
  puzzle: null,
  solution: null
)
```

