# Org.OpenAPITools.Api.PostalCodeApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1PlacesPostalBatchPost**](PostalCodeApi.md#v1placespostalbatchpost) | **POST** /v1/places/postal/batch | Batch Lookup Postal Codes |
| [**V1PlacesPostalCodeGet**](PostalCodeApi.md#v1placespostalcodeget) | **GET** /v1/places/postal/{code} | Lookup Postal Code |

<a id="v1placespostalbatchpost"></a>
# **V1PlacesPostalBatchPost**
> V1PlacesPostalBatchPost200Response V1PlacesPostalBatchPost (V1PlacesPostalBatchPostRequest v1PlacesPostalBatchPostRequest)

Batch Lookup Postal Codes

Look up city, state, and coordinates for up to 50 postal codes in a single request. Results are returned in input order.

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
    public class V1PlacesPostalBatchPostExample
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
            var apiInstance = new PostalCodeApi(httpClient, config, httpClientHandler);
            var v1PlacesPostalBatchPostRequest = new V1PlacesPostalBatchPostRequest(); // V1PlacesPostalBatchPostRequest | 

            try
            {
                // Batch Lookup Postal Codes
                V1PlacesPostalBatchPost200Response result = apiInstance.V1PlacesPostalBatchPost(v1PlacesPostalBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PostalCodeApi.V1PlacesPostalBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesPostalBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch Lookup Postal Codes
    ApiResponse<V1PlacesPostalBatchPost200Response> response = apiInstance.V1PlacesPostalBatchPostWithHttpInfo(v1PlacesPostalBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PostalCodeApi.V1PlacesPostalBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1PlacesPostalBatchPostRequest** | [**V1PlacesPostalBatchPostRequest**](V1PlacesPostalBatchPostRequest.md) |  |  |

### Return type

[**V1PlacesPostalBatchPost200Response**](V1PlacesPostalBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The request body is missing or malformed. |  -  |
| **422** | The items array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1placespostalcodeget"></a>
# **V1PlacesPostalCodeGet**
> V1PlacesPostalCodeGet200Response V1PlacesPostalCodeGet (string code, string? country = null)

Lookup Postal Code

Returns city, state, country, and coordinates for the given postal code.

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
    public class V1PlacesPostalCodeGetExample
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
            var apiInstance = new PostalCodeApi(httpClient, config, httpClientHandler);
            var code = 10001;  // string | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London)
            var country = US;  // string? | ISO 3166-1 alpha-2 country code (default: US) (optional) 

            try
            {
                // Lookup Postal Code
                V1PlacesPostalCodeGet200Response result = apiInstance.V1PlacesPostalCodeGet(code, country);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PostalCodeApi.V1PlacesPostalCodeGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1PlacesPostalCodeGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Lookup Postal Code
    ApiResponse<V1PlacesPostalCodeGet200Response> response = apiInstance.V1PlacesPostalCodeGetWithHttpInfo(code, country);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PostalCodeApi.V1PlacesPostalCodeGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **code** | **string** | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London) |  |
| **country** | **string?** | ISO 3166-1 alpha-2 country code (default: US) | [optional]  |

### Return type

[**V1PlacesPostalCodeGet200Response**](V1PlacesPostalCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **404** | The postal code was not found for the given country. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

