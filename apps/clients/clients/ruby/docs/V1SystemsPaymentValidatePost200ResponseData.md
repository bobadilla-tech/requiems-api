# OpenapiClient::V1SystemsPaymentValidatePost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **bin** | **String** | BIN result. Null when bin was not provided. Contains valid, scheme, card_type, card_level, country_code, issuer, prepaid, luhn. | [optional] |
| **iban** | **String** | IBAN result. Null when iban was not provided. Contains valid, country_code, bank_code, account_number. | [optional] |
| **swift** | **String** | SWIFT result. Null when swift was not provided. Contains valid, institution, country, branch. | [optional] |
| **consistency_ok** | **Boolean** | True when no cross-field mismatches were detected | [optional] |
| **consistency_flags** | **Array&lt;Object&gt;** | Mismatch flags. Possible values: country_mismatch_bin_iban, country_mismatch_bin_swift, country_mismatch_iban_swift, bank_mismatch_iban_swift | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsPaymentValidatePost200ResponseData.new(
  bin: null,
  iban: null,
  swift: null,
  consistency_ok: null,
  consistency_flags: null
)
```

