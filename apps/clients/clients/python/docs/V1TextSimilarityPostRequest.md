# V1TextSimilarityPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text1** | **str** | The first text to compare. | 
**text2** | **str** | The second text to compare. | 

## Example

```python
from openapi_client.models.v1_text_similarity_post_request import V1TextSimilarityPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSimilarityPostRequest from a JSON string
v1_text_similarity_post_request_instance = V1TextSimilarityPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextSimilarityPostRequest.to_json())

# convert the object into a dict
v1_text_similarity_post_request_dict = v1_text_similarity_post_request_instance.to_dict()
# create an instance of V1TextSimilarityPostRequest from a dict
v1_text_similarity_post_request_from_dict = V1TextSimilarityPostRequest.from_dict(v1_text_similarity_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


