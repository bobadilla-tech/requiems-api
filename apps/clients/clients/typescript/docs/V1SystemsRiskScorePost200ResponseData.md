# V1SystemsRiskScorePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**risk_score** | **string** | Composite risk score from 0.0 to 1.0 | [optional] [default to undefined]
**is_safe** | **boolean** | True when the risk score is below the safety threshold | [optional] [default to undefined]
**confidence** | **string** | Confidence in the score based on resolved signals | [optional] [default to undefined]
**flags** | **Array&lt;any&gt;** | Array of triggered risk flags | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsRiskScorePost200ResponseData } from 'api-client';

const instance: V1SystemsRiskScorePost200ResponseData = {
    risk_score,
    is_safe,
    confidence,
    flags,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
