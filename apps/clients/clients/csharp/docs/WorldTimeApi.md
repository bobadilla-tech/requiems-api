# Org.OpenAPITools.Api.WorldTimeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1PlacesTimeTimezoneGet**](WorldTimeApi.md#v1placestimetimezoneget) | **GET** /v1/places/time/{timezone} | Get Current Time by Timezone |

<a id="v1placestimetimezoneget"></a>
# **V1PlacesTimeTimezoneGet**
> V1PlacesTimeTimezoneGet200Response V1PlacesTimeTimezoneGet (string timezone)

Get Current Time by Timezone

Returns the current time for the given IANA timezone identifier. The timezone is supplied as a path parameter (e.g. `America/New_York`, `Europe/London`, `UTC`).

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
    public class V1PlacesTimeTimezoneGetExample
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
            var apiInstance = new WorldTimeApi(httpClient, config, httpClientHandler);
            var timezone = America/New_York;  // string | IANA timezone identifier (e.g. 'America/New_York', 'Europe/London', 'Asia/Kolkata')

            try
            {
                // Get Current Time by Timezone
                V1PlacesTimeTimezoneGet200Response result = apiInstance.V1PlacesTimeTimezoneGet(timezone);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling WorldTimeApi.V1PlacesTimeTimezoneGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesTimeTimezoneGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Current Time by Timezone
    ApiResponse<V1PlacesTimeTimezoneGet200Response> response = apiInstance.V1PlacesTimeTimezoneGetWithHttpInfo(timezone);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling WorldTimeApi.V1PlacesTimeTimezoneGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **timezone** | **string** | IANA timezone identifier (e.g. &#39;America/New_York&#39;, &#39;Europe/London&#39;, &#39;Asia/Kolkata&#39;) |  |

### Return type

[**V1PlacesTimeTimezoneGet200Response**](V1PlacesTimeTimezoneGet200Response.md)

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

