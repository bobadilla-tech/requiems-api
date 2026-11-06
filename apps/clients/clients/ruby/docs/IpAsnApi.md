# OpenapiClient::IpAsnApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_ip_asn_get**](IpAsnApi.md#v1_networking_ip_asn_get) | **GET** /v1/networking/ip/asn | Lookup ASN (Caller IP) |
| [**v1_networking_ip_asn_ip_get**](IpAsnApi.md#v1_networking_ip_asn_ip_get) | **GET** /v1/networking/ip/asn/{ip} | Lookup ASN for IP |


## v1_networking_ip_asn_get

> <V1NetworkingIpAsnGet200Response> v1_networking_ip_asn_get

Lookup ASN (Caller IP)

Look up ASN, organization, ISP, and network details for the requesting client's IP address. Useful when you want information about the user making the request without specifying an IP explicitly.

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

api_instance = OpenapiClient::IpAsnApi.new

begin
  # Lookup ASN (Caller IP)
  result = api_instance.v1_networking_ip_asn_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpAsnApi->v1_networking_ip_asn_get: #{e}"
end
```

#### Using the v1_networking_ip_asn_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingIpAsnGet200Response>, Integer, Hash)> v1_networking_ip_asn_get_with_http_info

```ruby
begin
  # Lookup ASN (Caller IP)
  data, status_code, headers = api_instance.v1_networking_ip_asn_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingIpAsnGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpAsnApi->v1_networking_ip_asn_get_with_http_info: #{e}"
end
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


## v1_networking_ip_asn_ip_get

> <V1NetworkingIpAsnIpGet200Response> v1_networking_ip_asn_ip_get(ip)

Lookup ASN for IP

Look up ASN, organization, ISP, and network details for a specific IP address.

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

api_instance = OpenapiClient::IpAsnApi.new
ip = '8.8.8.8' # String | The IP address to look up (supports IPv4 and IPv6)

begin
  # Lookup ASN for IP
  result = api_instance.v1_networking_ip_asn_ip_get(ip)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpAsnApi->v1_networking_ip_asn_ip_get: #{e}"
end
```

#### Using the v1_networking_ip_asn_ip_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingIpAsnIpGet200Response>, Integer, Hash)> v1_networking_ip_asn_ip_get_with_http_info(ip)

```ruby
begin
  # Lookup ASN for IP
  data, status_code, headers = api_instance.v1_networking_ip_asn_ip_get_with_http_info(ip)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingIpAsnIpGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IpAsnApi->v1_networking_ip_asn_ip_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ip** | **String** | The IP address to look up (supports IPv4 and IPv6) |  |

### Return type

[**V1NetworkingIpAsnIpGet200Response**](V1NetworkingIpAsnIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

