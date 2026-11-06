# V1SystemsTransactionRiskPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBin** | **string** | Card BIN (first 6–8 digits) to look up card issuing country. | 
**IpAddress** | **string** | Customer&#39;s IPv4 or IPv6 address. Must be a valid IP. | 
**BillingCountry** | Pointer to **string** | ISO 3166-1 alpha-2 billing country code. When provided, country mismatch checks are more precise. | [optional] 
**AmountUsd** | Pointer to **float32** | Transaction amount in USD. High-value transactions via VPN incur an additional risk score boost. | [optional] 

## Methods

### NewV1SystemsTransactionRiskPostRequest

`func NewV1SystemsTransactionRiskPostRequest(cardBin string, ipAddress string, ) *V1SystemsTransactionRiskPostRequest`

NewV1SystemsTransactionRiskPostRequest instantiates a new V1SystemsTransactionRiskPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsTransactionRiskPostRequestWithDefaults

`func NewV1SystemsTransactionRiskPostRequestWithDefaults() *V1SystemsTransactionRiskPostRequest`

NewV1SystemsTransactionRiskPostRequestWithDefaults instantiates a new V1SystemsTransactionRiskPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBin

`func (o *V1SystemsTransactionRiskPostRequest) GetCardBin() string`

GetCardBin returns the CardBin field if non-nil, zero value otherwise.

### GetCardBinOk

`func (o *V1SystemsTransactionRiskPostRequest) GetCardBinOk() (*string, bool)`

GetCardBinOk returns a tuple with the CardBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBin

`func (o *V1SystemsTransactionRiskPostRequest) SetCardBin(v string)`

SetCardBin sets CardBin field to given value.


### GetIpAddress

`func (o *V1SystemsTransactionRiskPostRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *V1SystemsTransactionRiskPostRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *V1SystemsTransactionRiskPostRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetBillingCountry

`func (o *V1SystemsTransactionRiskPostRequest) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *V1SystemsTransactionRiskPostRequest) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *V1SystemsTransactionRiskPostRequest) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.

### HasBillingCountry

`func (o *V1SystemsTransactionRiskPostRequest) HasBillingCountry() bool`

HasBillingCountry returns a boolean if a field has been set.

### GetAmountUsd

`func (o *V1SystemsTransactionRiskPostRequest) GetAmountUsd() float32`

GetAmountUsd returns the AmountUsd field if non-nil, zero value otherwise.

### GetAmountUsdOk

`func (o *V1SystemsTransactionRiskPostRequest) GetAmountUsdOk() (*float32, bool)`

GetAmountUsdOk returns a tuple with the AmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUsd

`func (o *V1SystemsTransactionRiskPostRequest) SetAmountUsd(v float32)`

SetAmountUsd sets AmountUsd field to given value.

### HasAmountUsd

`func (o *V1SystemsTransactionRiskPostRequest) HasAmountUsd() bool`

HasAmountUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


