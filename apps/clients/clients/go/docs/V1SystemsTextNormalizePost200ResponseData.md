# V1SystemsTextNormalizePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Original** | Pointer to **string** | The input text as provided | [optional] 
**Normalized** | Pointer to **string** | The text after all operations have been applied in order | [optional] 
**Operations** | Pointer to **[]interface{}** | The operations list echoed from the request | [optional] 

## Methods

### NewV1SystemsTextNormalizePost200ResponseData

`func NewV1SystemsTextNormalizePost200ResponseData() *V1SystemsTextNormalizePost200ResponseData`

NewV1SystemsTextNormalizePost200ResponseData instantiates a new V1SystemsTextNormalizePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsTextNormalizePost200ResponseDataWithDefaults

`func NewV1SystemsTextNormalizePost200ResponseDataWithDefaults() *V1SystemsTextNormalizePost200ResponseData`

NewV1SystemsTextNormalizePost200ResponseDataWithDefaults instantiates a new V1SystemsTextNormalizePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginal

`func (o *V1SystemsTextNormalizePost200ResponseData) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *V1SystemsTextNormalizePost200ResponseData) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *V1SystemsTextNormalizePost200ResponseData) SetOriginal(v string)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *V1SystemsTextNormalizePost200ResponseData) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetNormalized

`func (o *V1SystemsTextNormalizePost200ResponseData) GetNormalized() string`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *V1SystemsTextNormalizePost200ResponseData) GetNormalizedOk() (*string, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *V1SystemsTextNormalizePost200ResponseData) SetNormalized(v string)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *V1SystemsTextNormalizePost200ResponseData) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### GetOperations

`func (o *V1SystemsTextNormalizePost200ResponseData) GetOperations() []interface{}`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *V1SystemsTextNormalizePost200ResponseData) GetOperationsOk() (*[]interface{}, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *V1SystemsTextNormalizePost200ResponseData) SetOperations(v []interface{})`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *V1SystemsTextNormalizePost200ResponseData) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


