# V1SystemsPaymentValidatePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bin** | **string** | Card BIN (first 6–8 digits). At least one of bin, iban, or swift is required. | [optional] [default to undefined]
**iban** | **string** | IBAN string to validate. Checksum and format are verified. | [optional] [default to undefined]
**swift** | **string** | SWIFT/BIC code to look up institution and country. | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsPaymentValidatePostRequest } from 'api-client';

const instance: V1SystemsPaymentValidatePostRequest = {
    bin,
    iban,
    swift,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
