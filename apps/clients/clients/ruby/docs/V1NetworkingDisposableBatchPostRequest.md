# OpenapiClient::V1NetworkingDisposableBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **emails** | **Array&lt;Object&gt;** | Array of email addresses to check (max 100) |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingDisposableBatchPostRequest.new(
  emails: [&quot;user1@example.com&quot;,&quot;user2@tempmail.com&quot;]
)
```

