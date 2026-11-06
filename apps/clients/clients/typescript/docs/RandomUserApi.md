# RandomUserApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyRandomUserGet**](#v1technologyrandomuserget) | **GET** /v1/technology/random-user | Get Random User|

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

