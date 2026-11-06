# V1EntertainmentSudokuGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | Pointer to **string** | The difficulty level of the returned puzzle (easy, medium, or hard) | [optional] 
**Puzzle** | Pointer to **string** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] 
**Solution** | Pointer to **string** | 9×9 grid containing the complete, valid solution | [optional] 

## Methods

### NewV1EntertainmentSudokuGet200ResponseData

`func NewV1EntertainmentSudokuGet200ResponseData() *V1EntertainmentSudokuGet200ResponseData`

NewV1EntertainmentSudokuGet200ResponseData instantiates a new V1EntertainmentSudokuGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EntertainmentSudokuGet200ResponseDataWithDefaults

`func NewV1EntertainmentSudokuGet200ResponseDataWithDefaults() *V1EntertainmentSudokuGet200ResponseData`

NewV1EntertainmentSudokuGet200ResponseDataWithDefaults instantiates a new V1EntertainmentSudokuGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *V1EntertainmentSudokuGet200ResponseData) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *V1EntertainmentSudokuGet200ResponseData) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *V1EntertainmentSudokuGet200ResponseData) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *V1EntertainmentSudokuGet200ResponseData) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetPuzzle

`func (o *V1EntertainmentSudokuGet200ResponseData) GetPuzzle() string`

GetPuzzle returns the Puzzle field if non-nil, zero value otherwise.

### GetPuzzleOk

`func (o *V1EntertainmentSudokuGet200ResponseData) GetPuzzleOk() (*string, bool)`

GetPuzzleOk returns a tuple with the Puzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzle

`func (o *V1EntertainmentSudokuGet200ResponseData) SetPuzzle(v string)`

SetPuzzle sets Puzzle field to given value.

### HasPuzzle

`func (o *V1EntertainmentSudokuGet200ResponseData) HasPuzzle() bool`

HasPuzzle returns a boolean if a field has been set.

### GetSolution

`func (o *V1EntertainmentSudokuGet200ResponseData) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *V1EntertainmentSudokuGet200ResponseData) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *V1EntertainmentSudokuGet200ResponseData) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *V1EntertainmentSudokuGet200ResponseData) HasSolution() bool`

HasSolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


