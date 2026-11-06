# V1FinanceInflationBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**countries** | **List[object]** | Array of ISO 3166-1 alpha-2 country codes. Min: 1, Max: 50. | 

## Example

```python
from openapi_client.models.v1_finance_inflation_batch_post_request import V1FinanceInflationBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceInflationBatchPostRequest from a JSON string
v1_finance_inflation_batch_post_request_instance = V1FinanceInflationBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1FinanceInflationBatchPostRequest.to_json())

# convert the object into a dict
v1_finance_inflation_batch_post_request_dict = v1_finance_inflation_batch_post_request_instance.to_dict()
# create an instance of V1FinanceInflationBatchPostRequest from a dict
v1_finance_inflation_batch_post_request_from_dict = V1FinanceInflationBatchPostRequest.from_dict(v1_finance_inflation_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


