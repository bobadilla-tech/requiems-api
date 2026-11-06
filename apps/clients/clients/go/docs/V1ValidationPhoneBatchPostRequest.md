# V1ValidationPhoneBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numbers** | **[]interface{}** | Array of phone numbers to validate (min: 1, max: 50). Each must include the country calling code. | 

## Methods

### NewV1ValidationPhoneBatchPostRequest

`func NewV1ValidationPhoneBatchPostRequest(numbers []interface{}, ) *V1ValidationPhoneBatchPostRequest`

NewV1ValidationPhoneBatchPostRequest instantiates a new V1ValidationPhoneBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ValidationPhoneBatchPostRequestWithDefaults

`func NewV1ValidationPhoneBatchPostRequestWithDefaults() *V1ValidationPhoneBatchPostRequest`

NewV1ValidationPhoneBatchPostRequestWithDefaults instantiates a new V1ValidationPhoneBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumbers

`func (o *V1ValidationPhoneBatchPostRequest) GetNumbers() []interface{}`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *V1ValidationPhoneBatchPostRequest) GetNumbersOk() (*[]interface{}, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *V1ValidationPhoneBatchPostRequest) SetNumbers(v []interface{})`

SetNumbers sets Numbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


