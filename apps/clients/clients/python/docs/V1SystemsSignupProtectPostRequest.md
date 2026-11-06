# V1SystemsSignupProtectPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **str** | Email address to validate and score. At least one of email, phone, or ip_address is required. | [optional] 
**phone** | **str** | Phone number in E.164 format to validate and check for VoIP or virtual numbers. | [optional] 
**ip_address** | **str** | IPv4 or IPv6 address to check for VPN, proxy, TOR, and hosting status. | [optional] 

## Example

```python
from openapi_client.models.v1_systems_signup_protect_post_request import V1SystemsSignupProtectPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsSignupProtectPostRequest from a JSON string
v1_systems_signup_protect_post_request_instance = V1SystemsSignupProtectPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsSignupProtectPostRequest.to_json())

# convert the object into a dict
v1_systems_signup_protect_post_request_dict = v1_systems_signup_protect_post_request_instance.to_dict()
# create an instance of V1SystemsSignupProtectPostRequest from a dict
v1_systems_signup_protect_post_request_from_dict = V1SystemsSignupProtectPostRequest.from_dict(v1_systems_signup_protect_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


