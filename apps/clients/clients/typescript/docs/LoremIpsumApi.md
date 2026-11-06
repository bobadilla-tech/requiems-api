# LoremIpsumApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TextLoremGet**](#v1textloremget) | **GET** /v1/text/lorem | Generate Lorem Ipsum|

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

