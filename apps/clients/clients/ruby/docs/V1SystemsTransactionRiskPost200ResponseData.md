# OpenapiClient::V1SystemsTransactionRiskPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **risk_score** | **String** | Transaction risk score from 0.0 to 1.0 | [optional] |
| **is_safe** | **Boolean** | True when score is below 0.5 and no TOR or proxy detected | [optional] |
| **flags** | **Array&lt;Object&gt;** | Triggered flags. Possible values: ip_country_mismatch, bin_country_mismatch, country_mismatch, vpn_detected, proxy_detected, tor_detected, high_value_vpn | [optional] |
| **signals_ip_country** | **String** | Country code resolved from the IP address | [optional] |
| **signals_billing_country** | **String** | Billing country provided in the request, or null if omitted | [optional] |
| **signals_bin_country** | **String** | Country of the card issuer resolved from the BIN | [optional] |
| **signals_country_mismatch** | **Boolean** | True when any country comparison produced a mismatch | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsTransactionRiskPost200ResponseData.new(
  risk_score: null,
  is_safe: null,
  flags: null,
  signals_ip_country: null,
  signals_billing_country: null,
  signals_bin_country: null,
  signals_country_mismatch: null
)
```

