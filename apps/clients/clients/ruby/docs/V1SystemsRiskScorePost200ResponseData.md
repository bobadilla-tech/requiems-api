# OpenapiClient::V1SystemsRiskScorePost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **risk_score** | **String** | Composite risk score from 0.0 to 1.0 | [optional] |
| **is_safe** | **Boolean** | True when the risk score is below the safety threshold | [optional] |
| **confidence** | **String** | Confidence in the score based on resolved signals | [optional] |
| **flags** | **Array&lt;Object&gt;** | Array of triggered risk flags | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsRiskScorePost200ResponseData.new(
  risk_score: null,
  is_safe: null,
  confidence: null,
  flags: null
)
```

