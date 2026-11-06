# AdviceApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextAdviceGet**](#v1textadviceget) | **GET** /v1/text/advice | Get Random Advice|

# **v1TextAdviceGet**
> V1TextAdviceGet200Response v1TextAdviceGet()

Returns a random piece of advice

### Example

```typescript
import {
    AdviceApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new AdviceApi(configuration);

const { status, data } = await apiInstance.v1TextAdviceGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1TextAdviceGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

