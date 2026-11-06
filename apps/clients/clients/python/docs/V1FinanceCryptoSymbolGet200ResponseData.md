# V1FinanceCryptoSymbolGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**symbol** | **str** | Uppercase ticker symbol | [optional] 
**name** | **str** | Full coin name | [optional] 
**price_usd** | **float** | Current price in USD | [optional] 
**change_24h** | **float** | Price change over the last 24 hours as a percentage | [optional] 
**market_cap** | **float** | Total market capitalisation in USD | [optional] 
**volume_24h** | **float** | Total trading volume over the last 24 hours in USD | [optional] 

## Example

```python
from openapi_client.models.v1_finance_crypto_symbol_get200_response_data import V1FinanceCryptoSymbolGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceCryptoSymbolGet200ResponseData from a JSON string
v1_finance_crypto_symbol_get200_response_data_instance = V1FinanceCryptoSymbolGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceCryptoSymbolGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_crypto_symbol_get200_response_data_dict = v1_finance_crypto_symbol_get200_response_data_instance.to_dict()
# create an instance of V1FinanceCryptoSymbolGet200ResponseData from a dict
v1_finance_crypto_symbol_get200_response_data_from_dict = V1FinanceCryptoSymbolGet200ResponseData.from_dict(v1_finance_crypto_symbol_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


