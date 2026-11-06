# ColorConversionApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyColorGet**](#v1technologycolorget) | **GET** /v1/technology/color | Convert Color|

# **v1TechnologyColorGet**
> V1TechnologyColorGet200Response v1TechnologyColorGet()

Convert a color value from one format to another. The response always includes all four formats.

### Example

```typescript
import {
    ColorConversionApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new ColorConversionApi(configuration);

let from: string; //Source color format: hex, rgb, hsl, or cmyk (default to undefined)
let to: string; //Target color format: hex, rgb, hsl, or cmyk (default to undefined)
let value: string; //Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%)) (default to undefined)

const { status, data } = await apiInstance.v1TechnologyColorGet(
    from,
    to,
    value
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **from** | [**string**] | Source color format: hex, rgb, hsl, or cmyk | defaults to undefined|
| **to** | [**string**] | Target color format: hex, rgb, hsl, or cmyk | defaults to undefined|
| **value** | [**string**] | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%)) | defaults to undefined|


### Return type

**V1TechnologyColorGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | One or more of from, to, or value parameters is missing or the from/to value is not one of: hex, rgb, hsl, cmyk. |  -  |
|**422** | The value cannot be parsed in the specified from format. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

