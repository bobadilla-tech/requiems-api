# openapi_client.FactsApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_entertainment_facts_get**](FactsApi.md#v1_entertainment_facts_get) | **GET** /v1/entertainment/facts | Get Random Fact


# **v1_entertainment_facts_get**
> V1EntertainmentFactsGet200Response v1_entertainment_facts_get(category=category)

Get Random Fact

Returns a randomly selected fact, optionally filtered by category.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_facts_get200_response import V1EntertainmentFactsGet200Response
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
    api_instance = openapi_client.FactsApi(api_client)
    category = 'science' # str | Filter by category. Valid values: science, history, technology, nature, space, food (optional)

    try:
        # Get Random Fact
        api_response = api_instance.v1_entertainment_facts_get(category=category)
        print("The response of FactsApi->v1_entertainment_facts_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FactsApi->v1_entertainment_facts_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **str**| Filter by category. Valid values: science, history, technology, nature, space, food | [optional] 

### Return type

[**V1EntertainmentFactsGet200Response**](V1EntertainmentFactsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | bad_request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

