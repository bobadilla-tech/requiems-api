# V1PlacesGeocodeBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsAddress** | Pointer to **string** | The original address as provided in the request. | [optional] 
**ResultsResult** | Pointer to **map[string]interface{}** | Geocoding result (omitted on error). Same fields as the single-item endpoint. | [optional] 
**ResultsError** | Pointer to **string** | Error message if geocoding failed (omitted on success). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1PlacesGeocodeBatchPost200ResponseData

`func NewV1PlacesGeocodeBatchPost200ResponseData() *V1PlacesGeocodeBatchPost200ResponseData`

NewV1PlacesGeocodeBatchPost200ResponseData instantiates a new V1PlacesGeocodeBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesGeocodeBatchPost200ResponseDataWithDefaults

`func NewV1PlacesGeocodeBatchPost200ResponseDataWithDefaults() *V1PlacesGeocodeBatchPost200ResponseData`

NewV1PlacesGeocodeBatchPost200ResponseDataWithDefaults instantiates a new V1PlacesGeocodeBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1PlacesGeocodeBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1PlacesGeocodeBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsAddress

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResultsAddress() string`

GetResultsAddress returns the ResultsAddress field if non-nil, zero value otherwise.

### GetResultsAddressOk

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResultsAddressOk() (*string, bool)`

GetResultsAddressOk returns a tuple with the ResultsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsAddress

`func (o *V1PlacesGeocodeBatchPost200ResponseData) SetResultsAddress(v string)`

SetResultsAddress sets ResultsAddress field to given value.

### HasResultsAddress

`func (o *V1PlacesGeocodeBatchPost200ResponseData) HasResultsAddress() bool`

HasResultsAddress returns a boolean if a field has been set.

### GetResultsResult

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResultsResult() map[string]interface{}`

GetResultsResult returns the ResultsResult field if non-nil, zero value otherwise.

### GetResultsResultOk

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResultsResultOk() (*map[string]interface{}, bool)`

GetResultsResultOk returns a tuple with the ResultsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResult

`func (o *V1PlacesGeocodeBatchPost200ResponseData) SetResultsResult(v map[string]interface{})`

SetResultsResult sets ResultsResult field to given value.

### HasResultsResult

`func (o *V1PlacesGeocodeBatchPost200ResponseData) HasResultsResult() bool`

HasResultsResult returns a boolean if a field has been set.

### GetResultsError

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1PlacesGeocodeBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1PlacesGeocodeBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetTotal

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1PlacesGeocodeBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1PlacesGeocodeBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1PlacesGeocodeBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


