# OpenapiClient::DataIntegrityApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_systems_content_moderate_post**](DataIntegrityApi.md#v1_systems_content_moderate_post) | **POST** /v1/systems/content/moderate | Moderate Content |
| [**v1_systems_text_normalize_post**](DataIntegrityApi.md#v1_systems_text_normalize_post) | **POST** /v1/systems/text/normalize | Normalize Text |


## v1_systems_content_moderate_post

> <V1SystemsContentModeratePost200Response> v1_systems_content_moderate_post(v1_systems_content_moderate_post_request)

Moderate Content

Check text for profanity, toxicity, sentiment, and language. Returns per-category flags and an is_safe boolean for gating content.

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

api_instance = OpenapiClient::DataIntegrityApi.new
v1_systems_content_moderate_post_request = OpenapiClient::V1SystemsContentModeratePostRequest.new({text: 'This is some user-generated content.'}) # V1SystemsContentModeratePostRequest | 

begin
  # Moderate Content
  result = api_instance.v1_systems_content_moderate_post(v1_systems_content_moderate_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DataIntegrityApi->v1_systems_content_moderate_post: #{e}"
end
```

#### Using the v1_systems_content_moderate_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsContentModeratePost200Response>, Integer, Hash)> v1_systems_content_moderate_post_with_http_info(v1_systems_content_moderate_post_request)

```ruby
begin
  # Moderate Content
  data, status_code, headers = api_instance.v1_systems_content_moderate_post_with_http_info(v1_systems_content_moderate_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsContentModeratePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DataIntegrityApi->v1_systems_content_moderate_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_content_moderate_post_request** | [**V1SystemsContentModeratePostRequest**](V1SystemsContentModeratePostRequest.md) |  |  |

### Return type

[**V1SystemsContentModeratePost200Response**](V1SystemsContentModeratePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_systems_text_normalize_post

> <V1SystemsTextNormalizePost200Response> v1_systems_text_normalize_post(v1_systems_text_normalize_post_request)

Normalize Text

Apply a composable pipeline of normalization operations to a string. Trim whitespace, fix encoding, normalize unicode, strip HTML, and more — in a single call.

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

api_instance = OpenapiClient::DataIntegrityApi.new
v1_systems_text_normalize_post_request = OpenapiClient::V1SystemsTextNormalizePostRequest.new({text: '  Hello,   World!  ', operations: ["trim",  "lowercase",  "collapse_whitespace"]}) # V1SystemsTextNormalizePostRequest | 

begin
  # Normalize Text
  result = api_instance.v1_systems_text_normalize_post(v1_systems_text_normalize_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DataIntegrityApi->v1_systems_text_normalize_post: #{e}"
end
```

#### Using the v1_systems_text_normalize_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsTextNormalizePost200Response>, Integer, Hash)> v1_systems_text_normalize_post_with_http_info(v1_systems_text_normalize_post_request)

```ruby
begin
  # Normalize Text
  data, status_code, headers = api_instance.v1_systems_text_normalize_post_with_http_info(v1_systems_text_normalize_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsTextNormalizePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DataIntegrityApi->v1_systems_text_normalize_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_text_normalize_post_request** | [**V1SystemsTextNormalizePostRequest**](V1SystemsTextNormalizePostRequest.md) |  |  |

### Return type

[**V1SystemsTextNormalizePost200Response**](V1SystemsTextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

