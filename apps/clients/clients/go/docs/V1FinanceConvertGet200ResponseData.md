# V1FinanceConvertGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Source currency code (uppercased) | [optional] 
**To** | Pointer to **string** | Target currency code (uppercased) | [optional] 
**Rate** | Pointer to **float32** | Exchange rate used for the conversion | [optional] 
**Amount** | Pointer to **float32** | The original amount passed in the request | [optional] 
**Converted** | Pointer to **float32** | Result of amount × rate, rounded to 2 decimal places | [optional] 
**Timestamp** | Pointer to **string** | Date the rate was published by the ECB (ISO 8601) | [optional] 

## Methods

### NewV1FinanceConvertGet200ResponseData

`func NewV1FinanceConvertGet200ResponseData() *V1FinanceConvertGet200ResponseData`

NewV1FinanceConvertGet200ResponseData instantiates a new V1FinanceConvertGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceConvertGet200ResponseDataWithDefaults

`func NewV1FinanceConvertGet200ResponseDataWithDefaults() *V1FinanceConvertGet200ResponseData`

NewV1FinanceConvertGet200ResponseDataWithDefaults instantiates a new V1FinanceConvertGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *V1FinanceConvertGet200ResponseData) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1FinanceConvertGet200ResponseData) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1FinanceConvertGet200ResponseData) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1FinanceConvertGet200ResponseData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *V1FinanceConvertGet200ResponseData) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1FinanceConvertGet200ResponseData) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1FinanceConvertGet200ResponseData) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *V1FinanceConvertGet200ResponseData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetRate

`func (o *V1FinanceConvertGet200ResponseData) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *V1FinanceConvertGet200ResponseData) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *V1FinanceConvertGet200ResponseData) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *V1FinanceConvertGet200ResponseData) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetAmount

`func (o *V1FinanceConvertGet200ResponseData) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *V1FinanceConvertGet200ResponseData) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *V1FinanceConvertGet200ResponseData) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *V1FinanceConvertGet200ResponseData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetConverted

`func (o *V1FinanceConvertGet200ResponseData) GetConverted() float32`

GetConverted returns the Converted field if non-nil, zero value otherwise.

### GetConvertedOk

`func (o *V1FinanceConvertGet200ResponseData) GetConvertedOk() (*float32, bool)`

GetConvertedOk returns a tuple with the Converted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConverted

`func (o *V1FinanceConvertGet200ResponseData) SetConverted(v float32)`

SetConverted sets Converted field to given value.

### HasConverted

`func (o *V1FinanceConvertGet200ResponseData) HasConverted() bool`

HasConverted returns a boolean if a field has been set.

### GetTimestamp

`func (o *V1FinanceConvertGet200ResponseData) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *V1FinanceConvertGet200ResponseData) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *V1FinanceConvertGet200ResponseData) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *V1FinanceConvertGet200ResponseData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


