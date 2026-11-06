# Org.OpenAPITools.Api.TextSimilarityApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextSimilarityPost**](TextSimilarityApi.md#v1textsimilaritypost) | **POST** /v1/text/similarity | Compare Text Similarity |

<a id="v1textsimilaritypost"></a>
# **V1TextSimilarityPost**
> V1TextSimilarityPost200Response V1TextSimilarityPost (V1TextSimilarityPostRequest v1TextSimilarityPostRequest)

Compare Text Similarity

Compares two texts and returns a cosine similarity score.

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
    public class V1TextSimilarityPostExample
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
            var apiInstance = new TextSimilarityApi(httpClient, config, httpClientHandler);
            var v1TextSimilarityPostRequest = new V1TextSimilarityPostRequest(); // V1TextSimilarityPostRequest | 

            try
            {
                // Compare Text Similarity
                V1TextSimilarityPost200Response result = apiInstance.V1TextSimilarityPost(v1TextSimilarityPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling TextSimilarityApi.V1TextSimilarityPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextSimilarityPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Compare Text Similarity
    ApiResponse<V1TextSimilarityPost200Response> response = apiInstance.V1TextSimilarityPostWithHttpInfo(v1TextSimilarityPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling TextSimilarityApi.V1TextSimilarityPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TextSimilarityPostRequest** | [**V1TextSimilarityPostRequest**](V1TextSimilarityPostRequest.md) |  |  |

### Return type

[**V1TextSimilarityPost200Response**](V1TextSimilarityPost200Response.md)

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
| **422** | One or both text fields are missing or empty. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

