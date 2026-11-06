# V1NetworkingIpIpGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **str** | The IP address that was looked up | [optional] 
**country** | **str** | Country name where the IP is located | [optional] 
**country_code** | **str** | Two-letter ISO country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) | [optional] 
**city** | **str** | City name where the IP is located | [optional] 
**isp** | **str** | Internet Service Provider providing the IP | [optional] 
**is_vpn** | **bool** | True when the IP belongs to a known VPN | [optional] 

## Example

```python
from openapi_client.models.v1_networking_ip_ip_get200_response_data import V1NetworkingIpIpGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingIpIpGet200ResponseData from a JSON string
v1_networking_ip_ip_get200_response_data_instance = V1NetworkingIpIpGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingIpIpGet200ResponseData.to_json())

# convert the object into a dict
v1_networking_ip_ip_get200_response_data_dict = v1_networking_ip_ip_get200_response_data_instance.to_dict()
# create an instance of V1NetworkingIpIpGet200ResponseData from a dict
v1_networking_ip_ip_get200_response_data_from_dict = V1NetworkingIpIpGet200ResponseData.from_dict(v1_networking_ip_ip_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


