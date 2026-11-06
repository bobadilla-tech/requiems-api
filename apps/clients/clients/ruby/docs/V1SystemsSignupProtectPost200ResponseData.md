# OpenapiClient::V1SystemsSignupProtectPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **risk_score** | **String** | Composite risk score from 0.0 (clean) to 1.0 (high risk) | [optional] |
| **is_safe** | **Boolean** | True when risk_score is below the safety threshold with no critical flags | [optional] |
| **confidence** | **String** | Confidence in the decision from 0.0 to 1.0, based on which signals resolved successfully | [optional] |
| **flags** | **Array&lt;Object&gt;** | Array of triggered risk flags. Possible values: disposable_email, vpn_detected, proxy_detected, tor_detected, is_hosting | [optional] |
| **signals_email** | **Object** | Email validation result including disposable detection and MX check | [optional] |
| **signals_phone** | **Object** | Phone number validation result with country code and VoIP/virtual detection | [optional] |
| **signals_ip** | **Object** | IP intelligence including VPN/proxy/TOR status and fraud score | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsSignupProtectPost200ResponseData.new(
  risk_score: null,
  is_safe: null,
  confidence: null,
  flags: null,
  signals_email: null,
  signals_phone: null,
  signals_ip: null
)
```

