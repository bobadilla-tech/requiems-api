# SpellCheckApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextSpellcheckBatchPost**](#v1textspellcheckbatchpost) | **POST** /v1/text/spellcheck/batch | Check Spelling (Batch)|
|[**v1TextSpellcheckPost**](#v1textspellcheckpost) | **POST** /v1/text/spellcheck | Check Spelling|

# **v1TextSpellcheckBatchPost**
> V1TextSpellcheckBatchPost200Response v1TextSpellcheckBatchPost(v1TextSpellcheckBatchPostRequest)

Checks multiple texts for spelling mistakes in a single request. Returns a corrected version and per-word corrections for each input text. Results are returned in the same order as the input array.

### Example

```typescript
import {
    SpellCheckApi,
    Configuration,
    V1TextSpellcheckBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SpellCheckApi(configuration);

let v1TextSpellcheckBatchPostRequest: V1TextSpellcheckBatchPostRequest; //

const { status, data } = await apiInstance.v1TextSpellcheckBatchPost(
    v1TextSpellcheckBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextSpellcheckBatchPostRequest** | **V1TextSpellcheckBatchPostRequest**|  | |


### Return type

**V1TextSpellcheckBatchPost200Response**

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
|**422** | The texts field is missing, empty, or exceeds 50 items. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

