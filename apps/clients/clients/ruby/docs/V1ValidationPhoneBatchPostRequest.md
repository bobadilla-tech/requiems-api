# OpenapiClient::V1ValidationPhoneBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **numbers** | **Array&lt;Object&gt;** | Array of phone numbers to validate (min: 1, max: 50). Each must include the country calling code. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1ValidationPhoneBatchPostRequest.new(
  numbers: [&quot;+447400123456&quot;, &quot;+12015551234&quot;]
)
```

