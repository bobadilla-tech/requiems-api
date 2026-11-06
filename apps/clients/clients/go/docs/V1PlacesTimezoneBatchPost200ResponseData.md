# V1PlacesTimezoneBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsInfoTimezone** | Pointer to **string** | IANA timezone identifier. | [optional] 
**ResultsInfoOffset** | Pointer to **string** | UTC offset in +HH:MM or -HH:MM format. | [optional] 
**ResultsInfoCurrentTime** | Pointer to **string** | Current time in UTC (RFC 3339). | [optional] 
**ResultsInfoIsDst** | Pointer to **bool** | Whether daylight saving time is currently active. | [optional] 
**ResultsError** | Pointer to **string** | Error message if this item failed (omitted on success). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1PlacesTimezoneBatchPost200ResponseData

`func NewV1PlacesTimezoneBatchPost200ResponseData() *V1PlacesTimezoneBatchPost200ResponseData`

NewV1PlacesTimezoneBatchPost200ResponseData instantiates a new V1PlacesTimezoneBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesTimezoneBatchPost200ResponseDataWithDefaults

`func NewV1PlacesTimezoneBatchPost200ResponseDataWithDefaults() *V1PlacesTimezoneBatchPost200ResponseData`

NewV1PlacesTimezoneBatchPost200ResponseDataWithDefaults instantiates a new V1PlacesTimezoneBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1PlacesTimezoneBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1PlacesTimezoneBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsInfoTimezone

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoTimezone() string`

GetResultsInfoTimezone returns the ResultsInfoTimezone field if non-nil, zero value otherwise.

### GetResultsInfoTimezoneOk

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoTimezoneOk() (*string, bool)`

GetResultsInfoTimezoneOk returns a tuple with the ResultsInfoTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsInfoTimezone

`func (o *V1PlacesTimezoneBatchPost200ResponseData) SetResultsInfoTimezone(v string)`

SetResultsInfoTimezone sets ResultsInfoTimezone field to given value.

### HasResultsInfoTimezone

`func (o *V1PlacesTimezoneBatchPost200ResponseData) HasResultsInfoTimezone() bool`

HasResultsInfoTimezone returns a boolean if a field has been set.

### GetResultsInfoOffset

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoOffset() string`

GetResultsInfoOffset returns the ResultsInfoOffset field if non-nil, zero value otherwise.

### GetResultsInfoOffsetOk

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoOffsetOk() (*string, bool)`

GetResultsInfoOffsetOk returns a tuple with the ResultsInfoOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsInfoOffset

`func (o *V1PlacesTimezoneBatchPost200ResponseData) SetResultsInfoOffset(v string)`

SetResultsInfoOffset sets ResultsInfoOffset field to given value.

### HasResultsInfoOffset

`func (o *V1PlacesTimezoneBatchPost200ResponseData) HasResultsInfoOffset() bool`

HasResultsInfoOffset returns a boolean if a field has been set.

### GetResultsInfoCurrentTime

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoCurrentTime() string`

GetResultsInfoCurrentTime returns the ResultsInfoCurrentTime field if non-nil, zero value otherwise.

### GetResultsInfoCurrentTimeOk

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoCurrentTimeOk() (*string, bool)`

GetResultsInfoCurrentTimeOk returns a tuple with the ResultsInfoCurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsInfoCurrentTime

`func (o *V1PlacesTimezoneBatchPost200ResponseData) SetResultsInfoCurrentTime(v string)`

SetResultsInfoCurrentTime sets ResultsInfoCurrentTime field to given value.

### HasResultsInfoCurrentTime

`func (o *V1PlacesTimezoneBatchPost200ResponseData) HasResultsInfoCurrentTime() bool`

HasResultsInfoCurrentTime returns a boolean if a field has been set.

### GetResultsInfoIsDst

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoIsDst() bool`

GetResultsInfoIsDst returns the ResultsInfoIsDst field if non-nil, zero value otherwise.

### GetResultsInfoIsDstOk

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsInfoIsDstOk() (*bool, bool)`

GetResultsInfoIsDstOk returns a tuple with the ResultsInfoIsDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsInfoIsDst

`func (o *V1PlacesTimezoneBatchPost200ResponseData) SetResultsInfoIsDst(v bool)`

SetResultsInfoIsDst sets ResultsInfoIsDst field to given value.

### HasResultsInfoIsDst

`func (o *V1PlacesTimezoneBatchPost200ResponseData) HasResultsInfoIsDst() bool`

HasResultsInfoIsDst returns a boolean if a field has been set.

### GetResultsError

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1PlacesTimezoneBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1PlacesTimezoneBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetTotal

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1PlacesTimezoneBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1PlacesTimezoneBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1PlacesTimezoneBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


