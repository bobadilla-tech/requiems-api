# Org.OpenAPITools.Model.V1TechnologyPasswordBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Per-item results in input order. | [optional] 
**ResultsResultPassword** | **string** | The generated password (omitted on error). | [optional] 
**ResultsResultLength** | **int** | Length of the generated password. | [optional] 
**ResultsResultStrength** | **string** | Strength assessment (weak, medium, or strong). | [optional] 
**ResultsError** | **string** | Error message if generation failed (omitted on success). | [optional] 
**Total** | **int** | Total number of items processed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

