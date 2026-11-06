# V1TechnologyPasswordGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The generated password | [optional] 
**Length** | Pointer to **int32** | Length of the generated password | [optional] 
**Strength** | Pointer to **string** | Password strength assessment (weak, medium, or strong) | [optional] 

## Methods

### NewV1TechnologyPasswordGet200ResponseData

`func NewV1TechnologyPasswordGet200ResponseData() *V1TechnologyPasswordGet200ResponseData`

NewV1TechnologyPasswordGet200ResponseData instantiates a new V1TechnologyPasswordGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyPasswordGet200ResponseDataWithDefaults

`func NewV1TechnologyPasswordGet200ResponseDataWithDefaults() *V1TechnologyPasswordGet200ResponseData`

NewV1TechnologyPasswordGet200ResponseDataWithDefaults instantiates a new V1TechnologyPasswordGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *V1TechnologyPasswordGet200ResponseData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1TechnologyPasswordGet200ResponseData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1TechnologyPasswordGet200ResponseData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1TechnologyPasswordGet200ResponseData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetLength

`func (o *V1TechnologyPasswordGet200ResponseData) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V1TechnologyPasswordGet200ResponseData) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V1TechnologyPasswordGet200ResponseData) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V1TechnologyPasswordGet200ResponseData) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetStrength

`func (o *V1TechnologyPasswordGet200ResponseData) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *V1TechnologyPasswordGet200ResponseData) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *V1TechnologyPasswordGet200ResponseData) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *V1TechnologyPasswordGet200ResponseData) HasStrength() bool`

HasStrength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


