# Org.OpenAPITools.Api.IbanApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1FinanceIbanBatchPost**](IbanApi.md#v1financeibanbatchpost) | **POST** /v1/finance/iban/batch | Batch Validate IBANs |
| [**V1FinanceIbanIbanGet**](IbanApi.md#v1financeibanibanget) | **GET** /v1/finance/iban/{iban} | Validate IBAN |

<a id="v1financeibanbatchpost"></a>
# **V1FinanceIbanBatchPost**
> V1FinanceIbanBatchPost200Response V1FinanceIbanBatchPost (V1FinanceIbanBatchPostRequest v1FinanceIbanBatchPostRequest)

Batch Validate IBANs

Validates up to 50 iban numbers in a single request. Results are returned in the same order as the input.

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
    public class V1FinanceIbanBatchPostExample
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
            var apiInstance = new IbanApi(httpClient, config, httpClientHandler);
            var v1FinanceIbanBatchPostRequest = new V1FinanceIbanBatchPostRequest(); // V1FinanceIbanBatchPostRequest | 

            try
            {
                // Batch Validate IBANs
                V1FinanceIbanBatchPost200Response result = apiInstance.V1FinanceIbanBatchPost(v1FinanceIbanBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IbanApi.V1FinanceIbanBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceIbanBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch Validate IBANs
    ApiResponse<V1FinanceIbanBatchPost200Response> response = apiInstance.V1FinanceIbanBatchPostWithHttpInfo(v1FinanceIbanBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IbanApi.V1FinanceIbanBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1FinanceIbanBatchPostRequest** | [**V1FinanceIbanBatchPostRequest**](V1FinanceIbanBatchPostRequest.md) |  |  |

### Return type

[**V1FinanceIbanBatchPost200Response**](V1FinanceIbanBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **422** | The numbers array is missing, empty, or contains more than 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1financeibanibanget"></a>
# **V1FinanceIbanIbanGet**
> V1FinanceIbanIbanGet200Response V1FinanceIbanIbanGet (string iban)

Validate IBAN

Validates an IBAN and returns the country, bank code, and account number. Spaces in the input are stripped automatically. Always returns HTTP 200 — check the valid field to determine whether the IBAN is valid.

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
    public class V1FinanceIbanIbanGetExample
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
            var apiInstance = new IbanApi(httpClient, config, httpClientHandler);
            var iban = DE89370400440532013000;  // string | The IBAN to validate. Spaces are stripped. Case-insensitive.

            try
            {
                // Validate IBAN
                V1FinanceIbanIbanGet200Response result = apiInstance.V1FinanceIbanIbanGet(iban);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling IbanApi.V1FinanceIbanIbanGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceIbanIbanGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Validate IBAN
    ApiResponse<V1FinanceIbanIbanGet200Response> response = apiInstance.V1FinanceIbanIbanGetWithHttpInfo(iban);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling IbanApi.V1FinanceIbanIbanGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **iban** | **string** | The IBAN to validate. Spaces are stripped. Case-insensitive. |  |

### Return type

[**V1FinanceIbanIbanGet200Response**](V1FinanceIbanIbanGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **500** | Unexpected server error (e.g. database unreachable). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

