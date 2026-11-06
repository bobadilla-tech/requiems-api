# openapi_client.UseragentApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_useragent_get**](UseragentApi.md#v1_technology_useragent_get) | **GET** /v1/technology/useragent | Parse User Agent


# **v1_technology_useragent_get**
> V1TechnologyUseragentGet200Response v1_technology_useragent_get(ua)

Parse User Agent

Parses a user agent string and returns structured information about the browser, OS, device, and bot status.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_useragent_get200_response import V1TechnologyUseragentGet200Response
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
    api_instance = openapi_client.UseragentApi(api_client)
    ua = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36' # str | The user agent string to parse.

    try:
        # Parse User Agent
        api_response = api_instance.v1_technology_useragent_get(ua)
        print("The response of UseragentApi->v1_technology_useragent_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UseragentApi->v1_technology_useragent_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ua** | **str**| The user agent string to parse. | 

### Return type

[**V1TechnologyUseragentGet200Response**](V1TechnologyUseragentGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The ua query parameter is missing. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

