# openapi_client.PostalCodeApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_postal_batch_post**](PostalCodeApi.md#v1_places_postal_batch_post) | **POST** /v1/places/postal/batch | Batch Lookup Postal Codes
[**v1_places_postal_code_get**](PostalCodeApi.md#v1_places_postal_code_get) | **GET** /v1/places/postal/{code} | Lookup Postal Code


# **v1_places_postal_batch_post**
> V1PlacesPostalBatchPost200Response v1_places_postal_batch_post(v1_places_postal_batch_post_request)

Batch Lookup Postal Codes

Look up city, state, and coordinates for up to 50 postal codes in a single request. Results are returned in input order.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_postal_batch_post200_response import V1PlacesPostalBatchPost200Response
from openapi_client.models.v1_places_postal_batch_post_request import V1PlacesPostalBatchPostRequest
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
    api_instance = openapi_client.PostalCodeApi(api_client)
    v1_places_postal_batch_post_request = openapi_client.V1PlacesPostalBatchPostRequest() # V1PlacesPostalBatchPostRequest | 

    try:
        # Batch Lookup Postal Codes
        api_response = api_instance.v1_places_postal_batch_post(v1_places_postal_batch_post_request)
        print("The response of PostalCodeApi->v1_places_postal_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PostalCodeApi->v1_places_postal_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_places_postal_batch_post_request** | [**V1PlacesPostalBatchPostRequest**](V1PlacesPostalBatchPostRequest.md)|  | 

### Return type

[**V1PlacesPostalBatchPost200Response**](V1PlacesPostalBatchPost200Response.md)

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

