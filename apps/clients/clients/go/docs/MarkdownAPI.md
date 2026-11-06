# \MarkdownAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyMarkdownBatchPost**](MarkdownAPI.md#V1TechnologyMarkdownBatchPost) | **Post** /v1/technology/markdown/batch | Convert Markdown Batch
[**V1TechnologyMarkdownPost**](MarkdownAPI.md#V1TechnologyMarkdownPost) | **Post** /v1/technology/markdown | Convert Markdown to HTML



## V1TechnologyMarkdownBatchPost

> V1TechnologyMarkdownBatchPost200Response V1TechnologyMarkdownBatchPost(ctx).V1TechnologyMarkdownBatchPostRequest(v1TechnologyMarkdownBatchPostRequest).Execute()

Convert Markdown Batch



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
	v1TechnologyMarkdownBatchPostRequest := *openapiclient.NewV1TechnologyMarkdownBatchPostRequest([]interface{}{nil}) // V1TechnologyMarkdownBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarkdownAPI.V1TechnologyMarkdownBatchPost(context.Background()).V1TechnologyMarkdownBatchPostRequest(v1TechnologyMarkdownBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarkdownAPI.V1TechnologyMarkdownBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyMarkdownBatchPost`: V1TechnologyMarkdownBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MarkdownAPI.V1TechnologyMarkdownBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyMarkdownBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TechnologyMarkdownBatchPostRequest** | [**V1TechnologyMarkdownBatchPostRequest**](V1TechnologyMarkdownBatchPostRequest.md) |  | 

### Return type

[**V1TechnologyMarkdownBatchPost200Response**](V1TechnologyMarkdownBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyMarkdownPost

> V1TechnologyMarkdownPost200Response V1TechnologyMarkdownPost(ctx).V1TechnologyMarkdownPostRequest(v1TechnologyMarkdownPostRequest).Execute()

Convert Markdown to HTML



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
	v1TechnologyMarkdownPostRequest := *openapiclient.NewV1TechnologyMarkdownPostRequest("# Hello  This is **bold** and _italic_ text.") // V1TechnologyMarkdownPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarkdownAPI.V1TechnologyMarkdownPost(context.Background()).V1TechnologyMarkdownPostRequest(v1TechnologyMarkdownPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarkdownAPI.V1TechnologyMarkdownPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyMarkdownPost`: V1TechnologyMarkdownPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MarkdownAPI.V1TechnologyMarkdownPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyMarkdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TechnologyMarkdownPostRequest** | [**V1TechnologyMarkdownPostRequest**](V1TechnologyMarkdownPostRequest.md) |  | 

### Return type

[**V1TechnologyMarkdownPost200Response**](V1TechnologyMarkdownPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

