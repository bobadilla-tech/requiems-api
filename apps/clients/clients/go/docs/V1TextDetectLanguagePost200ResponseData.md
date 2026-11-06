# V1TextDetectLanguagePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** | Full name of the detected language (e.g. French, English, Spanish) | [optional] 
**Code** | Pointer to **string** | ISO 639-1 two-letter language code (e.g. fr, en, es). Empty string when detection is unreliable. | [optional] 
**Confidence** | Pointer to **string** | Confidence score between 0.0 and 1.0. 0.0 is returned when the language cannot be reliably detected. | [optional] 

## Methods

### NewV1TextDetectLanguagePost200ResponseData

`func NewV1TextDetectLanguagePost200ResponseData() *V1TextDetectLanguagePost200ResponseData`

NewV1TextDetectLanguagePost200ResponseData instantiates a new V1TextDetectLanguagePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextDetectLanguagePost200ResponseDataWithDefaults

`func NewV1TextDetectLanguagePost200ResponseDataWithDefaults() *V1TextDetectLanguagePost200ResponseData`

NewV1TextDetectLanguagePost200ResponseDataWithDefaults instantiates a new V1TextDetectLanguagePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *V1TextDetectLanguagePost200ResponseData) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *V1TextDetectLanguagePost200ResponseData) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *V1TextDetectLanguagePost200ResponseData) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *V1TextDetectLanguagePost200ResponseData) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCode

`func (o *V1TextDetectLanguagePost200ResponseData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *V1TextDetectLanguagePost200ResponseData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *V1TextDetectLanguagePost200ResponseData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *V1TextDetectLanguagePost200ResponseData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConfidence

`func (o *V1TextDetectLanguagePost200ResponseData) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *V1TextDetectLanguagePost200ResponseData) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *V1TextDetectLanguagePost200ResponseData) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *V1TextDetectLanguagePost200ResponseData) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


