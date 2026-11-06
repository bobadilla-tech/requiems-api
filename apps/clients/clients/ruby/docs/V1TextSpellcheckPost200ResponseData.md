# OpenapiClient::V1TextSpellcheckPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **corrected** | **String** | The full input text with all misspelled words replaced by their suggested corrections | [optional] |
| **corrections** | **String** | List of individual corrections. Each item contains: original (the misspelled word), suggested (the correction), and position (0-based character offset in the original text) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSpellcheckPost200ResponseData.new(
  corrected: null,
  corrections: null
)
```

