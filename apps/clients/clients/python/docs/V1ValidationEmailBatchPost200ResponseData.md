# V1ValidationEmailBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | List of validation results for each email, preserving input order | [optional] 
**results_email** | **str** | Original email input (null if invalid syntax) | [optional] 
**results_valid** | **bool** | Overall validity (syntax + MX record) | [optional] 
**results_syntax_valid** | **bool** | Whether the email is syntactically valid (RFC 5322) | [optional] 
**results_mx_valid** | **bool** | Whether the domain has valid MX records | [optional] 
**results_disposable** | **bool** | Whether the email comes from a disposable domain | [optional] 
**results_normalized** | **str** | Canonical normalized email (lowercase, alias handling, etc.) | [optional] 
**results_domain** | **str** | Extracted domain from email address | [optional] 
**results_suggestion** | **str** | Suggested correction for common domain typos | [optional] 
**total** | **int** | Number of emails processed in the batch | [optional] 

## Example

```python
from openapi_client.models.v1_validation_email_batch_post200_response_data import V1ValidationEmailBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationEmailBatchPost200ResponseData from a JSON string
v1_validation_email_batch_post200_response_data_instance = V1ValidationEmailBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1ValidationEmailBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_validation_email_batch_post200_response_data_dict = v1_validation_email_batch_post200_response_data_instance.to_dict()
# create an instance of V1ValidationEmailBatchPost200ResponseData from a dict
v1_validation_email_batch_post200_response_data_from_dict = V1ValidationEmailBatchPost200ResponseData.from_dict(v1_validation_email_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


