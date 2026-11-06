# V1SystemsBusinessCalendarCountryGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code echoed from the path parameter | [optional] 
**Year** | Pointer to **int32** | Year the calendar covers | [optional] 
**Month** | Pointer to **string** | Month scope. Present only when the month query parameter was provided. | [optional] 
**WorkingDays** | Pointer to **int32** | Number of working days (excluding weekends and public holidays) in the scope | [optional] 
**TotalDays** | Pointer to **string** | Total calendar days in the month. Present only when month is provided. | [optional] 
**WeekendDays** | Pointer to **string** | Number of weekend days in the month. Present only when month is provided. | [optional] 
**Holidays** | Pointer to **[]interface{}** | List of public holidays in the scope. Each item: {date: YYYY-MM-DD, name: string, type: national} | [optional] 
**HolidayCount** | Pointer to **int32** | Count of holidays in the holidays array | [optional] 
**NextHoliday** | Pointer to **string** | Next public holiday within 90 days from today. Shape: {date, name, type} | [optional] 

## Methods

### NewV1SystemsBusinessCalendarCountryGet200ResponseData

`func NewV1SystemsBusinessCalendarCountryGet200ResponseData() *V1SystemsBusinessCalendarCountryGet200ResponseData`

NewV1SystemsBusinessCalendarCountryGet200ResponseData instantiates a new V1SystemsBusinessCalendarCountryGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsBusinessCalendarCountryGet200ResponseDataWithDefaults

`func NewV1SystemsBusinessCalendarCountryGet200ResponseDataWithDefaults() *V1SystemsBusinessCalendarCountryGet200ResponseData`

NewV1SystemsBusinessCalendarCountryGet200ResponseDataWithDefaults instantiates a new V1SystemsBusinessCalendarCountryGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetYear

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetWorkingDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetWorkingDays() int32`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetWorkingDaysOk() (*int32, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetWorkingDays(v int32)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### GetTotalDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetTotalDays() string`

GetTotalDays returns the TotalDays field if non-nil, zero value otherwise.

### GetTotalDaysOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetTotalDaysOk() (*string, bool)`

GetTotalDaysOk returns a tuple with the TotalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetTotalDays(v string)`

SetTotalDays sets TotalDays field to given value.

### HasTotalDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasTotalDays() bool`

HasTotalDays returns a boolean if a field has been set.

### GetWeekendDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetWeekendDays() string`

GetWeekendDays returns the WeekendDays field if non-nil, zero value otherwise.

### GetWeekendDaysOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetWeekendDaysOk() (*string, bool)`

GetWeekendDaysOk returns a tuple with the WeekendDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekendDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetWeekendDays(v string)`

SetWeekendDays sets WeekendDays field to given value.

### HasWeekendDays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasWeekendDays() bool`

HasWeekendDays returns a boolean if a field has been set.

### GetHolidays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetHolidays() []interface{}`

GetHolidays returns the Holidays field if non-nil, zero value otherwise.

### GetHolidaysOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetHolidaysOk() (*[]interface{}, bool)`

GetHolidaysOk returns a tuple with the Holidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetHolidays(v []interface{})`

SetHolidays sets Holidays field to given value.

### HasHolidays

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasHolidays() bool`

HasHolidays returns a boolean if a field has been set.

### GetHolidayCount

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetHolidayCount() int32`

GetHolidayCount returns the HolidayCount field if non-nil, zero value otherwise.

### GetHolidayCountOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetHolidayCountOk() (*int32, bool)`

GetHolidayCountOk returns a tuple with the HolidayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayCount

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetHolidayCount(v int32)`

SetHolidayCount sets HolidayCount field to given value.

### HasHolidayCount

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasHolidayCount() bool`

HasHolidayCount returns a boolean if a field has been set.

### GetNextHoliday

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetNextHoliday() string`

GetNextHoliday returns the NextHoliday field if non-nil, zero value otherwise.

### GetNextHolidayOk

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) GetNextHolidayOk() (*string, bool)`

GetNextHolidayOk returns a tuple with the NextHoliday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHoliday

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) SetNextHoliday(v string)`

SetNextHoliday sets NextHoliday field to given value.

### HasNextHoliday

`func (o *V1SystemsBusinessCalendarCountryGet200ResponseData) HasNextHoliday() bool`

HasNextHoliday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


