# Org.OpenAPITools.Api.WhoisApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1NetworkingWhoisDomainGet**](WhoisApi.md#v1networkingwhoisdomainget) | **GET** /v1/networking/whois/{domain} | WHOIS Lookup |

<a id="v1networkingwhoisdomainget"></a>
# **V1NetworkingWhoisDomainGet**
> V1NetworkingWhoisDomainGet200Response V1NetworkingWhoisDomainGet (string domain)

WHOIS Lookup

Returns WHOIS registration information for a domain name.

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
    public class V1NetworkingWhoisDomainGetExample
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
            var apiInstance = new WhoisApi(httpClient, config, httpClientHandler);
            var domain = example.com;  // string | The domain name to look up (e.g. example.com)

            try
            {
                // WHOIS Lookup
                V1NetworkingWhoisDomainGet200Response result = apiInstance.V1NetworkingWhoisDomainGet(domain);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling WhoisApi.V1NetworkingWhoisDomainGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingWhoisDomainGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // WHOIS Lookup
    ApiResponse<V1NetworkingWhoisDomainGet200Response> response = apiInstance.V1NetworkingWhoisDomainGetWithHttpInfo(domain);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling WhoisApi.V1NetworkingWhoisDomainGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **domain** | **string** | The domain name to look up (e.g. example.com) |  |

### Return type

[**V1NetworkingWhoisDomainGet200Response**](V1NetworkingWhoisDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The domain name format is invalid. |  -  |
| **404** | No WHOIS record was found for the domain. |  -  |
| **500** | Unexpected server error or upstream WHOIS query failure. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

