# \MortgageAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceMortgageGet**](MortgageAPI.md#V1FinanceMortgageGet) | **Get** /v1/finance/mortgage | Calculate Mortgage



## V1FinanceMortgageGet

> V1FinanceMortgageGet200Response V1FinanceMortgageGet(ctx).Principal(principal).Rate(rate).Years(years).Execute()

Calculate Mortgage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	principal := float32(300000) // float32 | Loan amount in your chosen currency (e.g. 300000 for $300,000)
	rate := float32(6.5) // float32 | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0.
	years := int32(30) // int32 | Loan term in years (1–50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MortgageAPI.V1FinanceMortgageGet(context.Background()).Principal(principal).Rate(rate).Years(years).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MortgageAPI.V1FinanceMortgageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceMortgageGet`: V1FinanceMortgageGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MortgageAPI.V1FinanceMortgageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceMortgageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principal** | **float32** | Loan amount in your chosen currency (e.g. 300000 for $300,000) | 
 **rate** | **float32** | Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0. | 
 **years** | **int32** | Loan term in years (1–50) | 

### Return type

[**V1FinanceMortgageGet200Response**](V1FinanceMortgageGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

