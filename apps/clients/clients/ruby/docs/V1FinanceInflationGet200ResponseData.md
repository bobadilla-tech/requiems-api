# OpenapiClient::V1FinanceInflationGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **country** | **String** | ISO 3166-1 alpha-2 country code, uppercased | [optional] |
| **rate** | **Float** | Latest annual CPI inflation rate as a percentage (e.g. 2.9495 means 2.9495%) | [optional] |
| **period** | **String** | Year of the latest data point (e.g. 2024) | [optional] |
| **historical** | **Array&lt;Object&gt;** | Up to 10 previous years of inflation data, ordered newest to oldest | [optional] |
| **historical_period** | **String** | Year of the historical data point | [optional] |
| **historical_rate** | **Float** | Annual CPI inflation rate for that year | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceInflationGet200ResponseData.new(
  country: null,
  rate: null,
  period: null,
  historical: null,
  historical_period: null,
  historical_rate: null
)
```

