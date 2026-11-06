# V1SystemsSignupProtectPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScore** | Pointer to **string** | Composite risk score from 0.0 (clean) to 1.0 (high risk) | [optional] 
**IsSafe** | Pointer to **bool** | True when risk_score is below the safety threshold with no critical flags | [optional] 
**Confidence** | Pointer to **string** | Confidence in the decision from 0.0 to 1.0, based on which signals resolved successfully | [optional] 
**Flags** | Pointer to **[]interface{}** | Array of triggered risk flags. Possible values: disposable_email, vpn_detected, proxy_detected, tor_detected, is_hosting | [optional] 
**SignalsEmail** | Pointer to **map[string]interface{}** | Email validation result including disposable detection and MX check | [optional] 
**SignalsPhone** | Pointer to **map[string]interface{}** | Phone number validation result with country code and VoIP/virtual detection | [optional] 
**SignalsIp** | Pointer to **map[string]interface{}** | IP intelligence including VPN/proxy/TOR status and fraud score | [optional] 

## Methods

### NewV1SystemsSignupProtectPost200ResponseData

`func NewV1SystemsSignupProtectPost200ResponseData() *V1SystemsSignupProtectPost200ResponseData`

NewV1SystemsSignupProtectPost200ResponseData instantiates a new V1SystemsSignupProtectPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsSignupProtectPost200ResponseDataWithDefaults

`func NewV1SystemsSignupProtectPost200ResponseDataWithDefaults() *V1SystemsSignupProtectPost200ResponseData`

NewV1SystemsSignupProtectPost200ResponseDataWithDefaults instantiates a new V1SystemsSignupProtectPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskScore

`func (o *V1SystemsSignupProtectPost200ResponseData) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *V1SystemsSignupProtectPost200ResponseData) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *V1SystemsSignupProtectPost200ResponseData) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *V1SystemsSignupProtectPost200ResponseData) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetIsSafe

`func (o *V1SystemsSignupProtectPost200ResponseData) GetIsSafe() bool`

GetIsSafe returns the IsSafe field if non-nil, zero value otherwise.

### GetIsSafeOk

`func (o *V1SystemsSignupProtectPost200ResponseData) GetIsSafeOk() (*bool, bool)`

GetIsSafeOk returns a tuple with the IsSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafe

`func (o *V1SystemsSignupProtectPost200ResponseData) SetIsSafe(v bool)`

SetIsSafe sets IsSafe field to given value.

### HasIsSafe

`func (o *V1SystemsSignupProtectPost200ResponseData) HasIsSafe() bool`

HasIsSafe returns a boolean if a field has been set.

### GetConfidence

`func (o *V1SystemsSignupProtectPost200ResponseData) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *V1SystemsSignupProtectPost200ResponseData) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *V1SystemsSignupProtectPost200ResponseData) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *V1SystemsSignupProtectPost200ResponseData) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetFlags

`func (o *V1SystemsSignupProtectPost200ResponseData) GetFlags() []interface{}`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V1SystemsSignupProtectPost200ResponseData) GetFlagsOk() (*[]interface{}, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V1SystemsSignupProtectPost200ResponseData) SetFlags(v []interface{})`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V1SystemsSignupProtectPost200ResponseData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetSignalsEmail

`func (o *V1SystemsSignupProtectPost200ResponseData) GetSignalsEmail() map[string]interface{}`

GetSignalsEmail returns the SignalsEmail field if non-nil, zero value otherwise.

### GetSignalsEmailOk

`func (o *V1SystemsSignupProtectPost200ResponseData) GetSignalsEmailOk() (*map[string]interface{}, bool)`

GetSignalsEmailOk returns a tuple with the SignalsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsEmail

`func (o *V1SystemsSignupProtectPost200ResponseData) SetSignalsEmail(v map[string]interface{})`

SetSignalsEmail sets SignalsEmail field to given value.

### HasSignalsEmail

`func (o *V1SystemsSignupProtectPost200ResponseData) HasSignalsEmail() bool`

HasSignalsEmail returns a boolean if a field has been set.

### GetSignalsPhone

`func (o *V1SystemsSignupProtectPost200ResponseData) GetSignalsPhone() map[string]interface{}`

GetSignalsPhone returns the SignalsPhone field if non-nil, zero value otherwise.

### GetSignalsPhoneOk

`func (o *V1SystemsSignupProtectPost200ResponseData) GetSignalsPhoneOk() (*map[string]interface{}, bool)`

GetSignalsPhoneOk returns a tuple with the SignalsPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsPhone

`func (o *V1SystemsSignupProtectPost200ResponseData) SetSignalsPhone(v map[string]interface{})`

SetSignalsPhone sets SignalsPhone field to given value.

### HasSignalsPhone

`func (o *V1SystemsSignupProtectPost200ResponseData) HasSignalsPhone() bool`

HasSignalsPhone returns a boolean if a field has been set.

### GetSignalsIp

`func (o *V1SystemsSignupProtectPost200ResponseData) GetSignalsIp() map[string]interface{}`

GetSignalsIp returns the SignalsIp field if non-nil, zero value otherwise.

### GetSignalsIpOk

`func (o *V1SystemsSignupProtectPost200ResponseData) GetSignalsIpOk() (*map[string]interface{}, bool)`

GetSignalsIpOk returns a tuple with the SignalsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsIp

`func (o *V1SystemsSignupProtectPost200ResponseData) SetSignalsIp(v map[string]interface{})`

SetSignalsIp sets SignalsIp field to given value.

### HasSignalsIp

`func (o *V1SystemsSignupProtectPost200ResponseData) HasSignalsIp() bool`

HasSignalsIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


