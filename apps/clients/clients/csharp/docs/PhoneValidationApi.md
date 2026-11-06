# Org.OpenAPITools.Api.PhoneValidationApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1ValidationPhoneBatchPost**](PhoneValidationApi.md#v1validationphonebatchpost) | **POST** /v1/validation/phone/batch | Batch Validate Phone Numbers |
| [**V1ValidationPhoneGet**](PhoneValidationApi.md#v1validationphoneget) | **GET** /v1/validation/phone | Validate Phone Number |

<a id="v1validationphonebatchpost"></a>
# **V1ValidationPhoneBatchPost**
> V1FinanceIbanBatchPost200Response V1ValidationPhoneBatchPost (V1ValidationPhoneBatchPostRequest v1ValidationPhoneBatchPostRequest)

Batch Validate Phone Numbers

Validates up to 50 phone numbers in a single request. Results are returned in the same order as the input.

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
    public class V1ValidationPhoneBatchPostExample
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
            var apiInstance = new PhoneValidationApi(httpClient, config, httpClientHandler);
            var v1ValidationPhoneBatchPostRequest = new V1ValidationPhoneBatchPostRequest(); // V1ValidationPhoneBatchPostRequest | 

            try
            {
                // Batch Validate Phone Numbers
                V1FinanceIbanBatchPost200Response result = apiInstance.V1ValidationPhoneBatchPost(v1ValidationPhoneBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PhoneValidationApi.V1ValidationPhoneBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1ValidationPhoneBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch Validate Phone Numbers
    ApiResponse<V1FinanceIbanBatchPost200Response> response = apiInstance.V1ValidationPhoneBatchPostWithHttpInfo(v1ValidationPhoneBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PhoneValidationApi.V1ValidationPhoneBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1ValidationPhoneBatchPostRequest** | [**V1ValidationPhoneBatchPostRequest**](V1ValidationPhoneBatchPostRequest.md) |  |  |

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

<a id="v1validationphoneget"></a>
# **V1ValidationPhoneGet**
> V1ValidationPhoneGet200Response V1ValidationPhoneGet (string number)

Validate Phone Number

Validates a single phone number and returns its country, type, formatted representation, carrier, and VOIP/virtual risk flags.

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
    public class V1ValidationPhoneGetExample
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
            var apiInstance = new PhoneValidationApi(httpClient, config, httpClientHandler);
            var number = +447400123456;  // string | The phone number to validate. Must include the country calling code (e.g. +12015551234).

            try
            {
                // Validate Phone Number
                V1ValidationPhoneGet200Response result = apiInstance.V1ValidationPhoneGet(number);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PhoneValidationApi.V1ValidationPhoneGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1ValidationPhoneGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Validate Phone Number
    ApiResponse<V1ValidationPhoneGet200Response> response = apiInstance.V1ValidationPhoneGetWithHttpInfo(number);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PhoneValidationApi.V1ValidationPhoneGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **number** | **string** | The phone number to validate. Must include the country calling code (e.g. +12015551234). |  |

### Return type

[**V1ValidationPhoneGet200Response**](V1ValidationPhoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The number query parameter is missing. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

