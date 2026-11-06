# \InflationAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceInflationBatchPost**](InflationAPI.md#V1FinanceInflationBatchPost) | **Post** /v1/finance/inflation/batch | Batch Inflation Rates
[**V1FinanceInflationGet**](InflationAPI.md#V1FinanceInflationGet) | **Get** /v1/finance/inflation | Get Inflation Rate



## V1FinanceInflationBatchPost

> V1FinanceInflationBatchPost200Response V1FinanceInflationBatchPost(ctx).V1FinanceInflationBatchPostRequest(v1FinanceInflationBatchPostRequest).Execute()

Batch Inflation Rates



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
	v1FinanceInflationBatchPostRequest := *openapiclient.NewV1FinanceInflationBatchPostRequest([]interface{}{nil}) // V1FinanceInflationBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InflationAPI.V1FinanceInflationBatchPost(context.Background()).V1FinanceInflationBatchPostRequest(v1FinanceInflationBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InflationAPI.V1FinanceInflationBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceInflationBatchPost`: V1FinanceInflationBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InflationAPI.V1FinanceInflationBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceInflationBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1FinanceInflationBatchPostRequest** | [**V1FinanceInflationBatchPostRequest**](V1FinanceInflationBatchPostRequest.md) |  | 

### Return type

[**V1FinanceInflationBatchPost200Response**](V1FinanceInflationBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FinanceInflationGet

> V1FinanceInflationGet200Response V1FinanceInflationGet(ctx).Country(country).Execute()

Get Inflation Rate



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
	country := "US" // string | ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InflationAPI.V1FinanceInflationGet(context.Background()).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InflationAPI.V1FinanceInflationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceInflationGet`: V1FinanceInflationGet200Response
	fmt.Fprintf(os.Stdout, "Response from `InflationAPI.V1FinanceInflationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceInflationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive. | 

### Return type

[**V1FinanceInflationGet200Response**](V1FinanceInflationGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

