# openapi_client.QuotesApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_entertainment_quotes_random_batch_post**](QuotesApi.md#v1_entertainment_quotes_random_batch_post) | **POST** /v1/entertainment/quotes/random/batch | Get Random Quotes (Batch)
[**v1_entertainment_quotes_random_get**](QuotesApi.md#v1_entertainment_quotes_random_get) | **GET** /v1/entertainment/quotes/random | Get Random Quote


# **v1_entertainment_quotes_random_batch_post**
> V1EntertainmentQuotesRandomBatchPost200Response v1_entertainment_quotes_random_batch_post(v1_entertainment_quotes_random_batch_post_request)

Get Random Quotes (Batch)

Returns up to 50 random quotes in a single request. Each quote counts as one unit of usage.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_quotes_random_batch_post200_response import V1EntertainmentQuotesRandomBatchPost200Response
from openapi_client.models.v1_entertainment_quotes_random_batch_post_request import V1EntertainmentQuotesRandomBatchPostRequest
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
    api_instance = openapi_client.QuotesApi(api_client)
    v1_entertainment_quotes_random_batch_post_request = openapi_client.V1EntertainmentQuotesRandomBatchPostRequest() # V1EntertainmentQuotesRandomBatchPostRequest | 

    try:
        # Get Random Quotes (Batch)
        api_response = api_instance.v1_entertainment_quotes_random_batch_post(v1_entertainment_quotes_random_batch_post_request)
        print("The response of QuotesApi->v1_entertainment_quotes_random_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling QuotesApi->v1_entertainment_quotes_random_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_entertainment_quotes_random_batch_post_request** | [**V1EntertainmentQuotesRandomBatchPostRequest**](V1EntertainmentQuotesRandomBatchPostRequest.md)|  | 

### Return type

[**V1EntertainmentQuotesRandomBatchPost200Response**](V1EntertainmentQuotesRandomBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Request body is missing or not valid JSON. |  -  |
**422** | count is missing, less than 1, or greater than 50. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_entertainment_quotes_random_get**
> V1EntertainmentQuotesRandomGet200Response v1_entertainment_quotes_random_get()

Get Random Quote

Returns a random inspirational quote with author attribution

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_quotes_random_get200_response import V1EntertainmentQuotesRandomGet200Response
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
    api_instance = openapi_client.QuotesApi(api_client)

    try:
        # Get Random Quote
        api_response = api_instance.v1_entertainment_quotes_random_get()
        print("The response of QuotesApi->v1_entertainment_quotes_random_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling QuotesApi->v1_entertainment_quotes_random_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1EntertainmentQuotesRandomGet200Response**](V1EntertainmentQuotesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**503** | No quotes available in the database |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

