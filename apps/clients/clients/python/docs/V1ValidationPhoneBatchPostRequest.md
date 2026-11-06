# V1ValidationPhoneBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**numbers** | **List[object]** | Array of phone numbers to validate (min: 1, max: 50). Each must include the country calling code. | 

## Example

```python
from openapi_client.models.v1_validation_phone_batch_post_request import V1ValidationPhoneBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1ValidationPhoneBatchPostRequest from a JSON string
v1_validation_phone_batch_post_request_instance = V1ValidationPhoneBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1ValidationPhoneBatchPostRequest.to_json())

# convert the object into a dict
v1_validation_phone_batch_post_request_dict = v1_validation_phone_batch_post_request_instance.to_dict()
# create an instance of V1ValidationPhoneBatchPostRequest from a dict
v1_validation_phone_batch_post_request_from_dict = V1ValidationPhoneBatchPostRequest.from_dict(v1_validation_phone_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


