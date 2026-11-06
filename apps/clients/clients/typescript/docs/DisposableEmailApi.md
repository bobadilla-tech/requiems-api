# DisposableEmailApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingDisposableBatchPost**](#v1networkingdisposablebatchpost) | **POST** /v1/networking/disposable/batch | Check Batch Emails|
|[**v1NetworkingDisposableCheckPost**](#v1networkingdisposablecheckpost) | **POST** /v1/networking/disposable/check | Check Single Email|
|[**v1NetworkingDisposableDomainDomainGet**](#v1networkingdisposabledomaindomainget) | **GET** /v1/networking/disposable/domain/{domain} | Check Domain|
|[**v1NetworkingDisposableDomainsGet**](#v1networkingdisposabledomainsget) | **GET** /v1/networking/disposable/domains | List Domains (Paginated)|
|[**v1NetworkingDisposableStatsGet**](#v1networkingdisposablestatsget) | **GET** /v1/networking/disposable/stats | Get Statistics|

# **v1NetworkingDisposableBatchPost**
> V1NetworkingDisposableBatchPost200Response v1NetworkingDisposableBatchPost(v1NetworkingDisposableBatchPostRequest)

Validate multiple email addresses in a single request (max 100 emails)

### Example

```typescript
import {
    DisposableEmailApi,
    Configuration,
    V1NetworkingDisposableBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DisposableEmailApi(configuration);

let v1NetworkingDisposableBatchPostRequest: V1NetworkingDisposableBatchPostRequest; //

const { status, data } = await apiInstance.v1NetworkingDisposableBatchPost(
    v1NetworkingDisposableBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1NetworkingDisposableBatchPostRequest** | **V1NetworkingDisposableBatchPostRequest**|  | |


### Return type

**V1NetworkingDisposableBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing or malformed; The emails field is missing; Too many emails in the request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1NetworkingDisposableCheckPost**
> V1NetworkingDisposableCheckPost200Response v1NetworkingDisposableCheckPost(v1NetworkingDisposableCheckPostRequest)

Validate whether an email address uses a disposable domain

### Example

```typescript
import {
    DisposableEmailApi,
    Configuration,
    V1NetworkingDisposableCheckPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DisposableEmailApi(configuration);

let v1NetworkingDisposableCheckPostRequest: V1NetworkingDisposableCheckPostRequest; //

const { status, data } = await apiInstance.v1NetworkingDisposableCheckPost(
    v1NetworkingDisposableCheckPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1NetworkingDisposableCheckPostRequest** | **V1NetworkingDisposableCheckPostRequest**|  | |


### Return type

**V1NetworkingDisposableCheckPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing or malformed; The email address format is invalid |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1NetworkingDisposableDomainDomainGet**
> V1NetworkingDisposableDomainDomainGet200Response v1NetworkingDisposableDomainDomainGet()

Check if a specific domain is in the disposable blocklist

### Example

```typescript
import {
    DisposableEmailApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DisposableEmailApi(configuration);

let domain: string; //The domain to check (default to undefined)

const { status, data } = await apiInstance.v1NetworkingDisposableDomainDomainGet(
    domain
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **domain** | [**string**] | The domain to check | defaults to undefined|


### Return type

**V1NetworkingDisposableDomainDomainGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The domain parameter is missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1NetworkingDisposableDomainsGet**
> V1NetworkingDisposableDomainsGet200Response v1NetworkingDisposableDomainsGet()

Get a paginated list of all disposable domains in the blocklist

### Example

```typescript
import {
    DisposableEmailApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DisposableEmailApi(configuration);

let page: number; //Page number (default: 1) (optional) (default to undefined)
let perPage: number; //Items per page (default: 100, max: 1000) (optional) (default to undefined)

const { status, data } = await apiInstance.v1NetworkingDisposableDomainsGet(
    page,
    perPage
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **page** | [**number**] | Page number (default: 1) | (optional) defaults to undefined|
| **perPage** | [**number**] | Items per page (default: 100, max: 1000) | (optional) defaults to undefined|


### Return type

**V1NetworkingDisposableDomainsGet200Response**

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

# **v1NetworkingDisposableStatsGet**
> V1NetworkingDisposableStatsGet200Response v1NetworkingDisposableStatsGet()

Get statistics about the disposable email blocklist

### Example

```typescript
import {
    DisposableEmailApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DisposableEmailApi(configuration);

const { status, data } = await apiInstance.v1NetworkingDisposableStatsGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1NetworkingDisposableStatsGet200Response**

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

