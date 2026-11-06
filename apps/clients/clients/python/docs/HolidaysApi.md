# openapi_client.HolidaysApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_holidays_batch_post**](HolidaysApi.md#v1_places_holidays_batch_post) | **POST** /v1/places/holidays/batch | Batch Get Holidays
[**v1_places_holidays_get**](HolidaysApi.md#v1_places_holidays_get) | **GET** /v1/places/holidays | Get Holidays


# **v1_places_holidays_batch_post**
> V1PlacesHolidaysBatchPost200Response v1_places_holidays_batch_post(v1_places_holidays_batch_post_request)

Batch Get Holidays

Returns holidays for up to 50 (country, year) pairs in a single request. Each pair is processed independently — if one combination has no data, it returns found:false without failing the entire batch.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_holidays_batch_post200_response import V1PlacesHolidaysBatchPost200Response
from openapi_client.models.v1_places_holidays_batch_post_request import V1PlacesHolidaysBatchPostRequest
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
    api_instance = openapi_client.HolidaysApi(api_client)
    v1_places_holidays_batch_post_request = openapi_client.V1PlacesHolidaysBatchPostRequest() # V1PlacesHolidaysBatchPostRequest | 

    try:
        # Batch Get Holidays
        api_response = api_instance.v1_places_holidays_batch_post(v1_places_holidays_batch_post_request)
        print("The response of HolidaysApi->v1_places_holidays_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling HolidaysApi->v1_places_holidays_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_places_holidays_batch_post_request** | [**V1PlacesHolidaysBatchPostRequest**](V1PlacesHolidaysBatchPostRequest.md)|  | 

### Return type

[**V1PlacesHolidaysBatchPost200Response**](V1PlacesHolidaysBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Malformed request body |  -  |
**422** | queries is missing, empty, exceeds 50 items, or contains invalid country codes or years |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_places_holidays_get**
> V1PlacesHolidaysGet200Response v1_places_holidays_get(country, year)

Get Holidays

Returns a list of public holidays for the specified country and year

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_holidays_get200_response import V1PlacesHolidaysGet200Response
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
    api_instance = openapi_client.HolidaysApi(api_client)
    country = 'US' # str | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"DE\")
    year = 2025 # int | Year for which to retrieve holidays (e.g., 2025)

    try:
        # Get Holidays
        api_response = api_instance.v1_places_holidays_get(country, year)
        print("The response of HolidaysApi->v1_places_holidays_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling HolidaysApi->v1_places_holidays_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **str**| ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) | 
 **year** | **int**| Year for which to retrieve holidays (e.g., 2025) | 

### Return type

[**V1PlacesHolidaysGet200Response**](V1PlacesHolidaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Missing or invalid country code or year parameter |  -  |
**404** | No holidays found for the specified country and year |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

