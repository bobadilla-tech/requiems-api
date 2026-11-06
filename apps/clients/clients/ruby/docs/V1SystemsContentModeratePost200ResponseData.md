# OpenapiClient::V1SystemsContentModeratePost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **is_safe** | **Boolean** | False when any moderation category is triggered. True otherwise. | [optional] |
| **toxicity_score** | **String** | Reserved for future use. Currently always null. | [optional] |
| **sentiment** | **String** | Detected sentiment: positive, negative, or neutral | [optional] |
| **language** | **String** | Detected or provided ISO 639-1 language code. Null when detection confidence is 0. | [optional] |
| **language_confidence** | **String** | Confidence in the detected language from 0.0 to 1.0. 1.0 when language was explicitly provided. | [optional] |
| **flags** | **Array&lt;Object&gt;** | Triggered moderation flags. Currently possible: text_profanity | [optional] |
| **categories** | **Object** | Per-category boolean flags: profanity, hate_speech, spam, violence | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsContentModeratePost200ResponseData.new(
  is_safe: null,
  toxicity_score: null,
  sentiment: null,
  language: null,
  language_confidence: null,
  flags: null,
  categories: null
)
```

