# V1SystemsPaymentValidatePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | BIN result. Null when bin was not provided. Contains valid, scheme, card_type, card_level, country_code, issuer, prepaid, luhn. | [optional] 
**Iban** | Pointer to **string** | IBAN result. Null when iban was not provided. Contains valid, country_code, bank_code, account_number. | [optional] 
**Swift** | Pointer to **string** | SWIFT result. Null when swift was not provided. Contains valid, institution, country, branch. | [optional] 
**ConsistencyOk** | Pointer to **bool** | True when no cross-field mismatches were detected | [optional] 
**ConsistencyFlags** | Pointer to **[]interface{}** | Mismatch flags. Possible values: country_mismatch_bin_iban, country_mismatch_bin_swift, country_mismatch_iban_swift, bank_mismatch_iban_swift | [optional] 

## Methods

### NewV1SystemsPaymentValidatePost200ResponseData

`func NewV1SystemsPaymentValidatePost200ResponseData() *V1SystemsPaymentValidatePost200ResponseData`

NewV1SystemsPaymentValidatePost200ResponseData instantiates a new V1SystemsPaymentValidatePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsPaymentValidatePost200ResponseDataWithDefaults

`func NewV1SystemsPaymentValidatePost200ResponseDataWithDefaults() *V1SystemsPaymentValidatePost200ResponseData`

NewV1SystemsPaymentValidatePost200ResponseDataWithDefaults instantiates a new V1SystemsPaymentValidatePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *V1SystemsPaymentValidatePost200ResponseData) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *V1SystemsPaymentValidatePost200ResponseData) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetIban

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *V1SystemsPaymentValidatePost200ResponseData) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *V1SystemsPaymentValidatePost200ResponseData) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSwift

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetSwift() string`

GetSwift returns the Swift field if non-nil, zero value otherwise.

### GetSwiftOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetSwiftOk() (*string, bool)`

GetSwiftOk returns a tuple with the Swift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwift

`func (o *V1SystemsPaymentValidatePost200ResponseData) SetSwift(v string)`

SetSwift sets Swift field to given value.

### HasSwift

`func (o *V1SystemsPaymentValidatePost200ResponseData) HasSwift() bool`

HasSwift returns a boolean if a field has been set.

### GetConsistencyOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetConsistencyOk() bool`

GetConsistencyOk returns the ConsistencyOk field if non-nil, zero value otherwise.

### GetConsistencyOkOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetConsistencyOkOk() (*bool, bool)`

GetConsistencyOkOk returns a tuple with the ConsistencyOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) SetConsistencyOk(v bool)`

SetConsistencyOk sets ConsistencyOk field to given value.

### HasConsistencyOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) HasConsistencyOk() bool`

HasConsistencyOk returns a boolean if a field has been set.

### GetConsistencyFlags

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetConsistencyFlags() []interface{}`

GetConsistencyFlags returns the ConsistencyFlags field if non-nil, zero value otherwise.

### GetConsistencyFlagsOk

`func (o *V1SystemsPaymentValidatePost200ResponseData) GetConsistencyFlagsOk() (*[]interface{}, bool)`

GetConsistencyFlagsOk returns a tuple with the ConsistencyFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyFlags

`func (o *V1SystemsPaymentValidatePost200ResponseData) SetConsistencyFlags(v []interface{})`

SetConsistencyFlags sets ConsistencyFlags field to given value.

### HasConsistencyFlags

`func (o *V1SystemsPaymentValidatePost200ResponseData) HasConsistencyFlags() bool`

HasConsistencyFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


