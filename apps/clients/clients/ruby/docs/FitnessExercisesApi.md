# OpenapiClient::FitnessExercisesApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_health_body_parts_get**](FitnessExercisesApi.md#v1_health_body_parts_get) | **GET** /v1/health/body-parts | List Body Parts |
| [**v1_health_equipment_get**](FitnessExercisesApi.md#v1_health_equipment_get) | **GET** /v1/health/equipment | List Equipment |
| [**v1_health_exercises_get**](FitnessExercisesApi.md#v1_health_exercises_get) | **GET** /v1/health/exercises | List Exercises |
| [**v1_health_exercises_id_get**](FitnessExercisesApi.md#v1_health_exercises_id_get) | **GET** /v1/health/exercises/{id} | Get Exercise by ID |
| [**v1_health_exercises_random_get**](FitnessExercisesApi.md#v1_health_exercises_random_get) | **GET** /v1/health/exercises/random | Random Exercise |
| [**v1_health_muscles_get**](FitnessExercisesApi.md#v1_health_muscles_get) | **GET** /v1/health/muscles | List Muscles |


## v1_health_body_parts_get

> <V1HealthBodyPartsGet200Response> v1_health_body_parts_get

List Body Parts

Returns a sorted list of all distinct body part values present in the dataset. Use these as valid values for the body_part filter.

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

api_instance = OpenapiClient::FitnessExercisesApi.new

begin
  # List Body Parts
  result = api_instance.v1_health_body_parts_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_body_parts_get: #{e}"
end
```

#### Using the v1_health_body_parts_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1HealthBodyPartsGet200Response>, Integer, Hash)> v1_health_body_parts_get_with_http_info

```ruby
begin
  # List Body Parts
  data, status_code, headers = api_instance.v1_health_body_parts_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1HealthBodyPartsGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_body_parts_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1HealthBodyPartsGet200Response**](V1HealthBodyPartsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_health_equipment_get

> <V1HealthEquipmentGet200Response> v1_health_equipment_get

List Equipment

Returns a sorted list of all distinct equipment types. Use these as valid values for the equipment filter.

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

api_instance = OpenapiClient::FitnessExercisesApi.new

begin
  # List Equipment
  result = api_instance.v1_health_equipment_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_equipment_get: #{e}"
end
```

#### Using the v1_health_equipment_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1HealthEquipmentGet200Response>, Integer, Hash)> v1_health_equipment_get_with_http_info

```ruby
begin
  # List Equipment
  data, status_code, headers = api_instance.v1_health_equipment_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1HealthEquipmentGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_equipment_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1HealthEquipmentGet200Response**](V1HealthEquipmentGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_health_exercises_get

> <V1HealthExercisesGet200Response> v1_health_exercises_get(opts)

List Exercises

Returns a paginated list of exercises. All filter parameters are optional and combinable.

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

api_instance = OpenapiClient::FitnessExercisesApi.new
opts = {
  body_part: 'chest', # String | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values.
  equipment: 'barbell', # String | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values.
  muscle: 'biceps', # String | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values.
  search: 'bench press', # String | Full-text search on exercise name.
  page: 1, # Integer | Page number (default: 1)
  per_page: 20 # Integer | Results per page, 1–100 (default: 20)
}

begin
  # List Exercises
  result = api_instance.v1_health_exercises_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_exercises_get: #{e}"
end
```

#### Using the v1_health_exercises_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1HealthExercisesGet200Response>, Integer, Hash)> v1_health_exercises_get_with_http_info(opts)

```ruby
begin
  # List Exercises
  data, status_code, headers = api_instance.v1_health_exercises_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1HealthExercisesGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_exercises_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **body_part** | **String** | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. | [optional] |
| **equipment** | **String** | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. | [optional] |
| **muscle** | **String** | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. | [optional] |
| **search** | **String** | Full-text search on exercise name. | [optional] |
| **page** | **Integer** | Page number (default: 1) | [optional] |
| **per_page** | **Integer** | Results per page, 1–100 (default: 20) | [optional] |

### Return type

[**V1HealthExercisesGet200Response**](V1HealthExercisesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_health_exercises_id_get

> <V1HealthExercisesIdGet200Response> v1_health_exercises_id_get(id)

Get Exercise by ID

Returns a single exercise by its numeric ID.

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

api_instance = OpenapiClient::FitnessExercisesApi.new
id = 1 # Integer | Numeric exercise ID

begin
  # Get Exercise by ID
  result = api_instance.v1_health_exercises_id_get(id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_exercises_id_get: #{e}"
end
```

#### Using the v1_health_exercises_id_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1HealthExercisesIdGet200Response>, Integer, Hash)> v1_health_exercises_id_get_with_http_info(id)

```ruby
begin
  # Get Exercise by ID
  data, status_code, headers = api_instance.v1_health_exercises_id_get_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1HealthExercisesIdGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_exercises_id_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** | Numeric exercise ID |  |

### Return type

[**V1HealthExercisesIdGet200Response**](V1HealthExercisesIdGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_health_exercises_random_get

> <V1HealthExercisesRandomGet200Response> v1_health_exercises_random_get(opts)

Random Exercise

Returns a single randomly selected exercise. Accepts the same filter parameters as the list endpoint, so you can get a random chest exercise, a random bodyweight exercise, etc.

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

api_instance = OpenapiClient::FitnessExercisesApi.new
opts = {
  body_part: 'back', # String | Restrict random selection to this body part.
  equipment: 'body weight', # String | Restrict random selection to this equipment type.
  muscle: 'glutes', # String | Restrict random selection to exercises targeting this muscle.
  search: 'curl' # String | Restrict random selection to exercises matching this search term.
}

begin
  # Random Exercise
  result = api_instance.v1_health_exercises_random_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_exercises_random_get: #{e}"
end
```

#### Using the v1_health_exercises_random_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1HealthExercisesRandomGet200Response>, Integer, Hash)> v1_health_exercises_random_get_with_http_info(opts)

```ruby
begin
  # Random Exercise
  data, status_code, headers = api_instance.v1_health_exercises_random_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1HealthExercisesRandomGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_exercises_random_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **body_part** | **String** | Restrict random selection to this body part. | [optional] |
| **equipment** | **String** | Restrict random selection to this equipment type. | [optional] |
| **muscle** | **String** | Restrict random selection to exercises targeting this muscle. | [optional] |
| **search** | **String** | Restrict random selection to exercises matching this search term. | [optional] |

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## v1_health_muscles_get

> <V1HealthMusclesGet200Response> v1_health_muscles_get

List Muscles

Returns a sorted list of all distinct muscle names (combining target and secondary muscles). Use these as valid values for the muscle filter.

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

api_instance = OpenapiClient::FitnessExercisesApi.new

begin
  # List Muscles
  result = api_instance.v1_health_muscles_get
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_muscles_get: #{e}"
end
```

#### Using the v1_health_muscles_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1HealthMusclesGet200Response>, Integer, Hash)> v1_health_muscles_get_with_http_info

```ruby
begin
  # List Muscles
  data, status_code, headers = api_instance.v1_health_muscles_get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1HealthMusclesGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling FitnessExercisesApi->v1_health_muscles_get_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**V1HealthMusclesGet200Response**](V1HealthMusclesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

