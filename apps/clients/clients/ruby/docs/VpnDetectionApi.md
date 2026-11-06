# OpenapiClient::VpnDetectionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_ip_vpn_ip_get**](VpnDetectionApi.md#v1_networking_ip_vpn_ip_get) | **GET** /v1/networking/ip/vpn/{ip} | Check IP Address |


## v1_networking_ip_vpn_ip_get

> <V1NetworkingIpVpnIpGet200Response> v1_networking_ip_vpn_ip_get(ip)

Check IP Address

Analyze an IP address to determine if it's a VPN, proxy, Tor exit node, or hosting provider. Returns detailed threat indicators and scores.

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: requiems-api-key
  config.api_key['requiems-api-key'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['requiems-api-key'] = 'Bearer'
end

api_instance = OpenapiClient::VpnDetectionApi.new
ip = '8.8.8.8' # String | The IP address to check (supports IPv4 and IPv6)

begin
  # Check IP Address
  result = api_instance.v1_networking_ip_vpn_ip_get(ip)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling VpnDetectionApi->v1_networking_ip_vpn_ip_get: #{e}"
end
```

#### Using the v1_networking_ip_vpn_ip_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingIpVpnIpGet200Response>, Integer, Hash)> v1_networking_ip_vpn_ip_get_with_http_info(ip)

```ruby
begin
  # Check IP Address
  data, status_code, headers = api_instance.v1_networking_ip_vpn_ip_get_with_http_info(ip)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingIpVpnIpGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling VpnDetectionApi->v1_networking_ip_vpn_ip_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | The IP address to check (supports IPv4 and IPv6) |  |

### Return type

[**V1NetworkingIpVpnIpGet200Response**](V1NetworkingIpVpnIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

