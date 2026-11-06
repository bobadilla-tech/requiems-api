# V1SystemsContentModeratePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The text to moderate. Any language. Required. | 
**Language** | Pointer to **string** | ISO 639-1 language code (e.g. en, de, fr). When provided, language detection is skipped and confidence is set to 1.0. | [optional] 

## Methods

### NewV1SystemsContentModeratePostRequest

`func NewV1SystemsContentModeratePostRequest(text string, ) *V1SystemsContentModeratePostRequest`

NewV1SystemsContentModeratePostRequest instantiates a new V1SystemsContentModeratePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsContentModeratePostRequestWithDefaults

`func NewV1SystemsContentModeratePostRequestWithDefaults() *V1SystemsContentModeratePostRequest`

NewV1SystemsContentModeratePostRequestWithDefaults instantiates a new V1SystemsContentModeratePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *V1SystemsContentModeratePostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *V1SystemsContentModeratePostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *V1SystemsContentModeratePostRequest) SetText(v string)`

SetText sets Text field to given value.


### GetLanguage

`func (o *V1SystemsContentModeratePostRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *V1SystemsContentModeratePostRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *V1SystemsContentModeratePostRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *V1SystemsContentModeratePostRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


