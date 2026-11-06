# V1NetworkingWhoisDomainGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The domain name that was looked up | [optional] 
**Registrar** | Pointer to **string** | The name of the registrar holding the domain registration | [optional] 
**NameServers** | Pointer to **string** | List of authoritative name servers for the domain | [optional] 
**Status** | Pointer to **string** | EPP status codes for the domain (e.g. clientTransferProhibited) | [optional] 
**CreatedDate** | Pointer to **string** | Date the domain was first registered (ISO 8601) | [optional] 
**UpdatedDate** | Pointer to **string** | Date the domain record was last updated (ISO 8601) | [optional] 
**ExpiryDate** | Pointer to **string** | Date the domain registration expires (ISO 8601) | [optional] 
**Dnssec** | Pointer to **bool** | True when DNSSEC is enabled for the domain | [optional] 

## Methods

### NewV1NetworkingWhoisDomainGet200ResponseData

`func NewV1NetworkingWhoisDomainGet200ResponseData() *V1NetworkingWhoisDomainGet200ResponseData`

NewV1NetworkingWhoisDomainGet200ResponseData instantiates a new V1NetworkingWhoisDomainGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingWhoisDomainGet200ResponseDataWithDefaults

`func NewV1NetworkingWhoisDomainGet200ResponseDataWithDefaults() *V1NetworkingWhoisDomainGet200ResponseData`

NewV1NetworkingWhoisDomainGet200ResponseDataWithDefaults instantiates a new V1NetworkingWhoisDomainGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRegistrar

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetNameServers

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetNameServers() string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetNameServersOk() (*string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetNameServers(v string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetStatus

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDnssec

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetDnssec() bool`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *V1NetworkingWhoisDomainGet200ResponseData) GetDnssecOk() (*bool, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *V1NetworkingWhoisDomainGet200ResponseData) SetDnssec(v bool)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *V1NetworkingWhoisDomainGet200ResponseData) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


