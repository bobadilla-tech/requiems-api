# V1HealthExercisesBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Exercises found for the given IDs, in input order. IDs not found are omitted. | [optional] [default to undefined]
**results_id** | **number** | Unique exercise identifier | [optional] [default to undefined]
**results_name** | **string** | Exercise name | [optional] [default to undefined]
**results_body_parts** | **Array&lt;any&gt;** | Body part categories involved | [optional] [default to undefined]
**results_equipment** | **Array&lt;any&gt;** | Equipment required | [optional] [default to undefined]
**results_target_muscles** | **Array&lt;any&gt;** | Primary muscles targeted | [optional] [default to undefined]
**results_secondary_muscles** | **Array&lt;any&gt;** | Secondary muscles engaged | [optional] [default to undefined]
**results_instructions** | **Array&lt;any&gt;** | Ordered step-by-step instructions | [optional] [default to undefined]
**total** | **number** | Number of exercises returned | [optional] [default to undefined]

## Example

```typescript
import { V1HealthExercisesBatchPost200ResponseData } from 'api-client';

const instance: V1HealthExercisesBatchPost200ResponseData = {
    results,
    results_id,
    results_name,
    results_body_parts,
    results_equipment,
    results_target_muscles,
    results_secondary_muscles,
    results_instructions,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
