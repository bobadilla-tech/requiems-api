# V1HealthExercisesGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]interface{}** | Array of exercise objects for the current page | [optional] 
**ItemsId** | Pointer to **int32** | Unique exercise identifier | [optional] 
**ItemsName** | Pointer to **string** | Exercise name | [optional] 
**ItemsBodyParts** | Pointer to **[]interface{}** | Body part categories involved | [optional] 
**ItemsEquipment** | Pointer to **[]interface{}** | Equipment required | [optional] 
**ItemsTargetMuscles** | Pointer to **[]interface{}** | Primary muscles targeted | [optional] 
**ItemsSecondaryMuscles** | Pointer to **[]interface{}** | Secondary muscles engaged | [optional] 
**ItemsInstructions** | Pointer to **[]interface{}** | Ordered step-by-step instructions | [optional] 
**Total** | Pointer to **int32** | Total number of exercises matching the filters | [optional] 
**Page** | Pointer to **int32** | Current page number | [optional] 
**PerPage** | Pointer to **int32** | Number of results per page | [optional] 

## Methods

### NewV1HealthExercisesGet200ResponseData

`func NewV1HealthExercisesGet200ResponseData() *V1HealthExercisesGet200ResponseData`

NewV1HealthExercisesGet200ResponseData instantiates a new V1HealthExercisesGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HealthExercisesGet200ResponseDataWithDefaults

`func NewV1HealthExercisesGet200ResponseDataWithDefaults() *V1HealthExercisesGet200ResponseData`

NewV1HealthExercisesGet200ResponseDataWithDefaults instantiates a new V1HealthExercisesGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1HealthExercisesGet200ResponseData) GetItems() []interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsOk() (*[]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1HealthExercisesGet200ResponseData) SetItems(v []interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *V1HealthExercisesGet200ResponseData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetItemsId

`func (o *V1HealthExercisesGet200ResponseData) GetItemsId() int32`

GetItemsId returns the ItemsId field if non-nil, zero value otherwise.

### GetItemsIdOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsIdOk() (*int32, bool)`

GetItemsIdOk returns a tuple with the ItemsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsId

`func (o *V1HealthExercisesGet200ResponseData) SetItemsId(v int32)`

SetItemsId sets ItemsId field to given value.

### HasItemsId

`func (o *V1HealthExercisesGet200ResponseData) HasItemsId() bool`

HasItemsId returns a boolean if a field has been set.

### GetItemsName

`func (o *V1HealthExercisesGet200ResponseData) GetItemsName() string`

GetItemsName returns the ItemsName field if non-nil, zero value otherwise.

### GetItemsNameOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsNameOk() (*string, bool)`

GetItemsNameOk returns a tuple with the ItemsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsName

`func (o *V1HealthExercisesGet200ResponseData) SetItemsName(v string)`

SetItemsName sets ItemsName field to given value.

### HasItemsName

`func (o *V1HealthExercisesGet200ResponseData) HasItemsName() bool`

HasItemsName returns a boolean if a field has been set.

### GetItemsBodyParts

`func (o *V1HealthExercisesGet200ResponseData) GetItemsBodyParts() []interface{}`

GetItemsBodyParts returns the ItemsBodyParts field if non-nil, zero value otherwise.

### GetItemsBodyPartsOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsBodyPartsOk() (*[]interface{}, bool)`

GetItemsBodyPartsOk returns a tuple with the ItemsBodyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsBodyParts

`func (o *V1HealthExercisesGet200ResponseData) SetItemsBodyParts(v []interface{})`

SetItemsBodyParts sets ItemsBodyParts field to given value.

### HasItemsBodyParts

`func (o *V1HealthExercisesGet200ResponseData) HasItemsBodyParts() bool`

HasItemsBodyParts returns a boolean if a field has been set.

### GetItemsEquipment

`func (o *V1HealthExercisesGet200ResponseData) GetItemsEquipment() []interface{}`

GetItemsEquipment returns the ItemsEquipment field if non-nil, zero value otherwise.

### GetItemsEquipmentOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsEquipmentOk() (*[]interface{}, bool)`

GetItemsEquipmentOk returns a tuple with the ItemsEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsEquipment

`func (o *V1HealthExercisesGet200ResponseData) SetItemsEquipment(v []interface{})`

SetItemsEquipment sets ItemsEquipment field to given value.

### HasItemsEquipment

`func (o *V1HealthExercisesGet200ResponseData) HasItemsEquipment() bool`

HasItemsEquipment returns a boolean if a field has been set.

### GetItemsTargetMuscles

`func (o *V1HealthExercisesGet200ResponseData) GetItemsTargetMuscles() []interface{}`

GetItemsTargetMuscles returns the ItemsTargetMuscles field if non-nil, zero value otherwise.

### GetItemsTargetMusclesOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsTargetMusclesOk() (*[]interface{}, bool)`

GetItemsTargetMusclesOk returns a tuple with the ItemsTargetMuscles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTargetMuscles

`func (o *V1HealthExercisesGet200ResponseData) SetItemsTargetMuscles(v []interface{})`

SetItemsTargetMuscles sets ItemsTargetMuscles field to given value.

### HasItemsTargetMuscles

`func (o *V1HealthExercisesGet200ResponseData) HasItemsTargetMuscles() bool`

HasItemsTargetMuscles returns a boolean if a field has been set.

### GetItemsSecondaryMuscles

`func (o *V1HealthExercisesGet200ResponseData) GetItemsSecondaryMuscles() []interface{}`

GetItemsSecondaryMuscles returns the ItemsSecondaryMuscles field if non-nil, zero value otherwise.

### GetItemsSecondaryMusclesOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsSecondaryMusclesOk() (*[]interface{}, bool)`

GetItemsSecondaryMusclesOk returns a tuple with the ItemsSecondaryMuscles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsSecondaryMuscles

`func (o *V1HealthExercisesGet200ResponseData) SetItemsSecondaryMuscles(v []interface{})`

SetItemsSecondaryMuscles sets ItemsSecondaryMuscles field to given value.

### HasItemsSecondaryMuscles

`func (o *V1HealthExercisesGet200ResponseData) HasItemsSecondaryMuscles() bool`

HasItemsSecondaryMuscles returns a boolean if a field has been set.

### GetItemsInstructions

`func (o *V1HealthExercisesGet200ResponseData) GetItemsInstructions() []interface{}`

GetItemsInstructions returns the ItemsInstructions field if non-nil, zero value otherwise.

### GetItemsInstructionsOk

`func (o *V1HealthExercisesGet200ResponseData) GetItemsInstructionsOk() (*[]interface{}, bool)`

GetItemsInstructionsOk returns a tuple with the ItemsInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsInstructions

`func (o *V1HealthExercisesGet200ResponseData) SetItemsInstructions(v []interface{})`

SetItemsInstructions sets ItemsInstructions field to given value.

### HasItemsInstructions

`func (o *V1HealthExercisesGet200ResponseData) HasItemsInstructions() bool`

HasItemsInstructions returns a boolean if a field has been set.

### GetTotal

`func (o *V1HealthExercisesGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1HealthExercisesGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1HealthExercisesGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1HealthExercisesGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *V1HealthExercisesGet200ResponseData) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *V1HealthExercisesGet200ResponseData) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *V1HealthExercisesGet200ResponseData) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *V1HealthExercisesGet200ResponseData) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *V1HealthExercisesGet200ResponseData) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *V1HealthExercisesGet200ResponseData) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *V1HealthExercisesGet200ResponseData) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *V1HealthExercisesGet200ResponseData) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


