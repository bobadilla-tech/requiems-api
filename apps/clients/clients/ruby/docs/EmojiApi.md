# OpenapiClient::EmojiApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_entertainment_emoji_name_get**](EmojiApi.md#v1_entertainment_emoji_name_get) | **GET** /v1/entertainment/emoji/{name} | Get Emoji by Name |
| [**v1_entertainment_emoji_random_get**](EmojiApi.md#v1_entertainment_emoji_random_get) | **GET** /v1/entertainment/emoji/random | Get Random Emoji |
| [**v1_entertainment_emoji_search_get**](EmojiApi.md#v1_entertainment_emoji_search_get) | **GET** /v1/entertainment/emoji/search | Search Emoji |


## v1_entertainment_emoji_name_get

> <V1EntertainmentEmojiNameGet200Response> v1_entertainment_emoji_name_get(name)

Get Emoji by Name

Returns a specific emoji by its CLDR snake_case name. The name is case-insensitive.

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

api_instance = OpenapiClient::EmojiApi.new
name = 'grinning_face' # String | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up)

begin
  # Get Emoji by Name
  result = api_instance.v1_entertainment_emoji_name_get(name)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmojiApi->v1_entertainment_emoji_name_get: #{e}"
end
```

#### Using the v1_entertainment_emoji_name_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentEmojiNameGet200Response>, Integer, Hash)> v1_entertainment_emoji_name_get_with_http_info(name)

```ruby
begin
  # Get Emoji by Name
  data, status_code, headers = api_instance.v1_entertainment_emoji_name_get_with_http_info(name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentEmojiNameGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmojiApi->v1_entertainment_emoji_name_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **name** | **String** | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up) |  |

### Return type

[**V1EntertainmentEmojiNameGet200Response**](V1EntertainmentEmojiNameGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_entertainment_emoji_random_get

> <V1EntertainmentEmojiRandomGet200Response> v1_entertainment_emoji_random_get

Get Random Emoji

Returns a randomly selected emoji with its full metadata.

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

api_instance = OpenapiClient::EmojiApi.new

begin
  # Get Random Emoji
  result = api_instance.v1_entertainment_emoji_random_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmojiApi->v1_entertainment_emoji_random_get: #{e}"
end
```

#### Using the v1_entertainment_emoji_random_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentEmojiRandomGet200Response>, Integer, Hash)> v1_entertainment_emoji_random_get_with_http_info

```ruby
begin
  # Get Random Emoji
  data, status_code, headers = api_instance.v1_entertainment_emoji_random_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentEmojiRandomGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmojiApi->v1_entertainment_emoji_random_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1EntertainmentEmojiRandomGet200Response**](V1EntertainmentEmojiRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_entertainment_emoji_search_get

> <V1EntertainmentEmojiSearchGet200Response> v1_entertainment_emoji_search_get(q)

Search Emoji

Search for emojis whose name or category contains the given query string (case-insensitive). Returns a list of all matches.

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

api_instance = OpenapiClient::EmojiApi.new
q = 'happy' # String | Search term to match against emoji names and categories (e.g. happy, heart, food)

begin
  # Search Emoji
  result = api_instance.v1_entertainment_emoji_search_get(q)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmojiApi->v1_entertainment_emoji_search_get: #{e}"
end
```

#### Using the v1_entertainment_emoji_search_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1EntertainmentEmojiSearchGet200Response>, Integer, Hash)> v1_entertainment_emoji_search_get_with_http_info(q)

```ruby
begin
  # Search Emoji
  data, status_code, headers = api_instance.v1_entertainment_emoji_search_get_with_http_info(q)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1EntertainmentEmojiSearchGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmojiApi->v1_entertainment_emoji_search_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **q** | **String** | Search term to match against emoji names and categories (e.g. happy, heart, food) |  |

### Return type

[**V1EntertainmentEmojiSearchGet200Response**](V1EntertainmentEmojiSearchGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

