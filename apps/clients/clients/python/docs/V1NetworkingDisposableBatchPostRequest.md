# V1NetworkingDisposableBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**emails** | **List[object]** | Array of email addresses to check (max 100) | 

## Example

```python
from openapi_client.models.v1_networking_disposable_batch_post_request import V1NetworkingDisposableBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingDisposableBatchPostRequest from a JSON string
v1_networking_disposable_batch_post_request_instance = V1NetworkingDisposableBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingDisposableBatchPostRequest.to_json())

# convert the object into a dict
v1_networking_disposable_batch_post_request_dict = v1_networking_disposable_batch_post_request_instance.to_dict()
# create an instance of V1NetworkingDisposableBatchPostRequest from a dict
v1_networking_disposable_batch_post_request_from_dict = V1NetworkingDisposableBatchPostRequest.from_dict(v1_networking_disposable_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


