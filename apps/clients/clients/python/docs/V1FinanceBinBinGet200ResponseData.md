# V1FinanceBinBinGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bin** | **str** | The normalised BIN prefix used for the lookup | [optional] 
**scheme** | **str** | Card network: visa, mastercard, amex, discover, jcb, diners, unionpay, maestro, mir, rupay, private_label | [optional] 
**card_type** | **str** | credit, debit, prepaid, or charge | [optional] 
**card_level** | **str** | classic, gold, platinum, infinite, business, signature, or standard | [optional] 
**issuer_name** | **str** | Name of the card-issuing bank | [optional] 
**issuer_url** | **str** | Bank website URL | [optional] 
**issuer_phone** | **str** | Bank customer service phone number | [optional] 
**country_code** | **str** | ISO 3166-1 alpha-2 country code of the issuing bank (e.g. US, GB, DE) | [optional] 
**country_name** | **str** | Full country name of the issuing bank | [optional] 
**prepaid** | **bool** | Whether the card is a prepaid card | [optional] 
**luhn** | **bool** | Whether the BIN prefix passes the Luhn algorithm check | [optional] 
**confidence** | **float** | Data quality score (0.00–1.00). Multi-source confirmed records score higher. | [optional] 

## Example

```python
from openapi_client.models.v1_finance_bin_bin_get200_response_data import V1FinanceBinBinGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceBinBinGet200ResponseData from a JSON string
v1_finance_bin_bin_get200_response_data_instance = V1FinanceBinBinGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceBinBinGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_bin_bin_get200_response_data_dict = v1_finance_bin_bin_get200_response_data_instance.to_dict()
# create an instance of V1FinanceBinBinGet200ResponseData from a dict
v1_finance_bin_bin_get200_response_data_from_dict = V1FinanceBinBinGet200ResponseData.from_dict(v1_finance_bin_bin_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


