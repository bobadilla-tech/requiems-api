# Org.OpenAPITools.Api.ProfanityApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextProfanityPost**](ProfanityApi.md#v1textprofanitypost) | **POST** /v1/text/profanity | Check Profanity |

<a id="v1textprofanitypost"></a>
# **V1TextProfanityPost**
> V1TextProfanityPost200Response V1TextProfanityPost (V1TextProfanityPostRequest v1TextProfanityPostRequest)

Check Profanity

Checks text for profanity, returning a censored version and the list of flagged words.

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
    public class V1TextProfanityPostExample
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
            var apiInstance = new ProfanityApi(httpClient, config, httpClientHandler);
            var v1TextProfanityPostRequest = new V1TextProfanityPostRequest(); // V1TextProfanityPostRequest | 

            try
            {
                // Check Profanity
                V1TextProfanityPost200Response result = apiInstance.V1TextProfanityPost(v1TextProfanityPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ProfanityApi.V1TextProfanityPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextProfanityPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Check Profanity
    ApiResponse<V1TextProfanityPost200Response> response = apiInstance.V1TextProfanityPostWithHttpInfo(v1TextProfanityPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ProfanityApi.V1TextProfanityPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TextProfanityPostRequest** | [**V1TextProfanityPostRequest**](V1TextProfanityPostRequest.md) |  |  |

### Return type

[**V1TextProfanityPost200Response**](V1TextProfanityPost200Response.md)

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
| **422** | The text field is missing or empty. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

