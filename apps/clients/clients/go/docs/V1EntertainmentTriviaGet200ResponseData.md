# V1EntertainmentTriviaGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to **string** | The trivia question text | [optional] 
**Options** | Pointer to **string** | Four multiple-choice answer options | [optional] 
**Answer** | Pointer to **string** | The correct answer — always one of the values in options | [optional] 
**Category** | Pointer to **string** | The category the question belongs to | [optional] 
**Difficulty** | Pointer to **string** | The difficulty level of the question (easy, medium, or hard) | [optional] 

## Methods

### NewV1EntertainmentTriviaGet200ResponseData

`func NewV1EntertainmentTriviaGet200ResponseData() *V1EntertainmentTriviaGet200ResponseData`

NewV1EntertainmentTriviaGet200ResponseData instantiates a new V1EntertainmentTriviaGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EntertainmentTriviaGet200ResponseDataWithDefaults

`func NewV1EntertainmentTriviaGet200ResponseDataWithDefaults() *V1EntertainmentTriviaGet200ResponseData`

NewV1EntertainmentTriviaGet200ResponseDataWithDefaults instantiates a new V1EntertainmentTriviaGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *V1EntertainmentTriviaGet200ResponseData) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *V1EntertainmentTriviaGet200ResponseData) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *V1EntertainmentTriviaGet200ResponseData) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *V1EntertainmentTriviaGet200ResponseData) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetOptions

`func (o *V1EntertainmentTriviaGet200ResponseData) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *V1EntertainmentTriviaGet200ResponseData) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *V1EntertainmentTriviaGet200ResponseData) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *V1EntertainmentTriviaGet200ResponseData) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAnswer

`func (o *V1EntertainmentTriviaGet200ResponseData) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *V1EntertainmentTriviaGet200ResponseData) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *V1EntertainmentTriviaGet200ResponseData) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *V1EntertainmentTriviaGet200ResponseData) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetCategory

`func (o *V1EntertainmentTriviaGet200ResponseData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V1EntertainmentTriviaGet200ResponseData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V1EntertainmentTriviaGet200ResponseData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V1EntertainmentTriviaGet200ResponseData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDifficulty

`func (o *V1EntertainmentTriviaGet200ResponseData) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *V1EntertainmentTriviaGet200ResponseData) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *V1EntertainmentTriviaGet200ResponseData) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *V1EntertainmentTriviaGet200ResponseData) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


