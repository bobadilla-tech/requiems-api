# Org.OpenAPITools.Model.V1TechnologyBaseBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Per-item results in input order. | [optional] 
**ResultsFrom** | **int** | The source base used for conversion. | [optional] 
**ResultsTo** | **int** | The target base used for conversion. | [optional] 
**ResultsInput** | **string** | The original value as provided in the request. | [optional] 
**ResultsResult** | **string** | The converted value in the target base (omitted on error). | [optional] 
**ResultsError** | **string** | Error message if this item failed (omitted on success). | [optional] 
**Total** | **int** | Total number of items processed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

