# V1NetworkingWhoisBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | One entry per domain, in the same order as the input | [optional] 
**results_domain** | **str** | Domain name requested in the batch | [optional] 
**results_found** | **bool** | False when WHOIS data could not be found | [optional] 
**results_error** | **str** | Error message when found is false | [optional] 
**results_data** | **object** | WHOIS information for the domain when found is true | [optional] 
**total** | **int** | Total number of results returned | [optional] 

## Example

```python
from openapi_client.models.v1_networking_whois_batch_post200_response_data import V1NetworkingWhoisBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingWhoisBatchPost200ResponseData from a JSON string
v1_networking_whois_batch_post200_response_data_instance = V1NetworkingWhoisBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingWhoisBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_networking_whois_batch_post200_response_data_dict = v1_networking_whois_batch_post200_response_data_instance.to_dict()
# create an instance of V1NetworkingWhoisBatchPost200ResponseData from a dict
v1_networking_whois_batch_post200_response_data_from_dict = V1NetworkingWhoisBatchPost200ResponseData.from_dict(v1_networking_whois_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


