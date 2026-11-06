# OpenapiClient::IbanApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_iban_batch_post**](IbanApi.md#v1_finance_iban_batch_post) | **POST** /v1/finance/iban/batch | Batch Validate IBANs |
| [**v1_finance_iban_iban_get**](IbanApi.md#v1_finance_iban_iban_get) | **GET** /v1/finance/iban/{iban} | Validate IBAN |


## v1_finance_iban_batch_post

> <V1FinanceIbanBatchPost200Response> v1_finance_iban_batch_post(v1_finance_iban_batch_post_request)

Batch Validate IBANs

Validates up to 50 iban numbers in a single request. Results are returned in the same order as the input.

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

api_instance = OpenapiClient::IbanApi.new
v1_finance_iban_batch_post_request = OpenapiClient::V1FinanceIbanBatchPostRequest.new({numbers: ["GB29NWBK60161331926819",  "DE89370400440532013000"]}) # V1FinanceIbanBatchPostRequest | 

begin
  # Batch Validate IBANs
  result = api_instance.v1_finance_iban_batch_post(v1_finance_iban_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IbanApi->v1_finance_iban_batch_post: #{e}"
end
```

#### Using the v1_finance_iban_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceIbanBatchPost200Response>, Integer, Hash)> v1_finance_iban_batch_post_with_http_info(v1_finance_iban_batch_post_request)

```ruby
begin
  # Batch Validate IBANs
  data, status_code, headers = api_instance.v1_finance_iban_batch_post_with_http_info(v1_finance_iban_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceIbanBatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IbanApi->v1_finance_iban_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_finance_iban_batch_post_request** | [**V1FinanceIbanBatchPostRequest**](V1FinanceIbanBatchPostRequest.md) |  |  |

### Return type

[**V1FinanceIbanBatchPost200Response**](V1FinanceIbanBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_finance_iban_iban_get

> <V1FinanceIbanIbanGet200Response> v1_finance_iban_iban_get(iban)

Validate IBAN

Validates an IBAN and returns the country, bank code, and account number. Spaces in the input are stripped automatically. Always returns HTTP 200 — check the valid field to determine whether the IBAN is valid.

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

api_instance = OpenapiClient::IbanApi.new
iban = 'DE89370400440532013000' # String | The IBAN to validate. Spaces are stripped. Case-insensitive.

begin
  # Validate IBAN
  result = api_instance.v1_finance_iban_iban_get(iban)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IbanApi->v1_finance_iban_iban_get: #{e}"
end
```

#### Using the v1_finance_iban_iban_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceIbanIbanGet200Response>, Integer, Hash)> v1_finance_iban_iban_get_with_http_info(iban)

```ruby
begin
  # Validate IBAN
  data, status_code, headers = api_instance.v1_finance_iban_iban_get_with_http_info(iban)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceIbanIbanGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IbanApi->v1_finance_iban_iban_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **iban** | **String** | The IBAN to validate. Spaces are stripped. Case-insensitive. |  |

### Return type

[**V1FinanceIbanIbanGet200Response**](V1FinanceIbanIbanGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

