# V1SystemsPaymentValidatePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bin** | **str** | Card BIN (first 6–8 digits). At least one of bin, iban, or swift is required. | [optional] 
**iban** | **str** | IBAN string to validate. Checksum and format are verified. | [optional] 
**swift** | **str** | SWIFT/BIC code to look up institution and country. | [optional] 

## Example

```python
from openapi_client.models.v1_systems_payment_validate_post_request import V1SystemsPaymentValidatePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsPaymentValidatePostRequest from a JSON string
v1_systems_payment_validate_post_request_instance = V1SystemsPaymentValidatePostRequest.from_json(json)
# print the JSON string representation of the object
print(V1SystemsPaymentValidatePostRequest.to_json())

# convert the object into a dict
v1_systems_payment_validate_post_request_dict = v1_systems_payment_validate_post_request_instance.to_dict()
# create an instance of V1SystemsPaymentValidatePostRequest from a dict
v1_systems_payment_validate_post_request_from_dict = V1SystemsPaymentValidatePostRequest.from_dict(v1_systems_payment_validate_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


