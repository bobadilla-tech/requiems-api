# CounterApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyCounterNamespaceGet**](#v1technologycounternamespaceget) | **GET** /v1/technology/counter/{namespace} | Get Counter Value|
|[**v1TechnologyCounterNamespacePost**](#v1technologycounternamespacepost) | **POST** /v1/technology/counter/{namespace} | Increment Counter|

# **v1TechnologyCounterNamespaceGet**
> V1TechnologyCounterNamespaceGet200Response v1TechnologyCounterNamespaceGet()

Get the current value of a counter without incrementing it

### Example

```typescript
import {
    CounterApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new CounterApi(configuration);

let namespace: string; //Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) (default to undefined)

const { status, data } = await apiInstance.v1TechnologyCounterNamespaceGet(
    namespace
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **namespace** | [**string**] | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) | defaults to undefined|


### Return type

**V1TechnologyCounterNamespaceGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyCounterNamespacePost**
> V1TechnologyCounterNamespacePost200Response v1TechnologyCounterNamespacePost()

Atomically increment a counter in the specified namespace and return the new value

### Example

```typescript
import {
    CounterApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new CounterApi(configuration);

let namespace: string; //Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) (default to undefined)

const { status, data } = await apiInstance.v1TechnologyCounterNamespacePost(
    namespace
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **namespace** | [**string**] | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) | defaults to undefined|


### Return type

**V1TechnologyCounterNamespacePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

