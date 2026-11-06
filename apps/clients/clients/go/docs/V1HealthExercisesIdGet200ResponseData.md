# V1HealthExercisesIdGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique exercise identifier | [optional] 
**Name** | Pointer to **string** | Exercise name | [optional] 
**BodyParts** | Pointer to **[]interface{}** | Body part categories involved | [optional] 
**Equipment** | Pointer to **[]interface{}** | Equipment required | [optional] 
**TargetMuscles** | Pointer to **[]interface{}** | Primary muscles targeted | [optional] 
**SecondaryMuscles** | Pointer to **[]interface{}** | Secondary muscles engaged | [optional] 
**Instructions** | Pointer to **[]interface{}** | Ordered step-by-step instructions | [optional] 

## Methods

### NewV1HealthExercisesIdGet200ResponseData

`func NewV1HealthExercisesIdGet200ResponseData() *V1HealthExercisesIdGet200ResponseData`

NewV1HealthExercisesIdGet200ResponseData instantiates a new V1HealthExercisesIdGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HealthExercisesIdGet200ResponseDataWithDefaults

`func NewV1HealthExercisesIdGet200ResponseDataWithDefaults() *V1HealthExercisesIdGet200ResponseData`

NewV1HealthExercisesIdGet200ResponseDataWithDefaults instantiates a new V1HealthExercisesIdGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1HealthExercisesIdGet200ResponseData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1HealthExercisesIdGet200ResponseData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1HealthExercisesIdGet200ResponseData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V1HealthExercisesIdGet200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1HealthExercisesIdGet200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1HealthExercisesIdGet200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1HealthExercisesIdGet200ResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1HealthExercisesIdGet200ResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBodyParts

`func (o *V1HealthExercisesIdGet200ResponseData) GetBodyParts() []interface{}`

GetBodyParts returns the BodyParts field if non-nil, zero value otherwise.

### GetBodyPartsOk

`func (o *V1HealthExercisesIdGet200ResponseData) GetBodyPartsOk() (*[]interface{}, bool)`

GetBodyPartsOk returns a tuple with the BodyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyParts

`func (o *V1HealthExercisesIdGet200ResponseData) SetBodyParts(v []interface{})`

SetBodyParts sets BodyParts field to given value.

### HasBodyParts

`func (o *V1HealthExercisesIdGet200ResponseData) HasBodyParts() bool`

HasBodyParts returns a boolean if a field has been set.

### GetEquipment

`func (o *V1HealthExercisesIdGet200ResponseData) GetEquipment() []interface{}`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *V1HealthExercisesIdGet200ResponseData) GetEquipmentOk() (*[]interface{}, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *V1HealthExercisesIdGet200ResponseData) SetEquipment(v []interface{})`

SetEquipment sets Equipment field to given value.

### HasEquipment

`func (o *V1HealthExercisesIdGet200ResponseData) HasEquipment() bool`

HasEquipment returns a boolean if a field has been set.

### GetTargetMuscles

`func (o *V1HealthExercisesIdGet200ResponseData) GetTargetMuscles() []interface{}`

GetTargetMuscles returns the TargetMuscles field if non-nil, zero value otherwise.

### GetTargetMusclesOk

`func (o *V1HealthExercisesIdGet200ResponseData) GetTargetMusclesOk() (*[]interface{}, bool)`

GetTargetMusclesOk returns a tuple with the TargetMuscles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMuscles

`func (o *V1HealthExercisesIdGet200ResponseData) SetTargetMuscles(v []interface{})`

SetTargetMuscles sets TargetMuscles field to given value.

### HasTargetMuscles

`func (o *V1HealthExercisesIdGet200ResponseData) HasTargetMuscles() bool`

HasTargetMuscles returns a boolean if a field has been set.

### GetSecondaryMuscles

`func (o *V1HealthExercisesIdGet200ResponseData) GetSecondaryMuscles() []interface{}`

GetSecondaryMuscles returns the SecondaryMuscles field if non-nil, zero value otherwise.

### GetSecondaryMusclesOk

`func (o *V1HealthExercisesIdGet200ResponseData) GetSecondaryMusclesOk() (*[]interface{}, bool)`

GetSecondaryMusclesOk returns a tuple with the SecondaryMuscles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMuscles

`func (o *V1HealthExercisesIdGet200ResponseData) SetSecondaryMuscles(v []interface{})`

SetSecondaryMuscles sets SecondaryMuscles field to given value.

### HasSecondaryMuscles

`func (o *V1HealthExercisesIdGet200ResponseData) HasSecondaryMuscles() bool`

HasSecondaryMuscles returns a boolean if a field has been set.

### GetInstructions

`func (o *V1HealthExercisesIdGet200ResponseData) GetInstructions() []interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *V1HealthExercisesIdGet200ResponseData) GetInstructionsOk() (*[]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *V1HealthExercisesIdGet200ResponseData) SetInstructions(v []interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *V1HealthExercisesIdGet200ResponseData) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


