# V1FinanceIbanBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Validation result for each number in the same order as the input. Each item has the same fields as the single validate endpoint. | [optional] [default to undefined]
**total** | **number** | Number of results returned. Matches the length of the input array. | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceIbanBatchPost200ResponseData } from 'api-client';

const instance: V1FinanceIbanBatchPost200ResponseData = {
    results,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
