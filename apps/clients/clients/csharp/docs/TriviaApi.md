# Org.OpenAPITools.Api.TriviaApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1EntertainmentTriviaGet**](TriviaApi.md#v1entertainmenttriviaget) | **GET** /v1/entertainment/trivia | Get Trivia Question |

<a id="v1entertainmenttriviaget"></a>
# **V1EntertainmentTriviaGet**
> V1EntertainmentTriviaGet200Response V1EntertainmentTriviaGet (string? category = null, string? difficulty = null)

Get Trivia Question

Returns a random trivia question with multiple-choice answers. Use the optional category and difficulty query parameters to filter the question pool.

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
    public class V1EntertainmentTriviaGetExample
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
            var apiInstance = new TriviaApi(httpClient, config, httpClientHandler);
            var category = "category_example";  // string? | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. (optional) 
            var difficulty = "difficulty_example";  // string? | Filter by difficulty. One of: easy, medium, hard. (optional) 

            try
            {
                // Get Trivia Question
                V1EntertainmentTriviaGet200Response result = apiInstance.V1EntertainmentTriviaGet(category, difficulty);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling TriviaApi.V1EntertainmentTriviaGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentTriviaGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Trivia Question
    ApiResponse<V1EntertainmentTriviaGet200Response> response = apiInstance.V1EntertainmentTriviaGetWithHttpInfo(category, difficulty);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling TriviaApi.V1EntertainmentTriviaGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **category** | **string?** | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. | [optional]  |
| **difficulty** | **string?** | Filter by difficulty. One of: easy, medium, hard. | [optional]  |

### Return type

[**V1EntertainmentTriviaGet200Response**](V1EntertainmentTriviaGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | An invalid category or difficulty value was provided |  -  |
| **401** | Missing API key |  -  |
| **403** | Invalid or revoked API key |  -  |
| **404** | No questions match the given category and difficulty combination |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

