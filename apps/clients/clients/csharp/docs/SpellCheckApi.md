# Org.OpenAPITools.Api.SpellCheckApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextSpellcheckPost**](SpellCheckApi.md#v1textspellcheckpost) | **POST** /v1/text/spellcheck | Check Spelling |

<a id="v1textspellcheckpost"></a>
# **V1TextSpellcheckPost**
> V1TextSpellcheckPost200Response V1TextSpellcheckPost (V1TextSpellcheckPostRequest v1TextSpellcheckPostRequest)

Check Spelling

Checks the input text for spelling mistakes and returns a corrected version along with per-word corrections.

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
    public class V1TextSpellcheckPostExample
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
            var apiInstance = new SpellCheckApi(httpClient, config, httpClientHandler);
            var v1TextSpellcheckPostRequest = new V1TextSpellcheckPostRequest(); // V1TextSpellcheckPostRequest | 

            try
            {
                // Check Spelling
                V1TextSpellcheckPost200Response result = apiInstance.V1TextSpellcheckPost(v1TextSpellcheckPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling SpellCheckApi.V1TextSpellcheckPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextSpellcheckPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Check Spelling
    ApiResponse<V1TextSpellcheckPost200Response> response = apiInstance.V1TextSpellcheckPostWithHttpInfo(v1TextSpellcheckPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling SpellCheckApi.V1TextSpellcheckPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TextSpellcheckPostRequest** | [**V1TextSpellcheckPostRequest**](V1TextSpellcheckPostRequest.md) |  |  |

### Return type

[**V1TextSpellcheckPost200Response**](V1TextSpellcheckPost200Response.md)

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

