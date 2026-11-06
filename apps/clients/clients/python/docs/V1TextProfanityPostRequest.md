# V1TextProfanityPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **str** | The text to check for profanity. | 

## Example

```python
from openapi_client.models.v1_text_profanity_post_request import V1TextProfanityPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextProfanityPostRequest from a JSON string
v1_text_profanity_post_request_instance = V1TextProfanityPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextProfanityPostRequest.to_json())

# convert the object into a dict
v1_text_profanity_post_request_dict = v1_text_profanity_post_request_instance.to_dict()
# create an instance of V1TextProfanityPostRequest from a dict
v1_text_profanity_post_request_from_dict = V1TextProfanityPostRequest.from_dict(v1_text_profanity_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


