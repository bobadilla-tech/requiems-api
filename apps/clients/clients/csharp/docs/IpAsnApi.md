# Org.OpenAPITools.Api.IpAsnApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1NetworkingIpAsnGet**](IpAsnApi.md#v1networkingipasnget) | **GET** /v1/networking/ip/asn | Lookup ASN (Caller IP) |
| [**V1NetworkingIpAsnIpGet**](IpAsnApi.md#v1networkingipasnipget) | **GET** /v1/networking/ip/asn/{ip} | Lookup ASN for IP |

<a id="v1networkingipasnget"></a>
# **V1NetworkingIpAsnGet**
> V1NetworkingIpAsnGet200Response V1NetworkingIpAsnGet ()

Lookup ASN (Caller IP)

Look up ASN, organization, ISP, and network details for the requesting client's IP address. Useful when you want information about the user making the request without specifying an IP explicitly.

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
    public class V1NetworkingIpAsnGetExample
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
            var apiInstance = new IpAsnApi(httpClient, config, httpClientHandler);

            try
            {
                // Lookup ASN (Caller IP)
                V1NetworkingIpAsnGet200Response result = apiInstance.V1NetworkingIpAsnGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IpAsnApi.V1NetworkingIpAsnGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingIpAsnGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Lookup ASN (Caller IP)
    ApiResponse<V1NetworkingIpAsnGet200Response> response = apiInstance.V1NetworkingIpAsnGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IpAsnApi.V1NetworkingIpAsnGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1NetworkingIpAsnGet200Response**](V1NetworkingIpAsnGet200Response.md)

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

<a id="v1networkingipasnipget"></a>
# **V1NetworkingIpAsnIpGet**
> V1NetworkingIpAsnIpGet200Response V1NetworkingIpAsnIpGet (string ip)

Lookup ASN for IP

Look up ASN, organization, ISP, and network details for a specific IP address.

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
    public class V1NetworkingIpAsnIpGetExample
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
            var apiInstance = new IpAsnApi(httpClient, config, httpClientHandler);
            var ip = 8.8.8.8;  // string | The IP address to look up (supports IPv4 and IPv6)

            try
            {
                // Lookup ASN for IP
                V1NetworkingIpAsnIpGet200Response result = apiInstance.V1NetworkingIpAsnIpGet(ip);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IpAsnApi.V1NetworkingIpAsnIpGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingIpAsnIpGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Lookup ASN for IP
    ApiResponse<V1NetworkingIpAsnIpGet200Response> response = apiInstance.V1NetworkingIpAsnIpGetWithHttpInfo(ip);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IpAsnApi.V1NetworkingIpAsnIpGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **ip** | **string** | The IP address to look up (supports IPv4 and IPv6) |  |

### Return type

[**V1NetworkingIpAsnIpGet200Response**](V1NetworkingIpAsnIpGet200Response.md)

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

