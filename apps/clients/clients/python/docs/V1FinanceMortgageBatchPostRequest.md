# V1FinanceMortgageBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mortgages** | **List[object]** | Array of mortgages to calculate (min: 1, max: 50). | 

## Example

```python
from openapi_client.models.v1_finance_mortgage_batch_post_request import V1FinanceMortgageBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceMortgageBatchPostRequest from a JSON string
v1_finance_mortgage_batch_post_request_instance = V1FinanceMortgageBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1FinanceMortgageBatchPostRequest.to_json())

# convert the object into a dict
v1_finance_mortgage_batch_post_request_dict = v1_finance_mortgage_batch_post_request_instance.to_dict()
# create an instance of V1FinanceMortgageBatchPostRequest from a dict
v1_finance_mortgage_batch_post_request_from_dict = V1FinanceMortgageBatchPostRequest.from_dict(v1_finance_mortgage_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


