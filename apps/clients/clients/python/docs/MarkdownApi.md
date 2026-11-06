# openapi_client.MarkdownApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_technology_markdown_post**](MarkdownApi.md#v1_technology_markdown_post) | **POST** /v1/technology/markdown | Convert Markdown to HTML


# **v1_technology_markdown_post**
> V1TechnologyMarkdownPost200Response v1_technology_markdown_post(v1_technology_markdown_post_request)

Convert Markdown to HTML

Converts a Markdown string to HTML. Pass sanitize true to strip potentially unsafe tags like script and iframe from the output.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_technology_markdown_post200_response import V1TechnologyMarkdownPost200Response
from openapi_client.models.v1_technology_markdown_post_request import V1TechnologyMarkdownPostRequest
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
    api_instance = openapi_client.MarkdownApi(api_client)
    v1_technology_markdown_post_request = openapi_client.V1TechnologyMarkdownPostRequest() # V1TechnologyMarkdownPostRequest | 

    try:
        # Convert Markdown to HTML
        api_response = api_instance.v1_technology_markdown_post(v1_technology_markdown_post_request)
        print("The response of MarkdownApi->v1_technology_markdown_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MarkdownApi->v1_technology_markdown_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_technology_markdown_post_request** | [**V1TechnologyMarkdownPostRequest**](V1TechnologyMarkdownPostRequest.md)|  | 

### Return type

[**V1TechnologyMarkdownPost200Response**](V1TechnologyMarkdownPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**422** | unprocessable_entity |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

