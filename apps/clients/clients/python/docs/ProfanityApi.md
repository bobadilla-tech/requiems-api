# openapi_client.ProfanityApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_validation_profanity_batch_post**](ProfanityApi.md#v1_validation_profanity_batch_post) | **POST** /v1/validation/profanity/batch | Batch Check Profanity
[**v1_validation_profanity_post**](ProfanityApi.md#v1_validation_profanity_post) | **POST** /v1/validation/profanity | Check Profanity


# **v1_validation_profanity_batch_post**
> V1ValidationProfanityBatchPost200Response v1_validation_profanity_batch_post(v1_validation_profanity_batch_post_request)

Batch Check Profanity

Check up to 50 texts for profanity in a single request. Results are returned in input order.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_validation_profanity_batch_post200_response import V1ValidationProfanityBatchPost200Response
from openapi_client.models.v1_validation_profanity_batch_post_request import V1ValidationProfanityBatchPostRequest
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
    api_instance = openapi_client.ProfanityApi(api_client)
    v1_validation_profanity_batch_post_request = openapi_client.V1ValidationProfanityBatchPostRequest() # V1ValidationProfanityBatchPostRequest | 

    try:
        # Batch Check Profanity
        api_response = api_instance.v1_validation_profanity_batch_post(v1_validation_profanity_batch_post_request)
        print("The response of ProfanityApi->v1_validation_profanity_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProfanityApi->v1_validation_profanity_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_validation_profanity_batch_post_request** | [**V1ValidationProfanityBatchPostRequest**](V1ValidationProfanityBatchPostRequest.md)|  | 

### Return type

[**V1ValidationProfanityBatchPost200Response**](V1ValidationProfanityBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing or malformed. |  -  |
**422** | The texts array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_validation_profanity_post**
> V1ValidationProfanityPost200Response v1_validation_profanity_post(v1_validation_profanity_post_request)

Check Profanity

Checks text for profanity, returning a censored version and the list of flagged words.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_validation_profanity_post200_response import V1ValidationProfanityPost200Response
from openapi_client.models.v1_validation_profanity_post_request import V1ValidationProfanityPostRequest
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
    api_instance = openapi_client.ProfanityApi(api_client)
    v1_validation_profanity_post_request = openapi_client.V1ValidationProfanityPostRequest() # V1ValidationProfanityPostRequest | 

    try:
        # Check Profanity
        api_response = api_instance.v1_validation_profanity_post(v1_validation_profanity_post_request)
        print("The response of ProfanityApi->v1_validation_profanity_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProfanityApi->v1_validation_profanity_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_validation_profanity_post_request** | [**V1ValidationProfanityPostRequest**](V1ValidationProfanityPostRequest.md)|  | 

### Return type

[**V1ValidationProfanityPost200Response**](V1ValidationProfanityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing or malformed. |  -  |
**422** | The text field is missing or empty. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

