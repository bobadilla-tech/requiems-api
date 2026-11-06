# \EmailValidateAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ValidationEmailBatchPost**](EmailValidateAPI.md#V1ValidationEmailBatchPost) | **Post** /v1/validation/email/batch | Validate Emails (Batch)
[**V1ValidationEmailPost**](EmailValidateAPI.md#V1ValidationEmailPost) | **Post** /v1/validation/email | Validate Email



## V1ValidationEmailBatchPost

> V1ValidationEmailBatchPost200Response V1ValidationEmailBatchPost(ctx).V1ValidationEmailBatchPostRequest(v1ValidationEmailBatchPostRequest).Execute()

Validate Emails (Batch)



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
	v1ValidationEmailBatchPostRequest := *openapiclient.NewV1ValidationEmailBatchPostRequest([]interface{}{nil}) // V1ValidationEmailBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailValidateAPI.V1ValidationEmailBatchPost(context.Background()).V1ValidationEmailBatchPostRequest(v1ValidationEmailBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailValidateAPI.V1ValidationEmailBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ValidationEmailBatchPost`: V1ValidationEmailBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailValidateAPI.V1ValidationEmailBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ValidationEmailBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ValidationEmailBatchPostRequest** | [**V1ValidationEmailBatchPostRequest**](V1ValidationEmailBatchPostRequest.md) |  | 

### Return type

[**V1ValidationEmailBatchPost200Response**](V1ValidationEmailBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ValidationEmailPost

> V1ValidationEmailPost200Response V1ValidationEmailPost(ctx).V1ValidationEmailPostRequest(v1ValidationEmailPostRequest).Execute()

Validate Email



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
	v1ValidationEmailPostRequest := *openapiclient.NewV1ValidationEmailPostRequest("user@gmial.com") // V1ValidationEmailPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailValidateAPI.V1ValidationEmailPost(context.Background()).V1ValidationEmailPostRequest(v1ValidationEmailPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailValidateAPI.V1ValidationEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ValidationEmailPost`: V1ValidationEmailPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailValidateAPI.V1ValidationEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ValidationEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ValidationEmailPostRequest** | [**V1ValidationEmailPostRequest**](V1ValidationEmailPostRequest.md) |  | 

### Return type

[**V1ValidationEmailPost200Response**](V1ValidationEmailPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

