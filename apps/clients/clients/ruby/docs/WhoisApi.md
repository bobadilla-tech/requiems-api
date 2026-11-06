# OpenapiClient::WhoisApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_whois_batch_post**](WhoisApi.md#v1_networking_whois_batch_post) | **POST** /v1/networking/whois/batch | Batch WHOIS Lookup |
| [**v1_networking_whois_domain_get**](WhoisApi.md#v1_networking_whois_domain_get) | **GET** /v1/networking/whois/{domain} | WHOIS Lookup |


## v1_networking_whois_batch_post

> <V1NetworkingWhoisBatchPost200Response> v1_networking_whois_batch_post(v1_networking_whois_batch_post_request)

Batch WHOIS Lookup

Returns WHOIS information for up to 50 domains in a single request. Results are returned in the same order as the input array. Domains without WHOIS data return found: false instead of failing the entire request.

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
v1_networking_whois_batch_post_request = OpenapiClient::V1NetworkingWhoisBatchPostRequest.new({domains: ["example.com",  "google.com"]}) # V1NetworkingWhoisBatchPostRequest | 

begin
  # Batch WHOIS Lookup
  result = api_instance.v1_networking_whois_batch_post(v1_networking_whois_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling WhoisApi->v1_networking_whois_batch_post: #{e}"
end
```

#### Using the v1_networking_whois_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingWhoisBatchPost200Response>, Integer, Hash)> v1_networking_whois_batch_post_with_http_info(v1_networking_whois_batch_post_request)

```ruby
begin
  # Batch WHOIS Lookup
  data, status_code, headers = api_instance.v1_networking_whois_batch_post_with_http_info(v1_networking_whois_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingWhoisBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling WhoisApi->v1_networking_whois_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_networking_whois_batch_post_request** | [**V1NetworkingWhoisBatchPostRequest**](V1NetworkingWhoisBatchPostRequest.md) |  |  |

### Return type

[**V1NetworkingWhoisBatchPost200Response**](V1NetworkingWhoisBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


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

