# \IbanAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceIbanBatchPost**](IbanAPI.md#V1FinanceIbanBatchPost) | **Post** /v1/finance/iban/batch | Batch Validate IBANs
[**V1FinanceIbanIbanGet**](IbanAPI.md#V1FinanceIbanIbanGet) | **Get** /v1/finance/iban/{iban} | Validate IBAN



## V1FinanceIbanBatchPost

> V1FinanceIbanBatchPost200Response V1FinanceIbanBatchPost(ctx).V1FinanceIbanBatchPostRequest(v1FinanceIbanBatchPostRequest).Execute()

Batch Validate IBANs



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
	v1FinanceIbanBatchPostRequest := *openapiclient.NewV1FinanceIbanBatchPostRequest([]interface{}{nil}) // V1FinanceIbanBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbanAPI.V1FinanceIbanBatchPost(context.Background()).V1FinanceIbanBatchPostRequest(v1FinanceIbanBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbanAPI.V1FinanceIbanBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceIbanBatchPost`: V1FinanceIbanBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IbanAPI.V1FinanceIbanBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceIbanBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1FinanceIbanBatchPostRequest** | [**V1FinanceIbanBatchPostRequest**](V1FinanceIbanBatchPostRequest.md) |  | 

### Return type

[**V1FinanceIbanBatchPost200Response**](V1FinanceIbanBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FinanceIbanIbanGet

> V1FinanceIbanIbanGet200Response V1FinanceIbanIbanGet(ctx, iban).Execute()

Validate IBAN



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
	iban := "DE89370400440532013000" // string | The IBAN to validate. Spaces are stripped. Case-insensitive.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IbanAPI.V1FinanceIbanIbanGet(context.Background(), iban).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IbanAPI.V1FinanceIbanIbanGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceIbanIbanGet`: V1FinanceIbanIbanGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IbanAPI.V1FinanceIbanIbanGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iban** | **string** | The IBAN to validate. Spaces are stripped. Case-insensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceIbanIbanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1FinanceIbanIbanGet200Response**](V1FinanceIbanIbanGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

