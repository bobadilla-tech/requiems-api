# OpenapiClient::SudokuApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_sudoku_batch_post**](SudokuApi.md#v1_entertainment_sudoku_batch_post) | **POST** /v1/entertainment/sudoku/batch | Batch Generate Sudoku Puzzles |
| [**v1_entertainment_sudoku_get**](SudokuApi.md#v1_entertainment_sudoku_get) | **GET** /v1/entertainment/sudoku | Get Sudoku Puzzle |


## v1_entertainment_sudoku_batch_post

> <V1EntertainmentSudokuBatchPost200Response> v1_entertainment_sudoku_batch_post(v1_entertainment_sudoku_batch_post_request)

Batch Generate Sudoku Puzzles

Generate up to 20 Sudoku puzzles in a single request. Results are returned in the same order as the input array. Each puzzle in the batch counts as one unit of API usage.

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

api_instance = OpenapiClient::SudokuApi.new
v1_entertainment_sudoku_batch_post_request = OpenapiClient::V1EntertainmentSudokuBatchPostRequest.new({puzzles: ["easy",  "medium",  "hard"]}) # V1EntertainmentSudokuBatchPostRequest | 

begin
  # Batch Generate Sudoku Puzzles
  result = api_instance.v1_entertainment_sudoku_batch_post(v1_entertainment_sudoku_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling SudokuApi->v1_entertainment_sudoku_batch_post: #{e}"
end
```

#### Using the v1_entertainment_sudoku_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentSudokuBatchPost200Response>, Integer, Hash)> v1_entertainment_sudoku_batch_post_with_http_info(v1_entertainment_sudoku_batch_post_request)

```ruby
begin
  # Batch Generate Sudoku Puzzles
  data, status_code, headers = api_instance.v1_entertainment_sudoku_batch_post_with_http_info(v1_entertainment_sudoku_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentSudokuBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling SudokuApi->v1_entertainment_sudoku_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_entertainment_sudoku_batch_post_request** | [**V1EntertainmentSudokuBatchPostRequest**](V1EntertainmentSudokuBatchPostRequest.md) |  |  |

### Return type

[**V1EntertainmentSudokuBatchPost200Response**](V1EntertainmentSudokuBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_entertainment_sudoku_get

> <V1EntertainmentSudokuGet200Response> v1_entertainment_sudoku_get(opts)

Get Sudoku Puzzle

Returns a randomly generated Sudoku puzzle and its solution. Difficulty defaults to medium when not specified.

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

api_instance = OpenapiClient::SudokuApi.new
opts = {
  difficulty: 'difficulty_example' # String | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium.
}

begin
  # Get Sudoku Puzzle
  result = api_instance.v1_entertainment_sudoku_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling SudokuApi->v1_entertainment_sudoku_get: #{e}"
end
```

#### Using the v1_entertainment_sudoku_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentSudokuGet200Response>, Integer, Hash)> v1_entertainment_sudoku_get_with_http_info(opts)

```ruby
begin
  # Get Sudoku Puzzle
  data, status_code, headers = api_instance.v1_entertainment_sudoku_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentSudokuGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling SudokuApi->v1_entertainment_sudoku_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **difficulty** | **String** | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. | [optional] |

### Return type

[**V1EntertainmentSudokuGet200Response**](V1EntertainmentSudokuGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

