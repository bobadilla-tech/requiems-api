# V1TechnologyFormatPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**var_from** | **str** | Source format. One of: json, yaml, csv, xml, toml | 
**to** | **str** | Target format. One of: json, yaml, csv, xml, toml | 
**content** | **str** | The content to convert, serialized as a string in the source format. | 

## Example

```python
from openapi_client.models.v1_technology_format_post_request import V1TechnologyFormatPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyFormatPostRequest from a JSON string
v1_technology_format_post_request_instance = V1TechnologyFormatPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyFormatPostRequest.to_json())

# convert the object into a dict
v1_technology_format_post_request_dict = v1_technology_format_post_request_instance.to_dict()
# create an instance of V1TechnologyFormatPostRequest from a dict
v1_technology_format_post_request_from_dict = V1TechnologyFormatPostRequest.from_dict(v1_technology_format_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


