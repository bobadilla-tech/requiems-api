# V1SystemsLocationResolvePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Normalized full address returned by the geocoder | [optional] 
**City** | Pointer to **string** | City name | [optional] 
**Country** | Pointer to **string** | Full country name | [optional] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code | [optional] 
**Coordinates** | Pointer to **map[string]interface{}** | Resolved lat/lng coordinates | [optional] 
**Timezone** | Pointer to **string** | IANA timezone identifier (e.g. Europe/Berlin). Null when timezone resolution fails. | [optional] 
**UtcOffset** | Pointer to **string** | UTC offset string (e.g. +02:00). Null when timezone resolution fails. | [optional] 
**CurrentTime** | Pointer to **string** | Current local time in RFC3339 format. Null when timezone resolution fails. | [optional] 
**IsHolidayToday** | Pointer to **bool** | True when today is a public holiday in the resolved country | [optional] 
**WorkingDaysThisMonth** | Pointer to **int32** | Number of working days (excluding weekends and public holidays) in the current calendar month | [optional] 
**NextHoliday** | Pointer to **string** | Next public holiday within 90 days. Shape: {date: YYYY-MM-DD, name: string, type: national} | [optional] 
**Flags** | Pointer to **[]interface{}** | Non-fatal warnings. Possible values: timezone_unavailable, calendar_unavailable | [optional] 

## Methods

### NewV1SystemsLocationResolvePost200ResponseData

`func NewV1SystemsLocationResolvePost200ResponseData() *V1SystemsLocationResolvePost200ResponseData`

NewV1SystemsLocationResolvePost200ResponseData instantiates a new V1SystemsLocationResolvePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsLocationResolvePost200ResponseDataWithDefaults

`func NewV1SystemsLocationResolvePost200ResponseDataWithDefaults() *V1SystemsLocationResolvePost200ResponseData`

NewV1SystemsLocationResolvePost200ResponseDataWithDefaults instantiates a new V1SystemsLocationResolvePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1SystemsLocationResolvePost200ResponseData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1SystemsLocationResolvePost200ResponseData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1SystemsLocationResolvePost200ResponseData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *V1SystemsLocationResolvePost200ResponseData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *V1SystemsLocationResolvePost200ResponseData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V1SystemsLocationResolvePost200ResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V1SystemsLocationResolvePost200ResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V1SystemsLocationResolvePost200ResponseData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *V1SystemsLocationResolvePost200ResponseData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCoordinates

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCoordinates() map[string]interface{}`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCoordinatesOk() (*map[string]interface{}, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *V1SystemsLocationResolvePost200ResponseData) SetCoordinates(v map[string]interface{})`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *V1SystemsLocationResolvePost200ResponseData) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetTimezone

`func (o *V1SystemsLocationResolvePost200ResponseData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V1SystemsLocationResolvePost200ResponseData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V1SystemsLocationResolvePost200ResponseData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUtcOffset

`func (o *V1SystemsLocationResolvePost200ResponseData) GetUtcOffset() string`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetUtcOffsetOk() (*string, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *V1SystemsLocationResolvePost200ResponseData) SetUtcOffset(v string)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *V1SystemsLocationResolvePost200ResponseData) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.

### GetCurrentTime

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *V1SystemsLocationResolvePost200ResponseData) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *V1SystemsLocationResolvePost200ResponseData) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetIsHolidayToday

`func (o *V1SystemsLocationResolvePost200ResponseData) GetIsHolidayToday() bool`

GetIsHolidayToday returns the IsHolidayToday field if non-nil, zero value otherwise.

### GetIsHolidayTodayOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetIsHolidayTodayOk() (*bool, bool)`

GetIsHolidayTodayOk returns a tuple with the IsHolidayToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHolidayToday

`func (o *V1SystemsLocationResolvePost200ResponseData) SetIsHolidayToday(v bool)`

SetIsHolidayToday sets IsHolidayToday field to given value.

### HasIsHolidayToday

`func (o *V1SystemsLocationResolvePost200ResponseData) HasIsHolidayToday() bool`

HasIsHolidayToday returns a boolean if a field has been set.

### GetWorkingDaysThisMonth

`func (o *V1SystemsLocationResolvePost200ResponseData) GetWorkingDaysThisMonth() int32`

GetWorkingDaysThisMonth returns the WorkingDaysThisMonth field if non-nil, zero value otherwise.

### GetWorkingDaysThisMonthOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetWorkingDaysThisMonthOk() (*int32, bool)`

GetWorkingDaysThisMonthOk returns a tuple with the WorkingDaysThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDaysThisMonth

`func (o *V1SystemsLocationResolvePost200ResponseData) SetWorkingDaysThisMonth(v int32)`

SetWorkingDaysThisMonth sets WorkingDaysThisMonth field to given value.

### HasWorkingDaysThisMonth

`func (o *V1SystemsLocationResolvePost200ResponseData) HasWorkingDaysThisMonth() bool`

HasWorkingDaysThisMonth returns a boolean if a field has been set.

### GetNextHoliday

`func (o *V1SystemsLocationResolvePost200ResponseData) GetNextHoliday() string`

GetNextHoliday returns the NextHoliday field if non-nil, zero value otherwise.

### GetNextHolidayOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetNextHolidayOk() (*string, bool)`

GetNextHolidayOk returns a tuple with the NextHoliday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHoliday

`func (o *V1SystemsLocationResolvePost200ResponseData) SetNextHoliday(v string)`

SetNextHoliday sets NextHoliday field to given value.

### HasNextHoliday

`func (o *V1SystemsLocationResolvePost200ResponseData) HasNextHoliday() bool`

HasNextHoliday returns a boolean if a field has been set.

### GetFlags

`func (o *V1SystemsLocationResolvePost200ResponseData) GetFlags() []interface{}`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V1SystemsLocationResolvePost200ResponseData) GetFlagsOk() (*[]interface{}, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V1SystemsLocationResolvePost200ResponseData) SetFlags(v []interface{})`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V1SystemsLocationResolvePost200ResponseData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


