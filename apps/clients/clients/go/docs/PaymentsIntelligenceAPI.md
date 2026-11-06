# \PaymentsIntelligenceAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SystemsPaymentValidatePost**](PaymentsIntelligenceAPI.md#V1SystemsPaymentValidatePost) | **Post** /v1/systems/payment/validate | Validate Payment
[**V1SystemsTransactionRiskPost**](PaymentsIntelligenceAPI.md#V1SystemsTransactionRiskPost) | **Post** /v1/systems/transaction/risk | Score Transaction Risk



## V1SystemsPaymentValidatePost

> V1SystemsPaymentValidatePost200Response V1SystemsPaymentValidatePost(ctx).V1SystemsPaymentValidatePostRequest(v1SystemsPaymentValidatePostRequest).Execute()

Validate Payment



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
	v1SystemsPaymentValidatePostRequest := *openapiclient.NewV1SystemsPaymentValidatePostRequest() // V1SystemsPaymentValidatePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsIntelligenceAPI.V1SystemsPaymentValidatePost(context.Background()).V1SystemsPaymentValidatePostRequest(v1SystemsPaymentValidatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsIntelligenceAPI.V1SystemsPaymentValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsPaymentValidatePost`: V1SystemsPaymentValidatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentsIntelligenceAPI.V1SystemsPaymentValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsPaymentValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsPaymentValidatePostRequest** | [**V1SystemsPaymentValidatePostRequest**](V1SystemsPaymentValidatePostRequest.md) |  | 

### Return type

[**V1SystemsPaymentValidatePost200Response**](V1SystemsPaymentValidatePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SystemsTransactionRiskPost

> V1SystemsTransactionRiskPost200Response V1SystemsTransactionRiskPost(ctx).V1SystemsTransactionRiskPostRequest(v1SystemsTransactionRiskPostRequest).Execute()

Score Transaction Risk



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
	v1SystemsTransactionRiskPostRequest := *openapiclient.NewV1SystemsTransactionRiskPostRequest("424242", "92.168.1.1") // V1SystemsTransactionRiskPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsIntelligenceAPI.V1SystemsTransactionRiskPost(context.Background()).V1SystemsTransactionRiskPostRequest(v1SystemsTransactionRiskPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsIntelligenceAPI.V1SystemsTransactionRiskPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsTransactionRiskPost`: V1SystemsTransactionRiskPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentsIntelligenceAPI.V1SystemsTransactionRiskPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsTransactionRiskPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsTransactionRiskPostRequest** | [**V1SystemsTransactionRiskPostRequest**](V1SystemsTransactionRiskPostRequest.md) |  | 

### Return type

[**V1SystemsTransactionRiskPost200Response**](V1SystemsTransactionRiskPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

