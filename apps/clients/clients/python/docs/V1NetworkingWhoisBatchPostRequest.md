# V1NetworkingWhoisBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domains** | **List[object]** | Array of domain names. Min: 1, Max: 50. | 

## Example

```python
from openapi_client.models.v1_networking_whois_batch_post_request import V1NetworkingWhoisBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingWhoisBatchPostRequest from a JSON string
v1_networking_whois_batch_post_request_instance = V1NetworkingWhoisBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingWhoisBatchPostRequest.to_json())

# convert the object into a dict
v1_networking_whois_batch_post_request_dict = v1_networking_whois_batch_post_request_instance.to_dict()
# create an instance of V1NetworkingWhoisBatchPostRequest from a dict
v1_networking_whois_batch_post_request_from_dict = V1NetworkingWhoisBatchPostRequest.from_dict(v1_networking_whois_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


