# Org.OpenAPITools.Model.V1PlacesPostalBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Per-item results in input order. | [optional] 
**ResultsCode** | **string** | The postal code as provided in the request. | [optional] 
**ResultsCountry** | **string** | The country code used for the lookup (defaulted to US when omitted). | [optional] 
**ResultsFound** | **bool** | Whether the postal code was found in the dataset. | [optional] 
**ResultsResult** | **Object** | Postal code details (omitted when found is false). | [optional] 
**Total** | **int** | Total number of items processed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

