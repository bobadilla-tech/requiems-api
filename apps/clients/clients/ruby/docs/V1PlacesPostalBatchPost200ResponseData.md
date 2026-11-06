# OpenapiClient::V1PlacesPostalBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_code** | **String** | The postal code as provided in the request. | [optional] |
| **results_country** | **String** | The country code used for the lookup (defaulted to US when omitted). | [optional] |
| **results_found** | **Boolean** | Whether the postal code was found in the dataset. | [optional] |
| **results_result** | **Object** | Postal code details (omitted when found is false). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesPostalBatchPost200ResponseData.new(
  results: null,
  results_code: null,
  results_country: null,
  results_found: null,
  results_result: null,
  total: null
)
```

