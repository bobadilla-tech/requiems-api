# openapi_client.CitiesApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_cities_city_get**](CitiesApi.md#v1_places_cities_city_get) | **GET** /v1/places/cities/{city} | Get City Info


# **v1_places_cities_city_get**
> V1PlacesCitiesCityGet200Response v1_places_cities_city_get(city)

Get City Info

Returns metadata for a city by name. Lookup is case-insensitive.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_cities_city_get200_response import V1PlacesCitiesCityGet200Response
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
    api_instance = openapi_client.CitiesApi(api_client)
    city = 'london' # str | City name to look up (e.g. london, tokyo, new york city)

    try:
        # Get City Info
        api_response = api_instance.v1_places_cities_city_get(city)
        print("The response of CitiesApi->v1_places_cities_city_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CitiesApi->v1_places_cities_city_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **city** | **str**| City name to look up (e.g. london, tokyo, new york city) | 

### Return type

[**V1PlacesCitiesCityGet200Response**](V1PlacesCitiesCityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**404** | No city with that name was found in the dataset. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

