# OpenapiClient::V1FinanceCryptoSymbolGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **symbol** | **String** | Uppercase ticker symbol | [optional] |
| **name** | **String** | Full coin name | [optional] |
| **price_usd** | **Float** | Current price in USD | [optional] |
| **change_24h** | **Float** | Price change over the last 24 hours as a percentage | [optional] |
| **market_cap** | **Float** | Total market capitalisation in USD | [optional] |
| **volume_24h** | **Float** | Total trading volume over the last 24 hours in USD | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceCryptoSymbolGet200ResponseData.new(
  symbol: null,
  name: null,
  price_usd: null,
  change_24h: null,
  market_cap: null,
  volume_24h: null
)
```

