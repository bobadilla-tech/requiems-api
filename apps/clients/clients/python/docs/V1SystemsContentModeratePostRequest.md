# V1SystemsContentModeratePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **str** | The text to moderate. Any language. Required. | 
**language** | **str** | ISO 639-1 language code (e.g. en, de, fr). When provided, language detection is skipped and confidence is set to 1.0. | [optional] 

## Example

```python
from openapi_client.models.v1_systems_content_moderate_post_request import V1SystemsContentModeratePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsContentModeratePostRequest from a JSON string
v1_systems_content_moderate_post_request_instance = V1SystemsContentModeratePostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsContentModeratePostRequest.to_json())

# convert the object into a dict
v1_systems_content_moderate_post_request_dict = v1_systems_content_moderate_post_request_instance.to_dict()
# create an instance of V1SystemsContentModeratePostRequest from a dict
v1_systems_content_moderate_post_request_from_dict = V1SystemsContentModeratePostRequest.from_dict(v1_systems_content_moderate_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


