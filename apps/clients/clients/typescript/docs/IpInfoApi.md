# IpInfoApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingIpGet**](#v1networkingipget) | **GET** /v1/networking/ip | Get IP Info (Caller IP)|
|[**v1NetworkingIpInfoBatchPost**](#v1networkingipinfobatchpost) | **POST** /v1/networking/ip/info/batch | Batch IP Geolocation|
|[**v1NetworkingIpIpGet**](#v1networkingipipget) | **GET** /v1/networking/ip/{ip} | Get IP Info for IP|

# **v1NetworkingIpGet**
> V1NetworkingIpGet200Response v1NetworkingIpGet()

Get geolocation and network information for the requesting client\'s IP address. Useful when you want information about the user making the request without specifying an IP explicitly.

### Example

```typescript
import {
    IpInfoApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IpInfoApi(configuration);

const { status, data } = await apiInstance.v1NetworkingIpGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1NetworkingIpGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1NetworkingIpInfoBatchPost**
> V1NetworkingIpInfoBatchPost200Response v1NetworkingIpInfoBatchPost(v1NetworkingIpAsnBatchPostRequest)

Look up geolocation data for up to 50 IP addresses in a single request. Results are returned in input order. Per-item errors are reported inline.

### Example

```typescript
import {
    IpInfoApi,
    Configuration,
    V1NetworkingIpAsnBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IpInfoApi(configuration);

let v1NetworkingIpAsnBatchPostRequest: V1NetworkingIpAsnBatchPostRequest; //

const { status, data } = await apiInstance.v1NetworkingIpInfoBatchPost(
    v1NetworkingIpAsnBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1NetworkingIpAsnBatchPostRequest** | **V1NetworkingIpAsnBatchPostRequest**|  | |


### Return type

**V1NetworkingIpInfoBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing or malformed. |  -  |
|**422** | The ips array is missing, empty, exceeds 50 items, or contains an invalid IP address. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1NetworkingIpIpGet**
> V1NetworkingIpIpGet200Response v1NetworkingIpIpGet()

Get geolocation and network information for a specific IP address.

### Example

```typescript
import {
    IpInfoApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IpInfoApi(configuration);

let ip: string; //The IP address to look up (supports IPv4 and IPv6) (default to undefined)

const { status, data } = await apiInstance.v1NetworkingIpIpGet(
    ip
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ip** | [**string**] | The IP address to look up (supports IPv4 and IPv6) | defaults to undefined|


### Return type

**V1NetworkingIpIpGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The IP address is invalid |  -  |
|**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

