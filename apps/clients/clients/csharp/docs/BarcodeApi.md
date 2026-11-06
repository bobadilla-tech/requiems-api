# Org.OpenAPITools.Api.BarcodeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyBarcodeBase64Get**](BarcodeApi.md#v1technologybarcodebase64get) | **GET** /v1/technology/barcode/base64 | Generate Barcode (Base64 JSON) |
| [**V1TechnologyBarcodeGet**](BarcodeApi.md#v1technologybarcodeget) | **GET** /v1/technology/barcode | Generate Barcode (PNG) |

<a id="v1technologybarcodebase64get"></a>
# **V1TechnologyBarcodeBase64Get**
> V1TechnologyBarcodeBase64Get200Response V1TechnologyBarcodeBase64Get (string data, string type)

Generate Barcode (Base64 JSON)

Returns a JSON envelope containing the barcode as a base64-encoded PNG string, along with its type and dimensions.

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
    public class V1TechnologyBarcodeBase64GetExample
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
            var apiInstance = new BarcodeApi(httpClient, config, httpClientHandler);
            var data = 123456789;  // string | The text or numeric string to encode in the barcode
            var type = code128;  // string | Barcode format: code128, code93, code39, ean8, ean13

            try
            {
                // Generate Barcode (Base64 JSON)
                V1TechnologyBarcodeBase64Get200Response result = apiInstance.V1TechnologyBarcodeBase64Get(data, type);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling BarcodeApi.V1TechnologyBarcodeBase64Get: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyBarcodeBase64GetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Generate Barcode (Base64 JSON)
    ApiResponse<V1TechnologyBarcodeBase64Get200Response> response = apiInstance.V1TechnologyBarcodeBase64GetWithHttpInfo(data, type);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling BarcodeApi.V1TechnologyBarcodeBase64GetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **data** | **string** | The text or numeric string to encode in the barcode |  |
| **type** | **string** | Barcode format: code128, code93, code39, ean8, ean13 |  |

### Return type

[**V1TechnologyBarcodeBase64Get200Response**](V1TechnologyBarcodeBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Missing or invalid parameters (e.g. data not provided, unsupported type) |  -  |
| **422** | Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1technologybarcodeget"></a>
# **V1TechnologyBarcodeGet**
> void V1TechnologyBarcodeGet (string data, string type)

Generate Barcode (PNG)

Returns a raw PNG image of the barcode. Ideal for direct embedding or file download.

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
    public class V1TechnologyBarcodeGetExample
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
            var apiInstance = new BarcodeApi(httpClient, config, httpClientHandler);
            var data = 123456789;  // string | The text or numeric string to encode in the barcode
            var type = code128;  // string | Barcode format: code128, code93, code39, ean8, ean13

            try
            {
                // Generate Barcode (PNG)
                apiInstance.V1TechnologyBarcodeGet(data, type);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling BarcodeApi.V1TechnologyBarcodeGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyBarcodeGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Generate Barcode (PNG)
    apiInstance.V1TechnologyBarcodeGetWithHttpInfo(data, type);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling BarcodeApi.V1TechnologyBarcodeGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **data** | **string** | The text or numeric string to encode in the barcode |  |
| **type** | **string** | Barcode format: code128, code93, code39, ean8, ean13 |  |

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
| **400** | Missing or invalid parameters (e.g. data not provided, unsupported type) |  -  |
| **422** | Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

