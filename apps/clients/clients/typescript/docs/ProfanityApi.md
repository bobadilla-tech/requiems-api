# ProfanityApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextProfanityPost**](#v1textprofanitypost) | **POST** /v1/text/profanity | Check Profanity|

# **v1TextProfanityPost**
> V1TextProfanityPost200Response v1TextProfanityPost(v1TextProfanityPostRequest)

Checks text for profanity, returning a censored version and the list of flagged words.

### Example

```typescript
import {
    ProfanityApi,
    Configuration,
    V1TextProfanityPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ProfanityApi(configuration);

let v1TextProfanityPostRequest: V1TextProfanityPostRequest; //

const { status, data } = await apiInstance.v1TextProfanityPost(
    v1TextProfanityPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextProfanityPostRequest** | **V1TextProfanityPostRequest**|  | |


### Return type

**V1TextProfanityPost200Response**

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

