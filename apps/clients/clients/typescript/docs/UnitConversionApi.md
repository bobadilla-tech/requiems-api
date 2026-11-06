# UnitConversionApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyConvertBatchPost**](#v1technologyconvertbatchpost) | **POST** /v1/technology/convert/batch | Batch Convert Units|
|[**v1TechnologyConvertGet**](#v1technologyconvertget) | **GET** /v1/technology/convert | Convert Units|
|[**v1TechnologyConvertUnitsGet**](#v1technologyconvertunitsget) | **GET** /v1/technology/convert/units | List Available Units|

# **v1TechnologyConvertBatchPost**
> V1TechnologyConvertBatchPost200Response v1TechnologyConvertBatchPost(v1TechnologyConvertBatchPostRequest)

convert up to 50 unit conversion operations in a single request.

### Example

```typescript
import {
    UnitConversionApi,
    Configuration,
    V1TechnologyConvertBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new UnitConversionApi(configuration);

let v1TechnologyConvertBatchPostRequest: V1TechnologyConvertBatchPostRequest; //

const { status, data } = await apiInstance.v1TechnologyConvertBatchPost(
    v1TechnologyConvertBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1TechnologyConvertBatchPostRequest** | **V1TechnologyConvertBatchPostRequest**|  | |


### Return type

**V1TechnologyConvertBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid JSON, malformed request body, or unexpected field types. |  -  |
|**422** | The operations array is missing, empty, or contains more than 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyConvertGet**
> V1TechnologyConvertGet200Response v1TechnologyConvertGet()

Convert a value from one unit to another

### Example

```typescript
import {
    UnitConversionApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new UnitConversionApi(configuration);

let from: string; //Source unit key (e.g. miles, kg, c) (default to undefined)
let to: string; //Target unit key (e.g. km, lb, f) (default to undefined)
let value: number; //Numeric value to convert (default to undefined)

const { status, data } = await apiInstance.v1TechnologyConvertGet(
    from,
    to,
    value
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **from** | [**string**] | Source unit key (e.g. miles, kg, c) | defaults to undefined|
| **to** | [**string**] | Target unit key (e.g. km, lb, f) | defaults to undefined|
| **value** | [**number**] | Numeric value to convert | defaults to undefined|


### Return type

**V1TechnologyConvertGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyConvertUnitsGet**
> V1TechnologyConvertUnitsGet200Response v1TechnologyConvertUnitsGet()

Returns all available unit conversion types grouped by measurement category

### Example

```typescript
import {
    UnitConversionApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new UnitConversionApi(configuration);

const { status, data } = await apiInstance.v1TechnologyConvertUnitsGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1TechnologyConvertUnitsGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

