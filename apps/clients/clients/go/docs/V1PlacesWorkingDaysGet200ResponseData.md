# V1PlacesWorkingDaysGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkingDays** | Pointer to **int32** | Number of working days between the two dates (excluding weekends and optionally holidays) | [optional] 
**From** | Pointer to **string** | Start date (echoed from request) | [optional] 
**To** | Pointer to **string** | End date (echoed from request) | [optional] 
**Country** | Pointer to **string** | Country code (echoed from request, empty string if not provided) | [optional] 
**Subdivision** | Pointer to **string** | Subdivision code (echoed from request, empty string if not provided) | [optional] 

## Methods

### NewV1PlacesWorkingDaysGet200ResponseData

`func NewV1PlacesWorkingDaysGet200ResponseData() *V1PlacesWorkingDaysGet200ResponseData`

NewV1PlacesWorkingDaysGet200ResponseData instantiates a new V1PlacesWorkingDaysGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesWorkingDaysGet200ResponseDataWithDefaults

`func NewV1PlacesWorkingDaysGet200ResponseDataWithDefaults() *V1PlacesWorkingDaysGet200ResponseData`

NewV1PlacesWorkingDaysGet200ResponseDataWithDefaults instantiates a new V1PlacesWorkingDaysGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkingDays

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetWorkingDays() int32`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetWorkingDaysOk() (*int32, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *V1PlacesWorkingDaysGet200ResponseData) SetWorkingDays(v int32)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *V1PlacesWorkingDaysGet200ResponseData) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### GetFrom

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1PlacesWorkingDaysGet200ResponseData) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1PlacesWorkingDaysGet200ResponseData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1PlacesWorkingDaysGet200ResponseData) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *V1PlacesWorkingDaysGet200ResponseData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCountry

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V1PlacesWorkingDaysGet200ResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V1PlacesWorkingDaysGet200ResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSubdivision

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetSubdivision() string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *V1PlacesWorkingDaysGet200ResponseData) GetSubdivisionOk() (*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *V1PlacesWorkingDaysGet200ResponseData) SetSubdivision(v string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *V1PlacesWorkingDaysGet200ResponseData) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


