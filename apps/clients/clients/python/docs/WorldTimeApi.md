# openapi_client.WorldTimeApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_time_timezone_get**](WorldTimeApi.md#v1_places_time_timezone_get) | **GET** /v1/places/time/{timezone} | Get Current Time by Timezone


# **v1_places_time_timezone_get**
> V1PlacesTimeTimezoneGet200Response v1_places_time_timezone_get(timezone)

Get Current Time by Timezone

Returns the current time for the given IANA timezone identifier. The timezone is supplied as a path parameter (e.g. `America/New_York`, `Europe/London`, `UTC`).

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_time_timezone_get200_response import V1PlacesTimeTimezoneGet200Response
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
    api_instance = openapi_client.WorldTimeApi(api_client)
    timezone = 'America/New_York' # str | IANA timezone identifier (e.g. 'America/New_York', 'Europe/London', 'Asia/Kolkata')

    try:
        # Get Current Time by Timezone
        api_response = api_instance.v1_places_time_timezone_get(timezone)
        print("The response of WorldTimeApi->v1_places_time_timezone_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WorldTimeApi->v1_places_time_timezone_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timezone** | **str**| IANA timezone identifier (e.g. &#39;America/New_York&#39;, &#39;Europe/London&#39;, &#39;Asia/Kolkata&#39;) | 

### Return type

[**V1PlacesTimeTimezoneGet200Response**](V1PlacesTimeTimezoneGet200Response.md)

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

