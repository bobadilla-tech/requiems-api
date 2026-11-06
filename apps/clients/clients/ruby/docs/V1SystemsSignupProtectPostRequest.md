# OpenapiClient::V1SystemsSignupProtectPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **email** | **String** | Email address to validate and score. At least one of email, phone, or ip_address is required. | [optional] |
| **phone** | **String** | Phone number in E.164 format to validate and check for VoIP or virtual numbers. | [optional] |
| **ip_address** | **String** | IPv4 or IPv6 address to check for VPN, proxy, TOR, and hosting status. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsSignupProtectPostRequest.new(
  email: user@tempmail.io,
  phone: +14155552671,
  ip_address: 45.33.32.156
)
```

