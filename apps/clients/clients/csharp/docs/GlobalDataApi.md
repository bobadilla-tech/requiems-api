# Org.OpenAPITools.Api.GlobalDataApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1SystemsBusinessCalendarCountryGet**](GlobalDataApi.md#v1systemsbusinesscalendarcountryget) | **GET** /v1/systems/business-calendar/{country} | Business Calendar |
| [**V1SystemsLocationResolvePost**](GlobalDataApi.md#v1systemslocationresolvepost) | **POST** /v1/systems/location/resolve | Resolve Location |
| [**V1SystemsTimezoneFromIpIpGet**](GlobalDataApi.md#v1systemstimezonefromipipget) | **GET** /v1/systems/timezone/from-ip/{ip} | Timezone from IP |

<a id="v1systemsbusinesscalendarcountryget"></a>
# **V1SystemsBusinessCalendarCountryGet**
> V1SystemsBusinessCalendarCountryGet200Response V1SystemsBusinessCalendarCountryGet (string country, int? year = null, int? month = null)

Business Calendar

Retrieve working days, public holidays, and next upcoming holiday for any country and optional year/month. Covers 100+ countries.

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
    public class V1SystemsBusinessCalendarCountryGetExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://requiems.xyz";
            // Configure API key authorization: requiems-api-key
            config.AddApiKey("requiems-api-key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("requiems-api-key", "Bearer");

            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new GlobalDataApi(httpClient, config, httpClientHandler);
            var country = DE;  // string | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR)
            var year = 2026;  // int? | Year to retrieve the calendar for. Defaults to the current year. (optional) 
            var month = 6;  // int? | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. (optional) 

            try
            {
                // Business Calendar
                V1SystemsBusinessCalendarCountryGet200Response result = apiInstance.V1SystemsBusinessCalendarCountryGet(country, year, month);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling GlobalDataApi.V1SystemsBusinessCalendarCountryGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsBusinessCalendarCountryGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Business Calendar
    ApiResponse<V1SystemsBusinessCalendarCountryGet200Response> response = apiInstance.V1SystemsBusinessCalendarCountryGetWithHttpInfo(country, year, month);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling GlobalDataApi.V1SystemsBusinessCalendarCountryGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **country** | **string** | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR) |  |
| **year** | **int?** | Year to retrieve the calendar for. Defaults to the current year. | [optional]  |
| **month** | **int?** | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. | [optional]  |

### Return type

[**V1SystemsBusinessCalendarCountryGet200Response**](V1SystemsBusinessCalendarCountryGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | Country is not a 2-letter ISO code, or year/month is out of the valid range |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1systemslocationresolvepost"></a>
# **V1SystemsLocationResolvePost**
> V1SystemsLocationResolvePost200Response V1SystemsLocationResolvePost (V1SystemsLocationResolvePostRequest v1SystemsLocationResolvePostRequest)

Resolve Location

Resolve an address or coordinates into a full location profile — city, country, timezone, UTC offset, current time, working days this month, and next holiday.

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
    public class V1SystemsLocationResolvePostExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://requiems.xyz";
            // Configure API key authorization: requiems-api-key
            config.AddApiKey("requiems-api-key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("requiems-api-key", "Bearer");

            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new GlobalDataApi(httpClient, config, httpClientHandler);
            var v1SystemsLocationResolvePostRequest = new V1SystemsLocationResolvePostRequest(); // V1SystemsLocationResolvePostRequest | 

            try
            {
                // Resolve Location
                V1SystemsLocationResolvePost200Response result = apiInstance.V1SystemsLocationResolvePost(v1SystemsLocationResolvePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling GlobalDataApi.V1SystemsLocationResolvePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsLocationResolvePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Resolve Location
    ApiResponse<V1SystemsLocationResolvePost200Response> response = apiInstance.V1SystemsLocationResolvePostWithHttpInfo(v1SystemsLocationResolvePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling GlobalDataApi.V1SystemsLocationResolvePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsLocationResolvePostRequest** | [**V1SystemsLocationResolvePostRequest**](V1SystemsLocationResolvePostRequest.md) |  |  |

### Return type

[**V1SystemsLocationResolvePost200Response**](V1SystemsLocationResolvePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | Neither address nor coordinates was provided; Geocoding service was unavailable |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1systemstimezonefromipipget"></a>
# **V1SystemsTimezoneFromIpIpGet**
> V1SystemsTimezoneFromIpIpGet200Response V1SystemsTimezoneFromIpIpGet (string ip)

Timezone from IP

Resolve the timezone, UTC offset, and DST status for any IPv4 or IPv6 address without needing an address string.

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
    public class V1SystemsTimezoneFromIpIpGetExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://requiems.xyz";
            // Configure API key authorization: requiems-api-key
            config.AddApiKey("requiems-api-key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("requiems-api-key", "Bearer");

            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new GlobalDataApi(httpClient, config, httpClientHandler);
            var ip = 92.168.1.1;  // string | IPv4 or IPv6 address to resolve

            try
            {
                // Timezone from IP
                V1SystemsTimezoneFromIpIpGet200Response result = apiInstance.V1SystemsTimezoneFromIpIpGet(ip);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling GlobalDataApi.V1SystemsTimezoneFromIpIpGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsTimezoneFromIpIpGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Timezone from IP
    ApiResponse<V1SystemsTimezoneFromIpIpGet200Response> response = apiInstance.V1SystemsTimezoneFromIpIpGetWithHttpInfo(ip);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling GlobalDataApi.V1SystemsTimezoneFromIpIpGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **ip** | **string** | IPv4 or IPv6 address to resolve |  |

### Return type

[**V1SystemsTimezoneFromIpIpGet200Response**](V1SystemsTimezoneFromIpIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | The IP address could not be geolocated |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

