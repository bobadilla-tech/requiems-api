# V1TextNormalizePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Original** | Pointer to **string** | The email address exactly as supplied in the request body | [optional] 
**Normalized** | Pointer to **string** | The canonical form of the address after all transformations | [optional] 
**Local** | Pointer to **string** | The local part (before @) of the normalized address | [optional] 
**Domain** | Pointer to **string** | The domain part (after @) of the normalized address | [optional] 
**Changes** | Pointer to **[]interface{}** | Ordered list of transformations applied. Possible values: lowercased, trimmed_whitespace, removed_dots, removed_plus_tag, canonicalised_domain. Empty array when no changes were needed. | [optional] 

## Methods

### NewV1TextNormalizePost200ResponseData

`func NewV1TextNormalizePost200ResponseData() *V1TextNormalizePost200ResponseData`

NewV1TextNormalizePost200ResponseData instantiates a new V1TextNormalizePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextNormalizePost200ResponseDataWithDefaults

`func NewV1TextNormalizePost200ResponseDataWithDefaults() *V1TextNormalizePost200ResponseData`

NewV1TextNormalizePost200ResponseDataWithDefaults instantiates a new V1TextNormalizePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginal

`func (o *V1TextNormalizePost200ResponseData) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *V1TextNormalizePost200ResponseData) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *V1TextNormalizePost200ResponseData) SetOriginal(v string)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *V1TextNormalizePost200ResponseData) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetNormalized

`func (o *V1TextNormalizePost200ResponseData) GetNormalized() string`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *V1TextNormalizePost200ResponseData) GetNormalizedOk() (*string, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *V1TextNormalizePost200ResponseData) SetNormalized(v string)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *V1TextNormalizePost200ResponseData) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### GetLocal

`func (o *V1TextNormalizePost200ResponseData) GetLocal() string`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *V1TextNormalizePost200ResponseData) GetLocalOk() (*string, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *V1TextNormalizePost200ResponseData) SetLocal(v string)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *V1TextNormalizePost200ResponseData) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetDomain

`func (o *V1TextNormalizePost200ResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V1TextNormalizePost200ResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V1TextNormalizePost200ResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V1TextNormalizePost200ResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetChanges

`func (o *V1TextNormalizePost200ResponseData) GetChanges() []interface{}`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *V1TextNormalizePost200ResponseData) GetChangesOk() (*[]interface{}, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *V1TextNormalizePost200ResponseData) SetChanges(v []interface{})`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *V1TextNormalizePost200ResponseData) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


