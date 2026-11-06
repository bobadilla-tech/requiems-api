# V1SystemsSignupProtectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address to validate and score. At least one of email, phone, or ip_address is required. | [optional] 
**Phone** | Pointer to **string** | Phone number in E.164 format to validate and check for VoIP or virtual numbers. | [optional] 
**IpAddress** | Pointer to **string** | IPv4 or IPv6 address to check for VPN, proxy, TOR, and hosting status. | [optional] 

## Methods

### NewV1SystemsSignupProtectPostRequest

`func NewV1SystemsSignupProtectPostRequest() *V1SystemsSignupProtectPostRequest`

NewV1SystemsSignupProtectPostRequest instantiates a new V1SystemsSignupProtectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsSignupProtectPostRequestWithDefaults

`func NewV1SystemsSignupProtectPostRequestWithDefaults() *V1SystemsSignupProtectPostRequest`

NewV1SystemsSignupProtectPostRequestWithDefaults instantiates a new V1SystemsSignupProtectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *V1SystemsSignupProtectPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1SystemsSignupProtectPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1SystemsSignupProtectPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *V1SystemsSignupProtectPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *V1SystemsSignupProtectPostRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *V1SystemsSignupProtectPostRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *V1SystemsSignupProtectPostRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *V1SystemsSignupProtectPostRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetIpAddress

`func (o *V1SystemsSignupProtectPostRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *V1SystemsSignupProtectPostRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *V1SystemsSignupProtectPostRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *V1SystemsSignupProtectPostRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


