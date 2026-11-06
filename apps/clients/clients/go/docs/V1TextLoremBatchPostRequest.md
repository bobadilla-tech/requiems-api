# V1TextLoremBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]interface{}** | Array of text generation requests. Minimum 1 item, maximum 50 items. | 
**ItemsParagraphs** | Pointer to **int32** | Number of paragraphs to generate (1-20) | [optional] 
**ItemsSentences** | Pointer to **int32** | Number of sentences per paragraph (1-20) | [optional] 

## Methods

### NewV1TextLoremBatchPostRequest

`func NewV1TextLoremBatchPostRequest(items []interface{}, ) *V1TextLoremBatchPostRequest`

NewV1TextLoremBatchPostRequest instantiates a new V1TextLoremBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextLoremBatchPostRequestWithDefaults

`func NewV1TextLoremBatchPostRequestWithDefaults() *V1TextLoremBatchPostRequest`

NewV1TextLoremBatchPostRequestWithDefaults instantiates a new V1TextLoremBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1TextLoremBatchPostRequest) GetItems() []interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1TextLoremBatchPostRequest) GetItemsOk() (*[]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1TextLoremBatchPostRequest) SetItems(v []interface{})`

SetItems sets Items field to given value.


### GetItemsParagraphs

`func (o *V1TextLoremBatchPostRequest) GetItemsParagraphs() int32`

GetItemsParagraphs returns the ItemsParagraphs field if non-nil, zero value otherwise.

### GetItemsParagraphsOk

`func (o *V1TextLoremBatchPostRequest) GetItemsParagraphsOk() (*int32, bool)`

GetItemsParagraphsOk returns a tuple with the ItemsParagraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsParagraphs

`func (o *V1TextLoremBatchPostRequest) SetItemsParagraphs(v int32)`

SetItemsParagraphs sets ItemsParagraphs field to given value.

### HasItemsParagraphs

`func (o *V1TextLoremBatchPostRequest) HasItemsParagraphs() bool`

HasItemsParagraphs returns a boolean if a field has been set.

### GetItemsSentences

`func (o *V1TextLoremBatchPostRequest) GetItemsSentences() int32`

GetItemsSentences returns the ItemsSentences field if non-nil, zero value otherwise.

### GetItemsSentencesOk

`func (o *V1TextLoremBatchPostRequest) GetItemsSentencesOk() (*int32, bool)`

GetItemsSentencesOk returns a tuple with the ItemsSentences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsSentences

`func (o *V1TextLoremBatchPostRequest) SetItemsSentences(v int32)`

SetItemsSentences sets ItemsSentences field to given value.

### HasItemsSentences

`func (o *V1TextLoremBatchPostRequest) HasItemsSentences() bool`

HasItemsSentences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


