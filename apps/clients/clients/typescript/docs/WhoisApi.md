# WhoisApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingWhoisBatchPost**](#v1networkingwhoisbatchpost) | **POST** /v1/networking/whois/batch | Batch WHOIS Lookup|
|[**v1NetworkingWhoisDomainGet**](#v1networkingwhoisdomainget) | **GET** /v1/networking/whois/{domain} | WHOIS Lookup|

# **v1NetworkingWhoisBatchPost**
> V1NetworkingWhoisBatchPost200Response v1NetworkingWhoisBatchPost(v1NetworkingWhoisBatchPostRequest)

Returns WHOIS information for up to 50 domains in a single request. Results are returned in the same order as the input array. Domains without WHOIS data return found: false instead of failing the entire request.

### Example

```typescript
import {
    WhoisApi,
    Configuration,
    V1NetworkingWhoisBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new WhoisApi(configuration);

let v1NetworkingWhoisBatchPostRequest: V1NetworkingWhoisBatchPostRequest; //

const { status, data } = await apiInstance.v1NetworkingWhoisBatchPost(
    v1NetworkingWhoisBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1NetworkingWhoisBatchPostRequest** | **V1NetworkingWhoisBatchPostRequest**|  | |


### Return type

**V1NetworkingWhoisBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**422** | Invalid request body. This includes empty arrays, more than 50 domains, invalid domain names, or malformed payloads. |  -  |
|**500** | Unexpected server error during batch WHOIS lookup. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1NetworkingWhoisDomainGet**
> V1NetworkingWhoisDomainGet200Response v1NetworkingWhoisDomainGet()

Returns WHOIS registration information for a domain name.

### Example

```typescript
import {
    WhoisApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new WhoisApi(configuration);

let domain: string; //The domain name to look up (e.g. example.com) (default to undefined)

const { status, data } = await apiInstance.v1NetworkingWhoisDomainGet(
    domain
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **domain** | [**string**] | The domain name to look up (e.g. example.com) | defaults to undefined|


### Return type

**V1NetworkingWhoisDomainGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The domain name format is invalid. |  -  |
|**404** | No WHOIS record was found for the domain. |  -  |
|**500** | Unexpected server error or upstream WHOIS query failure. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

