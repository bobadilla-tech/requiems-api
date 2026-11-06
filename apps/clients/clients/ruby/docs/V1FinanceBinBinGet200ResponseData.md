# OpenapiClient::V1FinanceBinBinGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **bin** | **String** | The normalised BIN prefix used for the lookup | [optional] |
| **scheme** | **String** | Card network: visa, mastercard, amex, discover, jcb, diners, unionpay, maestro, mir, rupay, private_label | [optional] |
| **card_type** | **String** | credit, debit, prepaid, or charge | [optional] |
| **card_level** | **String** | classic, gold, platinum, infinite, business, signature, or standard | [optional] |
| **issuer_name** | **String** | Name of the card-issuing bank | [optional] |
| **issuer_url** | **String** | Bank website URL | [optional] |
| **issuer_phone** | **String** | Bank customer service phone number | [optional] |
| **country_code** | **String** | ISO 3166-1 alpha-2 country code of the issuing bank (e.g. US, GB, DE) | [optional] |
| **country_name** | **String** | Full country name of the issuing bank | [optional] |
| **prepaid** | **Boolean** | Whether the card is a prepaid card | [optional] |
| **luhn** | **Boolean** | Whether the BIN prefix passes the Luhn algorithm check | [optional] |
| **confidence** | **Float** | Data quality score (0.00–1.00). Multi-source confirmed records score higher. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceBinBinGet200ResponseData.new(
  bin: null,
  scheme: null,
  card_type: null,
  card_level: null,
  issuer_name: null,
  issuer_url: null,
  issuer_phone: null,
  country_code: null,
  country_name: null,
  prepaid: null,
  luhn: null,
  confidence: null
)
```

