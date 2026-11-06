# RandomUserApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyRandomUserBatchPost**](#v1technologyrandomuserbatchpost) | **POST** /v1/technology/random-user/batch | Batch Generate Users|
|[**v1TechnologyRandomUserGet**](#v1technologyrandomuserget) | **GET** /v1/technology/random-user | Get Random User|

# **v1TechnologyRandomUserBatchPost**
> V1TechnologyRandomUserBatchPost200Response v1TechnologyRandomUserBatchPost(v1TechnologyRandomUserBatchPostRequest)

Returns multiple randomly generated fake user profiles in a single request. Each call consumes count units of quota. Maximum 50 users per request.

### Example

```typescript
import {
    RandomUserApi,
    Configuration,
    V1TechnologyRandomUserBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new RandomUserApi(configuration);

let v1TechnologyRandomUserBatchPostRequest: V1TechnologyRandomUserBatchPostRequest; //

const { status, data } = await apiInstance.v1TechnologyRandomUserBatchPost(
    v1TechnologyRandomUserBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyRandomUserBatchPostRequest** | **V1TechnologyRandomUserBatchPostRequest**|  | |


### Return type

**V1TechnologyRandomUserBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Request body is malformed or not valid JSON |  -  |
|**422** | count is missing, less than 1, or greater than 50 |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyRandomUserGet**
> V1TechnologyRandomUserGet200Response v1TechnologyRandomUserGet()

Returns a randomly generated fake user profile.

### Example

```typescript
import {
    RandomUserApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new RandomUserApi(configuration);

const { status, data } = await apiInstance.v1TechnologyRandomUserGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1TechnologyRandomUserGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

