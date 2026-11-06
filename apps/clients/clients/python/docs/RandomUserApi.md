# openapi_client.RandomUserApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_random_user_get**](RandomUserApi.md#v1_technology_random_user_get) | **GET** /v1/technology/random-user | Get Random User


# **v1_technology_random_user_get**
> V1TechnologyRandomUserGet200Response v1_technology_random_user_get()

Get Random User

Returns a randomly generated fake user profile.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_random_user_get200_response import V1TechnologyRandomUserGet200Response
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
    api_instance = openapi_client.RandomUserApi(api_client)

    try:
        # Get Random User
        api_response = api_instance.v1_technology_random_user_get()
        print("The response of RandomUserApi->v1_technology_random_user_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RandomUserApi->v1_technology_random_user_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1TechnologyRandomUserGet200Response**](V1TechnologyRandomUserGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

