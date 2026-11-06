# OpenapiClient::DataFormatConversionApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_format_post**](DataFormatConversionApi.md#v1_technology_format_post) | **POST** /v1/technology/format | Convert Format |


## v1_technology_format_post

> <V1TechnologyFormatPost200Response> v1_technology_format_post(v1_technology_format_post_request)

Convert Format

Convert content from one structured data format to another. Supported formats: json, yaml, csv, xml, toml.

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

api_instance = OpenapiClient::DataFormatConversionApi.new
v1_technology_format_post_request = OpenapiClient::V1TechnologyFormatPostRequest.new({from: 'json', to: 'yaml', content: '{"name":"Alice","age":30}'}) # V1TechnologyFormatPostRequest | 

begin
  # Convert Format
  result = api_instance.v1_technology_format_post(v1_technology_format_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DataFormatConversionApi->v1_technology_format_post: #{e}"
end
```

#### Using the v1_technology_format_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyFormatPost200Response>, Integer, Hash)> v1_technology_format_post_with_http_info(v1_technology_format_post_request)

```ruby
begin
  # Convert Format
  data, status_code, headers = api_instance.v1_technology_format_post_with_http_info(v1_technology_format_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyFormatPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DataFormatConversionApi->v1_technology_format_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_technology_format_post_request** | [**V1TechnologyFormatPostRequest**](V1TechnologyFormatPostRequest.md) |  |  |

### Return type

[**V1TechnologyFormatPost200Response**](V1TechnologyFormatPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

