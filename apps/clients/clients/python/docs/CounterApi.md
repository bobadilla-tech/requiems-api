# openapi_client.CounterApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_counter_namespace_get**](CounterApi.md#v1_technology_counter_namespace_get) | **GET** /v1/technology/counter/{namespace} | Get Counter Value
[**v1_technology_counter_namespace_post**](CounterApi.md#v1_technology_counter_namespace_post) | **POST** /v1/technology/counter/{namespace} | Increment Counter


# **v1_technology_counter_namespace_get**
> V1TechnologyCounterNamespaceGet200Response v1_technology_counter_namespace_get(namespace)

Get Counter Value

Get the current value of a counter without incrementing it

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_counter_namespace_get200_response import V1TechnologyCounterNamespaceGet200Response
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
    api_instance = openapi_client.CounterApi(api_client)
    namespace = 'page-views' # str | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

    try:
        # Get Counter Value
        api_response = api_instance.v1_technology_counter_namespace_get(namespace)
        print("The response of CounterApi->v1_technology_counter_namespace_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CounterApi->v1_technology_counter_namespace_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) | 

### Return type

[**V1TechnologyCounterNamespaceGet200Response**](V1TechnologyCounterNamespaceGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_technology_counter_namespace_post**
> V1TechnologyCounterNamespacePost200Response v1_technology_counter_namespace_post(namespace)

Increment Counter

Atomically increment a counter in the specified namespace and return the new value

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_counter_namespace_post200_response import V1TechnologyCounterNamespacePost200Response
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
    api_instance = openapi_client.CounterApi(api_client)
    namespace = 'page-views' # str | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

    try:
        # Increment Counter
        api_response = api_instance.v1_technology_counter_namespace_post(namespace)
        print("The response of CounterApi->v1_technology_counter_namespace_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CounterApi->v1_technology_counter_namespace_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) | 

### Return type

[**V1TechnologyCounterNamespacePost200Response**](V1TechnologyCounterNamespacePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

