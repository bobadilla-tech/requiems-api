# V1NetworkingDisposableBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Array of check results for each email | [optional] 
**total** | **int** | Total number of emails checked | [optional] 

## Example

```python
from openapi_client.models.v1_networking_disposable_batch_post200_response_data import V1NetworkingDisposableBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingDisposableBatchPost200ResponseData from a JSON string
v1_networking_disposable_batch_post200_response_data_instance = V1NetworkingDisposableBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingDisposableBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_networking_disposable_batch_post200_response_data_dict = v1_networking_disposable_batch_post200_response_data_instance.to_dict()
# create an instance of V1NetworkingDisposableBatchPost200ResponseData from a dict
v1_networking_disposable_batch_post200_response_data_from_dict = V1NetworkingDisposableBatchPost200ResponseData.from_dict(v1_networking_disposable_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


