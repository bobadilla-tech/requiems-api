# Org.OpenAPITools.Api.DisposableEmailApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1NetworkingDisposableBatchPost**](DisposableEmailApi.md#v1networkingdisposablebatchpost) | **POST** /v1/networking/disposable/batch | Check Batch Emails |
| [**V1NetworkingDisposableCheckPost**](DisposableEmailApi.md#v1networkingdisposablecheckpost) | **POST** /v1/networking/disposable/check | Check Single Email |
| [**V1NetworkingDisposableDomainDomainGet**](DisposableEmailApi.md#v1networkingdisposabledomaindomainget) | **GET** /v1/networking/disposable/domain/{domain} | Check Domain |
| [**V1NetworkingDisposableDomainsGet**](DisposableEmailApi.md#v1networkingdisposabledomainsget) | **GET** /v1/networking/disposable/domains | List Domains (Paginated) |
| [**V1NetworkingDisposableStatsGet**](DisposableEmailApi.md#v1networkingdisposablestatsget) | **GET** /v1/networking/disposable/stats | Get Statistics |

<a id="v1networkingdisposablebatchpost"></a>
# **V1NetworkingDisposableBatchPost**
> V1NetworkingDisposableBatchPost200Response V1NetworkingDisposableBatchPost (V1NetworkingDisposableBatchPostRequest v1NetworkingDisposableBatchPostRequest)

Check Batch Emails

Validate multiple email addresses in a single request (max 100 emails)

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
    public class V1NetworkingDisposableBatchPostExample
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
            var apiInstance = new DisposableEmailApi(httpClient, config, httpClientHandler);
            var v1NetworkingDisposableBatchPostRequest = new V1NetworkingDisposableBatchPostRequest(); // V1NetworkingDisposableBatchPostRequest | 

            try
            {
                // Check Batch Emails
                V1NetworkingDisposableBatchPost200Response result = apiInstance.V1NetworkingDisposableBatchPost(v1NetworkingDisposableBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingDisposableBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Check Batch Emails
    ApiResponse<V1NetworkingDisposableBatchPost200Response> response = apiInstance.V1NetworkingDisposableBatchPostWithHttpInfo(v1NetworkingDisposableBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1NetworkingDisposableBatchPostRequest** | [**V1NetworkingDisposableBatchPostRequest**](V1NetworkingDisposableBatchPostRequest.md) |  |  |

### Return type

[**V1NetworkingDisposableBatchPost200Response**](V1NetworkingDisposableBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The request body is missing or malformed; The emails field is missing; Too many emails in the request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1networkingdisposablecheckpost"></a>
# **V1NetworkingDisposableCheckPost**
> V1NetworkingDisposableCheckPost200Response V1NetworkingDisposableCheckPost (V1NetworkingDisposableCheckPostRequest v1NetworkingDisposableCheckPostRequest)

Check Single Email

Validate whether an email address uses a disposable domain

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
    public class V1NetworkingDisposableCheckPostExample
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
            var apiInstance = new DisposableEmailApi(httpClient, config, httpClientHandler);
            var v1NetworkingDisposableCheckPostRequest = new V1NetworkingDisposableCheckPostRequest(); // V1NetworkingDisposableCheckPostRequest | 

            try
            {
                // Check Single Email
                V1NetworkingDisposableCheckPost200Response result = apiInstance.V1NetworkingDisposableCheckPost(v1NetworkingDisposableCheckPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableCheckPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingDisposableCheckPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Check Single Email
    ApiResponse<V1NetworkingDisposableCheckPost200Response> response = apiInstance.V1NetworkingDisposableCheckPostWithHttpInfo(v1NetworkingDisposableCheckPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableCheckPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1NetworkingDisposableCheckPostRequest** | [**V1NetworkingDisposableCheckPostRequest**](V1NetworkingDisposableCheckPostRequest.md) |  |  |

### Return type

[**V1NetworkingDisposableCheckPost200Response**](V1NetworkingDisposableCheckPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The request body is missing or malformed; The email address format is invalid |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1networkingdisposabledomaindomainget"></a>
# **V1NetworkingDisposableDomainDomainGet**
> V1NetworkingDisposableDomainDomainGet200Response V1NetworkingDisposableDomainDomainGet (string domain)

Check Domain

Check if a specific domain is in the disposable blocklist

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
    public class V1NetworkingDisposableDomainDomainGetExample
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
            var apiInstance = new DisposableEmailApi(httpClient, config, httpClientHandler);
            var domain = tempmail.com;  // string | The domain to check

            try
            {
                // Check Domain
                V1NetworkingDisposableDomainDomainGet200Response result = apiInstance.V1NetworkingDisposableDomainDomainGet(domain);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableDomainDomainGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingDisposableDomainDomainGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Check Domain
    ApiResponse<V1NetworkingDisposableDomainDomainGet200Response> response = apiInstance.V1NetworkingDisposableDomainDomainGetWithHttpInfo(domain);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableDomainDomainGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **domain** | **string** | The domain to check |  |

### Return type

[**V1NetworkingDisposableDomainDomainGet200Response**](V1NetworkingDisposableDomainDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The domain parameter is missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1networkingdisposabledomainsget"></a>
# **V1NetworkingDisposableDomainsGet**
> V1NetworkingDisposableDomainsGet200Response V1NetworkingDisposableDomainsGet (int? page = null, int? perPage = null)

List Domains (Paginated)

Get a paginated list of all disposable domains in the blocklist

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
    public class V1NetworkingDisposableDomainsGetExample
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
            var apiInstance = new DisposableEmailApi(httpClient, config, httpClientHandler);
            var page = 1;  // int? | Page number (default: 1) (optional) 
            var perPage = 100;  // int? | Items per page (default: 100, max: 1000) (optional) 

            try
            {
                // List Domains (Paginated)
                V1NetworkingDisposableDomainsGet200Response result = apiInstance.V1NetworkingDisposableDomainsGet(page, perPage);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableDomainsGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingDisposableDomainsGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // List Domains (Paginated)
    ApiResponse<V1NetworkingDisposableDomainsGet200Response> response = apiInstance.V1NetworkingDisposableDomainsGetWithHttpInfo(page, perPage);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableDomainsGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **page** | **int?** | Page number (default: 1) | [optional]  |
| **perPage** | **int?** | Items per page (default: 100, max: 1000) | [optional]  |

### Return type

[**V1NetworkingDisposableDomainsGet200Response**](V1NetworkingDisposableDomainsGet200Response.md)

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

<a id="v1networkingdisposablestatsget"></a>
# **V1NetworkingDisposableStatsGet**
> V1NetworkingDisposableStatsGet200Response V1NetworkingDisposableStatsGet ()

Get Statistics

Get statistics about the disposable email blocklist

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
    public class V1NetworkingDisposableStatsGetExample
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
            var apiInstance = new DisposableEmailApi(httpClient, config, httpClientHandler);

            try
            {
                // Get Statistics
                V1NetworkingDisposableStatsGet200Response result = apiInstance.V1NetworkingDisposableStatsGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableStatsGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1NetworkingDisposableStatsGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Statistics
    ApiResponse<V1NetworkingDisposableStatsGet200Response> response = apiInstance.V1NetworkingDisposableStatsGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DisposableEmailApi.V1NetworkingDisposableStatsGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1NetworkingDisposableStatsGet200Response**](V1NetworkingDisposableStatsGet200Response.md)

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

