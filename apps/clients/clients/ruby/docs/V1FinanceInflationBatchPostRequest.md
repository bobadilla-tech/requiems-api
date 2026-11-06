# OpenapiClient::V1FinanceInflationBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **countries** | **Array&lt;Object&gt;** | Array of ISO 3166-1 alpha-2 country codes. Min: 1, Max: 50. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceInflationBatchPostRequest.new(
  countries: [&quot;US&quot;, &quot;AR&quot;, &quot;DE&quot;]
)
```

