# OpenapiClient::ThesaurusApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_thesaurus_word_get**](ThesaurusApi.md#v1_text_thesaurus_word_get) | **GET** /v1/text/thesaurus/{word} | Thesaurus Lookup |


## v1_text_thesaurus_word_get

> <V1TextThesaurusWordGet200Response> v1_text_thesaurus_word_get(word)

Thesaurus Lookup

Returns synonyms and antonyms for the given word.

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

api_instance = OpenapiClient::ThesaurusApi.new
word = 'happy' # String | The word to look up in the thesaurus

begin
  # Thesaurus Lookup
  result = api_instance.v1_text_thesaurus_word_get(word)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling ThesaurusApi->v1_text_thesaurus_word_get: #{e}"
end
```

#### Using the v1_text_thesaurus_word_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextThesaurusWordGet200Response>, Integer, Hash)> v1_text_thesaurus_word_get_with_http_info(word)

```ruby
begin
  # Thesaurus Lookup
  data, status_code, headers = api_instance.v1_text_thesaurus_word_get_with_http_info(word)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextThesaurusWordGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling ThesaurusApi->v1_text_thesaurus_word_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **word** | **String** | The word to look up in the thesaurus |  |

### Return type

[**V1TextThesaurusWordGet200Response**](V1TextThesaurusWordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

