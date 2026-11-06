# ProfanityApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1ValidationProfanityBatchPost**](#v1validationprofanitybatchpost) | **POST** /v1/validation/profanity/batch | Batch Check Profanity|
|[**v1ValidationProfanityPost**](#v1validationprofanitypost) | **POST** /v1/validation/profanity | Check Profanity|

# **v1ValidationProfanityBatchPost**
> V1ValidationProfanityBatchPost200Response v1ValidationProfanityBatchPost(v1ValidationProfanityBatchPostRequest)

Check up to 50 texts for profanity in a single request. Results are returned in input order.

### Example

```typescript
import {
    ProfanityApi,
    Configuration,
    V1ValidationProfanityBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ProfanityApi(configuration);

let v1ValidationProfanityBatchPostRequest: V1ValidationProfanityBatchPostRequest; //

const { status, data } = await apiInstance.v1ValidationProfanityBatchPost(
    v1ValidationProfanityBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1ValidationProfanityBatchPostRequest** | **V1ValidationProfanityBatchPostRequest**|  | |


### Return type

**V1ValidationProfanityBatchPost200Response**

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
|**422** | The texts array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1ValidationProfanityPost**
> V1ValidationProfanityPost200Response v1ValidationProfanityPost(v1ValidationProfanityPostRequest)

Checks text for profanity, returning a censored version and the list of flagged words.

### Example

```typescript
import {
    ProfanityApi,
    Configuration,
    V1ValidationProfanityPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ProfanityApi(configuration);

let v1ValidationProfanityPostRequest: V1ValidationProfanityPostRequest; //

const { status, data } = await apiInstance.v1ValidationProfanityPost(
    v1ValidationProfanityPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1ValidationProfanityPostRequest** | **V1ValidationProfanityPostRequest**|  | |


### Return type

**V1ValidationProfanityPost200Response**

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
|**422** | The text field is missing or empty. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

