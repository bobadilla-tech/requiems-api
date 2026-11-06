# V1FinanceIbanIbanGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**iban** | **str** | The normalised IBAN (spaces stripped, uppercased) | [optional] 
**valid** | **bool** | true if the IBAN passed length and ISO 13616 checksum validation | [optional] 
**country** | **str** | Full country name (empty if the country code is not in the registry) | [optional] 
**bank_code** | **str** | Bank identifier extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] 
**account** | **str** | Account number extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] 

## Example

```python
from openapi_client.models.v1_finance_iban_iban_get200_response_data import V1FinanceIbanIbanGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceIbanIbanGet200ResponseData from a JSON string
v1_finance_iban_iban_get200_response_data_instance = V1FinanceIbanIbanGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceIbanIbanGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_iban_iban_get200_response_data_dict = v1_finance_iban_iban_get200_response_data_instance.to_dict()
# create an instance of V1FinanceIbanIbanGet200ResponseData from a dict
v1_finance_iban_iban_get200_response_data_from_dict = V1FinanceIbanIbanGet200ResponseData.from_dict(v1_finance_iban_iban_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


