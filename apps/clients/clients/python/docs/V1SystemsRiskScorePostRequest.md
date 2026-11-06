# V1SystemsRiskScorePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **str** | Email address to include in the risk score. | [optional] 
**phone** | **str** | Phone number in E.164 format. | [optional] 
**ip_address** | **str** | IPv4 or IPv6 address to check. | [optional] 

## Example

```python
from openapi_client.models.v1_systems_risk_score_post_request import V1SystemsRiskScorePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsRiskScorePostRequest from a JSON string
v1_systems_risk_score_post_request_instance = V1SystemsRiskScorePostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsRiskScorePostRequest.to_json())

# convert the object into a dict
v1_systems_risk_score_post_request_dict = v1_systems_risk_score_post_request_instance.to_dict()
# create an instance of V1SystemsRiskScorePostRequest from a dict
v1_systems_risk_score_post_request_from_dict = V1SystemsRiskScorePostRequest.from_dict(v1_systems_risk_score_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


