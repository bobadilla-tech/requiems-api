# V1TechnologyConvertGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**var_from** | **str** | Source unit key | [optional] 
**to** | **str** | Target unit key | [optional] 
**input** | **float** | The original input value | [optional] 
**result** | **float** | The converted value (rounded to 6 decimal places) | [optional] 
**formula** | **str** | Human-readable conversion formula | [optional] 

## Example

```python
from openapi_client.models.v1_technology_convert_get200_response_data import V1TechnologyConvertGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyConvertGet200ResponseData from a JSON string
v1_technology_convert_get200_response_data_instance = V1TechnologyConvertGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyConvertGet200ResponseData.to_json())

# convert the object into a dict
v1_technology_convert_get200_response_data_dict = v1_technology_convert_get200_response_data_instance.to_dict()
# create an instance of V1TechnologyConvertGet200ResponseData from a dict
v1_technology_convert_get200_response_data_from_dict = V1TechnologyConvertGet200ResponseData.from_dict(v1_technology_convert_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


