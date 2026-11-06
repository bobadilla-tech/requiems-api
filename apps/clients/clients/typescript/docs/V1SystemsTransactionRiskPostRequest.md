# V1SystemsTransactionRiskPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**card_bin** | **string** | Card BIN (first 6–8 digits) to look up card issuing country. | [default to undefined]
**ip_address** | **string** | Customer\&#39;s IPv4 or IPv6 address. Must be a valid IP. | [default to undefined]
**billing_country** | **string** | ISO 3166-1 alpha-2 billing country code. When provided, country mismatch checks are more precise. | [optional] [default to undefined]
**amount_usd** | **number** | Transaction amount in USD. High-value transactions via VPN incur an additional risk score boost. | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsTransactionRiskPostRequest } from 'api-client';

const instance: V1SystemsTransactionRiskPostRequest = {
    card_bin,
    ip_address,
    billing_country,
    amount_usd,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
