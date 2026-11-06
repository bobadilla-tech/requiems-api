# V1NetworkingDomainDomainGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The domain that was looked up | [optional] 
**Available** | Pointer to **bool** | True when the domain appears to be unregistered (NS lookup returns NXDOMAIN). False when name servers are delegated. | [optional] 
**DnsA** | Pointer to **string** | IPv4 addresses (A records) | [optional] 
**DnsAaaa** | Pointer to **string** | IPv6 addresses (AAAA records) | [optional] 
**DnsMx** | Pointer to **string** | Mail exchange records, each with host and priority fields | [optional] 
**DnsNs** | Pointer to **string** | Authoritative name server hostnames | [optional] 
**DnsTxt** | Pointer to **string** | TXT record values (SPF, DKIM, verification tokens, etc.) | [optional] 
**DnsCname** | Pointer to **string** | CNAME alias target, if the domain is an alias. Empty string when no alias exists. | [optional] 

## Methods

### NewV1NetworkingDomainDomainGet200ResponseData

`func NewV1NetworkingDomainDomainGet200ResponseData() *V1NetworkingDomainDomainGet200ResponseData`

NewV1NetworkingDomainDomainGet200ResponseData instantiates a new V1NetworkingDomainDomainGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingDomainDomainGet200ResponseDataWithDefaults

`func NewV1NetworkingDomainDomainGet200ResponseDataWithDefaults() *V1NetworkingDomainDomainGet200ResponseData`

NewV1NetworkingDomainDomainGet200ResponseDataWithDefaults instantiates a new V1NetworkingDomainDomainGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAvailable

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetDnsA

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsA() string`

GetDnsA returns the DnsA field if non-nil, zero value otherwise.

### GetDnsAOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsAOk() (*string, bool)`

GetDnsAOk returns a tuple with the DnsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsA

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetDnsA(v string)`

SetDnsA sets DnsA field to given value.

### HasDnsA

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasDnsA() bool`

HasDnsA returns a boolean if a field has been set.

### GetDnsAaaa

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsAaaa() string`

GetDnsAaaa returns the DnsAaaa field if non-nil, zero value otherwise.

### GetDnsAaaaOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsAaaaOk() (*string, bool)`

GetDnsAaaaOk returns a tuple with the DnsAaaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsAaaa

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetDnsAaaa(v string)`

SetDnsAaaa sets DnsAaaa field to given value.

### HasDnsAaaa

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasDnsAaaa() bool`

HasDnsAaaa returns a boolean if a field has been set.

### GetDnsMx

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsMx() string`

GetDnsMx returns the DnsMx field if non-nil, zero value otherwise.

### GetDnsMxOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsMxOk() (*string, bool)`

GetDnsMxOk returns a tuple with the DnsMx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMx

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetDnsMx(v string)`

SetDnsMx sets DnsMx field to given value.

### HasDnsMx

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasDnsMx() bool`

HasDnsMx returns a boolean if a field has been set.

### GetDnsNs

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsNs() string`

GetDnsNs returns the DnsNs field if non-nil, zero value otherwise.

### GetDnsNsOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsNsOk() (*string, bool)`

GetDnsNsOk returns a tuple with the DnsNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNs

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetDnsNs(v string)`

SetDnsNs sets DnsNs field to given value.

### HasDnsNs

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasDnsNs() bool`

HasDnsNs returns a boolean if a field has been set.

### GetDnsTxt

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsTxt() string`

GetDnsTxt returns the DnsTxt field if non-nil, zero value otherwise.

### GetDnsTxtOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsTxtOk() (*string, bool)`

GetDnsTxtOk returns a tuple with the DnsTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTxt

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetDnsTxt(v string)`

SetDnsTxt sets DnsTxt field to given value.

### HasDnsTxt

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasDnsTxt() bool`

HasDnsTxt returns a boolean if a field has been set.

### GetDnsCname

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsCname() string`

GetDnsCname returns the DnsCname field if non-nil, zero value otherwise.

### GetDnsCnameOk

`func (o *V1NetworkingDomainDomainGet200ResponseData) GetDnsCnameOk() (*string, bool)`

GetDnsCnameOk returns a tuple with the DnsCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCname

`func (o *V1NetworkingDomainDomainGet200ResponseData) SetDnsCname(v string)`

SetDnsCname sets DnsCname field to given value.

### HasDnsCname

`func (o *V1NetworkingDomainDomainGet200ResponseData) HasDnsCname() bool`

HasDnsCname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


