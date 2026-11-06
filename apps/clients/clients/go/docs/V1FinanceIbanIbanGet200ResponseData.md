# V1FinanceIbanIbanGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | Pointer to **string** | The normalised IBAN (spaces stripped, uppercased) | [optional] 
**Valid** | Pointer to **bool** | true if the IBAN passed length and ISO 13616 checksum validation | [optional] 
**Country** | Pointer to **string** | Full country name (empty if the country code is not in the registry) | [optional] 
**BankCode** | Pointer to **string** | Bank identifier extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] 
**Account** | Pointer to **string** | Account number extracted from the BBAN (empty if country not in registry or positions not defined) | [optional] 

## Methods

### NewV1FinanceIbanIbanGet200ResponseData

`func NewV1FinanceIbanIbanGet200ResponseData() *V1FinanceIbanIbanGet200ResponseData`

NewV1FinanceIbanIbanGet200ResponseData instantiates a new V1FinanceIbanIbanGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceIbanIbanGet200ResponseDataWithDefaults

`func NewV1FinanceIbanIbanGet200ResponseDataWithDefaults() *V1FinanceIbanIbanGet200ResponseData`

NewV1FinanceIbanIbanGet200ResponseDataWithDefaults instantiates a new V1FinanceIbanIbanGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *V1FinanceIbanIbanGet200ResponseData) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *V1FinanceIbanIbanGet200ResponseData) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *V1FinanceIbanIbanGet200ResponseData) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *V1FinanceIbanIbanGet200ResponseData) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetValid

`func (o *V1FinanceIbanIbanGet200ResponseData) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *V1FinanceIbanIbanGet200ResponseData) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *V1FinanceIbanIbanGet200ResponseData) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *V1FinanceIbanIbanGet200ResponseData) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetCountry

`func (o *V1FinanceIbanIbanGet200ResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V1FinanceIbanIbanGet200ResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V1FinanceIbanIbanGet200ResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V1FinanceIbanIbanGet200ResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetBankCode

`func (o *V1FinanceIbanIbanGet200ResponseData) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *V1FinanceIbanIbanGet200ResponseData) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *V1FinanceIbanIbanGet200ResponseData) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *V1FinanceIbanIbanGet200ResponseData) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetAccount

`func (o *V1FinanceIbanIbanGet200ResponseData) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V1FinanceIbanIbanGet200ResponseData) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V1FinanceIbanIbanGet200ResponseData) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V1FinanceIbanIbanGet200ResponseData) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


