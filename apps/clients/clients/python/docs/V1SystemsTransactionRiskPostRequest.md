# V1SystemsTransactionRiskPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**card_bin** | **str** | Card BIN (first 6–8 digits) to look up card issuing country. | 
**ip_address** | **str** | Customer&#39;s IPv4 or IPv6 address. Must be a valid IP. | 
**billing_country** | **str** | ISO 3166-1 alpha-2 billing country code. When provided, country mismatch checks are more precise. | [optional] 
**amount_usd** | **float** | Transaction amount in USD. High-value transactions via VPN incur an additional risk score boost. | [optional] 

## Example

```python
from openapi_client.models.v1_systems_transaction_risk_post_request import V1SystemsTransactionRiskPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsTransactionRiskPostRequest from a JSON string
v1_systems_transaction_risk_post_request_instance = V1SystemsTransactionRiskPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsTransactionRiskPostRequest.to_json())

# convert the object into a dict
v1_systems_transaction_risk_post_request_dict = v1_systems_transaction_risk_post_request_instance.to_dict()
# create an instance of V1SystemsTransactionRiskPostRequest from a dict
v1_systems_transaction_risk_post_request_from_dict = V1SystemsTransactionRiskPostRequest.from_dict(v1_systems_transaction_risk_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


