# V1NetworkingIpGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address that was looked up (the requesting client&#39;s IP) | [optional] 
**Country** | Pointer to **string** | Country name where the IP is located | [optional] 
**CountryCode** | Pointer to **string** | Two-letter ISO country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) | [optional] 
**City** | Pointer to **string** | City name where the IP is located | [optional] 
**Isp** | Pointer to **string** | Internet Service Provider providing the IP | [optional] 
**IsVpn** | Pointer to **bool** | True when the IP belongs to a known VPN | [optional] 

## Methods

### NewV1NetworkingIpGet200ResponseData

`func NewV1NetworkingIpGet200ResponseData() *V1NetworkingIpGet200ResponseData`

NewV1NetworkingIpGet200ResponseData instantiates a new V1NetworkingIpGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingIpGet200ResponseDataWithDefaults

`func NewV1NetworkingIpGet200ResponseDataWithDefaults() *V1NetworkingIpGet200ResponseData`

NewV1NetworkingIpGet200ResponseDataWithDefaults instantiates a new V1NetworkingIpGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *V1NetworkingIpGet200ResponseData) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V1NetworkingIpGet200ResponseData) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V1NetworkingIpGet200ResponseData) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V1NetworkingIpGet200ResponseData) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetCountry

`func (o *V1NetworkingIpGet200ResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V1NetworkingIpGet200ResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V1NetworkingIpGet200ResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V1NetworkingIpGet200ResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *V1NetworkingIpGet200ResponseData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V1NetworkingIpGet200ResponseData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V1NetworkingIpGet200ResponseData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *V1NetworkingIpGet200ResponseData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCity

`func (o *V1NetworkingIpGet200ResponseData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *V1NetworkingIpGet200ResponseData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *V1NetworkingIpGet200ResponseData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *V1NetworkingIpGet200ResponseData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetIsp

`func (o *V1NetworkingIpGet200ResponseData) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *V1NetworkingIpGet200ResponseData) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *V1NetworkingIpGet200ResponseData) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *V1NetworkingIpGet200ResponseData) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIsVpn

`func (o *V1NetworkingIpGet200ResponseData) GetIsVpn() bool`

GetIsVpn returns the IsVpn field if non-nil, zero value otherwise.

### GetIsVpnOk

`func (o *V1NetworkingIpGet200ResponseData) GetIsVpnOk() (*bool, bool)`

GetIsVpnOk returns a tuple with the IsVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpn

`func (o *V1NetworkingIpGet200ResponseData) SetIsVpn(v bool)`

SetIsVpn sets IsVpn field to given value.

### HasIsVpn

`func (o *V1NetworkingIpGet200ResponseData) HasIsVpn() bool`

HasIsVpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


