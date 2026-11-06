# V1NetworkingDisposableDomainsGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domains** | **List[object]** | Array of domain names | [optional] 
**total** | **int** | Total number of domains in blocklist | [optional] 
**page** | **int** | Current page number | [optional] 
**per_page** | **int** | Number of items per page | [optional] 
**has_more** | **bool** | Whether there are more pages available | [optional] 

## Example

```python
from openapi_client.models.v1_networking_disposable_domains_get200_response_data import V1NetworkingDisposableDomainsGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingDisposableDomainsGet200ResponseData from a JSON string
v1_networking_disposable_domains_get200_response_data_instance = V1NetworkingDisposableDomainsGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingDisposableDomainsGet200ResponseData.to_json())

# convert the object into a dict
v1_networking_disposable_domains_get200_response_data_dict = v1_networking_disposable_domains_get200_response_data_instance.to_dict()
# create an instance of V1NetworkingDisposableDomainsGet200ResponseData from a dict
v1_networking_disposable_domains_get200_response_data_from_dict = V1NetworkingDisposableDomainsGet200ResponseData.from_dict(v1_networking_disposable_domains_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


