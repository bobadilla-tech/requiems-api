# NumberBaseConversionApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyBaseGet**](#v1technologybaseget) | **GET** /v1/technology/base | Convert Base|

# **v1TechnologyBaseGet**
> V1TechnologyBaseGet200Response v1TechnologyBaseGet()

Convert an integer from one number base to another.

### Example

```typescript
import {
    NumberBaseConversionApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new NumberBaseConversionApi(configuration);

let from: number; //Source base (2, 8, 10, or 16) (default to undefined)
let to: number; //Target base (2, 8, 10, or 16) (default to undefined)
let value: string; //The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). (default to undefined)

const { status, data } = await apiInstance.v1TechnologyBaseGet(
    from,
    to,
    value
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **from** | [**number**] | Source base (2, 8, 10, or 16) | defaults to undefined|
| **to** | [**number**] | Target base (2, 8, 10, or 16) | defaults to undefined|
| **value** | [**string**] | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). | defaults to undefined|


### Return type

**V1TechnologyBaseGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | A required parameter is missing, the base is not one of 2/8/10/16, or value is not valid for the given base. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

