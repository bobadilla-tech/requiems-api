# Org.OpenAPITools.Model.V1FinanceConvertGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Source currency code (uppercased) | [optional] 
**To** | **string** | Target currency code (uppercased) | [optional] 
**Rate** | **decimal** | Exchange rate used for the conversion | [optional] 
**Amount** | **decimal** | The original amount passed in the request | [optional] 
**Converted** | **decimal** | Result of amount × rate, rounded to 2 decimal places | [optional] 
**Timestamp** | **string** | Date the rate was published by the ECB (ISO 8601) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

