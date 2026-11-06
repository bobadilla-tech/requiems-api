# OpenapiClient::V1FinanceInflationBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | One entry per country, in the same order as the input array | [optional] |
| **results_country** | **String** | ISO 3166-1 alpha-2 country code, uppercased | [optional] |
| **results_found** | **Boolean** | false when the country has no data in the World Bank set | [optional] |
| **results_rate** | **Float** | Latest CPI inflation rate. Omitted when found: false | [optional] |
| **results_period** | **String** | Year of the latest data point. Omitted when found: false | [optional] |
| **results_historical** | **Array&lt;Object&gt;** | Up to 10 previous years. Omitted when found: false | [optional] |
| **total** | **Integer** | Total number of results returned (equals number of countries sent) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceInflationBatchPost200ResponseData.new(
  results: null,
  results_country: null,
  results_found: null,
  results_rate: null,
  results_period: null,
  results_historical: null,
  total: null
)
```

