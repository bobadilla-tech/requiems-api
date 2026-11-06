# OpenapiClient::V1PlacesWorkingDaysGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **working_days** | **Integer** | Number of working days between the two dates (excluding weekends and optionally holidays) | [optional] |
| **from** | **String** | Start date (echoed from request) | [optional] |
| **to** | **String** | End date (echoed from request) | [optional] |
| **country** | **String** | Country code (echoed from request, empty string if not provided) | [optional] |
| **subdivision** | **String** | Subdivision code (echoed from request, empty string if not provided) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesWorkingDaysGet200ResponseData.new(
  working_days: null,
  from: null,
  to: null,
  country: null,
  subdivision: null
)
```

