# OpenapiClient::QuotesApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_quotes_random_batch_post**](QuotesApi.md#v1_entertainment_quotes_random_batch_post) | **POST** /v1/entertainment/quotes/random/batch | Get Random Quotes (Batch) |
| [**v1_entertainment_quotes_random_get**](QuotesApi.md#v1_entertainment_quotes_random_get) | **GET** /v1/entertainment/quotes/random | Get Random Quote |


## v1_entertainment_quotes_random_batch_post

> <V1EntertainmentQuotesRandomBatchPost200Response> v1_entertainment_quotes_random_batch_post(v1_entertainment_quotes_random_batch_post_request)

Get Random Quotes (Batch)

Returns up to 50 random quotes in a single request. Each quote counts as one unit of usage.

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

api_instance = OpenapiClient::QuotesApi.new
v1_entertainment_quotes_random_batch_post_request = OpenapiClient::V1EntertainmentQuotesRandomBatchPostRequest.new({count: 5}) # V1EntertainmentQuotesRandomBatchPostRequest | 

begin
  # Get Random Quotes (Batch)
  result = api_instance.v1_entertainment_quotes_random_batch_post(v1_entertainment_quotes_random_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling QuotesApi->v1_entertainment_quotes_random_batch_post: #{e}"
end
```

#### Using the v1_entertainment_quotes_random_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentQuotesRandomBatchPost200Response>, Integer, Hash)> v1_entertainment_quotes_random_batch_post_with_http_info(v1_entertainment_quotes_random_batch_post_request)

```ruby
begin
  # Get Random Quotes (Batch)
  data, status_code, headers = api_instance.v1_entertainment_quotes_random_batch_post_with_http_info(v1_entertainment_quotes_random_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentQuotesRandomBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling QuotesApi->v1_entertainment_quotes_random_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_entertainment_quotes_random_batch_post_request** | [**V1EntertainmentQuotesRandomBatchPostRequest**](V1EntertainmentQuotesRandomBatchPostRequest.md) |  |  |

### Return type

[**V1EntertainmentQuotesRandomBatchPost200Response**](V1EntertainmentQuotesRandomBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_entertainment_quotes_random_get

> <V1EntertainmentQuotesRandomGet200Response> v1_entertainment_quotes_random_get

Get Random Quote

Returns a random inspirational quote with author attribution

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

api_instance = OpenapiClient::QuotesApi.new

begin
  # Get Random Quote
  result = api_instance.v1_entertainment_quotes_random_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling QuotesApi->v1_entertainment_quotes_random_get: #{e}"
end
```

#### Using the v1_entertainment_quotes_random_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentQuotesRandomGet200Response>, Integer, Hash)> v1_entertainment_quotes_random_get_with_http_info

```ruby
begin
  # Get Random Quote
  data, status_code, headers = api_instance.v1_entertainment_quotes_random_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentQuotesRandomGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling QuotesApi->v1_entertainment_quotes_random_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1EntertainmentQuotesRandomGet200Response**](V1EntertainmentQuotesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

