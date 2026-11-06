# OpenapiClient::V1SystemsRiskScorePostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **email** | **String** | Email address to include in the risk score. | [optional] |
| **phone** | **String** | Phone number in E.164 format. | [optional] |
| **ip_address** | **String** | IPv4 or IPv6 address to check. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsRiskScorePostRequest.new(
  email: user@example.com,
  phone: +14155552671,
  ip_address: 45.33.32.156
)
```

