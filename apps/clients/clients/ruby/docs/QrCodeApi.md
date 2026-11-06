# OpenapiClient::QrCodeApi

All URIs are relative to *https://requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_technology_qr_base64_batch_post**](QrCodeApi.md#v1_technology_qr_base64_batch_post) | **POST** /v1/technology/qr/base64/batch | Batch Generate QR Codes (Base64) |
| [**v1_technology_qr_base64_get**](QrCodeApi.md#v1_technology_qr_base64_get) | **GET** /v1/technology/qr/base64 | Generate QR Code (Base64 JSON) |
| [**v1_technology_qr_get**](QrCodeApi.md#v1_technology_qr_get) | **GET** /v1/technology/qr | Generate QR Code (PNG) |


## v1_technology_qr_base64_batch_post

> <V1TechnologyQrBase64BatchPost200Response> v1_technology_qr_base64_batch_post(v1_technology_qr_base64_batch_post_request)

Batch Generate QR Codes (Base64)

Generate up to 50 base64-encoded QR codes in a single request. Results are returned in input order. The PNG endpoint has no batch variant — use this endpoint for batch generation.

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

api_instance = OpenapiClient::QrCodeApi.new
v1_technology_qr_base64_batch_post_request = OpenapiClient::V1TechnologyQrBase64BatchPostRequest.new({items: 'items_example', items_data: 'items_data_example'}) # V1TechnologyQrBase64BatchPostRequest | 

begin
  # Batch Generate QR Codes (Base64)
  result = api_instance.v1_technology_qr_base64_batch_post(v1_technology_qr_base64_batch_post_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling QrCodeApi->v1_technology_qr_base64_batch_post: #{e}"
end
```

#### Using the v1_technology_qr_base64_batch_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyQrBase64BatchPost200Response>, Integer, Hash)> v1_technology_qr_base64_batch_post_with_http_info(v1_technology_qr_base64_batch_post_request)

```ruby
begin
  # Batch Generate QR Codes (Base64)
  data, status_code, headers = api_instance.v1_technology_qr_base64_batch_post_with_http_info(v1_technology_qr_base64_batch_post_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyQrBase64BatchPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling QrCodeApi->v1_technology_qr_base64_batch_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **v1_technology_qr_base64_batch_post_request** | [**V1TechnologyQrBase64BatchPostRequest**](V1TechnologyQrBase64BatchPostRequest.md) |  |  |

### Return type

[**V1TechnologyQrBase64BatchPost200Response**](V1TechnologyQrBase64BatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## v1_technology_qr_base64_get

> <V1TechnologyQrBase64Get200Response> v1_technology_qr_base64_get(data, opts)

Generate QR Code (Base64 JSON)

Returns a JSON envelope containing the QR code as a base64-encoded PNG string, along with its dimensions.

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

api_instance = OpenapiClient::QrCodeApi.new
data = 'https://example.com' # String | The text or URL to encode in the QR code
opts = {
  size: 200, # Integer | Image size in pixels (default: 256, min: 50, max: 1000)
  recovery: 'highest' # String | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium
}

begin
  # Generate QR Code (Base64 JSON)
  result = api_instance.v1_technology_qr_base64_get(data, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling QrCodeApi->v1_technology_qr_base64_get: #{e}"
end
```

#### Using the v1_technology_qr_base64_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1TechnologyQrBase64Get200Response>, Integer, Hash)> v1_technology_qr_base64_get_with_http_info(data, opts)

```ruby
begin
  # Generate QR Code (Base64 JSON)
  data, status_code, headers = api_instance.v1_technology_qr_base64_get_with_http_info(data, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1TechnologyQrBase64Get200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling QrCodeApi->v1_technology_qr_base64_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data** | **String** | The text or URL to encode in the QR code |  |
| **size** | **Integer** | Image size in pixels (default: 256, min: 50, max: 1000) | [optional] |
| **recovery** | **String** | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium | [optional] |

### Return type

[**V1TechnologyQrBase64Get200Response**](V1TechnologyQrBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_technology_qr_get

> v1_technology_qr_get(data, opts)

Generate QR Code (PNG)

Returns a raw PNG image of the QR code. Ideal for direct embedding or file download.

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

api_instance = OpenapiClient::QrCodeApi.new
data = 'https://example.com' # String | The text or URL to encode in the QR code
opts = {
  size: 200, # Integer | Image size in pixels (default: 256, min: 50, max: 1000)
  recovery: 'high' # String | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium
}

begin
  # Generate QR Code (PNG)
  api_instance.v1_technology_qr_get(data, opts)
rescue OpenapiClient::ApiError => e
  puts "Error when calling QrCodeApi->v1_technology_qr_get: #{e}"
end
```

#### Using the v1_technology_qr_get_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> v1_technology_qr_get_with_http_info(data, opts)

```ruby
begin
  # Generate QR Code (PNG)
  data, status_code, headers = api_instance.v1_technology_qr_get_with_http_info(data, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling QrCodeApi->v1_technology_qr_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data** | **String** | The text or URL to encode in the QR code |  |
| **size** | **Integer** | Image size in pixels (default: 256, min: 50, max: 1000) | [optional] |
| **recovery** | **String** | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium | [optional] |

### Return type

nil (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

