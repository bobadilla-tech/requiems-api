# V1NetworkingIpAsnBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **string** | List of IP addresses to look up (1–50 items, each must be a valid IPv4 or IPv6 address). | 

## Methods

### NewV1NetworkingIpAsnBatchPostRequest

`func NewV1NetworkingIpAsnBatchPostRequest(ips string, ) *V1NetworkingIpAsnBatchPostRequest`

NewV1NetworkingIpAsnBatchPostRequest instantiates a new V1NetworkingIpAsnBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingIpAsnBatchPostRequestWithDefaults

`func NewV1NetworkingIpAsnBatchPostRequestWithDefaults() *V1NetworkingIpAsnBatchPostRequest`

NewV1NetworkingIpAsnBatchPostRequestWithDefaults instantiates a new V1NetworkingIpAsnBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *V1NetworkingIpAsnBatchPostRequest) GetIps() string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *V1NetworkingIpAsnBatchPostRequest) GetIpsOk() (*string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *V1NetworkingIpAsnBatchPostRequest) SetIps(v string)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


