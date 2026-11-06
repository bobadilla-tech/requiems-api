# V1SystemsTransactionRiskPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**risk_score** | **str** | Transaction risk score from 0.0 to 1.0 | [optional] 
**is_safe** | **bool** | True when score is below 0.5 and no TOR or proxy detected | [optional] 
**flags** | **List[object]** | Triggered flags. Possible values: ip_country_mismatch, bin_country_mismatch, country_mismatch, vpn_detected, proxy_detected, tor_detected, high_value_vpn | [optional] 
**signals_ip_country** | **str** | Country code resolved from the IP address | [optional] 
**signals_billing_country** | **str** | Billing country provided in the request, or null if omitted | [optional] 
**signals_bin_country** | **str** | Country of the card issuer resolved from the BIN | [optional] 
**signals_country_mismatch** | **bool** | True when any country comparison produced a mismatch | [optional] 

## Example

```python
from openapi_client.models.v1_systems_transaction_risk_post200_response_data import V1SystemsTransactionRiskPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsTransactionRiskPost200ResponseData from a JSON string
v1_systems_transaction_risk_post200_response_data_instance = V1SystemsTransactionRiskPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsTransactionRiskPost200ResponseData.to_json())

# convert the object into a dict
v1_systems_transaction_risk_post200_response_data_dict = v1_systems_transaction_risk_post200_response_data_instance.to_dict()
# create an instance of V1SystemsTransactionRiskPost200ResponseData from a dict
v1_systems_transaction_risk_post200_response_data_from_dict = V1SystemsTransactionRiskPost200ResponseData.from_dict(v1_systems_transaction_risk_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


