# V1TextSimilarityPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**similarity** | **float** | Cosine similarity score between the two texts, in the range [0, 1]. | [optional] 
**method** | **str** | The algorithm used. Currently always &#39;cosine&#39;. | [optional] 

## Example

```python
from openapi_client.models.v1_text_similarity_post200_response_data import V1TextSimilarityPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSimilarityPost200ResponseData from a JSON string
v1_text_similarity_post200_response_data_instance = V1TextSimilarityPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextSimilarityPost200ResponseData.to_json())

# convert the object into a dict
v1_text_similarity_post200_response_data_dict = v1_text_similarity_post200_response_data_instance.to_dict()
# create an instance of V1TextSimilarityPost200ResponseData from a dict
v1_text_similarity_post200_response_data_from_dict = V1TextSimilarityPost200ResponseData.from_dict(v1_text_similarity_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


