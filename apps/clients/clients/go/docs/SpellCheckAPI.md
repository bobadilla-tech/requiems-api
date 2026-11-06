# \SpellCheckAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextSpellcheckPost**](SpellCheckAPI.md#V1TextSpellcheckPost) | **Post** /v1/text/spellcheck | Check Spelling



## V1TextSpellcheckPost

> V1TextSpellcheckPost200Response V1TextSpellcheckPost(ctx).V1TextSpellcheckPostRequest(v1TextSpellcheckPostRequest).Execute()

Check Spelling



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
	v1TextSpellcheckPostRequest := *openapiclient.NewV1TextSpellcheckPostRequest("Ths is a smiple tset") // V1TextSpellcheckPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpellCheckAPI.V1TextSpellcheckPost(context.Background()).V1TextSpellcheckPostRequest(v1TextSpellcheckPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpellCheckAPI.V1TextSpellcheckPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextSpellcheckPost`: V1TextSpellcheckPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SpellCheckAPI.V1TextSpellcheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextSpellcheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextSpellcheckPostRequest** | [**V1TextSpellcheckPostRequest**](V1TextSpellcheckPostRequest.md) |  | 

### Return type

[**V1TextSpellcheckPost200Response**](V1TextSpellcheckPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

