# V1FinanceBinBinGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bin** | **string** | The normalised BIN prefix used for the lookup | [optional] [default to undefined]
**scheme** | **string** | Card network: visa, mastercard, amex, discover, jcb, diners, unionpay, maestro, mir, rupay, private_label | [optional] [default to undefined]
**card_type** | **string** | credit, debit, prepaid, or charge | [optional] [default to undefined]
**card_level** | **string** | classic, gold, platinum, infinite, business, signature, or standard | [optional] [default to undefined]
**issuer_name** | **string** | Name of the card-issuing bank | [optional] [default to undefined]
**issuer_url** | **string** | Bank website URL | [optional] [default to undefined]
**issuer_phone** | **string** | Bank customer service phone number | [optional] [default to undefined]
**country_code** | **string** | ISO 3166-1 alpha-2 country code of the issuing bank (e.g. US, GB, DE) | [optional] [default to undefined]
**country_name** | **string** | Full country name of the issuing bank | [optional] [default to undefined]
**prepaid** | **boolean** | Whether the card is a prepaid card | [optional] [default to undefined]
**luhn** | **boolean** | Whether the BIN prefix passes the Luhn algorithm check | [optional] [default to undefined]
**confidence** | **number** | Data quality score (0.00–1.00). Multi-source confirmed records score higher. | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceBinBinGet200ResponseData } from 'api-client';

const instance: V1FinanceBinBinGet200ResponseData = {
    bin,
    scheme,
    card_type,
    card_level,
    issuer_name,
    issuer_url,
    issuer_phone,
    country_code,
    country_name,
    prepaid,
    luhn,
    confidence,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
