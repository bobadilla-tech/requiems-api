# V1FinanceMortgageGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**principal** | **float** | The original loan amount passed in the request | [optional] 
**rate** | **float** | The annual interest rate passed in the request | [optional] 
**years** | **int** | The loan term in years passed in the request | [optional] 
**monthly_payment** | **float** | Fixed monthly payment amount (rounded to 2 decimal places) | [optional] 
**total_payment** | **float** | Total amount paid over the life of the loan | [optional] 
**total_interest** | **float** | Total interest paid (total_payment minus principal) | [optional] 
**schedule** | **List[object]** | Full amortization schedule — one entry per month (years × 12 entries) | [optional] 
**schedule_month** | **int** | Month number (1 to years × 12) | [optional] 
**schedule_payment** | **float** | Total payment for this month | [optional] 
**schedule_principal** | **float** | Portion of this month&#39;s payment applied to principal | [optional] 
**schedule_interest** | **float** | Portion of this month&#39;s payment applied to interest | [optional] 
**schedule_balance** | **float** | Remaining loan balance after this payment | [optional] 

## Example

```python
from openapi_client.models.v1_finance_mortgage_get200_response_data import V1FinanceMortgageGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1FinanceMortgageGet200ResponseData from a JSON string
v1_finance_mortgage_get200_response_data_instance = V1FinanceMortgageGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1FinanceMortgageGet200ResponseData.to_json())

# convert the object into a dict
v1_finance_mortgage_get200_response_data_dict = v1_finance_mortgage_get200_response_data_instance.to_dict()
# create an instance of V1FinanceMortgageGet200ResponseData from a dict
v1_finance_mortgage_get200_response_data_from_dict = V1FinanceMortgageGet200ResponseData.from_dict(v1_finance_mortgage_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


