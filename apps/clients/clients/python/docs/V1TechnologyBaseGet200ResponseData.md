# V1TechnologyBaseGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**input** | **str** | The original value as provided in the request | [optional] 
**var_from** | **int** | The source base | [optional] 
**to** | **int** | The target base | [optional] 
**result** | **str** | The converted value in the target base | [optional] 

## Example

```python
from openapi_client.models.v1_technology_base_get200_response_data import V1TechnologyBaseGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyBaseGet200ResponseData from a JSON string
v1_technology_base_get200_response_data_instance = V1TechnologyBaseGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyBaseGet200ResponseData.to_json())

# convert the object into a dict
v1_technology_base_get200_response_data_dict = v1_technology_base_get200_response_data_instance.to_dict()
# create an instance of V1TechnologyBaseGet200ResponseData from a dict
v1_technology_base_get200_response_data_from_dict = V1TechnologyBaseGet200ResponseData.from_dict(v1_technology_base_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


