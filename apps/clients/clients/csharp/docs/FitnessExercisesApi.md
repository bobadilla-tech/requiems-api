# Org.OpenAPITools.Api.FitnessExercisesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1HealthBodyPartsGet**](FitnessExercisesApi.md#v1healthbodypartsget) | **GET** /v1/health/body-parts | List Body Parts |
| [**V1HealthEquipmentGet**](FitnessExercisesApi.md#v1healthequipmentget) | **GET** /v1/health/equipment | List Equipment |
| [**V1HealthExercisesGet**](FitnessExercisesApi.md#v1healthexercisesget) | **GET** /v1/health/exercises | List Exercises |
| [**V1HealthExercisesIdGet**](FitnessExercisesApi.md#v1healthexercisesidget) | **GET** /v1/health/exercises/{id} | Get Exercise by ID |
| [**V1HealthExercisesRandomGet**](FitnessExercisesApi.md#v1healthexercisesrandomget) | **GET** /v1/health/exercises/random | Random Exercise |
| [**V1HealthMusclesGet**](FitnessExercisesApi.md#v1healthmusclesget) | **GET** /v1/health/muscles | List Muscles |

<a id="v1healthbodypartsget"></a>
# **V1HealthBodyPartsGet**
> V1HealthBodyPartsGet200Response V1HealthBodyPartsGet ()

List Body Parts

Returns a sorted list of all distinct body part values present in the dataset. Use these as valid values for the body_part filter.

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
    public class V1HealthBodyPartsGetExample
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
            var apiInstance = new FitnessExercisesApi(httpClient, config, httpClientHandler);

            try
            {
                // List Body Parts
                V1HealthBodyPartsGet200Response result = apiInstance.V1HealthBodyPartsGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling FitnessExercisesApi.V1HealthBodyPartsGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1HealthBodyPartsGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // List Body Parts
    ApiResponse<V1HealthBodyPartsGet200Response> response = apiInstance.V1HealthBodyPartsGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling FitnessExercisesApi.V1HealthBodyPartsGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1HealthBodyPartsGet200Response**](V1HealthBodyPartsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1healthequipmentget"></a>
# **V1HealthEquipmentGet**
> V1HealthEquipmentGet200Response V1HealthEquipmentGet ()

List Equipment

Returns a sorted list of all distinct equipment types. Use these as valid values for the equipment filter.

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
    public class V1HealthEquipmentGetExample
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
            var apiInstance = new FitnessExercisesApi(httpClient, config, httpClientHandler);

            try
            {
                // List Equipment
                V1HealthEquipmentGet200Response result = apiInstance.V1HealthEquipmentGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling FitnessExercisesApi.V1HealthEquipmentGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1HealthEquipmentGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // List Equipment
    ApiResponse<V1HealthEquipmentGet200Response> response = apiInstance.V1HealthEquipmentGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling FitnessExercisesApi.V1HealthEquipmentGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1HealthEquipmentGet200Response**](V1HealthEquipmentGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1healthexercisesget"></a>
# **V1HealthExercisesGet**
> V1HealthExercisesGet200Response V1HealthExercisesGet (string? bodyPart = null, string? equipment = null, string? muscle = null, string? search = null, int? page = null, int? perPage = null)

List Exercises

Returns a paginated list of exercises. All filter parameters are optional and combinable.

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
    public class V1HealthExercisesGetExample
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
            var apiInstance = new FitnessExercisesApi(httpClient, config, httpClientHandler);
            var bodyPart = chest;  // string? | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. (optional) 
            var equipment = barbell;  // string? | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. (optional) 
            var muscle = biceps;  // string? | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. (optional) 
            var search = bench press;  // string? | Full-text search on exercise name. (optional) 
            var page = 1;  // int? | Page number (default: 1) (optional) 
            var perPage = 20;  // int? | Results per page, 1–100 (default: 20) (optional) 

            try
            {
                // List Exercises
                V1HealthExercisesGet200Response result = apiInstance.V1HealthExercisesGet(bodyPart, equipment, muscle, search, page, perPage);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling FitnessExercisesApi.V1HealthExercisesGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1HealthExercisesGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // List Exercises
    ApiResponse<V1HealthExercisesGet200Response> response = apiInstance.V1HealthExercisesGetWithHttpInfo(bodyPart, equipment, muscle, search, page, perPage);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling FitnessExercisesApi.V1HealthExercisesGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **bodyPart** | **string?** | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. | [optional]  |
| **equipment** | **string?** | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. | [optional]  |
| **muscle** | **string?** | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. | [optional]  |
| **search** | **string?** | Full-text search on exercise name. | [optional]  |
| **page** | **int?** | Page number (default: 1) | [optional]  |
| **perPage** | **int?** | Results per page, 1–100 (default: 20) | [optional]  |

### Return type

[**V1HealthExercisesGet200Response**](V1HealthExercisesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | A query parameter has an invalid value (e.g. per_page out of range). |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1healthexercisesidget"></a>
# **V1HealthExercisesIdGet**
> V1HealthExercisesIdGet200Response V1HealthExercisesIdGet (int id)

Get Exercise by ID

Returns a single exercise by its numeric ID.

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
    public class V1HealthExercisesIdGetExample
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
            var apiInstance = new FitnessExercisesApi(httpClient, config, httpClientHandler);
            var id = 1;  // int | Numeric exercise ID

            try
            {
                // Get Exercise by ID
                V1HealthExercisesIdGet200Response result = apiInstance.V1HealthExercisesIdGet(id);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling FitnessExercisesApi.V1HealthExercisesIdGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1HealthExercisesIdGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Exercise by ID
    ApiResponse<V1HealthExercisesIdGet200Response> response = apiInstance.V1HealthExercisesIdGetWithHttpInfo(id);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling FitnessExercisesApi.V1HealthExercisesIdGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **id** | **int** | Numeric exercise ID |  |

### Return type

[**V1HealthExercisesIdGet200Response**](V1HealthExercisesIdGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The id parameter is not a positive integer. |  -  |
| **404** | No exercise exists with the given ID. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1healthexercisesrandomget"></a>
# **V1HealthExercisesRandomGet**
> V1HealthExercisesRandomGet200Response V1HealthExercisesRandomGet (string? bodyPart = null, string? equipment = null, string? muscle = null, string? search = null)

Random Exercise

Returns a single randomly selected exercise. Accepts the same filter parameters as the list endpoint, so you can get a random chest exercise, a random bodyweight exercise, etc.

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
    public class V1HealthExercisesRandomGetExample
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
            var apiInstance = new FitnessExercisesApi(httpClient, config, httpClientHandler);
            var bodyPart = back;  // string? | Restrict random selection to this body part. (optional) 
            var equipment = body weight;  // string? | Restrict random selection to this equipment type. (optional) 
            var muscle = glutes;  // string? | Restrict random selection to exercises targeting this muscle. (optional) 
            var search = curl;  // string? | Restrict random selection to exercises matching this search term. (optional) 

            try
            {
                // Random Exercise
                V1HealthExercisesRandomGet200Response result = apiInstance.V1HealthExercisesRandomGet(bodyPart, equipment, muscle, search);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling FitnessExercisesApi.V1HealthExercisesRandomGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1HealthExercisesRandomGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Random Exercise
    ApiResponse<V1HealthExercisesRandomGet200Response> response = apiInstance.V1HealthExercisesRandomGetWithHttpInfo(bodyPart, equipment, muscle, search);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling FitnessExercisesApi.V1HealthExercisesRandomGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **bodyPart** | **string?** | Restrict random selection to this body part. | [optional]  |
| **equipment** | **string?** | Restrict random selection to this equipment type. | [optional]  |
| **muscle** | **string?** | Restrict random selection to exercises targeting this muscle. | [optional]  |
| **search** | **string?** | Restrict random selection to exercises matching this search term. | [optional]  |

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **404** | No exercises match the given filters. |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1healthmusclesget"></a>
# **V1HealthMusclesGet**
> V1HealthMusclesGet200Response V1HealthMusclesGet ()

List Muscles

Returns a sorted list of all distinct muscle names (combining target and secondary muscles). Use these as valid values for the muscle filter.

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
    public class V1HealthMusclesGetExample
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
            var apiInstance = new FitnessExercisesApi(httpClient, config, httpClientHandler);

            try
            {
                // List Muscles
                V1HealthMusclesGet200Response result = apiInstance.V1HealthMusclesGet();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling FitnessExercisesApi.V1HealthMusclesGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1HealthMusclesGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // List Muscles
    ApiResponse<V1HealthMusclesGet200Response> response = apiInstance.V1HealthMusclesGetWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling FitnessExercisesApi.V1HealthMusclesGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**V1HealthMusclesGet200Response**](V1HealthMusclesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

