# OpenapiClient::V1TextSentimentPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **sentiment** | **String** | The dominant sentiment class: positive, negative, or neutral | [optional] |
| **score** | **Float** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] |
| **breakdown_positive** | **Float** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] |
| **breakdown_negative** | **Float** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] |
| **breakdown_neutral** | **Float** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextSentimentPost200ResponseData.new(
  sentiment: null,
  score: null,
  breakdown_positive: null,
  breakdown_negative: null,
  breakdown_neutral: null
)
```

