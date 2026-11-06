# Org.OpenAPITools.Api.DataFormatConversionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyFormatPost**](DataFormatConversionApi.md#v1technologyformatpost) | **POST** /v1/technology/format | Convert Format |

<a id="v1technologyformatpost"></a>
# **V1TechnologyFormatPost**
> V1TechnologyFormatPost200Response V1TechnologyFormatPost (V1TechnologyFormatPostRequest v1TechnologyFormatPostRequest)

Convert Format

Convert content from one structured data format to another. Supported formats: json, yaml, csv, xml, toml.

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
    public class V1TechnologyFormatPostExample
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
            var apiInstance = new DataFormatConversionApi(httpClient, config, httpClientHandler);
            var v1TechnologyFormatPostRequest = new V1TechnologyFormatPostRequest(); // V1TechnologyFormatPostRequest | 

            try
            {
                // Convert Format
                V1TechnologyFormatPost200Response result = apiInstance.V1TechnologyFormatPost(v1TechnologyFormatPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DataFormatConversionApi.V1TechnologyFormatPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyFormatPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Convert Format
    ApiResponse<V1TechnologyFormatPost200Response> response = apiInstance.V1TechnologyFormatPostWithHttpInfo(v1TechnologyFormatPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DataFormatConversionApi.V1TechnologyFormatPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TechnologyFormatPostRequest** | [**V1TechnologyFormatPostRequest**](V1TechnologyFormatPostRequest.md) |  |  |

### Return type

[**V1TechnologyFormatPost200Response**](V1TechnologyFormatPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **413** | The content field exceeds the 512 KB limit. |  -  |
| **422** | One of from, to, or content is missing, or from/to is not one of the supported format values.; The content field is not valid JSON (when from is json).; The content field is not valid YAML (when from is yaml).; The content field is not valid CSV, or a row has more columns than the header (when from is csv).; The content field is not valid XML (when from is xml).; The content field is not valid TOML (when from is toml).; The data structure is incompatible with the target format (e.g. converting a JSON array to TOML, which requires a top-level object). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

