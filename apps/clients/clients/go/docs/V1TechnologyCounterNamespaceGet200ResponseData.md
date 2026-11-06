# V1TechnologyCounterNamespaceGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | The counter namespace | [optional] 
**Value** | Pointer to **int32** | The current counter value (returns 0 if counter doesn&#39;t exist) | [optional] 

## Methods

### NewV1TechnologyCounterNamespaceGet200ResponseData

`func NewV1TechnologyCounterNamespaceGet200ResponseData() *V1TechnologyCounterNamespaceGet200ResponseData`

NewV1TechnologyCounterNamespaceGet200ResponseData instantiates a new V1TechnologyCounterNamespaceGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyCounterNamespaceGet200ResponseDataWithDefaults

`func NewV1TechnologyCounterNamespaceGet200ResponseDataWithDefaults() *V1TechnologyCounterNamespaceGet200ResponseData`

NewV1TechnologyCounterNamespaceGet200ResponseDataWithDefaults instantiates a new V1TechnologyCounterNamespaceGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetValue

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *V1TechnologyCounterNamespaceGet200ResponseData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


