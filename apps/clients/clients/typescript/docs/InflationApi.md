# InflationApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceInflationBatchPost**](#v1financeinflationbatchpost) | **POST** /v1/finance/inflation/batch | Batch Inflation Rates|
|[**v1FinanceInflationGet**](#v1financeinflationget) | **GET** /v1/finance/inflation | Get Inflation Rate|

# **v1FinanceInflationBatchPost**
> V1FinanceInflationBatchPost200Response v1FinanceInflationBatchPost(v1FinanceInflationBatchPostRequest)

Returns inflation data for up to 50 countries in a single request. Results are in the same order as the input. Countries with no data return found: false instead of failing the whole request. Billing: 1 credit per country (not per HTTP request).

### Example

```typescript
import {
    InflationApi,
    Configuration,
    V1FinanceInflationBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new InflationApi(configuration);

let v1FinanceInflationBatchPostRequest: V1FinanceInflationBatchPostRequest; //

const { status, data } = await apiInstance.v1FinanceInflationBatchPost(
    v1FinanceInflationBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1FinanceInflationBatchPostRequest** | **V1FinanceInflationBatchPostRequest**|  | |


### Return type

**V1FinanceInflationBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**422** | Body is invalid: empty array, more than 50 items, or a bad country code. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1FinanceInflationGet**
> V1FinanceInflationGet200Response v1FinanceInflationGet()

Returns the latest annual CPI inflation rate for a country plus the previous 10 years of historical data.

### Example

```typescript
import {
    InflationApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new InflationApi(configuration);

let country: string; //ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive. (default to undefined)

const { status, data } = await apiInstance.v1FinanceInflationGet(
    country
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **country** | [**string**] | ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive. | defaults to undefined|


### Return type

**V1FinanceInflationGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The country parameter is missing or is not a valid ISO 3166-1 alpha-2 code. |  -  |
|**404** | No inflation data found for the given country code. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

