# Org.OpenAPITools.Model.V1FinanceInflationBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | One entry per country, in the same order as the input array | [optional] 
**ResultsCountry** | **string** | ISO 3166-1 alpha-2 country code, uppercased | [optional] 
**ResultsFound** | **bool** | false when the country has no data in the World Bank set | [optional] 
**ResultsRate** | **decimal** | Latest CPI inflation rate. Omitted when found: false | [optional] 
**ResultsPeriod** | **string** | Year of the latest data point. Omitted when found: false | [optional] 
**ResultsHistorical** | **List&lt;Object&gt;** | Up to 10 previous years. Omitted when found: false | [optional] 
**Total** | **int** | Total number of results returned (equals number of countries sent) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

