# V1TextSpellcheckBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **string** | One entry per input text, in the same order as the input array. Each item contains corrected (the fixed text) and corrections (list of individual corrections with original, suggested, suggestions, and position). | [optional] 
**Total** | Pointer to **int32** | Number of texts processed. Equals the length of the input array. | [optional] 

## Methods

### NewV1TextSpellcheckBatchPost200ResponseData

`func NewV1TextSpellcheckBatchPost200ResponseData() *V1TextSpellcheckBatchPost200ResponseData`

NewV1TextSpellcheckBatchPost200ResponseData instantiates a new V1TextSpellcheckBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TextSpellcheckBatchPost200ResponseDataWithDefaults

`func NewV1TextSpellcheckBatchPost200ResponseDataWithDefaults() *V1TextSpellcheckBatchPost200ResponseData`

NewV1TextSpellcheckBatchPost200ResponseDataWithDefaults instantiates a new V1TextSpellcheckBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1TextSpellcheckBatchPost200ResponseData) GetResults() string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1TextSpellcheckBatchPost200ResponseData) GetResultsOk() (*string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1TextSpellcheckBatchPost200ResponseData) SetResults(v string)`

SetResults sets Results field to given value.

### HasResults

`func (o *V1TextSpellcheckBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *V1TextSpellcheckBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1TextSpellcheckBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1TextSpellcheckBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1TextSpellcheckBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


