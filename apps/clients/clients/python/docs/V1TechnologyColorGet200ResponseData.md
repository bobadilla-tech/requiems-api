# V1TechnologyColorGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**input** | **str** | The original value passed in the value parameter | [optional] 
**result** | **str** | The color expressed in the requested to format | [optional] 
**formats_hex** | **str** | HEX representation (#rrggbb) | [optional] 
**formats_rgb** | **str** | RGB representation (rgb(r, g, b)) | [optional] 
**formats_hsl** | **str** | HSL representation (hsl(h, s%, l%)) | [optional] 
**formats_cmyk** | **str** | CMYK representation (cmyk(c%, m%, y%, k%)) | [optional] 

## Example

```python
from openapi_client.models.v1_technology_color_get200_response_data import V1TechnologyColorGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyColorGet200ResponseData from a JSON string
v1_technology_color_get200_response_data_instance = V1TechnologyColorGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyColorGet200ResponseData.to_json())

# convert the object into a dict
v1_technology_color_get200_response_data_dict = v1_technology_color_get200_response_data_instance.to_dict()
# create an instance of V1TechnologyColorGet200ResponseData from a dict
v1_technology_color_get200_response_data_from_dict = V1TechnologyColorGet200ResponseData.from_dict(v1_technology_color_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


