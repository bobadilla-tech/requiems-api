# Org.OpenAPITools.Api.DataIntegrityApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1SystemsContentModeratePost**](DataIntegrityApi.md#v1systemscontentmoderatepost) | **POST** /v1/systems/content/moderate | Moderate Content |
| [**V1SystemsTextNormalizePost**](DataIntegrityApi.md#v1systemstextnormalizepost) | **POST** /v1/systems/text/normalize | Normalize Text |

<a id="v1systemscontentmoderatepost"></a>
# **V1SystemsContentModeratePost**
> V1SystemsContentModeratePost200Response V1SystemsContentModeratePost (V1SystemsContentModeratePostRequest v1SystemsContentModeratePostRequest)

Moderate Content

Check text for profanity, toxicity, sentiment, and language. Returns per-category flags and an is_safe boolean for gating content.

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
    public class V1SystemsContentModeratePostExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://requiems.xyz";
            // Configure API key authorization: requiems-api-key
            config.AddApiKey("requiems-api-key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("requiems-api-key", "Bearer");

            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new DataIntegrityApi(httpClient, config, httpClientHandler);
            var v1SystemsContentModeratePostRequest = new V1SystemsContentModeratePostRequest(); // V1SystemsContentModeratePostRequest | 

            try
            {
                // Moderate Content
                V1SystemsContentModeratePost200Response result = apiInstance.V1SystemsContentModeratePost(v1SystemsContentModeratePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DataIntegrityApi.V1SystemsContentModeratePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsContentModeratePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Moderate Content
    ApiResponse<V1SystemsContentModeratePost200Response> response = apiInstance.V1SystemsContentModeratePostWithHttpInfo(v1SystemsContentModeratePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DataIntegrityApi.V1SystemsContentModeratePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsContentModeratePostRequest** | [**V1SystemsContentModeratePostRequest**](V1SystemsContentModeratePostRequest.md) |  |  |

### Return type

[**V1SystemsContentModeratePost200Response**](V1SystemsContentModeratePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | text field is missing or empty |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1systemstextnormalizepost"></a>
# **V1SystemsTextNormalizePost**
> V1SystemsTextNormalizePost200Response V1SystemsTextNormalizePost (V1SystemsTextNormalizePostRequest v1SystemsTextNormalizePostRequest)

Normalize Text

Apply a composable pipeline of normalization operations to a string. Trim whitespace, fix encoding, normalize unicode, strip HTML, and more — in a single call.

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
    public class V1SystemsTextNormalizePostExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://requiems.xyz";
            // Configure API key authorization: requiems-api-key
            config.AddApiKey("requiems-api-key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("requiems-api-key", "Bearer");

            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new DataIntegrityApi(httpClient, config, httpClientHandler);
            var v1SystemsTextNormalizePostRequest = new V1SystemsTextNormalizePostRequest(); // V1SystemsTextNormalizePostRequest | 

            try
            {
                // Normalize Text
                V1SystemsTextNormalizePost200Response result = apiInstance.V1SystemsTextNormalizePost(v1SystemsTextNormalizePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DataIntegrityApi.V1SystemsTextNormalizePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsTextNormalizePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Normalize Text
    ApiResponse<V1SystemsTextNormalizePost200Response> response = apiInstance.V1SystemsTextNormalizePostWithHttpInfo(v1SystemsTextNormalizePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DataIntegrityApi.V1SystemsTextNormalizePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsTextNormalizePostRequest** | [**V1SystemsTextNormalizePostRequest**](V1SystemsTextNormalizePostRequest.md) |  |  |

### Return type

[**V1SystemsTextNormalizePost200Response**](V1SystemsTextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | text is empty or operations array is empty or missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

