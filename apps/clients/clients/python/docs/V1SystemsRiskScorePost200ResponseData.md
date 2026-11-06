# V1SystemsRiskScorePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**risk_score** | **str** | Composite risk score from 0.0 to 1.0 | [optional] 
**is_safe** | **bool** | True when the risk score is below the safety threshold | [optional] 
**confidence** | **str** | Confidence in the score based on resolved signals | [optional] 
**flags** | **List[object]** | Array of triggered risk flags | [optional] 

## Example

```python
from openapi_client.models.v1_systems_risk_score_post200_response_data import V1SystemsRiskScorePost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsRiskScorePost200ResponseData from a JSON string
v1_systems_risk_score_post200_response_data_instance = V1SystemsRiskScorePost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsRiskScorePost200ResponseData.to_json())

# convert the object into a dict
v1_systems_risk_score_post200_response_data_dict = v1_systems_risk_score_post200_response_data_instance.to_dict()
# create an instance of V1SystemsRiskScorePost200ResponseData from a dict
v1_systems_risk_score_post200_response_data_from_dict = V1SystemsRiskScorePost200ResponseData.from_dict(v1_systems_risk_score_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


