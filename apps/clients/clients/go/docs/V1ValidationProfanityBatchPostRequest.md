# V1ValidationProfanityBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Texts** | **string** | List of texts to check (1–50 items, each non-empty). | 

## Methods

### NewV1ValidationProfanityBatchPostRequest

`func NewV1ValidationProfanityBatchPostRequest(texts string, ) *V1ValidationProfanityBatchPostRequest`

NewV1ValidationProfanityBatchPostRequest instantiates a new V1ValidationProfanityBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ValidationProfanityBatchPostRequestWithDefaults

`func NewV1ValidationProfanityBatchPostRequestWithDefaults() *V1ValidationProfanityBatchPostRequest`

NewV1ValidationProfanityBatchPostRequestWithDefaults instantiates a new V1ValidationProfanityBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTexts

`func (o *V1ValidationProfanityBatchPostRequest) GetTexts() string`

GetTexts returns the Texts field if non-nil, zero value otherwise.

### GetTextsOk

`func (o *V1ValidationProfanityBatchPostRequest) GetTextsOk() (*string, bool)`

GetTextsOk returns a tuple with the Texts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexts

`func (o *V1ValidationProfanityBatchPostRequest) SetTexts(v string)`

SetTexts sets Texts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


