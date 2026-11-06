# Org.OpenAPITools.Api.IpInfoApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1NetworkingIpGet**](IpInfoApi.md#v1networkingipget) | **GET** /v1/networking/ip | Get IP Info (Caller IP) |
| [**V1NetworkingIpInfoBatchPost**](IpInfoApi.md#v1networkingipinfobatchpost) | **POST** /v1/networking/ip/info/batch | Batch IP Geolocation |
| [**V1NetworkingIpIpGet**](IpInfoApi.md#v1networkingipipget) | **GET** /v1/networking/ip/{ip} | Get IP Info for IP |

<a id="v1networkingipget"></a>
# **V1NetworkingIpGet**
> V1NetworkingIpGet200Response V1NetworkingIpGet ()

Get IP Info (Caller IP)

Get geolocation and network information for the requesting client's IP address. Useful when you want information about the user making the request without specifying an IP explicitly.

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
    public class V1NetworkingIpGetExample
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
            var apiInstance = new IpInfoApi(httpClient, config, httpClientHandler);

            try
            {
                // Get IP Info (Caller IP)
                V1NetworkingIpGet200Response result = apiInstance.V1NetworkingIpGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IpInfoApi.V1NetworkingIpGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingIpGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get IP Info (Caller IP)
    ApiResponse<V1NetworkingIpGet200Response> response = apiInstance.V1NetworkingIpGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IpInfoApi.V1NetworkingIpGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1NetworkingIpGet200Response**](V1NetworkingIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1networkingipinfobatchpost"></a>
# **V1NetworkingIpInfoBatchPost**
> V1NetworkingIpInfoBatchPost200Response V1NetworkingIpInfoBatchPost (V1NetworkingIpAsnBatchPostRequest v1NetworkingIpAsnBatchPostRequest)

Batch IP Geolocation

Look up geolocation data for up to 50 IP addresses in a single request. Results are returned in input order. Per-item errors are reported inline.

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
    public class V1NetworkingIpInfoBatchPostExample
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
            var apiInstance = new IpInfoApi(httpClient, config, httpClientHandler);
            var v1NetworkingIpAsnBatchPostRequest = new V1NetworkingIpAsnBatchPostRequest(); // V1NetworkingIpAsnBatchPostRequest | 

            try
            {
                // Batch IP Geolocation
                V1NetworkingIpInfoBatchPost200Response result = apiInstance.V1NetworkingIpInfoBatchPost(v1NetworkingIpAsnBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IpInfoApi.V1NetworkingIpInfoBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingIpInfoBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch IP Geolocation
    ApiResponse<V1NetworkingIpInfoBatchPost200Response> response = apiInstance.V1NetworkingIpInfoBatchPostWithHttpInfo(v1NetworkingIpAsnBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IpInfoApi.V1NetworkingIpInfoBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1NetworkingIpAsnBatchPostRequest** | [**V1NetworkingIpAsnBatchPostRequest**](V1NetworkingIpAsnBatchPostRequest.md) |  |  |

### Return type

[**V1NetworkingIpInfoBatchPost200Response**](V1NetworkingIpInfoBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The request body is missing or malformed. |  -  |
| **422** | The ips array is missing, empty, exceeds 50 items, or contains an invalid IP address. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1networkingipipget"></a>
# **V1NetworkingIpIpGet**
> V1NetworkingIpIpGet200Response V1NetworkingIpIpGet (string ip)

Get IP Info for IP

Get geolocation and network information for a specific IP address.

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
    public class V1NetworkingIpIpGetExample
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
            var apiInstance = new IpInfoApi(httpClient, config, httpClientHandler);
            var ip = 8.8.8.8;  // string | The IP address to look up (supports IPv4 and IPv6)

            try
            {
                // Get IP Info for IP
                V1NetworkingIpIpGet200Response result = apiInstance.V1NetworkingIpIpGet(ip);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IpInfoApi.V1NetworkingIpIpGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingIpIpGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get IP Info for IP
    ApiResponse<V1NetworkingIpIpGet200Response> response = apiInstance.V1NetworkingIpIpGetWithHttpInfo(ip);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IpInfoApi.V1NetworkingIpIpGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **ip** | **string** | The IP address to look up (supports IPv4 and IPv6) |  |

### Return type

[**V1NetworkingIpIpGet200Response**](V1NetworkingIpIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The IP address is invalid |  -  |
| **500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

