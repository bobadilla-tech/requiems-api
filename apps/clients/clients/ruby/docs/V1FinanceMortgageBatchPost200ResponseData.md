# OpenapiClient::V1FinanceMortgageBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Mortgage calculation result for each mortgage request in the same order as the input. Each item has the same fields as the single mortgage endpoint. | [optional] |
| **total** | **Integer** | Number of results returned. Matches the length of the input array. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceMortgageBatchPost200ResponseData.new(
  results: null,
  total: null
)
```

