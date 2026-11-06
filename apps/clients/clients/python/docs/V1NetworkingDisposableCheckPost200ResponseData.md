# V1NetworkingDisposableCheckPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **str** | The email address that was checked | [optional] 
**is_disposable** | **bool** | Whether the email uses a disposable domain | [optional] 
**domain** | **str** | The domain part of the email address | [optional] 

## Example

```python
from openapi_client.models.v1_networking_disposable_check_post200_response_data import V1NetworkingDisposableCheckPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingDisposableCheckPost200ResponseData from a JSON string
v1_networking_disposable_check_post200_response_data_instance = V1NetworkingDisposableCheckPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingDisposableCheckPost200ResponseData.to_json())

# convert the object into a dict
v1_networking_disposable_check_post200_response_data_dict = v1_networking_disposable_check_post200_response_data_instance.to_dict()
# create an instance of V1NetworkingDisposableCheckPost200ResponseData from a dict
v1_networking_disposable_check_post200_response_data_from_dict = V1NetworkingDisposableCheckPost200ResponseData.from_dict(v1_networking_disposable_check_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


