# OpenapiClient::V1TextSpellcheckBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **String** | One entry per input text, in the same order as the input array. Each item contains corrected (the fixed text) and corrections (list of individual corrections with original, suggested, suggestions, and position). | [optional] |
| **total** | **Integer** | Number of texts processed. Equals the length of the input array. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSpellcheckBatchPost200ResponseData.new(
  results: null,
  total: null
)
```

