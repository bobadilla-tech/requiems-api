# OpenapiClient::IpInfoApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_ip_get**](IpInfoApi.md#v1_networking_ip_get) | **GET** /v1/networking/ip | Get IP Info (Caller IP) |
| [**v1_networking_ip_ip_get**](IpInfoApi.md#v1_networking_ip_ip_get) | **GET** /v1/networking/ip/{ip} | Get IP Info for IP |


## v1_networking_ip_get

> <V1NetworkingIpGet200Response> v1_networking_ip_get

Get IP Info (Caller IP)

Get geolocation and network information for the requesting client's IP address. Useful when you want information about the user making the request without specifying an IP explicitly.

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

api_instance = OpenapiClient::IpInfoApi.new

begin
  # Get IP Info (Caller IP)
  result = api_instance.v1_networking_ip_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpInfoApi->v1_networking_ip_get: #{e}"
end
```

#### Using the v1_networking_ip_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingIpGet200Response>, Integer, Hash)> v1_networking_ip_get_with_http_info

```ruby
begin
  # Get IP Info (Caller IP)
  data, status_code, headers = api_instance.v1_networking_ip_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingIpGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpInfoApi->v1_networking_ip_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1NetworkingIpGet200Response**](V1NetworkingIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_networking_ip_ip_get

> <V1NetworkingIpIpGet200Response> v1_networking_ip_ip_get(ip)

Get IP Info for IP

Get geolocation and network information for a specific IP address.

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

api_instance = OpenapiClient::IpInfoApi.new
ip = '8.8.8.8' # String | The IP address to look up (supports IPv4 and IPv6)

begin
  # Get IP Info for IP
  result = api_instance.v1_networking_ip_ip_get(ip)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpInfoApi->v1_networking_ip_ip_get: #{e}"
end
```

#### Using the v1_networking_ip_ip_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingIpIpGet200Response>, Integer, Hash)> v1_networking_ip_ip_get_with_http_info(ip)

```ruby
begin
  # Get IP Info for IP
  data, status_code, headers = api_instance.v1_networking_ip_ip_get_with_http_info(ip)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingIpIpGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpInfoApi->v1_networking_ip_ip_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | The IP address to look up (supports IPv4 and IPv6) |  |

### Return type

[**V1NetworkingIpIpGet200Response**](V1NetworkingIpIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

