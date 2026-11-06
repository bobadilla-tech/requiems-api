# Org.OpenAPITools.Model.V1TextSentimentBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Sentiment results for each input text, in the same order as the input. | [optional] 
**ResultsSentiment** | **string** | The dominant sentiment class: positive, negative, or neutral | [optional] 
**ResultsScore** | **decimal** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] 
**ResultsBreakdownPositive** | **decimal** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] 
**ResultsBreakdownNegative** | **decimal** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] 
**ResultsBreakdownNeutral** | **decimal** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] 
**Total** | **decimal** | Total number of results returned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

