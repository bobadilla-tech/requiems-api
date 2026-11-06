# Org.OpenAPITools.Api.RandomUserApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyRandomUserBatchPost**](RandomUserApi.md#v1technologyrandomuserbatchpost) | **POST** /v1/technology/random-user/batch | Batch Generate Users |
| [**V1TechnologyRandomUserGet**](RandomUserApi.md#v1technologyrandomuserget) | **GET** /v1/technology/random-user | Get Random User |

<a id="v1technologyrandomuserbatchpost"></a>
# **V1TechnologyRandomUserBatchPost**
> V1TechnologyRandomUserBatchPost200Response V1TechnologyRandomUserBatchPost (V1TechnologyRandomUserBatchPostRequest v1TechnologyRandomUserBatchPostRequest)

Batch Generate Users

Returns multiple randomly generated fake user profiles in a single request. Each call consumes count units of quota. Maximum 50 users per request.

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
    public class V1TechnologyRandomUserBatchPostExample
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
            var apiInstance = new RandomUserApi(httpClient, config, httpClientHandler);
            var v1TechnologyRandomUserBatchPostRequest = new V1TechnologyRandomUserBatchPostRequest(); // V1TechnologyRandomUserBatchPostRequest | 

            try
            {
                // Batch Generate Users
                V1TechnologyRandomUserBatchPost200Response result = apiInstance.V1TechnologyRandomUserBatchPost(v1TechnologyRandomUserBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling RandomUserApi.V1TechnologyRandomUserBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyRandomUserBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch Generate Users
    ApiResponse<V1TechnologyRandomUserBatchPost200Response> response = apiInstance.V1TechnologyRandomUserBatchPostWithHttpInfo(v1TechnologyRandomUserBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling RandomUserApi.V1TechnologyRandomUserBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1TechnologyRandomUserBatchPostRequest** | [**V1TechnologyRandomUserBatchPostRequest**](V1TechnologyRandomUserBatchPostRequest.md) |  |  |

### Return type

[**V1TechnologyRandomUserBatchPost200Response**](V1TechnologyRandomUserBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Request body is malformed or not valid JSON |  -  |
| **422** | count is missing, less than 1, or greater than 50 |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1technologyrandomuserget"></a>
# **V1TechnologyRandomUserGet**
> V1TechnologyRandomUserGet200Response V1TechnologyRandomUserGet ()

Get Random User

Returns a randomly generated fake user profile.

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
    public class V1TechnologyRandomUserGetExample
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
            var apiInstance = new RandomUserApi(httpClient, config, httpClientHandler);

            try
            {
                // Get Random User
                V1TechnologyRandomUserGet200Response result = apiInstance.V1TechnologyRandomUserGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling RandomUserApi.V1TechnologyRandomUserGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyRandomUserGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Random User
    ApiResponse<V1TechnologyRandomUserGet200Response> response = apiInstance.V1TechnologyRandomUserGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling RandomUserApi.V1TechnologyRandomUserGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1TechnologyRandomUserGet200Response**](V1TechnologyRandomUserGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

