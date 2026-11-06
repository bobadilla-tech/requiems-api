# Org.OpenAPITools.Api.WorkingDaysApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1PlacesWorkingDaysGet**](WorkingDaysApi.md#v1placesworkingdaysget) | **GET** /v1/places/working-days | Calculate Working Days |

<a id="v1placesworkingdaysget"></a>
# **V1PlacesWorkingDaysGet**
> V1PlacesWorkingDaysGet200Response V1PlacesWorkingDaysGet (string from, string to, string? country = null, string? subdivision = null)

Calculate Working Days

Calculate the number of working days between two dates, optionally accounting for country-specific holidays

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
    public class V1PlacesWorkingDaysGetExample
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
            var apiInstance = new WorkingDaysApi(httpClient, config, httpClientHandler);
            var from = 2024-02-23;  // string | Start date in YYYY-MM-DD format (ISO 8601)
            var to = 2024-02-28;  // string | End date in YYYY-MM-DD format (ISO 8601). Must be >= from date.
            var country = US;  // string? | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"FR\"). When provided, country-specific holidays are excluded from working days count. (optional) 
            var subdivision = NY;  // string? | ISO 3166-2 subdivision code for state/region within the country (e.g., \"NY\" for New York, \"CA\" for California). Only used when country is provided. (optional) 

            try
            {
                // Calculate Working Days
                V1PlacesWorkingDaysGet200Response result = apiInstance.V1PlacesWorkingDaysGet(from, to, country, subdivision);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling WorkingDaysApi.V1PlacesWorkingDaysGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesWorkingDaysGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Calculate Working Days
    ApiResponse<V1PlacesWorkingDaysGet200Response> response = apiInstance.V1PlacesWorkingDaysGetWithHttpInfo(from, to, country, subdivision);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling WorkingDaysApi.V1PlacesWorkingDaysGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **from** | **string** | Start date in YYYY-MM-DD format (ISO 8601) |  |
| **to** | **string** | End date in YYYY-MM-DD format (ISO 8601). Must be &gt;&#x3D; from date. |  |
| **country** | **string?** | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;FR\&quot;). When provided, country-specific holidays are excluded from working days count. | [optional]  |
| **subdivision** | **string?** | ISO 3166-2 subdivision code for state/region within the country (e.g., \&quot;NY\&quot; for New York, \&quot;CA\&quot; for California). Only used when country is provided. | [optional]  |

### Return type

[**V1PlacesWorkingDaysGet200Response**](V1PlacesWorkingDaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The from and to parameters are required, or to date is before from date, or invalid date format |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

