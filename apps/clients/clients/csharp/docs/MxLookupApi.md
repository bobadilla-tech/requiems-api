# Org.OpenAPITools.Api.MxLookupApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1NetworkingMxDomainGet**](MxLookupApi.md#v1networkingmxdomainget) | **GET** /v1/networking/mx/{domain} | MX Lookup |

<a id="v1networkingmxdomainget"></a>
# **V1NetworkingMxDomainGet**
> V1NetworkingMxDomainGet200Response V1NetworkingMxDomainGet (string domain)

MX Lookup

Retrieve all MX records for a domain. Results are sorted by priority ascending (lowest numeric value has highest mail delivery priority per RFC 5321).

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
    public class V1NetworkingMxDomainGetExample
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
            var apiInstance = new MxLookupApi(httpClient, config, httpClientHandler);
            var domain = gmail.com;  // string | The domain name to look up MX records for (e.g. gmail.com)

            try
            {
                // MX Lookup
                V1NetworkingMxDomainGet200Response result = apiInstance.V1NetworkingMxDomainGet(domain);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling MxLookupApi.V1NetworkingMxDomainGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingMxDomainGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // MX Lookup
    ApiResponse<V1NetworkingMxDomainGet200Response> response = apiInstance.V1NetworkingMxDomainGetWithHttpInfo(domain);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling MxLookupApi.V1NetworkingMxDomainGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **domain** | **string** | The domain name to look up MX records for (e.g. gmail.com) |  |

### Return type

[**V1NetworkingMxDomainGet200Response**](V1NetworkingMxDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The domain parameter is not a valid domain name. |  -  |
| **404** | No MX records were found for the domain (domain may not accept email). |  -  |
| **500** | DNS lookup failed due to an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

