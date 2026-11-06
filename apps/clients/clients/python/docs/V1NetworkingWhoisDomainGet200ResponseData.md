# V1NetworkingWhoisDomainGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **str** | The domain name that was looked up | [optional] 
**registrar** | **str** | The name of the registrar holding the domain registration | [optional] 
**name_servers** | **str** | List of authoritative name servers for the domain | [optional] 
**status** | **str** | EPP status codes for the domain (e.g. clientTransferProhibited) | [optional] 
**created_date** | **str** | Date the domain was first registered (ISO 8601) | [optional] 
**updated_date** | **str** | Date the domain record was last updated (ISO 8601) | [optional] 
**expiry_date** | **str** | Date the domain registration expires (ISO 8601) | [optional] 
**dnssec** | **bool** | True when DNSSEC is enabled for the domain | [optional] 

## Example

```python
from openapi_client.models.v1_networking_whois_domain_get200_response_data import V1NetworkingWhoisDomainGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingWhoisDomainGet200ResponseData from a JSON string
v1_networking_whois_domain_get200_response_data_instance = V1NetworkingWhoisDomainGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingWhoisDomainGet200ResponseData.to_json())

# convert the object into a dict
v1_networking_whois_domain_get200_response_data_dict = v1_networking_whois_domain_get200_response_data_instance.to_dict()
# create an instance of V1NetworkingWhoisDomainGet200ResponseData from a dict
v1_networking_whois_domain_get200_response_data_from_dict = V1NetworkingWhoisDomainGet200ResponseData.from_dict(v1_networking_whois_domain_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


