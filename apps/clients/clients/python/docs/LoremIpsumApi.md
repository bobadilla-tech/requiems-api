# openapi_client.LoremIpsumApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_text_lorem_get**](LoremIpsumApi.md#v1_text_lorem_get) | **GET** /v1/text/lorem | Generate Lorem Ipsum


# **v1_text_lorem_get**
> V1TextLoremGet200Response v1_text_lorem_get(paragraphs=paragraphs, sentences=sentences)

Generate Lorem Ipsum

Generate Lorem Ipsum placeholder text with customizable length and format

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_lorem_get200_response import V1TextLoremGet200Response
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
    api_instance = openapi_client.LoremIpsumApi(api_client)
    paragraphs = 3 # int | Number of paragraphs to generate (1-20) (optional)
    sentences = 5 # int | Number of sentences per paragraph (1-20) (optional)

    try:
        # Generate Lorem Ipsum
        api_response = api_instance.v1_text_lorem_get(paragraphs=paragraphs, sentences=sentences)
        print("The response of LoremIpsumApi->v1_text_lorem_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LoremIpsumApi->v1_text_lorem_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paragraphs** | **int**| Number of paragraphs to generate (1-20) | [optional] 
 **sentences** | **int**| Number of sentences per paragraph (1-20) | [optional] 

### Return type

[**V1TextLoremGet200Response**](V1TextLoremGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The paragraphs parameter is out of valid range; The sentences parameter is out of valid range |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

