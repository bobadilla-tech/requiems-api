# OpenapiClient::V1TechnologyConvertBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **operations** | **Array&lt;Object&gt;** | Array of operations to convert (min: 1, max: 50). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyConvertBatchPostRequest.new(
  operations: [{&quot;from&quot;: &quot;kg&quot;, &quot;to&quot;: &quot;lb&quot;, &quot;value&quot;: 10}, {&quot;from&quot;: &quot;c&quot;, &quot;to&quot;: &quot;f&quot;, &quot;value&quot;: 5}]

)
```

