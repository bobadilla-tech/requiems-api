# \LoremIpsumAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextLoremGet**](LoremIpsumAPI.md#V1TextLoremGet) | **Get** /v1/text/lorem | Generate Lorem Ipsum



## V1TextLoremGet

> V1TextLoremGet200Response V1TextLoremGet(ctx).Paragraphs(paragraphs).Sentences(sentences).Execute()

Generate Lorem Ipsum



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
	paragraphs := int32(3) // int32 | Number of paragraphs to generate (1-20) (optional)
	sentences := int32(5) // int32 | Number of sentences per paragraph (1-20) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoremIpsumAPI.V1TextLoremGet(context.Background()).Paragraphs(paragraphs).Sentences(sentences).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoremIpsumAPI.V1TextLoremGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextLoremGet`: V1TextLoremGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LoremIpsumAPI.V1TextLoremGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextLoremGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paragraphs** | **int32** | Number of paragraphs to generate (1-20) | 
 **sentences** | **int32** | Number of sentences per paragraph (1-20) | 

### Return type

[**V1TextLoremGet200Response**](V1TextLoremGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

