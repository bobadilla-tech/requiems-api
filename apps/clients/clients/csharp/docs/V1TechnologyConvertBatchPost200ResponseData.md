# Org.OpenAPITools.Model.V1TechnologyConvertBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | List of unit conversion results returned in the same order as the input operations. | [optional] 
**ResultsFrom** | **string** | Source unit used for the conversion. | [optional] 
**ResultsTo** | **string** | Target unit used for the conversion. | [optional] 
**ResultsSuccess** | **bool** | Indicates whether the conversion operation was successful. | [optional] 
**ResultsError** | **string** | Error message if the conversion failed. Empty when success is true. | [optional] 
**ResultsData** | **Object** | Conversion result details. Contains empty values when the conversion fails. | [optional] 
**ResultsDataFrom** | **string** | Source unit. | [optional] 
**ResultsDataTo** | **string** | Target unit. | [optional] 
**ResultsDataInput** | **decimal** | Original numeric value provided for conversion. | [optional] 
**ResultsDataResult** | **decimal** | Converted numeric result. | [optional] 
**ResultsDataFormula** | **string** | Formula or conversion factor used to calculate the result. | [optional] 
**Total** | **int** | Total number of conversion results returned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

