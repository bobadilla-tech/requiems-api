# V1ValidationProfanityBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsText** | Pointer to **string** | The original text that was checked. | [optional] 
**ResultsResultHasProfanity** | Pointer to **bool** | Whether any profanity was detected. | [optional] 
**ResultsResultCensored** | Pointer to **string** | Text with profane words replaced by asterisks. | [optional] 
**ResultsResultFlaggedWords** | Pointer to **string** | Deduplicated list of detected profane words (lowercase). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1ValidationProfanityBatchPost200ResponseData

`func NewV1ValidationProfanityBatchPost200ResponseData() *V1ValidationProfanityBatchPost200ResponseData`

NewV1ValidationProfanityBatchPost200ResponseData instantiates a new V1ValidationProfanityBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ValidationProfanityBatchPost200ResponseDataWithDefaults

`func NewV1ValidationProfanityBatchPost200ResponseDataWithDefaults() *V1ValidationProfanityBatchPost200ResponseData`

NewV1ValidationProfanityBatchPost200ResponseDataWithDefaults instantiates a new V1ValidationProfanityBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1ValidationProfanityBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1ValidationProfanityBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsText

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsText() string`

GetResultsText returns the ResultsText field if non-nil, zero value otherwise.

### GetResultsTextOk

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsTextOk() (*string, bool)`

GetResultsTextOk returns a tuple with the ResultsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsText

`func (o *V1ValidationProfanityBatchPost200ResponseData) SetResultsText(v string)`

SetResultsText sets ResultsText field to given value.

### HasResultsText

`func (o *V1ValidationProfanityBatchPost200ResponseData) HasResultsText() bool`

HasResultsText returns a boolean if a field has been set.

### GetResultsResultHasProfanity

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsResultHasProfanity() bool`

GetResultsResultHasProfanity returns the ResultsResultHasProfanity field if non-nil, zero value otherwise.

### GetResultsResultHasProfanityOk

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsResultHasProfanityOk() (*bool, bool)`

GetResultsResultHasProfanityOk returns a tuple with the ResultsResultHasProfanity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResultHasProfanity

`func (o *V1ValidationProfanityBatchPost200ResponseData) SetResultsResultHasProfanity(v bool)`

SetResultsResultHasProfanity sets ResultsResultHasProfanity field to given value.

### HasResultsResultHasProfanity

`func (o *V1ValidationProfanityBatchPost200ResponseData) HasResultsResultHasProfanity() bool`

HasResultsResultHasProfanity returns a boolean if a field has been set.

### GetResultsResultCensored

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsResultCensored() string`

GetResultsResultCensored returns the ResultsResultCensored field if non-nil, zero value otherwise.

### GetResultsResultCensoredOk

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsResultCensoredOk() (*string, bool)`

GetResultsResultCensoredOk returns a tuple with the ResultsResultCensored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResultCensored

`func (o *V1ValidationProfanityBatchPost200ResponseData) SetResultsResultCensored(v string)`

SetResultsResultCensored sets ResultsResultCensored field to given value.

### HasResultsResultCensored

`func (o *V1ValidationProfanityBatchPost200ResponseData) HasResultsResultCensored() bool`

HasResultsResultCensored returns a boolean if a field has been set.

### GetResultsResultFlaggedWords

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsResultFlaggedWords() string`

GetResultsResultFlaggedWords returns the ResultsResultFlaggedWords field if non-nil, zero value otherwise.

### GetResultsResultFlaggedWordsOk

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetResultsResultFlaggedWordsOk() (*string, bool)`

GetResultsResultFlaggedWordsOk returns a tuple with the ResultsResultFlaggedWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResultFlaggedWords

`func (o *V1ValidationProfanityBatchPost200ResponseData) SetResultsResultFlaggedWords(v string)`

SetResultsResultFlaggedWords sets ResultsResultFlaggedWords field to given value.

### HasResultsResultFlaggedWords

`func (o *V1ValidationProfanityBatchPost200ResponseData) HasResultsResultFlaggedWords() bool`

HasResultsResultFlaggedWords returns a boolean if a field has been set.

### GetTotal

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1ValidationProfanityBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1ValidationProfanityBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1ValidationProfanityBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


