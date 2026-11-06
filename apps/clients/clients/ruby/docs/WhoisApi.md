# OpenapiClient::WhoisApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_whois_domain_get**](WhoisApi.md#v1_networking_whois_domain_get) | **GET** /v1/networking/whois/{domain} | WHOIS Lookup |


## v1_networking_whois_domain_get

> <V1NetworkingWhoisDomainGet200Response> v1_networking_whois_domain_get(domain)

WHOIS Lookup

Returns WHOIS registration information for a domain name.

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

api_instance = OpenapiClient::WhoisApi.new
domain = 'example.com' # String | The domain name to look up (e.g. example.com)

begin
  # WHOIS Lookup
  result = api_instance.v1_networking_whois_domain_get(domain)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling WhoisApi->v1_networking_whois_domain_get: #{e}"
end
```

#### Using the v1_networking_whois_domain_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingWhoisDomainGet200Response>, Integer, Hash)> v1_networking_whois_domain_get_with_http_info(domain)

```ruby
begin
  # WHOIS Lookup
  data, status_code, headers = api_instance.v1_networking_whois_domain_get_with_http_info(domain)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingWhoisDomainGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling WhoisApi->v1_networking_whois_domain_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domain** | **String** | The domain name to look up (e.g. example.com) |  |

### Return type

[**V1NetworkingWhoisDomainGet200Response**](V1NetworkingWhoisDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

