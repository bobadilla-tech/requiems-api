# V1FinanceInflationBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | One entry per country, in the same order as the input array | [optional] 
**ResultsCountry** | Pointer to **string** | ISO 3166-1 alpha-2 country code, uppercased | [optional] 
**ResultsFound** | Pointer to **bool** | false when the country has no data in the World Bank set | [optional] 
**ResultsRate** | Pointer to **float32** | Latest CPI inflation rate. Omitted when found: false | [optional] 
**ResultsPeriod** | Pointer to **string** | Year of the latest data point. Omitted when found: false | [optional] 
**ResultsHistorical** | Pointer to **[]interface{}** | Up to 10 previous years. Omitted when found: false | [optional] 
**Total** | Pointer to **int32** | Total number of results returned (equals number of countries sent) | [optional] 

## Methods

### NewV1FinanceInflationBatchPost200ResponseData

`func NewV1FinanceInflationBatchPost200ResponseData() *V1FinanceInflationBatchPost200ResponseData`

NewV1FinanceInflationBatchPost200ResponseData instantiates a new V1FinanceInflationBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceInflationBatchPost200ResponseDataWithDefaults

`func NewV1FinanceInflationBatchPost200ResponseDataWithDefaults() *V1FinanceInflationBatchPost200ResponseData`

NewV1FinanceInflationBatchPost200ResponseDataWithDefaults instantiates a new V1FinanceInflationBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1FinanceInflationBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1FinanceInflationBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsCountry

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsCountry() string`

GetResultsCountry returns the ResultsCountry field if non-nil, zero value otherwise.

### GetResultsCountryOk

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsCountryOk() (*string, bool)`

GetResultsCountryOk returns a tuple with the ResultsCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsCountry

`func (o *V1FinanceInflationBatchPost200ResponseData) SetResultsCountry(v string)`

SetResultsCountry sets ResultsCountry field to given value.

### HasResultsCountry

`func (o *V1FinanceInflationBatchPost200ResponseData) HasResultsCountry() bool`

HasResultsCountry returns a boolean if a field has been set.

### GetResultsFound

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsFound() bool`

GetResultsFound returns the ResultsFound field if non-nil, zero value otherwise.

### GetResultsFoundOk

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsFoundOk() (*bool, bool)`

GetResultsFoundOk returns a tuple with the ResultsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsFound

`func (o *V1FinanceInflationBatchPost200ResponseData) SetResultsFound(v bool)`

SetResultsFound sets ResultsFound field to given value.

### HasResultsFound

`func (o *V1FinanceInflationBatchPost200ResponseData) HasResultsFound() bool`

HasResultsFound returns a boolean if a field has been set.

### GetResultsRate

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsRate() float32`

GetResultsRate returns the ResultsRate field if non-nil, zero value otherwise.

### GetResultsRateOk

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsRateOk() (*float32, bool)`

GetResultsRateOk returns a tuple with the ResultsRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsRate

`func (o *V1FinanceInflationBatchPost200ResponseData) SetResultsRate(v float32)`

SetResultsRate sets ResultsRate field to given value.

### HasResultsRate

`func (o *V1FinanceInflationBatchPost200ResponseData) HasResultsRate() bool`

HasResultsRate returns a boolean if a field has been set.

### GetResultsPeriod

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsPeriod() string`

GetResultsPeriod returns the ResultsPeriod field if non-nil, zero value otherwise.

### GetResultsPeriodOk

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsPeriodOk() (*string, bool)`

GetResultsPeriodOk returns a tuple with the ResultsPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPeriod

`func (o *V1FinanceInflationBatchPost200ResponseData) SetResultsPeriod(v string)`

SetResultsPeriod sets ResultsPeriod field to given value.

### HasResultsPeriod

`func (o *V1FinanceInflationBatchPost200ResponseData) HasResultsPeriod() bool`

HasResultsPeriod returns a boolean if a field has been set.

### GetResultsHistorical

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsHistorical() []interface{}`

GetResultsHistorical returns the ResultsHistorical field if non-nil, zero value otherwise.

### GetResultsHistoricalOk

`func (o *V1FinanceInflationBatchPost200ResponseData) GetResultsHistoricalOk() (*[]interface{}, bool)`

GetResultsHistoricalOk returns a tuple with the ResultsHistorical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsHistorical

`func (o *V1FinanceInflationBatchPost200ResponseData) SetResultsHistorical(v []interface{})`

SetResultsHistorical sets ResultsHistorical field to given value.

### HasResultsHistorical

`func (o *V1FinanceInflationBatchPost200ResponseData) HasResultsHistorical() bool`

HasResultsHistorical returns a boolean if a field has been set.

### GetTotal

`func (o *V1FinanceInflationBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1FinanceInflationBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1FinanceInflationBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1FinanceInflationBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


