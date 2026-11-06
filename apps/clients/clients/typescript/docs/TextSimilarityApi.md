# TextSimilarityApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextSimilarityPost**](#v1textsimilaritypost) | **POST** /v1/text/similarity | Compare Text Similarity|

# **v1TextSimilarityPost**
> V1TextSimilarityPost200Response v1TextSimilarityPost(v1TextSimilarityPostRequest)

Compares two texts and returns a cosine similarity score.

### Example

```typescript
import {
    TextSimilarityApi,
    Configuration,
    V1TextSimilarityPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new TextSimilarityApi(configuration);

let v1TextSimilarityPostRequest: V1TextSimilarityPostRequest; //

const { status, data } = await apiInstance.v1TextSimilarityPost(
    v1TextSimilarityPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextSimilarityPostRequest** | **V1TextSimilarityPostRequest**|  | |


### Return type

**V1TextSimilarityPost200Response**

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
|**422** | One or both text fields are missing or empty. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

