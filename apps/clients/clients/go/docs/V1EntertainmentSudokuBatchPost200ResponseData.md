# V1EntertainmentSudokuBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Generated puzzles in the same order as the input array. Each item has the same fields as the single-puzzle endpoint. | [optional] 
**ResultsDifficulty** | Pointer to **string** | The difficulty level of the puzzle (easy, medium, or hard) | [optional] 
**ResultsPuzzle** | Pointer to **string** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] 
**ResultsSolution** | Pointer to **string** | 9×9 grid containing the complete, valid solution | [optional] 
**Total** | Pointer to **int32** | Number of puzzles returned. Matches the length of the input array. | [optional] 

## Methods

### NewV1EntertainmentSudokuBatchPost200ResponseData

`func NewV1EntertainmentSudokuBatchPost200ResponseData() *V1EntertainmentSudokuBatchPost200ResponseData`

NewV1EntertainmentSudokuBatchPost200ResponseData instantiates a new V1EntertainmentSudokuBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EntertainmentSudokuBatchPost200ResponseDataWithDefaults

`func NewV1EntertainmentSudokuBatchPost200ResponseDataWithDefaults() *V1EntertainmentSudokuBatchPost200ResponseData`

NewV1EntertainmentSudokuBatchPost200ResponseDataWithDefaults instantiates a new V1EntertainmentSudokuBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsDifficulty

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResultsDifficulty() string`

GetResultsDifficulty returns the ResultsDifficulty field if non-nil, zero value otherwise.

### GetResultsDifficultyOk

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResultsDifficultyOk() (*string, bool)`

GetResultsDifficultyOk returns a tuple with the ResultsDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDifficulty

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) SetResultsDifficulty(v string)`

SetResultsDifficulty sets ResultsDifficulty field to given value.

### HasResultsDifficulty

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) HasResultsDifficulty() bool`

HasResultsDifficulty returns a boolean if a field has been set.

### GetResultsPuzzle

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResultsPuzzle() string`

GetResultsPuzzle returns the ResultsPuzzle field if non-nil, zero value otherwise.

### GetResultsPuzzleOk

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResultsPuzzleOk() (*string, bool)`

GetResultsPuzzleOk returns a tuple with the ResultsPuzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPuzzle

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) SetResultsPuzzle(v string)`

SetResultsPuzzle sets ResultsPuzzle field to given value.

### HasResultsPuzzle

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) HasResultsPuzzle() bool`

HasResultsPuzzle returns a boolean if a field has been set.

### GetResultsSolution

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResultsSolution() string`

GetResultsSolution returns the ResultsSolution field if non-nil, zero value otherwise.

### GetResultsSolutionOk

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetResultsSolutionOk() (*string, bool)`

GetResultsSolutionOk returns a tuple with the ResultsSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsSolution

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) SetResultsSolution(v string)`

SetResultsSolution sets ResultsSolution field to given value.

### HasResultsSolution

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) HasResultsSolution() bool`

HasResultsSolution returns a boolean if a field has been set.

### GetTotal

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1EntertainmentSudokuBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


