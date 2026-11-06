# \CryptoAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceCryptoSymbolGet**](CryptoAPI.md#V1FinanceCryptoSymbolGet) | **Get** /v1/finance/crypto/{symbol} | Get Crypto Price



## V1FinanceCryptoSymbolGet

> V1FinanceCryptoSymbolGet200Response V1FinanceCryptoSymbolGet(ctx, symbol).Execute()

Get Crypto Price



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
	symbol := "BTC" // string | Uppercase ticker symbol (e.g. BTC, ETH, SOL)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoAPI.V1FinanceCryptoSymbolGet(context.Background(), symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoAPI.V1FinanceCryptoSymbolGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceCryptoSymbolGet`: V1FinanceCryptoSymbolGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CryptoAPI.V1FinanceCryptoSymbolGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** | Uppercase ticker symbol (e.g. BTC, ETH, SOL) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceCryptoSymbolGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1FinanceCryptoSymbolGet200Response**](V1FinanceCryptoSymbolGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

