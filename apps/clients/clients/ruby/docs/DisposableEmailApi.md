# OpenapiClient::DisposableEmailApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_networking_disposable_batch_post**](DisposableEmailApi.md#v1_networking_disposable_batch_post) | **POST** /v1/networking/disposable/batch | Check Batch Emails |
| [**v1_networking_disposable_check_post**](DisposableEmailApi.md#v1_networking_disposable_check_post) | **POST** /v1/networking/disposable/check | Check Single Email |
| [**v1_networking_disposable_domain_domain_get**](DisposableEmailApi.md#v1_networking_disposable_domain_domain_get) | **GET** /v1/networking/disposable/domain/{domain} | Check Domain |
| [**v1_networking_disposable_domains_get**](DisposableEmailApi.md#v1_networking_disposable_domains_get) | **GET** /v1/networking/disposable/domains | List Domains (Paginated) |
| [**v1_networking_disposable_stats_get**](DisposableEmailApi.md#v1_networking_disposable_stats_get) | **GET** /v1/networking/disposable/stats | Get Statistics |


## v1_networking_disposable_batch_post

> <V1NetworkingDisposableBatchPost200Response> v1_networking_disposable_batch_post(v1_networking_disposable_batch_post_request)

Check Batch Emails

Validate multiple email addresses in a single request (max 100 emails)

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

api_instance = OpenapiClient::DisposableEmailApi.new
v1_networking_disposable_batch_post_request = OpenapiClient::V1NetworkingDisposableBatchPostRequest.new({emails: ["user1@example.com", "user2@tempmail.com"]}) # V1NetworkingDisposableBatchPostRequest | 

begin
  # Check Batch Emails
  result = api_instance.v1_networking_disposable_batch_post(v1_networking_disposable_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_batch_post: #{e}"
end
```

#### Using the v1_networking_disposable_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingDisposableBatchPost200Response>, Integer, Hash)> v1_networking_disposable_batch_post_with_http_info(v1_networking_disposable_batch_post_request)

```ruby
begin
  # Check Batch Emails
  data, status_code, headers = api_instance.v1_networking_disposable_batch_post_with_http_info(v1_networking_disposable_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingDisposableBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_networking_disposable_batch_post_request** | [**V1NetworkingDisposableBatchPostRequest**](V1NetworkingDisposableBatchPostRequest.md) |  |  |

### Return type

[**V1NetworkingDisposableBatchPost200Response**](V1NetworkingDisposableBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_networking_disposable_check_post

> <V1NetworkingDisposableCheckPost200Response> v1_networking_disposable_check_post(v1_networking_disposable_check_post_request)

Check Single Email

Validate whether an email address uses a disposable domain

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

api_instance = OpenapiClient::DisposableEmailApi.new
v1_networking_disposable_check_post_request = OpenapiClient::V1NetworkingDisposableCheckPostRequest.new({email: 'test@example.com'}) # V1NetworkingDisposableCheckPostRequest | 

begin
  # Check Single Email
  result = api_instance.v1_networking_disposable_check_post(v1_networking_disposable_check_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_check_post: #{e}"
end
```

#### Using the v1_networking_disposable_check_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingDisposableCheckPost200Response>, Integer, Hash)> v1_networking_disposable_check_post_with_http_info(v1_networking_disposable_check_post_request)

```ruby
begin
  # Check Single Email
  data, status_code, headers = api_instance.v1_networking_disposable_check_post_with_http_info(v1_networking_disposable_check_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingDisposableCheckPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_check_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_networking_disposable_check_post_request** | [**V1NetworkingDisposableCheckPostRequest**](V1NetworkingDisposableCheckPostRequest.md) |  |  |

### Return type

[**V1NetworkingDisposableCheckPost200Response**](V1NetworkingDisposableCheckPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_networking_disposable_domain_domain_get

> <V1NetworkingDisposableDomainDomainGet200Response> v1_networking_disposable_domain_domain_get(domain)

Check Domain

Check if a specific domain is in the disposable blocklist

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

api_instance = OpenapiClient::DisposableEmailApi.new
domain = 'tempmail.com' # String | The domain to check

begin
  # Check Domain
  result = api_instance.v1_networking_disposable_domain_domain_get(domain)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_domain_domain_get: #{e}"
end
```

#### Using the v1_networking_disposable_domain_domain_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingDisposableDomainDomainGet200Response>, Integer, Hash)> v1_networking_disposable_domain_domain_get_with_http_info(domain)

```ruby
begin
  # Check Domain
  data, status_code, headers = api_instance.v1_networking_disposable_domain_domain_get_with_http_info(domain)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingDisposableDomainDomainGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_domain_domain_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domain** | **String** | The domain to check |  |

### Return type

[**V1NetworkingDisposableDomainDomainGet200Response**](V1NetworkingDisposableDomainDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_networking_disposable_domains_get

> <V1NetworkingDisposableDomainsGet200Response> v1_networking_disposable_domains_get(opts)

List Domains (Paginated)

Get a paginated list of all disposable domains in the blocklist

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

api_instance = OpenapiClient::DisposableEmailApi.new
opts = {
  page: 1, # Integer | Page number (default: 1)
  per_page: 100 # Integer | Items per page (default: 100, max: 1000)
}

begin
  # List Domains (Paginated)
  result = api_instance.v1_networking_disposable_domains_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_domains_get: #{e}"
end
```

#### Using the v1_networking_disposable_domains_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingDisposableDomainsGet200Response>, Integer, Hash)> v1_networking_disposable_domains_get_with_http_info(opts)

```ruby
begin
  # List Domains (Paginated)
  data, status_code, headers = api_instance.v1_networking_disposable_domains_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingDisposableDomainsGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_domains_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **page** | **Integer** | Page number (default: 1) | [optional] |
| **per_page** | **Integer** | Items per page (default: 100, max: 1000) | [optional] |

### Return type

[**V1NetworkingDisposableDomainsGet200Response**](V1NetworkingDisposableDomainsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_networking_disposable_stats_get

> <V1NetworkingDisposableStatsGet200Response> v1_networking_disposable_stats_get

Get Statistics

Get statistics about the disposable email blocklist

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

api_instance = OpenapiClient::DisposableEmailApi.new

begin
  # Get Statistics
  result = api_instance.v1_networking_disposable_stats_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_stats_get: #{e}"
end
```

#### Using the v1_networking_disposable_stats_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1NetworkingDisposableStatsGet200Response>, Integer, Hash)> v1_networking_disposable_stats_get_with_http_info

```ruby
begin
  # Get Statistics
  data, status_code, headers = api_instance.v1_networking_disposable_stats_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1NetworkingDisposableStatsGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DisposableEmailApi->v1_networking_disposable_stats_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1NetworkingDisposableStatsGet200Response**](V1NetworkingDisposableStatsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

