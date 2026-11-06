# OpenapiClient::V1TechnologyBaseBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_from** | **Integer** | The source base used for conversion. | [optional] |
| **results_to** | **Integer** | The target base used for conversion. | [optional] |
| **results_input** | **String** | The original value as provided in the request. | [optional] |
| **results_result** | **String** | The converted value in the target base (omitted on error). | [optional] |
| **results_error** | **String** | Error message if this item failed (omitted on success). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyBaseBatchPost200ResponseData.new(
  results: null,
  results_from: null,
  results_to: null,
  results_input: null,
  results_result: null,
  results_error: null,
  total: null
)
```

