# OpenapiClient::MxLookupApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_mx_batch_post**](MxLookupApi.md#v1_networking_mx_batch_post) | **POST** /v1/networking/mx/batch | Batch MX Lookup |
| [**v1_networking_mx_domain_get**](MxLookupApi.md#v1_networking_mx_domain_get) | **GET** /v1/networking/mx/{domain} | MX Lookup |


## v1_networking_mx_batch_post

> v1_networking_mx_batch_post(v1_networking_mx_batch_post_request)

Batch MX Lookup

Retrieve MX records for up to 50 domains in a single request. Results for each domain are sorted by priority ascending (lowest numeric value has the highest mail delivery priority per RFC 5321).

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

api_instance = OpenapiClient::MxLookupApi.new
v1_networking_mx_batch_post_request = OpenapiClient::V1NetworkingMxBatchPostRequest.new({domains: ["outlook.com",  "yahoo.com"]}) # V1NetworkingMxBatchPostRequest | 

begin
  # Batch MX Lookup
  api_instance.v1_networking_mx_batch_post(v1_networking_mx_batch_post_request)
rescue OpenapiClient::ApiError => e
  puts "Error when calling MxLookupApi->v1_networking_mx_batch_post: #{e}"
end
```

#### Using the v1_networking_mx_batch_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> v1_networking_mx_batch_post_with_http_info(v1_networking_mx_batch_post_request)

```ruby
begin
  # Batch MX Lookup
  data, status_code, headers = api_instance.v1_networking_mx_batch_post_with_http_info(v1_networking_mx_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling MxLookupApi->v1_networking_mx_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_networking_mx_batch_post_request** | [**V1NetworkingMxBatchPostRequest**](V1NetworkingMxBatchPostRequest.md) |  |  |

### Return type

nil (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## v1_networking_mx_domain_get

> <V1NetworkingMxDomainGet200Response> v1_networking_mx_domain_get(domain)

MX Lookup

Retrieve all MX records for a domain. Results are sorted by priority ascending (lowest numeric value has highest mail delivery priority per RFC 5321).

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

api_instance = OpenapiClient::MxLookupApi.new
domain = 'gmail.com' # String | The domain name to look up MX records for (e.g. gmail.com)

begin
  # MX Lookup
  result = api_instance.v1_networking_mx_domain_get(domain)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MxLookupApi->v1_networking_mx_domain_get: #{e}"
end
```

#### Using the v1_networking_mx_domain_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingMxDomainGet200Response>, Integer, Hash)> v1_networking_mx_domain_get_with_http_info(domain)

```ruby
begin
  # MX Lookup
  data, status_code, headers = api_instance.v1_networking_mx_domain_get_with_http_info(domain)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingMxDomainGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MxLookupApi->v1_networking_mx_domain_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domain** | **String** | The domain name to look up MX records for (e.g. gmail.com) |  |

### Return type

[**V1NetworkingMxDomainGet200Response**](V1NetworkingMxDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

