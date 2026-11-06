# \ExchangeRateAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceConvertGet**](ExchangeRateAPI.md#V1FinanceConvertGet) | **Get** /v1/finance/convert | Convert Currency
[**V1FinanceExchangeRateGet**](ExchangeRateAPI.md#V1FinanceExchangeRateGet) | **Get** /v1/finance/exchange-rate | Get Exchange Rate



## V1FinanceConvertGet

> V1FinanceConvertGet200Response V1FinanceConvertGet(ctx).From(from).To(to).Amount(amount).Execute()

Convert Currency



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
	from := "USD" // string | ISO 4217 source currency code (3 letters, e.g. USD)
	to := "EUR" // string | ISO 4217 target currency code (3 letters, e.g. EUR)
	amount := float32(100) // float32 | Amount to convert. Must be greater than 0.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangeRateAPI.V1FinanceConvertGet(context.Background()).From(from).To(to).Amount(amount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRateAPI.V1FinanceConvertGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceConvertGet`: V1FinanceConvertGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ExchangeRateAPI.V1FinanceConvertGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceConvertGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | ISO 4217 source currency code (3 letters, e.g. USD) | 
 **to** | **string** | ISO 4217 target currency code (3 letters, e.g. EUR) | 
 **amount** | **float32** | Amount to convert. Must be greater than 0. | 

### Return type

[**V1FinanceConvertGet200Response**](V1FinanceConvertGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FinanceExchangeRateGet

> V1FinanceExchangeRateGet200Response V1FinanceExchangeRateGet(ctx).From(from).To(to).Execute()

Get Exchange Rate



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
	from := "USD" // string | ISO 4217 source currency code (3 letters, e.g. USD)
	to := "EUR" // string | ISO 4217 target currency code (3 letters, e.g. EUR)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangeRateAPI.V1FinanceExchangeRateGet(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeRateAPI.V1FinanceExchangeRateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceExchangeRateGet`: V1FinanceExchangeRateGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ExchangeRateAPI.V1FinanceExchangeRateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceExchangeRateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | ISO 4217 source currency code (3 letters, e.g. USD) | 
 **to** | **string** | ISO 4217 target currency code (3 letters, e.g. EUR) | 

### Return type

[**V1FinanceExchangeRateGet200Response**](V1FinanceExchangeRateGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

