# ThesaurusApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextThesaurusWordGet**](#v1textthesauruswordget) | **GET** /v1/text/thesaurus/{word} | Thesaurus Lookup|

# **v1TextThesaurusWordGet**
> V1TextThesaurusWordGet200Response v1TextThesaurusWordGet()

Returns synonyms and antonyms for the given word.

### Example

```typescript
import {
    ThesaurusApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ThesaurusApi(configuration);

let word: string; //The word to look up in the thesaurus (default to undefined)

const { status, data } = await apiInstance.v1TextThesaurusWordGet(
    word
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **word** | [**string**] | The word to look up in the thesaurus | defaults to undefined|


### Return type

**V1TextThesaurusWordGet200Response**

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
|**404** | The word was not found in the thesaurus dataset. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

