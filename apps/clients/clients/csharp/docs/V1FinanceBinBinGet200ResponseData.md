# Org.OpenAPITools.Model.V1FinanceBinBinGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | **string** | The normalised BIN prefix used for the lookup | [optional] 
**Scheme** | **string** | Card network: visa, mastercard, amex, discover, jcb, diners, unionpay, maestro, mir, rupay, private_label | [optional] 
**CardType** | **string** | credit, debit, prepaid, or charge | [optional] 
**CardLevel** | **string** | classic, gold, platinum, infinite, business, signature, or standard | [optional] 
**IssuerName** | **string** | Name of the card-issuing bank | [optional] 
**IssuerUrl** | **string** | Bank website URL | [optional] 
**IssuerPhone** | **string** | Bank customer service phone number | [optional] 
**CountryCode** | **string** | ISO 3166-1 alpha-2 country code of the issuing bank (e.g. US, GB, DE) | [optional] 
**CountryName** | **string** | Full country name of the issuing bank | [optional] 
**Prepaid** | **bool** | Whether the card is a prepaid card | [optional] 
**Luhn** | **bool** | Whether the BIN prefix passes the Luhn algorithm check | [optional] 
**Confidence** | **decimal** | Data quality score (0.00–1.00). Multi-source confirmed records score higher. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

