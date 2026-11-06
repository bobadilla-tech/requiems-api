# OpenapiClient::DeveloperUtilitiesApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_base64_decode_post**](DeveloperUtilitiesApi.md#v1_technology_base64_decode_post) | **POST** /v1/technology/base64/decode | Base64 Decode |
| [**v1_technology_base64_encode_post**](DeveloperUtilitiesApi.md#v1_technology_base64_encode_post) | **POST** /v1/technology/base64/encode | Base64 Encode |


## v1_technology_base64_decode_post

> <V1TechnologyBase64DecodePost200Response> v1_technology_base64_decode_post(v1_technology_base64_decode_post_request)

Base64 Decode

Decode a Base64-encoded string back to plain text. Returns 422 for invalid or unpadded input.

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

api_instance = OpenapiClient::DeveloperUtilitiesApi.new
v1_technology_base64_decode_post_request = OpenapiClient::V1TechnologyBase64DecodePostRequest.new({value: 'SGVsbG8sIHdvcmxkIQ=='}) # V1TechnologyBase64DecodePostRequest | 

begin
  # Base64 Decode
  result = api_instance.v1_technology_base64_decode_post(v1_technology_base64_decode_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DeveloperUtilitiesApi->v1_technology_base64_decode_post: #{e}"
end
```

#### Using the v1_technology_base64_decode_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyBase64DecodePost200Response>, Integer, Hash)> v1_technology_base64_decode_post_with_http_info(v1_technology_base64_decode_post_request)

```ruby
begin
  # Base64 Decode
  data, status_code, headers = api_instance.v1_technology_base64_decode_post_with_http_info(v1_technology_base64_decode_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyBase64DecodePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DeveloperUtilitiesApi->v1_technology_base64_decode_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_technology_base64_decode_post_request** | [**V1TechnologyBase64DecodePostRequest**](V1TechnologyBase64DecodePostRequest.md) |  |  |

### Return type

[**V1TechnologyBase64DecodePost200Response**](V1TechnologyBase64DecodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_technology_base64_encode_post

> <V1TechnologyBase64EncodePost200Response> v1_technology_base64_encode_post(v1_technology_base64_encode_post_request)

Base64 Encode

Encode a plain-text string to Base64. Supports standard (RFC 4648) and URL-safe (base64url) variants.

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

api_instance = OpenapiClient::DeveloperUtilitiesApi.new
v1_technology_base64_encode_post_request = OpenapiClient::V1TechnologyBase64EncodePostRequest.new({value: 'Hello, world!'}) # V1TechnologyBase64EncodePostRequest | 

begin
  # Base64 Encode
  result = api_instance.v1_technology_base64_encode_post(v1_technology_base64_encode_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DeveloperUtilitiesApi->v1_technology_base64_encode_post: #{e}"
end
```

#### Using the v1_technology_base64_encode_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyBase64EncodePost200Response>, Integer, Hash)> v1_technology_base64_encode_post_with_http_info(v1_technology_base64_encode_post_request)

```ruby
begin
  # Base64 Encode
  data, status_code, headers = api_instance.v1_technology_base64_encode_post_with_http_info(v1_technology_base64_encode_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyBase64EncodePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DeveloperUtilitiesApi->v1_technology_base64_encode_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_technology_base64_encode_post_request** | [**V1TechnologyBase64EncodePostRequest**](V1TechnologyBase64EncodePostRequest.md) |  |  |

### Return type

[**V1TechnologyBase64EncodePost200Response**](V1TechnologyBase64EncodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

