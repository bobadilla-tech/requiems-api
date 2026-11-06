# Org.OpenAPITools.Model.V1PlacesTimezoneBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Per-item results in input order. | [optional] 
**ResultsInfoTimezone** | **string** | IANA timezone identifier. | [optional] 
**ResultsInfoOffset** | **string** | UTC offset in +HH:MM or -HH:MM format. | [optional] 
**ResultsInfoCurrentTime** | **string** | Current time in UTC (RFC 3339). | [optional] 
**ResultsInfoIsDst** | **bool** | Whether daylight saving time is currently active. | [optional] 
**ResultsError** | **string** | Error message if this item failed (omitted on success). | [optional] 
**Total** | **int** | Total number of items processed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

