# V1TextSentimentPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sentiment** | Pointer to **string** | The dominant sentiment class: positive, negative, or neutral | [optional] 
**Score** | Pointer to **float32** | Confidence score for the dominant sentiment, between 0.0 and 1.0 | [optional] 
**BreakdownPositive** | Pointer to **float32** | Proportional score for positive sentiment (sums to 1.0 with other classes) | [optional] 
**BreakdownNegative** | Pointer to **float32** | Proportional score for negative sentiment (sums to 1.0 with other classes) | [optional] 
**BreakdownNeutral** | Pointer to **float32** | Proportional score for neutral sentiment (sums to 1.0 with other classes) | [optional] 

## Methods

### NewV1TextSentimentPost200ResponseData

`func NewV1TextSentimentPost200ResponseData() *V1TextSentimentPost200ResponseData`

NewV1TextSentimentPost200ResponseData instantiates a new V1TextSentimentPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextSentimentPost200ResponseDataWithDefaults

`func NewV1TextSentimentPost200ResponseDataWithDefaults() *V1TextSentimentPost200ResponseData`

NewV1TextSentimentPost200ResponseDataWithDefaults instantiates a new V1TextSentimentPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSentiment

`func (o *V1TextSentimentPost200ResponseData) GetSentiment() string`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *V1TextSentimentPost200ResponseData) GetSentimentOk() (*string, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *V1TextSentimentPost200ResponseData) SetSentiment(v string)`

SetSentiment sets Sentiment field to given value.

### HasSentiment

`func (o *V1TextSentimentPost200ResponseData) HasSentiment() bool`

HasSentiment returns a boolean if a field has been set.

### GetScore

`func (o *V1TextSentimentPost200ResponseData) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *V1TextSentimentPost200ResponseData) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *V1TextSentimentPost200ResponseData) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *V1TextSentimentPost200ResponseData) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetBreakdownPositive

`func (o *V1TextSentimentPost200ResponseData) GetBreakdownPositive() float32`

GetBreakdownPositive returns the BreakdownPositive field if non-nil, zero value otherwise.

### GetBreakdownPositiveOk

`func (o *V1TextSentimentPost200ResponseData) GetBreakdownPositiveOk() (*float32, bool)`

GetBreakdownPositiveOk returns a tuple with the BreakdownPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownPositive

`func (o *V1TextSentimentPost200ResponseData) SetBreakdownPositive(v float32)`

SetBreakdownPositive sets BreakdownPositive field to given value.

### HasBreakdownPositive

`func (o *V1TextSentimentPost200ResponseData) HasBreakdownPositive() bool`

HasBreakdownPositive returns a boolean if a field has been set.

### GetBreakdownNegative

`func (o *V1TextSentimentPost200ResponseData) GetBreakdownNegative() float32`

GetBreakdownNegative returns the BreakdownNegative field if non-nil, zero value otherwise.

### GetBreakdownNegativeOk

`func (o *V1TextSentimentPost200ResponseData) GetBreakdownNegativeOk() (*float32, bool)`

GetBreakdownNegativeOk returns a tuple with the BreakdownNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownNegative

`func (o *V1TextSentimentPost200ResponseData) SetBreakdownNegative(v float32)`

SetBreakdownNegative sets BreakdownNegative field to given value.

### HasBreakdownNegative

`func (o *V1TextSentimentPost200ResponseData) HasBreakdownNegative() bool`

HasBreakdownNegative returns a boolean if a field has been set.

### GetBreakdownNeutral

`func (o *V1TextSentimentPost200ResponseData) GetBreakdownNeutral() float32`

GetBreakdownNeutral returns the BreakdownNeutral field if non-nil, zero value otherwise.

### GetBreakdownNeutralOk

`func (o *V1TextSentimentPost200ResponseData) GetBreakdownNeutralOk() (*float32, bool)`

GetBreakdownNeutralOk returns a tuple with the BreakdownNeutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownNeutral

`func (o *V1TextSentimentPost200ResponseData) SetBreakdownNeutral(v float32)`

SetBreakdownNeutral sets BreakdownNeutral field to given value.

### HasBreakdownNeutral

`func (o *V1TextSentimentPost200ResponseData) HasBreakdownNeutral() bool`

HasBreakdownNeutral returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


