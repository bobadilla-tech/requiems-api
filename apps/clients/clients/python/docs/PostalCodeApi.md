# openapi_client.PostalCodeApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_postal_code_get**](PostalCodeApi.md#v1_places_postal_code_get) | **GET** /v1/places/postal/{code} | Lookup Postal Code


# **v1_places_postal_code_get**
> V1PlacesPostalCodeGet200Response v1_places_postal_code_get(code, country=country)

Lookup Postal Code

Returns city, state, country, and coordinates for the given postal code.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_postal_code_get200_response import V1PlacesPostalCodeGet200Response
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
    api_instance = openapi_client.PostalCodeApi(api_client)
    code = '10001' # str | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London)
    country = 'US' # str | ISO 3166-1 alpha-2 country code (default: US) (optional)

    try:
        # Lookup Postal Code
        api_response = api_instance.v1_places_postal_code_get(code, country=country)
        print("The response of PostalCodeApi->v1_places_postal_code_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PostalCodeApi->v1_places_postal_code_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London) | 
 **country** | **str**| ISO 3166-1 alpha-2 country code (default: US) | [optional] 

### Return type

[**V1PlacesPostalCodeGet200Response**](V1PlacesPostalCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**404** | The postal code was not found for the given country. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

