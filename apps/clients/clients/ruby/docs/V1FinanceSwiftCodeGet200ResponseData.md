# OpenapiClient::V1FinanceSwiftCodeGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **swift_code** | **String** | Full 11-character SWIFT/BIC code | [optional] |
| **bank_code** | **String** | Institution code (characters 1-4) | [optional] |
| **country_code** | **String** | ISO 3166-1 alpha-2 country code (characters 5-6) | [optional] |
| **location_code** | **String** | Location code (characters 7-8) | [optional] |
| **branch_code** | **String** | Branch code (characters 9-11), XXX for primary office | [optional] |
| **bank_name** | **String** | Bank or institution name | [optional] |
| **city** | **String** | City of the branch or primary office | [optional] |
| **country_name** | **String** | Full country name | [optional] |
| **is_primary** | **Boolean** | true when branch_code is XXX | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceSwiftCodeGet200ResponseData.new(
  swift_code: null,
  bank_code: null,
  country_code: null,
  location_code: null,
  branch_code: null,
  bank_name: null,
  city: null,
  country_name: null,
  is_primary: null
)
```

