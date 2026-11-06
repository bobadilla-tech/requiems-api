# \BinLookupAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceBinBinGet**](BinLookupAPI.md#V1FinanceBinBinGet) | **Get** /v1/finance/bin/{bin} | BIN Lookup



## V1FinanceBinBinGet

> V1FinanceBinBinGet200Response V1FinanceBinBinGet(ctx, bin).Execute()

BIN Lookup



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
	bin := "424242" // string | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinLookupAPI.V1FinanceBinBinGet(context.Background(), bin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinLookupAPI.V1FinanceBinBinGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceBinBinGet`: V1FinanceBinBinGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BinLookupAPI.V1FinanceBinBinGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bin** | **string** | 6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceBinBinGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1FinanceBinBinGet200Response**](V1FinanceBinBinGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

