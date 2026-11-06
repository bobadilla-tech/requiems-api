# V1NetworkingIpAsnBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ips** | **str** | List of IP addresses to look up (1–50 items, each must be a valid IPv4 or IPv6 address). | 

## Example

```python
from openapi_client.models.v1_networking_ip_asn_batch_post_request import V1NetworkingIpAsnBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1NetworkingIpAsnBatchPostRequest from a JSON string
v1_networking_ip_asn_batch_post_request_instance = V1NetworkingIpAsnBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1NetworkingIpAsnBatchPostRequest.to_json())

# convert the object into a dict
v1_networking_ip_asn_batch_post_request_dict = v1_networking_ip_asn_batch_post_request_instance.to_dict()
# create an instance of V1NetworkingIpAsnBatchPostRequest from a dict
v1_networking_ip_asn_batch_post_request_from_dict = V1NetworkingIpAsnBatchPostRequest.from_dict(v1_networking_ip_asn_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


