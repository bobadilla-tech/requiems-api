# OpenapiClient::V1TextDetectLanguagePost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **language** | **String** | Full name of the detected language (e.g. French, English, Spanish) | [optional] |
| **code** | **String** | ISO 639-1 two-letter language code (e.g. fr, en, es). Empty string when detection is unreliable. | [optional] |
| **confidence** | **String** | Confidence score between 0.0 and 1.0. 0.0 is returned when the language cannot be reliably detected. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextDetectLanguagePost200ResponseData.new(
  language: null,
  code: null,
  confidence: null
)
```

