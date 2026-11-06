# V1SystemsPaymentValidatePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bin** | **str** | BIN result. Null when bin was not provided. Contains valid, scheme, card_type, card_level, country_code, issuer, prepaid, luhn. | [optional] 
**iban** | **str** | IBAN result. Null when iban was not provided. Contains valid, country_code, bank_code, account_number. | [optional] 
**swift** | **str** | SWIFT result. Null when swift was not provided. Contains valid, institution, country, branch. | [optional] 
**consistency_ok** | **bool** | True when no cross-field mismatches were detected | [optional] 
**consistency_flags** | **List[object]** | Mismatch flags. Possible values: country_mismatch_bin_iban, country_mismatch_bin_swift, country_mismatch_iban_swift, bank_mismatch_iban_swift | [optional] 

## Example

```python
from openapi_client.models.v1_systems_payment_validate_post200_response_data import V1SystemsPaymentValidatePost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsPaymentValidatePost200ResponseData from a JSON string
v1_systems_payment_validate_post200_response_data_instance = V1SystemsPaymentValidatePost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsPaymentValidatePost200ResponseData.to_json())

# convert the object into a dict
v1_systems_payment_validate_post200_response_data_dict = v1_systems_payment_validate_post200_response_data_instance.to_dict()
# create an instance of V1SystemsPaymentValidatePost200ResponseData from a dict
v1_systems_payment_validate_post200_response_data_from_dict = V1SystemsPaymentValidatePost200ResponseData.from_dict(v1_systems_payment_validate_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


