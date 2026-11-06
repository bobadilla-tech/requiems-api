# V1TechnologyFormatPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Source format. One of: json, yaml, csv, xml, toml | 
**To** | **string** | Target format. One of: json, yaml, csv, xml, toml | 
**Content** | **string** | The content to convert, serialized as a string in the source format. | 

## Methods

### NewV1TechnologyFormatPostRequest

`func NewV1TechnologyFormatPostRequest(from string, to string, content string, ) *V1TechnologyFormatPostRequest`

NewV1TechnologyFormatPostRequest instantiates a new V1TechnologyFormatPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyFormatPostRequestWithDefaults

`func NewV1TechnologyFormatPostRequestWithDefaults() *V1TechnologyFormatPostRequest`

NewV1TechnologyFormatPostRequestWithDefaults instantiates a new V1TechnologyFormatPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *V1TechnologyFormatPostRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1TechnologyFormatPostRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1TechnologyFormatPostRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *V1TechnologyFormatPostRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1TechnologyFormatPostRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1TechnologyFormatPostRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetContent

`func (o *V1TechnologyFormatPostRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V1TechnologyFormatPostRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V1TechnologyFormatPostRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


