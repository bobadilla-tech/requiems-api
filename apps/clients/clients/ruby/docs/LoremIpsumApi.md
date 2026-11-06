# OpenapiClient::LoremIpsumApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_lorem_batch_post**](LoremIpsumApi.md#v1_text_lorem_batch_post) | **POST** /v1/text/lorem/batch | Generate Lorem Ipsum (Batch) |
| [**v1_text_lorem_get**](LoremIpsumApi.md#v1_text_lorem_get) | **GET** /v1/text/lorem | Generate Lorem Ipsum |


## v1_text_lorem_batch_post

> <V1TextLoremBatchPost200Response> v1_text_lorem_batch_post(v1_text_lorem_batch_post_request)

Generate Lorem Ipsum (Batch)

Generate multiple Lorem Ipsum placeholder texts in a single request. Processes items in bulk and returns partial successes.

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

api_instance = OpenapiClient::LoremIpsumApi.new
v1_text_lorem_batch_post_request = OpenapiClient::V1TextLoremBatchPostRequest.new({items: [3.56]}) # V1TextLoremBatchPostRequest | 

begin
  # Generate Lorem Ipsum (Batch)
  result = api_instance.v1_text_lorem_batch_post(v1_text_lorem_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling LoremIpsumApi->v1_text_lorem_batch_post: #{e}"
end
```

#### Using the v1_text_lorem_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextLoremBatchPost200Response>, Integer, Hash)> v1_text_lorem_batch_post_with_http_info(v1_text_lorem_batch_post_request)

```ruby
begin
  # Generate Lorem Ipsum (Batch)
  data, status_code, headers = api_instance.v1_text_lorem_batch_post_with_http_info(v1_text_lorem_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextLoremBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling LoremIpsumApi->v1_text_lorem_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_lorem_batch_post_request** | [**V1TextLoremBatchPostRequest**](V1TextLoremBatchPostRequest.md) |  |  |

### Return type

[**V1TextLoremBatchPost200Response**](V1TextLoremBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_text_lorem_get

> <V1TextLoremGet200Response> v1_text_lorem_get(opts)

Generate Lorem Ipsum

Generate Lorem Ipsum placeholder text with customizable length and format

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

api_instance = OpenapiClient::LoremIpsumApi.new
opts = {
  paragraphs: 3, # Integer | Number of paragraphs to generate (1-20)
  sentences: 5 # Integer | Number of sentences per paragraph (1-20)
}

begin
  # Generate Lorem Ipsum
  result = api_instance.v1_text_lorem_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling LoremIpsumApi->v1_text_lorem_get: #{e}"
end
```

#### Using the v1_text_lorem_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextLoremGet200Response>, Integer, Hash)> v1_text_lorem_get_with_http_info(opts)

```ruby
begin
  # Generate Lorem Ipsum
  data, status_code, headers = api_instance.v1_text_lorem_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextLoremGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling LoremIpsumApi->v1_text_lorem_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **paragraphs** | **Integer** | Number of paragraphs to generate (1-20) | [optional] |
| **sentences** | **Integer** | Number of sentences per paragraph (1-20) | [optional] |

### Return type

[**V1TextLoremGet200Response**](V1TextLoremGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

