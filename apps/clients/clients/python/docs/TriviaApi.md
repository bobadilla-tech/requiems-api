# openapi_client.TriviaApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_entertainment_trivia_batch_post**](TriviaApi.md#v1_entertainment_trivia_batch_post) | **POST** /v1/entertainment/trivia/batch | Get Batch Trivia Questions
[**v1_entertainment_trivia_get**](TriviaApi.md#v1_entertainment_trivia_get) | **GET** /v1/entertainment/trivia | Get Trivia Question


# **v1_entertainment_trivia_batch_post**
> V1EntertainmentTriviaBatchPost200Response v1_entertainment_trivia_batch_post(v1_entertainment_trivia_batch_post_request)

Get Batch Trivia Questions

Returns up to 50 trivia questions in a single request.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_trivia_batch_post200_response import V1EntertainmentTriviaBatchPost200Response
from openapi_client.models.v1_entertainment_trivia_batch_post_request import V1EntertainmentTriviaBatchPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://requiems.xyz"
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
    api_instance = openapi_client.TriviaApi(api_client)
    v1_entertainment_trivia_batch_post_request = openapi_client.V1EntertainmentTriviaBatchPostRequest() # V1EntertainmentTriviaBatchPostRequest | 

    try:
        # Get Batch Trivia Questions
        api_response = api_instance.v1_entertainment_trivia_batch_post(v1_entertainment_trivia_batch_post_request)
        print("The response of TriviaApi->v1_entertainment_trivia_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TriviaApi->v1_entertainment_trivia_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_entertainment_trivia_batch_post_request** | [**V1EntertainmentTriviaBatchPostRequest**](V1EntertainmentTriviaBatchPostRequest.md)|  | 

### Return type

[**V1EntertainmentTriviaBatchPost200Response**](V1EntertainmentTriviaBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Invalid JSON, malformed request body, or unexpected field types. |  -  |
**401** | Missing API key. |  -  |
**403** | Invalid or revoked API key. |  -  |
**422** | The filters array is missing, empty, or contains more than 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_entertainment_trivia_get**
> V1EntertainmentTriviaGet200Response v1_entertainment_trivia_get(category=category, difficulty=difficulty)

Get Trivia Question

Returns a random trivia question with multiple-choice answers. Use the optional category and difficulty query parameters to filter the question pool.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_trivia_get200_response import V1EntertainmentTriviaGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://requiems.xyz"
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
    api_instance = openapi_client.TriviaApi(api_client)
    category = 'category_example' # str | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. (optional)
    difficulty = 'difficulty_example' # str | Filter by difficulty. One of: easy, medium, hard. (optional)

    try:
        # Get Trivia Question
        api_response = api_instance.v1_entertainment_trivia_get(category=category, difficulty=difficulty)
        print("The response of TriviaApi->v1_entertainment_trivia_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TriviaApi->v1_entertainment_trivia_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **str**| Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. | [optional] 
 **difficulty** | **str**| Filter by difficulty. One of: easy, medium, hard. | [optional] 

### Return type

[**V1EntertainmentTriviaGet200Response**](V1EntertainmentTriviaGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | An invalid category or difficulty value was provided |  -  |
**401** | Missing API key |  -  |
**403** | Invalid or revoked API key |  -  |
**404** | No questions match the given category and difficulty combination |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

