# V1FinanceIbanBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**numbers** | **List[object]** | Array of iban numbers to validate (min: 1, max: 50). | 

## Example

```python
from openapi_client.models.v1_finance_iban_batch_post_request import V1FinanceIbanBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceIbanBatchPostRequest from a JSON string
v1_finance_iban_batch_post_request_instance = V1FinanceIbanBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1FinanceIbanBatchPostRequest.to_json())

# convert the object into a dict
v1_finance_iban_batch_post_request_dict = v1_finance_iban_batch_post_request_instance.to_dict()
# create an instance of V1FinanceIbanBatchPostRequest from a dict
v1_finance_iban_batch_post_request_from_dict = V1FinanceIbanBatchPostRequest.from_dict(v1_finance_iban_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


