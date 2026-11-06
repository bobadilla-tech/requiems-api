# Org.OpenAPITools.Model.V1FinanceInflationGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | ISO 3166-1 alpha-2 country code, uppercased | [optional] 
**Rate** | **decimal** | Latest annual CPI inflation rate as a percentage (e.g. 2.9495 means 2.9495%) | [optional] 
**Period** | **string** | Year of the latest data point (e.g. 2024) | [optional] 
**Historical** | **List&lt;Object&gt;** | Up to 10 previous years of inflation data, ordered newest to oldest | [optional] 
**HistoricalPeriod** | **string** | Year of the historical data point | [optional] 
**HistoricalRate** | **decimal** | Annual CPI inflation rate for that year | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

