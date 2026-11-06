# Org.OpenAPITools.Api.GeocodeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1PlacesGeocodeGet**](GeocodeApi.md#v1placesgeocodeget) | **GET** /v1/places/geocode | Geocode Address |
| [**V1PlacesReverseGeocodeGet**](GeocodeApi.md#v1placesreversegeocodeget) | **GET** /v1/places/reverse-geocode | Reverse Geocode |

<a id="v1placesgeocodeget"></a>
# **V1PlacesGeocodeGet**
> V1PlacesGeocodeGet200Response V1PlacesGeocodeGet (string address)

Geocode Address

Converts a free-text address into latitude and longitude coordinates.

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
    public class V1PlacesGeocodeGetExample
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
            var apiInstance = new GeocodeApi(httpClient, config, httpClientHandler);
            var address = 1600 Pennsylvania Ave NW, Washington DC;  // string | The address to geocode (street, city, country, or any combination)

            try
            {
                // Geocode Address
                V1PlacesGeocodeGet200Response result = apiInstance.V1PlacesGeocodeGet(address);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling GeocodeApi.V1PlacesGeocodeGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesGeocodeGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Geocode Address
    ApiResponse<V1PlacesGeocodeGet200Response> response = apiInstance.V1PlacesGeocodeGetWithHttpInfo(address);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling GeocodeApi.V1PlacesGeocodeGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **address** | **string** | The address to geocode (street, city, country, or any combination) |  |

### Return type

[**V1PlacesGeocodeGet200Response**](V1PlacesGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The address parameter is missing. |  -  |
| **404** | No results found for the given address. |  -  |
| **503** | The geocoding service is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1placesreversegeocodeget"></a>
# **V1PlacesReverseGeocodeGet**
> V1PlacesReverseGeocodeGet200Response V1PlacesReverseGeocodeGet (decimal lat, decimal lon)

Reverse Geocode

Converts geographic coordinates into a human-readable address.

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
    public class V1PlacesReverseGeocodeGetExample
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
            var apiInstance = new GeocodeApi(httpClient, config, httpClientHandler);
            var lat = 38.8977;  // decimal | Latitude of the location (-90 to 90)
            var lon = -77.0365;  // decimal | Longitude of the location (-180 to 180)

            try
            {
                // Reverse Geocode
                V1PlacesReverseGeocodeGet200Response result = apiInstance.V1PlacesReverseGeocodeGet(lat, lon);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling GeocodeApi.V1PlacesReverseGeocodeGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesReverseGeocodeGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Reverse Geocode
    ApiResponse<V1PlacesReverseGeocodeGet200Response> response = apiInstance.V1PlacesReverseGeocodeGetWithHttpInfo(lat, lon);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling GeocodeApi.V1PlacesReverseGeocodeGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **lat** | **decimal** | Latitude of the location (-90 to 90) |  |
| **lon** | **decimal** | Longitude of the location (-180 to 180) |  |

### Return type

[**V1PlacesReverseGeocodeGet200Response**](V1PlacesReverseGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | lat or lon is missing or out of range (lat: -90..90, lon: -180..180). |  -  |
| **404** | No address found for the given coordinates. |  -  |
| **503** | The geocoding service is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

