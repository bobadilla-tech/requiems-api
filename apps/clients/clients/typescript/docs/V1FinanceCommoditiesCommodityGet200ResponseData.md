# V1FinanceCommoditiesCommodityGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**commodity** | **string** | The commodity slug as provided in the request path | [optional] [default to undefined]
**name** | **string** | Human-readable commodity name | [optional] [default to undefined]
**price** | **number** | Latest annual average price in the commodity\&#39;s display unit | [optional] [default to undefined]
**unit** | **string** | Price unit (oz, barrel, mmbtu, lb, or metric_ton) | [optional] [default to undefined]
**currency** | **string** | Currency code — always USD | [optional] [default to undefined]
**change_24h** | **number** | Year-over-year percentage change from the prior year\&#39;s annual average (positive &#x3D; price increased) | [optional] [default to undefined]
**historical** | **Array&lt;any&gt;** | Up to 10 prior years of annual average prices, ordered newest to oldest | [optional] [default to undefined]
**historical_period** | **string** | Year of the historical data point | [optional] [default to undefined]
**historical_price** | **number** | Annual average price for that year | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceCommoditiesCommodityGet200ResponseData } from 'api-client';

const instance: V1FinanceCommoditiesCommodityGet200ResponseData = {
    commodity,
    name,
    price,
    unit,
    currency,
    change_24h,
    historical,
    historical_period,
    historical_price,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
