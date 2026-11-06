# OpenapiClient::V1PlacesPostalBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **String** | List of postal code lookups (1–50 items). |  |
| **items_code** | **String** | The postal code to look up. |  |
| **items_country** | **String** | ISO 3166-1 alpha-2 country code (default: US). | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1PlacesPostalBatchPostRequest.new(
  items: null,
  items_code: null,
  items_country: null
)
```

