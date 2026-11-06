# V1TextNormalizePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**original** | **str** | The email address exactly as supplied in the request body | [optional] 
**normalized** | **str** | The canonical form of the address after all transformations | [optional] 
**local** | **str** | The local part (before @) of the normalized address | [optional] 
**domain** | **str** | The domain part (after @) of the normalized address | [optional] 
**changes** | **List[object]** | Ordered list of transformations applied. Possible values: lowercased, trimmed_whitespace, removed_dots, removed_plus_tag, canonicalised_domain. Empty array when no changes were needed. | [optional] 

## Example

```python
from openapi_client.models.v1_text_normalize_post200_response_data import V1TextNormalizePost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextNormalizePost200ResponseData from a JSON string
v1_text_normalize_post200_response_data_instance = V1TextNormalizePost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextNormalizePost200ResponseData.to_json())

# convert the object into a dict
v1_text_normalize_post200_response_data_dict = v1_text_normalize_post200_response_data_instance.to_dict()
# create an instance of V1TextNormalizePost200ResponseData from a dict
v1_text_normalize_post200_response_data_from_dict = V1TextNormalizePost200ResponseData.from_dict(v1_text_normalize_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


