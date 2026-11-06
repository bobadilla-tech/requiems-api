# V1NetworkingIpAsnGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **str** | The IP address that was looked up (the requesting client&#39;s IP) | [optional] 
**asn** | **str** | Autonomous System Number in format \&quot;ASxxxx\&quot; (e.g., \&quot;AS15169\&quot;) | [optional] 
**org** | **str** | Organization name owning the IP address range | [optional] 
**isp** | **str** | Internet Service Provider providing the IP | [optional] 
**domain** | **str** | Domain name associated with the IP or IP range | [optional] 
**route** | **str** | CIDR notation of the network route (e.g., \&quot;8.8.8.0/24\&quot;) | [optional] 
**type** | **str** | Type of network (e.g., \&quot;hosting\&quot;, \&quot;isp\&quot;, \&quot;business\&quot;, \&quot;cdn\&quot;) | [optional] 

## Example

```python
from openapi_client.models.v1_networking_ip_asn_get200_response_data import V1NetworkingIpAsnGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingIpAsnGet200ResponseData from a JSON string
v1_networking_ip_asn_get200_response_data_instance = V1NetworkingIpAsnGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingIpAsnGet200ResponseData.to_json())

# convert the object into a dict
v1_networking_ip_asn_get200_response_data_dict = v1_networking_ip_asn_get200_response_data_instance.to_dict()
# create an instance of V1NetworkingIpAsnGet200ResponseData from a dict
v1_networking_ip_asn_get200_response_data_from_dict = V1NetworkingIpAsnGet200ResponseData.from_dict(v1_networking_ip_asn_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


