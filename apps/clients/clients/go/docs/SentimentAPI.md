# \SentimentAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextSentimentBatchPost**](SentimentAPI.md#V1TextSentimentBatchPost) | **Post** /v1/text/sentiment/batch | Analyze Sentiment (Batch)
[**V1TextSentimentPost**](SentimentAPI.md#V1TextSentimentPost) | **Post** /v1/text/sentiment | Analyze Sentiment



## V1TextSentimentBatchPost

> V1TextSentimentBatchPost200Response V1TextSentimentBatchPost(ctx).V1TextSentimentBatchPostRequest(v1TextSentimentBatchPostRequest).Execute()

Analyze Sentiment (Batch)



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
	v1TextSentimentBatchPostRequest := *openapiclient.NewV1TextSentimentBatchPostRequest([]interface{}{nil}) // V1TextSentimentBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentimentAPI.V1TextSentimentBatchPost(context.Background()).V1TextSentimentBatchPostRequest(v1TextSentimentBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentimentAPI.V1TextSentimentBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextSentimentBatchPost`: V1TextSentimentBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SentimentAPI.V1TextSentimentBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextSentimentBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextSentimentBatchPostRequest** | [**V1TextSentimentBatchPostRequest**](V1TextSentimentBatchPostRequest.md) |  | 

### Return type

[**V1TextSentimentBatchPost200Response**](V1TextSentimentBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TextSentimentPost

> V1TextSentimentPost200Response V1TextSentimentPost(ctx).V1TextSentimentPostRequest(v1TextSentimentPostRequest).Execute()

Analyze Sentiment



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
	v1TextSentimentPostRequest := *openapiclient.NewV1TextSentimentPostRequest("I absolutely love this product, it exceeded my expectations!") // V1TextSentimentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SentimentAPI.V1TextSentimentPost(context.Background()).V1TextSentimentPostRequest(v1TextSentimentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SentimentAPI.V1TextSentimentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextSentimentPost`: V1TextSentimentPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SentimentAPI.V1TextSentimentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextSentimentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextSentimentPostRequest** | [**V1TextSentimentPostRequest**](V1TextSentimentPostRequest.md) |  | 

### Return type

[**V1TextSentimentPost200Response**](V1TextSentimentPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

