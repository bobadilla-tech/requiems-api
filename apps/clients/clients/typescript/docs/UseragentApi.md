# UseragentApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyUseragentBatchPost**](#v1technologyuseragentbatchpost) | **POST** /v1/technology/useragent/batch | Batch Parse User Agents|
|[**v1TechnologyUseragentGet**](#v1technologyuseragentget) | **GET** /v1/technology/useragent | Parse User Agent|

# **v1TechnologyUseragentBatchPost**
> V1TechnologyUseragentBatchPost200Response v1TechnologyUseragentBatchPost(v1TechnologyUseragentBatchPostRequest)

Parses up to 50 user agents in a single request.

### Example

```typescript
import {
    UseragentApi,
    Configuration,
    V1TechnologyUseragentBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new UseragentApi(configuration);

let v1TechnologyUseragentBatchPostRequest: V1TechnologyUseragentBatchPostRequest; //

const { status, data } = await apiInstance.v1TechnologyUseragentBatchPost(
    v1TechnologyUseragentBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyUseragentBatchPostRequest** | **V1TechnologyUseragentBatchPostRequest**|  | |


### Return type

**V1TechnologyUseragentBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid JSON, malformed request body, or unexpected field types. |  -  |
|**422** | The user_agents array is missing, empty, or contains more than 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyUseragentGet**
> V1TechnologyUseragentGet200Response v1TechnologyUseragentGet()

Parses a user agent string and returns structured information about the browser, OS, device, and bot status.

### Example

```typescript
import {
    UseragentApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new UseragentApi(configuration);

let ua: string; //The user agent string to parse. (default to undefined)

const { status, data } = await apiInstance.v1TechnologyUseragentGet(
    ua
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ua** | [**string**] | The user agent string to parse. | defaults to undefined|


### Return type

**V1TechnologyUseragentGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The ua query parameter is missing. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

