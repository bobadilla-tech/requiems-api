# OpenapiClient::PaymentsIntelligenceApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_systems_payment_validate_post**](PaymentsIntelligenceApi.md#v1_systems_payment_validate_post) | **POST** /v1/systems/payment/validate | Validate Payment |
| [**v1_systems_transaction_risk_post**](PaymentsIntelligenceApi.md#v1_systems_transaction_risk_post) | **POST** /v1/systems/transaction/risk | Score Transaction Risk |


## v1_systems_payment_validate_post

> <V1SystemsPaymentValidatePost200Response> v1_systems_payment_validate_post(v1_systems_payment_validate_post_request)

Validate Payment

Validate and cross-check BIN, IBAN, and SWIFT identifiers in one call. Returns per-field results and a consistency check that flags country and bank mismatches.

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

api_instance = OpenapiClient::PaymentsIntelligenceApi.new
v1_systems_payment_validate_post_request = OpenapiClient::V1SystemsPaymentValidatePostRequest.new # V1SystemsPaymentValidatePostRequest | 

begin
  # Validate Payment
  result = api_instance.v1_systems_payment_validate_post(v1_systems_payment_validate_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentsIntelligenceApi->v1_systems_payment_validate_post: #{e}"
end
```

#### Using the v1_systems_payment_validate_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsPaymentValidatePost200Response>, Integer, Hash)> v1_systems_payment_validate_post_with_http_info(v1_systems_payment_validate_post_request)

```ruby
begin
  # Validate Payment
  data, status_code, headers = api_instance.v1_systems_payment_validate_post_with_http_info(v1_systems_payment_validate_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsPaymentValidatePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentsIntelligenceApi->v1_systems_payment_validate_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_payment_validate_post_request** | [**V1SystemsPaymentValidatePostRequest**](V1SystemsPaymentValidatePostRequest.md) |  |  |

### Return type

[**V1SystemsPaymentValidatePost200Response**](V1SystemsPaymentValidatePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_systems_transaction_risk_post

> <V1SystemsTransactionRiskPost200Response> v1_systems_transaction_risk_post(v1_systems_transaction_risk_post_request)

Score Transaction Risk

Score a transaction for fraud risk by cross-checking the card BIN country against IP geolocation and billing country. Detects VPN, proxy, and TOR use.

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

api_instance = OpenapiClient::PaymentsIntelligenceApi.new
v1_systems_transaction_risk_post_request = OpenapiClient::V1SystemsTransactionRiskPostRequest.new({card_bin: '424242', ip_address: '92.168.1.1'}) # V1SystemsTransactionRiskPostRequest | 

begin
  # Score Transaction Risk
  result = api_instance.v1_systems_transaction_risk_post(v1_systems_transaction_risk_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentsIntelligenceApi->v1_systems_transaction_risk_post: #{e}"
end
```

#### Using the v1_systems_transaction_risk_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsTransactionRiskPost200Response>, Integer, Hash)> v1_systems_transaction_risk_post_with_http_info(v1_systems_transaction_risk_post_request)

```ruby
begin
  # Score Transaction Risk
  data, status_code, headers = api_instance.v1_systems_transaction_risk_post_with_http_info(v1_systems_transaction_risk_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsTransactionRiskPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentsIntelligenceApi->v1_systems_transaction_risk_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_transaction_risk_post_request** | [**V1SystemsTransactionRiskPostRequest**](V1SystemsTransactionRiskPostRequest.md) |  |  |

### Return type

[**V1SystemsTransactionRiskPost200Response**](V1SystemsTransactionRiskPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

