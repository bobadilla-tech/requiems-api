# V1SystemsTextNormalizePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**original** | **str** | The input text as provided | [optional] 
**normalized** | **str** | The text after all operations have been applied in order | [optional] 
**operations** | **List[object]** | The operations list echoed from the request | [optional] 

## Example

```python
from openapi_client.models.v1_systems_text_normalize_post200_response_data import V1SystemsTextNormalizePost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsTextNormalizePost200ResponseData from a JSON string
v1_systems_text_normalize_post200_response_data_instance = V1SystemsTextNormalizePost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsTextNormalizePost200ResponseData.to_json())

# convert the object into a dict
v1_systems_text_normalize_post200_response_data_dict = v1_systems_text_normalize_post200_response_data_instance.to_dict()
# create an instance of V1SystemsTextNormalizePost200ResponseData from a dict
v1_systems_text_normalize_post200_response_data_from_dict = V1SystemsTextNormalizePost200ResponseData.from_dict(v1_systems_text_normalize_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


