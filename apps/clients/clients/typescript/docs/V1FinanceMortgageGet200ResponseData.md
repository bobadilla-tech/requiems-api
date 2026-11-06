# V1FinanceMortgageGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**principal** | **number** | The original loan amount passed in the request | [optional] [default to undefined]
**rate** | **number** | The annual interest rate passed in the request | [optional] [default to undefined]
**years** | **number** | The loan term in years passed in the request | [optional] [default to undefined]
**monthly_payment** | **number** | Fixed monthly payment amount (rounded to 2 decimal places) | [optional] [default to undefined]
**total_payment** | **number** | Total amount paid over the life of the loan | [optional] [default to undefined]
**total_interest** | **number** | Total interest paid (total_payment minus principal) | [optional] [default to undefined]
**schedule** | **Array&lt;any&gt;** | Full amortization schedule — one entry per month (years × 12 entries) | [optional] [default to undefined]
**schedule_month** | **number** | Month number (1 to years × 12) | [optional] [default to undefined]
**schedule_payment** | **number** | Total payment for this month | [optional] [default to undefined]
**schedule_principal** | **number** | Portion of this month\&#39;s payment applied to principal | [optional] [default to undefined]
**schedule_interest** | **number** | Portion of this month\&#39;s payment applied to interest | [optional] [default to undefined]
**schedule_balance** | **number** | Remaining loan balance after this payment | [optional] [default to undefined]

## Example

```typescript
import { V1FinanceMortgageGet200ResponseData } from 'api-client';

const instance: V1FinanceMortgageGet200ResponseData = {
    principal,
    rate,
    years,
    monthly_payment,
    total_payment,
    total_interest,
    schedule,
    schedule_month,
    schedule_payment,
    schedule_principal,
    schedule_interest,
    schedule_balance,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
