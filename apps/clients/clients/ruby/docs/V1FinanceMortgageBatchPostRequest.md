# OpenapiClient::V1FinanceMortgageBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mortgages** | **Array&lt;Object&gt;** | Array of mortgages to calculate (min: 1, max: 50). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceMortgageBatchPostRequest.new(
  mortgages: [{&quot;principal&quot;:150000,&quot;rate&quot;:6.5,&quot;years&quot;:10},{&quot;principal&quot;:250000,&quot;rate&quot;:5.2,&quot;years&quot;:30}]
)
```

