# OpenapiClient::MortgageApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_finance_mortgage_get**](MortgageApi.md#v1_finance_mortgage_get) | **GET** /v1/finance/mortgage | Calculate Mortgage |


## v1_finance_mortgage_get

> <V1FinanceMortgageGet200Response> v1_finance_mortgage_get(principal, rate, years)

Calculate Mortgage

Returns the monthly payment, total cost, and full amortization schedule for a fixed-rate mortgage.

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

api_instance = OpenapiClient::MortgageApi.new
principal = 300000 # Float | Loan amount in your chosen currency (e.g. 300000 for $300,000)
rate = 6.5 # Float | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0.
years = 30 # Integer | Loan term in years (1–50)

begin
  # Calculate Mortgage
  result = api_instance.v1_finance_mortgage_get(principal, rate, years)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MortgageApi->v1_finance_mortgage_get: #{e}"
end
```

#### Using the v1_finance_mortgage_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1FinanceMortgageGet200Response>, Integer, Hash)> v1_finance_mortgage_get_with_http_info(principal, rate, years)

```ruby
begin
  # Calculate Mortgage
  data, status_code, headers = api_instance.v1_finance_mortgage_get_with_http_info(principal, rate, years)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1FinanceMortgageGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MortgageApi->v1_finance_mortgage_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **principal** | **Float** | Loan amount in your chosen currency (e.g. 300000 for $300,000) |  |
| **rate** | **Float** | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0. |  |
| **years** | **Integer** | Loan term in years (1–50) |  |

### Return type

[**V1FinanceMortgageGet200Response**](V1FinanceMortgageGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

