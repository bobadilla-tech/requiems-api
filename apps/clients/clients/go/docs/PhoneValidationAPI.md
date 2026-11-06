# \PhoneValidationAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ValidationPhoneBatchPost**](PhoneValidationAPI.md#V1ValidationPhoneBatchPost) | **Post** /v1/validation/phone/batch | Batch Validate Phone Numbers
[**V1ValidationPhoneGet**](PhoneValidationAPI.md#V1ValidationPhoneGet) | **Get** /v1/validation/phone | Validate Phone Number



## V1ValidationPhoneBatchPost

> V1FinanceIbanBatchPost200Response V1ValidationPhoneBatchPost(ctx).V1ValidationPhoneBatchPostRequest(v1ValidationPhoneBatchPostRequest).Execute()

Batch Validate Phone Numbers



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
	v1ValidationPhoneBatchPostRequest := *openapiclient.NewV1ValidationPhoneBatchPostRequest([]interface{}{nil}) // V1ValidationPhoneBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneValidationAPI.V1ValidationPhoneBatchPost(context.Background()).V1ValidationPhoneBatchPostRequest(v1ValidationPhoneBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneValidationAPI.V1ValidationPhoneBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ValidationPhoneBatchPost`: V1FinanceIbanBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneValidationAPI.V1ValidationPhoneBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ValidationPhoneBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ValidationPhoneBatchPostRequest** | [**V1ValidationPhoneBatchPostRequest**](V1ValidationPhoneBatchPostRequest.md) |  | 

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


## V1ValidationPhoneGet

> V1ValidationPhoneGet200Response V1ValidationPhoneGet(ctx).Number(number).Execute()

Validate Phone Number



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
	number := "+447400123456" // string | The phone number to validate. Must include the country calling code (e.g. +12015551234).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneValidationAPI.V1ValidationPhoneGet(context.Background()).Number(number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneValidationAPI.V1ValidationPhoneGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ValidationPhoneGet`: V1ValidationPhoneGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneValidationAPI.V1ValidationPhoneGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ValidationPhoneGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **number** | **string** | The phone number to validate. Must include the country calling code (e.g. +12015551234). | 

### Return type

[**V1ValidationPhoneGet200Response**](V1ValidationPhoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

