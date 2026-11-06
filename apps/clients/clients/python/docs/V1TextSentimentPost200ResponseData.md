# V1TextSentimentPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sentiment** | **str** | The dominant sentiment class: positive, negative, or neutral | [optional] 
**score** | **float** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] 
**breakdown_positive** | **float** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] 
**breakdown_negative** | **float** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] 
**breakdown_neutral** | **float** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] 

## Example

```python
from openapi_client.models.v1_text_sentiment_post200_response_data import V1TextSentimentPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TextSentimentPost200ResponseData from a JSON string
v1_text_sentiment_post200_response_data_instance = V1TextSentimentPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TextSentimentPost200ResponseData.to_json())

# convert the object into a dict
v1_text_sentiment_post200_response_data_dict = v1_text_sentiment_post200_response_data_instance.to_dict()
# create an instance of V1TextSentimentPost200ResponseData from a dict
v1_text_sentiment_post200_response_data_from_dict = V1TextSentimentPost200ResponseData.from_dict(v1_text_sentiment_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


