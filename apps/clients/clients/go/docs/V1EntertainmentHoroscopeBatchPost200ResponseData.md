# V1EntertainmentHoroscopeBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | List of daily horoscopes returned in the same order as the input signs. | [optional] 
**ResultsSign** | Pointer to **string** | Zodiac sign the horoscope corresponds to. | [optional] 
**ResultsDate** | Pointer to **string** | Date the horoscope is for, in YYYY-MM-DD format. | [optional] 
**ResultsHoroscope** | Pointer to **string** | Daily horoscope reading for the given sign. | [optional] 
**ResultsLuckyNumber** | Pointer to **int32** | Lucky number for the given sign on that date. | [optional] 
**ResultsMood** | Pointer to **string** | Predicted mood for the given sign on that date. | [optional] 
**Total** | Pointer to **int32** | Total number of horoscopes returned. | [optional] 

## Methods

### NewV1EntertainmentHoroscopeBatchPost200ResponseData

`func NewV1EntertainmentHoroscopeBatchPost200ResponseData() *V1EntertainmentHoroscopeBatchPost200ResponseData`

NewV1EntertainmentHoroscopeBatchPost200ResponseData instantiates a new V1EntertainmentHoroscopeBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EntertainmentHoroscopeBatchPost200ResponseDataWithDefaults

`func NewV1EntertainmentHoroscopeBatchPost200ResponseDataWithDefaults() *V1EntertainmentHoroscopeBatchPost200ResponseData`

NewV1EntertainmentHoroscopeBatchPost200ResponseDataWithDefaults instantiates a new V1EntertainmentHoroscopeBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsSign

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsSign() string`

GetResultsSign returns the ResultsSign field if non-nil, zero value otherwise.

### GetResultsSignOk

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsSignOk() (*string, bool)`

GetResultsSignOk returns a tuple with the ResultsSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsSign

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) SetResultsSign(v string)`

SetResultsSign sets ResultsSign field to given value.

### HasResultsSign

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) HasResultsSign() bool`

HasResultsSign returns a boolean if a field has been set.

### GetResultsDate

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsDate() string`

GetResultsDate returns the ResultsDate field if non-nil, zero value otherwise.

### GetResultsDateOk

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsDateOk() (*string, bool)`

GetResultsDateOk returns a tuple with the ResultsDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDate

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) SetResultsDate(v string)`

SetResultsDate sets ResultsDate field to given value.

### HasResultsDate

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) HasResultsDate() bool`

HasResultsDate returns a boolean if a field has been set.

### GetResultsHoroscope

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsHoroscope() string`

GetResultsHoroscope returns the ResultsHoroscope field if non-nil, zero value otherwise.

### GetResultsHoroscopeOk

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsHoroscopeOk() (*string, bool)`

GetResultsHoroscopeOk returns a tuple with the ResultsHoroscope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsHoroscope

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) SetResultsHoroscope(v string)`

SetResultsHoroscope sets ResultsHoroscope field to given value.

### HasResultsHoroscope

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) HasResultsHoroscope() bool`

HasResultsHoroscope returns a boolean if a field has been set.

### GetResultsLuckyNumber

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsLuckyNumber() int32`

GetResultsLuckyNumber returns the ResultsLuckyNumber field if non-nil, zero value otherwise.

### GetResultsLuckyNumberOk

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsLuckyNumberOk() (*int32, bool)`

GetResultsLuckyNumberOk returns a tuple with the ResultsLuckyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsLuckyNumber

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) SetResultsLuckyNumber(v int32)`

SetResultsLuckyNumber sets ResultsLuckyNumber field to given value.

### HasResultsLuckyNumber

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) HasResultsLuckyNumber() bool`

HasResultsLuckyNumber returns a boolean if a field has been set.

### GetResultsMood

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsMood() string`

GetResultsMood returns the ResultsMood field if non-nil, zero value otherwise.

### GetResultsMoodOk

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetResultsMoodOk() (*string, bool)`

GetResultsMoodOk returns a tuple with the ResultsMood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsMood

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) SetResultsMood(v string)`

SetResultsMood sets ResultsMood field to given value.

### HasResultsMood

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) HasResultsMood() bool`

HasResultsMood returns a boolean if a field has been set.

### GetTotal

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1EntertainmentHoroscopeBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


