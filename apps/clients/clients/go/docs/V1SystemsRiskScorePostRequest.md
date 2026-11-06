# V1SystemsRiskScorePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address to include in the risk score. | [optional] 
**Phone** | Pointer to **string** | Phone number in E.164 format. | [optional] 
**IpAddress** | Pointer to **string** | IPv4 or IPv6 address to check. | [optional] 

## Methods

### NewV1SystemsRiskScorePostRequest

`func NewV1SystemsRiskScorePostRequest() *V1SystemsRiskScorePostRequest`

NewV1SystemsRiskScorePostRequest instantiates a new V1SystemsRiskScorePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsRiskScorePostRequestWithDefaults

`func NewV1SystemsRiskScorePostRequestWithDefaults() *V1SystemsRiskScorePostRequest`

NewV1SystemsRiskScorePostRequestWithDefaults instantiates a new V1SystemsRiskScorePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *V1SystemsRiskScorePostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1SystemsRiskScorePostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1SystemsRiskScorePostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *V1SystemsRiskScorePostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *V1SystemsRiskScorePostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *V1SystemsRiskScorePostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *V1SystemsRiskScorePostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *V1SystemsRiskScorePostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetIpAddress

`func (o *V1SystemsRiskScorePostRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *V1SystemsRiskScorePostRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *V1SystemsRiskScorePostRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *V1SystemsRiskScorePostRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


