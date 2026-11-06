# V1SystemsTransactionRiskPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScore** | Pointer to **string** | Transaction risk score from 0.0 to 1.0 | [optional] 
**IsSafe** | Pointer to **bool** | True when score is below 0.5 and no TOR or proxy detected | [optional] 
**Flags** | Pointer to **[]interface{}** | Triggered flags. Possible values: ip_country_mismatch, bin_country_mismatch, country_mismatch, vpn_detected, proxy_detected, tor_detected, high_value_vpn | [optional] 
**SignalsIpCountry** | Pointer to **string** | Country code resolved from the IP address | [optional] 
**SignalsBillingCountry** | Pointer to **string** | Billing country provided in the request, or null if omitted | [optional] 
**SignalsBinCountry** | Pointer to **string** | Country of the card issuer resolved from the BIN | [optional] 
**SignalsCountryMismatch** | Pointer to **bool** | True when any country comparison produced a mismatch | [optional] 

## Methods

### NewV1SystemsTransactionRiskPost200ResponseData

`func NewV1SystemsTransactionRiskPost200ResponseData() *V1SystemsTransactionRiskPost200ResponseData`

NewV1SystemsTransactionRiskPost200ResponseData instantiates a new V1SystemsTransactionRiskPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsTransactionRiskPost200ResponseDataWithDefaults

`func NewV1SystemsTransactionRiskPost200ResponseDataWithDefaults() *V1SystemsTransactionRiskPost200ResponseData`

NewV1SystemsTransactionRiskPost200ResponseDataWithDefaults instantiates a new V1SystemsTransactionRiskPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskScore

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *V1SystemsTransactionRiskPost200ResponseData) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *V1SystemsTransactionRiskPost200ResponseData) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetIsSafe

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetIsSafe() bool`

GetIsSafe returns the IsSafe field if non-nil, zero value otherwise.

### GetIsSafeOk

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetIsSafeOk() (*bool, bool)`

GetIsSafeOk returns a tuple with the IsSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafe

`func (o *V1SystemsTransactionRiskPost200ResponseData) SetIsSafe(v bool)`

SetIsSafe sets IsSafe field to given value.

### HasIsSafe

`func (o *V1SystemsTransactionRiskPost200ResponseData) HasIsSafe() bool`

HasIsSafe returns a boolean if a field has been set.

### GetFlags

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetFlags() []interface{}`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetFlagsOk() (*[]interface{}, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V1SystemsTransactionRiskPost200ResponseData) SetFlags(v []interface{})`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V1SystemsTransactionRiskPost200ResponseData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetSignalsIpCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsIpCountry() string`

GetSignalsIpCountry returns the SignalsIpCountry field if non-nil, zero value otherwise.

### GetSignalsIpCountryOk

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsIpCountryOk() (*string, bool)`

GetSignalsIpCountryOk returns a tuple with the SignalsIpCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsIpCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) SetSignalsIpCountry(v string)`

SetSignalsIpCountry sets SignalsIpCountry field to given value.

### HasSignalsIpCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) HasSignalsIpCountry() bool`

HasSignalsIpCountry returns a boolean if a field has been set.

### GetSignalsBillingCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsBillingCountry() string`

GetSignalsBillingCountry returns the SignalsBillingCountry field if non-nil, zero value otherwise.

### GetSignalsBillingCountryOk

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsBillingCountryOk() (*string, bool)`

GetSignalsBillingCountryOk returns a tuple with the SignalsBillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsBillingCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) SetSignalsBillingCountry(v string)`

SetSignalsBillingCountry sets SignalsBillingCountry field to given value.

### HasSignalsBillingCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) HasSignalsBillingCountry() bool`

HasSignalsBillingCountry returns a boolean if a field has been set.

### GetSignalsBinCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsBinCountry() string`

GetSignalsBinCountry returns the SignalsBinCountry field if non-nil, zero value otherwise.

### GetSignalsBinCountryOk

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsBinCountryOk() (*string, bool)`

GetSignalsBinCountryOk returns a tuple with the SignalsBinCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsBinCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) SetSignalsBinCountry(v string)`

SetSignalsBinCountry sets SignalsBinCountry field to given value.

### HasSignalsBinCountry

`func (o *V1SystemsTransactionRiskPost200ResponseData) HasSignalsBinCountry() bool`

HasSignalsBinCountry returns a boolean if a field has been set.

### GetSignalsCountryMismatch

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsCountryMismatch() bool`

GetSignalsCountryMismatch returns the SignalsCountryMismatch field if non-nil, zero value otherwise.

### GetSignalsCountryMismatchOk

`func (o *V1SystemsTransactionRiskPost200ResponseData) GetSignalsCountryMismatchOk() (*bool, bool)`

GetSignalsCountryMismatchOk returns a tuple with the SignalsCountryMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsCountryMismatch

`func (o *V1SystemsTransactionRiskPost200ResponseData) SetSignalsCountryMismatch(v bool)`

SetSignalsCountryMismatch sets SignalsCountryMismatch field to given value.

### HasSignalsCountryMismatch

`func (o *V1SystemsTransactionRiskPost200ResponseData) HasSignalsCountryMismatch() bool`

HasSignalsCountryMismatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


