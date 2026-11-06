# V1TextSentimentBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Sentiment results for each input text, in the same order as the input. | [optional] 
**ResultsSentiment** | Pointer to **string** | The dominant sentiment class: positive, negative, or neutral | [optional] 
**ResultsScore** | Pointer to **float32** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] 
**ResultsBreakdownPositive** | Pointer to **float32** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] 
**ResultsBreakdownNegative** | Pointer to **float32** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] 
**ResultsBreakdownNeutral** | Pointer to **float32** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] 
**Total** | Pointer to **float32** | Total number of results returned. | [optional] 

## Methods

### NewV1TextSentimentBatchPost200ResponseData

`func NewV1TextSentimentBatchPost200ResponseData() *V1TextSentimentBatchPost200ResponseData`

NewV1TextSentimentBatchPost200ResponseData instantiates a new V1TextSentimentBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextSentimentBatchPost200ResponseDataWithDefaults

`func NewV1TextSentimentBatchPost200ResponseDataWithDefaults() *V1TextSentimentBatchPost200ResponseData`

NewV1TextSentimentBatchPost200ResponseDataWithDefaults instantiates a new V1TextSentimentBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1TextSentimentBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1TextSentimentBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1TextSentimentBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsSentiment

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsSentiment() string`

GetResultsSentiment returns the ResultsSentiment field if non-nil, zero value otherwise.

### GetResultsSentimentOk

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsSentimentOk() (*string, bool)`

GetResultsSentimentOk returns a tuple with the ResultsSentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsSentiment

`func (o *V1TextSentimentBatchPost200ResponseData) SetResultsSentiment(v string)`

SetResultsSentiment sets ResultsSentiment field to given value.

### HasResultsSentiment

`func (o *V1TextSentimentBatchPost200ResponseData) HasResultsSentiment() bool`

HasResultsSentiment returns a boolean if a field has been set.

### GetResultsScore

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsScore() float32`

GetResultsScore returns the ResultsScore field if non-nil, zero value otherwise.

### GetResultsScoreOk

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsScoreOk() (*float32, bool)`

GetResultsScoreOk returns a tuple with the ResultsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsScore

`func (o *V1TextSentimentBatchPost200ResponseData) SetResultsScore(v float32)`

SetResultsScore sets ResultsScore field to given value.

### HasResultsScore

`func (o *V1TextSentimentBatchPost200ResponseData) HasResultsScore() bool`

HasResultsScore returns a boolean if a field has been set.

### GetResultsBreakdownPositive

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsBreakdownPositive() float32`

GetResultsBreakdownPositive returns the ResultsBreakdownPositive field if non-nil, zero value otherwise.

### GetResultsBreakdownPositiveOk

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsBreakdownPositiveOk() (*float32, bool)`

GetResultsBreakdownPositiveOk returns a tuple with the ResultsBreakdownPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsBreakdownPositive

`func (o *V1TextSentimentBatchPost200ResponseData) SetResultsBreakdownPositive(v float32)`

SetResultsBreakdownPositive sets ResultsBreakdownPositive field to given value.

### HasResultsBreakdownPositive

`func (o *V1TextSentimentBatchPost200ResponseData) HasResultsBreakdownPositive() bool`

HasResultsBreakdownPositive returns a boolean if a field has been set.

### GetResultsBreakdownNegative

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsBreakdownNegative() float32`

GetResultsBreakdownNegative returns the ResultsBreakdownNegative field if non-nil, zero value otherwise.

### GetResultsBreakdownNegativeOk

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsBreakdownNegativeOk() (*float32, bool)`

GetResultsBreakdownNegativeOk returns a tuple with the ResultsBreakdownNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsBreakdownNegative

`func (o *V1TextSentimentBatchPost200ResponseData) SetResultsBreakdownNegative(v float32)`

SetResultsBreakdownNegative sets ResultsBreakdownNegative field to given value.

### HasResultsBreakdownNegative

`func (o *V1TextSentimentBatchPost200ResponseData) HasResultsBreakdownNegative() bool`

HasResultsBreakdownNegative returns a boolean if a field has been set.

### GetResultsBreakdownNeutral

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsBreakdownNeutral() float32`

GetResultsBreakdownNeutral returns the ResultsBreakdownNeutral field if non-nil, zero value otherwise.

### GetResultsBreakdownNeutralOk

`func (o *V1TextSentimentBatchPost200ResponseData) GetResultsBreakdownNeutralOk() (*float32, bool)`

GetResultsBreakdownNeutralOk returns a tuple with the ResultsBreakdownNeutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsBreakdownNeutral

`func (o *V1TextSentimentBatchPost200ResponseData) SetResultsBreakdownNeutral(v float32)`

SetResultsBreakdownNeutral sets ResultsBreakdownNeutral field to given value.

### HasResultsBreakdownNeutral

`func (o *V1TextSentimentBatchPost200ResponseData) HasResultsBreakdownNeutral() bool`

HasResultsBreakdownNeutral returns a boolean if a field has been set.

### GetTotal

`func (o *V1TextSentimentBatchPost200ResponseData) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1TextSentimentBatchPost200ResponseData) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1TextSentimentBatchPost200ResponseData) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1TextSentimentBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


