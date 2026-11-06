# openapi_client.VpnDetectionApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_networking_ip_vpn_ip_get**](VpnDetectionApi.md#v1_networking_ip_vpn_ip_get) | **GET** /v1/networking/ip/vpn/{ip} | Check IP Address


# **v1_networking_ip_vpn_ip_get**
> V1NetworkingIpVpnIpGet200Response v1_networking_ip_vpn_ip_get(ip)

Check IP Address

Analyze an IP address to determine if it's a VPN, proxy, Tor exit node, or hosting provider. Returns detailed threat indicators and scores.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_ip_vpn_ip_get200_response import V1NetworkingIpVpnIpGet200Response
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
    api_instance = openapi_client.VpnDetectionApi(api_client)
    ip = '8.8.8.8' # str | The IP address to check (supports IPv4 and IPv6)

    try:
        # Check IP Address
        api_response = api_instance.v1_networking_ip_vpn_ip_get(ip)
        print("The response of VpnDetectionApi->v1_networking_ip_vpn_ip_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VpnDetectionApi->v1_networking_ip_vpn_ip_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **str**| The IP address to check (supports IPv4 and IPv6) | 

### Return type

[**V1NetworkingIpVpnIpGet200Response**](V1NetworkingIpVpnIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The IP address is missing or invalid |  -  |
**500** | Unexpected server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

