# Org.OpenAPITools.Model.V1EntertainmentTriviaBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | List of trivia questions returned in the same order as the input filters. | [optional] 
**ResultsCategory** | **string** | The topic filter used to retrieve the trivia question. | [optional] 
**ResultsDifficulty** | **string** | The complexity level filter used to retrieve the trivia question. | [optional] 
**ResultsError** | **string** | Describes why the question could not be retrieved for this filter. Omitted if the query was successful. | [optional] 
**ResultsData** | **Object** | The trivia question returned for this filter. Empty if the query failed. | [optional] 
**ResultsDataQuestion** | **string** | The trivia question text. | [optional] 
**ResultsDataOptions** | **List&lt;Object&gt;** | List of possible answers for the question. | [optional] 
**ResultsDataAnswer** | **string** | The correct answer to the trivia question. | [optional] 
**ResultsDataCategory** | **string** | The category the trivia question belongs to. | [optional] 
**ResultsDataDifficulty** | **string** | The difficulty level of the trivia question. | [optional] 
**Total** | **int** | Total number of trivia results returned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

