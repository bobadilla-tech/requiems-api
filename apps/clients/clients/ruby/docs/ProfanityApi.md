# OpenapiClient::ProfanityApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_validation_profanity_batch_post**](ProfanityApi.md#v1_validation_profanity_batch_post) | **POST** /v1/validation/profanity/batch | Batch Check Profanity |
| [**v1_validation_profanity_post**](ProfanityApi.md#v1_validation_profanity_post) | **POST** /v1/validation/profanity | Check Profanity |


## v1_validation_profanity_batch_post

> <V1ValidationProfanityBatchPost200Response> v1_validation_profanity_batch_post(v1_validation_profanity_batch_post_request)

Batch Check Profanity

Check up to 50 texts for profanity in a single request. Results are returned in input order.

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

api_instance = OpenapiClient::ProfanityApi.new
v1_validation_profanity_batch_post_request = OpenapiClient::V1ValidationProfanityBatchPostRequest.new({texts: 'texts_example'}) # V1ValidationProfanityBatchPostRequest | 

begin
  # Batch Check Profanity
  result = api_instance.v1_validation_profanity_batch_post(v1_validation_profanity_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProfanityApi->v1_validation_profanity_batch_post: #{e}"
end
```

#### Using the v1_validation_profanity_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1ValidationProfanityBatchPost200Response>, Integer, Hash)> v1_validation_profanity_batch_post_with_http_info(v1_validation_profanity_batch_post_request)

```ruby
begin
  # Batch Check Profanity
  data, status_code, headers = api_instance.v1_validation_profanity_batch_post_with_http_info(v1_validation_profanity_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1ValidationProfanityBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProfanityApi->v1_validation_profanity_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_validation_profanity_batch_post_request** | [**V1ValidationProfanityBatchPostRequest**](V1ValidationProfanityBatchPostRequest.md) |  |  |

### Return type

[**V1ValidationProfanityBatchPost200Response**](V1ValidationProfanityBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_validation_profanity_post

> <V1ValidationProfanityPost200Response> v1_validation_profanity_post(v1_validation_profanity_post_request)

Check Profanity

Checks text for profanity, returning a censored version and the list of flagged words.

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

api_instance = OpenapiClient::ProfanityApi.new
v1_validation_profanity_post_request = OpenapiClient::V1ValidationProfanityPostRequest.new({text: 'What the heck is going on'}) # V1ValidationProfanityPostRequest | 

begin
  # Check Profanity
  result = api_instance.v1_validation_profanity_post(v1_validation_profanity_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProfanityApi->v1_validation_profanity_post: #{e}"
end
```

#### Using the v1_validation_profanity_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1ValidationProfanityPost200Response>, Integer, Hash)> v1_validation_profanity_post_with_http_info(v1_validation_profanity_post_request)

```ruby
begin
  # Check Profanity
  data, status_code, headers = api_instance.v1_validation_profanity_post_with_http_info(v1_validation_profanity_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1ValidationProfanityPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProfanityApi->v1_validation_profanity_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_validation_profanity_post_request** | [**V1ValidationProfanityPostRequest**](V1ValidationProfanityPostRequest.md) |  |  |

### Return type

[**V1ValidationProfanityPost200Response**](V1ValidationProfanityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

