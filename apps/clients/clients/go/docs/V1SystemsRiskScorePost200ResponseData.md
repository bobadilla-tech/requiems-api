# V1SystemsRiskScorePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScore** | Pointer to **string** | Composite risk score from 0.0 to 1.0 | [optional] 
**IsSafe** | Pointer to **bool** | True when the risk score is below the safety threshold | [optional] 
**Confidence** | Pointer to **string** | Confidence in the score based on resolved signals | [optional] 
**Flags** | Pointer to **[]interface{}** | Array of triggered risk flags | [optional] 

## Methods

### NewV1SystemsRiskScorePost200ResponseData

`func NewV1SystemsRiskScorePost200ResponseData() *V1SystemsRiskScorePost200ResponseData`

NewV1SystemsRiskScorePost200ResponseData instantiates a new V1SystemsRiskScorePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsRiskScorePost200ResponseDataWithDefaults

`func NewV1SystemsRiskScorePost200ResponseDataWithDefaults() *V1SystemsRiskScorePost200ResponseData`

NewV1SystemsRiskScorePost200ResponseDataWithDefaults instantiates a new V1SystemsRiskScorePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskScore

`func (o *V1SystemsRiskScorePost200ResponseData) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *V1SystemsRiskScorePost200ResponseData) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *V1SystemsRiskScorePost200ResponseData) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *V1SystemsRiskScorePost200ResponseData) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetIsSafe

`func (o *V1SystemsRiskScorePost200ResponseData) GetIsSafe() bool`

GetIsSafe returns the IsSafe field if non-nil, zero value otherwise.

### GetIsSafeOk

`func (o *V1SystemsRiskScorePost200ResponseData) GetIsSafeOk() (*bool, bool)`

GetIsSafeOk returns a tuple with the IsSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafe

`func (o *V1SystemsRiskScorePost200ResponseData) SetIsSafe(v bool)`

SetIsSafe sets IsSafe field to given value.

### HasIsSafe

`func (o *V1SystemsRiskScorePost200ResponseData) HasIsSafe() bool`

HasIsSafe returns a boolean if a field has been set.

### GetConfidence

`func (o *V1SystemsRiskScorePost200ResponseData) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *V1SystemsRiskScorePost200ResponseData) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *V1SystemsRiskScorePost200ResponseData) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *V1SystemsRiskScorePost200ResponseData) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetFlags

`func (o *V1SystemsRiskScorePost200ResponseData) GetFlags() []interface{}`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V1SystemsRiskScorePost200ResponseData) GetFlagsOk() (*[]interface{}, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V1SystemsRiskScorePost200ResponseData) SetFlags(v []interface{})`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V1SystemsRiskScorePost200ResponseData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


