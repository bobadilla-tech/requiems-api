# V1ValidationEmailPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address exactly as supplied in the request body; null when syntax is invalid | [optional] 
**Valid** | Pointer to **bool** | Overall validity. True only when the address passes syntax validation and the domain has at least one MX record | [optional] 
**SyntaxValid** | Pointer to **bool** | Whether the address is syntactically valid according to RFC 5322 | [optional] 
**MxValid** | Pointer to **bool** | Whether the domain has at least one MX record, meaning it can receive email | [optional] 
**Disposable** | Pointer to **bool** | Whether the address uses a known disposable or temporary email domain | [optional] 
**Normalized** | Pointer to **string** | The canonical form of the address after normalization (lowercase, plus-tag removal, alias-domain resolution). Null when syntax is invalid | [optional] 
**Domain** | Pointer to **string** | The domain part of the address (after @). Null when syntax is invalid | [optional] 
**Suggestion** | Pointer to **string** | A corrected domain name when the supplied domain looks like a typo of a well-known provider (e.g. gmial.com → gmail.com). Null when no close match is found or the domain is already correct | [optional] 

## Methods

### NewV1ValidationEmailPost200ResponseData

`func NewV1ValidationEmailPost200ResponseData() *V1ValidationEmailPost200ResponseData`

NewV1ValidationEmailPost200ResponseData instantiates a new V1ValidationEmailPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ValidationEmailPost200ResponseDataWithDefaults

`func NewV1ValidationEmailPost200ResponseDataWithDefaults() *V1ValidationEmailPost200ResponseData`

NewV1ValidationEmailPost200ResponseDataWithDefaults instantiates a new V1ValidationEmailPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *V1ValidationEmailPost200ResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1ValidationEmailPost200ResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1ValidationEmailPost200ResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *V1ValidationEmailPost200ResponseData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetValid

`func (o *V1ValidationEmailPost200ResponseData) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *V1ValidationEmailPost200ResponseData) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *V1ValidationEmailPost200ResponseData) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *V1ValidationEmailPost200ResponseData) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetSyntaxValid

`func (o *V1ValidationEmailPost200ResponseData) GetSyntaxValid() bool`

GetSyntaxValid returns the SyntaxValid field if non-nil, zero value otherwise.

### GetSyntaxValidOk

`func (o *V1ValidationEmailPost200ResponseData) GetSyntaxValidOk() (*bool, bool)`

GetSyntaxValidOk returns a tuple with the SyntaxValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxValid

`func (o *V1ValidationEmailPost200ResponseData) SetSyntaxValid(v bool)`

SetSyntaxValid sets SyntaxValid field to given value.

### HasSyntaxValid

`func (o *V1ValidationEmailPost200ResponseData) HasSyntaxValid() bool`

HasSyntaxValid returns a boolean if a field has been set.

### GetMxValid

`func (o *V1ValidationEmailPost200ResponseData) GetMxValid() bool`

GetMxValid returns the MxValid field if non-nil, zero value otherwise.

### GetMxValidOk

`func (o *V1ValidationEmailPost200ResponseData) GetMxValidOk() (*bool, bool)`

GetMxValidOk returns a tuple with the MxValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxValid

`func (o *V1ValidationEmailPost200ResponseData) SetMxValid(v bool)`

SetMxValid sets MxValid field to given value.

### HasMxValid

`func (o *V1ValidationEmailPost200ResponseData) HasMxValid() bool`

HasMxValid returns a boolean if a field has been set.

### GetDisposable

`func (o *V1ValidationEmailPost200ResponseData) GetDisposable() bool`

GetDisposable returns the Disposable field if non-nil, zero value otherwise.

### GetDisposableOk

`func (o *V1ValidationEmailPost200ResponseData) GetDisposableOk() (*bool, bool)`

GetDisposableOk returns a tuple with the Disposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposable

`func (o *V1ValidationEmailPost200ResponseData) SetDisposable(v bool)`

SetDisposable sets Disposable field to given value.

### HasDisposable

`func (o *V1ValidationEmailPost200ResponseData) HasDisposable() bool`

HasDisposable returns a boolean if a field has been set.

### GetNormalized

`func (o *V1ValidationEmailPost200ResponseData) GetNormalized() string`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *V1ValidationEmailPost200ResponseData) GetNormalizedOk() (*string, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *V1ValidationEmailPost200ResponseData) SetNormalized(v string)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *V1ValidationEmailPost200ResponseData) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### GetDomain

`func (o *V1ValidationEmailPost200ResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V1ValidationEmailPost200ResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V1ValidationEmailPost200ResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V1ValidationEmailPost200ResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSuggestion

`func (o *V1ValidationEmailPost200ResponseData) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *V1ValidationEmailPost200ResponseData) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *V1ValidationEmailPost200ResponseData) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.

### HasSuggestion

`func (o *V1ValidationEmailPost200ResponseData) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


