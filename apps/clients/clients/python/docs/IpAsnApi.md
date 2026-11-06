# openapi_client.IpAsnApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_networking_ip_asn_get**](IpAsnApi.md#v1_networking_ip_asn_get) | **GET** /v1/networking/ip/asn | Lookup ASN (Caller IP)
[**v1_networking_ip_asn_ip_get**](IpAsnApi.md#v1_networking_ip_asn_ip_get) | **GET** /v1/networking/ip/asn/{ip} | Lookup ASN for IP


# **v1_networking_ip_asn_get**
> V1NetworkingIpAsnGet200Response v1_networking_ip_asn_get()

Lookup ASN (Caller IP)

Look up ASN, organization, ISP, and network details for the requesting client's IP address. Useful when you want information about the user making the request without specifying an IP explicitly.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_ip_asn_get200_response import V1NetworkingIpAsnGet200Response
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
    api_instance = openapi_client.IpAsnApi(api_client)

    try:
        # Lookup ASN (Caller IP)
        api_response = api_instance.v1_networking_ip_asn_get()
        print("The response of IpAsnApi->v1_networking_ip_asn_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IpAsnApi->v1_networking_ip_asn_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1NetworkingIpAsnGet200Response**](V1NetworkingIpAsnGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_networking_ip_asn_ip_get**
> V1NetworkingIpAsnIpGet200Response v1_networking_ip_asn_ip_get(ip)

Lookup ASN for IP

Look up ASN, organization, ISP, and network details for a specific IP address.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_ip_asn_ip_get200_response import V1NetworkingIpAsnIpGet200Response
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
    api_instance = openapi_client.IpAsnApi(api_client)
    ip = '8.8.8.8' # str | The IP address to look up (supports IPv4 and IPv6)

    try:
        # Lookup ASN for IP
        api_response = api_instance.v1_networking_ip_asn_ip_get(ip)
        print("The response of IpAsnApi->v1_networking_ip_asn_ip_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IpAsnApi->v1_networking_ip_asn_ip_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **str**| The IP address to look up (supports IPv4 and IPv6) | 

### Return type

[**V1NetworkingIpAsnIpGet200Response**](V1NetworkingIpAsnIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The IP address is invalid |  -  |
**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

