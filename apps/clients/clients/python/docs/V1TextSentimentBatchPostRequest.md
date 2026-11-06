# V1TextSentimentBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**texts** | **List[object]** | The list of texts to analyze. Between 1 and 50 items. | 

## Example

```python
from openapi_client.models.v1_text_sentiment_batch_post_request import V1TextSentimentBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSentimentBatchPostRequest from a JSON string
v1_text_sentiment_batch_post_request_instance = V1TextSentimentBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TextSentimentBatchPostRequest.to_json())

# convert the object into a dict
v1_text_sentiment_batch_post_request_dict = v1_text_sentiment_batch_post_request_instance.to_dict()
# create an instance of V1TextSentimentBatchPostRequest from a dict
v1_text_sentiment_batch_post_request_from_dict = V1TextSentimentBatchPostRequest.from_dict(v1_text_sentiment_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


