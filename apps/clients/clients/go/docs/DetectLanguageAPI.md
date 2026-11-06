# \DetectLanguageAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextDetectLanguagePost**](DetectLanguageAPI.md#V1TextDetectLanguagePost) | **Post** /v1/text/detect-language | Detect Language



## V1TextDetectLanguagePost

> V1TextDetectLanguagePost200Response V1TextDetectLanguagePost(ctx).V1TextDetectLanguagePostRequest(v1TextDetectLanguagePostRequest).Execute()

Detect Language



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
	v1TextDetectLanguagePostRequest := *openapiclient.NewV1TextDetectLanguagePostRequest("Bonjour, comment ça va?") // V1TextDetectLanguagePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DetectLanguageAPI.V1TextDetectLanguagePost(context.Background()).V1TextDetectLanguagePostRequest(v1TextDetectLanguagePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DetectLanguageAPI.V1TextDetectLanguagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextDetectLanguagePost`: V1TextDetectLanguagePost200Response
	fmt.Fprintf(os.Stdout, "Response from `DetectLanguageAPI.V1TextDetectLanguagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextDetectLanguagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextDetectLanguagePostRequest** | [**V1TextDetectLanguagePostRequest**](V1TextDetectLanguagePostRequest.md) |  | 

### Return type

[**V1TextDetectLanguagePost200Response**](V1TextDetectLanguagePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

