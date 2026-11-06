# OpenapiClient::DetectLanguageApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_detect_language_post**](DetectLanguageApi.md#v1_text_detect_language_post) | **POST** /v1/text/detect-language | Detect Language |


## v1_text_detect_language_post

> <V1TextDetectLanguagePost200Response> v1_text_detect_language_post(v1_text_detect_language_post_request)

Detect Language

Identifies the language of the provided text and returns the language name, ISO 639-1 code, and confidence score.

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

api_instance = OpenapiClient::DetectLanguageApi.new
v1_text_detect_language_post_request = OpenapiClient::V1TextDetectLanguagePostRequest.new({text: 'Bonjour, comment ça va?'}) # V1TextDetectLanguagePostRequest | 

begin
  # Detect Language
  result = api_instance.v1_text_detect_language_post(v1_text_detect_language_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DetectLanguageApi->v1_text_detect_language_post: #{e}"
end
```

#### Using the v1_text_detect_language_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextDetectLanguagePost200Response>, Integer, Hash)> v1_text_detect_language_post_with_http_info(v1_text_detect_language_post_request)

```ruby
begin
  # Detect Language
  data, status_code, headers = api_instance.v1_text_detect_language_post_with_http_info(v1_text_detect_language_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextDetectLanguagePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DetectLanguageApi->v1_text_detect_language_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_detect_language_post_request** | [**V1TextDetectLanguagePostRequest**](V1TextDetectLanguagePostRequest.md) |  |  |

### Return type

[**V1TextDetectLanguagePost200Response**](V1TextDetectLanguagePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

