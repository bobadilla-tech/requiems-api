# V1PlacesPostalBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsCode** | Pointer to **string** | The postal code as provided in the request. | [optional] 
**ResultsCountry** | Pointer to **string** | The country code used for the lookup (defaulted to US when omitted). | [optional] 
**ResultsFound** | Pointer to **bool** | Whether the postal code was found in the dataset. | [optional] 
**ResultsResult** | Pointer to **map[string]interface{}** | Postal code details (omitted when found is false). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1PlacesPostalBatchPost200ResponseData

`func NewV1PlacesPostalBatchPost200ResponseData() *V1PlacesPostalBatchPost200ResponseData`

NewV1PlacesPostalBatchPost200ResponseData instantiates a new V1PlacesPostalBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesPostalBatchPost200ResponseDataWithDefaults

`func NewV1PlacesPostalBatchPost200ResponseDataWithDefaults() *V1PlacesPostalBatchPost200ResponseData`

NewV1PlacesPostalBatchPost200ResponseDataWithDefaults instantiates a new V1PlacesPostalBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1PlacesPostalBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1PlacesPostalBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsCode

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsCode() string`

GetResultsCode returns the ResultsCode field if non-nil, zero value otherwise.

### GetResultsCodeOk

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsCodeOk() (*string, bool)`

GetResultsCodeOk returns a tuple with the ResultsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsCode

`func (o *V1PlacesPostalBatchPost200ResponseData) SetResultsCode(v string)`

SetResultsCode sets ResultsCode field to given value.

### HasResultsCode

`func (o *V1PlacesPostalBatchPost200ResponseData) HasResultsCode() bool`

HasResultsCode returns a boolean if a field has been set.

### GetResultsCountry

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsCountry() string`

GetResultsCountry returns the ResultsCountry field if non-nil, zero value otherwise.

### GetResultsCountryOk

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsCountryOk() (*string, bool)`

GetResultsCountryOk returns a tuple with the ResultsCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsCountry

`func (o *V1PlacesPostalBatchPost200ResponseData) SetResultsCountry(v string)`

SetResultsCountry sets ResultsCountry field to given value.

### HasResultsCountry

`func (o *V1PlacesPostalBatchPost200ResponseData) HasResultsCountry() bool`

HasResultsCountry returns a boolean if a field has been set.

### GetResultsFound

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsFound() bool`

GetResultsFound returns the ResultsFound field if non-nil, zero value otherwise.

### GetResultsFoundOk

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsFoundOk() (*bool, bool)`

GetResultsFoundOk returns a tuple with the ResultsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFound

`func (o *V1PlacesPostalBatchPost200ResponseData) SetResultsFound(v bool)`

SetResultsFound sets ResultsFound field to given value.

### HasResultsFound

`func (o *V1PlacesPostalBatchPost200ResponseData) HasResultsFound() bool`

HasResultsFound returns a boolean if a field has been set.

### GetResultsResult

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsResult() map[string]interface{}`

GetResultsResult returns the ResultsResult field if non-nil, zero value otherwise.

### GetResultsResultOk

`func (o *V1PlacesPostalBatchPost200ResponseData) GetResultsResultOk() (*map[string]interface{}, bool)`

GetResultsResultOk returns a tuple with the ResultsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResult

`func (o *V1PlacesPostalBatchPost200ResponseData) SetResultsResult(v map[string]interface{})`

SetResultsResult sets ResultsResult field to given value.

### HasResultsResult

`func (o *V1PlacesPostalBatchPost200ResponseData) HasResultsResult() bool`

HasResultsResult returns a boolean if a field has been set.

### GetTotal

`func (o *V1PlacesPostalBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1PlacesPostalBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1PlacesPostalBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1PlacesPostalBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


