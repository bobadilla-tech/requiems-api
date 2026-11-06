# OpenapiClient::V1SystemsContentModeratePostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **text** | **String** | The text to moderate. Any language. Required. |  |
| **language** | **String** | ISO 639-1 language code (e.g. en, de, fr). When provided, language detection is skipped and confidence is set to 1.0. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsContentModeratePostRequest.new(
  text: This is some user-generated content.,
  language: en
)
```

