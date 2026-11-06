# V1ValidationEmailBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]interface{}** | Array of email addresses to validate. Min: 1, Max: 50. | 

## Methods

### NewV1ValidationEmailBatchPostRequest

`func NewV1ValidationEmailBatchPostRequest(emails []interface{}, ) *V1ValidationEmailBatchPostRequest`

NewV1ValidationEmailBatchPostRequest instantiates a new V1ValidationEmailBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ValidationEmailBatchPostRequestWithDefaults

`func NewV1ValidationEmailBatchPostRequestWithDefaults() *V1ValidationEmailBatchPostRequest`

NewV1ValidationEmailBatchPostRequestWithDefaults instantiates a new V1ValidationEmailBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *V1ValidationEmailBatchPostRequest) GetEmails() []interface{}`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *V1ValidationEmailBatchPostRequest) GetEmailsOk() (*[]interface{}, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *V1ValidationEmailBatchPostRequest) SetEmails(v []interface{})`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


