# Org.OpenAPITools.Model.V1FinanceCommoditiesCommodityGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commodity** | **string** | The commodity slug as provided in the request path | [optional] 
**Name** | **string** | Human-readable commodity name | [optional] 
**Price** | **decimal** | Latest annual average price in the commodity&#39;s display unit | [optional] 
**Unit** | **string** | Price unit (oz, barrel, mmbtu, lb, or metric_ton) | [optional] 
**Currency** | **string** | Currency code — always USD | [optional] 
**Change24h** | **decimal** | Year-over-year percentage change from the prior year&#39;s annual average (positive &#x3D; price increased) | [optional] 
**Historical** | **List&lt;Object&gt;** | Up to 10 prior years of annual average prices, ordered newest to oldest | [optional] 
**HistoricalPeriod** | **string** | Year of the historical data point | [optional] 
**HistoricalPrice** | **decimal** | Annual average price for that year | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

