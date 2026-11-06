# openapi_client.MxLookupApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_networking_mx_domain_get**](MxLookupApi.md#v1_networking_mx_domain_get) | **GET** /v1/networking/mx/{domain} | MX Lookup


# **v1_networking_mx_domain_get**
> V1NetworkingMxDomainGet200Response v1_networking_mx_domain_get(domain)

MX Lookup

Retrieve all MX records for a domain. Results are sorted by priority ascending (lowest numeric value has highest mail delivery priority per RFC 5321).

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_mx_domain_get200_response import V1NetworkingMxDomainGet200Response
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
    api_instance = openapi_client.MxLookupApi(api_client)
    domain = 'gmail.com' # str | The domain name to look up MX records for (e.g. gmail.com)

    try:
        # MX Lookup
        api_response = api_instance.v1_networking_mx_domain_get(domain)
        print("The response of MxLookupApi->v1_networking_mx_domain_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MxLookupApi->v1_networking_mx_domain_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **str**| The domain name to look up MX records for (e.g. gmail.com) | 

### Return type

[**V1NetworkingMxDomainGet200Response**](V1NetworkingMxDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The domain parameter is not a valid domain name. |  -  |
**404** | No MX records were found for the domain (domain may not accept email). |  -  |
**500** | DNS lookup failed due to an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

