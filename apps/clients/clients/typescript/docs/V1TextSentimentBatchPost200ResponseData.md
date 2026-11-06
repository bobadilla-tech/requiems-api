# V1TextSentimentBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Sentiment results for each input text, in the same order as the input. | [optional] [default to undefined]
**results_sentiment** | **string** | The dominant sentiment class: positive, negative, or neutral | [optional] [default to undefined]
**results_score** | **number** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] [default to undefined]
**results_breakdown_positive** | **number** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] [default to undefined]
**results_breakdown_negative** | **number** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] [default to undefined]
**results_breakdown_neutral** | **number** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] [default to undefined]
**total** | **number** | Total number of results returned. | [optional] [default to undefined]

## Example

```typescript
import { V1TextSentimentBatchPost200ResponseData } from 'api-client';

const instance: V1TextSentimentBatchPost200ResponseData = {
    results,
    results_sentiment,
    results_score,
    results_breakdown_positive,
    results_breakdown_negative,
    results_breakdown_neutral,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
