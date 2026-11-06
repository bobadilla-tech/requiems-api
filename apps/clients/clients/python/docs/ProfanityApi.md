# openapi_client.ProfanityApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_text_profanity_post**](ProfanityApi.md#v1_text_profanity_post) | **POST** /v1/text/profanity | Check Profanity


# **v1_text_profanity_post**
> V1TextProfanityPost200Response v1_text_profanity_post(v1_text_profanity_post_request)

Check Profanity

Checks text for profanity, returning a censored version and the list of flagged words.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_profanity_post200_response import V1TextProfanityPost200Response
from openapi_client.models.v1_text_profanity_post_request import V1TextProfanityPostRequest
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
    api_instance = openapi_client.ProfanityApi(api_client)
    v1_text_profanity_post_request = openapi_client.V1TextProfanityPostRequest() # V1TextProfanityPostRequest | 

    try:
        # Check Profanity
        api_response = api_instance.v1_text_profanity_post(v1_text_profanity_post_request)
        print("The response of ProfanityApi->v1_text_profanity_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProfanityApi->v1_text_profanity_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_text_profanity_post_request** | [**V1TextProfanityPostRequest**](V1TextProfanityPostRequest.md)|  | 

### Return type

[**V1TextProfanityPost200Response**](V1TextProfanityPost200Response.md)

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

