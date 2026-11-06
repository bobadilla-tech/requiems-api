# QuotesApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextQuotesRandomGet**](#v1textquotesrandomget) | **GET** /v1/text/quotes/random | Get Random Quote|

# **v1TextQuotesRandomGet**
> V1TextQuotesRandomGet200Response v1TextQuotesRandomGet()

Returns a random inspirational quote with author attribution

### Example

```typescript
import {
    QuotesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new QuotesApi(configuration);

const { status, data } = await apiInstance.v1TextQuotesRandomGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1TextQuotesRandomGet200Response**

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

