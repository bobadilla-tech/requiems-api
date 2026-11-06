# V1TextWordsBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **List[object]** | List of words to look up in batch (max 50 items) | 

## Example

```python
from openapi_client.models.v1_text_words_batch_post_request import V1TextWordsBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextWordsBatchPostRequest from a JSON string
v1_text_words_batch_post_request_instance = V1TextWordsBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextWordsBatchPostRequest.to_json())

# convert the object into a dict
v1_text_words_batch_post_request_dict = v1_text_words_batch_post_request_instance.to_dict()
# create an instance of V1TextWordsBatchPostRequest from a dict
v1_text_words_batch_post_request_from_dict = V1TextWordsBatchPostRequest.from_dict(v1_text_words_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


