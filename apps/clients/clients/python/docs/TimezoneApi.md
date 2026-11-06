# openapi_client.TimezoneApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_timezone_get**](TimezoneApi.md#v1_places_timezone_get) | **GET** /v1/places/timezone | Get Timezone


# **v1_places_timezone_get**
> V1PlacesTimezoneGet200Response v1_places_timezone_get(lat=lat, lon=lon, city=city)

Get Timezone

Returns timezone information for the given coordinates or city name. Provide either `city` or both `lat` and `lon`.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_timezone_get200_response import V1PlacesTimezoneGet200Response
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
    api_instance = openapi_client.TimezoneApi(api_client)
    lat = 'lat_example' # str | Latitude of the location (-90 to 90). Required when using coordinate-based lookup. (optional)
    lon = 'lon_example' # str | Longitude of the location (-180 to 180). Required when using coordinate-based lookup. (optional)
    city = 'city_example' # str | City name for city-based lookup (e.g. 'Tokyo', 'London'). Required when not using coordinates. (optional)

    try:
        # Get Timezone
        api_response = api_instance.v1_places_timezone_get(lat=lat, lon=lon, city=city)
        print("The response of TimezoneApi->v1_places_timezone_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TimezoneApi->v1_places_timezone_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lat** | **str**| Latitude of the location (-90 to 90). Required when using coordinate-based lookup. | [optional] 
 **lon** | **str**| Longitude of the location (-180 to 180). Required when using coordinate-based lookup. | [optional] 
 **city** | **str**| City name for city-based lookup (e.g. &#39;Tokyo&#39;, &#39;London&#39;). Required when not using coordinates. | [optional] 

### Return type

[**V1PlacesTimezoneGet200Response**](V1PlacesTimezoneGet200Response.md)

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

