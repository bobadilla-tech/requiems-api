# Org.OpenAPITools.Api.ProfanityApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1ValidationProfanityBatchPost**](ProfanityApi.md#v1validationprofanitybatchpost) | **POST** /v1/validation/profanity/batch | Batch Check Profanity |
| [**V1ValidationProfanityPost**](ProfanityApi.md#v1validationprofanitypost) | **POST** /v1/validation/profanity | Check Profanity |

<a id="v1validationprofanitybatchpost"></a>
# **V1ValidationProfanityBatchPost**
> V1ValidationProfanityBatchPost200Response V1ValidationProfanityBatchPost (V1ValidationProfanityBatchPostRequest v1ValidationProfanityBatchPostRequest)

Batch Check Profanity

Check up to 50 texts for profanity in a single request. Results are returned in input order.

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
    public class V1ValidationProfanityBatchPostExample
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
            var apiInstance = new ProfanityApi(httpClient, config, httpClientHandler);
            var v1ValidationProfanityBatchPostRequest = new V1ValidationProfanityBatchPostRequest(); // V1ValidationProfanityBatchPostRequest | 

            try
            {
                // Batch Check Profanity
                V1ValidationProfanityBatchPost200Response result = apiInstance.V1ValidationProfanityBatchPost(v1ValidationProfanityBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ProfanityApi.V1ValidationProfanityBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1ValidationProfanityBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch Check Profanity
    ApiResponse<V1ValidationProfanityBatchPost200Response> response = apiInstance.V1ValidationProfanityBatchPostWithHttpInfo(v1ValidationProfanityBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ProfanityApi.V1ValidationProfanityBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1ValidationProfanityBatchPostRequest** | [**V1ValidationProfanityBatchPostRequest**](V1ValidationProfanityBatchPostRequest.md) |  |  |

### Return type

[**V1ValidationProfanityBatchPost200Response**](V1ValidationProfanityBatchPost200Response.md)

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
| **422** | The texts array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1validationprofanitypost"></a>
# **V1ValidationProfanityPost**
> V1ValidationProfanityPost200Response V1ValidationProfanityPost (V1ValidationProfanityPostRequest v1ValidationProfanityPostRequest)

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
    public class V1ValidationProfanityPostExample
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
            var apiInstance = new ProfanityApi(httpClient, config, httpClientHandler);
            var v1ValidationProfanityPostRequest = new V1ValidationProfanityPostRequest(); // V1ValidationProfanityPostRequest | 

            try
            {
                // Check Profanity
                V1ValidationProfanityPost200Response result = apiInstance.V1ValidationProfanityPost(v1ValidationProfanityPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ProfanityApi.V1ValidationProfanityPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1ValidationProfanityPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Check Profanity
    ApiResponse<V1ValidationProfanityPost200Response> response = apiInstance.V1ValidationProfanityPostWithHttpInfo(v1ValidationProfanityPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ProfanityApi.V1ValidationProfanityPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1ValidationProfanityPostRequest** | [**V1ValidationProfanityPostRequest**](V1ValidationProfanityPostRequest.md) |  |  |

### Return type

[**V1ValidationProfanityPost200Response**](V1ValidationProfanityPost200Response.md)

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

