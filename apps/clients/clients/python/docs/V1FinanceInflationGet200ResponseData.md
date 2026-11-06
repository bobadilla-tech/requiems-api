# V1FinanceInflationGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**country** | **str** | ISO 3166-1 alpha-2 country code, uppercased | [optional] 
**rate** | **float** | Latest annual CPI inflation rate as a percentage (e.g. 2.9495 means 2.9495%) | [optional] 
**period** | **str** | Year of the latest data point (e.g. 2024) | [optional] 
**historical** | **List[object]** | Up to 10 previous years of inflation data, ordered newest to oldest | [optional] 
**historical_period** | **str** | Year of the historical data point | [optional] 
**historical_rate** | **float** | Annual CPI inflation rate for that year | [optional] 

## Example

```python
from openapi_client.models.v1_finance_inflation_get200_response_data import V1FinanceInflationGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceInflationGet200ResponseData from a JSON string
v1_finance_inflation_get200_response_data_instance = V1FinanceInflationGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceInflationGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_inflation_get200_response_data_dict = v1_finance_inflation_get200_response_data_instance.to_dict()
# create an instance of V1FinanceInflationGet200ResponseData from a dict
v1_finance_inflation_get200_response_data_from_dict = V1FinanceInflationGet200ResponseData.from_dict(v1_finance_inflation_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


