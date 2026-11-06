# Org.OpenAPITools.Api.DeveloperUtilitiesApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyBase64DecodePost**](DeveloperUtilitiesApi.md#v1technologybase64decodepost) | **POST** /v1/technology/base64/decode | Base64 Decode |
| [**V1TechnologyBase64EncodePost**](DeveloperUtilitiesApi.md#v1technologybase64encodepost) | **POST** /v1/technology/base64/encode | Base64 Encode |

<a id="v1technologybase64decodepost"></a>
# **V1TechnologyBase64DecodePost**
> V1TechnologyBase64DecodePost200Response V1TechnologyBase64DecodePost (V1TechnologyBase64DecodePostRequest v1TechnologyBase64DecodePostRequest)

Base64 Decode

Decode a Base64-encoded string back to plain text. Returns 422 for invalid or unpadded input.

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
    public class V1TechnologyBase64DecodePostExample
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
            var apiInstance = new DeveloperUtilitiesApi(httpClient, config, httpClientHandler);
            var v1TechnologyBase64DecodePostRequest = new V1TechnologyBase64DecodePostRequest(); // V1TechnologyBase64DecodePostRequest | 

            try
            {
                // Base64 Decode
                V1TechnologyBase64DecodePost200Response result = apiInstance.V1TechnologyBase64DecodePost(v1TechnologyBase64DecodePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeveloperUtilitiesApi.V1TechnologyBase64DecodePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyBase64DecodePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Base64 Decode
    ApiResponse<V1TechnologyBase64DecodePost200Response> response = apiInstance.V1TechnologyBase64DecodePostWithHttpInfo(v1TechnologyBase64DecodePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DeveloperUtilitiesApi.V1TechnologyBase64DecodePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TechnologyBase64DecodePostRequest** | [**V1TechnologyBase64DecodePostRequest**](V1TechnologyBase64DecodePostRequest.md) |  |  |

### Return type

[**V1TechnologyBase64DecodePost200Response**](V1TechnologyBase64DecodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | value field is missing or empty |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | The value is not valid Base64 and cannot be decoded |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1technologybase64encodepost"></a>
# **V1TechnologyBase64EncodePost**
> V1TechnologyBase64EncodePost200Response V1TechnologyBase64EncodePost (V1TechnologyBase64EncodePostRequest v1TechnologyBase64EncodePostRequest)

Base64 Encode

Encode a plain-text string to Base64. Supports standard (RFC 4648) and URL-safe (base64url) variants.

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
    public class V1TechnologyBase64EncodePostExample
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
            var apiInstance = new DeveloperUtilitiesApi(httpClient, config, httpClientHandler);
            var v1TechnologyBase64EncodePostRequest = new V1TechnologyBase64EncodePostRequest(); // V1TechnologyBase64EncodePostRequest | 

            try
            {
                // Base64 Encode
                V1TechnologyBase64EncodePost200Response result = apiInstance.V1TechnologyBase64EncodePost(v1TechnologyBase64EncodePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeveloperUtilitiesApi.V1TechnologyBase64EncodePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyBase64EncodePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Base64 Encode
    ApiResponse<V1TechnologyBase64EncodePost200Response> response = apiInstance.V1TechnologyBase64EncodePostWithHttpInfo(v1TechnologyBase64EncodePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DeveloperUtilitiesApi.V1TechnologyBase64EncodePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TechnologyBase64EncodePostRequest** | [**V1TechnologyBase64EncodePostRequest**](V1TechnologyBase64EncodePostRequest.md) |  |  |

### Return type

[**V1TechnologyBase64EncodePost200Response**](V1TechnologyBase64EncodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | value field is missing or empty |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | variant is not standard or url |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

