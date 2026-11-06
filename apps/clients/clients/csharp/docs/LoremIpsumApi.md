# Org.OpenAPITools.Api.LoremIpsumApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextLoremGet**](LoremIpsumApi.md#v1textloremget) | **GET** /v1/text/lorem | Generate Lorem Ipsum |

<a id="v1textloremget"></a>
# **V1TextLoremGet**
> V1TextLoremGet200Response V1TextLoremGet (int? paragraphs = null, int? sentences = null)

Generate Lorem Ipsum

Generate Lorem Ipsum placeholder text with customizable length and format

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
    public class V1TextLoremGetExample
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
            var apiInstance = new LoremIpsumApi(httpClient, config, httpClientHandler);
            var paragraphs = 3;  // int? | Number of paragraphs to generate (1-20) (optional) 
            var sentences = 5;  // int? | Number of sentences per paragraph (1-20) (optional) 

            try
            {
                // Generate Lorem Ipsum
                V1TextLoremGet200Response result = apiInstance.V1TextLoremGet(paragraphs, sentences);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling LoremIpsumApi.V1TextLoremGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextLoremGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Generate Lorem Ipsum
    ApiResponse<V1TextLoremGet200Response> response = apiInstance.V1TextLoremGetWithHttpInfo(paragraphs, sentences);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling LoremIpsumApi.V1TextLoremGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **paragraphs** | **int?** | Number of paragraphs to generate (1-20) | [optional]  |
| **sentences** | **int?** | Number of sentences per paragraph (1-20) | [optional]  |

### Return type

[**V1TextLoremGet200Response**](V1TextLoremGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The paragraphs parameter is out of valid range; The sentences parameter is out of valid range |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

