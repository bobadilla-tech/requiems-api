# V1PlacesHolidaysBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | One result per query, in the same order as the request | [optional] 
**ResultsCountry** | Pointer to **string** | ISO 3166-1 alpha-2 country code | [optional] 
**ResultsYear** | Pointer to **int32** | Year queried | [optional] 
**ResultsFound** | Pointer to **bool** | false when no holidays exist for that country/year combination | [optional] 
**ResultsHolidays** | Pointer to **[]interface{}** | List of holidays. Omitted when found is false. | [optional] 
**ResultsTotal** | Pointer to **int32** | Number of holidays. Omitted when found is false. | [optional] 
**Total** | Pointer to **int32** | Total number of results (equals the number of queries sent) | [optional] 

## Methods

### NewV1PlacesHolidaysBatchPost200ResponseData

`func NewV1PlacesHolidaysBatchPost200ResponseData() *V1PlacesHolidaysBatchPost200ResponseData`

NewV1PlacesHolidaysBatchPost200ResponseData instantiates a new V1PlacesHolidaysBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesHolidaysBatchPost200ResponseDataWithDefaults

`func NewV1PlacesHolidaysBatchPost200ResponseDataWithDefaults() *V1PlacesHolidaysBatchPost200ResponseData`

NewV1PlacesHolidaysBatchPost200ResponseDataWithDefaults instantiates a new V1PlacesHolidaysBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1PlacesHolidaysBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1PlacesHolidaysBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsCountry

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsCountry() string`

GetResultsCountry returns the ResultsCountry field if non-nil, zero value otherwise.

### GetResultsCountryOk

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsCountryOk() (*string, bool)`

GetResultsCountryOk returns a tuple with the ResultsCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsCountry

`func (o *V1PlacesHolidaysBatchPost200ResponseData) SetResultsCountry(v string)`

SetResultsCountry sets ResultsCountry field to given value.

### HasResultsCountry

`func (o *V1PlacesHolidaysBatchPost200ResponseData) HasResultsCountry() bool`

HasResultsCountry returns a boolean if a field has been set.

### GetResultsYear

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsYear() int32`

GetResultsYear returns the ResultsYear field if non-nil, zero value otherwise.

### GetResultsYearOk

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsYearOk() (*int32, bool)`

GetResultsYearOk returns a tuple with the ResultsYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsYear

`func (o *V1PlacesHolidaysBatchPost200ResponseData) SetResultsYear(v int32)`

SetResultsYear sets ResultsYear field to given value.

### HasResultsYear

`func (o *V1PlacesHolidaysBatchPost200ResponseData) HasResultsYear() bool`

HasResultsYear returns a boolean if a field has been set.

### GetResultsFound

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsFound() bool`

GetResultsFound returns the ResultsFound field if non-nil, zero value otherwise.

### GetResultsFoundOk

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsFoundOk() (*bool, bool)`

GetResultsFoundOk returns a tuple with the ResultsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFound

`func (o *V1PlacesHolidaysBatchPost200ResponseData) SetResultsFound(v bool)`

SetResultsFound sets ResultsFound field to given value.

### HasResultsFound

`func (o *V1PlacesHolidaysBatchPost200ResponseData) HasResultsFound() bool`

HasResultsFound returns a boolean if a field has been set.

### GetResultsHolidays

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsHolidays() []interface{}`

GetResultsHolidays returns the ResultsHolidays field if non-nil, zero value otherwise.

### GetResultsHolidaysOk

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsHolidaysOk() (*[]interface{}, bool)`

GetResultsHolidaysOk returns a tuple with the ResultsHolidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsHolidays

`func (o *V1PlacesHolidaysBatchPost200ResponseData) SetResultsHolidays(v []interface{})`

SetResultsHolidays sets ResultsHolidays field to given value.

### HasResultsHolidays

`func (o *V1PlacesHolidaysBatchPost200ResponseData) HasResultsHolidays() bool`

HasResultsHolidays returns a boolean if a field has been set.

### GetResultsTotal

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsTotal() int32`

GetResultsTotal returns the ResultsTotal field if non-nil, zero value otherwise.

### GetResultsTotalOk

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetResultsTotalOk() (*int32, bool)`

GetResultsTotalOk returns a tuple with the ResultsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsTotal

`func (o *V1PlacesHolidaysBatchPost200ResponseData) SetResultsTotal(v int32)`

SetResultsTotal sets ResultsTotal field to given value.

### HasResultsTotal

`func (o *V1PlacesHolidaysBatchPost200ResponseData) HasResultsTotal() bool`

HasResultsTotal returns a boolean if a field has been set.

### GetTotal

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1PlacesHolidaysBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1PlacesHolidaysBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1PlacesHolidaysBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


