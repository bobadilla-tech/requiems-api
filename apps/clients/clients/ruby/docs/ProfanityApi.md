# OpenapiClient::ProfanityApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_profanity_post**](ProfanityApi.md#v1_text_profanity_post) | **POST** /v1/text/profanity | Check Profanity |


## v1_text_profanity_post

> <V1TextProfanityPost200Response> v1_text_profanity_post(v1_text_profanity_post_request)

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
v1_text_profanity_post_request = OpenapiClient::V1TextProfanityPostRequest.new({text: 'What the heck is going on'}) # V1TextProfanityPostRequest | 

begin
  # Check Profanity
  result = api_instance.v1_text_profanity_post(v1_text_profanity_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProfanityApi->v1_text_profanity_post: #{e}"
end
```

#### Using the v1_text_profanity_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextProfanityPost200Response>, Integer, Hash)> v1_text_profanity_post_with_http_info(v1_text_profanity_post_request)

```ruby
begin
  # Check Profanity
  data, status_code, headers = api_instance.v1_text_profanity_post_with_http_info(v1_text_profanity_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextProfanityPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ProfanityApi->v1_text_profanity_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_profanity_post_request** | [**V1TextProfanityPostRequest**](V1TextProfanityPostRequest.md) |  |  |

### Return type

[**V1TextProfanityPost200Response**](V1TextProfanityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

