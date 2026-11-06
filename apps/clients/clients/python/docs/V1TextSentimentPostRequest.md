# V1TextSentimentPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**text** | **str** | The text to analyze. | 

## Example

```python
from openapi_client.models.v1_text_sentiment_post_request import V1TextSentimentPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSentimentPostRequest from a JSON string
v1_text_sentiment_post_request_instance = V1TextSentimentPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextSentimentPostRequest.to_json())

# convert the object into a dict
v1_text_sentiment_post_request_dict = v1_text_sentiment_post_request_instance.to_dict()
# create an instance of V1TextSentimentPostRequest from a dict
v1_text_sentiment_post_request_from_dict = V1TextSentimentPostRequest.from_dict(v1_text_sentiment_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


