# \TextSimilarityAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextSimilarityPost**](TextSimilarityAPI.md#V1TextSimilarityPost) | **Post** /v1/text/similarity | Compare Text Similarity



## V1TextSimilarityPost

> V1TextSimilarityPost200Response V1TextSimilarityPost(ctx).V1TextSimilarityPostRequest(v1TextSimilarityPostRequest).Execute()

Compare Text Similarity



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
	v1TextSimilarityPostRequest := *openapiclient.NewV1TextSimilarityPostRequest("The cat sat on the mat", "A cat was sitting on a mat") // V1TextSimilarityPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextSimilarityAPI.V1TextSimilarityPost(context.Background()).V1TextSimilarityPostRequest(v1TextSimilarityPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextSimilarityAPI.V1TextSimilarityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextSimilarityPost`: V1TextSimilarityPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TextSimilarityAPI.V1TextSimilarityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextSimilarityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextSimilarityPostRequest** | [**V1TextSimilarityPostRequest**](V1TextSimilarityPostRequest.md) |  | 

### Return type

[**V1TextSimilarityPost200Response**](V1TextSimilarityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

