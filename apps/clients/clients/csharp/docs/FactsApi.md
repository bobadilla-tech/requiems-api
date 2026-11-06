# Org.OpenAPITools.Api.FactsApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1EntertainmentFactsGet**](FactsApi.md#v1entertainmentfactsget) | **GET** /v1/entertainment/facts | Get Random Fact |

<a id="v1entertainmentfactsget"></a>
# **V1EntertainmentFactsGet**
> V1EntertainmentFactsGet200Response V1EntertainmentFactsGet (string? category = null)

Get Random Fact

Returns a randomly selected fact, optionally filtered by category.

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
    public class V1EntertainmentFactsGetExample
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
            var apiInstance = new FactsApi(httpClient, config, httpClientHandler);
            var category = science;  // string? | Filter by category. Valid values: science, history, technology, nature, space, food (optional) 

            try
            {
                // Get Random Fact
                V1EntertainmentFactsGet200Response result = apiInstance.V1EntertainmentFactsGet(category);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling FactsApi.V1EntertainmentFactsGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentFactsGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Random Fact
    ApiResponse<V1EntertainmentFactsGet200Response> response = apiInstance.V1EntertainmentFactsGetWithHttpInfo(category);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling FactsApi.V1EntertainmentFactsGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **category** | **string?** | Filter by category. Valid values: science, history, technology, nature, space, food | [optional]  |

### Return type

[**V1EntertainmentFactsGet200Response**](V1EntertainmentFactsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | bad_request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

