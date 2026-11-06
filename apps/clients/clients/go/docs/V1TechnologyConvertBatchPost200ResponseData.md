# V1TechnologyConvertBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | List of unit conversion results returned in the same order as the input operations. | [optional] 
**ResultsFrom** | Pointer to **string** | Source unit used for the conversion. | [optional] 
**ResultsTo** | Pointer to **string** | Target unit used for the conversion. | [optional] 
**ResultsSuccess** | Pointer to **bool** | Indicates whether the conversion operation was successful. | [optional] 
**ResultsError** | Pointer to **string** | Error message if the conversion failed. Empty when success is true. | [optional] 
**ResultsData** | Pointer to **map[string]interface{}** | Conversion result details. Contains empty values when the conversion fails. | [optional] 
**ResultsDataFrom** | Pointer to **string** | Source unit. | [optional] 
**ResultsDataTo** | Pointer to **string** | Target unit. | [optional] 
**ResultsDataInput** | Pointer to **float32** | Original numeric value provided for conversion. | [optional] 
**ResultsDataResult** | Pointer to **float32** | Converted numeric result. | [optional] 
**ResultsDataFormula** | Pointer to **string** | Formula or conversion factor used to calculate the result. | [optional] 
**Total** | Pointer to **int32** | Total number of conversion results returned. | [optional] 

## Methods

### NewV1TechnologyConvertBatchPost200ResponseData

`func NewV1TechnologyConvertBatchPost200ResponseData() *V1TechnologyConvertBatchPost200ResponseData`

NewV1TechnologyConvertBatchPost200ResponseData instantiates a new V1TechnologyConvertBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyConvertBatchPost200ResponseDataWithDefaults

`func NewV1TechnologyConvertBatchPost200ResponseDataWithDefaults() *V1TechnologyConvertBatchPost200ResponseData`

NewV1TechnologyConvertBatchPost200ResponseDataWithDefaults instantiates a new V1TechnologyConvertBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsFrom

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsFrom() string`

GetResultsFrom returns the ResultsFrom field if non-nil, zero value otherwise.

### GetResultsFromOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsFromOk() (*string, bool)`

GetResultsFromOk returns a tuple with the ResultsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFrom

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsFrom(v string)`

SetResultsFrom sets ResultsFrom field to given value.

### HasResultsFrom

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsFrom() bool`

HasResultsFrom returns a boolean if a field has been set.

### GetResultsTo

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsTo() string`

GetResultsTo returns the ResultsTo field if non-nil, zero value otherwise.

### GetResultsToOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsToOk() (*string, bool)`

GetResultsToOk returns a tuple with the ResultsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsTo

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsTo(v string)`

SetResultsTo sets ResultsTo field to given value.

### HasResultsTo

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsTo() bool`

HasResultsTo returns a boolean if a field has been set.

### GetResultsSuccess

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsSuccess() bool`

GetResultsSuccess returns the ResultsSuccess field if non-nil, zero value otherwise.

### GetResultsSuccessOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsSuccessOk() (*bool, bool)`

GetResultsSuccessOk returns a tuple with the ResultsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsSuccess

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsSuccess(v bool)`

SetResultsSuccess sets ResultsSuccess field to given value.

### HasResultsSuccess

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsSuccess() bool`

HasResultsSuccess returns a boolean if a field has been set.

### GetResultsError

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetResultsData

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsData() map[string]interface{}`

GetResultsData returns the ResultsData field if non-nil, zero value otherwise.

### GetResultsDataOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataOk() (*map[string]interface{}, bool)`

GetResultsDataOk returns a tuple with the ResultsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsData

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsData(v map[string]interface{})`

SetResultsData sets ResultsData field to given value.

### HasResultsData

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsData() bool`

HasResultsData returns a boolean if a field has been set.

### GetResultsDataFrom

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataFrom() string`

GetResultsDataFrom returns the ResultsDataFrom field if non-nil, zero value otherwise.

### GetResultsDataFromOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataFromOk() (*string, bool)`

GetResultsDataFromOk returns a tuple with the ResultsDataFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataFrom

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsDataFrom(v string)`

SetResultsDataFrom sets ResultsDataFrom field to given value.

### HasResultsDataFrom

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsDataFrom() bool`

HasResultsDataFrom returns a boolean if a field has been set.

### GetResultsDataTo

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataTo() string`

GetResultsDataTo returns the ResultsDataTo field if non-nil, zero value otherwise.

### GetResultsDataToOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataToOk() (*string, bool)`

GetResultsDataToOk returns a tuple with the ResultsDataTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataTo

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsDataTo(v string)`

SetResultsDataTo sets ResultsDataTo field to given value.

### HasResultsDataTo

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsDataTo() bool`

HasResultsDataTo returns a boolean if a field has been set.

### GetResultsDataInput

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataInput() float32`

GetResultsDataInput returns the ResultsDataInput field if non-nil, zero value otherwise.

### GetResultsDataInputOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataInputOk() (*float32, bool)`

GetResultsDataInputOk returns a tuple with the ResultsDataInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataInput

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsDataInput(v float32)`

SetResultsDataInput sets ResultsDataInput field to given value.

### HasResultsDataInput

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsDataInput() bool`

HasResultsDataInput returns a boolean if a field has been set.

### GetResultsDataResult

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataResult() float32`

GetResultsDataResult returns the ResultsDataResult field if non-nil, zero value otherwise.

### GetResultsDataResultOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataResultOk() (*float32, bool)`

GetResultsDataResultOk returns a tuple with the ResultsDataResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataResult

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsDataResult(v float32)`

SetResultsDataResult sets ResultsDataResult field to given value.

### HasResultsDataResult

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsDataResult() bool`

HasResultsDataResult returns a boolean if a field has been set.

### GetResultsDataFormula

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataFormula() string`

GetResultsDataFormula returns the ResultsDataFormula field if non-nil, zero value otherwise.

### GetResultsDataFormulaOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetResultsDataFormulaOk() (*string, bool)`

GetResultsDataFormulaOk returns a tuple with the ResultsDataFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataFormula

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetResultsDataFormula(v string)`

SetResultsDataFormula sets ResultsDataFormula field to given value.

### HasResultsDataFormula

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasResultsDataFormula() bool`

HasResultsDataFormula returns a boolean if a field has been set.

### GetTotal

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1TechnologyConvertBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1TechnologyConvertBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1TechnologyConvertBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


