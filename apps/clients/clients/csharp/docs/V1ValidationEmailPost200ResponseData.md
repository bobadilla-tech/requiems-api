# Org.OpenAPITools.Model.V1ValidationEmailPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address exactly as supplied in the request body; null when syntax is invalid | [optional] 
**Valid** | **bool** | Overall validity. True only when the address passes syntax validation and the domain has at least one MX record | [optional] 
**SyntaxValid** | **bool** | Whether the address is syntactically valid according to RFC 5322 | [optional] 
**MxValid** | **bool** | Whether the domain has at least one MX record, meaning it can receive email | [optional] 
**Disposable** | **bool** | Whether the address uses a known disposable or temporary email domain | [optional] 
**Normalized** | **string** | The canonical form of the address after normalization (lowercase, plus-tag removal, alias-domain resolution). Null when syntax is invalid | [optional] 
**Domain** | **string** | The domain part of the address (after @). Null when syntax is invalid | [optional] 
**Suggestion** | **string** | A corrected domain name when the supplied domain looks like a typo of a well-known provider (e.g. gmial.com → gmail.com). Null when no close match is found or the domain is already correct | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

