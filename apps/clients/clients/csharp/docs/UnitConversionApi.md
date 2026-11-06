# Org.OpenAPITools.Api.UnitConversionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyConvertGet**](UnitConversionApi.md#v1technologyconvertget) | **GET** /v1/technology/convert | Convert Units |
| [**V1TechnologyConvertUnitsGet**](UnitConversionApi.md#v1technologyconvertunitsget) | **GET** /v1/technology/convert/units | List Available Units |

<a id="v1technologyconvertget"></a>
# **V1TechnologyConvertGet**
> V1TechnologyConvertGet200Response V1TechnologyConvertGet (string from, string to, decimal value)

Convert Units

Convert a value from one unit to another

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
    public class V1TechnologyConvertGetExample
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
            var apiInstance = new UnitConversionApi(httpClient, config, httpClientHandler);
            var from = miles;  // string | Source unit key (e.g. miles, kg, c)
            var to = km;  // string | Target unit key (e.g. km, lb, f)
            var value = 10;  // decimal | Numeric value to convert

            try
            {
                // Convert Units
                V1TechnologyConvertGet200Response result = apiInstance.V1TechnologyConvertGet(from, to, value);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling UnitConversionApi.V1TechnologyConvertGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyConvertGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Convert Units
    ApiResponse<V1TechnologyConvertGet200Response> response = apiInstance.V1TechnologyConvertGetWithHttpInfo(from, to, value);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UnitConversionApi.V1TechnologyConvertGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **from** | **string** | Source unit key (e.g. miles, kg, c) |  |
| **to** | **string** | Target unit key (e.g. km, lb, f) |  |
| **value** | **decimal** | Numeric value to convert |  |

### Return type

[**V1TechnologyConvertGet200Response**](V1TechnologyConvertGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1technologyconvertunitsget"></a>
# **V1TechnologyConvertUnitsGet**
> V1TechnologyConvertUnitsGet200Response V1TechnologyConvertUnitsGet ()

List Available Units

Returns all available unit conversion types grouped by measurement category

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
    public class V1TechnologyConvertUnitsGetExample
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
            var apiInstance = new UnitConversionApi(httpClient, config, httpClientHandler);

            try
            {
                // List Available Units
                V1TechnologyConvertUnitsGet200Response result = apiInstance.V1TechnologyConvertUnitsGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling UnitConversionApi.V1TechnologyConvertUnitsGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyConvertUnitsGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // List Available Units
    ApiResponse<V1TechnologyConvertUnitsGet200Response> response = apiInstance.V1TechnologyConvertUnitsGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UnitConversionApi.V1TechnologyConvertUnitsGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1TechnologyConvertUnitsGet200Response**](V1TechnologyConvertUnitsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

