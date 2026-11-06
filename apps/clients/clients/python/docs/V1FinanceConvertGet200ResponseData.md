# V1FinanceConvertGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**var_from** | **str** | Source currency code (uppercased) | [optional] 
**to** | **str** | Target currency code (uppercased) | [optional] 
**rate** | **float** | Exchange rate used for the conversion | [optional] 
**amount** | **float** | The original amount passed in the request | [optional] 
**converted** | **float** | Result of amount × rate, rounded to 2 decimal places | [optional] 
**timestamp** | **str** | Date the rate was published by the ECB (ISO 8601) | [optional] 

## Example

```python
from openapi_client.models.v1_finance_convert_get200_response_data import V1FinanceConvertGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceConvertGet200ResponseData from a JSON string
v1_finance_convert_get200_response_data_instance = V1FinanceConvertGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceConvertGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_convert_get200_response_data_dict = v1_finance_convert_get200_response_data_instance.to_dict()
# create an instance of V1FinanceConvertGet200ResponseData from a dict
v1_finance_convert_get200_response_data_from_dict = V1FinanceConvertGet200ResponseData.from_dict(v1_finance_convert_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


