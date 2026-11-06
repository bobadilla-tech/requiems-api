# openapi_client.EmailNormalizeApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_text_normalize_batch_post**](EmailNormalizeApi.md#v1_text_normalize_batch_post) | **POST** /v1/text/normalize/batch | Normalize Email Batch
[**v1_text_normalize_post**](EmailNormalizeApi.md#v1_text_normalize_post) | **POST** /v1/text/normalize | Normalize Email


# **v1_text_normalize_batch_post**
> V1TextNormalizeBatchPost200Response v1_text_normalize_batch_post(v1_text_normalize_batch_post_request)

Normalize Email Batch

Normalizes up to 100 email addresses in one request. Results are in the same order as the input. Each item includes valid (boolean); when false, only original and message are set. Usage is billed per email processed (see gateway usage headers).

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_normalize_batch_post200_response import V1TextNormalizeBatchPost200Response
from openapi_client.models.v1_text_normalize_batch_post_request import V1TextNormalizeBatchPostRequest
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
    api_instance = openapi_client.EmailNormalizeApi(api_client)
    v1_text_normalize_batch_post_request = openapi_client.V1TextNormalizeBatchPostRequest() # V1TextNormalizeBatchPostRequest | 

    try:
        # Normalize Email Batch
        api_response = api_instance.v1_text_normalize_batch_post(v1_text_normalize_batch_post_request)
        print("The response of EmailNormalizeApi->v1_text_normalize_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EmailNormalizeApi->v1_text_normalize_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_text_normalize_batch_post_request** | [**V1TextNormalizeBatchPostRequest**](V1TextNormalizeBatchPostRequest.md)|  | 

### Return type

[**V1TextNormalizeBatchPost200Response**](V1TextNormalizeBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Invalid JSON or unknown fields in the body |  -  |
**422** | Missing emails, empty array, too many items, or empty string in the array |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_text_normalize_post**
> V1TextNormalizePost200Response v1_text_normalize_post(v1_text_normalize_post_request)

Normalize Email

Normalizes a single email address and returns the canonical form together with a breakdown of all transformations applied.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_normalize_post200_response import V1TextNormalizePost200Response
from openapi_client.models.v1_text_normalize_post_request import V1TextNormalizePostRequest
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
    api_instance = openapi_client.EmailNormalizeApi(api_client)
    v1_text_normalize_post_request = openapi_client.V1TextNormalizePostRequest() # V1TextNormalizePostRequest | 

    try:
        # Normalize Email
        api_response = api_instance.v1_text_normalize_post(v1_text_normalize_post_request)
        print("The response of EmailNormalizeApi->v1_text_normalize_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EmailNormalizeApi->v1_text_normalize_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_text_normalize_post_request** | [**V1TextNormalizePostRequest**](V1TextNormalizePostRequest.md)|  | 

### Return type

[**V1TextNormalizePost200Response**](V1TextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing, not valid JSON, or contains unknown fields. |  -  |
**422** | The email field is missing or not a valid email address format. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

