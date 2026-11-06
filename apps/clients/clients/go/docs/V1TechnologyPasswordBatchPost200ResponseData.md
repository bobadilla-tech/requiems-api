# V1TechnologyPasswordBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsResultPassword** | Pointer to **string** | The generated password (omitted on error). | [optional] 
**ResultsResultLength** | Pointer to **int32** | Length of the generated password. | [optional] 
**ResultsResultStrength** | Pointer to **string** | Strength assessment (weak, medium, or strong). | [optional] 
**ResultsError** | Pointer to **string** | Error message if generation failed (omitted on success). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1TechnologyPasswordBatchPost200ResponseData

`func NewV1TechnologyPasswordBatchPost200ResponseData() *V1TechnologyPasswordBatchPost200ResponseData`

NewV1TechnologyPasswordBatchPost200ResponseData instantiates a new V1TechnologyPasswordBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyPasswordBatchPost200ResponseDataWithDefaults

`func NewV1TechnologyPasswordBatchPost200ResponseDataWithDefaults() *V1TechnologyPasswordBatchPost200ResponseData`

NewV1TechnologyPasswordBatchPost200ResponseDataWithDefaults instantiates a new V1TechnologyPasswordBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1TechnologyPasswordBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1TechnologyPasswordBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsResultPassword

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsResultPassword() string`

GetResultsResultPassword returns the ResultsResultPassword field if non-nil, zero value otherwise.

### GetResultsResultPasswordOk

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsResultPasswordOk() (*string, bool)`

GetResultsResultPasswordOk returns a tuple with the ResultsResultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResultPassword

`func (o *V1TechnologyPasswordBatchPost200ResponseData) SetResultsResultPassword(v string)`

SetResultsResultPassword sets ResultsResultPassword field to given value.

### HasResultsResultPassword

`func (o *V1TechnologyPasswordBatchPost200ResponseData) HasResultsResultPassword() bool`

HasResultsResultPassword returns a boolean if a field has been set.

### GetResultsResultLength

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsResultLength() int32`

GetResultsResultLength returns the ResultsResultLength field if non-nil, zero value otherwise.

### GetResultsResultLengthOk

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsResultLengthOk() (*int32, bool)`

GetResultsResultLengthOk returns a tuple with the ResultsResultLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResultLength

`func (o *V1TechnologyPasswordBatchPost200ResponseData) SetResultsResultLength(v int32)`

SetResultsResultLength sets ResultsResultLength field to given value.

### HasResultsResultLength

`func (o *V1TechnologyPasswordBatchPost200ResponseData) HasResultsResultLength() bool`

HasResultsResultLength returns a boolean if a field has been set.

### GetResultsResultStrength

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsResultStrength() string`

GetResultsResultStrength returns the ResultsResultStrength field if non-nil, zero value otherwise.

### GetResultsResultStrengthOk

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsResultStrengthOk() (*string, bool)`

GetResultsResultStrengthOk returns a tuple with the ResultsResultStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResultStrength

`func (o *V1TechnologyPasswordBatchPost200ResponseData) SetResultsResultStrength(v string)`

SetResultsResultStrength sets ResultsResultStrength field to given value.

### HasResultsResultStrength

`func (o *V1TechnologyPasswordBatchPost200ResponseData) HasResultsResultStrength() bool`

HasResultsResultStrength returns a boolean if a field has been set.

### GetResultsError

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1TechnologyPasswordBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1TechnologyPasswordBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetTotal

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1TechnologyPasswordBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1TechnologyPasswordBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1TechnologyPasswordBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


