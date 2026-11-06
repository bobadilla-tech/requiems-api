# DetectLanguageApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextDetectLanguagePost**](#v1textdetectlanguagepost) | **POST** /v1/text/detect-language | Detect Language|

# **v1TextDetectLanguagePost**
> V1TextDetectLanguagePost200Response v1TextDetectLanguagePost(v1TextDetectLanguagePostRequest)

Identifies the language of the provided text and returns the language name, ISO 639-1 code, and confidence score.

### Example

```typescript
import {
    DetectLanguageApi,
    Configuration,
    V1TextDetectLanguagePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DetectLanguageApi(configuration);

let v1TextDetectLanguagePostRequest: V1TextDetectLanguagePostRequest; //

const { status, data } = await apiInstance.v1TextDetectLanguagePost(
    v1TextDetectLanguagePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextDetectLanguagePostRequest** | **V1TextDetectLanguagePostRequest**|  | |


### Return type

**V1TextDetectLanguagePost200Response**

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

