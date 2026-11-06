# V1FinanceInflationGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country code, uppercased | [optional] 
**Rate** | Pointer to **float32** | Latest annual CPI inflation rate as a percentage (e.g. 2.9495 means 2.9495%) | [optional] 
**Period** | Pointer to **string** | Year of the latest data point (e.g. 2024) | [optional] 
**Historical** | Pointer to **[]interface{}** | Up to 10 previous years of inflation data, ordered newest to oldest | [optional] 
**HistoricalPeriod** | Pointer to **string** | Year of the historical data point | [optional] 
**HistoricalRate** | Pointer to **float32** | Annual CPI inflation rate for that year | [optional] 

## Methods

### NewV1FinanceInflationGet200ResponseData

`func NewV1FinanceInflationGet200ResponseData() *V1FinanceInflationGet200ResponseData`

NewV1FinanceInflationGet200ResponseData instantiates a new V1FinanceInflationGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceInflationGet200ResponseDataWithDefaults

`func NewV1FinanceInflationGet200ResponseDataWithDefaults() *V1FinanceInflationGet200ResponseData`

NewV1FinanceInflationGet200ResponseDataWithDefaults instantiates a new V1FinanceInflationGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *V1FinanceInflationGet200ResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V1FinanceInflationGet200ResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V1FinanceInflationGet200ResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V1FinanceInflationGet200ResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRate

`func (o *V1FinanceInflationGet200ResponseData) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *V1FinanceInflationGet200ResponseData) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *V1FinanceInflationGet200ResponseData) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *V1FinanceInflationGet200ResponseData) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetPeriod

`func (o *V1FinanceInflationGet200ResponseData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *V1FinanceInflationGet200ResponseData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *V1FinanceInflationGet200ResponseData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *V1FinanceInflationGet200ResponseData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetHistorical

`func (o *V1FinanceInflationGet200ResponseData) GetHistorical() []interface{}`

GetHistorical returns the Historical field if non-nil, zero value otherwise.

### GetHistoricalOk

`func (o *V1FinanceInflationGet200ResponseData) GetHistoricalOk() (*[]interface{}, bool)`

GetHistoricalOk returns a tuple with the Historical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistorical

`func (o *V1FinanceInflationGet200ResponseData) SetHistorical(v []interface{})`

SetHistorical sets Historical field to given value.

### HasHistorical

`func (o *V1FinanceInflationGet200ResponseData) HasHistorical() bool`

HasHistorical returns a boolean if a field has been set.

### GetHistoricalPeriod

`func (o *V1FinanceInflationGet200ResponseData) GetHistoricalPeriod() string`

GetHistoricalPeriod returns the HistoricalPeriod field if non-nil, zero value otherwise.

### GetHistoricalPeriodOk

`func (o *V1FinanceInflationGet200ResponseData) GetHistoricalPeriodOk() (*string, bool)`

GetHistoricalPeriodOk returns a tuple with the HistoricalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalPeriod

`func (o *V1FinanceInflationGet200ResponseData) SetHistoricalPeriod(v string)`

SetHistoricalPeriod sets HistoricalPeriod field to given value.

### HasHistoricalPeriod

`func (o *V1FinanceInflationGet200ResponseData) HasHistoricalPeriod() bool`

HasHistoricalPeriod returns a boolean if a field has been set.

### GetHistoricalRate

`func (o *V1FinanceInflationGet200ResponseData) GetHistoricalRate() float32`

GetHistoricalRate returns the HistoricalRate field if non-nil, zero value otherwise.

### GetHistoricalRateOk

`func (o *V1FinanceInflationGet200ResponseData) GetHistoricalRateOk() (*float32, bool)`

GetHistoricalRateOk returns a tuple with the HistoricalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalRate

`func (o *V1FinanceInflationGet200ResponseData) SetHistoricalRate(v float32)`

SetHistoricalRate sets HistoricalRate field to given value.

### HasHistoricalRate

`func (o *V1FinanceInflationGet200ResponseData) HasHistoricalRate() bool`

HasHistoricalRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


