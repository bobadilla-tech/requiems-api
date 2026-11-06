# Org.OpenAPITools.Api.ExchangeRateApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1FinanceConvertGet**](ExchangeRateApi.md#v1financeconvertget) | **GET** /v1/finance/convert | Convert Currency |
| [**V1FinanceExchangeRateGet**](ExchangeRateApi.md#v1financeexchangerateget) | **GET** /v1/finance/exchange-rate | Get Exchange Rate |

<a id="v1financeconvertget"></a>
# **V1FinanceConvertGet**
> V1FinanceConvertGet200Response V1FinanceConvertGet (string from, string to, decimal amount)

Convert Currency

Converts an amount from one currency to another and returns the rate alongside the converted value.

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
    public class V1FinanceConvertGetExample
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
            var apiInstance = new ExchangeRateApi(httpClient, config, httpClientHandler);
            var from = USD;  // string | ISO 4217 source currency code (3 letters, e.g. USD)
            var to = EUR;  // string | ISO 4217 target currency code (3 letters, e.g. EUR)
            var amount = 100;  // decimal | Amount to convert. Must be greater than 0.

            try
            {
                // Convert Currency
                V1FinanceConvertGet200Response result = apiInstance.V1FinanceConvertGet(from, to, amount);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ExchangeRateApi.V1FinanceConvertGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceConvertGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Convert Currency
    ApiResponse<V1FinanceConvertGet200Response> response = apiInstance.V1FinanceConvertGetWithHttpInfo(from, to, amount);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ExchangeRateApi.V1FinanceConvertGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **from** | **string** | ISO 4217 source currency code (3 letters, e.g. USD) |  |
| **to** | **string** | ISO 4217 target currency code (3 letters, e.g. EUR) |  |
| **amount** | **decimal** | Amount to convert. Must be greater than 0. |  |

### Return type

[**V1FinanceConvertGet200Response**](V1FinanceConvertGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | A required parameter is missing, the currency code is not 3 alphabetic characters, or the amount is 0 or negative. |  -  |
| **422** | One or both currency codes are not recognised by the upstream data source. |  -  |
| **503** | The exchange rate data source is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1financeexchangerateget"></a>
# **V1FinanceExchangeRateGet**
> V1FinanceExchangeRateGet200Response V1FinanceExchangeRateGet (string from, string to)

Get Exchange Rate

Returns the current exchange rate between two currencies.

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
    public class V1FinanceExchangeRateGetExample
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
            var apiInstance = new ExchangeRateApi(httpClient, config, httpClientHandler);
            var from = USD;  // string | ISO 4217 source currency code (3 letters, e.g. USD)
            var to = EUR;  // string | ISO 4217 target currency code (3 letters, e.g. EUR)

            try
            {
                // Get Exchange Rate
                V1FinanceExchangeRateGet200Response result = apiInstance.V1FinanceExchangeRateGet(from, to);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ExchangeRateApi.V1FinanceExchangeRateGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceExchangeRateGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Exchange Rate
    ApiResponse<V1FinanceExchangeRateGet200Response> response = apiInstance.V1FinanceExchangeRateGetWithHttpInfo(from, to);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling ExchangeRateApi.V1FinanceExchangeRateGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **from** | **string** | ISO 4217 source currency code (3 letters, e.g. USD) |  |
| **to** | **string** | ISO 4217 target currency code (3 letters, e.g. EUR) |  |

### Return type

[**V1FinanceExchangeRateGet200Response**](V1FinanceExchangeRateGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | A required parameter is missing or the currency code is not exactly 3 alphabetic characters. |  -  |
| **422** | One or both currency codes are not recognised by the upstream data source. |  -  |
| **503** | The exchange rate data source is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

