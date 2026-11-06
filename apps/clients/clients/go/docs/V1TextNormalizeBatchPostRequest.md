# V1TextNormalizeBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]interface{}** | Array of addresses to normalize (min 1, max 100; each entry non-empty) | 

## Methods

### NewV1TextNormalizeBatchPostRequest

`func NewV1TextNormalizeBatchPostRequest(emails []interface{}, ) *V1TextNormalizeBatchPostRequest`

NewV1TextNormalizeBatchPostRequest instantiates a new V1TextNormalizeBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextNormalizeBatchPostRequestWithDefaults

`func NewV1TextNormalizeBatchPostRequestWithDefaults() *V1TextNormalizeBatchPostRequest`

NewV1TextNormalizeBatchPostRequestWithDefaults instantiates a new V1TextNormalizeBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *V1TextNormalizeBatchPostRequest) GetEmails() []interface{}`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *V1TextNormalizeBatchPostRequest) GetEmailsOk() (*[]interface{}, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *V1TextNormalizeBatchPostRequest) SetEmails(v []interface{})`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


