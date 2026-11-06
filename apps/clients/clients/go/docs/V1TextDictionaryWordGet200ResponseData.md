# V1TextDictionaryWordGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Word** | Pointer to **string** | The normalized (lowercased) word that was looked up | [optional] 
**Phonetic** | Pointer to **string** | IPA phonetic transcription of the word (may be omitted if unavailable) | [optional] 
**Definitions** | Pointer to **string** | One or more definitions for the word, each with partOfSpeech, definition, and an optional example | [optional] 
**DefinitionsPartOfSpeech** | Pointer to **string** | Grammatical category (e.g. noun, verb, adjective) | [optional] 
**DefinitionsDefinition** | Pointer to **string** | Plain-text definition of the word | [optional] 
**DefinitionsExample** | Pointer to **string** | Example sentence using the word (may be omitted) | [optional] 
**Synonyms** | Pointer to **string** | List of words with similar meaning | [optional] 

## Methods

### NewV1TextDictionaryWordGet200ResponseData

`func NewV1TextDictionaryWordGet200ResponseData() *V1TextDictionaryWordGet200ResponseData`

NewV1TextDictionaryWordGet200ResponseData instantiates a new V1TextDictionaryWordGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextDictionaryWordGet200ResponseDataWithDefaults

`func NewV1TextDictionaryWordGet200ResponseDataWithDefaults() *V1TextDictionaryWordGet200ResponseData`

NewV1TextDictionaryWordGet200ResponseDataWithDefaults instantiates a new V1TextDictionaryWordGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWord

`func (o *V1TextDictionaryWordGet200ResponseData) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *V1TextDictionaryWordGet200ResponseData) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *V1TextDictionaryWordGet200ResponseData) SetWord(v string)`

SetWord sets Word field to given value.

### HasWord

`func (o *V1TextDictionaryWordGet200ResponseData) HasWord() bool`

HasWord returns a boolean if a field has been set.

### GetPhonetic

`func (o *V1TextDictionaryWordGet200ResponseData) GetPhonetic() string`

GetPhonetic returns the Phonetic field if non-nil, zero value otherwise.

### GetPhoneticOk

`func (o *V1TextDictionaryWordGet200ResponseData) GetPhoneticOk() (*string, bool)`

GetPhoneticOk returns a tuple with the Phonetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonetic

`func (o *V1TextDictionaryWordGet200ResponseData) SetPhonetic(v string)`

SetPhonetic sets Phonetic field to given value.

### HasPhonetic

`func (o *V1TextDictionaryWordGet200ResponseData) HasPhonetic() bool`

HasPhonetic returns a boolean if a field has been set.

### GetDefinitions

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitions() string`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitionsOk() (*string, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *V1TextDictionaryWordGet200ResponseData) SetDefinitions(v string)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *V1TextDictionaryWordGet200ResponseData) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetDefinitionsPartOfSpeech

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitionsPartOfSpeech() string`

GetDefinitionsPartOfSpeech returns the DefinitionsPartOfSpeech field if non-nil, zero value otherwise.

### GetDefinitionsPartOfSpeechOk

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitionsPartOfSpeechOk() (*string, bool)`

GetDefinitionsPartOfSpeechOk returns a tuple with the DefinitionsPartOfSpeech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionsPartOfSpeech

`func (o *V1TextDictionaryWordGet200ResponseData) SetDefinitionsPartOfSpeech(v string)`

SetDefinitionsPartOfSpeech sets DefinitionsPartOfSpeech field to given value.

### HasDefinitionsPartOfSpeech

`func (o *V1TextDictionaryWordGet200ResponseData) HasDefinitionsPartOfSpeech() bool`

HasDefinitionsPartOfSpeech returns a boolean if a field has been set.

### GetDefinitionsDefinition

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitionsDefinition() string`

GetDefinitionsDefinition returns the DefinitionsDefinition field if non-nil, zero value otherwise.

### GetDefinitionsDefinitionOk

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitionsDefinitionOk() (*string, bool)`

GetDefinitionsDefinitionOk returns a tuple with the DefinitionsDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionsDefinition

`func (o *V1TextDictionaryWordGet200ResponseData) SetDefinitionsDefinition(v string)`

SetDefinitionsDefinition sets DefinitionsDefinition field to given value.

### HasDefinitionsDefinition

`func (o *V1TextDictionaryWordGet200ResponseData) HasDefinitionsDefinition() bool`

HasDefinitionsDefinition returns a boolean if a field has been set.

### GetDefinitionsExample

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitionsExample() string`

GetDefinitionsExample returns the DefinitionsExample field if non-nil, zero value otherwise.

### GetDefinitionsExampleOk

`func (o *V1TextDictionaryWordGet200ResponseData) GetDefinitionsExampleOk() (*string, bool)`

GetDefinitionsExampleOk returns a tuple with the DefinitionsExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionsExample

`func (o *V1TextDictionaryWordGet200ResponseData) SetDefinitionsExample(v string)`

SetDefinitionsExample sets DefinitionsExample field to given value.

### HasDefinitionsExample

`func (o *V1TextDictionaryWordGet200ResponseData) HasDefinitionsExample() bool`

HasDefinitionsExample returns a boolean if a field has been set.

### GetSynonyms

`func (o *V1TextDictionaryWordGet200ResponseData) GetSynonyms() string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *V1TextDictionaryWordGet200ResponseData) GetSynonymsOk() (*string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *V1TextDictionaryWordGet200ResponseData) SetSynonyms(v string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *V1TextDictionaryWordGet200ResponseData) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


