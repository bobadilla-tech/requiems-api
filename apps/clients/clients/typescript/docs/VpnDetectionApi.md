# VpnDetectionApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingIpVpnIpGet**](#v1networkingipvpnipget) | **GET** /v1/networking/ip/vpn/{ip} | Check IP Address|

# **v1NetworkingIpVpnIpGet**
> V1NetworkingIpVpnIpGet200Response v1NetworkingIpVpnIpGet()

Analyze an IP address to determine if it\'s a VPN, proxy, Tor exit node, or hosting provider. Returns detailed threat indicators and scores.

### Example

```typescript
import {
    VpnDetectionApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new VpnDetectionApi(configuration);

let ip: string; //The IP address to check (supports IPv4 and IPv6) (default to undefined)

const { status, data } = await apiInstance.v1NetworkingIpVpnIpGet(
    ip
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ip** | [**string**] | The IP address to check (supports IPv4 and IPv6) | defaults to undefined|


### Return type

**V1NetworkingIpVpnIpGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The IP address is missing or invalid |  -  |
|**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

