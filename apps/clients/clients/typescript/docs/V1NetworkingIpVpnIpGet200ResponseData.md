# V1NetworkingIpVpnIpGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **string** | The analyzed IP address | [optional] [default to undefined]
**is_vpn** | **boolean** | True when the IP belongs to a known VPN provider | [optional] [default to undefined]
**is_proxy** | **boolean** | True when the IP is a known public or web proxy | [optional] [default to undefined]
**is_tor** | **boolean** | True when the IP is a known Tor exit node | [optional] [default to undefined]
**is_hosting** | **boolean** | True when the IP belongs to a data-centre or hosting provider (DCH) | [optional] [default to undefined]
**score** | **number** | Raw threat score (0-9+). Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1 | [optional] [default to undefined]
**threat** | **number** | Threat level derived from score: 0 &#x3D; None, 1 &#x3D; Low, 2–3 &#x3D; Medium, 4–5 &#x3D; High, 6+ &#x3D; Critical | [optional] [default to undefined]
**fraud_score** | **number** | Fraud risk score from 0 (no risk) to 100 (high risk). Available when using IP2Proxy PX5 or higher | [optional] [default to undefined]
**asn_org** | **string** | Organization name owning the Autonomous System containing the IP (e.g. \&quot;DIGITALOCEAN-ASN\&quot;) | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingIpVpnIpGet200ResponseData } from 'api-client';

const instance: V1NetworkingIpVpnIpGet200ResponseData = {
    ip,
    is_vpn,
    is_proxy,
    is_tor,
    is_hosting,
    score,
    threat,
    fraud_score,
    asn_org,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
