# Org.OpenAPITools.Api.EmailNormalizeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextNormalizeBatchPost**](EmailNormalizeApi.md#v1textnormalizebatchpost) | **POST** /v1/text/normalize/batch | Normalize Email Batch |
| [**V1TextNormalizePost**](EmailNormalizeApi.md#v1textnormalizepost) | **POST** /v1/text/normalize | Normalize Email |

<a id="v1textnormalizebatchpost"></a>
# **V1TextNormalizeBatchPost**
> V1TextNormalizeBatchPost200Response V1TextNormalizeBatchPost (V1TextNormalizeBatchPostRequest v1TextNormalizeBatchPostRequest)

Normalize Email Batch

Normalizes up to 100 email addresses in one request. Results are in the same order as the input. Each item includes valid (boolean); when false, only original and message are set. Usage is billed per email processed (see gateway usage headers).

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
    public class V1TextNormalizeBatchPostExample
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
            var apiInstance = new EmailNormalizeApi(httpClient, config, httpClientHandler);
            var v1TextNormalizeBatchPostRequest = new V1TextNormalizeBatchPostRequest(); // V1TextNormalizeBatchPostRequest | 

            try
            {
                // Normalize Email Batch
                V1TextNormalizeBatchPost200Response result = apiInstance.V1TextNormalizeBatchPost(v1TextNormalizeBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling EmailNormalizeApi.V1TextNormalizeBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextNormalizeBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Normalize Email Batch
    ApiResponse<V1TextNormalizeBatchPost200Response> response = apiInstance.V1TextNormalizeBatchPostWithHttpInfo(v1TextNormalizeBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling EmailNormalizeApi.V1TextNormalizeBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TextNormalizeBatchPostRequest** | [**V1TextNormalizeBatchPostRequest**](V1TextNormalizeBatchPostRequest.md) |  |  |

### Return type

[**V1TextNormalizeBatchPost200Response**](V1TextNormalizeBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Invalid JSON or unknown fields in the body |  -  |
| **422** | Missing emails, empty array, too many items, or empty string in the array |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1textnormalizepost"></a>
# **V1TextNormalizePost**
> V1TextNormalizePost200Response V1TextNormalizePost (V1TextNormalizePostRequest v1TextNormalizePostRequest)

Normalize Email

Normalizes a single email address and returns the canonical form together with a breakdown of all transformations applied.

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
    public class V1TextNormalizePostExample
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
            var apiInstance = new EmailNormalizeApi(httpClient, config, httpClientHandler);
            var v1TextNormalizePostRequest = new V1TextNormalizePostRequest(); // V1TextNormalizePostRequest | 

            try
            {
                // Normalize Email
                V1TextNormalizePost200Response result = apiInstance.V1TextNormalizePost(v1TextNormalizePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling EmailNormalizeApi.V1TextNormalizePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextNormalizePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Normalize Email
    ApiResponse<V1TextNormalizePost200Response> response = apiInstance.V1TextNormalizePostWithHttpInfo(v1TextNormalizePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling EmailNormalizeApi.V1TextNormalizePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TextNormalizePostRequest** | [**V1TextNormalizePostRequest**](V1TextNormalizePostRequest.md) |  |  |

### Return type

[**V1TextNormalizePost200Response**](V1TextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The request body is missing, not valid JSON, or contains unknown fields. |  -  |
| **422** | The email field is missing or not a valid email address format. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

