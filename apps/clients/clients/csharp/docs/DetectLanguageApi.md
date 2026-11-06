# Org.OpenAPITools.Api.DetectLanguageApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextDetectLanguagePost**](DetectLanguageApi.md#v1textdetectlanguagepost) | **POST** /v1/text/detect-language | Detect Language |

<a id="v1textdetectlanguagepost"></a>
# **V1TextDetectLanguagePost**
> V1TextDetectLanguagePost200Response V1TextDetectLanguagePost (V1TextDetectLanguagePostRequest v1TextDetectLanguagePostRequest)

Detect Language

Identifies the language of the provided text and returns the language name, ISO 639-1 code, and confidence score.

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
    public class V1TextDetectLanguagePostExample
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
            var apiInstance = new DetectLanguageApi(httpClient, config, httpClientHandler);
            var v1TextDetectLanguagePostRequest = new V1TextDetectLanguagePostRequest(); // V1TextDetectLanguagePostRequest | 

            try
            {
                // Detect Language
                V1TextDetectLanguagePost200Response result = apiInstance.V1TextDetectLanguagePost(v1TextDetectLanguagePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DetectLanguageApi.V1TextDetectLanguagePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextDetectLanguagePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Detect Language
    ApiResponse<V1TextDetectLanguagePost200Response> response = apiInstance.V1TextDetectLanguagePostWithHttpInfo(v1TextDetectLanguagePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DetectLanguageApi.V1TextDetectLanguagePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TextDetectLanguagePostRequest** | [**V1TextDetectLanguagePostRequest**](V1TextDetectLanguagePostRequest.md) |  |  |

### Return type

[**V1TextDetectLanguagePost200Response**](V1TextDetectLanguagePost200Response.md)

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

