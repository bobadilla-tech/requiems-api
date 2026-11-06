# Org.OpenAPITools.Api.MarkdownApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyMarkdownPost**](MarkdownApi.md#v1technologymarkdownpost) | **POST** /v1/technology/markdown | Convert Markdown to HTML |

<a id="v1technologymarkdownpost"></a>
# **V1TechnologyMarkdownPost**
> V1TechnologyMarkdownPost200Response V1TechnologyMarkdownPost (V1TechnologyMarkdownPostRequest v1TechnologyMarkdownPostRequest)

Convert Markdown to HTML

Converts a Markdown string to HTML. Pass sanitize true to strip potentially unsafe tags like script and iframe from the output.

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
    public class V1TechnologyMarkdownPostExample
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
            var apiInstance = new MarkdownApi(httpClient, config, httpClientHandler);
            var v1TechnologyMarkdownPostRequest = new V1TechnologyMarkdownPostRequest(); // V1TechnologyMarkdownPostRequest | 

            try
            {
                // Convert Markdown to HTML
                V1TechnologyMarkdownPost200Response result = apiInstance.V1TechnologyMarkdownPost(v1TechnologyMarkdownPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling MarkdownApi.V1TechnologyMarkdownPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyMarkdownPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Convert Markdown to HTML
    ApiResponse<V1TechnologyMarkdownPost200Response> response = apiInstance.V1TechnologyMarkdownPostWithHttpInfo(v1TechnologyMarkdownPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling MarkdownApi.V1TechnologyMarkdownPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TechnologyMarkdownPostRequest** | [**V1TechnologyMarkdownPostRequest**](V1TechnologyMarkdownPostRequest.md) |  |  |

### Return type

[**V1TechnologyMarkdownPost200Response**](V1TechnologyMarkdownPost200Response.md)

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

