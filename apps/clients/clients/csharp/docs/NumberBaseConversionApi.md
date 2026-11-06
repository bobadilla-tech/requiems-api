# Org.OpenAPITools.Api.NumberBaseConversionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyBaseGet**](NumberBaseConversionApi.md#v1technologybaseget) | **GET** /v1/technology/base | Convert Base |

<a id="v1technologybaseget"></a>
# **V1TechnologyBaseGet**
> V1TechnologyBaseGet200Response V1TechnologyBaseGet (int from, int to, string value)

Convert Base

Convert an integer from one number base to another.

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
    public class V1TechnologyBaseGetExample
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
            var apiInstance = new NumberBaseConversionApi(httpClient, config, httpClientHandler);
            var from = 10;  // int | Source base (2, 8, 10, or 16)
            var to = 16;  // int | Target base (2, 8, 10, or 16)
            var value = 255;  // string | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal).

            try
            {
                // Convert Base
                V1TechnologyBaseGet200Response result = apiInstance.V1TechnologyBaseGet(from, to, value);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling NumberBaseConversionApi.V1TechnologyBaseGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyBaseGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Convert Base
    ApiResponse<V1TechnologyBaseGet200Response> response = apiInstance.V1TechnologyBaseGetWithHttpInfo(from, to, value);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling NumberBaseConversionApi.V1TechnologyBaseGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **from** | **int** | Source base (2, 8, 10, or 16) |  |
| **to** | **int** | Target base (2, 8, 10, or 16) |  |
| **value** | **string** | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). |  |

### Return type

[**V1TechnologyBaseGet200Response**](V1TechnologyBaseGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | A required parameter is missing, the base is not one of 2/8/10/16, or value is not valid for the given base. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

