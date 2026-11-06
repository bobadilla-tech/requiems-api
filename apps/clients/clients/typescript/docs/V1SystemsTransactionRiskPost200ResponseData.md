# V1SystemsTransactionRiskPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**risk_score** | **string** | Transaction risk score from 0.0 to 1.0 | [optional] [default to undefined]
**is_safe** | **boolean** | True when score is below 0.5 and no TOR or proxy detected | [optional] [default to undefined]
**flags** | **Array&lt;any&gt;** | Triggered flags. Possible values: ip_country_mismatch, bin_country_mismatch, country_mismatch, vpn_detected, proxy_detected, tor_detected, high_value_vpn | [optional] [default to undefined]
**signals_ip_country** | **string** | Country code resolved from the IP address | [optional] [default to undefined]
**signals_billing_country** | **string** | Billing country provided in the request, or null if omitted | [optional] [default to undefined]
**signals_bin_country** | **string** | Country of the card issuer resolved from the BIN | [optional] [default to undefined]
**signals_country_mismatch** | **boolean** | True when any country comparison produced a mismatch | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsTransactionRiskPost200ResponseData } from 'api-client';

const instance: V1SystemsTransactionRiskPost200ResponseData = {
    risk_score,
    is_safe,
    flags,
    signals_ip_country,
    signals_billing_country,
    signals_bin_country,
    signals_country_mismatch,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
