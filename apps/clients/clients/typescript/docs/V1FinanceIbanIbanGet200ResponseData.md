# V1FinanceIbanIbanGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**iban** | **string** | The normalised IBAN (spaces stripped, uppercased) | [optional] [default to undefined]
**valid** | **boolean** | true if the IBAN passed length and ISO 13616 checksum validation | [optional] [default to undefined]
**country** | **string** | Full country name (empty if the country code is not in the registry) | [optional] [default to undefined]
**bank_code** | **string** | Bank identifier extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] [default to undefined]
**account** | **string** | Account number extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceIbanIbanGet200ResponseData } from 'api-client';

const instance: V1FinanceIbanIbanGet200ResponseData = {
    iban,
    valid,
    country,
    bank_code,
    account,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
