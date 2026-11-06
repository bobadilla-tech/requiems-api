# Org.OpenAPITools.Model.V1FinanceMortgageGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | **decimal** | The original loan amount passed in the request | [optional] 
**Rate** | **decimal** | The annual interest rate passed in the request | [optional] 
**Years** | **int** | The loan term in years passed in the request | [optional] 
**MonthlyPayment** | **decimal** | Fixed monthly payment amount (rounded to 2 decimal places) | [optional] 
**TotalPayment** | **decimal** | Total amount paid over the life of the loan | [optional] 
**TotalInterest** | **decimal** | Total interest paid (total_payment minus principal) | [optional] 
**Schedule** | **List&lt;Object&gt;** | Full amortization schedule — one entry per month (years × 12 entries) | [optional] 
**ScheduleMonth** | **int** | Month number (1 to years × 12) | [optional] 
**SchedulePayment** | **decimal** | Total payment for this month | [optional] 
**SchedulePrincipal** | **decimal** | Portion of this month&#39;s payment applied to principal | [optional] 
**ScheduleInterest** | **decimal** | Portion of this month&#39;s payment applied to interest | [optional] 
**ScheduleBalance** | **decimal** | Remaining loan balance after this payment | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

