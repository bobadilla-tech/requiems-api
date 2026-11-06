# OpenapiClient::V1TextNormalizeBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **emails** | **Array&lt;Object&gt;** | Array of addresses to normalize (min 1, max 100; each entry non-empty) |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextNormalizeBatchPostRequest.new(
  emails: [&quot;user@example.com&quot;,&quot;not-an-email&quot;,&quot;te.st@gmail.com&quot;]
)
```

