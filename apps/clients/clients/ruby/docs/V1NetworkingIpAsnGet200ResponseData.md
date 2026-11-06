# OpenapiClient::V1NetworkingIpAsnGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | The IP address that was looked up (the requesting client&#39;s IP) | [optional] |
| **asn** | **String** | Autonomous System Number in format \&quot;ASxxxx\&quot; (e.g., \&quot;AS15169\&quot;) | [optional] |
| **org** | **String** | Organization name owning the IP address range | [optional] |
| **isp** | **String** | Internet Service Provider providing the IP | [optional] |
| **domain** | **String** | Domain name associated with the IP or IP range | [optional] |
| **route** | **String** | CIDR notation of the network route (e.g., \&quot;8.8.8.0/24\&quot;) | [optional] |
| **type** | **String** | Type of network (e.g., \&quot;hosting\&quot;, \&quot;isp\&quot;, \&quot;business\&quot;, \&quot;cdn\&quot;) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingIpAsnGet200ResponseData.new(
  ip: null,
  asn: null,
  org: null,
  isp: null,
  domain: null,
  route: null,
  type: null
)
```

