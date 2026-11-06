# Org.OpenAPITools.Api.ChuckNorrisApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1EntertainmentChuckNorrisGet**](ChuckNorrisApi.md#v1entertainmentchucknorrisget) | **GET** /v1/entertainment/chuck-norris | Get Random Chuck Norris Fact |

<a id="v1entertainmentchucknorrisget"></a>
# **V1EntertainmentChuckNorrisGet**
> V1EntertainmentChuckNorrisGet200Response V1EntertainmentChuckNorrisGet ()

Get Random Chuck Norris Fact

Returns a randomly selected Chuck Norris fact from the built-in database.

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
    public class V1EntertainmentChuckNorrisGetExample
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
            var apiInstance = new ChuckNorrisApi(httpClient, config, httpClientHandler);

            try
            {
                // Get Random Chuck Norris Fact
                V1EntertainmentChuckNorrisGet200Response result = apiInstance.V1EntertainmentChuckNorrisGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ChuckNorrisApi.V1EntertainmentChuckNorrisGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentChuckNorrisGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Random Chuck Norris Fact
    ApiResponse<V1EntertainmentChuckNorrisGet200Response> response = apiInstance.V1EntertainmentChuckNorrisGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ChuckNorrisApi.V1EntertainmentChuckNorrisGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1EntertainmentChuckNorrisGet200Response**](V1EntertainmentChuckNorrisGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

