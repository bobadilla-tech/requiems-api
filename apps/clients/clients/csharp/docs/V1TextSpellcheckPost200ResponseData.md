# Org.OpenAPITools.Model.V1TextSpellcheckPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Corrected** | **string** | The full input text with all misspelled words replaced by the top-ranked suggestion | [optional] 
**Corrections** | **string** | List of individual corrections. Each item contains: original (the misspelled word), suggested (the top-ranked replacement applied to corrected), suggestions (up to 3 ranked alternatives returned by LanguageTool, from most to least likely), and position (0-based Unicode character offset in the original text) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

