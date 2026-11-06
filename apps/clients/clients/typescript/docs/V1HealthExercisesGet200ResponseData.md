# V1HealthExercisesGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **Array&lt;any&gt;** | Array of exercise objects for the current page | [optional] [default to undefined]
**items_id** | **number** | Unique exercise identifier | [optional] [default to undefined]
**items_name** | **string** | Exercise name | [optional] [default to undefined]
**items_body_parts** | **Array&lt;any&gt;** | Body part categories involved | [optional] [default to undefined]
**items_equipment** | **Array&lt;any&gt;** | Equipment required | [optional] [default to undefined]
**items_target_muscles** | **Array&lt;any&gt;** | Primary muscles targeted | [optional] [default to undefined]
**items_secondary_muscles** | **Array&lt;any&gt;** | Secondary muscles engaged | [optional] [default to undefined]
**items_instructions** | **Array&lt;any&gt;** | Ordered step-by-step instructions | [optional] [default to undefined]
**total** | **number** | Total number of exercises matching the filters | [optional] [default to undefined]
**page** | **number** | Current page number | [optional] [default to undefined]
**per_page** | **number** | Number of results per page | [optional] [default to undefined]

## Example

```typescript
import { V1HealthExercisesGet200ResponseData } from 'api-client';

const instance: V1HealthExercisesGet200ResponseData = {
    items,
    items_id,
    items_name,
    items_body_parts,
    items_equipment,
    items_target_muscles,
    items_secondary_muscles,
    items_instructions,
    total,
    page,
    per_page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
