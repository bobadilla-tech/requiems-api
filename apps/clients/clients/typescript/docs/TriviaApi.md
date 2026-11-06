# TriviaApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentTriviaGet**](#v1entertainmenttriviaget) | **GET** /v1/entertainment/trivia | Get Trivia Question|

# **v1EntertainmentTriviaGet**
> V1EntertainmentTriviaGet200Response v1EntertainmentTriviaGet()

Returns a random trivia question with multiple-choice answers. Use the optional category and difficulty query parameters to filter the question pool.

### Example

```typescript
import {
    TriviaApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new TriviaApi(configuration);

let category: string; //Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. (optional) (default to undefined)
let difficulty: string; //Filter by difficulty. One of: easy, medium, hard. (optional) (default to undefined)

const { status, data } = await apiInstance.v1EntertainmentTriviaGet(
    category,
    difficulty
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **category** | [**string**] | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. | (optional) defaults to undefined|
| **difficulty** | [**string**] | Filter by difficulty. One of: easy, medium, hard. | (optional) defaults to undefined|


### Return type

**V1EntertainmentTriviaGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | An invalid category or difficulty value was provided |  -  |
|**401** | Missing API key |  -  |
|**403** | Invalid or revoked API key |  -  |
|**404** | No questions match the given category and difficulty combination |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

