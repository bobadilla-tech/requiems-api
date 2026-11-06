# \ProfanityAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ValidationProfanityBatchPost**](ProfanityAPI.md#V1ValidationProfanityBatchPost) | **Post** /v1/validation/profanity/batch | Batch Check Profanity
[**V1ValidationProfanityPost**](ProfanityAPI.md#V1ValidationProfanityPost) | **Post** /v1/validation/profanity | Check Profanity



## V1ValidationProfanityBatchPost

> V1ValidationProfanityBatchPost200Response V1ValidationProfanityBatchPost(ctx).V1ValidationProfanityBatchPostRequest(v1ValidationProfanityBatchPostRequest).Execute()

Batch Check Profanity



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
	v1ValidationProfanityBatchPostRequest := *openapiclient.NewV1ValidationProfanityBatchPostRequest("Texts_example") // V1ValidationProfanityBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfanityAPI.V1ValidationProfanityBatchPost(context.Background()).V1ValidationProfanityBatchPostRequest(v1ValidationProfanityBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfanityAPI.V1ValidationProfanityBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ValidationProfanityBatchPost`: V1ValidationProfanityBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ProfanityAPI.V1ValidationProfanityBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ValidationProfanityBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ValidationProfanityBatchPostRequest** | [**V1ValidationProfanityBatchPostRequest**](V1ValidationProfanityBatchPostRequest.md) |  | 

### Return type

[**V1ValidationProfanityBatchPost200Response**](V1ValidationProfanityBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ValidationProfanityPost

> V1ValidationProfanityPost200Response V1ValidationProfanityPost(ctx).V1ValidationProfanityPostRequest(v1ValidationProfanityPostRequest).Execute()

Check Profanity



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
	v1ValidationProfanityPostRequest := *openapiclient.NewV1ValidationProfanityPostRequest("What the heck is going on") // V1ValidationProfanityPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfanityAPI.V1ValidationProfanityPost(context.Background()).V1ValidationProfanityPostRequest(v1ValidationProfanityPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfanityAPI.V1ValidationProfanityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ValidationProfanityPost`: V1ValidationProfanityPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ProfanityAPI.V1ValidationProfanityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ValidationProfanityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ValidationProfanityPostRequest** | [**V1ValidationProfanityPostRequest**](V1ValidationProfanityPostRequest.md) |  | 

### Return type

[**V1ValidationProfanityPost200Response**](V1ValidationProfanityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

