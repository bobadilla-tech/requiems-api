# Org.OpenAPITools.Api.ThesaurusApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TextThesaurusWordGet**](ThesaurusApi.md#v1textthesauruswordget) | **GET** /v1/text/thesaurus/{word} | Thesaurus Lookup |

<a id="v1textthesauruswordget"></a>
# **V1TextThesaurusWordGet**
> V1TextThesaurusWordGet200Response V1TextThesaurusWordGet (string word)

Thesaurus Lookup

Returns synonyms and antonyms for the given word.

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
    public class V1TextThesaurusWordGetExample
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
            var apiInstance = new ThesaurusApi(httpClient, config, httpClientHandler);
            var word = happy;  // string | The word to look up in the thesaurus

            try
            {
                // Thesaurus Lookup
                V1TextThesaurusWordGet200Response result = apiInstance.V1TextThesaurusWordGet(word);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ThesaurusApi.V1TextThesaurusWordGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TextThesaurusWordGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Thesaurus Lookup
    ApiResponse<V1TextThesaurusWordGet200Response> response = apiInstance.V1TextThesaurusWordGetWithHttpInfo(word);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ThesaurusApi.V1TextThesaurusWordGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **word** | **string** | The word to look up in the thesaurus |  |

### Return type

[**V1TextThesaurusWordGet200Response**](V1TextThesaurusWordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The word path parameter is missing. |  -  |
| **404** | The word was not found in the thesaurus dataset. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

