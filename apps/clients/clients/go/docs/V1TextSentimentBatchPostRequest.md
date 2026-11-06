# V1TextSentimentBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Texts** | **[]interface{}** | The list of texts to analyze. Between 1 and 50 items. | 

## Methods

### NewV1TextSentimentBatchPostRequest

`func NewV1TextSentimentBatchPostRequest(texts []interface{}, ) *V1TextSentimentBatchPostRequest`

NewV1TextSentimentBatchPostRequest instantiates a new V1TextSentimentBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextSentimentBatchPostRequestWithDefaults

`func NewV1TextSentimentBatchPostRequestWithDefaults() *V1TextSentimentBatchPostRequest`

NewV1TextSentimentBatchPostRequestWithDefaults instantiates a new V1TextSentimentBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTexts

`func (o *V1TextSentimentBatchPostRequest) GetTexts() []interface{}`

GetTexts returns the Texts field if non-nil, zero value otherwise.

### GetTextsOk

`func (o *V1TextSentimentBatchPostRequest) GetTextsOk() (*[]interface{}, bool)`

GetTextsOk returns a tuple with the Texts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexts

`func (o *V1TextSentimentBatchPostRequest) SetTexts(v []interface{})`

SetTexts sets Texts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


