# \ProfanityAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextProfanityPost**](ProfanityAPI.md#V1TextProfanityPost) | **Post** /v1/text/profanity | Check Profanity



## V1TextProfanityPost

> V1TextProfanityPost200Response V1TextProfanityPost(ctx).V1TextProfanityPostRequest(v1TextProfanityPostRequest).Execute()

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
	v1TextProfanityPostRequest := *openapiclient.NewV1TextProfanityPostRequest("What the heck is going on") // V1TextProfanityPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfanityAPI.V1TextProfanityPost(context.Background()).V1TextProfanityPostRequest(v1TextProfanityPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfanityAPI.V1TextProfanityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextProfanityPost`: V1TextProfanityPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ProfanityAPI.V1TextProfanityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextProfanityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextProfanityPostRequest** | [**V1TextProfanityPostRequest**](V1TextProfanityPostRequest.md) |  | 

### Return type

[**V1TextProfanityPost200Response**](V1TextProfanityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

