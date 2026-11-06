# V1FinanceMortgageBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mortgages** | **[]interface{}** | Array of mortgages to calculate (min: 1, max: 50). | 

## Methods

### NewV1FinanceMortgageBatchPostRequest

`func NewV1FinanceMortgageBatchPostRequest(mortgages []interface{}, ) *V1FinanceMortgageBatchPostRequest`

NewV1FinanceMortgageBatchPostRequest instantiates a new V1FinanceMortgageBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceMortgageBatchPostRequestWithDefaults

`func NewV1FinanceMortgageBatchPostRequestWithDefaults() *V1FinanceMortgageBatchPostRequest`

NewV1FinanceMortgageBatchPostRequestWithDefaults instantiates a new V1FinanceMortgageBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMortgages

`func (o *V1FinanceMortgageBatchPostRequest) GetMortgages() []interface{}`

GetMortgages returns the Mortgages field if non-nil, zero value otherwise.

### GetMortgagesOk

`func (o *V1FinanceMortgageBatchPostRequest) GetMortgagesOk() (*[]interface{}, bool)`

GetMortgagesOk returns a tuple with the Mortgages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMortgages

`func (o *V1FinanceMortgageBatchPostRequest) SetMortgages(v []interface{})`

SetMortgages sets Mortgages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


