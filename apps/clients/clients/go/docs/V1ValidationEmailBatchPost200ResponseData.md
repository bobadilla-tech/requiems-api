# V1ValidationEmailBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | List of validation results for each email, preserving input order | [optional] 
**ResultsEmail** | Pointer to **string** | Original email input (null if invalid syntax) | [optional] 
**ResultsValid** | Pointer to **bool** | Overall validity (syntax + MX record) | [optional] 
**ResultsSyntaxValid** | Pointer to **bool** | Whether the email is syntactically valid (RFC 5322) | [optional] 
**ResultsMxValid** | Pointer to **bool** | Whether the domain has valid MX records | [optional] 
**ResultsDisposable** | Pointer to **bool** | Whether the email comes from a disposable domain | [optional] 
**ResultsNormalized** | Pointer to **string** | Canonical normalized email (lowercase, alias handling, etc.) | [optional] 
**ResultsDomain** | Pointer to **string** | Extracted domain from email address | [optional] 
**ResultsSuggestion** | Pointer to **string** | Suggested correction for common domain typos | [optional] 
**Total** | Pointer to **int32** | Number of emails processed in the batch | [optional] 

## Methods

### NewV1ValidationEmailBatchPost200ResponseData

`func NewV1ValidationEmailBatchPost200ResponseData() *V1ValidationEmailBatchPost200ResponseData`

NewV1ValidationEmailBatchPost200ResponseData instantiates a new V1ValidationEmailBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ValidationEmailBatchPost200ResponseDataWithDefaults

`func NewV1ValidationEmailBatchPost200ResponseDataWithDefaults() *V1ValidationEmailBatchPost200ResponseData`

NewV1ValidationEmailBatchPost200ResponseDataWithDefaults instantiates a new V1ValidationEmailBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsEmail

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsEmail() string`

GetResultsEmail returns the ResultsEmail field if non-nil, zero value otherwise.

### GetResultsEmailOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsEmailOk() (*string, bool)`

GetResultsEmailOk returns a tuple with the ResultsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsEmail

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsEmail(v string)`

SetResultsEmail sets ResultsEmail field to given value.

### HasResultsEmail

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsEmail() bool`

HasResultsEmail returns a boolean if a field has been set.

### GetResultsValid

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsValid() bool`

GetResultsValid returns the ResultsValid field if non-nil, zero value otherwise.

### GetResultsValidOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsValidOk() (*bool, bool)`

GetResultsValidOk returns a tuple with the ResultsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsValid

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsValid(v bool)`

SetResultsValid sets ResultsValid field to given value.

### HasResultsValid

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsValid() bool`

HasResultsValid returns a boolean if a field has been set.

### GetResultsSyntaxValid

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsSyntaxValid() bool`

GetResultsSyntaxValid returns the ResultsSyntaxValid field if non-nil, zero value otherwise.

### GetResultsSyntaxValidOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsSyntaxValidOk() (*bool, bool)`

GetResultsSyntaxValidOk returns a tuple with the ResultsSyntaxValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsSyntaxValid

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsSyntaxValid(v bool)`

SetResultsSyntaxValid sets ResultsSyntaxValid field to given value.

### HasResultsSyntaxValid

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsSyntaxValid() bool`

HasResultsSyntaxValid returns a boolean if a field has been set.

### GetResultsMxValid

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsMxValid() bool`

GetResultsMxValid returns the ResultsMxValid field if non-nil, zero value otherwise.

### GetResultsMxValidOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsMxValidOk() (*bool, bool)`

GetResultsMxValidOk returns a tuple with the ResultsMxValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsMxValid

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsMxValid(v bool)`

SetResultsMxValid sets ResultsMxValid field to given value.

### HasResultsMxValid

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsMxValid() bool`

HasResultsMxValid returns a boolean if a field has been set.

### GetResultsDisposable

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsDisposable() bool`

GetResultsDisposable returns the ResultsDisposable field if non-nil, zero value otherwise.

### GetResultsDisposableOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsDisposableOk() (*bool, bool)`

GetResultsDisposableOk returns a tuple with the ResultsDisposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDisposable

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsDisposable(v bool)`

SetResultsDisposable sets ResultsDisposable field to given value.

### HasResultsDisposable

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsDisposable() bool`

HasResultsDisposable returns a boolean if a field has been set.

### GetResultsNormalized

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsNormalized() string`

GetResultsNormalized returns the ResultsNormalized field if non-nil, zero value otherwise.

### GetResultsNormalizedOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsNormalizedOk() (*string, bool)`

GetResultsNormalizedOk returns a tuple with the ResultsNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsNormalized

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsNormalized(v string)`

SetResultsNormalized sets ResultsNormalized field to given value.

### HasResultsNormalized

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsNormalized() bool`

HasResultsNormalized returns a boolean if a field has been set.

### GetResultsDomain

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsDomain() string`

GetResultsDomain returns the ResultsDomain field if non-nil, zero value otherwise.

### GetResultsDomainOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsDomainOk() (*string, bool)`

GetResultsDomainOk returns a tuple with the ResultsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDomain

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsDomain(v string)`

SetResultsDomain sets ResultsDomain field to given value.

### HasResultsDomain

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsDomain() bool`

HasResultsDomain returns a boolean if a field has been set.

### GetResultsSuggestion

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsSuggestion() string`

GetResultsSuggestion returns the ResultsSuggestion field if non-nil, zero value otherwise.

### GetResultsSuggestionOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetResultsSuggestionOk() (*string, bool)`

GetResultsSuggestionOk returns a tuple with the ResultsSuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsSuggestion

`func (o *V1ValidationEmailBatchPost200ResponseData) SetResultsSuggestion(v string)`

SetResultsSuggestion sets ResultsSuggestion field to given value.

### HasResultsSuggestion

`func (o *V1ValidationEmailBatchPost200ResponseData) HasResultsSuggestion() bool`

HasResultsSuggestion returns a boolean if a field has been set.

### GetTotal

`func (o *V1ValidationEmailBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1ValidationEmailBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1ValidationEmailBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1ValidationEmailBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


