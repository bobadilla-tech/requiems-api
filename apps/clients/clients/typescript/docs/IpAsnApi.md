# IpAsnApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingIpAsnBatchPost**](#v1networkingipasnbatchpost) | **POST** /v1/networking/ip/asn/batch | Batch ASN Lookup|
|[**v1NetworkingIpAsnGet**](#v1networkingipasnget) | **GET** /v1/networking/ip/asn | Lookup ASN (Caller IP)|
|[**v1NetworkingIpAsnIpGet**](#v1networkingipasnipget) | **GET** /v1/networking/ip/asn/{ip} | Lookup ASN for IP|

# **v1NetworkingIpAsnBatchPost**
> V1NetworkingIpAsnBatchPost200Response v1NetworkingIpAsnBatchPost(v1NetworkingIpAsnBatchPostRequest)

Look up ASN data for up to 50 IP addresses in a single request. Results are returned in input order. Private and reserved IPs return an empty result with no error.

### Example

```typescript
import {
    IpAsnApi,
    Configuration,
    V1NetworkingIpAsnBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new IpAsnApi(configuration);

let v1NetworkingIpAsnBatchPostRequest: V1NetworkingIpAsnBatchPostRequest; //

const { status, data } = await apiInstance.v1NetworkingIpAsnBatchPost(
    v1NetworkingIpAsnBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1NetworkingIpAsnBatchPostRequest** | **V1NetworkingIpAsnBatchPostRequest**|  | |


### Return type

**V1NetworkingIpAsnBatchPost200Response**

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

