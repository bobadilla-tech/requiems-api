# V1FinanceInflationBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | One entry per country, in the same order as the input array | [optional] 
**results_country** | **str** | ISO 3166-1 alpha-2 country code, uppercased | [optional] 
**results_found** | **bool** | false when the country has no data in the World Bank set | [optional] 
**results_rate** | **float** | Latest CPI inflation rate. Omitted when found: false | [optional] 
**results_period** | **str** | Year of the latest data point. Omitted when found: false | [optional] 
**results_historical** | **List[object]** | Up to 10 previous years. Omitted when found: false | [optional] 
**total** | **int** | Total number of results returned (equals number of countries sent) | [optional] 

## Example

```python
from openapi_client.models.v1_finance_inflation_batch_post200_response_data import V1FinanceInflationBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceInflationBatchPost200ResponseData from a JSON string
v1_finance_inflation_batch_post200_response_data_instance = V1FinanceInflationBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceInflationBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_finance_inflation_batch_post200_response_data_dict = v1_finance_inflation_batch_post200_response_data_instance.to_dict()
# create an instance of V1FinanceInflationBatchPost200ResponseData from a dict
v1_finance_inflation_batch_post200_response_data_from_dict = V1FinanceInflationBatchPost200ResponseData.from_dict(v1_finance_inflation_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


