# V1FinanceCommoditiesCommodityGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commodity** | Pointer to **string** | The commodity slug as provided in the request path | [optional] 
**Name** | Pointer to **string** | Human-readable commodity name | [optional] 
**Price** | Pointer to **float32** | Latest annual average price in the commodity&#39;s display unit | [optional] 
**Unit** | Pointer to **string** | Price unit (oz, barrel, mmbtu, lb, or metric_ton) | [optional] 
**Currency** | Pointer to **string** | Currency code — always USD | [optional] 
**Change24h** | Pointer to **float32** | Year-over-year percentage change from the prior year&#39;s annual average (positive &#x3D; price increased) | [optional] 
**Historical** | Pointer to **[]interface{}** | Up to 10 prior years of annual average prices, ordered newest to oldest | [optional] 
**HistoricalPeriod** | Pointer to **string** | Year of the historical data point | [optional] 
**HistoricalPrice** | Pointer to **float32** | Annual average price for that year | [optional] 

## Methods

### NewV1FinanceCommoditiesCommodityGet200ResponseData

`func NewV1FinanceCommoditiesCommodityGet200ResponseData() *V1FinanceCommoditiesCommodityGet200ResponseData`

NewV1FinanceCommoditiesCommodityGet200ResponseData instantiates a new V1FinanceCommoditiesCommodityGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceCommoditiesCommodityGet200ResponseDataWithDefaults

`func NewV1FinanceCommoditiesCommodityGet200ResponseDataWithDefaults() *V1FinanceCommoditiesCommodityGet200ResponseData`

NewV1FinanceCommoditiesCommodityGet200ResponseDataWithDefaults instantiates a new V1FinanceCommoditiesCommodityGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommodity

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetCommodity() string`

GetCommodity returns the Commodity field if non-nil, zero value otherwise.

### GetCommodityOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetCommodityOk() (*string, bool)`

GetCommodityOk returns a tuple with the Commodity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommodity

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetCommodity(v string)`

SetCommodity sets Commodity field to given value.

### HasCommodity

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasCommodity() bool`

HasCommodity returns a boolean if a field has been set.

### GetName

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetUnit

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChange24h

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetChange24h() float32`

GetChange24h returns the Change24h field if non-nil, zero value otherwise.

### GetChange24hOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetChange24hOk() (*float32, bool)`

GetChange24hOk returns a tuple with the Change24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange24h

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetChange24h(v float32)`

SetChange24h sets Change24h field to given value.

### HasChange24h

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasChange24h() bool`

HasChange24h returns a boolean if a field has been set.

### GetHistorical

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetHistorical() []interface{}`

GetHistorical returns the Historical field if non-nil, zero value otherwise.

### GetHistoricalOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetHistoricalOk() (*[]interface{}, bool)`

GetHistoricalOk returns a tuple with the Historical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistorical

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetHistorical(v []interface{})`

SetHistorical sets Historical field to given value.

### HasHistorical

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasHistorical() bool`

HasHistorical returns a boolean if a field has been set.

### GetHistoricalPeriod

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetHistoricalPeriod() string`

GetHistoricalPeriod returns the HistoricalPeriod field if non-nil, zero value otherwise.

### GetHistoricalPeriodOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetHistoricalPeriodOk() (*string, bool)`

GetHistoricalPeriodOk returns a tuple with the HistoricalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalPeriod

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetHistoricalPeriod(v string)`

SetHistoricalPeriod sets HistoricalPeriod field to given value.

### HasHistoricalPeriod

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasHistoricalPeriod() bool`

HasHistoricalPeriod returns a boolean if a field has been set.

### GetHistoricalPrice

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetHistoricalPrice() float32`

GetHistoricalPrice returns the HistoricalPrice field if non-nil, zero value otherwise.

### GetHistoricalPriceOk

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) GetHistoricalPriceOk() (*float32, bool)`

GetHistoricalPriceOk returns a tuple with the HistoricalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalPrice

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) SetHistoricalPrice(v float32)`

SetHistoricalPrice sets HistoricalPrice field to given value.

### HasHistoricalPrice

`func (o *V1FinanceCommoditiesCommodityGet200ResponseData) HasHistoricalPrice() bool`

HasHistoricalPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


