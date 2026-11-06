# V1NetworkingIpAsnBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | Per-item results in input order. | [optional] 
**ResultsIp** | Pointer to **string** | The IP address that was looked up. | [optional] 
**ResultsResult** | Pointer to **map[string]interface{}** | ASN data (omitted on error). Same fields as the single-item endpoint. | [optional] 
**ResultsError** | Pointer to **string** | Error message if the lookup failed (omitted on success). | [optional] 
**Total** | Pointer to **int32** | Total number of items processed. | [optional] 

## Methods

### NewV1NetworkingIpAsnBatchPost200ResponseData

`func NewV1NetworkingIpAsnBatchPost200ResponseData() *V1NetworkingIpAsnBatchPost200ResponseData`

NewV1NetworkingIpAsnBatchPost200ResponseData instantiates a new V1NetworkingIpAsnBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingIpAsnBatchPost200ResponseDataWithDefaults

`func NewV1NetworkingIpAsnBatchPost200ResponseDataWithDefaults() *V1NetworkingIpAsnBatchPost200ResponseData`

NewV1NetworkingIpAsnBatchPost200ResponseDataWithDefaults instantiates a new V1NetworkingIpAsnBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsIp

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResultsIp() string`

GetResultsIp returns the ResultsIp field if non-nil, zero value otherwise.

### GetResultsIpOk

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResultsIpOk() (*string, bool)`

GetResultsIpOk returns a tuple with the ResultsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsIp

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) SetResultsIp(v string)`

SetResultsIp sets ResultsIp field to given value.

### HasResultsIp

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) HasResultsIp() bool`

HasResultsIp returns a boolean if a field has been set.

### GetResultsResult

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResultsResult() map[string]interface{}`

GetResultsResult returns the ResultsResult field if non-nil, zero value otherwise.

### GetResultsResultOk

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResultsResultOk() (*map[string]interface{}, bool)`

GetResultsResultOk returns a tuple with the ResultsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsResult

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) SetResultsResult(v map[string]interface{})`

SetResultsResult sets ResultsResult field to given value.

### HasResultsResult

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) HasResultsResult() bool`

HasResultsResult returns a boolean if a field has been set.

### GetResultsError

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetTotal

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1NetworkingIpAsnBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


