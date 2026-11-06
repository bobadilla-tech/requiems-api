# V1TextThesaurusWordGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Word** | Pointer to **string** | The normalized (lowercased) word that was looked up | [optional] 
**Synonyms** | Pointer to **string** | List of words with similar meaning | [optional] 
**Antonyms** | Pointer to **string** | List of words with opposite meaning | [optional] 

## Methods

### NewV1TextThesaurusWordGet200ResponseData

`func NewV1TextThesaurusWordGet200ResponseData() *V1TextThesaurusWordGet200ResponseData`

NewV1TextThesaurusWordGet200ResponseData instantiates a new V1TextThesaurusWordGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextThesaurusWordGet200ResponseDataWithDefaults

`func NewV1TextThesaurusWordGet200ResponseDataWithDefaults() *V1TextThesaurusWordGet200ResponseData`

NewV1TextThesaurusWordGet200ResponseDataWithDefaults instantiates a new V1TextThesaurusWordGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWord

`func (o *V1TextThesaurusWordGet200ResponseData) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *V1TextThesaurusWordGet200ResponseData) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *V1TextThesaurusWordGet200ResponseData) SetWord(v string)`

SetWord sets Word field to given value.

### HasWord

`func (o *V1TextThesaurusWordGet200ResponseData) HasWord() bool`

HasWord returns a boolean if a field has been set.

### GetSynonyms

`func (o *V1TextThesaurusWordGet200ResponseData) GetSynonyms() string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *V1TextThesaurusWordGet200ResponseData) GetSynonymsOk() (*string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *V1TextThesaurusWordGet200ResponseData) SetSynonyms(v string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *V1TextThesaurusWordGet200ResponseData) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetAntonyms

`func (o *V1TextThesaurusWordGet200ResponseData) GetAntonyms() string`

GetAntonyms returns the Antonyms field if non-nil, zero value otherwise.

### GetAntonymsOk

`func (o *V1TextThesaurusWordGet200ResponseData) GetAntonymsOk() (*string, bool)`

GetAntonymsOk returns a tuple with the Antonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntonyms

`func (o *V1TextThesaurusWordGet200ResponseData) SetAntonyms(v string)`

SetAntonyms sets Antonyms field to given value.

### HasAntonyms

`func (o *V1TextThesaurusWordGet200ResponseData) HasAntonyms() bool`

HasAntonyms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


