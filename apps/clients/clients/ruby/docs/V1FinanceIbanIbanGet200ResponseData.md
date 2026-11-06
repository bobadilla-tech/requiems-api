# OpenapiClient::V1FinanceIbanIbanGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **iban** | **String** | The normalised IBAN (spaces stripped, uppercased) | [optional] |
| **valid** | **Boolean** | true if the IBAN passed length and ISO 13616 checksum validation | [optional] |
| **country** | **String** | Full country name (empty if the country code is not in the registry) | [optional] |
| **bank_code** | **String** | Bank identifier extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] |
| **account** | **String** | Account number extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceIbanIbanGet200ResponseData.new(
  iban: null,
  valid: null,
  country: null,
  bank_code: null,
  account: null
)
```

