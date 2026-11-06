# V1HealthExercisesGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **List[object]** | Array of exercise objects for the current page | [optional] 
**items_id** | **int** | Unique exercise identifier | [optional] 
**items_name** | **str** | Exercise name | [optional] 
**items_body_parts** | **List[object]** | Body part categories involved | [optional] 
**items_equipment** | **List[object]** | Equipment required | [optional] 
**items_target_muscles** | **List[object]** | Primary muscles targeted | [optional] 
**items_secondary_muscles** | **List[object]** | Secondary muscles engaged | [optional] 
**items_instructions** | **List[object]** | Ordered step-by-step instructions | [optional] 
**total** | **int** | Total number of exercises matching the filters | [optional] 
**page** | **int** | Current page number | [optional] 
**per_page** | **int** | Number of results per page | [optional] 

## Example

```python
from openapi_client.models.v1_health_exercises_get200_response_data import V1HealthExercisesGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthExercisesGet200ResponseData from a JSON string
v1_health_exercises_get200_response_data_instance = V1HealthExercisesGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1HealthExercisesGet200ResponseData.to_json())

# convert the object into a dict
v1_health_exercises_get200_response_data_dict = v1_health_exercises_get200_response_data_instance.to_dict()
# create an instance of V1HealthExercisesGet200ResponseData from a dict
v1_health_exercises_get200_response_data_from_dict = V1HealthExercisesGet200ResponseData.from_dict(v1_health_exercises_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


