# V1SystemsPaymentValidatePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bin** | **string** | BIN result. Null when bin was not provided. Contains valid, scheme, card_type, card_level, country_code, issuer, prepaid, luhn. | [optional] [default to undefined]
**iban** | **string** | IBAN result. Null when iban was not provided. Contains valid, country_code, bank_code, account_number. | [optional] [default to undefined]
**swift** | **string** | SWIFT result. Null when swift was not provided. Contains valid, institution, country, branch. | [optional] [default to undefined]
**consistency_ok** | **boolean** | True when no cross-field mismatches were detected | [optional] [default to undefined]
**consistency_flags** | **Array&lt;any&gt;** | Mismatch flags. Possible values: country_mismatch_bin_iban, country_mismatch_bin_swift, country_mismatch_iban_swift, bank_mismatch_iban_swift | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsPaymentValidatePost200ResponseData } from 'api-client';

const instance: V1SystemsPaymentValidatePost200ResponseData = {
    bin,
    iban,
    swift,
    consistency_ok,
    consistency_flags,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
