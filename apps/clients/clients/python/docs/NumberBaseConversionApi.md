# openapi_client.NumberBaseConversionApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_base_batch_post**](NumberBaseConversionApi.md#v1_technology_base_batch_post) | **POST** /v1/technology/base/batch | Batch Convert Base
[**v1_technology_base_get**](NumberBaseConversionApi.md#v1_technology_base_get) | **GET** /v1/technology/base | Convert Base


# **v1_technology_base_batch_post**
> V1TechnologyBaseBatchPost200Response v1_technology_base_batch_post(v1_technology_base_batch_post_request)

Batch Convert Base

Convert up to 50 numbers between bases in a single request. Results are returned in input order. Per-item errors are reported inline.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_base_batch_post200_response import V1TechnologyBaseBatchPost200Response
from openapi_client.models.v1_technology_base_batch_post_request import V1TechnologyBaseBatchPostRequest
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
    api_instance = openapi_client.NumberBaseConversionApi(api_client)
    v1_technology_base_batch_post_request = openapi_client.V1TechnologyBaseBatchPostRequest() # V1TechnologyBaseBatchPostRequest | 

    try:
        # Batch Convert Base
        api_response = api_instance.v1_technology_base_batch_post(v1_technology_base_batch_post_request)
        print("The response of NumberBaseConversionApi->v1_technology_base_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NumberBaseConversionApi->v1_technology_base_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_technology_base_batch_post_request** | [**V1TechnologyBaseBatchPostRequest**](V1TechnologyBaseBatchPostRequest.md)|  | 

### Return type

[**V1TechnologyBaseBatchPost200Response**](V1TechnologyBaseBatchPost200Response.md)

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
**422** | The items array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_technology_base_get**
> V1TechnologyBaseGet200Response v1_technology_base_get(var_from, to, value)

Convert Base

Convert an integer from one number base to another.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_base_get200_response import V1TechnologyBaseGet200Response
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
    api_instance = openapi_client.NumberBaseConversionApi(api_client)
    var_from = 10 # int | Source base (2, 8, 10, or 16)
    to = 16 # int | Target base (2, 8, 10, or 16)
    value = '255' # str | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal).

    try:
        # Convert Base
        api_response = api_instance.v1_technology_base_get(var_from, to, value)
        print("The response of NumberBaseConversionApi->v1_technology_base_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NumberBaseConversionApi->v1_technology_base_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **var_from** | **int**| Source base (2, 8, 10, or 16) | 
 **to** | **int**| Target base (2, 8, 10, or 16) | 
 **value** | **str**| The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). | 

### Return type

[**V1TechnologyBaseGet200Response**](V1TechnologyBaseGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | A required parameter is missing, the base is not one of 2/8/10/16, or value is not valid for the given base. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

