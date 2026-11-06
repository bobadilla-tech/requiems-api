# V1FinanceConvertGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from** | **string** | Source currency code (uppercased) | [optional] [default to undefined]
**to** | **string** | Target currency code (uppercased) | [optional] [default to undefined]
**rate** | **number** | Exchange rate used for the conversion | [optional] [default to undefined]
**amount** | **number** | The original amount passed in the request | [optional] [default to undefined]
**converted** | **number** | Result of amount × rate, rounded to 2 decimal places | [optional] [default to undefined]
**timestamp** | **string** | Date the rate was published by the ECB (ISO 8601) | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceConvertGet200ResponseData } from 'api-client';

const instance: V1FinanceConvertGet200ResponseData = {
    from,
    to,
    rate,
    amount,
    converted,
    timestamp,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
