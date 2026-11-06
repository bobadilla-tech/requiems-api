# V1ValidationProfanityBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**texts** | **str** | List of texts to check (1–50 items, each non-empty). | 

## Example

```python
from openapi_client.models.v1_validation_profanity_batch_post_request import V1ValidationProfanityBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationProfanityBatchPostRequest from a JSON string
v1_validation_profanity_batch_post_request_instance = V1ValidationProfanityBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1ValidationProfanityBatchPostRequest.to_json())

# convert the object into a dict
v1_validation_profanity_batch_post_request_dict = v1_validation_profanity_batch_post_request_instance.to_dict()
# create an instance of V1ValidationProfanityBatchPostRequest from a dict
v1_validation_profanity_batch_post_request_from_dict = V1ValidationProfanityBatchPostRequest.from_dict(v1_validation_profanity_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


