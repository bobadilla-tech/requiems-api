# V1SystemsTextNormalizePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **str** | The text to normalize. Required. | 
**operations** | **List[object]** | Ordered list of normalization operations to apply. At least one required. Valid values: trim, lowercase, uppercase, normalize_unicode, collapse_whitespace, strip_html, remove_punctuation | 

## Example

```python
from openapi_client.models.v1_systems_text_normalize_post_request import V1SystemsTextNormalizePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsTextNormalizePostRequest from a JSON string
v1_systems_text_normalize_post_request_instance = V1SystemsTextNormalizePostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsTextNormalizePostRequest.to_json())

# convert the object into a dict
v1_systems_text_normalize_post_request_dict = v1_systems_text_normalize_post_request_instance.to_dict()
# create an instance of V1SystemsTextNormalizePostRequest from a dict
v1_systems_text_normalize_post_request_from_dict = V1SystemsTextNormalizePostRequest.from_dict(v1_systems_text_normalize_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


