# OpenapiClient::V1NetworkingMxBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domains** | **Array&lt;Object&gt;** | Array of domains to lookup (min: 1, max: 50). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingMxBatchPostRequest.new(
  domains: [&quot;outlook.com&quot;, &quot;yahoo.com&quot;]
)
```

