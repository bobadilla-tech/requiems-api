# OpenapiClient::V1NetworkingIpVpnIpGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | The analyzed IP address | [optional] |
| **is_vpn** | **Boolean** | True when the IP belongs to a known VPN provider | [optional] |
| **is_proxy** | **Boolean** | True when the IP is a known public or web proxy | [optional] |
| **is_tor** | **Boolean** | True when the IP is a known Tor exit node | [optional] |
| **is_hosting** | **Boolean** | True when the IP belongs to a data-centre or hosting provider (DCH) | [optional] |
| **score** | **Integer** | Raw threat score (0-9+). Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1 | [optional] |
| **threat** | **Integer** | Threat level derived from score: 0 &#x3D; None, 1 &#x3D; Low, 2–3 &#x3D; Medium, 4–5 &#x3D; High, 6+ &#x3D; Critical | [optional] |
| **fraud_score** | **Integer** | Fraud risk score from 0 (no risk) to 100 (high risk). Available when using IP2Proxy PX5 or higher | [optional] |
| **asn_org** | **String** | Organization name owning the Autonomous System containing the IP (e.g. \&quot;DIGITALOCEAN-ASN\&quot;) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingIpVpnIpGet200ResponseData.new(
  ip: null,
  is_vpn: null,
  is_proxy: null,
  is_tor: null,
  is_hosting: null,
  score: null,
  threat: null,
  fraud_score: null,
  asn_org: null
)
```

