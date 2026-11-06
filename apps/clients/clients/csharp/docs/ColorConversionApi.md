# Org.OpenAPITools.Api.ColorConversionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyColorGet**](ColorConversionApi.md#v1technologycolorget) | **GET** /v1/technology/color | Convert Color |

<a id="v1technologycolorget"></a>
# **V1TechnologyColorGet**
> V1TechnologyColorGet200Response V1TechnologyColorGet (string from, string to, string value)

Convert Color

Convert a color value from one format to another. The response always includes all four formats.

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
    public class V1TechnologyColorGetExample
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
            var apiInstance = new ColorConversionApi(httpClient, config, httpClientHandler);
            var from = hex;  // string | Source color format: hex, rgb, hsl, or cmyk
            var to = hsl;  // string | Target color format: hex, rgb, hsl, or cmyk
            var value = #ff5733;  // string | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%))

            try
            {
                // Convert Color
                V1TechnologyColorGet200Response result = apiInstance.V1TechnologyColorGet(from, to, value);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ColorConversionApi.V1TechnologyColorGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyColorGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Convert Color
    ApiResponse<V1TechnologyColorGet200Response> response = apiInstance.V1TechnologyColorGetWithHttpInfo(from, to, value);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ColorConversionApi.V1TechnologyColorGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **from** | **string** | Source color format: hex, rgb, hsl, or cmyk |  |
| **to** | **string** | Target color format: hex, rgb, hsl, or cmyk |  |
| **value** | **string** | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%)) |  |

### Return type

[**V1TechnologyColorGet200Response**](V1TechnologyColorGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | One or more of from, to, or value parameters is missing or the from/to value is not one of: hex, rgb, hsl, cmyk. |  -  |
| **422** | The value cannot be parsed in the specified from format. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

