# Org.OpenAPITools.Api.QuotesApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1EntertainmentQuotesRandomBatchPost**](QuotesApi.md#v1entertainmentquotesrandombatchpost) | **POST** /v1/entertainment/quotes/random/batch | Get Random Quotes (Batch) |
| [**V1EntertainmentQuotesRandomGet**](QuotesApi.md#v1entertainmentquotesrandomget) | **GET** /v1/entertainment/quotes/random | Get Random Quote |

<a id="v1entertainmentquotesrandombatchpost"></a>
# **V1EntertainmentQuotesRandomBatchPost**
> V1EntertainmentQuotesRandomBatchPost200Response V1EntertainmentQuotesRandomBatchPost (V1EntertainmentQuotesRandomBatchPostRequest v1EntertainmentQuotesRandomBatchPostRequest)

Get Random Quotes (Batch)

Returns up to 50 random quotes in a single request. Each quote counts as one unit of usage.

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
    public class V1EntertainmentQuotesRandomBatchPostExample
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
            var apiInstance = new QuotesApi(httpClient, config, httpClientHandler);
            var v1EntertainmentQuotesRandomBatchPostRequest = new V1EntertainmentQuotesRandomBatchPostRequest(); // V1EntertainmentQuotesRandomBatchPostRequest | 

            try
            {
                // Get Random Quotes (Batch)
                V1EntertainmentQuotesRandomBatchPost200Response result = apiInstance.V1EntertainmentQuotesRandomBatchPost(v1EntertainmentQuotesRandomBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling QuotesApi.V1EntertainmentQuotesRandomBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentQuotesRandomBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Random Quotes (Batch)
    ApiResponse<V1EntertainmentQuotesRandomBatchPost200Response> response = apiInstance.V1EntertainmentQuotesRandomBatchPostWithHttpInfo(v1EntertainmentQuotesRandomBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling QuotesApi.V1EntertainmentQuotesRandomBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1EntertainmentQuotesRandomBatchPostRequest** | [**V1EntertainmentQuotesRandomBatchPostRequest**](V1EntertainmentQuotesRandomBatchPostRequest.md) |  |  |

### Return type

[**V1EntertainmentQuotesRandomBatchPost200Response**](V1EntertainmentQuotesRandomBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Request body is missing or not valid JSON. |  -  |
| **422** | count is missing, less than 1, or greater than 50. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1entertainmentquotesrandomget"></a>
# **V1EntertainmentQuotesRandomGet**
> V1EntertainmentQuotesRandomGet200Response V1EntertainmentQuotesRandomGet ()

Get Random Quote

Returns a random inspirational quote with author attribution

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
    public class V1EntertainmentQuotesRandomGetExample
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
            var apiInstance = new QuotesApi(httpClient, config, httpClientHandler);

            try
            {
                // Get Random Quote
                V1EntertainmentQuotesRandomGet200Response result = apiInstance.V1EntertainmentQuotesRandomGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling QuotesApi.V1EntertainmentQuotesRandomGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentQuotesRandomGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Random Quote
    ApiResponse<V1EntertainmentQuotesRandomGet200Response> response = apiInstance.V1EntertainmentQuotesRandomGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling QuotesApi.V1EntertainmentQuotesRandomGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1EntertainmentQuotesRandomGet200Response**](V1EntertainmentQuotesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **503** | No quotes available in the database |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

