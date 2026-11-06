# V1TechnologyBaseBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_from** | **number** | The source base used for conversion. | [optional] [default to undefined]
**results_to** | **number** | The target base used for conversion. | [optional] [default to undefined]
**results_input** | **string** | The original value as provided in the request. | [optional] [default to undefined]
**results_result** | **string** | The converted value in the target base (omitted on error). | [optional] [default to undefined]
**results_error** | **string** | Error message if this item failed (omitted on success). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyBaseBatchPost200ResponseData } from 'api-client';

const instance: V1TechnologyBaseBatchPost200ResponseData = {
    results,
    results_from,
    results_to,
    results_input,
    results_result,
    results_error,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
