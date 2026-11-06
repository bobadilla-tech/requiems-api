# \SwiftCodeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceSwiftCodeGet**](SwiftCodeAPI.md#V1FinanceSwiftCodeGet) | **Get** /v1/finance/swift/{code} | Get SWIFT Code
[**V1FinanceSwiftGet**](SwiftCodeAPI.md#V1FinanceSwiftGet) | **Get** /v1/finance/swift | List SWIFT Codes



## V1FinanceSwiftCodeGet

> V1FinanceSwiftCodeGet200Response V1FinanceSwiftCodeGet(ctx, code).Execute()

Get SWIFT Code



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
	code := "DEUTDEDB" // string | SWIFT/BIC code (8 or 11 alphanumeric characters)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwiftCodeAPI.V1FinanceSwiftCodeGet(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwiftCodeAPI.V1FinanceSwiftCodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceSwiftCodeGet`: V1FinanceSwiftCodeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SwiftCodeAPI.V1FinanceSwiftCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | SWIFT/BIC code (8 or 11 alphanumeric characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceSwiftCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1FinanceSwiftCodeGet200Response**](V1FinanceSwiftCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FinanceSwiftGet

> V1HealthExercisesRandomGet200Response V1FinanceSwiftGet(ctx).CountryCode(countryCode).BankCode(bankCode).Q(q).Limit(limit).Offset(offset).Execute()

List SWIFT Codes



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
	countryCode := "DE" // string | Optional 2-letter country code filter (e.g. DE, US) (optional)
	bankCode := "DEUT" // string | Optional 4-letter bank code filter (e.g. DEUT) (optional)
	q := "deutsche" // string | Optional text search across swift_code, bank_name, and city (optional)
	limit := int32(50) // int32 | Max rows to return (default 50, max 200) (optional)
	offset := int32(0) // int32 | Number of rows to skip (default 0) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwiftCodeAPI.V1FinanceSwiftGet(context.Background()).CountryCode(countryCode).BankCode(bankCode).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwiftCodeAPI.V1FinanceSwiftGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceSwiftGet`: V1HealthExercisesRandomGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SwiftCodeAPI.V1FinanceSwiftGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceSwiftGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | Optional 2-letter country code filter (e.g. DE, US) | 
 **bankCode** | **string** | Optional 4-letter bank code filter (e.g. DEUT) | 
 **q** | **string** | Optional text search across swift_code, bank_name, and city | 
 **limit** | **int32** | Max rows to return (default 50, max 200) | 
 **offset** | **int32** | Number of rows to skip (default 0) | 

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

