# V1NetworkingDomainDomainGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **str** | The domain that was looked up | [optional] 
**available** | **bool** | True when the domain appears to be unregistered (NS lookup returns NXDOMAIN). False when name servers are delegated. | [optional] 
**dns_a** | **str** | IPv4 addresses (A records) | [optional] 
**dns_aaaa** | **str** | IPv6 addresses (AAAA records) | [optional] 
**dns_mx** | **str** | Mail exchange records, each with host and priority fields | [optional] 
**dns_ns** | **str** | Authoritative name server hostnames | [optional] 
**dns_txt** | **str** | TXT record values (SPF, DKIM, verification tokens, etc.) | [optional] 
**dns_cname** | **str** | CNAME alias target, if the domain is an alias. Empty string when no alias exists. | [optional] 

## Example

```python
from openapi_client.models.v1_networking_domain_domain_get200_response_data import V1NetworkingDomainDomainGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingDomainDomainGet200ResponseData from a JSON string
v1_networking_domain_domain_get200_response_data_instance = V1NetworkingDomainDomainGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingDomainDomainGet200ResponseData.to_json())

# convert the object into a dict
v1_networking_domain_domain_get200_response_data_dict = v1_networking_domain_domain_get200_response_data_instance.to_dict()
# create an instance of V1NetworkingDomainDomainGet200ResponseData from a dict
v1_networking_domain_domain_get200_response_data_from_dict = V1NetworkingDomainDomainGet200ResponseData.from_dict(v1_networking_domain_domain_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


