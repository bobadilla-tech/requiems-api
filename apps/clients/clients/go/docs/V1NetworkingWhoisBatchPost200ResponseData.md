# V1NetworkingWhoisBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | One entry per domain, in the same order as the input | [optional] 
**ResultsDomain** | Pointer to **string** | Domain name requested in the batch | [optional] 
**ResultsFound** | Pointer to **bool** | False when WHOIS data could not be found | [optional] 
**ResultsError** | Pointer to **string** | Error message when found is false | [optional] 
**ResultsData** | Pointer to **map[string]interface{}** | WHOIS information for the domain when found is true | [optional] 
**Total** | Pointer to **int32** | Total number of results returned | [optional] 

## Methods

### NewV1NetworkingWhoisBatchPost200ResponseData

`func NewV1NetworkingWhoisBatchPost200ResponseData() *V1NetworkingWhoisBatchPost200ResponseData`

NewV1NetworkingWhoisBatchPost200ResponseData instantiates a new V1NetworkingWhoisBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingWhoisBatchPost200ResponseDataWithDefaults

`func NewV1NetworkingWhoisBatchPost200ResponseDataWithDefaults() *V1NetworkingWhoisBatchPost200ResponseData`

NewV1NetworkingWhoisBatchPost200ResponseDataWithDefaults instantiates a new V1NetworkingWhoisBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1NetworkingWhoisBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1NetworkingWhoisBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsDomain

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsDomain() string`

GetResultsDomain returns the ResultsDomain field if non-nil, zero value otherwise.

### GetResultsDomainOk

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsDomainOk() (*string, bool)`

GetResultsDomainOk returns a tuple with the ResultsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDomain

`func (o *V1NetworkingWhoisBatchPost200ResponseData) SetResultsDomain(v string)`

SetResultsDomain sets ResultsDomain field to given value.

### HasResultsDomain

`func (o *V1NetworkingWhoisBatchPost200ResponseData) HasResultsDomain() bool`

HasResultsDomain returns a boolean if a field has been set.

### GetResultsFound

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsFound() bool`

GetResultsFound returns the ResultsFound field if non-nil, zero value otherwise.

### GetResultsFoundOk

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsFoundOk() (*bool, bool)`

GetResultsFoundOk returns a tuple with the ResultsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFound

`func (o *V1NetworkingWhoisBatchPost200ResponseData) SetResultsFound(v bool)`

SetResultsFound sets ResultsFound field to given value.

### HasResultsFound

`func (o *V1NetworkingWhoisBatchPost200ResponseData) HasResultsFound() bool`

HasResultsFound returns a boolean if a field has been set.

### GetResultsError

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsError() string`

GetResultsError returns the ResultsError field if non-nil, zero value otherwise.

### GetResultsErrorOk

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsErrorOk() (*string, bool)`

GetResultsErrorOk returns a tuple with the ResultsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsError

`func (o *V1NetworkingWhoisBatchPost200ResponseData) SetResultsError(v string)`

SetResultsError sets ResultsError field to given value.

### HasResultsError

`func (o *V1NetworkingWhoisBatchPost200ResponseData) HasResultsError() bool`

HasResultsError returns a boolean if a field has been set.

### GetResultsData

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsData() map[string]interface{}`

GetResultsData returns the ResultsData field if non-nil, zero value otherwise.

### GetResultsDataOk

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetResultsDataOk() (*map[string]interface{}, bool)`

GetResultsDataOk returns a tuple with the ResultsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsData

`func (o *V1NetworkingWhoisBatchPost200ResponseData) SetResultsData(v map[string]interface{})`

SetResultsData sets ResultsData field to given value.

### HasResultsData

`func (o *V1NetworkingWhoisBatchPost200ResponseData) HasResultsData() bool`

HasResultsData returns a boolean if a field has been set.

### GetTotal

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1NetworkingWhoisBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1NetworkingWhoisBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1NetworkingWhoisBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


