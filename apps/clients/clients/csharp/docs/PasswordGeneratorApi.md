# Org.OpenAPITools.Api.PasswordGeneratorApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1TechnologyPasswordGet**](PasswordGeneratorApi.md#v1technologypasswordget) | **GET** /v1/technology/password | Generate Password |

<a id="v1technologypasswordget"></a>
# **V1TechnologyPasswordGet**
> V1TechnologyPasswordGet200Response V1TechnologyPasswordGet (int? length = null, bool? uppercase = null, bool? numbers = null, bool? symbols = null)

Generate Password

Generate a cryptographically secure random password with customizable character sets and length

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
    public class V1TechnologyPasswordGetExample
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
            var apiInstance = new PasswordGeneratorApi(httpClient, config, httpClientHandler);
            var length = 16;  // int? | Password length (8-128 characters) (optional) 
            var uppercase = true;  // bool? | Include uppercase letters (A-Z) (optional) 
            var numbers = true;  // bool? | Include numbers (0-9) (optional) 
            var symbols = true;  // bool? | Include special characters (!@#$%^&*()-_=+[]{}|;:,.<>?) (optional) 

            try
            {
                // Generate Password
                V1TechnologyPasswordGet200Response result = apiInstance.V1TechnologyPasswordGet(length, uppercase, numbers, symbols);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PasswordGeneratorApi.V1TechnologyPasswordGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1TechnologyPasswordGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Generate Password
    ApiResponse<V1TechnologyPasswordGet200Response> response = apiInstance.V1TechnologyPasswordGetWithHttpInfo(length, uppercase, numbers, symbols);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PasswordGeneratorApi.V1TechnologyPasswordGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **length** | **int?** | Password length (8-128 characters) | [optional]  |
| **uppercase** | **bool?** | Include uppercase letters (A-Z) | [optional]  |
| **numbers** | **bool?** | Include numbers (0-9) | [optional]  |
| **symbols** | **bool?** | Include special characters (!@#$%^&amp;*()-_&#x3D;+[]{}|;:,.&lt;&gt;?) | [optional]  |

### Return type

[**V1TechnologyPasswordGet200Response**](V1TechnologyPasswordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The length parameter is out of valid range (8-128) |  -  |
| **500** | Failed to generate password (rare cryptographic failure) |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

