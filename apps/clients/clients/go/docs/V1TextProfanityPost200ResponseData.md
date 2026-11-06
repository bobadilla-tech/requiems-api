# V1TextProfanityPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasProfanity** | Pointer to **bool** | Whether any profanity was detected in the text | [optional] 
**Censored** | Pointer to **string** | The input text with profane words replaced by asterisks | [optional] 
**FlaggedWords** | Pointer to **string** | Deduplicated list of profane words found (lowercase) | [optional] 

## Methods

### NewV1TextProfanityPost200ResponseData

`func NewV1TextProfanityPost200ResponseData() *V1TextProfanityPost200ResponseData`

NewV1TextProfanityPost200ResponseData instantiates a new V1TextProfanityPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextProfanityPost200ResponseDataWithDefaults

`func NewV1TextProfanityPost200ResponseDataWithDefaults() *V1TextProfanityPost200ResponseData`

NewV1TextProfanityPost200ResponseDataWithDefaults instantiates a new V1TextProfanityPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasProfanity

`func (o *V1TextProfanityPost200ResponseData) GetHasProfanity() bool`

GetHasProfanity returns the HasProfanity field if non-nil, zero value otherwise.

### GetHasProfanityOk

`func (o *V1TextProfanityPost200ResponseData) GetHasProfanityOk() (*bool, bool)`

GetHasProfanityOk returns a tuple with the HasProfanity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProfanity

`func (o *V1TextProfanityPost200ResponseData) SetHasProfanity(v bool)`

SetHasProfanity sets HasProfanity field to given value.

### HasHasProfanity

`func (o *V1TextProfanityPost200ResponseData) HasHasProfanity() bool`

HasHasProfanity returns a boolean if a field has been set.

### GetCensored

`func (o *V1TextProfanityPost200ResponseData) GetCensored() string`

GetCensored returns the Censored field if non-nil, zero value otherwise.

### GetCensoredOk

`func (o *V1TextProfanityPost200ResponseData) GetCensoredOk() (*string, bool)`

GetCensoredOk returns a tuple with the Censored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCensored

`func (o *V1TextProfanityPost200ResponseData) SetCensored(v string)`

SetCensored sets Censored field to given value.

### HasCensored

`func (o *V1TextProfanityPost200ResponseData) HasCensored() bool`

HasCensored returns a boolean if a field has been set.

### GetFlaggedWords

`func (o *V1TextProfanityPost200ResponseData) GetFlaggedWords() string`

GetFlaggedWords returns the FlaggedWords field if non-nil, zero value otherwise.

### GetFlaggedWordsOk

`func (o *V1TextProfanityPost200ResponseData) GetFlaggedWordsOk() (*string, bool)`

GetFlaggedWordsOk returns a tuple with the FlaggedWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedWords

`func (o *V1TextProfanityPost200ResponseData) SetFlaggedWords(v string)`

SetFlaggedWords sets FlaggedWords field to given value.

### HasFlaggedWords

`func (o *V1TextProfanityPost200ResponseData) HasFlaggedWords() bool`

HasFlaggedWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


