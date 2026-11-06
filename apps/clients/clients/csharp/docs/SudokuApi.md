# Org.OpenAPITools.Api.SudokuApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**V1EntertainmentSudokuBatchPost**](SudokuApi.md#v1entertainmentsudokubatchpost) | **POST** /v1/entertainment/sudoku/batch | Batch Generate Sudoku Puzzles |
| [**V1EntertainmentSudokuGet**](SudokuApi.md#v1entertainmentsudokuget) | **GET** /v1/entertainment/sudoku | Get Sudoku Puzzle |

<a id="v1entertainmentsudokubatchpost"></a>
# **V1EntertainmentSudokuBatchPost**
> V1EntertainmentSudokuBatchPost200Response V1EntertainmentSudokuBatchPost (V1EntertainmentSudokuBatchPostRequest v1EntertainmentSudokuBatchPostRequest)

Batch Generate Sudoku Puzzles

Generate up to 20 Sudoku puzzles in a single request. Results are returned in the same order as the input array. Each puzzle in the batch counts as one unit of API usage.

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
    public class V1EntertainmentSudokuBatchPostExample
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
            var apiInstance = new SudokuApi(httpClient, config, httpClientHandler);
            var v1EntertainmentSudokuBatchPostRequest = new V1EntertainmentSudokuBatchPostRequest(); // V1EntertainmentSudokuBatchPostRequest | 

            try
            {
                // Batch Generate Sudoku Puzzles
                V1EntertainmentSudokuBatchPost200Response result = apiInstance.V1EntertainmentSudokuBatchPost(v1EntertainmentSudokuBatchPostRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling SudokuApi.V1EntertainmentSudokuBatchPost: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentSudokuBatchPostWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Batch Generate Sudoku Puzzles
    ApiResponse<V1EntertainmentSudokuBatchPost200Response> response = apiInstance.V1EntertainmentSudokuBatchPostWithHttpInfo(v1EntertainmentSudokuBatchPostRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling SudokuApi.V1EntertainmentSudokuBatchPostWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **v1EntertainmentSudokuBatchPostRequest** | [**V1EntertainmentSudokuBatchPostRequest**](V1EntertainmentSudokuBatchPostRequest.md) |  |  |

### Return type

[**V1EntertainmentSudokuBatchPost200Response**](V1EntertainmentSudokuBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The request body is missing or contains malformed JSON. |  -  |
| **401** | Missing API key |  -  |
| **403** | Invalid or revoked API key |  -  |
| **422** | The puzzles array is missing, empty, exceeds 20 items, or contains a value other than easy, medium, or hard. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="v1entertainmentsudokuget"></a>
# **V1EntertainmentSudokuGet**
> V1EntertainmentSudokuGet200Response V1EntertainmentSudokuGet (string? difficulty = null)

Get Sudoku Puzzle

Returns a randomly generated Sudoku puzzle and its solution. Difficulty defaults to medium when not specified.

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
    public class V1EntertainmentSudokuGetExample
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
            var apiInstance = new SudokuApi(httpClient, config, httpClientHandler);
            var difficulty = "difficulty_example";  // string? | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. (optional) 

            try
            {
                // Get Sudoku Puzzle
                V1EntertainmentSudokuGet200Response result = apiInstance.V1EntertainmentSudokuGet(difficulty);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling SudokuApi.V1EntertainmentSudokuGet: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the V1EntertainmentSudokuGetWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get Sudoku Puzzle
    ApiResponse<V1EntertainmentSudokuGet200Response> response = apiInstance.V1EntertainmentSudokuGetWithHttpInfo(difficulty);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling SudokuApi.V1EntertainmentSudokuGetWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **difficulty** | **string?** | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. | [optional]  |

### Return type

[**V1EntertainmentSudokuGet200Response**](V1EntertainmentSudokuGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful response |  -  |
| **400** | The difficulty parameter is not one of easy, medium, or hard |  -  |
| **401** | Missing API key |  -  |
| **403** | Invalid or revoked API key |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

