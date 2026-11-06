# OpenapiClient::V1NetworkingWhoisBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domains** | **Array&lt;Object&gt;** | Array of domain names. Min: 1, Max: 50. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingWhoisBatchPostRequest.new(
  domains: [&quot;example.com&quot;, &quot;google.com&quot;]
)
```

