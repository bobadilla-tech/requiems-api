# Org.OpenAPITools.Model.V1ValidationProfanityBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Per-item results in input order. | [optional] 
**ResultsText** | **string** | The original text that was checked. | [optional] 
**ResultsResultHasProfanity** | **bool** | Whether any profanity was detected. | [optional] 
**ResultsResultCensored** | **string** | Text with profane words replaced by asterisks. | [optional] 
**ResultsResultFlaggedWords** | **string** | Deduplicated list of detected profane words (lowercase). | [optional] 
**Total** | **int** | Total number of items processed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

