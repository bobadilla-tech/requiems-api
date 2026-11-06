# QuotesApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentQuotesRandomBatchPost**](#v1entertainmentquotesrandombatchpost) | **POST** /v1/entertainment/quotes/random/batch | Get Random Quotes (Batch)|
|[**v1EntertainmentQuotesRandomGet**](#v1entertainmentquotesrandomget) | **GET** /v1/entertainment/quotes/random | Get Random Quote|

# **v1EntertainmentQuotesRandomBatchPost**
> V1EntertainmentQuotesRandomBatchPost200Response v1EntertainmentQuotesRandomBatchPost(v1EntertainmentQuotesRandomBatchPostRequest)

Returns up to 50 random quotes in a single request. Each quote counts as one unit of usage.

### Example

```typescript
import {
    QuotesApi,
    Configuration,
    V1EntertainmentQuotesRandomBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new QuotesApi(configuration);

let v1EntertainmentQuotesRandomBatchPostRequest: V1EntertainmentQuotesRandomBatchPostRequest; //

const { status, data } = await apiInstance.v1EntertainmentQuotesRandomBatchPost(
    v1EntertainmentQuotesRandomBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1EntertainmentQuotesRandomBatchPostRequest** | **V1EntertainmentQuotesRandomBatchPostRequest**|  | |


### Return type

**V1EntertainmentQuotesRandomBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Request body is missing or not valid JSON. |  -  |
|**422** | count is missing, less than 1, or greater than 50. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1EntertainmentQuotesRandomGet**
> V1EntertainmentQuotesRandomGet200Response v1EntertainmentQuotesRandomGet()

Returns a random inspirational quote with author attribution

### Example

```typescript
import {
    QuotesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new QuotesApi(configuration);

const { status, data } = await apiInstance.v1EntertainmentQuotesRandomGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1EntertainmentQuotesRandomGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**503** | No quotes available in the database |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

