# V1SystemsSignupProtectPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**risk_score** | **str** | Composite risk score from 0.0 (clean) to 1.0 (high risk) | [optional] 
**is_safe** | **bool** | True when risk_score is below the safety threshold with no critical flags | [optional] 
**confidence** | **str** | Confidence in the decision from 0.0 to 1.0, based on which signals resolved successfully | [optional] 
**flags** | **List[object]** | Array of triggered risk flags. Possible values: disposable_email, vpn_detected, proxy_detected, tor_detected, is_hosting | [optional] 
**signals_email** | **object** | Email validation result including disposable detection and MX check | [optional] 
**signals_phone** | **object** | Phone number validation result with country code and VoIP/virtual detection | [optional] 
**signals_ip** | **object** | IP intelligence including VPN/proxy/TOR status and fraud score | [optional] 

## Example

```python
from openapi_client.models.v1_systems_signup_protect_post200_response_data import V1SystemsSignupProtectPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsSignupProtectPost200ResponseData from a JSON string
v1_systems_signup_protect_post200_response_data_instance = V1SystemsSignupProtectPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsSignupProtectPost200ResponseData.to_json())

# convert the object into a dict
v1_systems_signup_protect_post200_response_data_dict = v1_systems_signup_protect_post200_response_data_instance.to_dict()
# create an instance of V1SystemsSignupProtectPost200ResponseData from a dict
v1_systems_signup_protect_post200_response_data_from_dict = V1SystemsSignupProtectPost200ResponseData.from_dict(v1_systems_signup_protect_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


