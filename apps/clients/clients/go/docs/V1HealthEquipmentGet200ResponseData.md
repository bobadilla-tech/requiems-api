# V1HealthEquipmentGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]interface{}** | Sorted list of all distinct equipment names | [optional] 
**Total** | Pointer to **int32** | Total number of distinct equipment types | [optional] 

## Methods

### NewV1HealthEquipmentGet200ResponseData

`func NewV1HealthEquipmentGet200ResponseData() *V1HealthEquipmentGet200ResponseData`

NewV1HealthEquipmentGet200ResponseData instantiates a new V1HealthEquipmentGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HealthEquipmentGet200ResponseDataWithDefaults

`func NewV1HealthEquipmentGet200ResponseDataWithDefaults() *V1HealthEquipmentGet200ResponseData`

NewV1HealthEquipmentGet200ResponseDataWithDefaults instantiates a new V1HealthEquipmentGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1HealthEquipmentGet200ResponseData) GetItems() []interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1HealthEquipmentGet200ResponseData) GetItemsOk() (*[]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1HealthEquipmentGet200ResponseData) SetItems(v []interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *V1HealthEquipmentGet200ResponseData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *V1HealthEquipmentGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1HealthEquipmentGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1HealthEquipmentGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1HealthEquipmentGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


