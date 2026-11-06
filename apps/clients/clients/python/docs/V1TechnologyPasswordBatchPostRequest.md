# V1TechnologyPasswordBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **str** | List of password generation requests (1–50 items). | 
**items_length** | **int** | Password length (8–128). Defaults to 16 when omitted. | [optional] 
**items_uppercase** | **bool** | Include uppercase letters (A–Z). Defaults to false. | [optional] 
**items_numbers** | **bool** | Include numbers (0–9). Defaults to false. | [optional] 
**items_symbols** | **bool** | Include special characters. Defaults to false. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_password_batch_post_request import V1TechnologyPasswordBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyPasswordBatchPostRequest from a JSON string
v1_technology_password_batch_post_request_instance = V1TechnologyPasswordBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyPasswordBatchPostRequest.to_json())

# convert the object into a dict
v1_technology_password_batch_post_request_dict = v1_technology_password_batch_post_request_instance.to_dict()
# create an instance of V1TechnologyPasswordBatchPostRequest from a dict
v1_technology_password_batch_post_request_from_dict = V1TechnologyPasswordBatchPostRequest.from_dict(v1_technology_password_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


