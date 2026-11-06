# EmailNormalizeApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextNormalizeBatchPost**](#v1textnormalizebatchpost) | **POST** /v1/text/normalize/batch | Normalize Email Batch|
|[**v1TextNormalizePost**](#v1textnormalizepost) | **POST** /v1/text/normalize | Normalize Email|

# **v1TextNormalizeBatchPost**
> V1TextNormalizeBatchPost200Response v1TextNormalizeBatchPost(v1TextNormalizeBatchPostRequest)

Normalizes up to 100 email addresses in one request. Results are in the same order as the input. Each item includes valid (boolean); when false, only original and message are set. Usage is billed per email processed (see gateway usage headers).

### Example

```typescript
import {
    EmailNormalizeApi,
    Configuration,
    V1TextNormalizeBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new EmailNormalizeApi(configuration);

let v1TextNormalizeBatchPostRequest: V1TextNormalizeBatchPostRequest; //

const { status, data } = await apiInstance.v1TextNormalizeBatchPost(
    v1TextNormalizeBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextNormalizeBatchPostRequest** | **V1TextNormalizeBatchPostRequest**|  | |


### Return type

**V1TextNormalizeBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid JSON or unknown fields in the body |  -  |
|**422** | Missing emails, empty array, too many items, or empty string in the array |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TextNormalizePost**
> V1TextNormalizePost200Response v1TextNormalizePost(v1TextNormalizePostRequest)

Normalizes a single email address and returns the canonical form together with a breakdown of all transformations applied.

### Example

```typescript
import {
    EmailNormalizeApi,
    Configuration,
    V1TextNormalizePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new EmailNormalizeApi(configuration);

let v1TextNormalizePostRequest: V1TextNormalizePostRequest; //

const { status, data } = await apiInstance.v1TextNormalizePost(
    v1TextNormalizePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextNormalizePostRequest** | **V1TextNormalizePostRequest**|  | |


### Return type

**V1TextNormalizePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing, not valid JSON, or contains unknown fields. |  -  |
|**422** | The email field is missing or not a valid email address format. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

