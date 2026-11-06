# OpenapiClient::SentimentApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_text_sentiment_post**](SentimentApi.md#v1_text_sentiment_post) | **POST** /v1/text/sentiment | Analyze Sentiment |


## v1_text_sentiment_post

> <V1TextSentimentPost200Response> v1_text_sentiment_post(v1_text_sentiment_post_request)

Analyze Sentiment

Analyzes the sentiment of the provided text and returns a classification, confidence score, and a full breakdown across all three sentiment classes.

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

api_instance = OpenapiClient::SentimentApi.new
v1_text_sentiment_post_request = OpenapiClient::V1TextSentimentPostRequest.new({text: 'I absolutely love this product, it exceeded my expectations!'}) # V1TextSentimentPostRequest | 

begin
  # Analyze Sentiment
  result = api_instance.v1_text_sentiment_post(v1_text_sentiment_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling SentimentApi->v1_text_sentiment_post: #{e}"
end
```

#### Using the v1_text_sentiment_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TextSentimentPost200Response>, Integer, Hash)> v1_text_sentiment_post_with_http_info(v1_text_sentiment_post_request)

```ruby
begin
  # Analyze Sentiment
  data, status_code, headers = api_instance.v1_text_sentiment_post_with_http_info(v1_text_sentiment_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TextSentimentPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling SentimentApi->v1_text_sentiment_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_text_sentiment_post_request** | [**V1TextSentimentPostRequest**](V1TextSentimentPostRequest.md) |  |  |

### Return type

[**V1TextSentimentPost200Response**](V1TextSentimentPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

