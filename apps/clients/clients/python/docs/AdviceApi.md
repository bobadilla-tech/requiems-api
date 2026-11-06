# openapi_client.AdviceApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_text_advice_get**](AdviceApi.md#v1_text_advice_get) | **GET** /v1/text/advice | Get Random Advice


# **v1_text_advice_get**
> V1TextAdviceGet200Response v1_text_advice_get()

Get Random Advice

Returns a random piece of advice

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_advice_get200_response import V1TextAdviceGet200Response
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
    api_instance = openapi_client.AdviceApi(api_client)

    try:
        # Get Random Advice
        api_response = api_instance.v1_text_advice_get()
        print("The response of AdviceApi->v1_text_advice_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AdviceApi->v1_text_advice_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1TextAdviceGet200Response**](V1TextAdviceGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

