# Org.OpenAPITools.Api.HolidaysApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1PlacesHolidaysBatchPost**](HolidaysApi.md#v1placesholidaysbatchpost) | **POST** /v1/places/holidays/batch | Batch Get Holidays |
| [**V1PlacesHolidaysGet**](HolidaysApi.md#v1placesholidaysget) | **GET** /v1/places/holidays | Get Holidays |

<a id="v1placesholidaysbatchpost"></a>
# **V1PlacesHolidaysBatchPost**
> V1PlacesHolidaysBatchPost200Response V1PlacesHolidaysBatchPost (V1PlacesHolidaysBatchPostRequest v1PlacesHolidaysBatchPostRequest)

Batch Get Holidays

Returns holidays for up to 50 (country, year) pairs in a single request. Each pair is processed independently — if one combination has no data, it returns found:false without failing the entire batch.

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
    public class V1PlacesHolidaysBatchPostExample
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
            var apiInstance = new HolidaysApi(httpClient, config, httpClientHandler);
            var v1PlacesHolidaysBatchPostRequest = new V1PlacesHolidaysBatchPostRequest(); // V1PlacesHolidaysBatchPostRequest | 

            try
            {
                // Batch Get Holidays
                V1PlacesHolidaysBatchPost200Response result = apiInstance.V1PlacesHolidaysBatchPost(v1PlacesHolidaysBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling HolidaysApi.V1PlacesHolidaysBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesHolidaysBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch Get Holidays
    ApiResponse<V1PlacesHolidaysBatchPost200Response> response = apiInstance.V1PlacesHolidaysBatchPostWithHttpInfo(v1PlacesHolidaysBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling HolidaysApi.V1PlacesHolidaysBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1PlacesHolidaysBatchPostRequest** | [**V1PlacesHolidaysBatchPostRequest**](V1PlacesHolidaysBatchPostRequest.md) |  |  |

### Return type

[**V1PlacesHolidaysBatchPost200Response**](V1PlacesHolidaysBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Malformed request body |  -  |
| **422** | queries is missing, empty, exceeds 50 items, or contains invalid country codes or years |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1placesholidaysget"></a>
# **V1PlacesHolidaysGet**
> V1PlacesHolidaysGet200Response V1PlacesHolidaysGet (string country, int year)

Get Holidays

Returns a list of public holidays for the specified country and year

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
    public class V1PlacesHolidaysGetExample
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
            var apiInstance = new HolidaysApi(httpClient, config, httpClientHandler);
            var country = US;  // string | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"DE\")
            var year = 2025;  // int | Year for which to retrieve holidays (e.g., 2025)

            try
            {
                // Get Holidays
                V1PlacesHolidaysGet200Response result = apiInstance.V1PlacesHolidaysGet(country, year);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling HolidaysApi.V1PlacesHolidaysGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesHolidaysGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Holidays
    ApiResponse<V1PlacesHolidaysGet200Response> response = apiInstance.V1PlacesHolidaysGetWithHttpInfo(country, year);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling HolidaysApi.V1PlacesHolidaysGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **country** | **string** | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) |  |
| **year** | **int** | Year for which to retrieve holidays (e.g., 2025) |  |

### Return type

[**V1PlacesHolidaysGet200Response**](V1PlacesHolidaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Missing or invalid country code or year parameter |  -  |
| **404** | No holidays found for the specified country and year |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

