# OpenapiClient::V1TextProfanityPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **has_profanity** | **Boolean** | Whether any profanity was detected in the text | [optional] |
| **censored** | **String** | The input text with profane words replaced by asterisks | [optional] |
| **flagged_words** | **String** | Deduplicated list of profane words found (lowercase) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextProfanityPost200ResponseData.new(
  has_profanity: null,
  censored: null,
  flagged_words: null
)
```

