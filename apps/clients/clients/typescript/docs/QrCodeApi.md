# QrCodeApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1TechnologyQrBase64Get**](#v1technologyqrbase64get) | **GET** /v1/technology/qr/base64 | Generate QR Code (Base64 JSON)|
|[**v1TechnologyQrGet**](#v1technologyqrget) | **GET** /v1/technology/qr | Generate QR Code (PNG)|

# **v1TechnologyQrBase64Get**
> V1TechnologyQrBase64Get200Response v1TechnologyQrBase64Get()

Returns a JSON envelope containing the QR code as a base64-encoded PNG string, along with its dimensions.

### Example

```typescript
import {
    QrCodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new QrCodeApi(configuration);

let data: string; //The text or URL to encode in the QR code (default to undefined)
let size: number; //Image size in pixels (default: 256, min: 50, max: 1000) (optional) (default to undefined)
let recovery: string; //Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium (optional) (default to undefined)

const { status, data } = await apiInstance.v1TechnologyQrBase64Get(
    data,
    size,
    recovery
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **data** | [**string**] | The text or URL to encode in the QR code | defaults to undefined|
| **size** | [**number**] | Image size in pixels (default: 256, min: 50, max: 1000) | (optional) defaults to undefined|
| **recovery** | [**string**] | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium | (optional) defaults to undefined|


### Return type

**V1TechnologyQrBase64Get200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level) |  -  |
|**500** | Failed to generate QR code |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1TechnologyQrGet**
> v1TechnologyQrGet()

Returns a raw PNG image of the QR code. Ideal for direct embedding or file download.

### Example

```typescript
import {
    QrCodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new QrCodeApi(configuration);

let data: string; //The text or URL to encode in the QR code (default to undefined)
let size: number; //Image size in pixels (default: 256, min: 50, max: 1000) (optional) (default to undefined)
let recovery: string; //Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium (optional) (default to undefined)

const { status, data } = await apiInstance.v1TechnologyQrGet(
    data,
    size,
    recovery
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **data** | [**string**] | The text or URL to encode in the QR code | defaults to undefined|
| **size** | [**number**] | Image size in pixels (default: 256, min: 50, max: 1000) | (optional) defaults to undefined|
| **recovery** | [**string**] | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium | (optional) defaults to undefined|


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
|**400** | Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level) |  -  |
|**500** | Failed to generate QR code |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

