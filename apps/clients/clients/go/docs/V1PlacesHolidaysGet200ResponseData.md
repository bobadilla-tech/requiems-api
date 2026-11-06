# V1PlacesHolidaysGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country code | [optional] 
**Year** | Pointer to **int32** | Year for which holidays are returned | [optional] 
**Holidays** | Pointer to **[]interface{}** | Array of holiday objects | [optional] 
**HolidaysDate** | Pointer to **string** | Holiday date in YYYY-MM-DD format | [optional] 
**HolidaysName** | Pointer to **string** | Name of the holiday | [optional] 
**Total** | Pointer to **int32** | Total number of holidays for the country/year | [optional] 

## Methods

### NewV1PlacesHolidaysGet200ResponseData

`func NewV1PlacesHolidaysGet200ResponseData() *V1PlacesHolidaysGet200ResponseData`

NewV1PlacesHolidaysGet200ResponseData instantiates a new V1PlacesHolidaysGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesHolidaysGet200ResponseDataWithDefaults

`func NewV1PlacesHolidaysGet200ResponseDataWithDefaults() *V1PlacesHolidaysGet200ResponseData`

NewV1PlacesHolidaysGet200ResponseDataWithDefaults instantiates a new V1PlacesHolidaysGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *V1PlacesHolidaysGet200ResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V1PlacesHolidaysGet200ResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V1PlacesHolidaysGet200ResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V1PlacesHolidaysGet200ResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetYear

`func (o *V1PlacesHolidaysGet200ResponseData) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *V1PlacesHolidaysGet200ResponseData) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *V1PlacesHolidaysGet200ResponseData) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *V1PlacesHolidaysGet200ResponseData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetHolidays

`func (o *V1PlacesHolidaysGet200ResponseData) GetHolidays() []interface{}`

GetHolidays returns the Holidays field if non-nil, zero value otherwise.

### GetHolidaysOk

`func (o *V1PlacesHolidaysGet200ResponseData) GetHolidaysOk() (*[]interface{}, bool)`

GetHolidaysOk returns a tuple with the Holidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidays

`func (o *V1PlacesHolidaysGet200ResponseData) SetHolidays(v []interface{})`

SetHolidays sets Holidays field to given value.

### HasHolidays

`func (o *V1PlacesHolidaysGet200ResponseData) HasHolidays() bool`

HasHolidays returns a boolean if a field has been set.

### GetHolidaysDate

`func (o *V1PlacesHolidaysGet200ResponseData) GetHolidaysDate() string`

GetHolidaysDate returns the HolidaysDate field if non-nil, zero value otherwise.

### GetHolidaysDateOk

`func (o *V1PlacesHolidaysGet200ResponseData) GetHolidaysDateOk() (*string, bool)`

GetHolidaysDateOk returns a tuple with the HolidaysDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidaysDate

`func (o *V1PlacesHolidaysGet200ResponseData) SetHolidaysDate(v string)`

SetHolidaysDate sets HolidaysDate field to given value.

### HasHolidaysDate

`func (o *V1PlacesHolidaysGet200ResponseData) HasHolidaysDate() bool`

HasHolidaysDate returns a boolean if a field has been set.

### GetHolidaysName

`func (o *V1PlacesHolidaysGet200ResponseData) GetHolidaysName() string`

GetHolidaysName returns the HolidaysName field if non-nil, zero value otherwise.

### GetHolidaysNameOk

`func (o *V1PlacesHolidaysGet200ResponseData) GetHolidaysNameOk() (*string, bool)`

GetHolidaysNameOk returns a tuple with the HolidaysName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidaysName

`func (o *V1PlacesHolidaysGet200ResponseData) SetHolidaysName(v string)`

SetHolidaysName sets HolidaysName field to given value.

### HasHolidaysName

`func (o *V1PlacesHolidaysGet200ResponseData) HasHolidaysName() bool`

HasHolidaysName returns a boolean if a field has been set.

### GetTotal

`func (o *V1PlacesHolidaysGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1PlacesHolidaysGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1PlacesHolidaysGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1PlacesHolidaysGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


