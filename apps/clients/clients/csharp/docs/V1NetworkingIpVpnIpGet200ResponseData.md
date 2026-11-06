# Org.OpenAPITools.Model.V1NetworkingIpVpnIpGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | The analyzed IP address | [optional] 
**IsVpn** | **bool** | True when the IP belongs to a known VPN provider | [optional] 
**IsProxy** | **bool** | True when the IP is a known public or web proxy | [optional] 
**IsTor** | **bool** | True when the IP is a known Tor exit node | [optional] 
**IsHosting** | **bool** | True when the IP belongs to a data-centre or hosting provider (DCH) | [optional] 
**Score** | **int** | Raw threat score (0-9+). Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1 | [optional] 
**Threat** | **int** | Threat level derived from score: 0 &#x3D; None, 1 &#x3D; Low, 2–3 &#x3D; Medium, 4–5 &#x3D; High, 6+ &#x3D; Critical | [optional] 
**FraudScore** | **int** | Fraud risk score from 0 (no risk) to 100 (high risk). Available when using IP2Proxy PX5 or higher | [optional] 
**AsnOrg** | **string** | Organization name owning the Autonomous System containing the IP (e.g. \&quot;DIGITALOCEAN-ASN\&quot;) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

