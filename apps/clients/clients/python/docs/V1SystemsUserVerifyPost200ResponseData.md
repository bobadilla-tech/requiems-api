# V1SystemsUserVerifyPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**verified** | **bool** | True when risk score is below 0.3, confidence is above 0.5, email is valid, domain has MX, and domain is registered | [optional] 
**confidence** | **str** | Confidence in the verification result, based on how many signals resolved | [optional] 
**risk_score** | **str** | Composite risk score from 0.0 to 1.0 | [optional] 
**flags** | **List[object]** | Triggered risk flags. Possible values: email_invalid, disposable_email, no_mx, domain_not_registered, young_domain, whois_unavailable, ip_risk | [optional] 
**signals_email** | **object** | Email validation result | [optional] 
**signals_domain** | **object** | Domain intelligence including WHOIS age, MX and A record presence, and registration status | [optional] 
**signals_ip** | **str** | IP risk signals. Null when ip_address was not provided | [optional] 

## Example

```python
from openapi_client.models.v1_systems_user_verify_post200_response_data import V1SystemsUserVerifyPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsUserVerifyPost200ResponseData from a JSON string
v1_systems_user_verify_post200_response_data_instance = V1SystemsUserVerifyPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsUserVerifyPost200ResponseData.to_json())

# convert the object into a dict
v1_systems_user_verify_post200_response_data_dict = v1_systems_user_verify_post200_response_data_instance.to_dict()
# create an instance of V1SystemsUserVerifyPost200ResponseData from a dict
v1_systems_user_verify_post200_response_data_from_dict = V1SystemsUserVerifyPost200ResponseData.from_dict(v1_systems_user_verify_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


