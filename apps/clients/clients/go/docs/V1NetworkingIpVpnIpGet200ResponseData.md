# V1NetworkingIpVpnIpGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The analyzed IP address | [optional] 
**IsVpn** | Pointer to **bool** | True when the IP belongs to a known VPN provider | [optional] 
**IsProxy** | Pointer to **bool** | True when the IP is a known public or web proxy | [optional] 
**IsTor** | Pointer to **bool** | True when the IP is a known Tor exit node | [optional] 
**IsHosting** | Pointer to **bool** | True when the IP belongs to a data-centre or hosting provider (DCH) | [optional] 
**Score** | Pointer to **int32** | Raw threat score (0-9+). Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1 | [optional] 
**Threat** | Pointer to **int32** | Threat level derived from score: 0 &#x3D; None, 1 &#x3D; Low, 2–3 &#x3D; Medium, 4–5 &#x3D; High, 6+ &#x3D; Critical | [optional] 
**FraudScore** | Pointer to **int32** | Fraud risk score from 0 (no risk) to 100 (high risk). Available when using IP2Proxy PX5 or higher | [optional] 
**AsnOrg** | Pointer to **string** | Organization name owning the Autonomous System containing the IP (e.g. \&quot;DIGITALOCEAN-ASN\&quot;) | [optional] 

## Methods

### NewV1NetworkingIpVpnIpGet200ResponseData

`func NewV1NetworkingIpVpnIpGet200ResponseData() *V1NetworkingIpVpnIpGet200ResponseData`

NewV1NetworkingIpVpnIpGet200ResponseData instantiates a new V1NetworkingIpVpnIpGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingIpVpnIpGet200ResponseDataWithDefaults

`func NewV1NetworkingIpVpnIpGet200ResponseDataWithDefaults() *V1NetworkingIpVpnIpGet200ResponseData`

NewV1NetworkingIpVpnIpGet200ResponseDataWithDefaults instantiates a new V1NetworkingIpVpnIpGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsVpn

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsVpn() bool`

GetIsVpn returns the IsVpn field if non-nil, zero value otherwise.

### GetIsVpnOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsVpnOk() (*bool, bool)`

GetIsVpnOk returns a tuple with the IsVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpn

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetIsVpn(v bool)`

SetIsVpn sets IsVpn field to given value.

### HasIsVpn

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasIsVpn() bool`

HasIsVpn returns a boolean if a field has been set.

### GetIsProxy

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetIsTor

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsTor() bool`

GetIsTor returns the IsTor field if non-nil, zero value otherwise.

### GetIsTorOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsTorOk() (*bool, bool)`

GetIsTorOk returns a tuple with the IsTor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTor

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetIsTor(v bool)`

SetIsTor sets IsTor field to given value.

### HasIsTor

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasIsTor() bool`

HasIsTor returns a boolean if a field has been set.

### GetIsHosting

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsHosting() bool`

GetIsHosting returns the IsHosting field if non-nil, zero value otherwise.

### GetIsHostingOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetIsHostingOk() (*bool, bool)`

GetIsHostingOk returns a tuple with the IsHosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHosting

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetIsHosting(v bool)`

SetIsHosting sets IsHosting field to given value.

### HasIsHosting

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasIsHosting() bool`

HasIsHosting returns a boolean if a field has been set.

### GetScore

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetThreat

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetThreat() int32`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetThreatOk() (*int32, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetThreat(v int32)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetFraudScore

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetFraudScore() int32`

GetFraudScore returns the FraudScore field if non-nil, zero value otherwise.

### GetFraudScoreOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetFraudScoreOk() (*int32, bool)`

GetFraudScoreOk returns a tuple with the FraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudScore

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetFraudScore(v int32)`

SetFraudScore sets FraudScore field to given value.

### HasFraudScore

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasFraudScore() bool`

HasFraudScore returns a boolean if a field has been set.

### GetAsnOrg

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetAsnOrg() string`

GetAsnOrg returns the AsnOrg field if non-nil, zero value otherwise.

### GetAsnOrgOk

`func (o *V1NetworkingIpVpnIpGet200ResponseData) GetAsnOrgOk() (*string, bool)`

GetAsnOrgOk returns a tuple with the AsnOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnOrg

`func (o *V1NetworkingIpVpnIpGet200ResponseData) SetAsnOrg(v string)`

SetAsnOrg sets AsnOrg field to given value.

### HasAsnOrg

`func (o *V1NetworkingIpVpnIpGet200ResponseData) HasAsnOrg() bool`

HasAsnOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


