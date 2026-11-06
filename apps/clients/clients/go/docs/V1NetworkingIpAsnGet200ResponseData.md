# V1NetworkingIpAsnGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address that was looked up (the requesting client&#39;s IP) | [optional] 
**Asn** | Pointer to **string** | Autonomous System Number in format \&quot;ASxxxx\&quot; (e.g., \&quot;AS15169\&quot;) | [optional] 
**Org** | Pointer to **string** | Organization name owning the IP address range | [optional] 
**Isp** | Pointer to **string** | Internet Service Provider providing the IP | [optional] 
**Domain** | Pointer to **string** | Domain name associated with the IP or IP range | [optional] 
**Route** | Pointer to **string** | CIDR notation of the network route (e.g., \&quot;8.8.8.0/24\&quot;) | [optional] 
**Type** | Pointer to **string** | Type of network (e.g., \&quot;hosting\&quot;, \&quot;isp\&quot;, \&quot;business\&quot;, \&quot;cdn\&quot;) | [optional] 

## Methods

### NewV1NetworkingIpAsnGet200ResponseData

`func NewV1NetworkingIpAsnGet200ResponseData() *V1NetworkingIpAsnGet200ResponseData`

NewV1NetworkingIpAsnGet200ResponseData instantiates a new V1NetworkingIpAsnGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingIpAsnGet200ResponseDataWithDefaults

`func NewV1NetworkingIpAsnGet200ResponseDataWithDefaults() *V1NetworkingIpAsnGet200ResponseData`

NewV1NetworkingIpAsnGet200ResponseDataWithDefaults instantiates a new V1NetworkingIpAsnGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *V1NetworkingIpAsnGet200ResponseData) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V1NetworkingIpAsnGet200ResponseData) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V1NetworkingIpAsnGet200ResponseData) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V1NetworkingIpAsnGet200ResponseData) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAsn

`func (o *V1NetworkingIpAsnGet200ResponseData) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *V1NetworkingIpAsnGet200ResponseData) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *V1NetworkingIpAsnGet200ResponseData) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *V1NetworkingIpAsnGet200ResponseData) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetOrg

`func (o *V1NetworkingIpAsnGet200ResponseData) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *V1NetworkingIpAsnGet200ResponseData) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *V1NetworkingIpAsnGet200ResponseData) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *V1NetworkingIpAsnGet200ResponseData) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetIsp

`func (o *V1NetworkingIpAsnGet200ResponseData) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *V1NetworkingIpAsnGet200ResponseData) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *V1NetworkingIpAsnGet200ResponseData) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *V1NetworkingIpAsnGet200ResponseData) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetDomain

`func (o *V1NetworkingIpAsnGet200ResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V1NetworkingIpAsnGet200ResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V1NetworkingIpAsnGet200ResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V1NetworkingIpAsnGet200ResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRoute

`func (o *V1NetworkingIpAsnGet200ResponseData) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *V1NetworkingIpAsnGet200ResponseData) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *V1NetworkingIpAsnGet200ResponseData) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *V1NetworkingIpAsnGet200ResponseData) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetType

`func (o *V1NetworkingIpAsnGet200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1NetworkingIpAsnGet200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1NetworkingIpAsnGet200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1NetworkingIpAsnGet200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


