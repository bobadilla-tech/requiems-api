# OpenapiClient::V1TechnologyPasswordBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_result_password** | **String** | The generated password (omitted on error). | [optional] |
| **results_result_length** | **Integer** | Length of the generated password. | [optional] |
| **results_result_strength** | **String** | Strength assessment (weak, medium, or strong). | [optional] |
| **results_error** | **String** | Error message if generation failed (omitted on success). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyPasswordBatchPost200ResponseData.new(
  results: null,
  results_result_password: null,
  results_result_length: null,
  results_result_strength: null,
  results_error: null,
  total: null
)
```

