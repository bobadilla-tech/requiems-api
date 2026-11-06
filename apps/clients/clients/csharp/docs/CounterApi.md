# Org.OpenAPITools.Api.CounterApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyCounterNamespaceGet**](CounterApi.md#v1technologycounternamespaceget) | **GET** /v1/technology/counter/{namespace} | Get Counter Value |
| [**V1TechnologyCounterNamespacePost**](CounterApi.md#v1technologycounternamespacepost) | **POST** /v1/technology/counter/{namespace} | Increment Counter |

<a id="v1technologycounternamespaceget"></a>
# **V1TechnologyCounterNamespaceGet**
> V1TechnologyCounterNamespaceGet200Response V1TechnologyCounterNamespaceGet (string varNamespace)

Get Counter Value

Get the current value of a counter without incrementing it

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
    public class V1TechnologyCounterNamespaceGetExample
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
            var apiInstance = new CounterApi(httpClient, config, httpClientHandler);
            var varNamespace = page-views;  // string | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

            try
            {
                // Get Counter Value
                V1TechnologyCounterNamespaceGet200Response result = apiInstance.V1TechnologyCounterNamespaceGet(varNamespace);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling CounterApi.V1TechnologyCounterNamespaceGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyCounterNamespaceGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Counter Value
    ApiResponse<V1TechnologyCounterNamespaceGet200Response> response = apiInstance.V1TechnologyCounterNamespaceGetWithHttpInfo(varNamespace);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling CounterApi.V1TechnologyCounterNamespaceGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **varNamespace** | **string** | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) |  |

### Return type

[**V1TechnologyCounterNamespaceGet200Response**](V1TechnologyCounterNamespaceGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1technologycounternamespacepost"></a>
# **V1TechnologyCounterNamespacePost**
> V1TechnologyCounterNamespacePost200Response V1TechnologyCounterNamespacePost (string varNamespace)

Increment Counter

Atomically increment a counter in the specified namespace and return the new value

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
    public class V1TechnologyCounterNamespacePostExample
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
            var apiInstance = new CounterApi(httpClient, config, httpClientHandler);
            var varNamespace = page-views;  // string | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

            try
            {
                // Increment Counter
                V1TechnologyCounterNamespacePost200Response result = apiInstance.V1TechnologyCounterNamespacePost(varNamespace);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling CounterApi.V1TechnologyCounterNamespacePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyCounterNamespacePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Increment Counter
    ApiResponse<V1TechnologyCounterNamespacePost200Response> response = apiInstance.V1TechnologyCounterNamespacePostWithHttpInfo(varNamespace);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling CounterApi.V1TechnologyCounterNamespacePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **varNamespace** | **string** | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) |  |

### Return type

[**V1TechnologyCounterNamespacePost200Response**](V1TechnologyCounterNamespacePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

