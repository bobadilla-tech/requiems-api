# Org.OpenAPITools.Model.V1ValidationEmailBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | List of validation results for each email, preserving input order | [optional] 
**ResultsEmail** | **string** | Original email input (null if invalid syntax) | [optional] 
**ResultsValid** | **bool** | Overall validity (syntax + MX record) | [optional] 
**ResultsSyntaxValid** | **bool** | Whether the email is syntactically valid (RFC 5322) | [optional] 
**ResultsMxValid** | **bool** | Whether the domain has valid MX records | [optional] 
**ResultsDisposable** | **bool** | Whether the email comes from a disposable domain | [optional] 
**ResultsNormalized** | **string** | Canonical normalized email (lowercase, alias handling, etc.) | [optional] 
**ResultsDomain** | **string** | Extracted domain from email address | [optional] 
**ResultsSuggestion** | **string** | Suggested correction for common domain typos | [optional] 
**Total** | **int** | Number of emails processed in the batch | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

