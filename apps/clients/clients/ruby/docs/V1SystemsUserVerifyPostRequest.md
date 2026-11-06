# OpenapiClient::V1SystemsUserVerifyPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **email** | **String** | Email address to verify. Domain-level WHOIS, MX, and DNS checks are run automatically. |  |
| **ip_address** | **String** | Optional IPv4 or IPv6 address. When provided, VPN and proxy signals are added to the result. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsUserVerifyPostRequest.new(
  email: user@example.com,
  ip_address: 45.33.32.156
)
```

