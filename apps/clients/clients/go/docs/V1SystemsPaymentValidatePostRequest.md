# V1SystemsPaymentValidatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | Card BIN (first 6–8 digits). At least one of bin, iban, or swift is required. | [optional] 
**Iban** | Pointer to **string** | IBAN string to validate. Checksum and format are verified. | [optional] 
**Swift** | Pointer to **string** | SWIFT/BIC code to look up institution and country. | [optional] 

## Methods

### NewV1SystemsPaymentValidatePostRequest

`func NewV1SystemsPaymentValidatePostRequest() *V1SystemsPaymentValidatePostRequest`

NewV1SystemsPaymentValidatePostRequest instantiates a new V1SystemsPaymentValidatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsPaymentValidatePostRequestWithDefaults

`func NewV1SystemsPaymentValidatePostRequestWithDefaults() *V1SystemsPaymentValidatePostRequest`

NewV1SystemsPaymentValidatePostRequestWithDefaults instantiates a new V1SystemsPaymentValidatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *V1SystemsPaymentValidatePostRequest) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *V1SystemsPaymentValidatePostRequest) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *V1SystemsPaymentValidatePostRequest) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *V1SystemsPaymentValidatePostRequest) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetIban

`func (o *V1SystemsPaymentValidatePostRequest) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *V1SystemsPaymentValidatePostRequest) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *V1SystemsPaymentValidatePostRequest) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *V1SystemsPaymentValidatePostRequest) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSwift

`func (o *V1SystemsPaymentValidatePostRequest) GetSwift() string`

GetSwift returns the Swift field if non-nil, zero value otherwise.

### GetSwiftOk

`func (o *V1SystemsPaymentValidatePostRequest) GetSwiftOk() (*string, bool)`

GetSwiftOk returns a tuple with the Swift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwift

`func (o *V1SystemsPaymentValidatePostRequest) SetSwift(v string)`

SetSwift sets Swift field to given value.

### HasSwift

`func (o *V1SystemsPaymentValidatePostRequest) HasSwift() bool`

HasSwift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


