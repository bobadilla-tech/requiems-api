# Org.OpenAPITools.Api.VpnDetectionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1NetworkingIpVpnIpGet**](VpnDetectionApi.md#v1networkingipvpnipget) | **GET** /v1/networking/ip/vpn/{ip} | Check IP Address |

<a id="v1networkingipvpnipget"></a>
# **V1NetworkingIpVpnIpGet**
> V1NetworkingIpVpnIpGet200Response V1NetworkingIpVpnIpGet (string ip)

Check IP Address

Analyze an IP address to determine if it's a VPN, proxy, Tor exit node, or hosting provider. Returns detailed threat indicators and scores.

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
    public class V1NetworkingIpVpnIpGetExample
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
            var apiInstance = new VpnDetectionApi(httpClient, config, httpClientHandler);
            var ip = 8.8.8.8;  // string | The IP address to check (supports IPv4 and IPv6)

            try
            {
                // Check IP Address
                V1NetworkingIpVpnIpGet200Response result = apiInstance.V1NetworkingIpVpnIpGet(ip);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling VpnDetectionApi.V1NetworkingIpVpnIpGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingIpVpnIpGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Check IP Address
    ApiResponse<V1NetworkingIpVpnIpGet200Response> response = apiInstance.V1NetworkingIpVpnIpGetWithHttpInfo(ip);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling VpnDetectionApi.V1NetworkingIpVpnIpGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **ip** | **string** | The IP address to check (supports IPv4 and IPv6) |  |

### Return type

[**V1NetworkingIpVpnIpGet200Response**](V1NetworkingIpVpnIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The IP address is missing or invalid |  -  |
| **500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

