# CommoditiesApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceCommoditiesCommodityGet**](#v1financecommoditiescommodityget) | **GET** /v1/finance/commodities/{commodity} | Get Commodity Price|

# **v1FinanceCommoditiesCommodityGet**
> V1FinanceCommoditiesCommodityGet200Response v1FinanceCommoditiesCommodityGet()

Returns the latest annual average price and up to 10 years of historical data for the requested commodity slug.

### Example

```typescript
import {
    CommoditiesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new CommoditiesApi(configuration);

let commodity: string; //Commodity slug (e.g. gold, silver, oil). See supported slugs below. (default to undefined)

const { status, data } = await apiInstance.v1FinanceCommoditiesCommodityGet(
    commodity
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **commodity** | [**string**] | Commodity slug (e.g. gold, silver, oil). See supported slugs below. | defaults to undefined|


### Return type

**V1FinanceCommoditiesCommodityGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**404** | No data found for the given commodity slug. Check the supported slugs list. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

