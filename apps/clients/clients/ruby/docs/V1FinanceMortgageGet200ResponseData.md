# OpenapiClient::V1FinanceMortgageGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **principal** | **Float** | The original loan amount passed in the request | [optional] |
| **rate** | **Float** | The annual interest rate passed in the request | [optional] |
| **years** | **Integer** | The loan term in years passed in the request | [optional] |
| **monthly_payment** | **Float** | Fixed monthly payment amount (rounded to 2 decimal places) | [optional] |
| **total_payment** | **Float** | Total amount paid over the life of the loan | [optional] |
| **total_interest** | **Float** | Total interest paid (total_payment minus principal) | [optional] |
| **schedule** | **Array&lt;Object&gt;** | Full amortization schedule — one entry per month (years × 12 entries) | [optional] |
| **schedule_month** | **Integer** | Month number (1 to years × 12) | [optional] |
| **schedule_payment** | **Float** | Total payment for this month | [optional] |
| **schedule_principal** | **Float** | Portion of this month&#39;s payment applied to principal | [optional] |
| **schedule_interest** | **Float** | Portion of this month&#39;s payment applied to interest | [optional] |
| **schedule_balance** | **Float** | Remaining loan balance after this payment | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceMortgageGet200ResponseData.new(
  principal: null,
  rate: null,
  years: null,
  monthly_payment: null,
  total_payment: null,
  total_interest: null,
  schedule: null,
  schedule_month: null,
  schedule_payment: null,
  schedule_principal: null,
  schedule_interest: null,
  schedule_balance: null
)
```

