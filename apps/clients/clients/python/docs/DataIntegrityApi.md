# openapi_client.DataIntegrityApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_systems_content_moderate_post**](DataIntegrityApi.md#v1_systems_content_moderate_post) | **POST** /v1/systems/content/moderate | Moderate Content
[**v1_systems_text_normalize_post**](DataIntegrityApi.md#v1_systems_text_normalize_post) | **POST** /v1/systems/text/normalize | Normalize Text


# **v1_systems_content_moderate_post**
> V1SystemsContentModeratePost200Response v1_systems_content_moderate_post(v1_systems_content_moderate_post_request)

Moderate Content

Check text for profanity, toxicity, sentiment, and language. Returns per-category flags and an is_safe boolean for gating content.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_content_moderate_post200_response import V1SystemsContentModeratePost200Response
from openapi_client.models.v1_systems_content_moderate_post_request import V1SystemsContentModeratePostRequest
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
    api_instance = openapi_client.DataIntegrityApi(api_client)
    v1_systems_content_moderate_post_request = openapi_client.V1SystemsContentModeratePostRequest() # V1SystemsContentModeratePostRequest | 

    try:
        # Moderate Content
        api_response = api_instance.v1_systems_content_moderate_post(v1_systems_content_moderate_post_request)
        print("The response of DataIntegrityApi->v1_systems_content_moderate_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DataIntegrityApi->v1_systems_content_moderate_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_content_moderate_post_request** | [**V1SystemsContentModeratePostRequest**](V1SystemsContentModeratePostRequest.md)|  | 

### Return type

[**V1SystemsContentModeratePost200Response**](V1SystemsContentModeratePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | text field is missing or empty |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_systems_text_normalize_post**
> V1SystemsTextNormalizePost200Response v1_systems_text_normalize_post(v1_systems_text_normalize_post_request)

Normalize Text

Apply a composable pipeline of normalization operations to a string. Trim whitespace, fix encoding, normalize unicode, strip HTML, and more — in a single call.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_text_normalize_post200_response import V1SystemsTextNormalizePost200Response
from openapi_client.models.v1_systems_text_normalize_post_request import V1SystemsTextNormalizePostRequest
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
    api_instance = openapi_client.DataIntegrityApi(api_client)
    v1_systems_text_normalize_post_request = openapi_client.V1SystemsTextNormalizePostRequest() # V1SystemsTextNormalizePostRequest | 

    try:
        # Normalize Text
        api_response = api_instance.v1_systems_text_normalize_post(v1_systems_text_normalize_post_request)
        print("The response of DataIntegrityApi->v1_systems_text_normalize_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DataIntegrityApi->v1_systems_text_normalize_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_text_normalize_post_request** | [**V1SystemsTextNormalizePostRequest**](V1SystemsTextNormalizePostRequest.md)|  | 

### Return type

[**V1SystemsTextNormalizePost200Response**](V1SystemsTextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | text is empty or operations array is empty or missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

