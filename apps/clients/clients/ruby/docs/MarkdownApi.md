# OpenapiClient::MarkdownApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_markdown_batch_post**](MarkdownApi.md#v1_technology_markdown_batch_post) | **POST** /v1/technology/markdown/batch | Convert Markdown Batch |
| [**v1_technology_markdown_post**](MarkdownApi.md#v1_technology_markdown_post) | **POST** /v1/technology/markdown | Convert Markdown to HTML |


## v1_technology_markdown_batch_post

> <V1TechnologyMarkdownBatchPost200Response> v1_technology_markdown_batch_post(v1_technology_markdown_batch_post_request)

Convert Markdown Batch

Converts multiple Markdown strings into HTML in a single request. Results are returned in the same order as the input array. Supports optional sanitization to remove unsafe HTML tags such as script and iframe.

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

api_instance = OpenapiClient::MarkdownApi.new
v1_technology_markdown_batch_post_request = OpenapiClient::V1TechnologyMarkdownBatchPostRequest.new({markdowns: ["# Hello",  "**bold**",  "- item one"]}) # V1TechnologyMarkdownBatchPostRequest | 

begin
  # Convert Markdown Batch
  result = api_instance.v1_technology_markdown_batch_post(v1_technology_markdown_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MarkdownApi->v1_technology_markdown_batch_post: #{e}"
end
```

#### Using the v1_technology_markdown_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyMarkdownBatchPost200Response>, Integer, Hash)> v1_technology_markdown_batch_post_with_http_info(v1_technology_markdown_batch_post_request)

```ruby
begin
  # Convert Markdown Batch
  data, status_code, headers = api_instance.v1_technology_markdown_batch_post_with_http_info(v1_technology_markdown_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyMarkdownBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MarkdownApi->v1_technology_markdown_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_technology_markdown_batch_post_request** | [**V1TechnologyMarkdownBatchPostRequest**](V1TechnologyMarkdownBatchPostRequest.md) |  |  |

### Return type

[**V1TechnologyMarkdownBatchPost200Response**](V1TechnologyMarkdownBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_technology_markdown_post

> <V1TechnologyMarkdownPost200Response> v1_technology_markdown_post(v1_technology_markdown_post_request)

Convert Markdown to HTML

Converts a Markdown string to HTML. Pass sanitize true to strip potentially unsafe tags like script and iframe from the output.

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

api_instance = OpenapiClient::MarkdownApi.new
v1_technology_markdown_post_request = OpenapiClient::V1TechnologyMarkdownPostRequest.new({markdown: '# Hello

This is **bold** and _italic_ text.'}) # V1TechnologyMarkdownPostRequest | 

begin
  # Convert Markdown to HTML
  result = api_instance.v1_technology_markdown_post(v1_technology_markdown_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MarkdownApi->v1_technology_markdown_post: #{e}"
end
```

#### Using the v1_technology_markdown_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyMarkdownPost200Response>, Integer, Hash)> v1_technology_markdown_post_with_http_info(v1_technology_markdown_post_request)

```ruby
begin
  # Convert Markdown to HTML
  data, status_code, headers = api_instance.v1_technology_markdown_post_with_http_info(v1_technology_markdown_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyMarkdownPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MarkdownApi->v1_technology_markdown_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_technology_markdown_post_request** | [**V1TechnologyMarkdownPostRequest**](V1TechnologyMarkdownPostRequest.md) |  |  |

### Return type

[**V1TechnologyMarkdownPost200Response**](V1TechnologyMarkdownPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

