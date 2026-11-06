# V1HealthBodyPartsGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **List[object]** | Sorted list of all distinct body part names | [optional] 
**total** | **int** | Total number of distinct body parts | [optional] 

## Example

```python
from openapi_client.models.v1_health_body_parts_get200_response_data import V1HealthBodyPartsGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthBodyPartsGet200ResponseData from a JSON string
v1_health_body_parts_get200_response_data_instance = V1HealthBodyPartsGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1HealthBodyPartsGet200ResponseData.to_json())

# convert the object into a dict
v1_health_body_parts_get200_response_data_dict = v1_health_body_parts_get200_response_data_instance.to_dict()
# create an instance of V1HealthBodyPartsGet200ResponseData from a dict
v1_health_body_parts_get200_response_data_from_dict = V1HealthBodyPartsGet200ResponseData.from_dict(v1_health_body_parts_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


