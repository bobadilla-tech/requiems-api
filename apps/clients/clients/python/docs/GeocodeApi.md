# openapi_client.GeocodeApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_geocode_get**](GeocodeApi.md#v1_places_geocode_get) | **GET** /v1/places/geocode | Geocode Address
[**v1_places_reverse_geocode_get**](GeocodeApi.md#v1_places_reverse_geocode_get) | **GET** /v1/places/reverse-geocode | Reverse Geocode


# **v1_places_geocode_get**
> V1PlacesGeocodeGet200Response v1_places_geocode_get(address)

Geocode Address

Converts a free-text address into latitude and longitude coordinates.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_geocode_get200_response import V1PlacesGeocodeGet200Response
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
    api_instance = openapi_client.GeocodeApi(api_client)
    address = '1600 Pennsylvania Ave NW, Washington DC' # str | The address to geocode (street, city, country, or any combination)

    try:
        # Geocode Address
        api_response = api_instance.v1_places_geocode_get(address)
        print("The response of GeocodeApi->v1_places_geocode_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GeocodeApi->v1_places_geocode_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **str**| The address to geocode (street, city, country, or any combination) | 

### Return type

[**V1PlacesGeocodeGet200Response**](V1PlacesGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The address parameter is missing. |  -  |
**404** | No results found for the given address. |  -  |
**503** | The geocoding service is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_places_reverse_geocode_get**
> V1PlacesReverseGeocodeGet200Response v1_places_reverse_geocode_get(lat, lon)

Reverse Geocode

Converts geographic coordinates into a human-readable address.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_reverse_geocode_get200_response import V1PlacesReverseGeocodeGet200Response
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
    api_instance = openapi_client.GeocodeApi(api_client)
    lat = 38.8977 # float | Latitude of the location (-90 to 90)
    lon = -77.0365 # float | Longitude of the location (-180 to 180)

    try:
        # Reverse Geocode
        api_response = api_instance.v1_places_reverse_geocode_get(lat, lon)
        print("The response of GeocodeApi->v1_places_reverse_geocode_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GeocodeApi->v1_places_reverse_geocode_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lat** | **float**| Latitude of the location (-90 to 90) | 
 **lon** | **float**| Longitude of the location (-180 to 180) | 

### Return type

[**V1PlacesReverseGeocodeGet200Response**](V1PlacesReverseGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | lat or lon is missing or out of range (lat: -90..90, lon: -180..180). |  -  |
**404** | No address found for the given coordinates. |  -  |
**503** | The geocoding service is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

