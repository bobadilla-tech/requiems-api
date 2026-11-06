# OpenapiClient::V1ValidationProfanityBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_text** | **String** | The original text that was checked. | [optional] |
| **results_result_has_profanity** | **Boolean** | Whether any profanity was detected. | [optional] |
| **results_result_censored** | **String** | Text with profane words replaced by asterisks. | [optional] |
| **results_result_flagged_words** | **String** | Deduplicated list of detected profane words (lowercase). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1ValidationProfanityBatchPost200ResponseData.new(
  results: null,
  results_text: null,
  results_result_has_profanity: null,
  results_result_censored: null,
  results_result_flagged_words: null,
  total: null
)
```

