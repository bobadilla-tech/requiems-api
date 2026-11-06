# V1SystemsUserVerifyPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**verified** | **boolean** | True when risk score is below 0.3, confidence is above 0.5, email is valid, domain has MX, and domain is registered | [optional] [default to undefined]
**confidence** | **string** | Confidence in the verification result, based on how many signals resolved | [optional] [default to undefined]
**risk_score** | **string** | Composite risk score from 0.0 to 1.0 | [optional] [default to undefined]
**flags** | **Array&lt;any&gt;** | Triggered risk flags. Possible values: email_invalid, disposable_email, no_mx, domain_not_registered, young_domain, whois_unavailable, ip_risk | [optional] [default to undefined]
**signals_email** | **object** | Email validation result | [optional] [default to undefined]
**signals_domain** | **object** | Domain intelligence including WHOIS age, MX and A record presence, and registration status | [optional] [default to undefined]
**signals_ip** | **string** | IP risk signals. Null when ip_address was not provided | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsUserVerifyPost200ResponseData } from 'api-client';

const instance: V1SystemsUserVerifyPost200ResponseData = {
    verified,
    confidence,
    risk_score,
    flags,
    signals_email,
    signals_domain,
    signals_ip,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
