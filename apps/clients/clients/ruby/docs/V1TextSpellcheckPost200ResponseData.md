# OpenapiClient::V1TextSpellcheckPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **corrected** | **String** | The full input text with all misspelled words replaced by the top-ranked suggestion | [optional] |
| **corrections** | **String** | List of individual corrections. Each item contains: original (the misspelled word), suggested (the top-ranked replacement applied to corrected), suggestions (up to 3 ranked alternatives returned by LanguageTool, from most to least likely), and position (0-based Unicode character offset in the original text) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSpellcheckPost200ResponseData.new(
  corrected: null,
  corrections: null
)
```

