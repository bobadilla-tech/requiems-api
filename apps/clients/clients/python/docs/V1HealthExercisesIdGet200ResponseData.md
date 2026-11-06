# V1HealthExercisesIdGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique exercise identifier | [optional] 
**name** | **str** | Exercise name | [optional] 
**body_parts** | **List[object]** | Body part categories involved | [optional] 
**equipment** | **List[object]** | Equipment required | [optional] 
**target_muscles** | **List[object]** | Primary muscles targeted | [optional] 
**secondary_muscles** | **List[object]** | Secondary muscles engaged | [optional] 
**instructions** | **List[object]** | Ordered step-by-step instructions | [optional] 

## Example

```python
from openapi_client.models.v1_health_exercises_id_get200_response_data import V1HealthExercisesIdGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthExercisesIdGet200ResponseData from a JSON string
v1_health_exercises_id_get200_response_data_instance = V1HealthExercisesIdGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1HealthExercisesIdGet200ResponseData.to_json())

# convert the object into a dict
v1_health_exercises_id_get200_response_data_dict = v1_health_exercises_id_get200_response_data_instance.to_dict()
# create an instance of V1HealthExercisesIdGet200ResponseData from a dict
v1_health_exercises_id_get200_response_data_from_dict = V1HealthExercisesIdGet200ResponseData.from_dict(v1_health_exercises_id_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


