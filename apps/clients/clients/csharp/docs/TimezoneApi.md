# Org.OpenAPITools.Api.TimezoneApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1PlacesTimezoneGet**](TimezoneApi.md#v1placestimezoneget) | **GET** /v1/places/timezone | Get Timezone |

<a id="v1placestimezoneget"></a>
# **V1PlacesTimezoneGet**
> V1PlacesTimezoneGet200Response V1PlacesTimezoneGet (string? lat = null, string? lon = null, string? city = null)

Get Timezone

Returns timezone information for the given coordinates or city name. Provide either `city` or both `lat` and `lon`.

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
    public class V1PlacesTimezoneGetExample
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
            var apiInstance = new TimezoneApi(httpClient, config, httpClientHandler);
            var lat = "lat_example";  // string? | Latitude of the location (-90 to 90). Required when using coordinate-based lookup. (optional) 
            var lon = "lon_example";  // string? | Longitude of the location (-180 to 180). Required when using coordinate-based lookup. (optional) 
            var city = "city_example";  // string? | City name for city-based lookup (e.g. 'Tokyo', 'London'). Required when not using coordinates. (optional) 

            try
            {
                // Get Timezone
                V1PlacesTimezoneGet200Response result = apiInstance.V1PlacesTimezoneGet(lat, lon, city);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling TimezoneApi.V1PlacesTimezoneGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesTimezoneGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Timezone
    ApiResponse<V1PlacesTimezoneGet200Response> response = apiInstance.V1PlacesTimezoneGetWithHttpInfo(lat, lon, city);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling TimezoneApi.V1PlacesTimezoneGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **lat** | **string?** | Latitude of the location (-90 to 90). Required when using coordinate-based lookup. | [optional]  |
| **lon** | **string?** | Longitude of the location (-180 to 180). Required when using coordinate-based lookup. | [optional]  |
| **city** | **string?** | City name for city-based lookup (e.g. &#39;Tokyo&#39;, &#39;London&#39;). Required when not using coordinates. | [optional]  |

### Return type

[**V1PlacesTimezoneGet200Response**](V1PlacesTimezoneGet200Response.md)

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

