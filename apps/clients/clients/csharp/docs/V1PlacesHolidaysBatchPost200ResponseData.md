# Org.OpenAPITools.Model.V1PlacesHolidaysBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | One result per query, in the same order as the request | [optional] 
**ResultsCountry** | **string** | ISO 3166-1 alpha-2 country code | [optional] 
**ResultsYear** | **int** | Year queried | [optional] 
**ResultsFound** | **bool** | false when no holidays exist for that country/year combination | [optional] 
**ResultsHolidays** | **List&lt;Object&gt;** | List of holidays. Omitted when found is false. | [optional] 
**ResultsTotal** | **int** | Number of holidays. Omitted when found is false. | [optional] 
**Total** | **int** | Total number of results (equals the number of queries sent) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

