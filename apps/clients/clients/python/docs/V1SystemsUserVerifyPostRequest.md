# V1SystemsUserVerifyPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **str** | Email address to verify. Domain-level WHOIS, MX, and DNS checks are run automatically. | 
**ip_address** | **str** | Optional IPv4 or IPv6 address. When provided, VPN and proxy signals are added to the result. | [optional] 

## Example

```python
from openapi_client.models.v1_systems_user_verify_post_request import V1SystemsUserVerifyPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsUserVerifyPostRequest from a JSON string
v1_systems_user_verify_post_request_instance = V1SystemsUserVerifyPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsUserVerifyPostRequest.to_json())

# convert the object into a dict
v1_systems_user_verify_post_request_dict = v1_systems_user_verify_post_request_instance.to_dict()
# create an instance of V1SystemsUserVerifyPostRequest from a dict
v1_systems_user_verify_post_request_from_dict = V1SystemsUserVerifyPostRequest.from_dict(v1_systems_user_verify_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


