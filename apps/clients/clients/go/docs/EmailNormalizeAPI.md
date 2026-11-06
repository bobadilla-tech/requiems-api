# \EmailNormalizeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextNormalizeBatchPost**](EmailNormalizeAPI.md#V1TextNormalizeBatchPost) | **Post** /v1/text/normalize/batch | Normalize Email Batch
[**V1TextNormalizePost**](EmailNormalizeAPI.md#V1TextNormalizePost) | **Post** /v1/text/normalize | Normalize Email



## V1TextNormalizeBatchPost

> V1TextNormalizeBatchPost200Response V1TextNormalizeBatchPost(ctx).V1TextNormalizeBatchPostRequest(v1TextNormalizeBatchPostRequest).Execute()

Normalize Email Batch



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
	v1TextNormalizeBatchPostRequest := *openapiclient.NewV1TextNormalizeBatchPostRequest([]interface{}{nil}) // V1TextNormalizeBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailNormalizeAPI.V1TextNormalizeBatchPost(context.Background()).V1TextNormalizeBatchPostRequest(v1TextNormalizeBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailNormalizeAPI.V1TextNormalizeBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextNormalizeBatchPost`: V1TextNormalizeBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailNormalizeAPI.V1TextNormalizeBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextNormalizeBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextNormalizeBatchPostRequest** | [**V1TextNormalizeBatchPostRequest**](V1TextNormalizeBatchPostRequest.md) |  | 

### Return type

[**V1TextNormalizeBatchPost200Response**](V1TextNormalizeBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TextNormalizePost

> V1TextNormalizePost200Response V1TextNormalizePost(ctx).V1TextNormalizePostRequest(v1TextNormalizePostRequest).Execute()

Normalize Email



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
	v1TextNormalizePostRequest := *openapiclient.NewV1TextNormalizePostRequest("Te.st.User+spam@Googlemail.com") // V1TextNormalizePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailNormalizeAPI.V1TextNormalizePost(context.Background()).V1TextNormalizePostRequest(v1TextNormalizePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailNormalizeAPI.V1TextNormalizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextNormalizePost`: V1TextNormalizePost200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailNormalizeAPI.V1TextNormalizePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextNormalizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextNormalizePostRequest** | [**V1TextNormalizePostRequest**](V1TextNormalizePostRequest.md) |  | 

### Return type

[**V1TextNormalizePost200Response**](V1TextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

