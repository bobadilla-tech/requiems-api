# Org.OpenAPITools.Model.V1NetworkingWhoisBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | One entry per domain, in the same order as the input | [optional] 
**ResultsDomain** | **string** | Domain name requested in the batch | [optional] 
**ResultsFound** | **bool** | False when WHOIS data could not be found | [optional] 
**ResultsError** | **string** | Error message when found is false | [optional] 
**ResultsData** | **Object** | WHOIS information for the domain when found is true | [optional] 
**Total** | **int** | Total number of results returned | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

