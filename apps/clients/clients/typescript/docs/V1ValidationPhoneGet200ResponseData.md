# V1ValidationPhoneGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**number** | **string** | The original number as supplied in the request | [optional] [default to undefined]
**valid** | **boolean** | Whether the number is a valid, dialable phone number | [optional] [default to undefined]
**country** | **string** | ISO 3166-1 alpha-2 country code (omitted when valid is false) | [optional] [default to undefined]
**type** | **string** | Number type: mobile, landline, landline_or_mobile, toll_free, voip, premium_rate, shared_cost, personal_number, pager, uan, voicemail, or unknown (omitted when valid is false) | [optional] [default to undefined]
**formatted** | **string** | International format of the number, e.g. +44 7400 123456 (omitted when valid is false) | [optional] [default to undefined]
**carrier_name** | **string** | Carrier name from phone prefix metadata (omitted when carrier cannot be determined) | [optional] [default to undefined]
**carrier_source** | **string** | How the carrier was determined. Always \&quot;metadata\&quot; when present | [optional] [default to undefined]
**risk_is_voip** | **boolean** | true when the number type is voip | [optional] [default to undefined]
**risk_is_virtual** | **boolean** | true when the number is not tied to a physical SIM or fixed line: voip, personal_number, uan, pager, or voicemail | [optional] [default to undefined]

## Example

```typescript
import { V1ValidationPhoneGet200ResponseData } from 'api-client';

const instance: V1ValidationPhoneGet200ResponseData = {
    number,
    valid,
    country,
    type,
    formatted,
    carrier_name,
    carrier_source,
    risk_is_voip,
    risk_is_virtual,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
