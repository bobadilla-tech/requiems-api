# openapi_client.HoroscopeApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_entertainment_horoscope_sign_get**](HoroscopeApi.md#v1_entertainment_horoscope_sign_get) | **GET** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope


# **v1_entertainment_horoscope_sign_get**
> V1EntertainmentHoroscopeSignGet200Response v1_entertainment_horoscope_sign_get(sign)

Get Daily Horoscope

Returns a daily horoscope reading for the specified zodiac sign.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_horoscope_sign_get200_response import V1EntertainmentHoroscopeSignGet200Response
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
    api_instance = openapi_client.HoroscopeApi(api_client)
    sign = 'sign_example' # str | Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces

    try:
        # Get Daily Horoscope
        api_response = api_instance.v1_entertainment_horoscope_sign_get(sign)
        print("The response of HoroscopeApi->v1_entertainment_horoscope_sign_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling HoroscopeApi->v1_entertainment_horoscope_sign_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sign** | **str**| Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces | 

### Return type

[**V1EntertainmentHoroscopeSignGet200Response**](V1EntertainmentHoroscopeSignGet200Response.md)

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

