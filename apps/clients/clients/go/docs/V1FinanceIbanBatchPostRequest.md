# V1FinanceIbanBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numbers** | **[]interface{}** | Array of iban numbers to validate (min: 1, max: 50). | 

## Methods

### NewV1FinanceIbanBatchPostRequest

`func NewV1FinanceIbanBatchPostRequest(numbers []interface{}, ) *V1FinanceIbanBatchPostRequest`

NewV1FinanceIbanBatchPostRequest instantiates a new V1FinanceIbanBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceIbanBatchPostRequestWithDefaults

`func NewV1FinanceIbanBatchPostRequestWithDefaults() *V1FinanceIbanBatchPostRequest`

NewV1FinanceIbanBatchPostRequestWithDefaults instantiates a new V1FinanceIbanBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumbers

`func (o *V1FinanceIbanBatchPostRequest) GetNumbers() []interface{}`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *V1FinanceIbanBatchPostRequest) GetNumbersOk() (*[]interface{}, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *V1FinanceIbanBatchPostRequest) SetNumbers(v []interface{})`

SetNumbers sets Numbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


