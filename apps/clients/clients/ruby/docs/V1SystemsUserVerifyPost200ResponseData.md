# OpenapiClient::V1SystemsUserVerifyPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **verified** | **Boolean** | True when risk score is below 0.3, confidence is above 0.5, email is valid, domain has MX, and domain is registered | [optional] |
| **confidence** | **String** | Confidence in the verification result, based on how many signals resolved | [optional] |
| **risk_score** | **String** | Composite risk score from 0.0 to 1.0 | [optional] |
| **flags** | **Array&lt;Object&gt;** | Triggered risk flags. Possible values: email_invalid, disposable_email, no_mx, domain_not_registered, young_domain, whois_unavailable, ip_risk | [optional] |
| **signals_email** | **Object** | Email validation result | [optional] |
| **signals_domain** | **Object** | Domain intelligence including WHOIS age, MX and A record presence, and registration status | [optional] |
| **signals_ip** | **String** | IP risk signals. Null when ip_address was not provided | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsUserVerifyPost200ResponseData.new(
  verified: null,
  confidence: null,
  risk_score: null,
  flags: null,
  signals_email: null,
  signals_domain: null,
  signals_ip: null
)
```

