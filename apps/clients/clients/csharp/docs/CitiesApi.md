# Org.OpenAPITools.Api.CitiesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1PlacesCitiesCityGet**](CitiesApi.md#v1placescitiescityget) | **GET** /v1/places/cities/{city} | Get City Info |

<a id="v1placescitiescityget"></a>
# **V1PlacesCitiesCityGet**
> V1PlacesCitiesCityGet200Response V1PlacesCitiesCityGet (string city)

Get City Info

Returns metadata for a city by name. Lookup is case-insensitive.

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
    public class V1PlacesCitiesCityGetExample
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
            var apiInstance = new CitiesApi(httpClient, config, httpClientHandler);
            var city = london;  // string | City name to look up (e.g. london, tokyo, new york city)

            try
            {
                // Get City Info
                V1PlacesCitiesCityGet200Response result = apiInstance.V1PlacesCitiesCityGet(city);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling CitiesApi.V1PlacesCitiesCityGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesCitiesCityGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get City Info
    ApiResponse<V1PlacesCitiesCityGet200Response> response = apiInstance.V1PlacesCitiesCityGetWithHttpInfo(city);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling CitiesApi.V1PlacesCitiesCityGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **city** | **string** | City name to look up (e.g. london, tokyo, new york city) |  |

### Return type

[**V1PlacesCitiesCityGet200Response**](V1PlacesCitiesCityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **404** | No city with that name was found in the dataset. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

