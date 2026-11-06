# OpenapiClient::TextSimilarityApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_similarity_post**](TextSimilarityApi.md#v1_text_similarity_post) | **POST** /v1/text/similarity | Compare Text Similarity |


## v1_text_similarity_post

> <V1TextSimilarityPost200Response> v1_text_similarity_post(v1_text_similarity_post_request)

Compare Text Similarity

Compares two texts and returns a cosine similarity score.

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

api_instance = OpenapiClient::TextSimilarityApi.new
v1_text_similarity_post_request = OpenapiClient::V1TextSimilarityPostRequest.new({text1: 'The cat sat on the mat', text2: 'A cat was sitting on a mat'}) # V1TextSimilarityPostRequest | 

begin
  # Compare Text Similarity
  result = api_instance.v1_text_similarity_post(v1_text_similarity_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling TextSimilarityApi->v1_text_similarity_post: #{e}"
end
```

#### Using the v1_text_similarity_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextSimilarityPost200Response>, Integer, Hash)> v1_text_similarity_post_with_http_info(v1_text_similarity_post_request)

```ruby
begin
  # Compare Text Similarity
  data, status_code, headers = api_instance.v1_text_similarity_post_with_http_info(v1_text_similarity_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextSimilarityPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling TextSimilarityApi->v1_text_similarity_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_similarity_post_request** | [**V1TextSimilarityPostRequest**](V1TextSimilarityPostRequest.md) |  |  |

### Return type

[**V1TextSimilarityPost200Response**](V1TextSimilarityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

