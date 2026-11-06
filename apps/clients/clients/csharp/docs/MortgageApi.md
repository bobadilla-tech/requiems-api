# Org.OpenAPITools.Api.MortgageApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1FinanceMortgageGet**](MortgageApi.md#v1financemortgageget) | **GET** /v1/finance/mortgage | Calculate Mortgage |

<a id="v1financemortgageget"></a>
# **V1FinanceMortgageGet**
> V1FinanceMortgageGet200Response V1FinanceMortgageGet (decimal principal, decimal rate, int years)

Calculate Mortgage

Returns the monthly payment, total cost, and full amortization schedule for a fixed-rate mortgage.

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
    public class V1FinanceMortgageGetExample
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
            var apiInstance = new MortgageApi(httpClient, config, httpClientHandler);
            var principal = 300000;  // decimal | Loan amount in your chosen currency (e.g. 300000 for $300,000)
            var rate = 6.5;  // decimal | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0.
            var years = 30;  // int | Loan term in years (1–50)

            try
            {
                // Calculate Mortgage
                V1FinanceMortgageGet200Response result = apiInstance.V1FinanceMortgageGet(principal, rate, years);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling MortgageApi.V1FinanceMortgageGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceMortgageGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Calculate Mortgage
    ApiResponse<V1FinanceMortgageGet200Response> response = apiInstance.V1FinanceMortgageGetWithHttpInfo(principal, rate, years);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling MortgageApi.V1FinanceMortgageGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **principal** | **decimal** | Loan amount in your chosen currency (e.g. 300000 for $300,000) |  |
| **rate** | **decimal** | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0. |  |
| **years** | **int** | Loan term in years (1–50) |  |

### Return type

[**V1FinanceMortgageGet200Response**](V1FinanceMortgageGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | A required parameter is missing, not a valid number, or out of range (e.g. years &gt; 50 or rate &lt;&#x3D; 0). |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

