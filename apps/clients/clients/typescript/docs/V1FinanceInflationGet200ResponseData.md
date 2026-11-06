# V1FinanceInflationGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**country** | **string** | ISO 3166-1 alpha-2 country code, uppercased | [optional] [default to undefined]
**rate** | **number** | Latest annual CPI inflation rate as a percentage (e.g. 2.9495 means 2.9495%) | [optional] [default to undefined]
**period** | **string** | Year of the latest data point (e.g. 2024) | [optional] [default to undefined]
**historical** | **Array&lt;any&gt;** | Up to 10 previous years of inflation data, ordered newest to oldest | [optional] [default to undefined]
**historical_period** | **string** | Year of the historical data point | [optional] [default to undefined]
**historical_rate** | **number** | Annual CPI inflation rate for that year | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceInflationGet200ResponseData } from 'api-client';

const instance: V1FinanceInflationGet200ResponseData = {
    country,
    rate,
    period,
    historical,
    historical_period,
    historical_rate,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
