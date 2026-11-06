# SpellCheckApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextSpellcheckPost**](#v1textspellcheckpost) | **POST** /v1/text/spellcheck | Check Spelling|

# **v1TextSpellcheckPost**
> V1TextSpellcheckPost200Response v1TextSpellcheckPost(v1TextSpellcheckPostRequest)

Checks the input text for spelling mistakes and returns a corrected version along with per-word corrections.

### Example

```typescript
import {
    SpellCheckApi,
    Configuration,
    V1TextSpellcheckPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SpellCheckApi(configuration);

let v1TextSpellcheckPostRequest: V1TextSpellcheckPostRequest; //

const { status, data } = await apiInstance.v1TextSpellcheckPost(
    v1TextSpellcheckPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextSpellcheckPostRequest** | **V1TextSpellcheckPostRequest**|  | |


### Return type

**V1TextSpellcheckPost200Response**

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

