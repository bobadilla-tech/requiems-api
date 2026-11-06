# OpenapiClient::V1TextSentimentBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Sentiment results for each input text, in the same order as the input. | [optional] |
| **results_sentiment** | **String** | The dominant sentiment class: positive, negative, or neutral | [optional] |
| **results_score** | **Float** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] |
| **results_breakdown_positive** | **Float** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] |
| **results_breakdown_negative** | **Float** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] |
| **results_breakdown_neutral** | **Float** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] |
| **total** | **Float** | Total number of results returned. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSentimentBatchPost200ResponseData.new(
  results: null,
  results_sentiment: null,
  results_score: null,
  results_breakdown_positive: null,
  results_breakdown_negative: null,
  results_breakdown_neutral: null,
  total: null
)
```

