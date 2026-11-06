# OpenapiClient::V1NetworkingWhoisBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | One entry per domain, in the same order as the input | [optional] |
| **results_domain** | **String** | Domain name requested in the batch | [optional] |
| **results_found** | **Boolean** | False when WHOIS data could not be found | [optional] |
| **results_error** | **String** | Error message when found is false | [optional] |
| **results_data** | **Object** | WHOIS information for the domain when found is true | [optional] |
| **total** | **Integer** | Total number of results returned | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingWhoisBatchPost200ResponseData.new(
  results: null,
  results_domain: null,
  results_found: null,
  results_error: null,
  results_data: null,
  total: null
)
```

