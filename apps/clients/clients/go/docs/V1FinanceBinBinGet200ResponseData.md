# V1FinanceBinBinGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | The normalised BIN prefix used for the lookup | [optional] 
**Scheme** | Pointer to **string** | Card network: visa, mastercard, amex, discover, jcb, diners, unionpay, maestro, mir, rupay, private_label | [optional] 
**CardType** | Pointer to **string** | credit, debit, prepaid, or charge | [optional] 
**CardLevel** | Pointer to **string** | classic, gold, platinum, infinite, business, signature, or standard | [optional] 
**IssuerName** | Pointer to **string** | Name of the card-issuing bank | [optional] 
**IssuerUrl** | Pointer to **string** | Bank website URL | [optional] 
**IssuerPhone** | Pointer to **string** | Bank customer service phone number | [optional] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code of the issuing bank (e.g. US, GB, DE) | [optional] 
**CountryName** | Pointer to **string** | Full country name of the issuing bank | [optional] 
**Prepaid** | Pointer to **bool** | Whether the card is a prepaid card | [optional] 
**Luhn** | Pointer to **bool** | Whether the BIN prefix passes the Luhn algorithm check | [optional] 
**Confidence** | Pointer to **float32** | Data quality score (0.00–1.00). Multi-source confirmed records score higher. | [optional] 

## Methods

### NewV1FinanceBinBinGet200ResponseData

`func NewV1FinanceBinBinGet200ResponseData() *V1FinanceBinBinGet200ResponseData`

NewV1FinanceBinBinGet200ResponseData instantiates a new V1FinanceBinBinGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceBinBinGet200ResponseDataWithDefaults

`func NewV1FinanceBinBinGet200ResponseDataWithDefaults() *V1FinanceBinBinGet200ResponseData`

NewV1FinanceBinBinGet200ResponseDataWithDefaults instantiates a new V1FinanceBinBinGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *V1FinanceBinBinGet200ResponseData) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *V1FinanceBinBinGet200ResponseData) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *V1FinanceBinBinGet200ResponseData) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *V1FinanceBinBinGet200ResponseData) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetScheme

`func (o *V1FinanceBinBinGet200ResponseData) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *V1FinanceBinBinGet200ResponseData) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *V1FinanceBinBinGet200ResponseData) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *V1FinanceBinBinGet200ResponseData) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetCardType

`func (o *V1FinanceBinBinGet200ResponseData) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *V1FinanceBinBinGet200ResponseData) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *V1FinanceBinBinGet200ResponseData) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *V1FinanceBinBinGet200ResponseData) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardLevel

`func (o *V1FinanceBinBinGet200ResponseData) GetCardLevel() string`

GetCardLevel returns the CardLevel field if non-nil, zero value otherwise.

### GetCardLevelOk

`func (o *V1FinanceBinBinGet200ResponseData) GetCardLevelOk() (*string, bool)`

GetCardLevelOk returns a tuple with the CardLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLevel

`func (o *V1FinanceBinBinGet200ResponseData) SetCardLevel(v string)`

SetCardLevel sets CardLevel field to given value.

### HasCardLevel

`func (o *V1FinanceBinBinGet200ResponseData) HasCardLevel() bool`

HasCardLevel returns a boolean if a field has been set.

### GetIssuerName

`func (o *V1FinanceBinBinGet200ResponseData) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *V1FinanceBinBinGet200ResponseData) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *V1FinanceBinBinGet200ResponseData) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *V1FinanceBinBinGet200ResponseData) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *V1FinanceBinBinGet200ResponseData) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *V1FinanceBinBinGet200ResponseData) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *V1FinanceBinBinGet200ResponseData) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *V1FinanceBinBinGet200ResponseData) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetIssuerPhone

`func (o *V1FinanceBinBinGet200ResponseData) GetIssuerPhone() string`

GetIssuerPhone returns the IssuerPhone field if non-nil, zero value otherwise.

### GetIssuerPhoneOk

`func (o *V1FinanceBinBinGet200ResponseData) GetIssuerPhoneOk() (*string, bool)`

GetIssuerPhoneOk returns a tuple with the IssuerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerPhone

`func (o *V1FinanceBinBinGet200ResponseData) SetIssuerPhone(v string)`

SetIssuerPhone sets IssuerPhone field to given value.

### HasIssuerPhone

`func (o *V1FinanceBinBinGet200ResponseData) HasIssuerPhone() bool`

HasIssuerPhone returns a boolean if a field has been set.

### GetCountryCode

`func (o *V1FinanceBinBinGet200ResponseData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V1FinanceBinBinGet200ResponseData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V1FinanceBinBinGet200ResponseData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *V1FinanceBinBinGet200ResponseData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *V1FinanceBinBinGet200ResponseData) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *V1FinanceBinBinGet200ResponseData) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *V1FinanceBinBinGet200ResponseData) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *V1FinanceBinBinGet200ResponseData) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetPrepaid

`func (o *V1FinanceBinBinGet200ResponseData) GetPrepaid() bool`

GetPrepaid returns the Prepaid field if non-nil, zero value otherwise.

### GetPrepaidOk

`func (o *V1FinanceBinBinGet200ResponseData) GetPrepaidOk() (*bool, bool)`

GetPrepaidOk returns a tuple with the Prepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaid

`func (o *V1FinanceBinBinGet200ResponseData) SetPrepaid(v bool)`

SetPrepaid sets Prepaid field to given value.

### HasPrepaid

`func (o *V1FinanceBinBinGet200ResponseData) HasPrepaid() bool`

HasPrepaid returns a boolean if a field has been set.

### GetLuhn

`func (o *V1FinanceBinBinGet200ResponseData) GetLuhn() bool`

GetLuhn returns the Luhn field if non-nil, zero value otherwise.

### GetLuhnOk

`func (o *V1FinanceBinBinGet200ResponseData) GetLuhnOk() (*bool, bool)`

GetLuhnOk returns a tuple with the Luhn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLuhn

`func (o *V1FinanceBinBinGet200ResponseData) SetLuhn(v bool)`

SetLuhn sets Luhn field to given value.

### HasLuhn

`func (o *V1FinanceBinBinGet200ResponseData) HasLuhn() bool`

HasLuhn returns a boolean if a field has been set.

### GetConfidence

`func (o *V1FinanceBinBinGet200ResponseData) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *V1FinanceBinBinGet200ResponseData) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *V1FinanceBinBinGet200ResponseData) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *V1FinanceBinBinGet200ResponseData) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


