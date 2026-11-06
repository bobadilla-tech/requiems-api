# V1ValidationPhoneGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | The original number as supplied in the request | [optional] 
**Valid** | Pointer to **bool** | Whether the number is a valid, dialable phone number | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country code (omitted when valid is false) | [optional] 
**Type** | Pointer to **string** | Number type: mobile, landline, landline_or_mobile, toll_free, voip, premium_rate, shared_cost, personal_number, pager, uan, voicemail, or unknown (omitted when valid is false) | [optional] 
**Formatted** | Pointer to **string** | International format of the number, e.g. +44 7400 123456 (omitted when valid is false) | [optional] 
**CarrierName** | Pointer to **string** | Carrier name from phone prefix metadata (omitted when carrier cannot be determined) | [optional] 
**CarrierSource** | Pointer to **string** | How the carrier was determined. Always \&quot;metadata\&quot; when present | [optional] 
**RiskIsVoip** | Pointer to **bool** | true when the number type is voip | [optional] 
**RiskIsVirtual** | Pointer to **bool** | true when the number is not tied to a physical SIM or fixed line: voip, personal_number, uan, pager, or voicemail | [optional] 

## Methods

### NewV1ValidationPhoneGet200ResponseData

`func NewV1ValidationPhoneGet200ResponseData() *V1ValidationPhoneGet200ResponseData`

NewV1ValidationPhoneGet200ResponseData instantiates a new V1ValidationPhoneGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ValidationPhoneGet200ResponseDataWithDefaults

`func NewV1ValidationPhoneGet200ResponseDataWithDefaults() *V1ValidationPhoneGet200ResponseData`

NewV1ValidationPhoneGet200ResponseDataWithDefaults instantiates a new V1ValidationPhoneGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *V1ValidationPhoneGet200ResponseData) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *V1ValidationPhoneGet200ResponseData) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *V1ValidationPhoneGet200ResponseData) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *V1ValidationPhoneGet200ResponseData) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetValid

`func (o *V1ValidationPhoneGet200ResponseData) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *V1ValidationPhoneGet200ResponseData) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *V1ValidationPhoneGet200ResponseData) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *V1ValidationPhoneGet200ResponseData) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetCountry

`func (o *V1ValidationPhoneGet200ResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V1ValidationPhoneGet200ResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V1ValidationPhoneGet200ResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V1ValidationPhoneGet200ResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetType

`func (o *V1ValidationPhoneGet200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ValidationPhoneGet200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ValidationPhoneGet200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ValidationPhoneGet200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFormatted

`func (o *V1ValidationPhoneGet200ResponseData) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *V1ValidationPhoneGet200ResponseData) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *V1ValidationPhoneGet200ResponseData) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *V1ValidationPhoneGet200ResponseData) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### GetCarrierName

`func (o *V1ValidationPhoneGet200ResponseData) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *V1ValidationPhoneGet200ResponseData) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *V1ValidationPhoneGet200ResponseData) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *V1ValidationPhoneGet200ResponseData) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetCarrierSource

`func (o *V1ValidationPhoneGet200ResponseData) GetCarrierSource() string`

GetCarrierSource returns the CarrierSource field if non-nil, zero value otherwise.

### GetCarrierSourceOk

`func (o *V1ValidationPhoneGet200ResponseData) GetCarrierSourceOk() (*string, bool)`

GetCarrierSourceOk returns a tuple with the CarrierSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSource

`func (o *V1ValidationPhoneGet200ResponseData) SetCarrierSource(v string)`

SetCarrierSource sets CarrierSource field to given value.

### HasCarrierSource

`func (o *V1ValidationPhoneGet200ResponseData) HasCarrierSource() bool`

HasCarrierSource returns a boolean if a field has been set.

### GetRiskIsVoip

`func (o *V1ValidationPhoneGet200ResponseData) GetRiskIsVoip() bool`

GetRiskIsVoip returns the RiskIsVoip field if non-nil, zero value otherwise.

### GetRiskIsVoipOk

`func (o *V1ValidationPhoneGet200ResponseData) GetRiskIsVoipOk() (*bool, bool)`

GetRiskIsVoipOk returns a tuple with the RiskIsVoip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskIsVoip

`func (o *V1ValidationPhoneGet200ResponseData) SetRiskIsVoip(v bool)`

SetRiskIsVoip sets RiskIsVoip field to given value.

### HasRiskIsVoip

`func (o *V1ValidationPhoneGet200ResponseData) HasRiskIsVoip() bool`

HasRiskIsVoip returns a boolean if a field has been set.

### GetRiskIsVirtual

`func (o *V1ValidationPhoneGet200ResponseData) GetRiskIsVirtual() bool`

GetRiskIsVirtual returns the RiskIsVirtual field if non-nil, zero value otherwise.

### GetRiskIsVirtualOk

`func (o *V1ValidationPhoneGet200ResponseData) GetRiskIsVirtualOk() (*bool, bool)`

GetRiskIsVirtualOk returns a tuple with the RiskIsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskIsVirtual

`func (o *V1ValidationPhoneGet200ResponseData) SetRiskIsVirtual(v bool)`

SetRiskIsVirtual sets RiskIsVirtual field to given value.

### HasRiskIsVirtual

`func (o *V1ValidationPhoneGet200ResponseData) HasRiskIsVirtual() bool`

HasRiskIsVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


