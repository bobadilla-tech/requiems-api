# IpAsnApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingIpAsnGet**](#v1networkingipasnget) | **GET** /v1/networking/ip/asn | Lookup ASN (Caller IP)|
|[**v1NetworkingIpAsnIpGet**](#v1networkingipasnipget) | **GET** /v1/networking/ip/asn/{ip} | Lookup ASN for IP|

# **v1NetworkingIpAsnGet**
> V1NetworkingIpAsnGet200Response v1NetworkingIpAsnGet()

Look up ASN, organization, ISP, and network details for the requesting client\'s IP address. Useful when you want information about the user making the request without specifying an IP explicitly.

### Example

```typescript
import {
    IpAsnApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IpAsnApi(configuration);

const { status, data } = await apiInstance.v1NetworkingIpAsnGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1NetworkingIpAsnGet200Response**

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

# **v1NetworkingIpAsnIpGet**
> V1NetworkingIpAsnIpGet200Response v1NetworkingIpAsnIpGet()

Look up ASN, organization, ISP, and network details for a specific IP address.

### Example

```typescript
import {
    IpAsnApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IpAsnApi(configuration);

let ip: string; //The IP address to look up (supports IPv4 and IPv6) (default to undefined)

const { status, data } = await apiInstance.v1NetworkingIpAsnIpGet(
    ip
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ip** | [**string**] | The IP address to look up (supports IPv4 and IPv6) | defaults to undefined|


### Return type

**V1NetworkingIpAsnIpGet200Response**

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

