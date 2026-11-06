# V1FinanceMortgageBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Mortgage calculation result for each mortgage request in the same order as the input. Each item has the same fields as the single mortgage endpoint. | [optional] 
**total** | **int** | Number of results returned. Matches the length of the input array. | [optional] 

## Example

```python
from openapi_client.models.v1_finance_mortgage_batch_post200_response_data import V1FinanceMortgageBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceMortgageBatchPost200ResponseData from a JSON string
v1_finance_mortgage_batch_post200_response_data_instance = V1FinanceMortgageBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceMortgageBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_finance_mortgage_batch_post200_response_data_dict = v1_finance_mortgage_batch_post200_response_data_instance.to_dict()
# create an instance of V1FinanceMortgageBatchPost200ResponseData from a dict
v1_finance_mortgage_batch_post200_response_data_from_dict = V1FinanceMortgageBatchPost200ResponseData.from_dict(v1_finance_mortgage_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


