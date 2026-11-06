# Org.OpenAPITools.Model.V1ValidationPhoneGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The original number as supplied in the request | [optional] 
**Valid** | **bool** | Whether the number is a valid, dialable phone number | [optional] 
**Country** | **string** | ISO 3166-1 alpha-2 country code (omitted when valid is false) | [optional] 
**Type** | **string** | Number type: mobile, landline, landline_or_mobile, toll_free, voip, premium_rate, shared_cost, personal_number, pager, uan, voicemail, or unknown (omitted when valid is false) | [optional] 
**Formatted** | **string** | International format of the number, e.g. +44 7400 123456 (omitted when valid is false) | [optional] 
**CarrierName** | **string** | Carrier name from phone prefix metadata (omitted when carrier cannot be determined) | [optional] 
**CarrierSource** | **string** | How the carrier was determined. Always \&quot;metadata\&quot; when present | [optional] 
**RiskIsVoip** | **bool** | true when the number type is voip | [optional] 
**RiskIsVirtual** | **bool** | true when the number is not tied to a physical SIM or fixed line: voip, personal_number, uan, pager, or voicemail | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

