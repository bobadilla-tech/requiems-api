# Org.OpenAPITools.Model.V1PlacesReverseGeocodeBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Per-item results in input order. | [optional] 
**ResultsLat** | **decimal** | Latitude as provided in the request. | [optional] 
**ResultsLon** | **decimal** | Longitude as provided in the request. | [optional] 
**ResultsResult** | **Object** | Reverse geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] 
**ResultsError** | **string** | Error message if reverse geocoding failed (omitted on success). | [optional] 
**Total** | **int** | Total number of items processed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

