# V1FinanceInflationBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | One entry per country, in the same order as the input array | [optional] [default to undefined]
**results_country** | **string** | ISO 3166-1 alpha-2 country code, uppercased | [optional] [default to undefined]
**results_found** | **boolean** | false when the country has no data in the World Bank set | [optional] [default to undefined]
**results_rate** | **number** | Latest CPI inflation rate. Omitted when found: false | [optional] [default to undefined]
**results_period** | **string** | Year of the latest data point. Omitted when found: false | [optional] [default to undefined]
**results_historical** | **Array&lt;any&gt;** | Up to 10 previous years. Omitted when found: false | [optional] [default to undefined]
**total** | **number** | Total number of results returned (equals number of countries sent) | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceInflationBatchPost200ResponseData } from 'api-client';

const instance: V1FinanceInflationBatchPost200ResponseData = {
    results,
    results_country,
    results_found,
    results_rate,
    results_period,
    results_historical,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
