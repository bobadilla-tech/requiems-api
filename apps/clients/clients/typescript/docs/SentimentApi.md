# SentimentApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextSentimentBatchPost**](#v1textsentimentbatchpost) | **POST** /v1/text/sentiment/batch | Analyze Sentiment (Batch)|
|[**v1TextSentimentPost**](#v1textsentimentpost) | **POST** /v1/text/sentiment | Analyze Sentiment|

# **v1TextSentimentBatchPost**
> V1TextSentimentBatchPost200Response v1TextSentimentBatchPost(v1TextSentimentBatchPostRequest)

Analyzes the sentiment of up to 50 texts in a single request. Results are returned in the same order as the input. Each text counts as one unit of usage.

### Example

```typescript
import {
    SentimentApi,
    Configuration,
    V1TextSentimentBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SentimentApi(configuration);

let v1TextSentimentBatchPostRequest: V1TextSentimentBatchPostRequest; //

const { status, data } = await apiInstance.v1TextSentimentBatchPost(
    v1TextSentimentBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextSentimentBatchPostRequest** | **V1TextSentimentBatchPostRequest**|  | |


### Return type

**V1TextSentimentBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | bad_request |  -  |
|**422** | unprocessable_entity |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TextSentimentPost**
> V1TextSentimentPost200Response v1TextSentimentPost(v1TextSentimentPostRequest)

Analyzes the sentiment of the provided text and returns a classification, confidence score, and a full breakdown across all three sentiment classes.

### Example

```typescript
import {
    SentimentApi,
    Configuration,
    V1TextSentimentPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SentimentApi(configuration);

let v1TextSentimentPostRequest: V1TextSentimentPostRequest; //

const { status, data } = await apiInstance.v1TextSentimentPost(
    v1TextSentimentPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextSentimentPostRequest** | **V1TextSentimentPostRequest**|  | |


### Return type

**V1TextSentimentPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**422** | unprocessable_entity |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

