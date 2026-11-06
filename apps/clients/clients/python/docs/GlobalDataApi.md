# openapi_client.GlobalDataApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_systems_business_calendar_country_get**](GlobalDataApi.md#v1_systems_business_calendar_country_get) | **GET** /v1/systems/business-calendar/{country} | Business Calendar
[**v1_systems_location_resolve_post**](GlobalDataApi.md#v1_systems_location_resolve_post) | **POST** /v1/systems/location/resolve | Resolve Location
[**v1_systems_timezone_from_ip_ip_get**](GlobalDataApi.md#v1_systems_timezone_from_ip_ip_get) | **GET** /v1/systems/timezone/from-ip/{ip} | Timezone from IP


# **v1_systems_business_calendar_country_get**
> V1SystemsBusinessCalendarCountryGet200Response v1_systems_business_calendar_country_get(country, year=year, month=month)

Business Calendar

Retrieve working days, public holidays, and next upcoming holiday for any country and optional year/month. Covers 100+ countries.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_business_calendar_country_get200_response import V1SystemsBusinessCalendarCountryGet200Response
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
    api_instance = openapi_client.GlobalDataApi(api_client)
    country = 'DE' # str | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR)
    year = 2026 # int | Year to retrieve the calendar for. Defaults to the current year. (optional)
    month = 6 # int | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. (optional)

    try:
        # Business Calendar
        api_response = api_instance.v1_systems_business_calendar_country_get(country, year=year, month=month)
        print("The response of GlobalDataApi->v1_systems_business_calendar_country_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GlobalDataApi->v1_systems_business_calendar_country_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **str**| ISO 3166-1 alpha-2 country code (e.g. DE, US, FR) | 
 **year** | **int**| Year to retrieve the calendar for. Defaults to the current year. | [optional] 
 **month** | **int**| Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. | [optional] 

### Return type

[**V1SystemsBusinessCalendarCountryGet200Response**](V1SystemsBusinessCalendarCountryGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | Country is not a 2-letter ISO code, or year/month is out of the valid range |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_systems_location_resolve_post**
> V1SystemsLocationResolvePost200Response v1_systems_location_resolve_post(v1_systems_location_resolve_post_request)

Resolve Location

Resolve an address or coordinates into a full location profile — city, country, timezone, UTC offset, current time, working days this month, and next holiday.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_location_resolve_post200_response import V1SystemsLocationResolvePost200Response
from openapi_client.models.v1_systems_location_resolve_post_request import V1SystemsLocationResolvePostRequest
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
    api_instance = openapi_client.GlobalDataApi(api_client)
    v1_systems_location_resolve_post_request = openapi_client.V1SystemsLocationResolvePostRequest() # V1SystemsLocationResolvePostRequest | 

    try:
        # Resolve Location
        api_response = api_instance.v1_systems_location_resolve_post(v1_systems_location_resolve_post_request)
        print("The response of GlobalDataApi->v1_systems_location_resolve_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GlobalDataApi->v1_systems_location_resolve_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_location_resolve_post_request** | [**V1SystemsLocationResolvePostRequest**](V1SystemsLocationResolvePostRequest.md)|  | 

### Return type

[**V1SystemsLocationResolvePost200Response**](V1SystemsLocationResolvePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | Neither address nor coordinates was provided; Geocoding service was unavailable |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_systems_timezone_from_ip_ip_get**
> V1SystemsTimezoneFromIpIpGet200Response v1_systems_timezone_from_ip_ip_get(ip)

Timezone from IP

Resolve the timezone, UTC offset, and DST status for any IPv4 or IPv6 address without needing an address string.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_timezone_from_ip_ip_get200_response import V1SystemsTimezoneFromIpIpGet200Response
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
    api_instance = openapi_client.GlobalDataApi(api_client)
    ip = '92.168.1.1' # str | IPv4 or IPv6 address to resolve

    try:
        # Timezone from IP
        api_response = api_instance.v1_systems_timezone_from_ip_ip_get(ip)
        print("The response of GlobalDataApi->v1_systems_timezone_from_ip_ip_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GlobalDataApi->v1_systems_timezone_from_ip_ip_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **str**| IPv4 or IPv6 address to resolve | 

### Return type

[**V1SystemsTimezoneFromIpIpGet200Response**](V1SystemsTimezoneFromIpIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | The IP address could not be geolocated |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

