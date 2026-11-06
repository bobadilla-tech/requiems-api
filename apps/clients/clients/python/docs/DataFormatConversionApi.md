# openapi_client.DataFormatConversionApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_format_post**](DataFormatConversionApi.md#v1_technology_format_post) | **POST** /v1/technology/format | Convert Format


# **v1_technology_format_post**
> V1TechnologyFormatPost200Response v1_technology_format_post(v1_technology_format_post_request)

Convert Format

Convert content from one structured data format to another. Supported formats: json, yaml, csv, xml, toml.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_format_post200_response import V1TechnologyFormatPost200Response
from openapi_client.models.v1_technology_format_post_request import V1TechnologyFormatPostRequest
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
    api_instance = openapi_client.DataFormatConversionApi(api_client)
    v1_technology_format_post_request = openapi_client.V1TechnologyFormatPostRequest() # V1TechnologyFormatPostRequest | 

    try:
        # Convert Format
        api_response = api_instance.v1_technology_format_post(v1_technology_format_post_request)
        print("The response of DataFormatConversionApi->v1_technology_format_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DataFormatConversionApi->v1_technology_format_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_technology_format_post_request** | [**V1TechnologyFormatPostRequest**](V1TechnologyFormatPostRequest.md)|  | 

### Return type

[**V1TechnologyFormatPost200Response**](V1TechnologyFormatPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**413** | The content field exceeds the 512 KB limit. |  -  |
**422** | One of from, to, or content is missing, or from/to is not one of the supported format values.; The content field is not valid JSON (when from is json).; The content field is not valid YAML (when from is yaml).; The content field is not valid CSV, or a row has more columns than the header (when from is csv).; The content field is not valid XML (when from is xml).; The content field is not valid TOML (when from is toml).; The data structure is incompatible with the target format (e.g. converting a JSON array to TOML, which requires a top-level object). |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

