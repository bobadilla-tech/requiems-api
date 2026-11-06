# V1PlacesTimezoneGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** | IANA timezone identifier (e.g. \&quot;Europe/London\&quot;, \&quot;Asia/Tokyo\&quot;) | [optional] 
**Offset** | Pointer to **string** | UTC offset in +HH:MM or -HH:MM format (e.g. &#39;+05:30&#39;, &#39;-05:00&#39;) | [optional] 
**CurrentTime** | Pointer to **string** | Current time in UTC, formatted as RFC 3339 (e.g. \&quot;2024-12-15T14:30:00Z\&quot;) | [optional] 
**IsDst** | Pointer to **bool** | Whether the location is currently observing daylight saving time | [optional] 

## Methods

### NewV1PlacesTimezoneGet200ResponseData

`func NewV1PlacesTimezoneGet200ResponseData() *V1PlacesTimezoneGet200ResponseData`

NewV1PlacesTimezoneGet200ResponseData instantiates a new V1PlacesTimezoneGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesTimezoneGet200ResponseDataWithDefaults

`func NewV1PlacesTimezoneGet200ResponseDataWithDefaults() *V1PlacesTimezoneGet200ResponseData`

NewV1PlacesTimezoneGet200ResponseDataWithDefaults instantiates a new V1PlacesTimezoneGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *V1PlacesTimezoneGet200ResponseData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V1PlacesTimezoneGet200ResponseData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V1PlacesTimezoneGet200ResponseData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V1PlacesTimezoneGet200ResponseData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetOffset

`func (o *V1PlacesTimezoneGet200ResponseData) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *V1PlacesTimezoneGet200ResponseData) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *V1PlacesTimezoneGet200ResponseData) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *V1PlacesTimezoneGet200ResponseData) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCurrentTime

`func (o *V1PlacesTimezoneGet200ResponseData) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *V1PlacesTimezoneGet200ResponseData) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *V1PlacesTimezoneGet200ResponseData) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *V1PlacesTimezoneGet200ResponseData) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetIsDst

`func (o *V1PlacesTimezoneGet200ResponseData) GetIsDst() bool`

GetIsDst returns the IsDst field if non-nil, zero value otherwise.

### GetIsDstOk

`func (o *V1PlacesTimezoneGet200ResponseData) GetIsDstOk() (*bool, bool)`

GetIsDstOk returns a tuple with the IsDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDst

`func (o *V1PlacesTimezoneGet200ResponseData) SetIsDst(v bool)`

SetIsDst sets IsDst field to given value.

### HasIsDst

`func (o *V1PlacesTimezoneGet200ResponseData) HasIsDst() bool`

HasIsDst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


