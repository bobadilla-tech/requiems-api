# LoremIpsumApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextLoremBatchPost**](#v1textlorembatchpost) | **POST** /v1/text/lorem/batch | Generate Lorem Ipsum (Batch)|
|[**v1TextLoremGet**](#v1textloremget) | **GET** /v1/text/lorem | Generate Lorem Ipsum|

# **v1TextLoremBatchPost**
> V1TextLoremBatchPost200Response v1TextLoremBatchPost(v1TextLoremBatchPostRequest)

Generate multiple Lorem Ipsum placeholder texts in a single request. Processes items in bulk and returns partial successes.

### Example

```typescript
import {
    LoremIpsumApi,
    Configuration,
    V1TextLoremBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new LoremIpsumApi(configuration);

let v1TextLoremBatchPostRequest: V1TextLoremBatchPostRequest; //

const { status, data } = await apiInstance.v1TextLoremBatchPost(
    v1TextLoremBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TextLoremBatchPostRequest** | **V1TextLoremBatchPostRequest**|  | |


### Return type

**V1TextLoremBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**422** | The request body is malformed, missing the items array, or exceeds the maximum batch size of 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TextLoremGet**
> V1TextLoremGet200Response v1TextLoremGet()

Generate Lorem Ipsum placeholder text with customizable length and format

### Example

```typescript
import {
    LoremIpsumApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new LoremIpsumApi(configuration);

let paragraphs: number; //Number of paragraphs to generate (1-20) (optional) (default to undefined)
let sentences: number; //Number of sentences per paragraph (1-20) (optional) (default to undefined)

const { status, data } = await apiInstance.v1TextLoremGet(
    paragraphs,
    sentences
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **paragraphs** | [**number**] | Number of paragraphs to generate (1-20) | (optional) defaults to undefined|
| **sentences** | [**number**] | Number of sentences per paragraph (1-20) | (optional) defaults to undefined|


### Return type

**V1TextLoremGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The paragraphs parameter is out of valid range; The sentences parameter is out of valid range |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

