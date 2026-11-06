# BarcodeApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyBarcodeBase64Get**](#v1technologybarcodebase64get) | **GET** /v1/technology/barcode/base64 | Generate Barcode (Base64 JSON)|
|[**v1TechnologyBarcodeGet**](#v1technologybarcodeget) | **GET** /v1/technology/barcode | Generate Barcode (PNG)|

# **v1TechnologyBarcodeBase64Get**
> V1TechnologyBarcodeBase64Get200Response v1TechnologyBarcodeBase64Get()

Returns a JSON envelope containing the barcode as a base64-encoded PNG string, along with its type and dimensions.

### Example

```typescript
import {
    BarcodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new BarcodeApi(configuration);

let data: string; //The text or numeric string to encode in the barcode (default to undefined)
let type: string; //Barcode format: code128, code93, code39, ean8, ean13 (default to undefined)

const { status, data } = await apiInstance.v1TechnologyBarcodeBase64Get(
    data,
    type
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **data** | [**string**] | The text or numeric string to encode in the barcode | defaults to undefined|
| **type** | [**string**] | Barcode format: code128, code93, code39, ean8, ean13 | defaults to undefined|


### Return type

**V1TechnologyBarcodeBase64Get200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Missing or invalid parameters (e.g. data not provided, unsupported type) |  -  |
|**422** | Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyBarcodeGet**
> v1TechnologyBarcodeGet()

Returns a raw PNG image of the barcode. Ideal for direct embedding or file download.

### Example

```typescript
import {
    BarcodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new BarcodeApi(configuration);

let data: string; //The text or numeric string to encode in the barcode (default to undefined)
let type: string; //Barcode format: code128, code93, code39, ean8, ean13 (default to undefined)

const { status, data } = await apiInstance.v1TechnologyBarcodeGet(
    data,
    type
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **data** | [**string**] | The text or numeric string to encode in the barcode | defaults to undefined|
| **type** | [**string**] | Barcode format: code128, code93, code39, ean8, ean13 | defaults to undefined|


### Return type

void (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Missing or invalid parameters (e.g. data not provided, unsupported type) |  -  |
|**422** | Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

