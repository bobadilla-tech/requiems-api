# V1FinanceExchangeRateGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from** | **string** | Source currency code (uppercased) | [optional] [default to undefined]
**to** | **string** | Target currency code (uppercased) | [optional] [default to undefined]
**rate** | **number** | Exchange rate — how many units of &#x60;to&#x60; equal 1 unit of &#x60;from&#x60; | [optional] [default to undefined]
**timestamp** | **string** | Date the rate was published by the ECB (ISO 8601) | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceExchangeRateGet200ResponseData } from 'api-client';

const instance: V1FinanceExchangeRateGet200ResponseData = {
    from,
    to,
    rate,
    timestamp,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
