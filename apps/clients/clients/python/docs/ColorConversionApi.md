# openapi_client.ColorConversionApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_color_get**](ColorConversionApi.md#v1_technology_color_get) | **GET** /v1/technology/color | Convert Color


# **v1_technology_color_get**
> V1TechnologyColorGet200Response v1_technology_color_get(var_from, to, value)

Convert Color

Convert a color value from one format to another. The response always includes all four formats.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_color_get200_response import V1TechnologyColorGet200Response
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
    api_instance = openapi_client.ColorConversionApi(api_client)
    var_from = 'hex' # str | Source color format: hex, rgb, hsl, or cmyk
    to = 'hsl' # str | Target color format: hex, rgb, hsl, or cmyk
    value = '#ff5733' # str | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%))

    try:
        # Convert Color
        api_response = api_instance.v1_technology_color_get(var_from, to, value)
        print("The response of ColorConversionApi->v1_technology_color_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ColorConversionApi->v1_technology_color_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **var_from** | **str**| Source color format: hex, rgb, hsl, or cmyk | 
 **to** | **str**| Target color format: hex, rgb, hsl, or cmyk | 
 **value** | **str**| Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%)) | 

### Return type

[**V1TechnologyColorGet200Response**](V1TechnologyColorGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | One or more of from, to, or value parameters is missing or the from/to value is not one of: hex, rgb, hsl, cmyk. |  -  |
**422** | The value cannot be parsed in the specified from format. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

