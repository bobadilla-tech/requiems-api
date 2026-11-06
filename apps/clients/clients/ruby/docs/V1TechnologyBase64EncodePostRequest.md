# OpenapiClient::V1TechnologyBase64EncodePostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **value** | **String** | The string to encode |  |
| **variant** | **String** | Encoding variant: standard (default) or url (URL-safe base64url, RFC 4648 §5) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyBase64EncodePostRequest.new(
  value: Hello, world!,
  variant: standard
)
```

