# openapi_client.SudokuApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_entertainment_sudoku_batch_post**](SudokuApi.md#v1_entertainment_sudoku_batch_post) | **POST** /v1/entertainment/sudoku/batch | Batch Generate Sudoku Puzzles
[**v1_entertainment_sudoku_get**](SudokuApi.md#v1_entertainment_sudoku_get) | **GET** /v1/entertainment/sudoku | Get Sudoku Puzzle


# **v1_entertainment_sudoku_batch_post**
> V1EntertainmentSudokuBatchPost200Response v1_entertainment_sudoku_batch_post(v1_entertainment_sudoku_batch_post_request)

Batch Generate Sudoku Puzzles

Generate up to 20 Sudoku puzzles in a single request. Results are returned in the same order as the input array. Each puzzle in the batch counts as one unit of API usage.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_sudoku_batch_post200_response import V1EntertainmentSudokuBatchPost200Response
from openapi_client.models.v1_entertainment_sudoku_batch_post_request import V1EntertainmentSudokuBatchPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.SudokuApi(api_client)
    v1_entertainment_sudoku_batch_post_request = openapi_client.V1EntertainmentSudokuBatchPostRequest() # V1EntertainmentSudokuBatchPostRequest | 

    try:
        # Batch Generate Sudoku Puzzles
        api_response = api_instance.v1_entertainment_sudoku_batch_post(v1_entertainment_sudoku_batch_post_request)
        print("The response of SudokuApi->v1_entertainment_sudoku_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SudokuApi->v1_entertainment_sudoku_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_entertainment_sudoku_batch_post_request** | [**V1EntertainmentSudokuBatchPostRequest**](V1EntertainmentSudokuBatchPostRequest.md)|  | 

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
**200** | Successful response |  -  |
**400** | The request body is missing or contains malformed JSON. |  -  |
**401** | Missing API key |  -  |
**403** | Invalid or revoked API key |  -  |
**422** | The puzzles array is missing, empty, exceeds 20 items, or contains a value other than easy, medium, or hard. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_entertainment_sudoku_get**
> V1EntertainmentSudokuGet200Response v1_entertainment_sudoku_get(difficulty=difficulty)

Get Sudoku Puzzle

Returns a randomly generated Sudoku puzzle and its solution. Difficulty defaults to medium when not specified.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_sudoku_get200_response import V1EntertainmentSudokuGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.SudokuApi(api_client)
    difficulty = 'difficulty_example' # str | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. (optional)

    try:
        # Get Sudoku Puzzle
        api_response = api_instance.v1_entertainment_sudoku_get(difficulty=difficulty)
        print("The response of SudokuApi->v1_entertainment_sudoku_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SudokuApi->v1_entertainment_sudoku_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **difficulty** | **str**| Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. | [optional] 

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
**200** | Successful response |  -  |
**400** | The difficulty parameter is not one of easy, medium, or hard |  -  |
**401** | Missing API key |  -  |
**403** | Invalid or revoked API key |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

