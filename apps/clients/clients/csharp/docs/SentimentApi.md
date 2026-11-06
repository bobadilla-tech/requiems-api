# Org.OpenAPITools.Api.SentimentApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextSentimentPost**](SentimentApi.md#v1textsentimentpost) | **POST** /v1/text/sentiment | Analyze Sentiment |

<a id="v1textsentimentpost"></a>
# **V1TextSentimentPost**
> V1TextSentimentPost200Response V1TextSentimentPost (V1TextSentimentPostRequest v1TextSentimentPostRequest)

Analyze Sentiment

Analyzes the sentiment of the provided text and returns a classification, confidence score, and a full breakdown across all three sentiment classes.

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
    public class V1TextSentimentPostExample
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
            var apiInstance = new SentimentApi(httpClient, config, httpClientHandler);
            var v1TextSentimentPostRequest = new V1TextSentimentPostRequest(); // V1TextSentimentPostRequest | 

            try
            {
                // Analyze Sentiment
                V1TextSentimentPost200Response result = apiInstance.V1TextSentimentPost(v1TextSentimentPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling SentimentApi.V1TextSentimentPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextSentimentPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Analyze Sentiment
    ApiResponse<V1TextSentimentPost200Response> response = apiInstance.V1TextSentimentPostWithHttpInfo(v1TextSentimentPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling SentimentApi.V1TextSentimentPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TextSentimentPostRequest** | [**V1TextSentimentPostRequest**](V1TextSentimentPostRequest.md) |  |  |

### Return type

[**V1TextSentimentPost200Response**](V1TextSentimentPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **422** | unprocessable_entity |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

