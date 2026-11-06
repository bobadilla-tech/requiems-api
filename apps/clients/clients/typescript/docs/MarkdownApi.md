# MarkdownApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyMarkdownPost**](#v1technologymarkdownpost) | **POST** /v1/technology/markdown | Convert Markdown to HTML|

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

