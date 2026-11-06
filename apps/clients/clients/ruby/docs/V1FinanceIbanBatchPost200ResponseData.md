# OpenapiClient::V1FinanceIbanBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Validation result for each number in the same order as the input. Each item has the same fields as the single validate endpoint. | [optional] |
| **total** | **Integer** | Number of results returned. Matches the length of the input array. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceIbanBatchPost200ResponseData.new(
  results: null,
  total: null
)
```

