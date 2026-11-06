# V1TechnologyQrBase64BatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsData** | Pointer to **string** | The input text or URL that was encoded. | [optional] 
**ResultsImage** | Pointer to **string** | Base64-encoded PNG image data (omitted on error). | [optional] 
**ResultsWidth** | Pointer to **int32** | Width of the generated image in pixels (omitted on error). | [optional] 
**ResultsHeight** | Pointer to **int32** | Height of the generated image in pixels (omitted on error). | [optional] 
**ResultsError** | Pointer to **string** | Error message if generation failed (omitted on success). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1TechnologyQrBase64BatchPost200ResponseData

`func NewV1TechnologyQrBase64BatchPost200ResponseData() *V1TechnologyQrBase64BatchPost200ResponseData`

NewV1TechnologyQrBase64BatchPost200ResponseData instantiates a new V1TechnologyQrBase64BatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyQrBase64BatchPost200ResponseDataWithDefaults

`func NewV1TechnologyQrBase64BatchPost200ResponseDataWithDefaults() *V1TechnologyQrBase64BatchPost200ResponseData`

NewV1TechnologyQrBase64BatchPost200ResponseDataWithDefaults instantiates a new V1TechnologyQrBase64BatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsData

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsData() string`

GetResultsData returns the ResultsData field if non-nil, zero value otherwise.

### GetResultsDataOk

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsDataOk() (*string, bool)`

GetResultsDataOk returns a tuple with the ResultsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsData

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) SetResultsData(v string)`

SetResultsData sets ResultsData field to given value.

### HasResultsData

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) HasResultsData() bool`

HasResultsData returns a boolean if a field has been set.

### GetResultsImage

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsImage() string`

GetResultsImage returns the ResultsImage field if non-nil, zero value otherwise.

### GetResultsImageOk

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsImageOk() (*string, bool)`

GetResultsImageOk returns a tuple with the ResultsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsImage

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) SetResultsImage(v string)`

SetResultsImage sets ResultsImage field to given value.

### HasResultsImage

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) HasResultsImage() bool`

HasResultsImage returns a boolean if a field has been set.

### GetResultsWidth

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsWidth() int32`

GetResultsWidth returns the ResultsWidth field if non-nil, zero value otherwise.

### GetResultsWidthOk

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsWidthOk() (*int32, bool)`

GetResultsWidthOk returns a tuple with the ResultsWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsWidth

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) SetResultsWidth(v int32)`

SetResultsWidth sets ResultsWidth field to given value.

### HasResultsWidth

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) HasResultsWidth() bool`

HasResultsWidth returns a boolean if a field has been set.

### GetResultsHeight

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsHeight() int32`

GetResultsHeight returns the ResultsHeight field if non-nil, zero value otherwise.

### GetResultsHeightOk

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsHeightOk() (*int32, bool)`

GetResultsHeightOk returns a tuple with the ResultsHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsHeight

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) SetResultsHeight(v int32)`

SetResultsHeight sets ResultsHeight field to given value.

### HasResultsHeight

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) HasResultsHeight() bool`

HasResultsHeight returns a boolean if a field has been set.

### GetResultsError

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetTotal

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1TechnologyQrBase64BatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


