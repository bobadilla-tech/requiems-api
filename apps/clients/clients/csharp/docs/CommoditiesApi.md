# Org.OpenAPITools.Api.CommoditiesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1FinanceCommoditiesCommodityGet**](CommoditiesApi.md#v1financecommoditiescommodityget) | **GET** /v1/finance/commodities/{commodity} | Get Commodity Price |

<a id="v1financecommoditiescommodityget"></a>
# **V1FinanceCommoditiesCommodityGet**
> V1FinanceCommoditiesCommodityGet200Response V1FinanceCommoditiesCommodityGet (string commodity)

Get Commodity Price

Returns the latest annual average price and up to 10 years of historical data for the requested commodity slug.

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
    public class V1FinanceCommoditiesCommodityGetExample
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
            var apiInstance = new CommoditiesApi(httpClient, config, httpClientHandler);
            var commodity = gold;  // string | Commodity slug (e.g. gold, silver, oil). See supported slugs below.

            try
            {
                // Get Commodity Price
                V1FinanceCommoditiesCommodityGet200Response result = apiInstance.V1FinanceCommoditiesCommodityGet(commodity);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling CommoditiesApi.V1FinanceCommoditiesCommodityGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1FinanceCommoditiesCommodityGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Commodity Price
    ApiResponse<V1FinanceCommoditiesCommodityGet200Response> response = apiInstance.V1FinanceCommoditiesCommodityGetWithHttpInfo(commodity);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling CommoditiesApi.V1FinanceCommoditiesCommodityGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **commodity** | **string** | Commodity slug (e.g. gold, silver, oil). See supported slugs below. |  |

### Return type

[**V1FinanceCommoditiesCommodityGet200Response**](V1FinanceCommoditiesCommodityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **404** | No data found for the given commodity slug. Check the supported slugs list. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

