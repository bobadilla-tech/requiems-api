# Org.OpenAPITools.Model.V1TextNormalizePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Original** | **string** | The email address exactly as supplied in the request body | [optional] 
**Normalized** | **string** | The canonical form of the address after all transformations | [optional] 
**Local** | **string** | The local part (before @) of the normalized address | [optional] 
**Domain** | **string** | The domain part (after @) of the normalized address | [optional] 
**Changes** | **List&lt;Object&gt;** | Ordered list of transformations applied. Possible values: lowercased, trimmed_whitespace, removed_dots, removed_plus_tag, canonicalised_domain. Empty array when no changes were needed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

