# OpenapiClient::V1SystemsPaymentValidatePostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **bin** | **String** | Card BIN (first 6–8 digits). At least one of bin, iban, or swift is required. | [optional] |
| **iban** | **String** | IBAN string to validate. Checksum and format are verified. | [optional] |
| **swift** | **String** | SWIFT/BIC code to look up institution and country. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsPaymentValidatePostRequest.new(
  bin: 424242,
  iban: DE89370400440532013000,
  swift: DEUTDEDB
)
```

