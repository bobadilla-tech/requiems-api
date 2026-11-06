# OpenapiClient::V1ValidationEmailBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **emails** | **Array&lt;Object&gt;** | Array of email addresses to validate. Min: 1, Max: 50. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1ValidationEmailBatchPostRequest.new(
  emails: [&quot;user@gmail.com&quot;, &quot;user@gmial.com&quot;]
)
```

