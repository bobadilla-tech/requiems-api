# OpenapiClient::V1FinanceConvertGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | Source currency code (uppercased) | [optional] |
| **to** | **String** | Target currency code (uppercased) | [optional] |
| **rate** | **Float** | Exchange rate used for the conversion | [optional] |
| **amount** | **Float** | The original amount passed in the request | [optional] |
| **converted** | **Float** | Result of amount × rate, rounded to 2 decimal places | [optional] |
| **timestamp** | **String** | Date the rate was published by the ECB (ISO 8601) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceConvertGet200ResponseData.new(
  from: null,
  to: null,
  rate: null,
  amount: null,
  converted: null,
  timestamp: null
)
```

