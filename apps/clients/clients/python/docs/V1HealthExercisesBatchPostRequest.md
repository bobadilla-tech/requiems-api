# V1HealthExercisesBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ids** | **List[object]** | Array of numeric exercise IDs to fetch. Must contain between 1 and 50 items. Each ID must be a positive integer. | 

## Example

```python
from openapi_client.models.v1_health_exercises_batch_post_request import V1HealthExercisesBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthExercisesBatchPostRequest from a JSON string
v1_health_exercises_batch_post_request_instance = V1HealthExercisesBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1HealthExercisesBatchPostRequest.to_json())

# convert the object into a dict
v1_health_exercises_batch_post_request_dict = v1_health_exercises_batch_post_request_instance.to_dict()
# create an instance of V1HealthExercisesBatchPostRequest from a dict
v1_health_exercises_batch_post_request_from_dict = V1HealthExercisesBatchPostRequest.from_dict(v1_health_exercises_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


