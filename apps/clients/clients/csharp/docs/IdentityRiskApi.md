# Org.OpenAPITools.Api.IdentityRiskApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1SystemsRiskScorePost**](IdentityRiskApi.md#v1systemsriskscorepost) | **POST** /v1/systems/risk/score | Score Risk |
| [**V1SystemsSignupProtectPost**](IdentityRiskApi.md#v1systemssignupprotectpost) | **POST** /v1/systems/signup/protect | Protect Signup |
| [**V1SystemsUserVerifyPost**](IdentityRiskApi.md#v1systemsuserverifypost) | **POST** /v1/systems/user/verify | Verify User |

<a id="v1systemsriskscorepost"></a>
# **V1SystemsRiskScorePost**
> V1SystemsRiskScorePost200Response V1SystemsRiskScorePost (V1SystemsRiskScorePostRequest v1SystemsRiskScorePostRequest)

Score Risk

Score a user for risk without the full signal breakdown. Lower latency than /signup/protect, suited for background re-scoring and rate limiting.

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
    public class V1SystemsRiskScorePostExample
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
            var apiInstance = new IdentityRiskApi(httpClient, config, httpClientHandler);
            var v1SystemsRiskScorePostRequest = new V1SystemsRiskScorePostRequest(); // V1SystemsRiskScorePostRequest | 

            try
            {
                // Score Risk
                V1SystemsRiskScorePost200Response result = apiInstance.V1SystemsRiskScorePost(v1SystemsRiskScorePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IdentityRiskApi.V1SystemsRiskScorePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsRiskScorePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Score Risk
    ApiResponse<V1SystemsRiskScorePost200Response> response = apiInstance.V1SystemsRiskScorePostWithHttpInfo(v1SystemsRiskScorePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IdentityRiskApi.V1SystemsRiskScorePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsRiskScorePostRequest** | [**V1SystemsRiskScorePostRequest**](V1SystemsRiskScorePostRequest.md) |  |  |

### Return type

[**V1SystemsRiskScorePost200Response**](V1SystemsRiskScorePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | All of email, phone, and ip_address were omitted |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1systemssignupprotectpost"></a>
# **V1SystemsSignupProtectPost**
> V1SystemsSignupProtectPost200Response V1SystemsSignupProtectPost (V1SystemsSignupProtectPostRequest v1SystemsSignupProtectPostRequest)

Protect Signup

Evaluate a new user at signup. Returns a full risk decision with per-signal breakdown across email, phone, and IP.

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
    public class V1SystemsSignupProtectPostExample
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
            var apiInstance = new IdentityRiskApi(httpClient, config, httpClientHandler);
            var v1SystemsSignupProtectPostRequest = new V1SystemsSignupProtectPostRequest(); // V1SystemsSignupProtectPostRequest | 

            try
            {
                // Protect Signup
                V1SystemsSignupProtectPost200Response result = apiInstance.V1SystemsSignupProtectPost(v1SystemsSignupProtectPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IdentityRiskApi.V1SystemsSignupProtectPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsSignupProtectPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Protect Signup
    ApiResponse<V1SystemsSignupProtectPost200Response> response = apiInstance.V1SystemsSignupProtectPostWithHttpInfo(v1SystemsSignupProtectPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IdentityRiskApi.V1SystemsSignupProtectPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsSignupProtectPostRequest** | [**V1SystemsSignupProtectPostRequest**](V1SystemsSignupProtectPostRequest.md) |  |  |

### Return type

[**V1SystemsSignupProtectPost200Response**](V1SystemsSignupProtectPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key in the requiems-api-key header |  -  |
| **422** | All of email, phone, and ip_address were omitted. At least one is required. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1systemsuserverifypost"></a>
# **V1SystemsUserVerifyPost**
> V1SystemsUserVerifyPost200Response V1SystemsUserVerifyPost (V1SystemsUserVerifyPostRequest v1SystemsUserVerifyPostRequest)

Verify User

Deep-verify an email address using domain-level signals including WHOIS age, MX records, and domain availability. Optional IP check. Use for high-value or suspicious accounts.

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
    public class V1SystemsUserVerifyPostExample
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
            var apiInstance = new IdentityRiskApi(httpClient, config, httpClientHandler);
            var v1SystemsUserVerifyPostRequest = new V1SystemsUserVerifyPostRequest(); // V1SystemsUserVerifyPostRequest | 

            try
            {
                // Verify User
                V1SystemsUserVerifyPost200Response result = apiInstance.V1SystemsUserVerifyPost(v1SystemsUserVerifyPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IdentityRiskApi.V1SystemsUserVerifyPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsUserVerifyPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Verify User
    ApiResponse<V1SystemsUserVerifyPost200Response> response = apiInstance.V1SystemsUserVerifyPostWithHttpInfo(v1SystemsUserVerifyPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IdentityRiskApi.V1SystemsUserVerifyPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsUserVerifyPostRequest** | [**V1SystemsUserVerifyPostRequest**](V1SystemsUserVerifyPostRequest.md) |  |  |

### Return type

[**V1SystemsUserVerifyPost200Response**](V1SystemsUserVerifyPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **401** | Missing or invalid API key |  -  |
| **422** | email field is missing or empty |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

