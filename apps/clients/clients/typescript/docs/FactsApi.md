# FactsApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentFactsGet**](#v1entertainmentfactsget) | **GET** /v1/entertainment/facts | Get Random Fact|

# **v1EntertainmentFactsGet**
> V1EntertainmentFactsGet200Response v1EntertainmentFactsGet()

Returns a randomly selected fact, optionally filtered by category.

### Example

```typescript
import {
    FactsApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new FactsApi(configuration);

let category: string; //Filter by category. Valid values: science, history, technology, nature, space, food (optional) (default to undefined)

const { status, data } = await apiInstance.v1EntertainmentFactsGet(
    category
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **category** | [**string**] | Filter by category. Valid values: science, history, technology, nature, space, food | (optional) defaults to undefined|


### Return type

**V1EntertainmentFactsGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | bad_request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

