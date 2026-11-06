# V1TextSimilarityPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Similarity** | Pointer to **float32** | Cosine similarity score between the two texts, in the range [0, 1]. | [optional] 
**Method** | Pointer to **string** | The algorithm used. Currently always &#39;cosine&#39;. | [optional] 

## Methods

### NewV1TextSimilarityPost200ResponseData

`func NewV1TextSimilarityPost200ResponseData() *V1TextSimilarityPost200ResponseData`

NewV1TextSimilarityPost200ResponseData instantiates a new V1TextSimilarityPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextSimilarityPost200ResponseDataWithDefaults

`func NewV1TextSimilarityPost200ResponseDataWithDefaults() *V1TextSimilarityPost200ResponseData`

NewV1TextSimilarityPost200ResponseDataWithDefaults instantiates a new V1TextSimilarityPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimilarity

`func (o *V1TextSimilarityPost200ResponseData) GetSimilarity() float32`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *V1TextSimilarityPost200ResponseData) GetSimilarityOk() (*float32, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *V1TextSimilarityPost200ResponseData) SetSimilarity(v float32)`

SetSimilarity sets Similarity field to given value.

### HasSimilarity

`func (o *V1TextSimilarityPost200ResponseData) HasSimilarity() bool`

HasSimilarity returns a boolean if a field has been set.

### GetMethod

`func (o *V1TextSimilarityPost200ResponseData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V1TextSimilarityPost200ResponseData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V1TextSimilarityPost200ResponseData) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V1TextSimilarityPost200ResponseData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


