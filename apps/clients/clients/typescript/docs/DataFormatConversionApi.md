# DataFormatConversionApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyFormatPost**](#v1technologyformatpost) | **POST** /v1/technology/format | Convert Format|

# **v1TechnologyFormatPost**
> V1TechnologyFormatPost200Response v1TechnologyFormatPost(v1TechnologyFormatPostRequest)

Convert content from one structured data format to another. Supported formats: json, yaml, csv, xml, toml.

### Example

```typescript
import {
    DataFormatConversionApi,
    Configuration,
    V1TechnologyFormatPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DataFormatConversionApi(configuration);

let v1TechnologyFormatPostRequest: V1TechnologyFormatPostRequest; //

const { status, data } = await apiInstance.v1TechnologyFormatPost(
    v1TechnologyFormatPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyFormatPostRequest** | **V1TechnologyFormatPostRequest**|  | |


### Return type

**V1TechnologyFormatPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**413** | The content field exceeds the 512 KB limit. |  -  |
|**422** | One of from, to, or content is missing, or from/to is not one of the supported format values.; The content field is not valid JSON (when from is json).; The content field is not valid YAML (when from is yaml).; The content field is not valid CSV, or a row has more columns than the header (when from is csv).; The content field is not valid XML (when from is xml).; The content field is not valid TOML (when from is toml).; The data structure is incompatible with the target format (e.g. converting a JSON array to TOML, which requires a top-level object). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

