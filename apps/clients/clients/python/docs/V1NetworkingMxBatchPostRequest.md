# V1NetworkingMxBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domains** | **List[object]** | Array of domains to lookup (min: 1, max: 50). | 

## Example

```python
from openapi_client.models.v1_networking_mx_batch_post_request import V1NetworkingMxBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingMxBatchPostRequest from a JSON string
v1_networking_mx_batch_post_request_instance = V1NetworkingMxBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingMxBatchPostRequest.to_json())

# convert the object into a dict
v1_networking_mx_batch_post_request_dict = v1_networking_mx_batch_post_request_instance.to_dict()
# create an instance of V1NetworkingMxBatchPostRequest from a dict
v1_networking_mx_batch_post_request_from_dict = V1NetworkingMxBatchPostRequest.from_dict(v1_networking_mx_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


