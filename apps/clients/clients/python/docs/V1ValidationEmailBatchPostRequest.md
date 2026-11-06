# V1ValidationEmailBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**emails** | **List[object]** | Array of email addresses to validate. Min: 1, Max: 50. | 

## Example

```python
from openapi_client.models.v1_validation_email_batch_post_request import V1ValidationEmailBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationEmailBatchPostRequest from a JSON string
v1_validation_email_batch_post_request_instance = V1ValidationEmailBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1ValidationEmailBatchPostRequest.to_json())

# convert the object into a dict
v1_validation_email_batch_post_request_dict = v1_validation_email_batch_post_request_instance.to_dict()
# create an instance of V1ValidationEmailBatchPostRequest from a dict
v1_validation_email_batch_post_request_from_dict = V1ValidationEmailBatchPostRequest.from_dict(v1_validation_email_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


