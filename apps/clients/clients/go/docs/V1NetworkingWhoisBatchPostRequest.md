# V1NetworkingWhoisBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | **[]interface{}** | Array of domain names. Min: 1, Max: 50. | 

## Methods

### NewV1NetworkingWhoisBatchPostRequest

`func NewV1NetworkingWhoisBatchPostRequest(domains []interface{}, ) *V1NetworkingWhoisBatchPostRequest`

NewV1NetworkingWhoisBatchPostRequest instantiates a new V1NetworkingWhoisBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingWhoisBatchPostRequestWithDefaults

`func NewV1NetworkingWhoisBatchPostRequestWithDefaults() *V1NetworkingWhoisBatchPostRequest`

NewV1NetworkingWhoisBatchPostRequestWithDefaults instantiates a new V1NetworkingWhoisBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *V1NetworkingWhoisBatchPostRequest) GetDomains() []interface{}`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *V1NetworkingWhoisBatchPostRequest) GetDomainsOk() (*[]interface{}, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *V1NetworkingWhoisBatchPostRequest) SetDomains(v []interface{})`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


