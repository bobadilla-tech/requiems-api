# OpenapiClient::PhoneValidationApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_validation_phone_batch_post**](PhoneValidationApi.md#v1_validation_phone_batch_post) | **POST** /v1/validation/phone/batch | Batch Validate Phone Numbers |
| [**v1_validation_phone_get**](PhoneValidationApi.md#v1_validation_phone_get) | **GET** /v1/validation/phone | Validate Phone Number |


## v1_validation_phone_batch_post

> <V1FinanceIbanBatchPost200Response> v1_validation_phone_batch_post(v1_validation_phone_batch_post_request)

Batch Validate Phone Numbers

Validates up to 50 phone numbers in a single request. Results are returned in the same order as the input.

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

api_instance = OpenapiClient::PhoneValidationApi.new
v1_validation_phone_batch_post_request = OpenapiClient::V1ValidationPhoneBatchPostRequest.new({numbers: ["+447400123456",  "+12015551234"]}) # V1ValidationPhoneBatchPostRequest | 

begin
  # Batch Validate Phone Numbers
  result = api_instance.v1_validation_phone_batch_post(v1_validation_phone_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PhoneValidationApi->v1_validation_phone_batch_post: #{e}"
end
```

#### Using the v1_validation_phone_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceIbanBatchPost200Response>, Integer, Hash)> v1_validation_phone_batch_post_with_http_info(v1_validation_phone_batch_post_request)

```ruby
begin
  # Batch Validate Phone Numbers
  data, status_code, headers = api_instance.v1_validation_phone_batch_post_with_http_info(v1_validation_phone_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceIbanBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PhoneValidationApi->v1_validation_phone_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_validation_phone_batch_post_request** | [**V1ValidationPhoneBatchPostRequest**](V1ValidationPhoneBatchPostRequest.md) |  |  |

### Return type

[**V1FinanceIbanBatchPost200Response**](V1FinanceIbanBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_validation_phone_get

> <V1ValidationPhoneGet200Response> v1_validation_phone_get(number)

Validate Phone Number

Validates a single phone number and returns its country, type, formatted representation, carrier, and VOIP/virtual risk flags.

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

api_instance = OpenapiClient::PhoneValidationApi.new
number = '+447400123456' # String | The phone number to validate. Must include the country calling code (e.g. +12015551234).

begin
  # Validate Phone Number
  result = api_instance.v1_validation_phone_get(number)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PhoneValidationApi->v1_validation_phone_get: #{e}"
end
```

#### Using the v1_validation_phone_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1ValidationPhoneGet200Response>, Integer, Hash)> v1_validation_phone_get_with_http_info(number)

```ruby
begin
  # Validate Phone Number
  data, status_code, headers = api_instance.v1_validation_phone_get_with_http_info(number)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1ValidationPhoneGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PhoneValidationApi->v1_validation_phone_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **number** | **String** | The phone number to validate. Must include the country calling code (e.g. +12015551234). |  |

### Return type

[**V1ValidationPhoneGet200Response**](V1ValidationPhoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

