# Org.OpenAPITools.Api.DomainInfoApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1NetworkingDomainDomainGet**](DomainInfoApi.md#v1networkingdomaindomainget) | **GET** /v1/networking/domain/{domain} | Get Domain Info |

<a id="v1networkingdomaindomainget"></a>
# **V1NetworkingDomainDomainGet**
> V1NetworkingDomainDomainGet200Response V1NetworkingDomainDomainGet (string domain)

Get Domain Info

Returns DNS records and availability status for the given domain.

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
    public class V1NetworkingDomainDomainGetExample
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
            var apiInstance = new DomainInfoApi(httpClient, config, httpClientHandler);
            var domain = example.com;  // string | The domain to look up (e.g. example.com)

            try
            {
                // Get Domain Info
                V1NetworkingDomainDomainGet200Response result = apiInstance.V1NetworkingDomainDomainGet(domain);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DomainInfoApi.V1NetworkingDomainDomainGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingDomainDomainGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Domain Info
    ApiResponse<V1NetworkingDomainDomainGet200Response> response = apiInstance.V1NetworkingDomainDomainGetWithHttpInfo(domain);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DomainInfoApi.V1NetworkingDomainDomainGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **domain** | **string** | The domain to look up (e.g. example.com) |  |

### Return type

[**V1NetworkingDomainDomainGet200Response**](V1NetworkingDomainDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The domain parameter is not a valid hostname (e.g. missing TLD, invalid characters, or leading/trailing hyphens). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

