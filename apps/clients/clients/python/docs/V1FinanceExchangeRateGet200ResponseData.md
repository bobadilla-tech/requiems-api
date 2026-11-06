# V1FinanceExchangeRateGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**var_from** | **str** | Source currency code (uppercased) | [optional] 
**to** | **str** | Target currency code (uppercased) | [optional] 
**rate** | **float** | Exchange rate — how many units of &#x60;to&#x60; equal 1 unit of &#x60;from&#x60; | [optional] 
**timestamp** | **str** | Date the rate was published by the ECB (ISO 8601) | [optional] 

## Example

```python
from openapi_client.models.v1_finance_exchange_rate_get200_response_data import V1FinanceExchangeRateGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceExchangeRateGet200ResponseData from a JSON string
v1_finance_exchange_rate_get200_response_data_instance = V1FinanceExchangeRateGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceExchangeRateGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_exchange_rate_get200_response_data_dict = v1_finance_exchange_rate_get200_response_data_instance.to_dict()
# create an instance of V1FinanceExchangeRateGet200ResponseData from a dict
v1_finance_exchange_rate_get200_response_data_from_dict = V1FinanceExchangeRateGet200ResponseData.from_dict(v1_finance_exchange_rate_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


