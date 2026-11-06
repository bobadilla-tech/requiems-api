# V1NetworkingMxDomainGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **str** | The domain that was queried | [optional] 
**records** | **List[object]** | List of MX records, sorted by priority ascending (lowest number &#x3D; highest priority) | [optional] 
**records_host** | **str** | Fully-qualified hostname of the mail server (typically ends with a trailing dot) | [optional] 
**records_priority** | **int** | MX priority value. Lower values have higher delivery priority per RFC 5321. | [optional] 

## Example

```python
from openapi_client.models.v1_networking_mx_domain_get200_response_data import V1NetworkingMxDomainGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingMxDomainGet200ResponseData from a JSON string
v1_networking_mx_domain_get200_response_data_instance = V1NetworkingMxDomainGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingMxDomainGet200ResponseData.to_json())

# convert the object into a dict
v1_networking_mx_domain_get200_response_data_dict = v1_networking_mx_domain_get200_response_data_instance.to_dict()
# create an instance of V1NetworkingMxDomainGet200ResponseData from a dict
v1_networking_mx_domain_get200_response_data_from_dict = V1NetworkingMxDomainGet200ResponseData.from_dict(v1_networking_mx_domain_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


