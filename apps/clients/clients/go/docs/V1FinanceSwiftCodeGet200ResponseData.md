# V1FinanceSwiftCodeGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwiftCode** | Pointer to **string** | Full 11-character SWIFT/BIC code | [optional] 
**BankCode** | Pointer to **string** | Institution code (characters 1-4) | [optional] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code (characters 5-6) | [optional] 
**LocationCode** | Pointer to **string** | Location code (characters 7-8) | [optional] 
**BranchCode** | Pointer to **string** | Branch code (characters 9-11), XXX for primary office | [optional] 
**BankName** | Pointer to **string** | Bank or institution name | [optional] 
**City** | Pointer to **string** | City of the branch or primary office | [optional] 
**CountryName** | Pointer to **string** | Full country name | [optional] 
**IsPrimary** | Pointer to **bool** | true when branch_code is XXX | [optional] 

## Methods

### NewV1FinanceSwiftCodeGet200ResponseData

`func NewV1FinanceSwiftCodeGet200ResponseData() *V1FinanceSwiftCodeGet200ResponseData`

NewV1FinanceSwiftCodeGet200ResponseData instantiates a new V1FinanceSwiftCodeGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceSwiftCodeGet200ResponseDataWithDefaults

`func NewV1FinanceSwiftCodeGet200ResponseDataWithDefaults() *V1FinanceSwiftCodeGet200ResponseData`

NewV1FinanceSwiftCodeGet200ResponseDataWithDefaults instantiates a new V1FinanceSwiftCodeGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwiftCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetBankCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLocationCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetLocationCode() string`

GetLocationCode returns the LocationCode field if non-nil, zero value otherwise.

### GetLocationCodeOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetLocationCodeOk() (*string, bool)`

GetLocationCodeOk returns a tuple with the LocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetLocationCode(v string)`

SetLocationCode sets LocationCode field to given value.

### HasLocationCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasLocationCode() bool`

HasLocationCode returns a boolean if a field has been set.

### GetBranchCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### GetBankName

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetCity

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryName

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetIsPrimary

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *V1FinanceSwiftCodeGet200ResponseData) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *V1FinanceSwiftCodeGet200ResponseData) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *V1FinanceSwiftCodeGet200ResponseData) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


