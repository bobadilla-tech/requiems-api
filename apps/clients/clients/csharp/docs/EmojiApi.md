# Org.OpenAPITools.Api.EmojiApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1EntertainmentEmojiNameGet**](EmojiApi.md#v1entertainmentemojinameget) | **GET** /v1/entertainment/emoji/{name} | Get Emoji by Name |
| [**V1EntertainmentEmojiRandomGet**](EmojiApi.md#v1entertainmentemojirandomget) | **GET** /v1/entertainment/emoji/random | Get Random Emoji |
| [**V1EntertainmentEmojiSearchGet**](EmojiApi.md#v1entertainmentemojisearchget) | **GET** /v1/entertainment/emoji/search | Search Emoji |

<a id="v1entertainmentemojinameget"></a>
# **V1EntertainmentEmojiNameGet**
> V1EntertainmentEmojiNameGet200Response V1EntertainmentEmojiNameGet (string name)

Get Emoji by Name

Returns a specific emoji by its CLDR snake_case name. The name is case-insensitive.

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
    public class V1EntertainmentEmojiNameGetExample
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
            var apiInstance = new EmojiApi(httpClient, config, httpClientHandler);
            var name = grinning_face;  // string | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up)

            try
            {
                // Get Emoji by Name
                V1EntertainmentEmojiNameGet200Response result = apiInstance.V1EntertainmentEmojiNameGet(name);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling EmojiApi.V1EntertainmentEmojiNameGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentEmojiNameGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Emoji by Name
    ApiResponse<V1EntertainmentEmojiNameGet200Response> response = apiInstance.V1EntertainmentEmojiNameGetWithHttpInfo(name);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling EmojiApi.V1EntertainmentEmojiNameGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **name** | **string** | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up) |  |

### Return type

[**V1EntertainmentEmojiNameGet200Response**](V1EntertainmentEmojiNameGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **404** | No emoji found with the given name. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1entertainmentemojirandomget"></a>
# **V1EntertainmentEmojiRandomGet**
> V1EntertainmentEmojiRandomGet200Response V1EntertainmentEmojiRandomGet ()

Get Random Emoji

Returns a randomly selected emoji with its full metadata.

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
    public class V1EntertainmentEmojiRandomGetExample
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
            var apiInstance = new EmojiApi(httpClient, config, httpClientHandler);

            try
            {
                // Get Random Emoji
                V1EntertainmentEmojiRandomGet200Response result = apiInstance.V1EntertainmentEmojiRandomGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling EmojiApi.V1EntertainmentEmojiRandomGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentEmojiRandomGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Random Emoji
    ApiResponse<V1EntertainmentEmojiRandomGet200Response> response = apiInstance.V1EntertainmentEmojiRandomGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling EmojiApi.V1EntertainmentEmojiRandomGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1EntertainmentEmojiRandomGet200Response**](V1EntertainmentEmojiRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1entertainmentemojisearchget"></a>
# **V1EntertainmentEmojiSearchGet**
> V1EntertainmentEmojiSearchGet200Response V1EntertainmentEmojiSearchGet (string q)

Search Emoji

Search for emojis whose name or category contains the given query string (case-insensitive). Returns a list of all matches.

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
    public class V1EntertainmentEmojiSearchGetExample
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
            var apiInstance = new EmojiApi(httpClient, config, httpClientHandler);
            var q = happy;  // string | Search term to match against emoji names and categories (e.g. happy, heart, food)

            try
            {
                // Search Emoji
                V1EntertainmentEmojiSearchGet200Response result = apiInstance.V1EntertainmentEmojiSearchGet(q);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling EmojiApi.V1EntertainmentEmojiSearchGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentEmojiSearchGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Search Emoji
    ApiResponse<V1EntertainmentEmojiSearchGet200Response> response = apiInstance.V1EntertainmentEmojiSearchGetWithHttpInfo(q);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling EmojiApi.V1EntertainmentEmojiSearchGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **q** | **string** | Search term to match against emoji names and categories (e.g. happy, heart, food) |  |

### Return type

[**V1EntertainmentEmojiSearchGet200Response**](V1EntertainmentEmojiSearchGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The q query parameter is missing or empty. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

