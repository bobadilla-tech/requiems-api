# OpenapiClient::V1FinanceExchangeRateGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | Source currency code (uppercased) | [optional] |
| **to** | **String** | Target currency code (uppercased) | [optional] |
| **rate** | **Float** | Exchange rate — how many units of &#x60;to&#x60; equal 1 unit of &#x60;from&#x60; | [optional] |
| **timestamp** | **String** | Date the rate was published by the ECB (ISO 8601) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceExchangeRateGet200ResponseData.new(
  from: null,
  to: null,
  rate: null,
  timestamp: null
)
```

