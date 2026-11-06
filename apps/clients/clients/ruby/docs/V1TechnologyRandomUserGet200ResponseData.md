# OpenapiClient::V1TechnologyRandomUserGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **name** | **String** | Full name of the generated user | [optional] |
| **email** | **String** | Email address of the generated user | [optional] |
| **phone** | **String** | Phone number of the generated user | [optional] |
| **address_street** | **String** | Street address | [optional] |
| **address_city** | **String** | City name | [optional] |
| **address_state** | **String** | State or region | [optional] |
| **address_zip** | **String** | Postal / ZIP code | [optional] |
| **address_country** | **String** | Country name | [optional] |
| **avatar** | **String** | URL to a unique identicon avatar for the generated user (DiceBear) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyRandomUserGet200ResponseData.new(
  name: null,
  email: null,
  phone: null,
  address_street: null,
  address_city: null,
  address_state: null,
  address_zip: null,
  address_country: null,
  avatar: null
)
```

