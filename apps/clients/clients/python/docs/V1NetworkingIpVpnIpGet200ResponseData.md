# V1NetworkingIpVpnIpGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **str** | The analyzed IP address | [optional] 
**is_vpn** | **bool** | True when the IP belongs to a known VPN provider | [optional] 
**is_proxy** | **bool** | True when the IP is a known public or web proxy | [optional] 
**is_tor** | **bool** | True when the IP is a known Tor exit node | [optional] 
**is_hosting** | **bool** | True when the IP belongs to a data-centre or hosting provider (DCH) | [optional] 
**score** | **int** | Raw threat score (0-9+). Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1 | [optional] 
**threat** | **int** | Threat level derived from score: 0 &#x3D; None, 1 &#x3D; Low, 2–3 &#x3D; Medium, 4–5 &#x3D; High, 6+ &#x3D; Critical | [optional] 
**fraud_score** | **int** | Fraud risk score from 0 (no risk) to 100 (high risk). Available when using IP2Proxy PX5 or higher | [optional] 
**asn_org** | **str** | Organization name owning the Autonomous System containing the IP (e.g. \&quot;DIGITALOCEAN-ASN\&quot;) | [optional] 

## Example

```python
from openapi_client.models.v1_networking_ip_vpn_ip_get200_response_data import V1NetworkingIpVpnIpGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingIpVpnIpGet200ResponseData from a JSON string
v1_networking_ip_vpn_ip_get200_response_data_instance = V1NetworkingIpVpnIpGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingIpVpnIpGet200ResponseData.to_json())

# convert the object into a dict
v1_networking_ip_vpn_ip_get200_response_data_dict = v1_networking_ip_vpn_ip_get200_response_data_instance.to_dict()
# create an instance of V1NetworkingIpVpnIpGet200ResponseData from a dict
v1_networking_ip_vpn_ip_get200_response_data_from_dict = V1NetworkingIpVpnIpGet200ResponseData.from_dict(v1_networking_ip_vpn_ip_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


