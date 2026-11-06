# V1HealthExercisesBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Exercises found for the given IDs, in input order. IDs not found are omitted. | [optional] 
**results_id** | **int** | Unique exercise identifier | [optional] 
**results_name** | **str** | Exercise name | [optional] 
**results_body_parts** | **List[object]** | Body part categories involved | [optional] 
**results_equipment** | **List[object]** | Equipment required | [optional] 
**results_target_muscles** | **List[object]** | Primary muscles targeted | [optional] 
**results_secondary_muscles** | **List[object]** | Secondary muscles engaged | [optional] 
**results_instructions** | **List[object]** | Ordered step-by-step instructions | [optional] 
**total** | **int** | Number of exercises returned | [optional] 

## Example

```python
from openapi_client.models.v1_health_exercises_batch_post200_response_data import V1HealthExercisesBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthExercisesBatchPost200ResponseData from a JSON string
v1_health_exercises_batch_post200_response_data_instance = V1HealthExercisesBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1HealthExercisesBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_health_exercises_batch_post200_response_data_dict = v1_health_exercises_batch_post200_response_data_instance.to_dict()
# create an instance of V1HealthExercisesBatchPost200ResponseData from a dict
v1_health_exercises_batch_post200_response_data_from_dict = V1HealthExercisesBatchPost200ResponseData.from_dict(v1_health_exercises_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


