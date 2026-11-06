# OpenapiClient::V1FinanceCommoditiesCommodityGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **commodity** | **String** | The commodity slug as provided in the request path | [optional] |
| **name** | **String** | Human-readable commodity name | [optional] |
| **price** | **Float** | Latest annual average price in the commodity&#39;s display unit | [optional] |
| **unit** | **String** | Price unit (oz, barrel, mmbtu, lb, or metric_ton) | [optional] |
| **currency** | **String** | Currency code — always USD | [optional] |
| **change_24h** | **Float** | Year-over-year percentage change from the prior year&#39;s annual average (positive &#x3D; price increased) | [optional] |
| **historical** | **Array&lt;Object&gt;** | Up to 10 prior years of annual average prices, ordered newest to oldest | [optional] |
| **historical_period** | **String** | Year of the historical data point | [optional] |
| **historical_price** | **Float** | Annual average price for that year | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceCommoditiesCommodityGet200ResponseData.new(
  commodity: null,
  name: null,
  price: null,
  unit: null,
  currency: null,
  change_24h: null,
  historical: null,
  historical_period: null,
  historical_price: null
)
```

