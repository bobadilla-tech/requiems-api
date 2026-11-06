# V1ValidationProfanityPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **str** | The text to check for profanity. | 

## Example

```python
from openapi_client.models.v1_validation_profanity_post_request import V1ValidationProfanityPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationProfanityPostRequest from a JSON string
v1_validation_profanity_post_request_instance = V1ValidationProfanityPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1ValidationProfanityPostRequest.to_json())

# convert the object into a dict
v1_validation_profanity_post_request_dict = v1_validation_profanity_post_request_instance.to_dict()
# create an instance of V1ValidationProfanityPostRequest from a dict
v1_validation_profanity_post_request_from_dict = V1ValidationProfanityPostRequest.from_dict(v1_validation_profanity_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


