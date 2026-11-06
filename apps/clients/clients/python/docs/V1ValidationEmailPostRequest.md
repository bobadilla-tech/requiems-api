# V1ValidationEmailPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **str** | The email address to validate. | 

## Example

```python
from openapi_client.models.v1_validation_email_post_request import V1ValidationEmailPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationEmailPostRequest from a JSON string
v1_validation_email_post_request_instance = V1ValidationEmailPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1ValidationEmailPostRequest.to_json())

# convert the object into a dict
v1_validation_email_post_request_dict = v1_validation_email_post_request_instance.to_dict()
# create an instance of V1ValidationEmailPostRequest from a dict
v1_validation_email_post_request_from_dict = V1ValidationEmailPostRequest.from_dict(v1_validation_email_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


