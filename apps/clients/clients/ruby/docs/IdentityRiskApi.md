# OpenapiClient::IdentityRiskApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_systems_risk_score_post**](IdentityRiskApi.md#v1_systems_risk_score_post) | **POST** /v1/systems/risk/score | Score Risk |
| [**v1_systems_signup_protect_post**](IdentityRiskApi.md#v1_systems_signup_protect_post) | **POST** /v1/systems/signup/protect | Protect Signup |
| [**v1_systems_user_verify_post**](IdentityRiskApi.md#v1_systems_user_verify_post) | **POST** /v1/systems/user/verify | Verify User |


## v1_systems_risk_score_post

> <V1SystemsRiskScorePost200Response> v1_systems_risk_score_post(v1_systems_risk_score_post_request)

Score Risk

Score a user for risk without the full signal breakdown. Lower latency than /signup/protect, suited for background re-scoring and rate limiting.

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

api_instance = OpenapiClient::IdentityRiskApi.new
v1_systems_risk_score_post_request = OpenapiClient::V1SystemsRiskScorePostRequest.new # V1SystemsRiskScorePostRequest | 

begin
  # Score Risk
  result = api_instance.v1_systems_risk_score_post(v1_systems_risk_score_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IdentityRiskApi->v1_systems_risk_score_post: #{e}"
end
```

#### Using the v1_systems_risk_score_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsRiskScorePost200Response>, Integer, Hash)> v1_systems_risk_score_post_with_http_info(v1_systems_risk_score_post_request)

```ruby
begin
  # Score Risk
  data, status_code, headers = api_instance.v1_systems_risk_score_post_with_http_info(v1_systems_risk_score_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsRiskScorePost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IdentityRiskApi->v1_systems_risk_score_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_risk_score_post_request** | [**V1SystemsRiskScorePostRequest**](V1SystemsRiskScorePostRequest.md) |  |  |

### Return type

[**V1SystemsRiskScorePost200Response**](V1SystemsRiskScorePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_systems_signup_protect_post

> <V1SystemsSignupProtectPost200Response> v1_systems_signup_protect_post(v1_systems_signup_protect_post_request)

Protect Signup

Evaluate a new user at signup. Returns a full risk decision with per-signal breakdown across email, phone, and IP.

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

api_instance = OpenapiClient::IdentityRiskApi.new
v1_systems_signup_protect_post_request = OpenapiClient::V1SystemsSignupProtectPostRequest.new # V1SystemsSignupProtectPostRequest | 

begin
  # Protect Signup
  result = api_instance.v1_systems_signup_protect_post(v1_systems_signup_protect_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IdentityRiskApi->v1_systems_signup_protect_post: #{e}"
end
```

#### Using the v1_systems_signup_protect_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsSignupProtectPost200Response>, Integer, Hash)> v1_systems_signup_protect_post_with_http_info(v1_systems_signup_protect_post_request)

```ruby
begin
  # Protect Signup
  data, status_code, headers = api_instance.v1_systems_signup_protect_post_with_http_info(v1_systems_signup_protect_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsSignupProtectPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IdentityRiskApi->v1_systems_signup_protect_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_signup_protect_post_request** | [**V1SystemsSignupProtectPostRequest**](V1SystemsSignupProtectPostRequest.md) |  |  |

### Return type

[**V1SystemsSignupProtectPost200Response**](V1SystemsSignupProtectPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_systems_user_verify_post

> <V1SystemsUserVerifyPost200Response> v1_systems_user_verify_post(v1_systems_user_verify_post_request)

Verify User

Deep-verify an email address using domain-level signals including WHOIS age, MX records, and domain availability. Optional IP check. Use for high-value or suspicious accounts.

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

api_instance = OpenapiClient::IdentityRiskApi.new
v1_systems_user_verify_post_request = OpenapiClient::V1SystemsUserVerifyPostRequest.new({email: 'user@example.com'}) # V1SystemsUserVerifyPostRequest | 

begin
  # Verify User
  result = api_instance.v1_systems_user_verify_post(v1_systems_user_verify_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling IdentityRiskApi->v1_systems_user_verify_post: #{e}"
end
```

#### Using the v1_systems_user_verify_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1SystemsUserVerifyPost200Response>, Integer, Hash)> v1_systems_user_verify_post_with_http_info(v1_systems_user_verify_post_request)

```ruby
begin
  # Verify User
  data, status_code, headers = api_instance.v1_systems_user_verify_post_with_http_info(v1_systems_user_verify_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1SystemsUserVerifyPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling IdentityRiskApi->v1_systems_user_verify_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_systems_user_verify_post_request** | [**V1SystemsUserVerifyPostRequest**](V1SystemsUserVerifyPostRequest.md) |  |  |

### Return type

[**V1SystemsUserVerifyPost200Response**](V1SystemsUserVerifyPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

