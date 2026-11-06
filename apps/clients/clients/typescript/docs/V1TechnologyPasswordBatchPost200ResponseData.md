# V1TechnologyPasswordBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_result_password** | **string** | The generated password (omitted on error). | [optional] [default to undefined]
**results_result_length** | **number** | Length of the generated password. | [optional] [default to undefined]
**results_result_strength** | **string** | Strength assessment (weak, medium, or strong). | [optional] [default to undefined]
**results_error** | **string** | Error message if generation failed (omitted on success). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyPasswordBatchPost200ResponseData } from 'api-client';

const instance: V1TechnologyPasswordBatchPost200ResponseData = {
    results,
    results_result_password,
    results_result_length,
    results_result_strength,
    results_error,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
