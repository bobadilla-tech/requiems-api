# DomainInfoApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1NetworkingDomainDomainGet**](#v1networkingdomaindomainget) | **GET** /v1/networking/domain/{domain} | Get Domain Info|

# **v1NetworkingDomainDomainGet**
> V1NetworkingDomainDomainGet200Response v1NetworkingDomainDomainGet()

Returns DNS records and availability status for the given domain.

### Example

```typescript
import {
    DomainInfoApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DomainInfoApi(configuration);

let domain: string; //The domain to look up (e.g. example.com) (default to undefined)

const { status, data } = await apiInstance.v1NetworkingDomainDomainGet(
    domain
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **domain** | [**string**] | The domain to look up (e.g. example.com) | defaults to undefined|


### Return type

**V1NetworkingDomainDomainGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The domain parameter is not a valid hostname (e.g. missing TLD, invalid characters, or leading/trailing hyphens). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

