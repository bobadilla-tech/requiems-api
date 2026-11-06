# OpenapiClient::TriviaApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_trivia_get**](TriviaApi.md#v1_entertainment_trivia_get) | **GET** /v1/entertainment/trivia | Get Trivia Question |


## v1_entertainment_trivia_get

> <V1EntertainmentTriviaGet200Response> v1_entertainment_trivia_get(opts)

Get Trivia Question

Returns a random trivia question with multiple-choice answers. Use the optional category and difficulty query parameters to filter the question pool.

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

api_instance = OpenapiClient::TriviaApi.new
opts = {
  category: 'category_example', # String | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature.
  difficulty: 'difficulty_example' # String | Filter by difficulty. One of: easy, medium, hard.
}

begin
  # Get Trivia Question
  result = api_instance.v1_entertainment_trivia_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling TriviaApi->v1_entertainment_trivia_get: #{e}"
end
```

#### Using the v1_entertainment_trivia_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentTriviaGet200Response>, Integer, Hash)> v1_entertainment_trivia_get_with_http_info(opts)

```ruby
begin
  # Get Trivia Question
  data, status_code, headers = api_instance.v1_entertainment_trivia_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentTriviaGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling TriviaApi->v1_entertainment_trivia_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **category** | **String** | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. | [optional] |
| **difficulty** | **String** | Filter by difficulty. One of: easy, medium, hard. | [optional] |

### Return type

[**V1EntertainmentTriviaGet200Response**](V1EntertainmentTriviaGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

