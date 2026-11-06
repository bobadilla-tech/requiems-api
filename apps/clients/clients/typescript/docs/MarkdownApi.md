# MarkdownApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyMarkdownBatchPost**](#v1technologymarkdownbatchpost) | **POST** /v1/technology/markdown/batch | Convert Markdown Batch|
|[**v1TechnologyMarkdownPost**](#v1technologymarkdownpost) | **POST** /v1/technology/markdown | Convert Markdown to HTML|

# **v1TechnologyMarkdownBatchPost**
> V1TechnologyMarkdownBatchPost200Response v1TechnologyMarkdownBatchPost(v1TechnologyMarkdownBatchPostRequest)

Converts multiple Markdown strings into HTML in a single request. Results are returned in the same order as the input array. Supports optional sanitization to remove unsafe HTML tags such as script and iframe.

### Example

```typescript
import {
    MarkdownApi,
    Configuration,
    V1TechnologyMarkdownBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new MarkdownApi(configuration);

let v1TechnologyMarkdownBatchPostRequest: V1TechnologyMarkdownBatchPostRequest; //

const { status, data } = await apiInstance.v1TechnologyMarkdownBatchPost(
    v1TechnologyMarkdownBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyMarkdownBatchPostRequest** | **V1TechnologyMarkdownBatchPostRequest**|  | |


### Return type

**V1TechnologyMarkdownBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Missing or malformed request body. |  -  |
|**422** | Body is invalid: empty array, more than 50 items, or invalid Markdown entry. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyMarkdownPost**
> V1TechnologyMarkdownPost200Response v1TechnologyMarkdownPost(v1TechnologyMarkdownPostRequest)

Converts a Markdown string to HTML. Pass sanitize true to strip potentially unsafe tags like script and iframe from the output.

### Example

```typescript
import {
    MarkdownApi,
    Configuration,
    V1TechnologyMarkdownPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new MarkdownApi(configuration);

let v1TechnologyMarkdownPostRequest: V1TechnologyMarkdownPostRequest; //

const { status, data } = await apiInstance.v1TechnologyMarkdownPost(
    v1TechnologyMarkdownPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyMarkdownPostRequest** | **V1TechnologyMarkdownPostRequest**|  | |


### Return type

**V1TechnologyMarkdownPost200Response**

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

