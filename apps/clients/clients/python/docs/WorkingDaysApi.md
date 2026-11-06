# openapi_client.WorkingDaysApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_places_working_days_get**](WorkingDaysApi.md#v1_places_working_days_get) | **GET** /v1/places/working-days | Calculate Working Days


# **v1_places_working_days_get**
> V1PlacesWorkingDaysGet200Response v1_places_working_days_get(var_from, to, country=country, subdivision=subdivision)

Calculate Working Days

Calculate the number of working days between two dates, optionally accounting for country-specific holidays

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_places_working_days_get200_response import V1PlacesWorkingDaysGet200Response
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
    api_instance = openapi_client.WorkingDaysApi(api_client)
    var_from = '2024-02-23' # str | Start date in YYYY-MM-DD format (ISO 8601)
    to = '2024-02-28' # str | End date in YYYY-MM-DD format (ISO 8601). Must be >= from date.
    country = 'US' # str | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"FR\"). When provided, country-specific holidays are excluded from working days count. (optional)
    subdivision = 'NY' # str | ISO 3166-2 subdivision code for state/region within the country (e.g., \"NY\" for New York, \"CA\" for California). Only used when country is provided. (optional)

    try:
        # Calculate Working Days
        api_response = api_instance.v1_places_working_days_get(var_from, to, country=country, subdivision=subdivision)
        print("The response of WorkingDaysApi->v1_places_working_days_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WorkingDaysApi->v1_places_working_days_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **var_from** | **str**| Start date in YYYY-MM-DD format (ISO 8601) | 
 **to** | **str**| End date in YYYY-MM-DD format (ISO 8601). Must be &gt;&#x3D; from date. | 
 **country** | **str**| ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;FR\&quot;). When provided, country-specific holidays are excluded from working days count. | [optional] 
 **subdivision** | **str**| ISO 3166-2 subdivision code for state/region within the country (e.g., \&quot;NY\&quot; for New York, \&quot;CA\&quot; for California). Only used when country is provided. | [optional] 

### Return type

[**V1PlacesWorkingDaysGet200Response**](V1PlacesWorkingDaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The from and to parameters are required, or to date is before from date, or invalid date format |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

