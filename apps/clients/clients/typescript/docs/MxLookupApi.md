# MxLookupApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingMxDomainGet**](#v1networkingmxdomainget) | **GET** /v1/networking/mx/{domain} | MX Lookup|

# **v1NetworkingMxDomainGet**
> V1NetworkingMxDomainGet200Response v1NetworkingMxDomainGet()

Retrieve all MX records for a domain. Results are sorted by priority ascending (lowest numeric value has highest mail delivery priority per RFC 5321).

### Example

```typescript
import {
    MxLookupApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new MxLookupApi(configuration);

let domain: string; //The domain name to look up MX records for (e.g. gmail.com) (default to undefined)

const { status, data } = await apiInstance.v1NetworkingMxDomainGet(
    domain
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **domain** | [**string**] | The domain name to look up MX records for (e.g. gmail.com) | defaults to undefined|


### Return type

**V1NetworkingMxDomainGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The domain parameter is not a valid domain name. |  -  |
|**404** | No MX records were found for the domain (domain may not accept email). |  -  |
|**500** | DNS lookup failed due to an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

