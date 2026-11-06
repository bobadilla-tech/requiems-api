# OpenapiClient::V1EntertainmentQuotesRandomBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Array of random quotes, one per requested count. | [optional] |
| **results_id** | **Integer** | Unique identifier for the quote. | [optional] |
| **results_text** | **String** | The quote text. | [optional] |
| **results_author** | **String** | Name of the person who said or wrote the quote. May be empty. | [optional] |
| **total** | **Integer** | Total number of quotes returned. Always equals the requested count. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentQuotesRandomBatchPost200ResponseData.new(
  results: null,
  results_id: null,
  results_text: null,
  results_author: null,
  total: null
)
```

