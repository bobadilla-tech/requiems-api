# V1HealthExercisesBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Exercises found for the given IDs, in input order. IDs not found are omitted. | [optional] 
**ResultsId** | Pointer to **int32** | Unique exercise identifier | [optional] 
**ResultsName** | Pointer to **string** | Exercise name | [optional] 
**ResultsBodyParts** | Pointer to **[]interface{}** | Body part categories involved | [optional] 
**ResultsEquipment** | Pointer to **[]interface{}** | Equipment required | [optional] 
**ResultsTargetMuscles** | Pointer to **[]interface{}** | Primary muscles targeted | [optional] 
**ResultsSecondaryMuscles** | Pointer to **[]interface{}** | Secondary muscles engaged | [optional] 
**ResultsInstructions** | Pointer to **[]interface{}** | Ordered step-by-step instructions | [optional] 
**Total** | Pointer to **int32** | Number of exercises returned | [optional] 

## Methods

### NewV1HealthExercisesBatchPost200ResponseData

`func NewV1HealthExercisesBatchPost200ResponseData() *V1HealthExercisesBatchPost200ResponseData`

NewV1HealthExercisesBatchPost200ResponseData instantiates a new V1HealthExercisesBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HealthExercisesBatchPost200ResponseDataWithDefaults

`func NewV1HealthExercisesBatchPost200ResponseDataWithDefaults() *V1HealthExercisesBatchPost200ResponseData`

NewV1HealthExercisesBatchPost200ResponseDataWithDefaults instantiates a new V1HealthExercisesBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsId

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsId() int32`

GetResultsId returns the ResultsId field if non-nil, zero value otherwise.

### GetResultsIdOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsIdOk() (*int32, bool)`

GetResultsIdOk returns a tuple with the ResultsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsId

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResultsId(v int32)`

SetResultsId sets ResultsId field to given value.

### HasResultsId

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResultsId() bool`

HasResultsId returns a boolean if a field has been set.

### GetResultsName

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsName() string`

GetResultsName returns the ResultsName field if non-nil, zero value otherwise.

### GetResultsNameOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsNameOk() (*string, bool)`

GetResultsNameOk returns a tuple with the ResultsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsName

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResultsName(v string)`

SetResultsName sets ResultsName field to given value.

### HasResultsName

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResultsName() bool`

HasResultsName returns a boolean if a field has been set.

### GetResultsBodyParts

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsBodyParts() []interface{}`

GetResultsBodyParts returns the ResultsBodyParts field if non-nil, zero value otherwise.

### GetResultsBodyPartsOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsBodyPartsOk() (*[]interface{}, bool)`

GetResultsBodyPartsOk returns a tuple with the ResultsBodyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsBodyParts

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResultsBodyParts(v []interface{})`

SetResultsBodyParts sets ResultsBodyParts field to given value.

### HasResultsBodyParts

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResultsBodyParts() bool`

HasResultsBodyParts returns a boolean if a field has been set.

### GetResultsEquipment

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsEquipment() []interface{}`

GetResultsEquipment returns the ResultsEquipment field if non-nil, zero value otherwise.

### GetResultsEquipmentOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsEquipmentOk() (*[]interface{}, bool)`

GetResultsEquipmentOk returns a tuple with the ResultsEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsEquipment

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResultsEquipment(v []interface{})`

SetResultsEquipment sets ResultsEquipment field to given value.

### HasResultsEquipment

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResultsEquipment() bool`

HasResultsEquipment returns a boolean if a field has been set.

### GetResultsTargetMuscles

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsTargetMuscles() []interface{}`

GetResultsTargetMuscles returns the ResultsTargetMuscles field if non-nil, zero value otherwise.

### GetResultsTargetMusclesOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsTargetMusclesOk() (*[]interface{}, bool)`

GetResultsTargetMusclesOk returns a tuple with the ResultsTargetMuscles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsTargetMuscles

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResultsTargetMuscles(v []interface{})`

SetResultsTargetMuscles sets ResultsTargetMuscles field to given value.

### HasResultsTargetMuscles

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResultsTargetMuscles() bool`

HasResultsTargetMuscles returns a boolean if a field has been set.

### GetResultsSecondaryMuscles

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsSecondaryMuscles() []interface{}`

GetResultsSecondaryMuscles returns the ResultsSecondaryMuscles field if non-nil, zero value otherwise.

### GetResultsSecondaryMusclesOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsSecondaryMusclesOk() (*[]interface{}, bool)`

GetResultsSecondaryMusclesOk returns a tuple with the ResultsSecondaryMuscles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsSecondaryMuscles

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResultsSecondaryMuscles(v []interface{})`

SetResultsSecondaryMuscles sets ResultsSecondaryMuscles field to given value.

### HasResultsSecondaryMuscles

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResultsSecondaryMuscles() bool`

HasResultsSecondaryMuscles returns a boolean if a field has been set.

### GetResultsInstructions

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsInstructions() []interface{}`

GetResultsInstructions returns the ResultsInstructions field if non-nil, zero value otherwise.

### GetResultsInstructionsOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetResultsInstructionsOk() (*[]interface{}, bool)`

GetResultsInstructionsOk returns a tuple with the ResultsInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsInstructions

`func (o *V1HealthExercisesBatchPost200ResponseData) SetResultsInstructions(v []interface{})`

SetResultsInstructions sets ResultsInstructions field to given value.

### HasResultsInstructions

`func (o *V1HealthExercisesBatchPost200ResponseData) HasResultsInstructions() bool`

HasResultsInstructions returns a boolean if a field has been set.

### GetTotal

`func (o *V1HealthExercisesBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1HealthExercisesBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1HealthExercisesBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1HealthExercisesBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


