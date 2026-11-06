# V1SystemsUserVerifyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address to verify. Domain-level WHOIS, MX, and DNS checks are run automatically. | 
**IpAddress** | Pointer to **string** | Optional IPv4 or IPv6 address. When provided, VPN and proxy signals are added to the result. | [optional] 

## Methods

### NewV1SystemsUserVerifyPostRequest

`func NewV1SystemsUserVerifyPostRequest(email string, ) *V1SystemsUserVerifyPostRequest`

NewV1SystemsUserVerifyPostRequest instantiates a new V1SystemsUserVerifyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsUserVerifyPostRequestWithDefaults

`func NewV1SystemsUserVerifyPostRequestWithDefaults() *V1SystemsUserVerifyPostRequest`

NewV1SystemsUserVerifyPostRequestWithDefaults instantiates a new V1SystemsUserVerifyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *V1SystemsUserVerifyPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1SystemsUserVerifyPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1SystemsUserVerifyPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIpAddress

`func (o *V1SystemsUserVerifyPostRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *V1SystemsUserVerifyPostRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *V1SystemsUserVerifyPostRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *V1SystemsUserVerifyPostRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


