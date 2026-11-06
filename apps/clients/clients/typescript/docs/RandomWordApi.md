# RandomWordApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextWordsRandomGet**](#v1textwordsrandomget) | **GET** /v1/text/words/random | Get Random Word|

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

