# OpenapiClient::DomainInfoApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_domain_domain_get**](DomainInfoApi.md#v1_networking_domain_domain_get) | **GET** /v1/networking/domain/{domain} | Get Domain Info |


## v1_networking_domain_domain_get

> <V1NetworkingDomainDomainGet200Response> v1_networking_domain_domain_get(domain)

Get Domain Info

Returns DNS records and availability status for the given domain.

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

api_instance = OpenapiClient::DomainInfoApi.new
domain = 'example.com' # String | The domain to look up (e.g. example.com)

begin
  # Get Domain Info
  result = api_instance.v1_networking_domain_domain_get(domain)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DomainInfoApi->v1_networking_domain_domain_get: #{e}"
end
```

#### Using the v1_networking_domain_domain_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingDomainDomainGet200Response>, Integer, Hash)> v1_networking_domain_domain_get_with_http_info(domain)

```ruby
begin
  # Get Domain Info
  data, status_code, headers = api_instance.v1_networking_domain_domain_get_with_http_info(domain)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingDomainDomainGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DomainInfoApi->v1_networking_domain_domain_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domain** | **String** | The domain to look up (e.g. example.com) |  |

### Return type

[**V1NetworkingDomainDomainGet200Response**](V1NetworkingDomainDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

