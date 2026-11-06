# OpenapiClient::V1NetworkingIpIpGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | The IP address that was looked up | [optional] |
| **country** | **String** | Country name where the IP is located | [optional] |
| **country_code** | **String** | Two-letter ISO country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) | [optional] |
| **city** | **String** | City name where the IP is located | [optional] |
| **isp** | **String** | Internet Service Provider providing the IP | [optional] |
| **is_vpn** | **Boolean** | True when the IP belongs to a known VPN | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingIpIpGet200ResponseData.new(
  ip: null,
  country: null,
  country_code: null,
  city: null,
  isp: null,
  is_vpn: null
)
```

