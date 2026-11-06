# V1TechnologyMarkdownPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Markdown** | **string** | The Markdown text to convert. | 
**Sanitize** | Pointer to **bool** | When true, sanitizes the HTML output to remove unsafe tags and attributes. | [optional] 

## Methods

### NewV1TechnologyMarkdownPostRequest

`func NewV1TechnologyMarkdownPostRequest(markdown string, ) *V1TechnologyMarkdownPostRequest`

NewV1TechnologyMarkdownPostRequest instantiates a new V1TechnologyMarkdownPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyMarkdownPostRequestWithDefaults

`func NewV1TechnologyMarkdownPostRequestWithDefaults() *V1TechnologyMarkdownPostRequest`

NewV1TechnologyMarkdownPostRequestWithDefaults instantiates a new V1TechnologyMarkdownPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkdown

`func (o *V1TechnologyMarkdownPostRequest) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *V1TechnologyMarkdownPostRequest) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *V1TechnologyMarkdownPostRequest) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetSanitize

`func (o *V1TechnologyMarkdownPostRequest) GetSanitize() bool`

GetSanitize returns the Sanitize field if non-nil, zero value otherwise.

### GetSanitizeOk

`func (o *V1TechnologyMarkdownPostRequest) GetSanitizeOk() (*bool, bool)`

GetSanitizeOk returns a tuple with the Sanitize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitize

`func (o *V1TechnologyMarkdownPostRequest) SetSanitize(v bool)`

SetSanitize sets Sanitize field to given value.

### HasSanitize

`func (o *V1TechnologyMarkdownPostRequest) HasSanitize() bool`

HasSanitize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


