# OpenapiClient::EmailNormalizeApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_normalize_batch_post**](EmailNormalizeApi.md#v1_text_normalize_batch_post) | **POST** /v1/text/normalize/batch | Normalize Email Batch |
| [**v1_text_normalize_post**](EmailNormalizeApi.md#v1_text_normalize_post) | **POST** /v1/text/normalize | Normalize Email |


## v1_text_normalize_batch_post

> <V1TextNormalizeBatchPost200Response> v1_text_normalize_batch_post(v1_text_normalize_batch_post_request)

Normalize Email Batch

Normalizes up to 100 email addresses in one request. Results are in the same order as the input. Each item includes valid (boolean); when false, only original and message are set. Usage is billed per email processed (see gateway usage headers).

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

api_instance = OpenapiClient::EmailNormalizeApi.new
v1_text_normalize_batch_post_request = OpenapiClient::V1TextNormalizeBatchPostRequest.new({emails: ["user@example.com", "not-an-email", "te.st@gmail.com"]}) # V1TextNormalizeBatchPostRequest | 

begin
  # Normalize Email Batch
  result = api_instance.v1_text_normalize_batch_post(v1_text_normalize_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailNormalizeApi->v1_text_normalize_batch_post: #{e}"
end
```

#### Using the v1_text_normalize_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextNormalizeBatchPost200Response>, Integer, Hash)> v1_text_normalize_batch_post_with_http_info(v1_text_normalize_batch_post_request)

```ruby
begin
  # Normalize Email Batch
  data, status_code, headers = api_instance.v1_text_normalize_batch_post_with_http_info(v1_text_normalize_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextNormalizeBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailNormalizeApi->v1_text_normalize_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_normalize_batch_post_request** | [**V1TextNormalizeBatchPostRequest**](V1TextNormalizeBatchPostRequest.md) |  |  |

### Return type

[**V1TextNormalizeBatchPost200Response**](V1TextNormalizeBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_text_normalize_post

> <V1TextNormalizePost200Response> v1_text_normalize_post(v1_text_normalize_post_request)

Normalize Email

Normalizes a single email address and returns the canonical form together with a breakdown of all transformations applied.

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

api_instance = OpenapiClient::EmailNormalizeApi.new
v1_text_normalize_post_request = OpenapiClient::V1TextNormalizePostRequest.new({email: 'Te.st.User+spam@Googlemail.com'}) # V1TextNormalizePostRequest | 

begin
  # Normalize Email
  result = api_instance.v1_text_normalize_post(v1_text_normalize_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailNormalizeApi->v1_text_normalize_post: #{e}"
end
```

#### Using the v1_text_normalize_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextNormalizePost200Response>, Integer, Hash)> v1_text_normalize_post_with_http_info(v1_text_normalize_post_request)

```ruby
begin
  # Normalize Email
  data, status_code, headers = api_instance.v1_text_normalize_post_with_http_info(v1_text_normalize_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextNormalizePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailNormalizeApi->v1_text_normalize_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_normalize_post_request** | [**V1TextNormalizePostRequest**](V1TextNormalizePostRequest.md) |  |  |

### Return type

[**V1TextNormalizePost200Response**](V1TextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

