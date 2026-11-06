# OpenapiClient::WorkingDaysApi

All URIs are relative to *https://api.requiems.xyz*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**v1_places_working_days_get**](WorkingDaysApi.md#v1_places_working_days_get) | **GET** /v1/places/working-days | Calculate Working Days |


## v1_places_working_days_get

> <V1PlacesWorkingDaysGet200Response> v1_places_working_days_get(from, to, opts)

Calculate Working Days

Calculate the number of working days between two dates, optionally accounting for country-specific holidays

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

api_instance = OpenapiClient::WorkingDaysApi.new
from = '2024-02-23' # String | Start date in YYYY-MM-DD format (ISO 8601)
to = '2024-02-28' # String | End date in YYYY-MM-DD format (ISO 8601). Must be >= from date.
opts = {
  country: 'US', # String | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"FR\"). When provided, country-specific holidays are excluded from working days count.
  subdivision: 'NY' # String | ISO 3166-2 subdivision code for state/region within the country (e.g., \"NY\" for New York, \"CA\" for California). Only used when country is provided.
}

begin
  # Calculate Working Days
  result = api_instance.v1_places_working_days_get(from, to, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling WorkingDaysApi->v1_places_working_days_get: #{e}"
end
```

#### Using the v1_places_working_days_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1PlacesWorkingDaysGet200Response>, Integer, Hash)> v1_places_working_days_get_with_http_info(from, to, opts)

```ruby
begin
  # Calculate Working Days
  data, status_code, headers = api_instance.v1_places_working_days_get_with_http_info(from, to, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1PlacesWorkingDaysGet200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling WorkingDaysApi->v1_places_working_days_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **from** | **String** | Start date in YYYY-MM-DD format (ISO 8601) |  |
| **to** | **String** | End date in YYYY-MM-DD format (ISO 8601). Must be &gt;&#x3D; from date. |  |
| **country** | **String** | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;FR\&quot;). When provided, country-specific holidays are excluded from working days count. | [optional] |
| **subdivision** | **String** | ISO 3166-2 subdivision code for state/region within the country (e.g., \&quot;NY\&quot; for New York, \&quot;CA\&quot; for California). Only used when country is provided. | [optional] |

### Return type

[**V1PlacesWorkingDaysGet200Response**](V1PlacesWorkingDaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

