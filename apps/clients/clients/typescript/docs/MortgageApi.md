# MortgageApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1FinanceMortgageGet**](#v1financemortgageget) | **GET** /v1/finance/mortgage | Calculate Mortgage|

# **v1FinanceMortgageGet**
> V1FinanceMortgageGet200Response v1FinanceMortgageGet()

Returns the monthly payment, total cost, and full amortization schedule for a fixed-rate mortgage.

### Example

```typescript
import {
    MortgageApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new MortgageApi(configuration);

let principal: number; //Loan amount in your chosen currency (e.g. 300000 for $300,000) (default to undefined)
let rate: number; //Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0. (default to undefined)
let years: number; //Loan term in years (1–50) (default to undefined)

const { status, data } = await apiInstance.v1FinanceMortgageGet(
    principal,
    rate,
    years
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **principal** | [**number**] | Loan amount in your chosen currency (e.g. 300000 for $300,000) | defaults to undefined|
| **rate** | [**number**] | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0. | defaults to undefined|
| **years** | [**number**] | Loan term in years (1–50) | defaults to undefined|


### Return type

**V1FinanceMortgageGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | A required parameter is missing, not a valid number, or out of range (e.g. years &gt; 50 or rate &lt;&#x3D; 0). |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

