# Org.OpenAPITools.Api.CryptoApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1FinanceCryptoSymbolGet**](CryptoApi.md#v1financecryptosymbolget) | **GET** /v1/finance/crypto/{symbol} | Get Crypto Price |

<a id="v1financecryptosymbolget"></a>
# **V1FinanceCryptoSymbolGet**
> V1FinanceCryptoSymbolGet200Response V1FinanceCryptoSymbolGet (string symbol)

Get Crypto Price

Returns current price data for the given cryptocurrency symbol.

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
    public class V1FinanceCryptoSymbolGetExample
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
            var apiInstance = new CryptoApi(httpClient, config, httpClientHandler);
            var symbol = BTC;  // string | Uppercase ticker symbol (e.g. BTC, ETH, SOL)

            try
            {
                // Get Crypto Price
                V1FinanceCryptoSymbolGet200Response result = apiInstance.V1FinanceCryptoSymbolGet(symbol);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling CryptoApi.V1FinanceCryptoSymbolGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceCryptoSymbolGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Crypto Price
    ApiResponse<V1FinanceCryptoSymbolGet200Response> response = apiInstance.V1FinanceCryptoSymbolGetWithHttpInfo(symbol);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling CryptoApi.V1FinanceCryptoSymbolGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **symbol** | **string** | Uppercase ticker symbol (e.g. BTC, ETH, SOL) |  |

### Return type

[**V1FinanceCryptoSymbolGet200Response**](V1FinanceCryptoSymbolGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **422** | The symbol is not in the supported coin list. |  -  |
| **503** | CoinGecko is unavailable or returned an unexpected response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

