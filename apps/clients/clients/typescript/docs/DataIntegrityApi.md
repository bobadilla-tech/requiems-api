# DataIntegrityApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1SystemsContentModeratePost**](#v1systemscontentmoderatepost) | **POST** /v1/systems/content/moderate | Moderate Content|
|[**v1SystemsTextNormalizePost**](#v1systemstextnormalizepost) | **POST** /v1/systems/text/normalize | Normalize Text|

# **v1SystemsContentModeratePost**
> V1SystemsContentModeratePost200Response v1SystemsContentModeratePost(v1SystemsContentModeratePostRequest)

Check text for profanity, toxicity, sentiment, and language. Returns per-category flags and an is_safe boolean for gating content.

### Example

```typescript
import {
    DataIntegrityApi,
    Configuration,
    V1SystemsContentModeratePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DataIntegrityApi(configuration);

let v1SystemsContentModeratePostRequest: V1SystemsContentModeratePostRequest; //

const { status, data } = await apiInstance.v1SystemsContentModeratePost(
    v1SystemsContentModeratePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsContentModeratePostRequest** | **V1SystemsContentModeratePostRequest**|  | |


### Return type

**V1SystemsContentModeratePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key |  -  |
|**422** | text field is missing or empty |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1SystemsTextNormalizePost**
> V1SystemsTextNormalizePost200Response v1SystemsTextNormalizePost(v1SystemsTextNormalizePostRequest)

Apply a composable pipeline of normalization operations to a string. Trim whitespace, fix encoding, normalize unicode, strip HTML, and more — in a single call.

### Example

```typescript
import {
    DataIntegrityApi,
    Configuration,
    V1SystemsTextNormalizePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DataIntegrityApi(configuration);

let v1SystemsTextNormalizePostRequest: V1SystemsTextNormalizePostRequest; //

const { status, data } = await apiInstance.v1SystemsTextNormalizePost(
    v1SystemsTextNormalizePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsTextNormalizePostRequest** | **V1SystemsTextNormalizePostRequest**|  | |


### Return type

**V1SystemsTextNormalizePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key |  -  |
|**422** | text is empty or operations array is empty or missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

