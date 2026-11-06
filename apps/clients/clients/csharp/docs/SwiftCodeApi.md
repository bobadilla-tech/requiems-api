# Org.OpenAPITools.Api.SwiftCodeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1FinanceSwiftCodeGet**](SwiftCodeApi.md#v1financeswiftcodeget) | **GET** /v1/finance/swift/{code} | Get SWIFT Code |
| [**V1FinanceSwiftGet**](SwiftCodeApi.md#v1financeswiftget) | **GET** /v1/finance/swift | List SWIFT Codes |

<a id="v1financeswiftcodeget"></a>
# **V1FinanceSwiftCodeGet**
> V1FinanceSwiftCodeGet200Response V1FinanceSwiftCodeGet (string code)

Get SWIFT Code

Look up bank metadata for a SWIFT/BIC code.

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
    public class V1FinanceSwiftCodeGetExample
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
            var apiInstance = new SwiftCodeApi(httpClient, config, httpClientHandler);
            var code = DEUTDEDB;  // string | SWIFT/BIC code (8 or 11 alphanumeric characters)

            try
            {
                // Get SWIFT Code
                V1FinanceSwiftCodeGet200Response result = apiInstance.V1FinanceSwiftCodeGet(code);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling SwiftCodeApi.V1FinanceSwiftCodeGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceSwiftCodeGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get SWIFT Code
    ApiResponse<V1FinanceSwiftCodeGet200Response> response = apiInstance.V1FinanceSwiftCodeGetWithHttpInfo(code);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling SwiftCodeApi.V1FinanceSwiftCodeGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **code** | **string** | SWIFT/BIC code (8 or 11 alphanumeric characters) |  |

### Return type

[**V1FinanceSwiftCodeGet200Response**](V1FinanceSwiftCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Invalid SWIFT/BIC format (must be 8 or 11 valid characters). |  -  |
| **404** | SWIFT/BIC code not found in the dataset. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1financeswiftget"></a>
# **V1FinanceSwiftGet**
> V1HealthExercisesRandomGet200Response V1FinanceSwiftGet (string? countryCode = null, string? bankCode = null, string? q = null, int? limit = null, int? offset = null)

List SWIFT Codes

List SWIFT records with optional filters and pagination.

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
    public class V1FinanceSwiftGetExample
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
            var apiInstance = new SwiftCodeApi(httpClient, config, httpClientHandler);
            var countryCode = DE;  // string? | Optional 2-letter country code filter (e.g. DE, US) (optional) 
            var bankCode = DEUT;  // string? | Optional 4-letter bank code filter (e.g. DEUT) (optional) 
            var q = deutsche;  // string? | Optional text search across swift_code, bank_name, and city (optional) 
            var limit = 50;  // int? | Max rows to return (default 50, max 200) (optional) 
            var offset = 0;  // int? | Number of rows to skip (default 0) (optional) 

            try
            {
                // List SWIFT Codes
                V1HealthExercisesRandomGet200Response result = apiInstance.V1FinanceSwiftGet(countryCode, bankCode, q, limit, offset);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling SwiftCodeApi.V1FinanceSwiftGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceSwiftGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // List SWIFT Codes
    ApiResponse<V1HealthExercisesRandomGet200Response> response = apiInstance.V1FinanceSwiftGetWithHttpInfo(countryCode, bankCode, q, limit, offset);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling SwiftCodeApi.V1FinanceSwiftGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **countryCode** | **string?** | Optional 2-letter country code filter (e.g. DE, US) | [optional]  |
| **bankCode** | **string?** | Optional 4-letter bank code filter (e.g. DEUT) | [optional]  |
| **q** | **string?** | Optional text search across swift_code, bank_name, and city | [optional]  |
| **limit** | **int?** | Max rows to return (default 50, max 200) | [optional]  |
| **offset** | **int?** | Number of rows to skip (default 0) | [optional]  |

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | Invalid filter or pagination parameter. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

