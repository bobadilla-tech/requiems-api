# openapi_client.SpellCheckApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_text_spellcheck_batch_post**](SpellCheckApi.md#v1_text_spellcheck_batch_post) | **POST** /v1/text/spellcheck/batch | Check Spelling (Batch)
[**v1_text_spellcheck_post**](SpellCheckApi.md#v1_text_spellcheck_post) | **POST** /v1/text/spellcheck | Check Spelling


# **v1_text_spellcheck_batch_post**
> V1TextSpellcheckBatchPost200Response v1_text_spellcheck_batch_post(v1_text_spellcheck_batch_post_request)

Check Spelling (Batch)

Checks multiple texts for spelling mistakes in a single request. Returns a corrected version and per-word corrections for each input text. Results are returned in the same order as the input array.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_spellcheck_batch_post200_response import V1TextSpellcheckBatchPost200Response
from openapi_client.models.v1_text_spellcheck_batch_post_request import V1TextSpellcheckBatchPostRequest
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
    api_instance = openapi_client.SpellCheckApi(api_client)
    v1_text_spellcheck_batch_post_request = openapi_client.V1TextSpellcheckBatchPostRequest() # V1TextSpellcheckBatchPostRequest | 

    try:
        # Check Spelling (Batch)
        api_response = api_instance.v1_text_spellcheck_batch_post(v1_text_spellcheck_batch_post_request)
        print("The response of SpellCheckApi->v1_text_spellcheck_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SpellCheckApi->v1_text_spellcheck_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_text_spellcheck_batch_post_request** | [**V1TextSpellcheckBatchPostRequest**](V1TextSpellcheckBatchPostRequest.md)|  | 

### Return type

[**V1TextSpellcheckBatchPost200Response**](V1TextSpellcheckBatchPost200Response.md)

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
**422** | The texts field is missing, empty, or exceeds 50 items. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_text_spellcheck_post**
> V1TextSpellcheckPost200Response v1_text_spellcheck_post(v1_text_spellcheck_post_request)

Check Spelling

Checks the input text for spelling mistakes and returns a corrected version along with per-word corrections.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_spellcheck_post200_response import V1TextSpellcheckPost200Response
from openapi_client.models.v1_text_spellcheck_post_request import V1TextSpellcheckPostRequest
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
    api_instance = openapi_client.SpellCheckApi(api_client)
    v1_text_spellcheck_post_request = openapi_client.V1TextSpellcheckPostRequest() # V1TextSpellcheckPostRequest | 

    try:
        # Check Spelling
        api_response = api_instance.v1_text_spellcheck_post(v1_text_spellcheck_post_request)
        print("The response of SpellCheckApi->v1_text_spellcheck_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SpellCheckApi->v1_text_spellcheck_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_text_spellcheck_post_request** | [**V1TextSpellcheckPostRequest**](V1TextSpellcheckPostRequest.md)|  | 

### Return type

[**V1TextSpellcheckPost200Response**](V1TextSpellcheckPost200Response.md)

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

