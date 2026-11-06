# OpenapiClient::DictionaryApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_dictionary_word_get**](DictionaryApi.md#v1_text_dictionary_word_get) | **GET** /v1/text/dictionary/{word} | Dictionary Lookup |


## v1_text_dictionary_word_get

> <V1TextDictionaryWordGet200Response> v1_text_dictionary_word_get(word)

Dictionary Lookup

Returns the definition, phonetics, examples, and synonyms for the given word.

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

api_instance = OpenapiClient::DictionaryApi.new
word = 'ephemeral' # String | The word to look up in the dictionary

begin
  # Dictionary Lookup
  result = api_instance.v1_text_dictionary_word_get(word)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DictionaryApi->v1_text_dictionary_word_get: #{e}"
end
```

#### Using the v1_text_dictionary_word_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextDictionaryWordGet200Response>, Integer, Hash)> v1_text_dictionary_word_get_with_http_info(word)

```ruby
begin
  # Dictionary Lookup
  data, status_code, headers = api_instance.v1_text_dictionary_word_get_with_http_info(word)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextDictionaryWordGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DictionaryApi->v1_text_dictionary_word_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **word** | **String** | The word to look up in the dictionary |  |

### Return type

[**V1TextDictionaryWordGet200Response**](V1TextDictionaryWordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

