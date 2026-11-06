# V1PlacesReverseGeocodeBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsLat** | Pointer to **float32** | Latitude as provided in the request. | [optional] 
**ResultsLon** | Pointer to **float32** | Longitude as provided in the request. | [optional] 
**ResultsResult** | Pointer to **map[string]interface{}** | Reverse geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] 
**ResultsError** | Pointer to **string** | Error message if reverse geocoding failed (omitted on success). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1PlacesReverseGeocodeBatchPost200ResponseData

`func NewV1PlacesReverseGeocodeBatchPost200ResponseData() *V1PlacesReverseGeocodeBatchPost200ResponseData`

NewV1PlacesReverseGeocodeBatchPost200ResponseData instantiates a new V1PlacesReverseGeocodeBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesReverseGeocodeBatchPost200ResponseDataWithDefaults

`func NewV1PlacesReverseGeocodeBatchPost200ResponseDataWithDefaults() *V1PlacesReverseGeocodeBatchPost200ResponseData`

NewV1PlacesReverseGeocodeBatchPost200ResponseDataWithDefaults instantiates a new V1PlacesReverseGeocodeBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsLat

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsLat() float32`

GetResultsLat returns the ResultsLat field if non-nil, zero value otherwise.

### GetResultsLatOk

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsLatOk() (*float32, bool)`

GetResultsLatOk returns a tuple with the ResultsLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsLat

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) SetResultsLat(v float32)`

SetResultsLat sets ResultsLat field to given value.

### HasResultsLat

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) HasResultsLat() bool`

HasResultsLat returns a boolean if a field has been set.

### GetResultsLon

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsLon() float32`

GetResultsLon returns the ResultsLon field if non-nil, zero value otherwise.

### GetResultsLonOk

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsLonOk() (*float32, bool)`

GetResultsLonOk returns a tuple with the ResultsLon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsLon

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) SetResultsLon(v float32)`

SetResultsLon sets ResultsLon field to given value.

### HasResultsLon

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) HasResultsLon() bool`

HasResultsLon returns a boolean if a field has been set.

### GetResultsResult

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsResult() map[string]interface{}`

GetResultsResult returns the ResultsResult field if non-nil, zero value otherwise.

### GetResultsResultOk

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsResultOk() (*map[string]interface{}, bool)`

GetResultsResultOk returns a tuple with the ResultsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResult

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) SetResultsResult(v map[string]interface{})`

SetResultsResult sets ResultsResult field to given value.

### HasResultsResult

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) HasResultsResult() bool`

HasResultsResult returns a boolean if a field has been set.

### GetResultsError

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetTotal

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1PlacesReverseGeocodeBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


