# OpenapiClient::RandomWordApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_words_batch_post**](RandomWordApi.md#v1_text_words_batch_post) | **POST** /v1/text/words/batch | Batch Define Words |
| [**v1_text_words_random_get**](RandomWordApi.md#v1_text_words_random_get) | **GET** /v1/text/words/random | Get Random Word |


## v1_text_words_batch_post

> <V1TextWordsBatchPost200Response> v1_text_words_batch_post(v1_text_words_batch_post_request)

Batch Define Words

Resolve multiple words in a single request. Returns dictionary entries when found, or error information when a word is not in the dataset. 

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

api_instance = OpenapiClient::RandomWordApi.new
v1_text_words_batch_post_request = OpenapiClient::V1TextWordsBatchPostRequest.new({items: [3.56]}) # V1TextWordsBatchPostRequest | 

begin
  # Batch Define Words
  result = api_instance.v1_text_words_batch_post(v1_text_words_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling RandomWordApi->v1_text_words_batch_post: #{e}"
end
```

#### Using the v1_text_words_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextWordsBatchPost200Response>, Integer, Hash)> v1_text_words_batch_post_with_http_info(v1_text_words_batch_post_request)

```ruby
begin
  # Batch Define Words
  data, status_code, headers = api_instance.v1_text_words_batch_post_with_http_info(v1_text_words_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextWordsBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling RandomWordApi->v1_text_words_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_words_batch_post_request** | [**V1TextWordsBatchPostRequest**](V1TextWordsBatchPostRequest.md) |  |  |

### Return type

[**V1TextWordsBatchPost200Response**](V1TextWordsBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_text_words_random_get

> <V1TextWordsRandomGet200Response> v1_text_words_random_get

Get Random Word

Returns a random word with its definition and part of speech

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

api_instance = OpenapiClient::RandomWordApi.new

begin
  # Get Random Word
  result = api_instance.v1_text_words_random_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling RandomWordApi->v1_text_words_random_get: #{e}"
end
```

#### Using the v1_text_words_random_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextWordsRandomGet200Response>, Integer, Hash)> v1_text_words_random_get_with_http_info

```ruby
begin
  # Get Random Word
  data, status_code, headers = api_instance.v1_text_words_random_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextWordsRandomGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling RandomWordApi->v1_text_words_random_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1TextWordsRandomGet200Response**](V1TextWordsRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

