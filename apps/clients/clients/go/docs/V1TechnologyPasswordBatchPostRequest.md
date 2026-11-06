# V1TechnologyPasswordBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **string** | List of password generation requests (1–50 items). | 
**ItemsLength** | Pointer to **int32** | Password length (8–128). Defaults to 16 when omitted. | [optional] 
**ItemsUppercase** | Pointer to **bool** | Include uppercase letters (A–Z). Defaults to false. | [optional] 
**ItemsNumbers** | Pointer to **bool** | Include numbers (0–9). Defaults to false. | [optional] 
**ItemsSymbols** | Pointer to **bool** | Include special characters. Defaults to false. | [optional] 

## Methods

### NewV1TechnologyPasswordBatchPostRequest

`func NewV1TechnologyPasswordBatchPostRequest(items string, ) *V1TechnologyPasswordBatchPostRequest`

NewV1TechnologyPasswordBatchPostRequest instantiates a new V1TechnologyPasswordBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyPasswordBatchPostRequestWithDefaults

`func NewV1TechnologyPasswordBatchPostRequestWithDefaults() *V1TechnologyPasswordBatchPostRequest`

NewV1TechnologyPasswordBatchPostRequestWithDefaults instantiates a new V1TechnologyPasswordBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1TechnologyPasswordBatchPostRequest) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1TechnologyPasswordBatchPostRequest) SetItems(v string)`

SetItems sets Items field to given value.


### GetItemsLength

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsLength() int32`

GetItemsLength returns the ItemsLength field if non-nil, zero value otherwise.

### GetItemsLengthOk

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsLengthOk() (*int32, bool)`

GetItemsLengthOk returns a tuple with the ItemsLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsLength

`func (o *V1TechnologyPasswordBatchPostRequest) SetItemsLength(v int32)`

SetItemsLength sets ItemsLength field to given value.

### HasItemsLength

`func (o *V1TechnologyPasswordBatchPostRequest) HasItemsLength() bool`

HasItemsLength returns a boolean if a field has been set.

### GetItemsUppercase

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsUppercase() bool`

GetItemsUppercase returns the ItemsUppercase field if non-nil, zero value otherwise.

### GetItemsUppercaseOk

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsUppercaseOk() (*bool, bool)`

GetItemsUppercaseOk returns a tuple with the ItemsUppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsUppercase

`func (o *V1TechnologyPasswordBatchPostRequest) SetItemsUppercase(v bool)`

SetItemsUppercase sets ItemsUppercase field to given value.

### HasItemsUppercase

`func (o *V1TechnologyPasswordBatchPostRequest) HasItemsUppercase() bool`

HasItemsUppercase returns a boolean if a field has been set.

### GetItemsNumbers

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsNumbers() bool`

GetItemsNumbers returns the ItemsNumbers field if non-nil, zero value otherwise.

### GetItemsNumbersOk

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsNumbersOk() (*bool, bool)`

GetItemsNumbersOk returns a tuple with the ItemsNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsNumbers

`func (o *V1TechnologyPasswordBatchPostRequest) SetItemsNumbers(v bool)`

SetItemsNumbers sets ItemsNumbers field to given value.

### HasItemsNumbers

`func (o *V1TechnologyPasswordBatchPostRequest) HasItemsNumbers() bool`

HasItemsNumbers returns a boolean if a field has been set.

### GetItemsSymbols

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsSymbols() bool`

GetItemsSymbols returns the ItemsSymbols field if non-nil, zero value otherwise.

### GetItemsSymbolsOk

`func (o *V1TechnologyPasswordBatchPostRequest) GetItemsSymbolsOk() (*bool, bool)`

GetItemsSymbolsOk returns a tuple with the ItemsSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsSymbols

`func (o *V1TechnologyPasswordBatchPostRequest) SetItemsSymbols(v bool)`

SetItemsSymbols sets ItemsSymbols field to given value.

### HasItemsSymbols

`func (o *V1TechnologyPasswordBatchPostRequest) HasItemsSymbols() bool`

HasItemsSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


