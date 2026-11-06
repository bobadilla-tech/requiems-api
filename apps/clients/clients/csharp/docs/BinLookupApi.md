# Org.OpenAPITools.Api.BinLookupApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1FinanceBinBinGet**](BinLookupApi.md#v1financebinbinget) | **GET** /v1/finance/bin/{bin} | BIN Lookup |

<a id="v1financebinbinget"></a>
# **V1FinanceBinBinGet**
> V1FinanceBinBinGet200Response V1FinanceBinBinGet (string bin)

BIN Lookup

Returns card metadata for the given 6–8 digit BIN prefix.

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
    public class V1FinanceBinBinGetExample
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
            var apiInstance = new BinLookupApi(httpClient, config, httpClientHandler);
            var bin = 424242;  // string | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically.

            try
            {
                // BIN Lookup
                V1FinanceBinBinGet200Response result = apiInstance.V1FinanceBinBinGet(bin);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling BinLookupApi.V1FinanceBinBinGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceBinBinGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // BIN Lookup
    ApiResponse<V1FinanceBinBinGet200Response> response = apiInstance.V1FinanceBinBinGetWithHttpInfo(bin);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling BinLookupApi.V1FinanceBinBinGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **bin** | **string** | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically. |  |

### Return type

[**V1FinanceBinBinGet200Response**](V1FinanceBinBinGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | BIN is not 6–8 digits or contains non-digit characters. |  -  |
| **404** | BIN prefix not found in the database. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

