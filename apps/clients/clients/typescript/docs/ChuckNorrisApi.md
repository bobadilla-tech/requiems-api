# ChuckNorrisApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentChuckNorrisGet**](#v1entertainmentchucknorrisget) | **GET** /v1/entertainment/chuck-norris | Get Random Chuck Norris Fact|

# **v1EntertainmentChuckNorrisGet**
> V1EntertainmentChuckNorrisGet200Response v1EntertainmentChuckNorrisGet()

Returns a randomly selected Chuck Norris fact from the built-in database.

### Example

```typescript
import {
    ChuckNorrisApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ChuckNorrisApi(configuration);

const { status, data } = await apiInstance.v1EntertainmentChuckNorrisGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1EntertainmentChuckNorrisGet200Response**

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

