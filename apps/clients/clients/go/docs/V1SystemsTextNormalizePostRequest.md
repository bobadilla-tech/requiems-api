# V1SystemsTextNormalizePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The text to normalize. Required. | 
**Operations** | **[]interface{}** | Ordered list of normalization operations to apply. At least one required. Valid values: trim, lowercase, uppercase, normalize_unicode, collapse_whitespace, strip_html, remove_punctuation | 

## Methods

### NewV1SystemsTextNormalizePostRequest

`func NewV1SystemsTextNormalizePostRequest(text string, operations []interface{}, ) *V1SystemsTextNormalizePostRequest`

NewV1SystemsTextNormalizePostRequest instantiates a new V1SystemsTextNormalizePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsTextNormalizePostRequestWithDefaults

`func NewV1SystemsTextNormalizePostRequestWithDefaults() *V1SystemsTextNormalizePostRequest`

NewV1SystemsTextNormalizePostRequestWithDefaults instantiates a new V1SystemsTextNormalizePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *V1SystemsTextNormalizePostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *V1SystemsTextNormalizePostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *V1SystemsTextNormalizePostRequest) SetText(v string)`

SetText sets Text field to given value.


### GetOperations

`func (o *V1SystemsTextNormalizePostRequest) GetOperations() []interface{}`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *V1SystemsTextNormalizePostRequest) GetOperationsOk() (*[]interface{}, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *V1SystemsTextNormalizePostRequest) SetOperations(v []interface{})`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


