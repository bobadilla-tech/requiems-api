# V1TextSentimentPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sentiment** | **string** | The dominant sentiment class: positive, negative, or neutral | [optional] [default to undefined]
**score** | **number** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] [default to undefined]
**breakdown_positive** | **number** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] [default to undefined]
**breakdown_negative** | **number** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] [default to undefined]
**breakdown_neutral** | **number** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] [default to undefined]

## Example

```typescript
import { V1TextSentimentPost200ResponseData } from 'api-client';

const instance: V1TextSentimentPost200ResponseData = {
    sentiment,
    score,
    breakdown_positive,
    breakdown_negative,
    breakdown_neutral,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
