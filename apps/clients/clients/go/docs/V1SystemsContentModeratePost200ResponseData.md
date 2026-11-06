# V1SystemsContentModeratePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSafe** | Pointer to **bool** | False when any moderation category is triggered. True otherwise. | [optional] 
**ToxicityScore** | Pointer to **string** | Reserved for future use. Currently always null. | [optional] 
**Sentiment** | Pointer to **string** | Detected sentiment: positive, negative, or neutral | [optional] 
**Language** | Pointer to **string** | Detected or provided ISO 639-1 language code. Null when detection confidence is 0. | [optional] 
**LanguageConfidence** | Pointer to **string** | Confidence in the detected language from 0.0 to 1.0. 1.0 when language was explicitly provided. | [optional] 
**Flags** | Pointer to **[]interface{}** | Triggered moderation flags. Currently possible: text_profanity | [optional] 
**Categories** | Pointer to **map[string]interface{}** | Per-category boolean flags: profanity, hate_speech, spam, violence | [optional] 

## Methods

### NewV1SystemsContentModeratePost200ResponseData

`func NewV1SystemsContentModeratePost200ResponseData() *V1SystemsContentModeratePost200ResponseData`

NewV1SystemsContentModeratePost200ResponseData instantiates a new V1SystemsContentModeratePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsContentModeratePost200ResponseDataWithDefaults

`func NewV1SystemsContentModeratePost200ResponseDataWithDefaults() *V1SystemsContentModeratePost200ResponseData`

NewV1SystemsContentModeratePost200ResponseDataWithDefaults instantiates a new V1SystemsContentModeratePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSafe

`func (o *V1SystemsContentModeratePost200ResponseData) GetIsSafe() bool`

GetIsSafe returns the IsSafe field if non-nil, zero value otherwise.

### GetIsSafeOk

`func (o *V1SystemsContentModeratePost200ResponseData) GetIsSafeOk() (*bool, bool)`

GetIsSafeOk returns a tuple with the IsSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafe

`func (o *V1SystemsContentModeratePost200ResponseData) SetIsSafe(v bool)`

SetIsSafe sets IsSafe field to given value.

### HasIsSafe

`func (o *V1SystemsContentModeratePost200ResponseData) HasIsSafe() bool`

HasIsSafe returns a boolean if a field has been set.

### GetToxicityScore

`func (o *V1SystemsContentModeratePost200ResponseData) GetToxicityScore() string`

GetToxicityScore returns the ToxicityScore field if non-nil, zero value otherwise.

### GetToxicityScoreOk

`func (o *V1SystemsContentModeratePost200ResponseData) GetToxicityScoreOk() (*string, bool)`

GetToxicityScoreOk returns a tuple with the ToxicityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToxicityScore

`func (o *V1SystemsContentModeratePost200ResponseData) SetToxicityScore(v string)`

SetToxicityScore sets ToxicityScore field to given value.

### HasToxicityScore

`func (o *V1SystemsContentModeratePost200ResponseData) HasToxicityScore() bool`

HasToxicityScore returns a boolean if a field has been set.

### GetSentiment

`func (o *V1SystemsContentModeratePost200ResponseData) GetSentiment() string`

GetSentiment returns the Sentiment field if non-nil, zero value otherwise.

### GetSentimentOk

`func (o *V1SystemsContentModeratePost200ResponseData) GetSentimentOk() (*string, bool)`

GetSentimentOk returns a tuple with the Sentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentiment

`func (o *V1SystemsContentModeratePost200ResponseData) SetSentiment(v string)`

SetSentiment sets Sentiment field to given value.

### HasSentiment

`func (o *V1SystemsContentModeratePost200ResponseData) HasSentiment() bool`

HasSentiment returns a boolean if a field has been set.

### GetLanguage

`func (o *V1SystemsContentModeratePost200ResponseData) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *V1SystemsContentModeratePost200ResponseData) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *V1SystemsContentModeratePost200ResponseData) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *V1SystemsContentModeratePost200ResponseData) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLanguageConfidence

`func (o *V1SystemsContentModeratePost200ResponseData) GetLanguageConfidence() string`

GetLanguageConfidence returns the LanguageConfidence field if non-nil, zero value otherwise.

### GetLanguageConfidenceOk

`func (o *V1SystemsContentModeratePost200ResponseData) GetLanguageConfidenceOk() (*string, bool)`

GetLanguageConfidenceOk returns a tuple with the LanguageConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageConfidence

`func (o *V1SystemsContentModeratePost200ResponseData) SetLanguageConfidence(v string)`

SetLanguageConfidence sets LanguageConfidence field to given value.

### HasLanguageConfidence

`func (o *V1SystemsContentModeratePost200ResponseData) HasLanguageConfidence() bool`

HasLanguageConfidence returns a boolean if a field has been set.

### GetFlags

`func (o *V1SystemsContentModeratePost200ResponseData) GetFlags() []interface{}`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V1SystemsContentModeratePost200ResponseData) GetFlagsOk() (*[]interface{}, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V1SystemsContentModeratePost200ResponseData) SetFlags(v []interface{})`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V1SystemsContentModeratePost200ResponseData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetCategories

`func (o *V1SystemsContentModeratePost200ResponseData) GetCategories() map[string]interface{}`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *V1SystemsContentModeratePost200ResponseData) GetCategoriesOk() (*map[string]interface{}, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *V1SystemsContentModeratePost200ResponseData) SetCategories(v map[string]interface{})`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *V1SystemsContentModeratePost200ResponseData) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


