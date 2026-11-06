# V1FinanceMortgageGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | Pointer to **float32** | The original loan amount passed in the request | [optional] 
**Rate** | Pointer to **float32** | The annual interest rate passed in the request | [optional] 
**Years** | Pointer to **int32** | The loan term in years passed in the request | [optional] 
**MonthlyPayment** | Pointer to **float32** | Fixed monthly payment amount (rounded to 2 decimal places) | [optional] 
**TotalPayment** | Pointer to **float32** | Total amount paid over the life of the loan | [optional] 
**TotalInterest** | Pointer to **float32** | Total interest paid (total_payment minus principal) | [optional] 
**Schedule** | Pointer to **[]interface{}** | Full amortization schedule — one entry per month (years × 12 entries) | [optional] 
**ScheduleMonth** | Pointer to **int32** | Month number (1 to years × 12) | [optional] 
**SchedulePayment** | Pointer to **float32** | Total payment for this month | [optional] 
**SchedulePrincipal** | Pointer to **float32** | Portion of this month&#39;s payment applied to principal | [optional] 
**ScheduleInterest** | Pointer to **float32** | Portion of this month&#39;s payment applied to interest | [optional] 
**ScheduleBalance** | Pointer to **float32** | Remaining loan balance after this payment | [optional] 

## Methods

### NewV1FinanceMortgageGet200ResponseData

`func NewV1FinanceMortgageGet200ResponseData() *V1FinanceMortgageGet200ResponseData`

NewV1FinanceMortgageGet200ResponseData instantiates a new V1FinanceMortgageGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FinanceMortgageGet200ResponseDataWithDefaults

`func NewV1FinanceMortgageGet200ResponseDataWithDefaults() *V1FinanceMortgageGet200ResponseData`

NewV1FinanceMortgageGet200ResponseDataWithDefaults instantiates a new V1FinanceMortgageGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *V1FinanceMortgageGet200ResponseData) GetPrincipal() float32`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *V1FinanceMortgageGet200ResponseData) GetPrincipalOk() (*float32, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *V1FinanceMortgageGet200ResponseData) SetPrincipal(v float32)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *V1FinanceMortgageGet200ResponseData) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetRate

`func (o *V1FinanceMortgageGet200ResponseData) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *V1FinanceMortgageGet200ResponseData) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *V1FinanceMortgageGet200ResponseData) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *V1FinanceMortgageGet200ResponseData) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetYears

`func (o *V1FinanceMortgageGet200ResponseData) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *V1FinanceMortgageGet200ResponseData) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *V1FinanceMortgageGet200ResponseData) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *V1FinanceMortgageGet200ResponseData) HasYears() bool`

HasYears returns a boolean if a field has been set.

### GetMonthlyPayment

`func (o *V1FinanceMortgageGet200ResponseData) GetMonthlyPayment() float32`

GetMonthlyPayment returns the MonthlyPayment field if non-nil, zero value otherwise.

### GetMonthlyPaymentOk

`func (o *V1FinanceMortgageGet200ResponseData) GetMonthlyPaymentOk() (*float32, bool)`

GetMonthlyPaymentOk returns a tuple with the MonthlyPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPayment

`func (o *V1FinanceMortgageGet200ResponseData) SetMonthlyPayment(v float32)`

SetMonthlyPayment sets MonthlyPayment field to given value.

### HasMonthlyPayment

`func (o *V1FinanceMortgageGet200ResponseData) HasMonthlyPayment() bool`

HasMonthlyPayment returns a boolean if a field has been set.

### GetTotalPayment

`func (o *V1FinanceMortgageGet200ResponseData) GetTotalPayment() float32`

GetTotalPayment returns the TotalPayment field if non-nil, zero value otherwise.

### GetTotalPaymentOk

`func (o *V1FinanceMortgageGet200ResponseData) GetTotalPaymentOk() (*float32, bool)`

GetTotalPaymentOk returns a tuple with the TotalPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPayment

`func (o *V1FinanceMortgageGet200ResponseData) SetTotalPayment(v float32)`

SetTotalPayment sets TotalPayment field to given value.

### HasTotalPayment

`func (o *V1FinanceMortgageGet200ResponseData) HasTotalPayment() bool`

HasTotalPayment returns a boolean if a field has been set.

### GetTotalInterest

`func (o *V1FinanceMortgageGet200ResponseData) GetTotalInterest() float32`

GetTotalInterest returns the TotalInterest field if non-nil, zero value otherwise.

### GetTotalInterestOk

`func (o *V1FinanceMortgageGet200ResponseData) GetTotalInterestOk() (*float32, bool)`

GetTotalInterestOk returns a tuple with the TotalInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInterest

