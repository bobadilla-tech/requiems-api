# OpenapiClient::V1ValidationPhoneGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **number** | **String** | The original number as supplied in the request | [optional] |
| **valid** | **Boolean** | Whether the number is a valid, dialable phone number | [optional] |
| **country** | **String** | ISO 3166-1 alpha-2 country code (omitted when valid is false) | [optional] |
| **type** | **String** | Number type: mobile, landline, landline_or_mobile, toll_free, voip, premium_rate, shared_cost, personal_number, pager, uan, voicemail, or unknown (omitted when valid is false) | [optional] |
| **formatted** | **String** | International format of the number, e.g. +44 7400 123456 (omitted when valid is false) | [optional] |
| **carrier_name** | **String** | Carrier name from phone prefix metadata (omitted when carrier cannot be determined) | [optional] |
| **carrier_source** | **String** | How the carrier was determined. Always \&quot;metadata\&quot; when present | [optional] |
| **risk_is_voip** | **Boolean** | true when the number type is voip | [optional] |
| **risk_is_virtual** | **Boolean** | true when the number is not tied to a physical SIM or fixed line: voip, personal_number, uan, pager, or voicemail | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1ValidationPhoneGet200ResponseData.new(
  number: null,
  valid: null,
  country: null,
  type: null,
  formatted: null,
  carrier_name: null,
  carrier_source: null,
  risk_is_voip: null,
  risk_is_virtual: null
)
```

