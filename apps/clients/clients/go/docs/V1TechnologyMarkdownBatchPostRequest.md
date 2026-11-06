# V1TechnologyMarkdownBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Markdowns** | **[]interface{}** | Array of Markdown strings. Min: 1, Max: 50. | 
**Sanitize** | Pointer to **bool** | When true, sanitizes HTML output to remove unsafe tags and attributes. | [optional] 

## Methods

### NewV1TechnologyMarkdownBatchPostRequest

`func NewV1TechnologyMarkdownBatchPostRequest(markdowns []interface{}, ) *V1TechnologyMarkdownBatchPostRequest`

NewV1TechnologyMarkdownBatchPostRequest instantiates a new V1TechnologyMarkdownBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyMarkdownBatchPostRequestWithDefaults

`func NewV1TechnologyMarkdownBatchPostRequestWithDefaults() *V1TechnologyMarkdownBatchPostRequest`

NewV1TechnologyMarkdownBatchPostRequestWithDefaults instantiates a new V1TechnologyMarkdownBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkdowns

`func (o *V1TechnologyMarkdownBatchPostRequest) GetMarkdowns() []interface{}`

GetMarkdowns returns the Markdowns field if non-nil, zero value otherwise.

### GetMarkdownsOk

`func (o *V1TechnologyMarkdownBatchPostRequest) GetMarkdownsOk() (*[]interface{}, bool)`

GetMarkdownsOk returns a tuple with the Markdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdowns

`func (o *V1TechnologyMarkdownBatchPostRequest) SetMarkdowns(v []interface{})`

SetMarkdowns sets Markdowns field to given value.


### GetSanitize

`func (o *V1TechnologyMarkdownBatchPostRequest) GetSanitize() bool`

GetSanitize returns the Sanitize field if non-nil, zero value otherwise.

### GetSanitizeOk

`func (o *V1TechnologyMarkdownBatchPostRequest) GetSanitizeOk() (*bool, bool)`

GetSanitizeOk returns a tuple with the Sanitize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitize

`func (o *V1TechnologyMarkdownBatchPostRequest) SetSanitize(v bool)`

SetSanitize sets Sanitize field to given value.

### HasSanitize

`func (o *V1TechnologyMarkdownBatchPostRequest) HasSanitize() bool`

HasSanitize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


