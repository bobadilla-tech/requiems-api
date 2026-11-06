# V1EntertainmentSudokuBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Puzzles** | **[]interface{}** | Array of difficulty levels to generate (min: 1, max: 20). Each must be one of: easy, medium, hard. | 

## Methods

### NewV1EntertainmentSudokuBatchPostRequest

`func NewV1EntertainmentSudokuBatchPostRequest(puzzles []interface{}, ) *V1EntertainmentSudokuBatchPostRequest`

NewV1EntertainmentSudokuBatchPostRequest instantiates a new V1EntertainmentSudokuBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EntertainmentSudokuBatchPostRequestWithDefaults

`func NewV1EntertainmentSudokuBatchPostRequestWithDefaults() *V1EntertainmentSudokuBatchPostRequest`

NewV1EntertainmentSudokuBatchPostRequestWithDefaults instantiates a new V1EntertainmentSudokuBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPuzzles

`func (o *V1EntertainmentSudokuBatchPostRequest) GetPuzzles() []interface{}`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *V1EntertainmentSudokuBatchPostRequest) GetPuzzlesOk() (*[]interface{}, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *V1EntertainmentSudokuBatchPostRequest) SetPuzzles(v []interface{})`

SetPuzzles sets Puzzles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


