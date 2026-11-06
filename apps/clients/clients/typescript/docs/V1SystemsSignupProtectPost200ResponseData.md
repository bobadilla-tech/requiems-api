# V1SystemsSignupProtectPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**risk_score** | **string** | Composite risk score from 0.0 (clean) to 1.0 (high risk) | [optional] [default to undefined]
**is_safe** | **boolean** | True when risk_score is below the safety threshold with no critical flags | [optional] [default to undefined]
**confidence** | **string** | Confidence in the decision from 0.0 to 1.0, based on which signals resolved successfully | [optional] [default to undefined]
**flags** | **Array&lt;any&gt;** | Array of triggered risk flags. Possible values: disposable_email, vpn_detected, proxy_detected, tor_detected, is_hosting | [optional] [default to undefined]
**signals_email** | **object** | Email validation result including disposable detection and MX check | [optional] [default to undefined]
**signals_phone** | **object** | Phone number validation result with country code and VoIP/virtual detection | [optional] [default to undefined]
**signals_ip** | **object** | IP intelligence including VPN/proxy/TOR status and fraud score | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsSignupProtectPost200ResponseData } from 'api-client';

const instance: V1SystemsSignupProtectPost200ResponseData = {
    risk_score,
    is_safe,
    confidence,
    flags,
    signals_email,
    signals_phone,
    signals_ip,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
