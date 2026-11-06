# V1TechnologyPasswordGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**password** | **str** | The generated password | [optional] 
**length** | **int** | Length of the generated password | [optional] 
**strength** | **str** | Password strength assessment (weak, medium, or strong) | [optional] 

## Example

```python
from openapi_client.models.v1_technology_password_get200_response_data import V1TechnologyPasswordGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyPasswordGet200ResponseData from a JSON string
v1_technology_password_get200_response_data_instance = V1TechnologyPasswordGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyPasswordGet200ResponseData.to_json())

# convert the object into a dict
v1_technology_password_get200_response_data_dict = v1_technology_password_get200_response_data_instance.to_dict()
# create an instance of V1TechnologyPasswordGet200ResponseData from a dict
v1_technology_password_get200_response_data_from_dict = V1TechnologyPasswordGet200ResponseData.from_dict(v1_technology_password_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


