# Org.OpenAPITools.Model.V1PlacesTimezoneBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **string** | List of timezone lookup requests (1–50 items). | 
**ItemsTimezone** | **string** | IANA timezone name (e.g. &#39;America/New_York&#39;). Highest priority. | [optional] 
**ItemsCity** | **string** | City name for city-based lookup. Used when timezone is not set. | [optional] 
**ItemsLat** | **decimal** | Latitude (-90 to 90). Used when neither timezone nor city is set. | [optional] 
**ItemsLon** | **decimal** | Longitude (-180 to 180). Used alongside lat for coordinate-based lookup. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

