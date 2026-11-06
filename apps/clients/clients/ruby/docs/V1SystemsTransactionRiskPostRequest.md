# OpenapiClient::V1SystemsTransactionRiskPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **card_bin** | **String** | Card BIN (first 6–8 digits) to look up card issuing country. |  |
| **ip_address** | **String** | Customer&#39;s IPv4 or IPv6 address. Must be a valid IP. |  |
| **billing_country** | **String** | ISO 3166-1 alpha-2 billing country code. When provided, country mismatch checks are more precise. | [optional] |
| **amount_usd** | **Float** | Transaction amount in USD. High-value transactions via VPN incur an additional risk score boost. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsTransactionRiskPostRequest.new(
  card_bin: 424242,
  ip_address: 92.168.1.1,
  billing_country: DE,
  amount_usd: 250
)
```

