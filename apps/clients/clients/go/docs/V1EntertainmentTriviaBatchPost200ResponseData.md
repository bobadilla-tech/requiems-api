# V1EntertainmentTriviaBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | List of trivia questions returned in the same order as the input filters. | [optional] 
**ResultsCategory** | Pointer to **string** | The topic filter used to retrieve the trivia question. | [optional] 
**ResultsDifficulty** | Pointer to **string** | The complexity level filter used to retrieve the trivia question. | [optional] 
**ResultsError** | Pointer to **string** | Describes why the question could not be retrieved for this filter. Omitted if the query was successful. | [optional] 
**ResultsData** | Pointer to **map[string]interface{}** | The trivia question returned for this filter. Empty if the query failed. | [optional] 
**ResultsDataQuestion** | Pointer to **string** | The trivia question text. | [optional] 
**ResultsDataOptions** | Pointer to **[]interface{}** | List of possible answers for the question. | [optional] 
**ResultsDataAnswer** | Pointer to **string** | The correct answer to the trivia question. | [optional] 
**ResultsDataCategory** | Pointer to **string** | The category the trivia question belongs to. | [optional] 
**ResultsDataDifficulty** | Pointer to **string** | The difficulty level of the trivia question. | [optional] 
**Total** | Pointer to **int32** | Total number of trivia results returned. | [optional] 

## Methods

### NewV1EntertainmentTriviaBatchPost200ResponseData

`func NewV1EntertainmentTriviaBatchPost200ResponseData() *V1EntertainmentTriviaBatchPost200ResponseData`

NewV1EntertainmentTriviaBatchPost200ResponseData instantiates a new V1EntertainmentTriviaBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EntertainmentTriviaBatchPost200ResponseDataWithDefaults

`func NewV1EntertainmentTriviaBatchPost200ResponseDataWithDefaults() *V1EntertainmentTriviaBatchPost200ResponseData`

NewV1EntertainmentTriviaBatchPost200ResponseDataWithDefaults instantiates a new V1EntertainmentTriviaBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsCategory

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsCategory() string`

GetResultsCategory returns the ResultsCategory field if non-nil, zero value otherwise.

### GetResultsCategoryOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsCategoryOk() (*string, bool)`

GetResultsCategoryOk returns a tuple with the ResultsCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsCategory

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsCategory(v string)`

SetResultsCategory sets ResultsCategory field to given value.

### HasResultsCategory

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsCategory() bool`

HasResultsCategory returns a boolean if a field has been set.

### GetResultsDifficulty

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDifficulty() string`

GetResultsDifficulty returns the ResultsDifficulty field if non-nil, zero value otherwise.

### GetResultsDifficultyOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDifficultyOk() (*string, bool)`

GetResultsDifficultyOk returns a tuple with the ResultsDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDifficulty

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsDifficulty(v string)`

SetResultsDifficulty sets ResultsDifficulty field to given value.

### HasResultsDifficulty

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsDifficulty() bool`

HasResultsDifficulty returns a boolean if a field has been set.

### GetResultsError

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetResultsData

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsData() map[string]interface{}`

GetResultsData returns the ResultsData field if non-nil, zero value otherwise.

### GetResultsDataOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataOk() (*map[string]interface{}, bool)`

GetResultsDataOk returns a tuple with the ResultsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsData

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsData(v map[string]interface{})`

SetResultsData sets ResultsData field to given value.

### HasResultsData

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsData() bool`

HasResultsData returns a boolean if a field has been set.

### GetResultsDataQuestion

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataQuestion() string`

GetResultsDataQuestion returns the ResultsDataQuestion field if non-nil, zero value otherwise.

### GetResultsDataQuestionOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataQuestionOk() (*string, bool)`

GetResultsDataQuestionOk returns a tuple with the ResultsDataQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataQuestion

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsDataQuestion(v string)`

SetResultsDataQuestion sets ResultsDataQuestion field to given value.

### HasResultsDataQuestion

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsDataQuestion() bool`

HasResultsDataQuestion returns a boolean if a field has been set.

### GetResultsDataOptions

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataOptions() []interface{}`

GetResultsDataOptions returns the ResultsDataOptions field if non-nil, zero value otherwise.

### GetResultsDataOptionsOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataOptionsOk() (*[]interface{}, bool)`

GetResultsDataOptionsOk returns a tuple with the ResultsDataOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataOptions

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsDataOptions(v []interface{})`

SetResultsDataOptions sets ResultsDataOptions field to given value.

### HasResultsDataOptions

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsDataOptions() bool`

HasResultsDataOptions returns a boolean if a field has been set.

### GetResultsDataAnswer

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataAnswer() string`

GetResultsDataAnswer returns the ResultsDataAnswer field if non-nil, zero value otherwise.

### GetResultsDataAnswerOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataAnswerOk() (*string, bool)`

GetResultsDataAnswerOk returns a tuple with the ResultsDataAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataAnswer

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsDataAnswer(v string)`

SetResultsDataAnswer sets ResultsDataAnswer field to given value.

### HasResultsDataAnswer

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsDataAnswer() bool`

HasResultsDataAnswer returns a boolean if a field has been set.

### GetResultsDataCategory

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataCategory() string`

GetResultsDataCategory returns the ResultsDataCategory field if non-nil, zero value otherwise.

### GetResultsDataCategoryOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataCategoryOk() (*string, bool)`

GetResultsDataCategoryOk returns a tuple with the ResultsDataCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataCategory

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsDataCategory(v string)`

SetResultsDataCategory sets ResultsDataCategory field to given value.

### HasResultsDataCategory

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsDataCategory() bool`

HasResultsDataCategory returns a boolean if a field has been set.

### GetResultsDataDifficulty

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataDifficulty() string`

GetResultsDataDifficulty returns the ResultsDataDifficulty field if non-nil, zero value otherwise.

### GetResultsDataDifficultyOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetResultsDataDifficultyOk() (*string, bool)`

GetResultsDataDifficultyOk returns a tuple with the ResultsDataDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataDifficulty

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetResultsDataDifficulty(v string)`

SetResultsDataDifficulty sets ResultsDataDifficulty field to given value.

### HasResultsDataDifficulty

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasResultsDataDifficulty() bool`

HasResultsDataDifficulty returns a boolean if a field has been set.

### GetTotal

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1EntertainmentTriviaBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


