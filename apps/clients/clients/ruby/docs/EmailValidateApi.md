# OpenapiClient::EmailValidateApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_validation_email_batch_post**](EmailValidateApi.md#v1_validation_email_batch_post) | **POST** /v1/validation/email/batch | Validate Emails (Batch) |
| [**v1_validation_email_post**](EmailValidateApi.md#v1_validation_email_post) | **POST** /v1/validation/email | Validate Email |


## v1_validation_email_batch_post

> <V1ValidationEmailBatchPost200Response> v1_validation_email_batch_post(v1_validation_email_batch_post_request)

Validate Emails (Batch)

Validates up to 50 email addresses in a single request. Each email is processed independently and returns a full validation breakdown (syntax, MX record, disposable check, normalization, and typo suggestion). Invalid emails do not fail the request. Billing: 1 credit per email.

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

api_instance = OpenapiClient::EmailValidateApi.new
v1_validation_email_batch_post_request = OpenapiClient::V1ValidationEmailBatchPostRequest.new({emails: ["user@gmail.com",  "user@gmial.com"]}) # V1ValidationEmailBatchPostRequest | 

begin
  # Validate Emails (Batch)
  result = api_instance.v1_validation_email_batch_post(v1_validation_email_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailValidateApi->v1_validation_email_batch_post: #{e}"
end
```

#### Using the v1_validation_email_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1ValidationEmailBatchPost200Response>, Integer, Hash)> v1_validation_email_batch_post_with_http_info(v1_validation_email_batch_post_request)

```ruby
begin
  # Validate Emails (Batch)
  data, status_code, headers = api_instance.v1_validation_email_batch_post_with_http_info(v1_validation_email_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1ValidationEmailBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailValidateApi->v1_validation_email_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_validation_email_batch_post_request** | [**V1ValidationEmailBatchPostRequest**](V1ValidationEmailBatchPostRequest.md) |  |  |

### Return type

[**V1ValidationEmailBatchPost200Response**](V1ValidationEmailBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_validation_email_post

> <V1ValidationEmailPost200Response> v1_validation_email_post(v1_validation_email_post_request)

Validate Email

Validates a single email address and returns a full breakdown of syntax validity, MX record status, disposable domain check, normalized form, and any typo suggestion.

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

api_instance = OpenapiClient::EmailValidateApi.new
v1_validation_email_post_request = OpenapiClient::V1ValidationEmailPostRequest.new({email: 'user@gmial.com'}) # V1ValidationEmailPostRequest | 

begin
  # Validate Email
  result = api_instance.v1_validation_email_post(v1_validation_email_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailValidateApi->v1_validation_email_post: #{e}"
end
```

#### Using the v1_validation_email_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1ValidationEmailPost200Response>, Integer, Hash)> v1_validation_email_post_with_http_info(v1_validation_email_post_request)

```ruby
begin
  # Validate Email
  data, status_code, headers = api_instance.v1_validation_email_post_with_http_info(v1_validation_email_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1ValidationEmailPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling EmailValidateApi->v1_validation_email_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_validation_email_post_request** | [**V1ValidationEmailPostRequest**](V1ValidationEmailPostRequest.md) |  |  |

### Return type

[**V1ValidationEmailPost200Response**](V1ValidationEmailPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

