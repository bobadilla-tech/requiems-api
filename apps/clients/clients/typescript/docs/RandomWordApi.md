# RandomWordApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextWordsBatchPost**](#v1textwordsbatchpost) | **POST** /v1/text/words/batch | Batch Define Words|
|[**v1TextWordsRandomGet**](#v1textwordsrandomget) | **GET** /v1/text/words/random | Get Random Word|

# **v1TextWordsBatchPost**
> V1TextWordsBatchPost200Response v1TextWordsBatchPost(v1TextWordsBatchPostRequest)

Resolve multiple words in a single request. Returns dictionary entries when found, or error information when a word is not in the dataset. 

### Example

```typescript
import {
    RandomWordApi,
    Configuration,
    V1TextWordsBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new RandomWordApi(configuration);

let v1TextWordsBatchPostRequest: V1TextWordsBatchPostRequest; //

const { status, data } = await apiInstance.v1TextWordsBatchPost(
    v1TextWordsBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextWordsBatchPostRequest** | **V1TextWordsBatchPostRequest**|  | |


### Return type

**V1TextWordsBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**422** | Invalid or empty items array |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TextWordsRandomGet**
> V1TextWordsRandomGet200Response v1TextWordsRandomGet()

Returns a random word with its definition and part of speech

### Example

```typescript
import {
    RandomWordApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new RandomWordApi(configuration);

const { status, data } = await apiInstance.v1TextWordsRandomGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1TextWordsRandomGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**503** | No words available in the database |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

