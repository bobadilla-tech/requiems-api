# Org.OpenAPITools.Model.V1FinanceIbanIbanGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | The normalised IBAN (spaces stripped, uppercased) | [optional] 
**Valid** | **bool** | true if the IBAN passed length and ISO 13616 checksum validation | [optional] 
**Country** | **string** | Full country name (empty if the country code is not in the registry) | [optional] 
**BankCode** | **string** | Bank identifier extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] 
**Account** | **string** | Account number extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