`func (o *V1FinanceMortgageGet200ResponseData) SetTotalInterest(v float32)`

SetTotalInterest sets TotalInterest field to given value.

### HasTotalInterest

`func (o *V1FinanceMortgageGet200ResponseData) HasTotalInterest() bool`

HasTotalInterest returns a boolean if a field has been set.

### GetSchedule

`func (o *V1FinanceMortgageGet200ResponseData) GetSchedule() []interface{}`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *V1FinanceMortgageGet200ResponseData) GetScheduleOk() (*[]interface{}, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *V1FinanceMortgageGet200ResponseData) SetSchedule(v []interface{})`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *V1FinanceMortgageGet200ResponseData) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScheduleMonth

`func (o *V1FinanceMortgageGet200ResponseData) GetScheduleMonth() int32`

GetScheduleMonth returns the ScheduleMonth field if non-nil, zero value otherwise.

### GetScheduleMonthOk

`func (o *V1FinanceMortgageGet200ResponseData) GetScheduleMonthOk() (*int32, bool)`

GetScheduleMonthOk returns a tuple with the ScheduleMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMonth

`func (o *V1FinanceMortgageGet200ResponseData) SetScheduleMonth(v int32)`

SetScheduleMonth sets ScheduleMonth field to given value.

### HasScheduleMonth

`func (o *V1FinanceMortgageGet200ResponseData) HasScheduleMonth() bool`

HasScheduleMonth returns a boolean if a field has been set.

### GetSchedulePayment

`func (o *V1FinanceMortgageGet200ResponseData) GetSchedulePayment() float32`

GetSchedulePayment returns the SchedulePayment field if non-nil, zero value otherwise.

### GetSchedulePaymentOk

`func (o *V1FinanceMortgageGet200ResponseData) GetSchedulePaymentOk() (*float32, bool)`

GetSchedulePaymentOk returns a tuple with the SchedulePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulePayment

`func (o *V1FinanceMortgageGet200ResponseData) SetSchedulePayment(v float32)`

SetSchedulePayment sets SchedulePayment field to given value.

### HasSchedulePayment

`func (o *V1FinanceMortgageGet200ResponseData) HasSchedulePayment() bool`

HasSchedulePayment returns a boolean if a field has been set.

### GetSchedulePrincipal

`func (o *V1FinanceMortgageGet200ResponseData) GetSchedulePrincipal() float32`

GetSchedulePrincipal returns the SchedulePrincipal field if non-nil, zero value otherwise.

### GetSchedulePrincipalOk

`func (o *V1FinanceMortgageGet200ResponseData) GetSchedulePrincipalOk() (*float32, bool)`

GetSchedulePrincipalOk returns a tuple with the SchedulePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulePrincipal

`func (o *V1FinanceMortgageGet200ResponseData) SetSchedulePrincipal(v float32)`

SetSchedulePrincipal sets SchedulePrincipal field to given value.

### HasSchedulePrincipal

`func (o *V1FinanceMortgageGet200ResponseData) HasSchedulePrincipal() bool`

HasSchedulePrincipal returns a boolean if a field has been set.

### GetScheduleInterest

`func (o *V1FinanceMortgageGet200ResponseData) GetScheduleInterest() float32`

GetScheduleInterest returns the ScheduleInterest field if non-nil, zero value otherwise.

### GetScheduleInterestOk

`func (o *V1FinanceMortgageGet200ResponseData) GetScheduleInterestOk() (*float32, bool)`

GetScheduleInterestOk returns a tuple with the ScheduleInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterest

`func (o *V1FinanceMortgageGet200ResponseData) SetScheduleInterest(v float32)`

SetScheduleInterest sets ScheduleInterest field to given value.

### HasScheduleInterest

`func (o *V1FinanceMortgageGet200ResponseData) HasScheduleInterest() bool`

HasScheduleInterest returns a boolean if a field has been set.

### GetScheduleBalance

`func (o *V1FinanceMortgageGet200ResponseData) GetScheduleBalance() float32`

GetScheduleBalance returns the ScheduleBalance field if non-nil, zero value otherwise.

### GetScheduleBalanceOk

`func (o *V1FinanceMortgageGet200ResponseData) GetScheduleBalanceOk() (*float32, bool)`

GetScheduleBalanceOk returns a tuple with the ScheduleBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleBalance

`func (o *V1FinanceMortgageGet200ResponseData) SetScheduleBalance(v float32)`

SetScheduleBalance sets ScheduleBalance field to given value.

### HasScheduleBalance

`func (o *V1FinanceMortgageGet200ResponseData) HasScheduleBalance() bool`

HasScheduleBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


