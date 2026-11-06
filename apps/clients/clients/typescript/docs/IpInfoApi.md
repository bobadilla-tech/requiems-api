# IpInfoApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingIpGet**](#v1networkingipget) | **GET** /v1/networking/ip | Get IP Info (Caller IP)|
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

