# V1TextSpellcheckPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Corrected** | Pointer to **string** | The full input text with all misspelled words replaced by their suggested corrections | [optional] 
**Corrections** | Pointer to **string** | List of individual corrections. Each item contains: original (the misspelled word), suggested (the correction), and position (0-based character offset in the original text) | [optional] 

## Methods

### NewV1TextSpellcheckPost200ResponseData

`func NewV1TextSpellcheckPost200ResponseData() *V1TextSpellcheckPost200ResponseData`

NewV1TextSpellcheckPost200ResponseData instantiates a new V1TextSpellcheckPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextSpellcheckPost200ResponseDataWithDefaults

`func NewV1TextSpellcheckPost200ResponseDataWithDefaults() *V1TextSpellcheckPost200ResponseData`

NewV1TextSpellcheckPost200ResponseDataWithDefaults instantiates a new V1TextSpellcheckPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrected

`func (o *V1TextSpellcheckPost200ResponseData) GetCorrected() string`

GetCorrected returns the Corrected field if non-nil, zero value otherwise.

### GetCorrectedOk

`func (o *V1TextSpellcheckPost200ResponseData) GetCorrectedOk() (*string, bool)`

GetCorrectedOk returns a tuple with the Corrected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrected

`func (o *V1TextSpellcheckPost200ResponseData) SetCorrected(v string)`

SetCorrected sets Corrected field to given value.

### HasCorrected

`func (o *V1TextSpellcheckPost200ResponseData) HasCorrected() bool`

HasCorrected returns a boolean if a field has been set.

### GetCorrections

`func (o *V1TextSpellcheckPost200ResponseData) GetCorrections() string`

GetCorrections returns the Corrections field if non-nil, zero value otherwise.

### GetCorrectionsOk

`func (o *V1TextSpellcheckPost200ResponseData) GetCorrectionsOk() (*string, bool)`

GetCorrectionsOk returns a tuple with the Corrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrections

`func (o *V1TextSpellcheckPost200ResponseData) SetCorrections(v string)`

SetCorrections sets Corrections field to given value.

### HasCorrections

`func (o *V1TextSpellcheckPost200ResponseData) HasCorrections() bool`

HasCorrections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


