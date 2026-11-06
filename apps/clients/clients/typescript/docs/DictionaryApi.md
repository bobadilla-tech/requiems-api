# DictionaryApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextDictionaryWordGet**](#v1textdictionarywordget) | **GET** /v1/text/dictionary/{word} | Dictionary Lookup|

# **v1TextDictionaryWordGet**
> V1TextDictionaryWordGet200Response v1TextDictionaryWordGet()

Returns the definition, phonetics, examples, and synonyms for the given word.

### Example

```typescript
import {
    DictionaryApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DictionaryApi(configuration);

let word: string; //The word to look up in the dictionary (default to undefined)

const { status, data } = await apiInstance.v1TextDictionaryWordGet(
    word
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **word** | [**string**] | The word to look up in the dictionary | defaults to undefined|


### Return type

**V1TextDictionaryWordGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The word path parameter is missing. |  -  |
|**404** | The word was not found in the dictionary dataset. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

