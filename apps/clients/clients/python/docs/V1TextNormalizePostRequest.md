# V1TextNormalizePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **str** | The email address to normalize. Must be a syntactically valid address. | 

## Example

```python
from openapi_client.models.v1_text_normalize_post_request import V1TextNormalizePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextNormalizePostRequest from a JSON string
v1_text_normalize_post_request_instance = V1TextNormalizePostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextNormalizePostRequest.to_json())

# convert the object into a dict
v1_text_normalize_post_request_dict = v1_text_normalize_post_request_instance.to_dict()
# create an instance of V1TextNormalizePostRequest from a dict
v1_text_normalize_post_request_from_dict = V1TextNormalizePostRequest.from_dict(v1_text_normalize_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


