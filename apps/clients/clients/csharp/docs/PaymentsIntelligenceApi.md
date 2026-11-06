# Org.OpenAPITools.Api.PaymentsIntelligenceApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1SystemsPaymentValidatePost**](PaymentsIntelligenceApi.md#v1systemspaymentvalidatepost) | **POST** /v1/systems/payment/validate | Validate Payment |
| [**V1SystemsTransactionRiskPost**](PaymentsIntelligenceApi.md#v1systemstransactionriskpost) | **POST** /v1/systems/transaction/risk | Score Transaction Risk |

<a id="v1systemspaymentvalidatepost"></a>
# **V1SystemsPaymentValidatePost**
> V1SystemsPaymentValidatePost200Response V1SystemsPaymentValidatePost (V1SystemsPaymentValidatePostRequest v1SystemsPaymentValidatePostRequest)

Validate Payment

Validate and cross-check BIN, IBAN, and SWIFT identifiers in one call. Returns per-field results and a consistency check that flags country and bank mismatches.

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
    public class V1SystemsPaymentValidatePostExample
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
            var apiInstance = new PaymentsIntelligenceApi(httpClient, config, httpClientHandler);
            var v1SystemsPaymentValidatePostRequest = new V1SystemsPaymentValidatePostRequest(); // V1SystemsPaymentValidatePostRequest | 

            try
            {
                // Validate Payment
                V1SystemsPaymentValidatePost200Response result = apiInstance.V1SystemsPaymentValidatePost(v1SystemsPaymentValidatePostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PaymentsIntelligenceApi.V1SystemsPaymentValidatePost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsPaymentValidatePostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Validate Payment
    ApiResponse<V1SystemsPaymentValidatePost200Response> response = apiInstance.V1SystemsPaymentValidatePostWithHttpInfo(v1SystemsPaymentValidatePostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PaymentsIntelligenceApi.V1SystemsPaymentValidatePostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsPaymentValidatePostRequest** | [**V1SystemsPaymentValidatePostRequest**](V1SystemsPaymentValidatePostRequest.md) |  |  |

### Return type

[**V1SystemsPaymentValidatePost200Response**](V1SystemsPaymentValidatePost200Response.md)

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
| **422** | All of bin, iban, and swift were omitted. At least one is required. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1systemstransactionriskpost"></a>
# **V1SystemsTransactionRiskPost**
> V1SystemsTransactionRiskPost200Response V1SystemsTransactionRiskPost (V1SystemsTransactionRiskPostRequest v1SystemsTransactionRiskPostRequest)

Score Transaction Risk

Score a transaction for fraud risk by cross-checking the card BIN country against IP geolocation and billing country. Detects VPN, proxy, and TOR use.

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
    public class V1SystemsTransactionRiskPostExample
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
            var apiInstance = new PaymentsIntelligenceApi(httpClient, config, httpClientHandler);
            var v1SystemsTransactionRiskPostRequest = new V1SystemsTransactionRiskPostRequest(); // V1SystemsTransactionRiskPostRequest | 

            try
            {
                // Score Transaction Risk
                V1SystemsTransactionRiskPost200Response result = apiInstance.V1SystemsTransactionRiskPost(v1SystemsTransactionRiskPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PaymentsIntelligenceApi.V1SystemsTransactionRiskPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1SystemsTransactionRiskPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Score Transaction Risk
    ApiResponse<V1SystemsTransactionRiskPost200Response> response = apiInstance.V1SystemsTransactionRiskPostWithHttpInfo(v1SystemsTransactionRiskPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PaymentsIntelligenceApi.V1SystemsTransactionRiskPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1SystemsTransactionRiskPostRequest** | [**V1SystemsTransactionRiskPostRequest**](V1SystemsTransactionRiskPostRequest.md) |  |  |

### Return type

[**V1SystemsTransactionRiskPost200Response**](V1SystemsTransactionRiskPost200Response.md)

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
| **422** | ip_address is not a valid IPv4 or IPv6 address, or card_bin is missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

