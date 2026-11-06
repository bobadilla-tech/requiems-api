# OpenapiClient::SpellCheckApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_spellcheck_post**](SpellCheckApi.md#v1_text_spellcheck_post) | **POST** /v1/text/spellcheck | Check Spelling |


## v1_text_spellcheck_post

> <V1TextSpellcheckPost200Response> v1_text_spellcheck_post(v1_text_spellcheck_post_request)

Check Spelling

Checks the input text for spelling mistakes and returns a corrected version along with per-word corrections.

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

api_instance = OpenapiClient::SpellCheckApi.new
v1_text_spellcheck_post_request = OpenapiClient::V1TextSpellcheckPostRequest.new({text: 'Ths is a smiple tset'}) # V1TextSpellcheckPostRequest | 

begin
  # Check Spelling
  result = api_instance.v1_text_spellcheck_post(v1_text_spellcheck_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling SpellCheckApi->v1_text_spellcheck_post: #{e}"
end
```

#### Using the v1_text_spellcheck_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextSpellcheckPost200Response>, Integer, Hash)> v1_text_spellcheck_post_with_http_info(v1_text_spellcheck_post_request)

```ruby
begin
  # Check Spelling
  data, status_code, headers = api_instance.v1_text_spellcheck_post_with_http_info(v1_text_spellcheck_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextSpellcheckPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling SpellCheckApi->v1_text_spellcheck_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_spellcheck_post_request** | [**V1TextSpellcheckPostRequest**](V1TextSpellcheckPostRequest.md) |  |  |

### Return type

[**V1TextSpellcheckPost200Response**](V1TextSpellcheckPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

