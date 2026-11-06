# V1SystemsTimezoneFromIpIpGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address as provided | [optional] 
**City** | Pointer to **string** | City resolved from the IP | [optional] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code | [optional] 
**Timezone** | Pointer to **string** | IANA timezone identifier. Null when timezone resolution fails. | [optional] 
**UtcOffset** | Pointer to **string** | UTC offset string. Null when timezone resolution fails. | [optional] 
**DstActive** | Pointer to **string** | True when daylight saving time is currently active. Null when timezone resolution fails. | [optional] 
**CurrentTime** | Pointer to **string** | Current local time in RFC3339 format | [optional] 

## Methods

### NewV1SystemsTimezoneFromIpIpGet200ResponseData

`func NewV1SystemsTimezoneFromIpIpGet200ResponseData() *V1SystemsTimezoneFromIpIpGet200ResponseData`

NewV1SystemsTimezoneFromIpIpGet200ResponseData instantiates a new V1SystemsTimezoneFromIpIpGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsTimezoneFromIpIpGet200ResponseDataWithDefaults

`func NewV1SystemsTimezoneFromIpIpGet200ResponseDataWithDefaults() *V1SystemsTimezoneFromIpIpGet200ResponseData`

NewV1SystemsTimezoneFromIpIpGet200ResponseDataWithDefaults instantiates a new V1SystemsTimezoneFromIpIpGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetCity

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetTimezone

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUtcOffset

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetUtcOffset() string`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetUtcOffsetOk() (*string, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) SetUtcOffset(v string)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.

### GetDstActive

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetDstActive() string`

GetDstActive returns the DstActive field if non-nil, zero value otherwise.

### GetDstActiveOk

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetDstActiveOk() (*string, bool)`

GetDstActiveOk returns a tuple with the DstActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstActive

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) SetDstActive(v string)`

SetDstActive sets DstActive field to given value.

### HasDstActive

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) HasDstActive() bool`

HasDstActive returns a boolean if a field has been set.

### GetCurrentTime

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *V1SystemsTimezoneFromIpIpGet200ResponseData) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


