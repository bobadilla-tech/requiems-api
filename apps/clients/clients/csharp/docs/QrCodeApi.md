# Org.OpenAPITools.Api.QrCodeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyQrBase64Get**](QrCodeApi.md#v1technologyqrbase64get) | **GET** /v1/technology/qr/base64 | Generate QR Code (Base64 JSON) |
| [**V1TechnologyQrGet**](QrCodeApi.md#v1technologyqrget) | **GET** /v1/technology/qr | Generate QR Code (PNG) |

<a id="v1technologyqrbase64get"></a>
# **V1TechnologyQrBase64Get**
> V1TechnologyQrBase64Get200Response V1TechnologyQrBase64Get (string data, int? size = null, string? recovery = null)

Generate QR Code (Base64 JSON)

Returns a JSON envelope containing the QR code as a base64-encoded PNG string, along with its dimensions.

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using System.Net.Http;
using Org.OpenAPITools.Api;
using Org.OpenAPITools.Client;
using Org.OpenAPITools.Model;

namespace Example
{
    public class V1TechnologyQrBase64GetExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://api.requiems.xyz";
            // Configure API key authorization: requiems-api-key
            config.AddApiKey("requiems-api-key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("requiems-api-key", "Bearer");

            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new QrCodeApi(httpClient, config, httpClientHandler);
            var data = https://example.com;  // string | The text or URL to encode in the QR code
            var size = 200;  // int? | Image size in pixels (default: 256, min: 50, max: 1000) (optional) 
            var recovery = highest;  // string? | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium (optional) 

            try
            {
                // Generate QR Code (Base64 JSON)
                V1TechnologyQrBase64Get200Response result = apiInstance.V1TechnologyQrBase64Get(data, size, recovery);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling QrCodeApi.V1TechnologyQrBase64Get: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyQrBase64GetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Generate QR Code (Base64 JSON)
    ApiResponse<V1TechnologyQrBase64Get200Response> response = apiInstance.V1TechnologyQrBase64GetWithHttpInfo(data, size, recovery);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling QrCodeApi.V1TechnologyQrBase64GetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **data** | **string** | The text or URL to encode in the QR code |  |
| **size** | **int?** | Image size in pixels (default: 256, min: 50, max: 1000) | [optional]  |
| **recovery** | **string?** | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium | [optional]  |

### Return type

[**V1TechnologyQrBase64Get200Response**](V1TechnologyQrBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level) |  -  |
| **500** | Failed to generate QR code |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1technologyqrget"></a>
# **V1TechnologyQrGet**
> void V1TechnologyQrGet (string data, int? size = null, string? recovery = null)

Generate QR Code (PNG)

Returns a raw PNG image of the QR code. Ideal for direct embedding or file download.

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using System.Net.Http;
using Org.OpenAPITools.Api;
using Org.OpenAPITools.Client;
using Org.OpenAPITools.Model;

namespace Example
{
    public class V1TechnologyQrGetExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://api.requiems.xyz";
            // Configure API key authorization: requiems-api-key
            config.AddApiKey("requiems-api-key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("requiems-api-key", "Bearer");

            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new QrCodeApi(httpClient, config, httpClientHandler);
            var data = https://example.com;  // string | The text or URL to encode in the QR code
            var size = 200;  // int? | Image size in pixels (default: 256, min: 50, max: 1000) (optional) 
            var recovery = high;  // string? | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium (optional) 

            try
            {
                // Generate QR Code (PNG)
                apiInstance.V1TechnologyQrGet(data, size, recovery);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling QrCodeApi.V1TechnologyQrGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyQrGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Generate QR Code (PNG)
    apiInstance.V1TechnologyQrGetWithHttpInfo(data, size, recovery);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling QrCodeApi.V1TechnologyQrGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **data** | **string** | The text or URL to encode in the QR code |  |
| **size** | **int?** | Image size in pixels (default: 256, min: 50, max: 1000) | [optional]  |
| **recovery** | **string?** | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium | [optional]  |

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
| **200** | Successful response |  -  |
| **400** | Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level) |  -  |
| **500** | Failed to generate QR code |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

