# OpenapiClient::V1TechnologyBase64DecodePostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **value** | **String** | The Base64-encoded string to decode |  |
| **variant** | **String** | Encoding variant: standard (default) or url (URL-safe base64url) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyBase64DecodePostRequest.new(
  value: SGVsbG8sIHdvcmxkIQ&#x3D;&#x3D;,
  variant: standard
)
```

