# OpenapiClient::PasswordGeneratorApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_password_get**](PasswordGeneratorApi.md#v1_technology_password_get) | **GET** /v1/technology/password | Generate Password |


## v1_technology_password_get

> <V1TechnologyPasswordGet200Response> v1_technology_password_get(opts)

Generate Password

Generate a cryptographically secure random password with customizable character sets and length

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

api_instance = OpenapiClient::PasswordGeneratorApi.new
opts = {
  length: 16, # Integer | Password length (8-128 characters)
  uppercase: true, # Boolean | Include uppercase letters (A-Z)
  numbers: true, # Boolean | Include numbers (0-9)
  symbols: true # Boolean | Include special characters (!@#$%^&*()-_=+[]{}|;:,.<>?)
}

begin
  # Generate Password
  result = api_instance.v1_technology_password_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PasswordGeneratorApi->v1_technology_password_get: #{e}"
end
```

#### Using the v1_technology_password_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyPasswordGet200Response>, Integer, Hash)> v1_technology_password_get_with_http_info(opts)

```ruby
begin
  # Generate Password
  data, status_code, headers = api_instance.v1_technology_password_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyPasswordGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PasswordGeneratorApi->v1_technology_password_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **length** | **Integer** | Password length (8-128 characters) | [optional] |
| **uppercase** | **Boolean** | Include uppercase letters (A-Z) | [optional] |
| **numbers** | **Boolean** | Include numbers (0-9) | [optional] |
| **symbols** | **Boolean** | Include special characters (!@#$%^&amp;*()-_&#x3D;+[]{}|;:,.&lt;&gt;?) | [optional] |

### Return type

[**V1TechnologyPasswordGet200Response**](V1TechnologyPasswordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

