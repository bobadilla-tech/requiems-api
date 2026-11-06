# V1SystemsUserVerifyPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | Pointer to **bool** | True when risk score is below 0.3, confidence is above 0.5, email is valid, domain has MX, and domain is registered | [optional] 
**Confidence** | Pointer to **string** | Confidence in the verification result, based on how many signals resolved | [optional] 
**RiskScore** | Pointer to **string** | Composite risk score from 0.0 to 1.0 | [optional] 
**Flags** | Pointer to **[]interface{}** | Triggered risk flags. Possible values: email_invalid, disposable_email, no_mx, domain_not_registered, young_domain, whois_unavailable, ip_risk | [optional] 
**SignalsEmail** | Pointer to **map[string]interface{}** | Email validation result | [optional] 
**SignalsDomain** | Pointer to **map[string]interface{}** | Domain intelligence including WHOIS age, MX and A record presence, and registration status | [optional] 
**SignalsIp** | Pointer to **string** | IP risk signals. Null when ip_address was not provided | [optional] 

## Methods

### NewV1SystemsUserVerifyPost200ResponseData

`func NewV1SystemsUserVerifyPost200ResponseData() *V1SystemsUserVerifyPost200ResponseData`

NewV1SystemsUserVerifyPost200ResponseData instantiates a new V1SystemsUserVerifyPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsUserVerifyPost200ResponseDataWithDefaults

`func NewV1SystemsUserVerifyPost200ResponseDataWithDefaults() *V1SystemsUserVerifyPost200ResponseData`

NewV1SystemsUserVerifyPost200ResponseDataWithDefaults instantiates a new V1SystemsUserVerifyPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerified

`func (o *V1SystemsUserVerifyPost200ResponseData) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *V1SystemsUserVerifyPost200ResponseData) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *V1SystemsUserVerifyPost200ResponseData) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *V1SystemsUserVerifyPost200ResponseData) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetConfidence

`func (o *V1SystemsUserVerifyPost200ResponseData) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *V1SystemsUserVerifyPost200ResponseData) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *V1SystemsUserVerifyPost200ResponseData) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *V1SystemsUserVerifyPost200ResponseData) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetRiskScore

`func (o *V1SystemsUserVerifyPost200ResponseData) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *V1SystemsUserVerifyPost200ResponseData) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *V1SystemsUserVerifyPost200ResponseData) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *V1SystemsUserVerifyPost200ResponseData) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetFlags

`func (o *V1SystemsUserVerifyPost200ResponseData) GetFlags() []interface{}`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V1SystemsUserVerifyPost200ResponseData) GetFlagsOk() (*[]interface{}, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V1SystemsUserVerifyPost200ResponseData) SetFlags(v []interface{})`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V1SystemsUserVerifyPost200ResponseData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetSignalsEmail

`func (o *V1SystemsUserVerifyPost200ResponseData) GetSignalsEmail() map[string]interface{}`

GetSignalsEmail returns the SignalsEmail field if non-nil, zero value otherwise.

### GetSignalsEmailOk

`func (o *V1SystemsUserVerifyPost200ResponseData) GetSignalsEmailOk() (*map[string]interface{}, bool)`

GetSignalsEmailOk returns a tuple with the SignalsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsEmail

`func (o *V1SystemsUserVerifyPost200ResponseData) SetSignalsEmail(v map[string]interface{})`

SetSignalsEmail sets SignalsEmail field to given value.

### HasSignalsEmail

`func (o *V1SystemsUserVerifyPost200ResponseData) HasSignalsEmail() bool`

HasSignalsEmail returns a boolean if a field has been set.

### GetSignalsDomain

`func (o *V1SystemsUserVerifyPost200ResponseData) GetSignalsDomain() map[string]interface{}`

GetSignalsDomain returns the SignalsDomain field if non-nil, zero value otherwise.

### GetSignalsDomainOk

`func (o *V1SystemsUserVerifyPost200ResponseData) GetSignalsDomainOk() (*map[string]interface{}, bool)`

GetSignalsDomainOk returns a tuple with the SignalsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsDomain

`func (o *V1SystemsUserVerifyPost200ResponseData) SetSignalsDomain(v map[string]interface{})`

SetSignalsDomain sets SignalsDomain field to given value.

### HasSignalsDomain

`func (o *V1SystemsUserVerifyPost200ResponseData) HasSignalsDomain() bool`

HasSignalsDomain returns a boolean if a field has been set.

### GetSignalsIp

`func (o *V1SystemsUserVerifyPost200ResponseData) GetSignalsIp() string`

GetSignalsIp returns the SignalsIp field if non-nil, zero value otherwise.

### GetSignalsIpOk

`func (o *V1SystemsUserVerifyPost200ResponseData) GetSignalsIpOk() (*string, bool)`

GetSignalsIpOk returns a tuple with the SignalsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsIp

`func (o *V1SystemsUserVerifyPost200ResponseData) SetSignalsIp(v string)`

SetSignalsIp sets SignalsIp field to given value.

### HasSignalsIp

`func (o *V1SystemsUserVerifyPost200ResponseData) HasSignalsIp() bool`

HasSignalsIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


