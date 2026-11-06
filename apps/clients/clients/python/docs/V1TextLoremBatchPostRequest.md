# V1TextLoremBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **List[object]** | Array of text generation requests. Minimum 1 item, maximum 50 items. | 
**items_paragraphs** | **int** | Number of paragraphs to generate (1-20) | [optional] 
**items_sentences** | **int** | Number of sentences per paragraph (1-20) | [optional] 

## Example

```python
from openapi_client.models.v1_text_lorem_batch_post_request import V1TextLoremBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextLoremBatchPostRequest from a JSON string
v1_text_lorem_batch_post_request_instance = V1TextLoremBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextLoremBatchPostRequest.to_json())

# convert the object into a dict
v1_text_lorem_batch_post_request_dict = v1_text_lorem_batch_post_request_instance.to_dict()
# create an instance of V1TextLoremBatchPostRequest from a dict
v1_text_lorem_batch_post_request_from_dict = V1TextLoremBatchPostRequest.from_dict(v1_text_lorem_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


